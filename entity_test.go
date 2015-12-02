package data

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Entity(t *testing.T) {
	var (
		assertion                     = assert.New(t)
		endpoint, tags, metric string = "test-endpoint", "test", "test-metric"
	)

	entity := NewEntity(endpoint, tags, metric)
	entity.SetValueAndType(float64(2.333), CounterTypeGAUGE)

	assertion.EqualValues(fmt.Sprintf("{\"endpoint\":\"test-endpoint\",\"tags\":\"test\",\"metric\":\"test-metric\",\"value\":2.333,\"counterType\":\"GAUGE\",\"timestamp\":%d,\"step\":60}", entity.Timestamp), entity.String())
}
