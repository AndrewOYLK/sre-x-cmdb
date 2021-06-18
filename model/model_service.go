package model

func CreateModel(model Model) error {
	return mydb.CreateModel(model)
}

func DeleteModel(id int64) error {
	return mydb.DeleteModel(id)
}

func UpdateModel(model Model) error {
	return mydb.UpdateModel(model)
}

func GetModel(id int64) (Model, error) {
	return mydb.GetModel(id)
}

func ListModel() ([]Model, error) {
	return mydb.ListModel()
}
