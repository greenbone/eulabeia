// Package target implements handler for targets
package target

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/models"
	"github.com/tidwall/gjson"
)

type Storage interface {
	Put(models.Target) error
	Get(string) (*models.Target, error)
}

type NoopStorage struct{}

func (n NoopStorage) Put(target models.Target) error {
	return nil
}
func (n NoopStorage) Get(id string) (*models.Target, error) {
	return &models.Target{ID: id}, nil
}

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
	} else {
		return models.GotTarget{
			Message: messages.NewMessage("got.target", get.MessageID, get.GroupID),
			Target:  *target,
		}, nil
	}
}

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
	target, e := omt.storage.Get(modify.ID)
	if e != nil {
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
