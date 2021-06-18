package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Number struct {
	Metadata struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	}
}

func (n *Number) Name() string {
	t := reflect.TypeOf(n)
	return t.Name()
}

func (n *Number) SetMetadata(metadata string) error {
	if err := json.Unmarshal([]byte(metadata), &n.Metadata); err != nil {
		return err
	}
	return nil
}

func (n *Number) GetMetadata() interface{} {
	return n.Metadata
}

func (n *Number) ValidateValue(value interface{}) error {
	defer func() error {
		if v := recover(); v != nil {
			return fmt.Errorf("panic: %v", v)
		}
		return nil
	}()

	v := reflect.ValueOf(value)
	number := v.Float()

	if number < n.Metadata.Min {
		return fmt.Errorf("输入的数值[%v]小于属性限定的最小值[%v]", number, n.Metadata.Min)
	}

	if number > n.Metadata.Max {
		return fmt.Errorf("输入的数值[%v]大于属性限定的最大值[%v]", number, n.Metadata.Max)
	}

	return nil
}

func (n *Number) ToString(value interface{}) string {
	v := reflect.ValueOf(value)
	return strconv.FormatFloat(v.Float(), 'g', -1, 64)
}

func (n *Number) ParseValue(value string) (interface{}, error) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (n *Number) GetDefaultValue() interface{} {
	return 0
}

func (n *Number) DeepCopy() TypeInterface {
	if n == nil {
		return nil
	}
	out := new(Number)
	*out = *n
	out.Metadata = n.Metadata
	return out
}
