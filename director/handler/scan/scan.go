// Package scan implements handler for scans
package scan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/director/handler/target"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
)

// Storage is for poutting and getting a models.Scan
type Storage interface {
	Put(models.Scan) error            // Overrides existing or creates a models.Scan
	Get(string) (*models.Scan, error) // Gets a models.Scan via ID
}

// NoopStorage is used when Storage should not have an effect
type NoopStorage struct{}

func (n NoopStorage) Put(scan models.Scan) error {
	return nil
}
func (n NoopStorage) Get(id string) (*models.Scan, error) {
	return &models.Scan{ID: id}, nil
}

// FileStorage stores models.Scan as json within a given StorageDir
// The filename is a uuid without suffix.
type FileStorage struct {
	StorageDir string
}

func (ts FileStorage) Put(scan models.Scan) error {
	b, err := json.Marshal(scan)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ts.StorageDir+"/"+scan.ID, b, 0640)
}

func (ts FileStorage) Get(id string) (*models.Scan, error) {
	var scan models.Scan
	b, err := ioutil.ReadFile(ts.StorageDir + "/" + id)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return nil, nil
		}
		return nil, err
	}
	err = json.Unmarshal(b, &scan)
	return &scan, err
}

type scanAggregate struct {
	storage Storage
	target  target.Storage
}

func (t scanAggregate) Create(c messages.Create) (*messages.Created, error) {
	scan := models.Scan{
		ID: uuid.NewString(),
	}
	if err := t.storage.Put(scan); err != nil {
		return nil, err
	}
	return &messages.Created{
		ID:      scan.ID,
		Message: messages.NewMessage("created.scan", c.MessageID, c.GroupID),
	}, nil
}

func (t scanAggregate) Modify(m messages.Modify) (*messages.Modified, *messages.Failure, error) {
	var scan *models.Scan
	scan, err := t.storage.Get(m.ID)
	if err != nil {
		return nil, nil, err
	} else if scan == nil {
		log.Printf("Scan %s not found, creating a new one.", m.ID)
		scan = &models.Scan{
			ID: m.ID,
		}
	}
	for k, v := range m.Values {
		switch k {
		case "exclude":
			// due to map[string]interface{} []string can be detected as []interface{} instead
			if cv, ok := v.([]string); ok {
				scan.Exclude = cv
			} else {
				log.Printf("Unable to cast %s to []string", v)
			}
		case "target_id", "targetId", "targetID":
			if str, ok := v.(string); ok {
				target, err := t.target.Get(str)
				if err != nil {
					return nil, nil, err
				}
				if target == nil {
					return nil, &messages.Failure{
						Message: messages.NewMessage("failure.modify.scan", m.MessageID, m.GroupID),
						Error:   fmt.Sprintf("Unable to find target: %s", v),
					}, nil
				}
				scan.Target = *target
			} else {
				return nil, &messages.Failure{
					Message: messages.NewMessage("failure.modify.scan", m.MessageID, m.GroupID),
					Error:   fmt.Sprintf("To cast %v to string", v),
				}, nil

			}
		}

	}
	if err := t.storage.Put(*scan); err != nil {
		return nil, nil, err
	}

	return &messages.Modified{
		ID:      m.ID,
		Message: messages.NewMessage("modified.scan", m.MessageID, m.GroupID),
	}, nil, nil

}
func (t scanAggregate) Get(g messages.Get) (interface{}, *messages.Failure, error) {
	if scan, err := t.storage.Get(g.ID); err != nil {
		return nil, nil, err
	} else if scan == nil {
		return nil, &messages.Failure{
			Message: messages.NewMessage("failure.get.scan", g.MessageID, g.GroupID),
			Error:   fmt.Sprintf("%s not found.", g.ID),
		}, nil
	} else {
		return &models.GotScan{
			Message: messages.NewMessage("got.scan", g.MessageID, g.GroupID),
			Scan:    *scan,
		}, nil, nil
	}
}

// New returns the type of aggregate as string and Aggregate
func New(storage Storage, target target.Storage) (string, handler.Aggregate) {

	if storage == nil {
		storage = NoopStorage{}
	}
	return "scan", scanAggregate{storage: storage, target: target}
}
