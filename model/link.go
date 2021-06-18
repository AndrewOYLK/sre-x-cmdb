package model

/*
	example:
	ID	Key			StoD		DtoS		CreatedTime						LastUpdTime
	1   belong	属于		包含		 2021-05-30 11:53:20   2021-05-30 11:53:20
	2		connect	上联		下联		 2021-05-30 11:53:20   2021-05-30 11:53:20
	3   run			运行于	 运行		 2021-05-30 11:53:20   2021-05-30 11:53:20
	4   group		组成		组成于	 2021-05-30 11:53:20   2021-05-30 11:53:20
*/

// LinkType 关联类型（全局）
type LinkType struct {
	ID          int64  `json:"id,omitempty" db:"id"`
	Key         string `json:"key,omitempty" db:"key"`
	From        string `json:"from,omitempty" db:"from"` // 源模型->目标模型的描述
	To          string `json:"to,omitempty" db:"to"`     // 目标模型->源模型的描述
	CreatedTime string `json:"createdTime,omitempty" db:"createdTime"`
	LastUpdTime string `json:"lastUpdTime,omitempty" db:"lastUpdTime"`
}

/*
	LinkModel里的mapping决定关联实例这个动作要做的“验证”，譬如:
	1. src:dst --> 1:1
		源实例在关联目标实例的时候，如果LinkModel实例是1:1的mapping关系，那么该源在LinkEntity表中SrcEntityID中只能出现一次
	2. src:dst --> 1:N
		一个源可以绑定多个目标，但是一个目标不能被多个源绑定，即目标里的实例只能被一个源实例绑定
	3. src:dst --> N:N
		一个源实例可以绑定多个目标实例，一个目标实例又同时可以被多个源实例绑定
*/

// LinkModel 模型关联（模型有哪些关联）
type LinkModel struct {
	ID          int64  `json:"id,omitempty" db:"id"`
	LinkTypeID  int64  `json:"linkTypeID,omitempty" db:"linkTypeID"`   // 关联类型
	FromModelID int64  `json:"fromModelID,omitempty" db:"fromModelID"` // 源模型
	ToModelID   int64  `json:"toModelID,omitempty" db:"toModelID"`     // 目标模型
	Mapping     string `json:"mapping,omitempty" db:"mapping"`         // 约束关系（1:1、1:N、N:N）
	CreatedTime string `json:"createdTime,omitempty" db:"createdTime"`
	LastUpdTime string `json:"lastUpdTime,omitempty" db:"lastUpdTime"`
}

func (l *LinkModel) Valid(linkentity LinkEntity) error {
	mapping := NewMappingValidation(l.Mapping)
	if err := mapping.Valid(linkentity); err != nil {
		return err
	}
	return nil
}

// LinkEntity 模型关联实例
type LinkEntity struct {
	ID           int64  `json:"id,omitempty" db:"id"`
	LinkModelID  int64  `json:"linkModelID,omitempty" db:"linkModelID"`
	FromEntityID int64  `json:"fromEntityID,omitempty" db:"fromEntityID"` // 源模型记录ID
	ToEntityID   int64  `json:"toEntityID,omitempty" db:"toEntityID"`     // 目标模型记录ID
	CreatedTime  string `json:"createdTime,omitempty" db:"createdTime"`
	LastUpdTime  string `json:"lastUpdTime,omitempty" db:"lastUpdTime"`
}
