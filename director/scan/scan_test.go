package scan

import (
	"testing"

	"github.com/greenbone/eulabeia/internal/test"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/handler"
	"github.com/greenbone/eulabeia/storage"
)

func TestCreateScan(t *testing.T) {
	h := handler.New(New(storage.Noop{}))
	tests := []test.HandleTests{
		{
			Input: cmds.Create{
				Message: messages.NewMessage("create.scan", "1", "1"),
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("created.scan", "1", "1"),
		},
		{
			Input: cmds.Start{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("start.scan", "1", "1"),
					ID:      "1234",
				},
			},
			Handler: h,
			// although NoopStorage for target doesn't have sensor it should just
			// empty string and extend it that way
			ExpectedMessage: messages.NewMessage("start.scan.", "1", "1"),
		},
		{
			Input: cmds.Modify{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("modify.scan", "1", "2"),
					ID:      "123",
				},
				Values: map[string]interface{}{
					"finished":  []string{"1", "2"},
					"target_id": "1",
				},
			},
			ExpectedMessage: messages.NewMessage("modified.scan", "1", "2"),
			Handler:         h,
		},
		{
			Input: cmds.Modify{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("modify.scan", "1", "2"),
					ID:      "123",
				},
				Values: map[string]interface{}{
					"exclude":   []string{"1", "2"},
					"target_id": 1,
				},
			},
			Handler:         h,
			ExpectedMessage: messages.NewMessage("failure.modify.scan", "1", "2"),
		},
		{
			Input: cmds.Get{
				Identifier: messages.Identifier{
					Message: messages.NewMessage("get.scan", "1", "2"),
					ID:      "123",
				},
			},
			ExpectedMessage: messages.NewMessage("got.scan", "1", "2"),
			Handler:         h,
		},
	}
	for _, test := range tests {
		test.Verify(t)
	}

}
