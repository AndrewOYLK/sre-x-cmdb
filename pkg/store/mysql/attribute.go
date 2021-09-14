package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AndrewOYLK/ou-cmdb/types"
)

/*
	example:
	ID	Key			ModelID	Type		IsRequired	Metadata																Description	CreatedTime						LastUpdTime
	1		name		1				string	true				{"minLength": 8, "maxLength": 20}				主机名称		2021-05-30 11:53:20   2021-05-30 11:53:20
	2 	ip			1				string	false				{"minLength": 8, "maxLength": 20}				ip地址			2021-05-30 11:53:20   2021-05-30 11:53:20
	3 	type		1				string	false				{"minLength": 8, "maxLength": 20}				类型				2021-05-30 11:53:20   2021-05-30 11:53:20
	4		name 		2				string 	true				{"minLength": 8, "maxLength": 20}				应用名称		2021-05-30 11:53:20   2021-05-30 11:53:20
	5		project	2				string	true				{"minLength": 8, "maxLength": 20}				所属项目		2021-05-30 11:53:20   2021-05-30 11:53:20
	6		owner		3				string	false				{"minLength": 8, "maxLength": 20}				负责人			2021-05-30 11:53:20   2021-05-30 11:53:20
*/

// Attribute 属性
type Attribute struct {
	ID          int64  `json:"id,omitempty" db:"id"`
	Key         string `json:"key,omitempty" db:"key"`
	ModelID     int64  `json:"modelID,omitempty" db:"modelID"`
	Type        string `json:"type,omitempty" db:"type"`
	IsRequired  byte   `json:"isRequired,omitempty" db:"isRequired"`
	Matedata    string `json:"metadata,omitempty" db:"metadata"` // 字段元数据（用于验证字段值有效性、设定默认值），存放字段值的约束条件
	Description string `json:"description,omitempty" db:"description"`
	CreatedTime string `json:"createdTime,omitempty" db:"createdTime"`
	LastUpdTime string `json:"lastUpdTime,omitempty" db:"lastUpdTime"`
}

// ValidAttrValue 对每个字段值进行验证，验证通过后以字符串形式返回各个类型值，方便插入数据库
func (attr *Attribute) ValidateValue(value interface{}) (string, error) {
	// 1. 根据字段的类型值，来获取类型统一管理接口
	t := types.NewType(attr.Type)

	// 2. 装载字段的Metadata, 这里存放字段值的约束条件
	if err := t.SetMetadata(attr.Matedata); err != nil {
		return "", err
	}

	// 3. 根据配置的Metadata，进行验证字段value的值
	if err := t.ValidateValue(value); err != nil {
		return "", err
	}

	return t.ToString(value), nil
}

// GetAttrDefaultValue 获取字段的默认值
func (attr *Attribute) GetDefaultValue() (interface{}, error) {
	t := types.NewType(attr.Type)

	if err := t.SetMetadata(attr.Matedata); err != nil {
		return "", err
	}

	return t.GetDefaultValue(), nil
}

// ParseAttrValue 解析字段的值（string -> 实际的字段类型值）
func (attr *Attribute) ParseValue(value string) (interface{}, error) {
	t := types.NewType(attr.Type)

	if err := t.SetMetadata(attr.Matedata); err != nil {
		return nil, err
	}

	v, err := t.ParseValue(value)
	if err != nil {
		return nil, err
	}
	return v, nil
}

/*
	example:
	ID	ModelID		AttributeIDs
	1		1					[1,2]
	2		1					[4,5]
*/

// UniqueAttrs 字段组合唯一校验
type UniqueAttrs struct {
	ID           int64  `json:"id,omitempty" db:"id"`
	ModelID      int64  `json:"modelID,omitempty" db:"modelID"`
	AttributeIDs string `json:"attributeIDs,omitempty" db:"attributeIDs"`
	CreatedTime  string `json:"createdTime,omitempty" db:"createdTime"`
	LastUpdTime  string `json:"lastUpdTime,omitempty" db:"lastUpdTime"`
}

func (u *UniqueAttrs) String() string {
	str := ""
	attributes, _ := mydb.ListAttribute(u.ModelID)
	attrIDs := strings.Split(u.AttributeIDs, ",")
	for _, attr := range attributes {
		for _, attrID := range attrIDs {
			if fmt.Sprint(attr.ID) == attrID {
				if str == "" {
					str = attr.Key
					continue
				}
				str += fmt.Sprintf("-%s", attr.Key)
			}
		}
	}
	return str
}

func (u *UniqueAttrs) ValidateValues(values []Value) error {
	uniqueAttrsValues := []Value{}
	attrIDs := strings.Split(u.AttributeIDs, ",")
	for _, attrID := range attrIDs {
		for _, v := range values {
			id, _ := strconv.ParseInt(attrID, 0, 64)
			if id == v.AttributeID {
				uniqueAttrsValues = append(uniqueAttrsValues, v)
			}
		}
	}

	if len(uniqueAttrsValues) == 0 {
		return nil
	}

	count, err := mydb.CountEntityByUniqueAttrsValues(uniqueAttrsValues)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("[%s] unique index limit", u.String())
	}
	return nil
}
