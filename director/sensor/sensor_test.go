package sensor

import (
	"testing"

	"github.com/greenbone/eulabeia/internal/test"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/storage"
)

func TestSensor(t *testing.T) {
	h := handler.New(New(storage.Noop{}))
	tests := []test.HandleTests{
		{
			Input: messages.Create{
				Message: messages.NewMessage("create.sensor", "1", "1"),
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("created.sensor", "1", "1"),
		},
		{
			Input: messages.Modify{
				Message: messages.NewMessage("modify.sensor", "1", "2"),
				ID:      "123",
				Values: map[string]interface{}{
					"type": "openvas",
				},
			},
			ExpectedMessage: messages.NewMessage("modified.sensor", "1", "2"),
			Handler:         h,
		},
		{
			Input: messages.Get{
				Message: messages.NewMessage("get.sensor", "1", "2"),
				ID:      "123",
			},
			ExpectedMessage: messages.NewMessage("got.sensor", "1", "2"),
			Handler:         h,
		},
	}
	for _, test := range tests {
		test.Verify(t)
	}

}
