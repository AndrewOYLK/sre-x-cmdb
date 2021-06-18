package model

import (
	"fmt"
	"time"
)

func (m *DB) CreateEntityAndValues(modelID int64, values []Value) error {
	createdTime := time.Now().Format(TimeFormat)
	tx := m.MustBegin()

	sql := fmt.Sprintf("insert into entity(modelID, createdTime, lastUpdTime) values(%d, '%s', '%s')", modelID, createdTime, createdTime)
	result, err := tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		return err
	}

	entityID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// insert into `value`(entityID, attributeID, `value`) values(4, 9, "vm164"),(4, 3, "192.168.101.164"),(4, 1, "kvm")
	t := ""
	for _, value := range values {
		if t == "" {
			t = fmt.Sprintf("(%d, %d, '%s')", entityID, value.AttributeID, value.Value)
			continue
		}
		t += fmt.Sprintf(",(%d, %d, '%s')", entityID, value.AttributeID, value.Value)
	}
	sql = fmt.Sprintf("insert into `value`(entityID, attributeID, `value`) values %s", t)
	if _, err := tx.Exec(sql); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (m *DB) DeleteEntityAndValues(entityID int64) error {
	tx := m.MustBegin()

	sql := fmt.Sprintf("delete from entity where id=%d", entityID)
	if _, err := tx.Exec(sql); err != nil {
		tx.Rollback()
		return err
	}

	sql = fmt.Sprintf("delete from value where entityID=%d", entityID)
	if _, err := tx.Exec(sql); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

type ExtValue struct {
	Value `json:",inline"`
	Type  string `json:"type"`
}

func (m *DB) ListValuesByModelID(modelID int64) ([]ExtValue, error) {
	values := []ExtValue{}

	sql := fmt.Sprintf("select v.id, v.entityID, v.attributeID, v.value, a.type from `value` v "+
		"left join entity e on v.entityID = e.id "+
		"left join attribute a on v.attributeID = a.id "+
		"where e.modelID=%d order by v.entityID, v.attributeID", modelID)

	if err := m.Select(&values, sql); err != nil {
		return values, err
	}
	return values, nil
}

func (m *DB) CountEntityByUniqueAttrsValues(values []Value) (int64, error) {
	/*
		select count(*)
		from `value` as v1
		left join `value` as v2 on v.entityID = v2.entityID
		left join `value` as v3 on v.entityID = v3.entityID
		where v1.attributeID = 9 and v1.`value` = "vm161"
			and v2.attributeID = 3 and v2.`value` = "192.168.101.161"
			and v3.attributeID = 1 and v3.`value` = "kvm"
	*/

	var count int64 = 0
	raw := ""
	when := ""
	and := ""
	leftJoin := ""
	for index, value := range values {
		if index == 0 {
			raw = "select count(*) from `value` as a"
			when = fmt.Sprintf("where a.attributeID=%d and a.`value`='%s'", value.AttributeID, value.Value)
			continue
		}

		alias := string(98 + index)
		leftJoin += fmt.Sprintf("left join `value` as %s on a.entityID=%s.entityID", alias, alias)
		and += fmt.Sprintf("and %s.attributeID=%d and %s.`value`='%s'", alias, value.AttributeID, alias, value.Value)
	}

	sql := fmt.Sprintf("%s %s %s %s", raw, leftJoin, when, and)
	fmt.Println(sql)
	if err := m.Get(&count, sql); err != nil {
		return count, err
	}
	return count, nil
}
