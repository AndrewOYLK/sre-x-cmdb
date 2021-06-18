package types

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// String 字符串类型
type String struct {
	Metadata struct {
		MinLength int `json:"minLength"`
		Maxlength int `json:"maxLength"`
	}
}

func (s *String) Name() string {
	t := reflect.TypeOf(s)
	return t.Name()
}

func (s *String) SetMetadata(metadata string) error {
	if err := json.Unmarshal([]byte(metadata), &s.Metadata); err != nil {
		return err
	}
	return nil
}

func (s *String) GetMetadata() interface{} {
	return s.Metadata
}

func (s *String) ValidateValue(value interface{}) error {
	v := reflect.ValueOf(value)
	if len(v.String()) > s.Metadata.Maxlength {
		return fmt.Errorf("字符串长度[%d]大于属性限定的最大长度[%d]", len(v.String()), s.Metadata.Maxlength)
	}

	if len(v.String()) < s.Metadata.MinLength {
		return fmt.Errorf("字符串长度[%d]小于属性限定的最小长度[%d]", len(v.String()), s.Metadata.MinLength)
	}
	return nil
}

func (s *String) ToString(value interface{}) string {
	v := reflect.ValueOf(value)
	return v.String()
}

func (s *String) ParseValue(value string) (interface{}, error) {
	return value, nil
}

func (s *String) GetDefaultValue() interface{} {
	return ""
}

func (s *String) DeepCopy() TypeInterface {
	if s == nil {
		return nil
	}
	out := new(String)
	*out = *s
	out.Metadata = s.Metadata
	return out
}
