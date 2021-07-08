package handler

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/messages/cmds"
	"github.com/greenbone/eulabeia/messages/info"
	"github.com/greenbone/eulabeia/models"
)

type exampleAggregate struct {
}

func (t exampleAggregate) ErrorOnKeyword(m messages.Message) error {
	if strings.HasSuffix(m.MessageID, "error") {
		return errors.New("something occured")
	}
	return nil
}
func (t exampleAggregate) FailureOnKeyword(m messages.Message) *info.Failure {
	if strings.HasSuffix(m.MessageID, "failure") {
		return &info.Failure{Message: m, Error: "some failure"}
	}
	return nil
}

func (t exampleAggregate) Create(c cmds.Create) (*info.Created, error) {

	if err := t.ErrorOnKeyword(c.Message); err != nil {
		return nil, err
	}
	return &info.Created{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("created.target", c.MessageID, c.GroupID),
			ID:      "fakeid",
		},
	}, nil
}

func (t exampleAggregate) Modify(m cmds.Modify) (*info.Modified, *info.Failure, error) {
	if err := t.ErrorOnKeyword(m.Message); err != nil {
		return nil, nil, err
	}
	if failure := t.FailureOnKeyword(m.Message); failure != nil {
		return nil, failure, nil
	}

	return &info.Modified{
		Identifier: messages.Identifier{
			Message: messages.NewMessage("modified.target", m.MessageID, m.GroupID),
			ID:      m.ID,
		},
	}, nil, nil

}
func (t exampleAggregate) Get(g cmds.Get) (interface{}, *info.Failure, error) {
	if err := t.ErrorOnKeyword(g.Message); err != nil {
		return nil, nil, err
	}
	if failure := t.FailureOnKeyword(g.Message); failure != nil {
		return nil, failure, nil
	}

	return &models.GotTarget{
		Message: g.Message,
	}, nil, nil
}
func (t exampleAggregate) Delete(g cmds.Delete) (*info.Deleted, *info.Failure, error) {
	if err := t.ErrorOnKeyword(g.Message); err != nil {
		return nil, nil, err
	}
	if failure := t.FailureOnKeyword(g.Message); failure != nil {
		return nil, failure, nil
	}

	return &info.Deleted{
		Identifier: messages.Identifier{
			Message: g.Message,
			ID:      g.ID,
		},
	}, nil, nil
}

func createMessage(mt string, tt string) messages.Message {
	return messages.NewMessage(mt+".target", "1234"+tt, "")
}

func createEvent(mt string, tt string) interface{} {
	switch mt {
	case "create":
		return &cmds.Create{
			Message: createMessage(mt, tt),
		}
	case "modify":
		return &cmds.Modify{
			Identifier: messages.Identifier{
				Message: createMessage(mt, tt),
				ID:      "1234",
			},
		}
	case "get":
		return &cmds.Get{
			Identifier: messages.Identifier{
				Message: createMessage(mt, tt),
				ID:      "1234",
			},
		}
	case "delete":
		return &cmds.Delete{
			Identifier: messages.Identifier{
				Message: createMessage(mt, tt),
				ID:      "1234",
			},
		}
	default:
		return &info.Failure{
			Message: createMessage("", "failure"),
		}
	}
}

const (
	SUCCESS string = "success"
	FAILURE string = "failure"
	ERROR   string = "error"
)

func TestAggragteHandler(t *testing.T) {
	all := []string{SUCCESS, FAILURE, ERROR}
	var tests = map[string][]string{
		"create": all,
		"modify": all,
		"get":    all,
		"delete": all,
	}
	for k, test := range tests {
		for _, j := range test {
			b, err := json.Marshal(createEvent(k, j))
			if err != nil {
				t.Errorf("[%s][%s] failed to create json", k, j)
			}
			h := New(FromAggregate("target", exampleAggregate{}))
			r, err := h.On("", b)
			switch j {
			case SUCCESS:
				switch k {
				case "delete":
					if _, ok := r.MSG.(*info.Deleted); !ok {
						t.Errorf("[%s][%s] expected models.GotTarget but got %T", k, j, r)
					}
				case "get":
					if _, ok := r.MSG.(*models.GotTarget); !ok {
						t.Errorf("[%s][%s] expected models.GotTarget but got %T", k, j, r)
					}
				case "create":
					if _, ok := r.MSG.(*info.Created); !ok {
						t.Errorf("[%s][%s] expected info.Created but got %T", k, j, r)
					}
				case "modify":
					if _, ok := r.MSG.(*info.Modified); !ok {
						t.Errorf("[%s][%s] expected info.Modified but got %T", k, j, r)
					}

				}
			case FAILURE:
				if k != "create" {
					if _, ok := r.MSG.(*info.Failure); !ok {
						t.Errorf("[%s][%s] expected info.Failure but got %T", k, j, r)
					}
				}
			case ERROR:
				if err == nil {
					t.Errorf("[%s][%s] expected error but is nil; got msg %T instead", k, j, r)
				}

			}
		}
	}
}
