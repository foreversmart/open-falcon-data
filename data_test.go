package data

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Data(t *testing.T) {
	var (
		assertion                     = assert.New(t)
		endpoint, tags, metric string = "test-endpoint", "test", "test-metric"
	)

	entity := NewEntity(endpoint, tags, metric)
	entity.SetValueAndType(float64(6.33), CounterTypeCOUNTER)

	data := NewData()
	data.AddEntity(entity)
	data.AddEntity(entity)
	ret, err := data.String()
	exp := "[{\"endpoint\":\"test-endpoint\",\"tags\":\"test\",\"metric\":\"test-metric\",\"value\":6.33,\"counterType\":\"COUNTER\",\"timestamp\":%d,\"step\":60},{\"endpoint\":\"test-endpoint\",\"tags\":\"test\",\"metric\":\"test-metric\",\"value\":6.33,\"counterType\":\"COUNTER\",\"timestamp\":%d,\"step\":60}]"

	assertion.Nil(err)
	assertion.Equal(fmt.Sprintf(exp, entity.Timestamp, entity.Timestamp), ret)
}
