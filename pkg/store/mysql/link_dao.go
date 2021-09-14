package model

import (
	"fmt"
	"time"
)

// LinkType

func (m *DB) CreateLinkType(linkType LinkType) error {
	createdTime := time.Now().Format(TimeFormat)

	sql := fmt.Sprintf("insert into linktype(key, from, to, createdTime, lastUpdTime) "+
		"values('%s', '%s', '%s', '%s', '%s')", linkType.Key, linkType.From, linkType.To, createdTime, createdTime)

	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (m *DB) DeleteLinkType(id int64) error {
	sql := fmt.Sprintf("delete from linktype where id=%d", id)

	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (m *DB) UpdateLinkType(linkType LinkType) error {
	updateTime := time.Now().Format(TimeFormat)
	sql := fmt.Sprintf("update linktype set key='%s',from='%s',to='%s',lastUpdTime='%s' where id=%d", linkType.Key, linkType.From, linkType.To, updateTime, linkType.ID)

	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (m *DB) ListLinkTypes() ([]LinkType, error) {
	linkTypes := make([]LinkType, 0)

	sql := "select * from linktype"
	if err := m.Select(&linkTypes, sql); err != nil {
		return linkTypes, err
	}
	return linkTypes, nil
}

// LinkModel

func (m *DB) CreateLinkModel(linkModel LinkModel) error {
	createdTime := time.Now().Format(TimeFormat)
	sql := fmt.Sprintf("insert into linkmodel(linkTypeID, fromModelID, toModelID, mapping, createdTime, lastUpdTime) "+
		"values(%d, %d, %d, '%s', '%s', '%s')", linkModel.LinkTypeID, linkModel.FromModelID, linkModel.ToModelID, linkModel.Mapping, createdTime, createdTime)

	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (m *DB) DeleteLinkModel(id int64) error {
	sql := fmt.Sprintf("delete from linkmodel where id=%d", id)
	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (m *DB) ListLinkModels(modelID int64) ([]LinkModel, error) {
	linkModels := make([]LinkModel, 0)

	sql := fmt.Sprintf("select * from linkmodel where fromModelID=%d or toModelID=%d", modelID, modelID)
	if err := m.Select(&linkModels, sql); err != nil {
		return linkModels, err
	}
	return linkModels, nil
}

// LinkEntity

func (m *DB) CreateLinkEntity(linkEntity LinkEntity) error {
	createdTime := time.Now().Format(TimeFormat)
	sql := fmt.Sprintf("insert into linkentity(linkModelID, fromEntityID, toEntityID, createdTime, lastUpdTime) "+
		"values(%d, %d, %d, '%s', '%s')", linkEntity.LinkModelID, linkEntity.FromEntityID, linkEntity.ToEntityID, createdTime, createdTime)

	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (m *DB) DeleteLinkEntity(id int64) error {
	sql := fmt.Sprintf("delete from linkentity where id=%d", id)
	if _, err := m.Exec(sql); err != nil {
		return err
	}
	return nil
}
