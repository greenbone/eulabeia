package messagehandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/models"
	"github.com/tidwall/gjson"
)

type OnChangeEvent interface {
	Change(message_type string, message []byte, writer io.Writer) (interface{}, error)
}

type OnCreateTarget struct{}

func (oct OnCreateTarget) Change(message_type string, message []byte, writer io.Writer) (interface{}, error) {
	if message_type != "create.target" {
		return nil, nil
	}
	target := models.Target{
		ID: uuid.NewString(),
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(target)
	if err != nil {
		return nil, err
	}
	_, err = writer.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	groupID := gjson.GetBytes(message, "group_id")
	messageID := gjson.GetBytes(message, "message_id")
	returnMessage := messages.Created{
		CreatedID: target.ID,
		Message: messages.Message{
			Created:     time.Now().Nanosecond(),
			MessageType: "created.target",
			MessageID:   messageID.String(),
			GroupID:     groupID.String(),
		},
	}

	return returnMessage, nil
}

type OnMessage interface {
	On(message []byte) (interface{}, error)
}

type MessageHandler struct {
	changeEventHandler []OnChangeEvent
	writer             io.Writer
}

func (mh MessageHandler) On(message []byte) (result interface{}, err error) {
	messageType := gjson.GetBytes(message, "message_type")
	if messageType.Type == gjson.Null {
		fmt.Printf("message: %s does not contain message_type!", string(message))
		return
	}
	for _, i := range mh.changeEventHandler {
		result, err = i.Change(messageType.String(), message, mh.writer)
		if result != nil || err != nil {
			return
		}
	}
	return nil, fmt.Errorf("unable to find handler for %s", messageType.String())
}

func NewMessageHandler(writer io.Writer, changeEventHandler []OnChangeEvent) (OnMessage, error) {
	return MessageHandler{changeEventHandler: changeEventHandler, writer: writer}, nil
}
