package entity

type User struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Name   string `gorm:"type:varchar(20);not null" json:"name"`
	Age    int    `gorm:"type:int" json:"age"`
	Gender string `gorm:"type:varchar(20)" json:"gender"`
}
