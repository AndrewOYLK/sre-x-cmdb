package model

import (
	builtinSQL "database/sql"
	"fmt"
	"time"
)

func (m *DB) CreateModel(model Model) error {
	createdTime := time.Now().Format(TimeFormat)

	sql := fmt.Sprintf("insert into model(`key`, description, createdTime, lastUpdTime) values('%v', '%v', '%v', '%v');", model.Key, model.Description, createdTime, createdTime)

	_, err := m.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (m *DB) GetModel(id int64) (Model, error) {
	var model = Model{}

	sql := fmt.Sprintf(`
		select * from model where id = %v;
	`, id)

	err := m.Get(&model, sql)
	if err == builtinSQL.ErrNoRows {
		return model, nil
	}

	if err != nil {
		return model, err
	}

	return model, nil
}

func (m *DB) ListModel() ([]Model, error) {
	models := make([]Model, 0)
	err := m.Select(&models, "select * from model;")
	if err != nil {
		return models, err
	}

	return models, nil
}

func (m *DB) DeleteModel(id int64) error {
	sql := fmt.Sprintf("delete from model where id=%v", id)
	_, err := m.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (m *DB) UpdateModel(model Model) error {
	lastUpdTime := time.Now().Format(TimeFormat)

	sql := fmt.Sprintf("update model set `key`='%s', description='%s', lastUpdTime='%s' where id = %d", model.Key, model.Description, lastUpdTime, model.ID)
	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}
