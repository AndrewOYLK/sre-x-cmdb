package store

// Store 存储接口
type Store interface {
	ModelGroup
	Model
	AttributeGroup
	Attribute
	LinkRelation
	Resource
}

type ModelGroup interface{}

type Model interface{}

type AttributeGroup interface{}

type Attribute interface{}

type LinkRelation interface{}

type Resource interface{}
