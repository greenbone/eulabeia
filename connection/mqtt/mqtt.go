package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"github.com/greenbone/eulabeia/connection"
	"log"
	"net"
)

type MQTT struct {
	client            *paho.Client
	connectProperties *paho.Connect
	qos               byte
}

func (m MQTT) Close() error {
	return m.client.Disconnect(&paho.Disconnect{ReasonCode: 0})
}

func (m MQTT) Subscribe(handler map[string]connection.OnMessage) error {
	for topic, h := range handler {
		m.client.Router.RegisterHandler(topic, func(p *paho.Publish) {
			log.Printf("Retrieved message: %s", string(p.Payload))
			message, err := h.On(p.Payload)
			if err != nil {
				panic(err)
			}
			if message != nil {
				err = m.Publish(topic, message)
				if err != nil {
					panic(err)
				}
			}
		})

		_, err := m.client.Subscribe(context.Background(), &paho.Subscribe{
			// we need NoLocal otherwise we would consum our own messages
			// again and ack them.
			Subscriptions: map[string]paho.SubscribeOptions{
				topic: {QoS: m.qos, NoLocal: true},
			},
		},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m MQTT) Publish(topic string, message interface{}) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	log.Printf("Publishing message: %s", string(b))
	props := &paho.PublishProperties{}
	pb := &paho.Publish{
		Topic:      topic,
		QoS:        m.qos,
		Payload:    b,
		Properties: props,
	}
	_, err = m.client.Publish(context.Background(), pb)
	return err
}

func (m MQTT) Connect() error {
	ca, err := m.client.Connect(context.Background(), m.connectProperties)
	if err != nil {
		return err
	}
	if ca.ReasonCode != 0 {
		return fmt.Errorf("failed to connect to %s : %d - %s", m.client.Conn.RemoteAddr().String(), ca.ReasonCode, ca.Properties.ReasonString)
	}
	return nil
}

func New(server string,
	clientid string,
	username string,
	password string) (connection.PubSub, error) {

	conn, err := net.Dial("tcp", server)
	if err != nil {
		return nil, err
	}
	c := paho.NewClient(paho.ClientConfig{
		Router: paho.NewStandardRouter(),
		Conn:   conn,
	})

	cp := &paho.Connect{
		KeepAlive:  30,
		ClientID:   clientid,
		CleanStart: true,
		Username:   username,
		Password:   []byte(password),
	}
	if username != "" {
		cp.UsernameFlag = true
	}
	if password != "" {
		cp.PasswordFlag = true
	}
	return &MQTT{
		client:            c,
		connectProperties: cp,
		qos:               1,
	}, nil
}
