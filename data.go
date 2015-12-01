package data

import (
	"encoding/json"
)

type Data struct {
	entities []*Entity
}

func NewData() *Data {
	return &Data{
		entities: make([]*Entity, 0, 10),
	}
}

func (data *Data) AddEntity(entity *Entity) {
	data.entities = append(data.entities, entity)
}

func (data *Data) String() (string, error) {
	res, err := json.Marshal(data.entities)

	if err != nil {
		return "", err
	}

	return string(res), nil
}
