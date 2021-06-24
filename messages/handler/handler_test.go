package handler

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/models"
)

type targetAggregate struct {
}

func (t targetAggregate) ErrorOnKeyword(m messages.Message) error {
	if strings.HasSuffix(m.MessageID, "error") {
		return errors.New("something occured")
	}
	return nil
}
func (t targetAggregate) FailureOnKeyword(m messages.Message) *messages.Failure {
	if strings.HasSuffix(m.MessageID, "failure") {
		return &messages.Failure{Message: m, Error: "some failure"}
	}
	return nil
}

func (t targetAggregate) Create(c messages.Create) (*messages.Created, error) {

	if err := t.ErrorOnKeyword(c.Message); err != nil {
		return nil, err
	}
	return &messages.Created{
		ID:      "fakeid",
		Message: messages.NewMessage("created.target", c.MessageID, c.GroupID),
	}, nil
}

func (t targetAggregate) Modify(m messages.Modify) (*messages.Modified, *messages.Failure, error) {
	if err := t.ErrorOnKeyword(m.Message); err != nil {
		return nil, nil, err
	}
	if failure := t.FailureOnKeyword(m.Message); failure != nil {
		return nil, failure, nil
	}

	return &messages.Modified{
		ID:      m.ID,
		Message: messages.NewMessage("modified.target", m.MessageID, m.GroupID),
	}, nil, nil

}
func (t targetAggregate) Get(g messages.Get) (interface{}, *messages.Failure, error) {
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
func createMessage(mt string, tt string) messages.Message {
	return messages.NewMessage(mt+".target", "1234"+tt, "")
}

func createEvent(mt string, tt string) interface{} {
	switch mt {
	case "create":
		return &messages.Create{
			Message: createMessage(mt, tt),
		}
	case "modify":
		return &messages.Modify{
			Message: createMessage(mt, tt),
			ID:      "1234",
		}
	case "get":
		return &messages.Get{
			Message: createMessage(mt, tt),
			ID:      "1234",
		}
	default:
		return &messages.Failure{
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
	}
	for k, test := range tests {
		for _, j := range test {
			b, err := json.Marshal(createEvent(k, j))
			if err != nil {
				t.Errorf("[%s][%s] failed to create json", k, j)
			}
			h := New(map[string]Aggregate{
				"target": targetAggregate{},
			})
			r, err := h.On(b)
			switch j {
			case SUCCESS:
				switch k {
				case "get":
					if _, ok := r.(*models.GotTarget); !ok {
						t.Errorf("[%s][%s] expected models.GotTarget but got %T", k, j, r)
					}
				case "create":
					if _, ok := r.(*messages.Created); !ok {
						t.Errorf("[%s][%s] expected messages.Created but got %T", k, j, r)
					}
				case "modify":
					if _, ok := r.(*messages.Modified); !ok {
						t.Errorf("[%s][%s] expected messages.Modified but got %T", k, j, r)
					}

				}
			case FAILURE:
				if k != "create" {
					if _, ok := r.(*messages.Failure); !ok {
						t.Errorf("[%s][%s] expected messages.Failure but got %T", k, j, r)
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
