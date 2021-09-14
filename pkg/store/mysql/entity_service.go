package model

import (
	"fmt"
	"sync"

	"github.com/AndrewOYLK/ou-cmdb/types"
)

// 类似表级锁
var ModelLockMap = make(map[int64]*sync.Mutex)

// 验证并存储记录
func ValidAndSaveEntity(modelID int64, values []Value) error {
	modelMu, exist := ModelLockMap[modelID]
	if !exist {
		modelMu = &sync.Mutex{}
		ModelLockMap[modelID] = modelMu
	}
	// 加锁解锁
	modelMu.Lock()
	defer modelMu.Unlock()

	// 循环验证字段值
	validedValues := []Value{}
	for _, value := range values {
		attribute, err := GetAttribute(value.AttributeID)
		if err != nil {
			return err
		}

		str, err := attribute.ValidateValue(value.Value)
		if err != nil {
			return fmt.Errorf("attribute[%s]: %s", attribute.Key, err.Error())
		}

		validedValues = append(validedValues, Value{
			AttributeID: value.AttributeID,
			Value:       str, // 注意这里存储的是字符串值（验证通过时，返回字符串类型的值），主要方便直接插入数据库
		})
	}

	// 唯一性约束验证，列出模型ID的唯一约束字段组
	uniqueAttrsList, err := ListUniqueAttrs(modelID)
	if err != nil {
		return err
	}

	for _, uniqueAttrs := range uniqueAttrsList {
		err := uniqueAttrs.ValidateValues(validedValues)
		if err != nil {
			return err
		}
	}

	// 保存记录
	if err := mydb.CreateEntityAndValues(modelID, values); err != nil {
		return err
	}
	return nil
}

func CheckLinkAndDeleteEntity(entityID int64) error {
	// TODO 检查关联情况
	// 删除
	if err := mydb.DeleteEntityAndValues(entityID); err != nil {
		return err
	}
	return nil
}

type EntityWithValues struct {
	EntityID int64   `json:"entityID"`
	Values   []Value `json:"values"`
}

func ListEntityWithValues(modelID int64) ([]EntityWithValues, error) {
	data := make([]EntityWithValues, 0)

	// 值
	extValues, err := mydb.ListValuesByModelID(modelID)
	if err != nil {
		return data, err
	}

	t := EntityWithValues{}
	for index, extValue := range extValues {
		if t.EntityID == 0 {
			t.EntityID = extValue.EntityID
		}

		if t.EntityID != extValue.EntityID {
			data = append(data, t)

			// reset
			t = EntityWithValues{}
			t.EntityID = extValue.EntityID
		}

		// transfer value 数据库的value值是varchar类型的，而前端需要字段类型对应的值
		ty := types.NewType(extValue.Type)
		v, err := ty.ParseValue(fmt.Sprintf("%s", extValue.Value.Value))
		if err != nil {
			return data, err
		}
		extValue.Value.Value = v
		t.Values = append(t.Values, extValue.Value)

		if index == len(extValues)-1 {
			data = append(data, t)
		}
	}
	return data, nil
}
