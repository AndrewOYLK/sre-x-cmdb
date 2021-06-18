package types

import (
	"encoding/json"
	"reflect"
	"strconv"
)

type Bool struct {
	Metadata struct {
		Default bool `json:"default"`
	}
}

func (b *Bool) Name() string {
	t := reflect.TypeOf(b)
	return t.Name()
}

func (b *Bool) SetMetadata(metadata string) error {
	if err := json.Unmarshal([]byte(metadata), &b.Metadata); err != nil {
		return err
	}
	return nil
}

func (b *Bool) GetMetadata() interface{} {
	return b.Metadata
}

func (b *Bool) ValidateValue(value interface{}) error {
	return nil
}

func (b *Bool) ToString(value interface{}) string {
	v := reflect.ValueOf(value)
	return strconv.FormatBool(v.Bool())
}

func (b *Bool) ParseValue(value string) (interface{}, error) {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (b *Bool) GetDefaultValue() interface{} {
	return b.Metadata.Default
}

func (b *Bool) DeepCopy() TypeInterface {
	if b == nil {
		return nil
	}
	out := new(Bool)
	*out = *b
	out.Metadata = b.Metadata
	return out
}
