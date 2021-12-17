package tunnel

import (
	"strings"

	"github.com/greenbone/eulabeia/connection"
)

func CreateRepublish(filter string) func(connection.TopicData) *connection.SendResponse {
	return func(td connection.TopicData) *connection.SendResponse {
		for _, s := range td.Sender {
			if strings.HasPrefix(s, filter) {
				return &connection.SendResponse{
					Topic: td.Topic,
					MSG:   td.Message,
				}
			}
		}
		return nil
	}
}
