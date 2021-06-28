// Package handler contains various message handler for directors
package handler

import (
	"fmt"
	"github.com/greenbone/eulabeia/messages"
	"github.com/greenbone/eulabeia/models"
	"log"
	"strings"
)

func GenericSetValueOf(target interface{}, m messages.Modify) *messages.Failure {
	for k, v := range m.Values {
		// normalize field name
		nk := strings.Title(k)
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
			return &messages.Failure{
				Error:   fmt.Sprintf("Unable to set %s on target to %s: %v", nk, v, failure),
				Message: messages.NewMessage("failure.modified.target", m.MessageID, m.GroupID),
			}
		}
	}
	return nil
}
