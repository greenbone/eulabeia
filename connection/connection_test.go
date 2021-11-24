package connection

import (
	"errors"
	"testing"
)

var testData = TopicData{Topic: "test", Message: []byte{0, 1, 2}}

func TestIncomingSuccess(t *testing.T) {
	handled := 0
	h := map[string]OnMessage{
		"test": ClosureOnMessage{func(td TopicData) (*SendResponse, error) {
			handled = handled + 1
			return nil, nil
		}},
		"respond": ClosureOnMessage{func(td TopicData) (*SendResponse, error) {
			handled = handled + 1
			return &SendResponse{Topic: "test", MSG: nil}, nil
		}},
	}
	published := 0
	publisher := []Publisher{
		ClosurePublisher{func(s string, i interface{}) error {
			published = published + 1
			return nil
		}},
	}
	in := make(chan *TopicData, 1)
	in <- &testData
	out := make(chan *SendResponse, 1)
	mh := NewMessageHandler(h, []Preprocessor{}, publisher, in, out)
	if !mh.Check() {
		t.Fatalf("Expected check to be true")
	}
	if handled != 1 {
		t.Fatalf("Expected test handler to be called once")
	}
	in <- &TopicData{Topic: "respond", Message: []byte{0, 1, 2}}
	if !mh.Check() {
		t.Fatalf("Expected check to be true")
	}
	if handled != 2 {
		t.Fatalf("Expected test handler to be called twice")
	}
	if published != 1 {
		t.Fatalf("Expected publisher to be called once")
	}
}
func TestIncomingClosed(t *testing.T) {
	h := map[string]OnMessage{}
	publisher := []Publisher{}
	in := make(chan *TopicData, 1)
	out := make(chan *SendResponse, 1)
	mh := NewMessageHandler(h, []Preprocessor{}, publisher, in, out)
	close(in)
	if mh.Check() {
		t.Fatalf("Expected check to be false (due to closed in)")
	}
}
func TestIncomingFailureNotPanic(t *testing.T) {
	handled := 0
	h := map[string]OnMessage{
		"test": ClosureOnMessage{func(td TopicData) (*SendResponse, error) {
			handled = handled + 1
			return nil, errors.New("Horrible")
		}},
	}
	publisher := []Publisher{}
	in := make(chan *TopicData, 1)
	in <- &testData
	out := make(chan *SendResponse, 1)
	mh := NewMessageHandler(h, []Preprocessor{}, publisher, in, out)
	if !mh.Check() {
		t.Fatalf("Expected check to be true")
	}
	if handled != 1 {
		t.Fatalf("Expected test handler to be called once")
	}
}
func TestOutgoingSuccess(t *testing.T) {
	h := map[string]OnMessage{}
	published := 0
	publisher := []Publisher{
		ClosurePublisher{func(s string, i interface{}) error {
			published = published + 1
			return nil
		}},
	}
	in := make(chan *TopicData, 1)
	out := make(chan *SendResponse, 1)
	out <- &SendResponse{Topic: "Test"}
	mh := NewMessageHandler(h, []Preprocessor{}, publisher, in, out)
	if !mh.Check() {
		t.Fatalf("Expected check to be true")
	}
	if published != 1 {
		t.Fatalf("Expected publisher to be called once")
	}

}
func TestOutgoingClosed(t *testing.T) {
	h := map[string]OnMessage{}
	publisher := []Publisher{}
	in := make(chan *TopicData, 1)
	out := make(chan *SendResponse, 1)
	close(out)
	mh := NewMessageHandler(h, []Preprocessor{}, publisher, in, out)
	if mh.Check() {
		t.Fatalf("Expected check to be false")
	}
}
func TestOutgoingFailureNotPanic(t *testing.T) {
	h := map[string]OnMessage{}
	published := 0
	publisher := []Publisher{
		ClosurePublisher{func(s string, i interface{}) error {
			published = published + 1
			return errors.New("Something")
		}},
	}
	in := make(chan *TopicData, 1)
	out := make(chan *SendResponse, 1)
	out <- &SendResponse{Topic: "Test"}
	mh := NewMessageHandler(h, []Preprocessor{}, publisher, in, out)
	if !mh.Check() {
		t.Fatalf("Expected check to be true")
	}
	if published != 1 {
		t.Fatalf("Expected publisher to be called once")
	}

}
func TestPreprocessor(t *testing.T) {
	handled := 0
	h := map[string]OnMessage{
		"test": ClosureOnMessage{func(td TopicData) (*SendResponse, error) {
			handled = handled + 1
			return nil, nil
		}},
	}
	publisher := []Publisher{}
	in := make(chan *TopicData, 1)
	out := make(chan *SendResponse, 1)
	Preprocessor := []Preprocessor{
		ClosurePreprocessor{func(td TopicData) ([]TopicData, bool) {
			return []TopicData{testData}, true
		}},
	}
	in <- &TopicData{Topic: "holla"}
	mh := NewMessageHandler(h, Preprocessor, publisher, in, out)
	if !mh.Check() {
		t.Fatalf("Expected check to be true")
	}
	if handled != 1 {
		t.Fatalf("Expected test handler to be called once")
	}
}
