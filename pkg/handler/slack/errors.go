package slack

import (
	"fmt"

	"github.com/mattmoor/knobots/pkg/botinfo"
	"github.com/mattmoor/knobots/pkg/handler"
)

func ErrorReport(message string, attributes map[string]string) handler.Response {
	lines := []string{
		fmt.Sprintf("bot: %s", botinfo.GetName()),
		fmt.Sprintf("message: %s", message),
	}

	for k, v := range attributes {
		lines = append(lines, fmt.Sprintf("%s: %s", k, v))
	}

	return &DirectMessage{
		Emails:  []string{"mattomata@gmail.com"},
		Message: lines,
	}
}
