package model

// Attribute

func CreateAttribute(attr Attribute) error {
	return mydb.CreateAttribute(attr)
}

func DeleteAttribute(id int64) error {
	return mydb.DeleteAttribute(id)
}

func UpdateAttribute(attr Attribute) error {
	return mydb.UpdateAttribute(attr)
}

func GetAttribute(id int64) (Attribute, error) {
	return mydb.GetAttribute(id)
}

func ListAttribute(modelID int64) ([]Attribute, error) {
	return mydb.ListAttribute(modelID)
}

// UniqueAttrs

func CreateUniqueAttrs(uniqueAttrs UniqueAttrs) error {
	return mydb.CreateUniqueAttrs(uniqueAttrs)
}

func DeleteUniqueAttrs(id int64) error {
	return mydb.DeleteAttribute(id)
}

func UpdateUniqueAttrs(uniqueAttrs UniqueAttrs) error {
	return mydb.UpdateUniqueAttrs(uniqueAttrs)
}

func ListUniqueAttrs(modelID int64) ([]UniqueAttrs, error) {
	return mydb.ListUniqueAttrs(modelID)
}
