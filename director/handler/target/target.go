// Package target implements handler for targets
package target

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	dh "github.com/greenbone/eulabeia/director/handler"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
)

// Storage is for poutting and getting a models.Target
type Storage interface {
	Put(models.Target) error            // Overrides existing or creates a models.Target
	Get(string) (*models.Target, error) // Gets a models.Target via ID
}

// NoopStorage is used when Storage should not have an effect
type NoopStorage struct{}

func (n NoopStorage) Put(target models.Target) error {
	return nil
}
func (n NoopStorage) Get(id string) (*models.Target, error) {
	return &models.Target{ID: id}, nil
}

// FileStorage stores models.Target as json within a given StorageDir
// The filename is a uuid without suffix.
type FileStorage struct {
	StorageDir string
}

func (ts FileStorage) Put(target models.Target) error {
	b, err := json.Marshal(target)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ts.StorageDir+"/"+target.ID, b, 0640)
}

func (ts FileStorage) Get(id string) (*models.Target, error) {
	var target models.Target
	b, err := ioutil.ReadFile(ts.StorageDir + "/" + id)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return nil, nil
		}
		return nil, err
	}
	err = json.Unmarshal(b, &target)
	return &target, err
}

type targetAggregate struct {
	storage Storage
}

func (t targetAggregate) Create(c messages.Create) (*messages.Created, error) {
	target := models.Target{
		ID: uuid.NewString(),
	}
	if err := t.storage.Put(target); err != nil {
		return nil, err
	}
	return &messages.Created{
		ID:      target.ID,
		Message: messages.NewMessage("created.target", c.MessageID, c.GroupID),
	}, nil
}

func (t targetAggregate) Modify(m messages.Modify) (*messages.Modified, *messages.Failure, error) {
	var target *models.Target
	target, err := t.storage.Get(m.ID)
	if err != nil {
		return nil, nil, err
	} else if target == nil {
		log.Printf("Target %s not found, creating a new one.", m.ID)
		target = &models.Target{
			ID: m.ID,
		}
	}
	if f := dh.GenericSetValueOf(target, m); f != nil {
		return nil, f, nil
	}

	if err := t.storage.Put(*target); err != nil {
		return nil, nil, err
	}

	return &messages.Modified{
		ID:      m.ID,
		Message: messages.NewMessage("modified.target", m.MessageID, m.GroupID),
	}, nil, nil

}
func (t targetAggregate) Get(g messages.Get) (interface{}, *messages.Failure, error) {
	if target, err := t.storage.Get(g.ID); err != nil {
		return nil, nil, err
	} else if target == nil {
		return nil, &messages.Failure{
			Message: messages.NewMessage("failure.get.target", g.MessageID, g.GroupID),
			Error:   fmt.Sprintf("%s not found.", g.ID),
		}, nil
	} else {
		return &models.GotTarget{
			Message: messages.NewMessage("got.target", g.MessageID, g.GroupID),
			Target:  *target,
		}, nil, nil
	}
}

// New returns the type of aggregate as string and Aggregate
func New(storage Storage) (string, handler.Aggregate) {

	if storage == nil {
		storage = NoopStorage{}
	}
	return "target", targetAggregate{storage: storage}
}
