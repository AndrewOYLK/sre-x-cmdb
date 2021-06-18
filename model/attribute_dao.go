package model

import (
	builtinSQL "database/sql"
	"fmt"
	"time"
)

func (m *DB) CreateAttribute(attr Attribute) error {
	now := time.Now().Format(TimeFormat)
	sql := fmt.Sprintf("insert into attribute(`key`, modelID, type, isRequired, metadata, description, createdTime, lastUpdTime) "+
		"values('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v');", attr.Key, attr.ModelID, attr.Type, attr.IsRequired, attr.Matedata, attr.Description, now, now)

	_, err := m.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (m *DB) DeleteAttribute(id int64) error {
	sql := fmt.Sprintf("delete from attribute where id=%d;", id)

	_, err := m.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (m *DB) UpdateAttribute(attr Attribute) error {
	now := time.Now().Format(TimeFormat)

	sql := fmt.Sprintf("update attribute set `key`='%v', type='%v', isRequired='%v', metadata='%v', description='%v', lastUpdTime='%v'"+
		"where id=%d;", attr.Key, attr.Type, attr.IsRequired, attr.Matedata, attr.Description, now, attr.ID)

	_, err := m.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (m *DB) GetAttribute(id int64) (Attribute, error) {
	model := Attribute{}

	sql := fmt.Sprintf("select * from attribute where id='%d';", id)

	err := m.Get(&model, sql)
	if err == builtinSQL.ErrNoRows {
		return model, nil
	}

	if err != nil {
		return model, err
	}
	return model, nil
}

func (m *DB) ListAttribute(modelID int64) ([]Attribute, error) {
	attrs := make([]Attribute, 0)

	sql := fmt.Sprintf("select * from attribute where modelID=%d;", modelID)

	fmt.Println(sql)
	if err := m.Select(&attrs, sql); err != nil {
		return attrs, err
	}
	return attrs, nil
}

func (m *DB) CreateUniqueAttrs(uniqueAttrs UniqueAttrs) error {
	sql := fmt.Sprintf("insert into uniqueattrs(modelID, attributeIDs) values(%d, '%v');", uniqueAttrs.ModelID, uniqueAttrs.AttributeIDs)

	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (m *DB) DeleteUniqueAttrs(id int64) error {
	sql := fmt.Sprintf("delete from uniqueattrs where id=%d;", id)

	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (m *DB) UpdateUniqueAttrs(uniqueAttrs UniqueAttrs) error {
	sql := fmt.Sprintf("update uniqueattrs set attributeIDs='%v' where id=%d", uniqueAttrs.AttributeIDs, uniqueAttrs.ID)

	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (m *DB) ListUniqueAttrs(modelID int64) ([]UniqueAttrs, error) {
	uniqueAttrs := make([]UniqueAttrs, 0)

	sql := fmt.Sprintf("select * from uniqueattrs where modelID=%d;", modelID)
	if err := m.Select(&uniqueAttrs, sql); err != nil {
		if err != builtinSQL.ErrNoRows {
			return uniqueAttrs, err
		}
	}

	return uniqueAttrs, nil
}
