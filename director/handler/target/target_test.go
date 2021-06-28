package target

import (
	"encoding/json"
	"fmt"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/handler"
	"testing"
)

func TestSuccessResponse(t *testing.T) {
	var tests = []struct {
		on   interface{}
		then string
	}{
		{messages.Get{
			Message: messages.Message{
				MessageType: "get.target",
				Created:     7774,
				MessageID:   "1",
				GroupID:     "12",
			},
			ID: "someid",
		},
			"*models.GotTarget",
		},
		{messages.Create{
			Message: messages.Message{
				MessageType: "create.target",
				Created:     7774,
				MessageID:   "1",
				GroupID:     "12",
			},
		},
			"*messages.Created",
		},
		{messages.Create{
			Message: messages.Message{
				MessageType: "created.target",
				Created:     7774,
				MessageID:   "1",
				GroupID:     "12",
			},
		},
			"<nil>",
		},
		{messages.Modify{
			Message: messages.Message{
				MessageType: "modify.target",
				Created:     7774,
				MessageID:   "1",
				GroupID:     "12",
			},
			ID: "1",
			Values: map[string]interface{}{
				"scanner":  "openvas",
				"hosts":    []string{"a", "b"},
				"plugins":  []string{"a", "b"},
				"alive":    true,
				"parallel": false,
			},
		},
			"*messages.Modified",
		},
	}
	for i, test := range tests {
		b, err := json.Marshal(test.on)
		if err != nil {
			t.Errorf("[%d] marshalling [%v] to json failed", i, test.on)
		}
		h := handler.New(handler.FromAggregate(New(NoopStorage{})))
		r, err := h.On(b)
		if err != nil {
			t.Errorf("[%d] returned err (%v) on: %v", i, err, test.on)
		}
		ts := fmt.Sprintf("%T", r)
		if ts != test.then {
			t.Errorf("[%d] returned %v while expecting %v on: %v", i, ts, test.then, test.on)
		}
	}
}
