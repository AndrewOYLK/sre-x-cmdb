package model

/*
	example:
	ID	EntityID	AttributeID	Value
	1		1					1						vm161
	2		1					2						192.168.101.161
	3		1					3						kvm
	4		2					1						vm162
	5		2					2						192.168.101.162
	6		2					3						kvm
	7		3					4						serverA
	8		3					5						beijing
	9		3					6						andrew
	10	4					4						serverB
	11	4					5						xinye
	12	4					6						david
*/

// Value 实例值
type Value struct {
	ID          int64       `json:"id,omitempty" db:"id"`
	EntityID    int64       `json:"entityID,omitempty" db:"entityID"`
	AttributeID int64       `json:"attributeID,omitempty" db:"attributeID"`
	Value       interface{} `json:"value,omitempty" db:"value"`
	// CreatedTime string      `json:"createdTime,omitempty" db:"createdTime"`
	// LastUpdTime string      `json:"lastUpdTime,omitempty" db:"lastUpdTime"`
}
