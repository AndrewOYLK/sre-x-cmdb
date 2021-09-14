package model

/*
	example:
	ID	ModelID
	1		1
	2		1
	3		2
	4		2
*/

// Entity 实例
type Entity struct {
	ID int64 `json:"id,omitempty" db:"id"`
	// Key         string `json:"key,omitempty" db:"id"`
	ModelID     int64  `json:"modelID,omitempty" db:"modelID"`
	CreatedTime string `json:"createdTime,omitempty" db:"createdTime"`
	LastUpdTime string `json:"lastUpdTime,omitempty" db:"lastUpdTime"`
}
