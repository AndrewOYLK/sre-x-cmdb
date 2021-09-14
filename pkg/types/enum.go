package types

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Enum struct {
	Metadata []struct {
		Key     string `json:"key"`
		Value   string `json:"value"`
		Default string `json:"default"`
	}
}

func (e *Enum) Name() string {
	t := reflect.TypeOf(e)
	return t.Name()
}

func (e *Enum) SetMetadata(metadata string) error {
	if err := json.Unmarshal([]byte(metadata), &e.Metadata); err != nil {
		return err
	}
	return nil
}

func (e *Enum) GetMetadata() interface{} {
	return e.Metadata
}

func (e *Enum) ValidateValue(value interface{}) error {
	flag := false
	for _, v := range e.Metadata {
		if v.Key == value {
			flag = true
		}
	}

	if !flag {
		return fmt.Errorf("选择的枚举key[%v]不在有效范围内", value)
	}
	return nil
}

func (e *Enum) ToString(value interface{}) string {
	return ""
}

func (e *Enum) ParseValue(value string) (interface{}, error) {
	return nil, nil
}

func (e *Enum) GetDefaultValue() interface{} {
	return nil
}

func (e *Enum) DeepCopy() TypeInterface {
	if e == nil {
		return nil
	}
	out := new(Enum)
	*out = *e
	out.Metadata = e.Metadata
	return out
}
