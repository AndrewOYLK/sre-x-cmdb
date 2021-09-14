package model

/*
	example:
	ID 	Key					Description CreatedTime 					LastUpdTime
	1		host				主机表			2021-05-30 11:53:20   2021-05-30 11:53:20
	2		application	应用表			2021-05-30 11:53:20   2021-05-30 11:53:20
*/

// Model 模型
type Model struct {
	ID          int64  `json:"id,omitempty" db:"id"`
	Key         string `json:"key,omitempty" db:"key"`
	Description string `json:"description,omitempty" db:"description"`
	CreatedTime string `json:"createdTime,omitempty" db:"createdTime"`
	LastUpdTime string `json:"lastUpdTime,omitempty" db:"lastUpdTime"`
}
