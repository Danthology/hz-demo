// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package mysql

const TableNameMsg = "msg"

// Msg mapped from table <msg>
type Msg struct {
	ID   int32   `gorm:"column:Id;type:int(11);primaryKey;autoIncrement:true" json:"Id"`
	Num  *int32  `gorm:"column:num;type:int(11)" json:"num"`
	Code *string `gorm:"column:code;type:varchar(255)" json:"code"`
}

// TableName Msg's table name
func (*Msg) TableName() string {
	return TableNameMsg
}
