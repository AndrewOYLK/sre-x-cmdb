package types

import (
	"encoding/json"
	"reflect"
)

type Date struct {
	Metadata struct{}
}

func (d *Date) Name() string {
	t := reflect.TypeOf(d)
	return t.Name()
}

func (d *Date) SetMetadata(metadata string) error {
	if err := json.Unmarshal([]byte(metadata), &d.Metadata); err != nil {
		return err
	}
	return nil
}

func (d *Date) GetMetadata() interface{} {
	return d.Metadata
}

func (d *Date) ValidateValue(value interface{}) error {
	return nil
}

func (d *Date) ToString(value interface{}) string {
	v := reflect.ValueOf(value)
	return v.String()
}

func (d *Date) ParseValue(value string) (interface{}, error) {
	return value, nil
}

func (d *Date) GetDefaultValue() interface{} {
	return ""
}

func (d *Date) DeepCopy() TypeInterface {
	if d == nil {
		return nil
	}
	out := new(Date)
	*out = *d
	out.Metadata = d.Metadata
	return out
}
