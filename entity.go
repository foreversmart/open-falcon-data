package data

import (
	"encoding/json"
	"time"
)

const (
	DefaultStep int = 60
)

type Entity struct {
	Endpoint    string      `json:"endpoint"`
	Tags        string      `json:"tags"`
	Metric      string      `json:"metric"`
	Value       float64     `json:"value"`
	CounterType CounterType `json:"counterType"`
	Timestamp   int64       `json:"timestamp"`
	Step        int         `json:"step"`
}

func NewEntity(endpoint, tags, metric string) *Entity {
	now := time.Now().Unix()
	return &Entity{
		Endpoint:  endpoint,
		Tags:      tags,
		Metric:    metric,
		Timestamp: now,
		Step:      DefaultStep,
	}
}

func (entity *Entity) SetValueAndType(value float64, counterType CounterType) {
	entity.Value = value
	entity.CounterType = counterType
}

func (entity *Entity) SetStep(step int) {
	entity.Step = step
}

func (entity *Entity) String() string {
	ret, err := json.Marshal(&entity)

	if err != nil {
		return ""
	}

	return string(ret)
}
