package day

import (
	"time"

	"github.com/dpb587/slack-alias-bot/condition"
	"github.com/dpb587/slack-alias-bot/message"
)

type Condition struct {
	Location *time.Location
	Days     []string
}

var _ condition.Condition = &Condition{}

func (c Condition) Evaluate(m message.Message) (bool, error) {
	actual := m.Timestamp.In(c.Location).Format("Mon")

	for _, expected := range c.Days {
		if expected == actual {
			return true, nil
		}
	}

	return false, nil
}
