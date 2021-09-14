package types

/*
	在全局来看对不同的类型的业务逻辑操作是要一样的，但是不同的类型有不同的验证细节，
	所以分开来通过一个接口来统一管理类型会比较方便和直观
*/

// 允许的类型
var AllowedType = map[string]TypeInterface{
	"string": &String{},
	"number": &Number{},
	"bool":   &Bool{},
	"date":   &Date{},
	"enum":   &Enum{},
}

// TypeInterface 各类型操作接口
type TypeInterface interface {
	Name() string
	SetMetadata(metadata string) error
	GetMetadata() interface{}              // 由于每个字段的Metadata都无确定的结构体，此处是json格式数据
	ValidateValue(value interface{}) error // 字段值验证
	ToString(value interface{}) string
	ParseValue(value string) (interface{}, error) // 根据字段值解析为相应字段的真实类型值
	GetDefaultValue() interface{}
	DeepCopy() TypeInterface
}

// NewType 获取类型接口
func NewType(name string) TypeInterface {
	v, exist := AllowedType[name]
	if exist {
		return v.DeepCopy()
	}
	return nil
}
