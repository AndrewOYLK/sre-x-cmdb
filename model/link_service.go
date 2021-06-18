package model

// LinkType

func CreateLinkType(linkType LinkType) error {
	if err := mydb.CreateLinkType(linkType); err != nil {
		return err
	}
	return nil
}

func DeleteLinkType(id int64) error {
	if err := mydb.DeleteLinkType(id); err != nil {
		return err
	}
	return nil
}

func UpdateLinkType(linkType LinkType) error {
	if err := mydb.UpdateLinkType(linkType); err != nil {
		return err
	}
	return nil
}

func ListLinkTypes() ([]LinkType, error) {
	linkTypes, err := mydb.ListLinkTypes()
	if err != nil {
		return nil, err
	}
	return linkTypes, nil
}

// LinkModel

func CreateLinkModel(linkModel LinkModel) error {
	if err := mydb.CreateLinkModel(linkModel); err != nil {
		return err
	}
	return nil
}

func DeleteLinkModel(id int64) error {
	if err := mydb.DeleteLinkModel(id); err != nil {
		return err
	}
	return nil
}

func ListLinkModels(modelID int64) ([]LinkModel, error) {
	linkModels, err := mydb.ListLinkModels(modelID)
	if err != nil {
		return linkModels, err
	}
	return linkModels, nil
}

// LinkEntity

func ValidAndSaveLinkEntity(linkEntity LinkEntity) error {
	if err := mydb.CreateLinkEntity(linkEntity); err != nil {
		return err
	}
	return nil
}

func DeleteLinkEntity(id int64) error {
	if err := mydb.DeleteLinkEntity(id); err != nil {
		return err
	}
	return nil
}

/*
	mapping的作用是用来检查源ID和目标ID之间的约束关系
	情况1 - 1:1
	情况2 - 1:N
	情况3 - N:N
*/

type MappingInterface interface {
	Valid(LinkEntity) error
}

func NewMappingValidation(mapping string) MappingInterface {
	switch mapping {
	case "1:1":
		return &OneToOne{}
	case "1:n":
		return &OneToMany{}
	case "n:n":
		return &ManyToMany{}
	}
	return nil
}

// OneToOne 1:1
type OneToOne struct{}

func (o *OneToOne) Valid(linkEntity LinkEntity) error {
	// TODO
	return nil
}

// OneToMany 1:N
type OneToMany struct{}

func (o *OneToMany) Valid(linkEntity LinkEntity) error {
	// TODO
	return nil
}

// ManyToMany N:N
type ManyToMany struct{}

func (m *ManyToMany) Valid(linkEntity LinkEntity) error {
	// TODO
	return nil
}
