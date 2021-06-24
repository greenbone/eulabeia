// Package target implements handler for targets
package target

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
	"github.com/tidwall/gjson"
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

// OnCreate handles the message create.target it will respond with a created.target message containing a uuid
type OnCreate struct {
	storage Storage
}

func (oct OnCreate) On(messageType string, message []byte) (interface{}, error) {
	if messageType != "create.target" {
		return nil, nil
	}
	target := models.Target{
		ID: uuid.NewString(),
	}
	if err := oct.storage.Put(target); err != nil {
		return nil, err
	}
	groupID := gjson.GetBytes(message, "groupId")
	messageID := gjson.GetBytes(message, "messageId")
	returnMessage := messages.Created{
		ID:      target.ID,
		Message: messages.NewMessage("created.target", messageID.String(), groupID.String()),
	}

	return returnMessage, nil
}

// OnGet handles the message get.target it will respond with models.GotTarget
type OnGet struct {
	storage Storage
}

func (ogt OnGet) On(messageType string, message []byte) (interface{}, error) {
	if messageType != "get.target" {
		return nil, nil
	}
	var get messages.Get
	if err := json.Unmarshal(message, &get); err != nil {
		return nil, err
	}
	if target, err := ogt.storage.Get(get.ID); err != nil {
		return nil, err
	} else if target == nil {
		return messages.Failure{
			Message: messages.NewMessage("failure.get.target", get.MessageID, get.GroupID),
			Error:   fmt.Sprintf("%s not found.", get.ID),
		}, nil
	} else {
		return models.GotTarget{
			Message: messages.NewMessage("got.target", get.MessageID, get.GroupID),
			Target:  *target,
		}, nil
	}
}

// OnModify handles the message modify.target; it will respond with messages.Modified
type OnModify struct {
	storage Storage
}

func (omt OnModify) On(messageType string, message []byte) (msg interface{}, err error) {
	if messageType != "modify.target" {
		return nil, nil
	}
	var modify messages.Modify
	err = json.Unmarshal(message, &modify)
	if err != nil {
		return
	}
	var target *models.Target
	target, err = omt.storage.Get(modify.ID)
	if err != nil {
		return
	} else if target == nil {
		log.Printf("Target %s not found, creating a new one.", modify.ID)
		target = &models.Target{
			ID: modify.ID,
		}
	}
	for k, v := range modify.Values {
		// normalize field name
		nk := strings.Title(k)
		// don't return failure as err since it is based on the message
		// and cannot be handled by the program itself anyway
		var failure error
		// due to map[string]interface{} []string can be detected as []interface{} instead
		switch cv := v.(type) {
		case []interface{}:
			strings := make([]string, len(cv), cap(cv))
			for i, j := range cv {
				if s, ok := j.(string); ok {
					strings[i] = s
				}
			}
			failure = models.SetValueOf(target, nk, strings)
		default:
			failure = models.SetValueOf(target, nk, cv)
		}
		if failure != nil {
			log.Printf("Failure while processing field %v: %v", nk, failure)
			msg = messages.Failure{
				Error:   fmt.Sprintf("Unable to set %s on target to %s: %v", nk, v, failure),
				Message: messages.NewMessage("failure.modified.target", modify.MessageID, modify.GroupID),
			}
			return
		}
	}
	omt.storage.Put(*target)
	msg = messages.Modified{
		ID:      modify.ID,
		Message: messages.NewMessage("modified.target", modify.MessageID, modify.GroupID),
	}

	return
}

// New returns a list of all known *.target handler.
// A handler returns nil, nil when it either got handled or when it did not handle the message
// Returns either an message.Failure or the corresponding response.
// Returns an error if something unexpected did happen and the containing program should react on that
func New(storage Storage) []handler.OnEvent {
	if storage == nil {
		storage = NoopStorage{}
	}
	return []handler.OnEvent{
		OnCreate{storage: storage},
		OnModify{storage: storage},
		OnGet{storage: storage},
	}
}
