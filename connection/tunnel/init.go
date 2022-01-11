// tunnel provides functionalities to
// - create a ssh net.connection for further usage
// - create a timer based establishing of a mqtt connection using net.connection
// - provides a possibility to publish messages to the ssh host
package tunnel

import (
	"io"
	"io/ioutil"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/config"
	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/connection/mqtt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ssh"
)

// SSHCredential contains data for connect to an remote ssh
type SSHCredential struct {
	Protocol string        // The protocol to be used; usually tcp
	Address  string        // The address to connect to
	User     string        // User to use when connecting
	Key      ssh.Signer    // User private key
	HostKey  ssh.PublicKey // PublicKey of the host
}

// MQTT contains data for the MQTT server behind the tunnel
type MQTT struct {
	Address  string // Address of the MQTT server
	Protocol string // Protocol of the address (TCP, UNIX)
}

// Credential contains needed information when a tunneled connection is requested
type Credential struct {
	ID   string         // Internally used identifier
	MQTT MQTT           // MQTT credentials
	SSH  *SSHCredential // Credentials to establish SSH connection
	Wait time.Duration  // The duration to keep the connection open, when 0 a default will be used
}
type reader struct {
	lookup *sync.Map
}

func (r *reader) Get(ID string) []Credential {
	var results []Credential
	if ID == "" {
		r.lookup.Range(func(key, value interface{}) bool {
			if v, ok := value.(Credential); ok {
				results = append(results, v)
			}
			// we want all values
			return true
		})
	} else {
		if v, ok := r.lookup.Load(ID); ok {
			if s, ok := v.(Credential); ok {
				results = append(results, s)
			}
		}
	}
	return results
}

type CredentialsLookup func(string) []Credential

// DefaultSSHCredentials creates SSHCredentials with tcp and user gvm
func DefaultSSHCredentials(address string, prvUserKey, hostKey []byte) (*SSHCredential, error) {
	prv, err := ssh.ParsePrivateKey(prvUserKey)
	if err != nil {
		return nil, err
	}
	pk, _, _, _, err := ssh.ParseAuthorizedKey(hostKey)
	if err != nil {
		return nil, err
	}
	return &SSHCredential{
		Protocol: "tcp",
		Address:  address,
		User:     "gvm",
		Key:      prv,
		HostKey:  pk,
	}, nil
}

func InMemoryLookupCreater(c []config.SSHSensor) (CredentialsLookup, error) {
	l := sync.Map{}

	for _, ssc := range c {

		// meine Seele brennt im Autokino! Move to SSHCredentials function
		uk, err := ioutil.ReadFile(ssc.User.KeyPath)
		if err != nil {
			return nil, err
		}
		hpk, err := ioutil.ReadFile(ssc.Host.PublicKeyPath)
		if err != nil {
			return nil, err
		}

		c, err := DefaultSSHCredentials(ssc.Host.Address, uk, hpk)
		if err != nil {
			return nil, err
		}
		if ssc.User.Name != "" {
			c.User = ssc.User.Name
		}
		if ssc.Host.Protocol != "" {
			c.Protocol = ssc.Host.Protocol
		}
		mqtt := MQTT{
			Address:  "localhost:1883",
			Protocol: "tcp",
		}

		l.Store(ssc.ID, Credential{
			ID:   ssc.ID,
			MQTT: mqtt,
			SSH:  c,
		})
	}
	r := &reader{lookup: &l}
	return r.Get, nil
}

type Connecter func(mqtt.Configuration, Credential) (io.Closer, connection.PubSub, error)

func SSHConnecter(mqttc mqtt.Configuration, c Credential) (client io.Closer, ps connection.PubSub, err error) {
	config := &ssh.ClientConfig{
		User:            c.SSH.User,
		HostKeyCallback: ssh.FixedHostKey(c.SSH.HostKey),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(c.SSH.Key),
		},
	}
	ssh_client, err := ssh.Dial(c.SSH.Protocol, c.SSH.Address, config)
	if err != nil {
		log.Warn().Err(err).Msgf("Unable to dial %s:%s", c.SSH.Address, c.SSH.Protocol)
		return
	}
	conn, err := ssh_client.Dial(c.MQTT.Protocol, c.MQTT.Address)
	if err != nil {
		log.Warn().Err(err).Msgf("Unable to connect to mqtt %+v", c.MQTT)
		ssh_client.Close()
		return
	}
	ps, err = mqtt.New(conn, mqttc)
	client = ssh_client
	err = ps.Connect()
	return
}

type pubsub struct {
	sync.RWMutex
	connecter Connecter
	lookup    CredentialsLookup
	config    mqtt.Configuration
	topics    []string
	connected bool
}

func (p *pubsub) blockUntilConnectionFalse() {
	slept := 0
	for p.connected {
		slept = slept + 1
		time.Sleep(1 * time.Second)
		log.Trace().Msgf("Slept for %d seconds", slept)
	}
}

func (p *pubsub) publishSingle(c Credential, topic string, message interface{}) error {
	config := p.config
	config.ClientID = p.config.ClientID + "Pub" + uuid.NewString()
	sshc, ps, err := p.connecter(config, c)
	if err != nil {
		return err
	}
	defer func() {
		sshc.Close()
		ps.Close()
	}()
	log.Trace().Msgf("Sending %+T to %s in %s", message, topic, c.SSH.Address)
	return ps.Publish(topic, message)

}

func (p *pubsub) Publish(topic string, message interface{}) error {
	p.Lock()
	defer p.Unlock()
	var wg sync.WaitGroup
	for _, c := range p.lookup("") {
		go func(c Credential) {
			wg.Add(1)
			defer wg.Done()
			if err := p.publishSingle(c, topic, message); err != nil {
				log.Warn().Err(err).Msgf("Unable to publish %+T on %s to %s (%s->%s://%s)",
					message,
					topic,
					c.ID,
					c.SSH.Address,
					c.MQTT.Protocol,
					c.MQTT.Address,
				)
			}
		}(c)

	}
	wg.Wait()
	return nil
}

func (p *pubsub) In() <-chan *connection.TopicData {
	return p.config.In
}

func (p *pubsub) Subscribe(topics []string) error {
	p.Lock()
	defer p.Unlock()
	p.topics = topics
	return nil
}
func (p *pubsub) Connect() error {
	return nil
}

func (p *pubsub) Close() error {
	return nil
}

// Receive is opening a connection to ssh to the broker on there and wait for duration to receive messages
func (p *pubsub) Receive(wait time.Duration) error {
	p.Lock()
	defer p.Unlock()
	var wg sync.WaitGroup
	for _, c := range p.lookup("") {
		go func(wg *sync.WaitGroup, c Credential, t []string, d time.Duration) {
			wg.Add(1)
			defer wg.Done()
			p.blockUntilConnectionFalse()
			p.connected = true
			sshc, ps, err := p.connecter(p.config, c)
			if err != nil {
				log.Warn().Err(err).Msgf("Unable to establish connection to %s", c.ID)
				return
			}
			defer func() {
				ps.Close()
				sshc.Close()
				p.connected = false
			}()
			if err = ps.Subscribe(t); err != nil {
				log.Warn().Err(err).Msgf("Unable to subscribe to %s", c.ID)
				return
			}
			time.Sleep(d)
			p.connected = false

		}(&wg, c, p.topics, wait)
	}
	log.Trace().Msg("Waiting for tunnel sync to finish")
	wg.Wait()
	return nil
}

func NewTunnelPubSub(sensors []config.SSHSensor, config mqtt.Configuration, topics []string) (connection.PubSub, error) {
	l, err := InMemoryLookupCreater(sensors)
	if err != nil {
		return nil, err
	}
	return &pubsub{
		connecter: SSHConnecter,
		lookup:    l,
		config:    config,
		topics:    topics,
	}, nil
}

// Watcher periodically fetches messages from configured hosts
type Watcher struct {
	sync.RWMutex
	ps       *pubsub
	keepOpen time.Duration
	waitFor  time.Duration
	run      bool
}

func (w *Watcher) Start() {
	w.Lock()
	defer w.Unlock()
	w.run = true
	go func(run bool) {
		for run {
			for w.ps.connected {
				log.Trace().Msg("Waiting until a connection opens up")
				time.Sleep(10 * time.Millisecond)
			}
			log.Trace().Msg("Receiving")
			w.ps.Receive(w.keepOpen)
			time.Sleep(w.waitFor)
			w.RLock()
			run = w.run
			w.RUnlock()

		}

	}(true)
}

func (w *Watcher) Close() {
	w.Lock()
	defer w.Unlock()
	w.run = false
}

func NewWatcher(
	keepOpen time.Duration,
	waitFor time.Duration,
	sensors []config.SSHSensor,
	config mqtt.Configuration,
	topics []string) (connection.PubSub, *Watcher, error) {
	l, err := InMemoryLookupCreater(sensors)
	if err != nil {
		return nil, nil, err
	}
	ps := &pubsub{
		connecter: SSHConnecter,
		lookup:    l,
		config:    config,
		topics:    topics,
	}
	w := Watcher{
		ps:       ps,
		keepOpen: keepOpen,
		waitFor:  waitFor,
	}
	return ps, &w, nil
}
