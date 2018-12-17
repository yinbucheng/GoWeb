package entity

type ClassRoom struct {
	ID        int    `gorm:"primary_key" json:"id"`
	ClassRoom string `gorm:"type:varchar(20);not null" json:"classRoom"`
	Address   string `gorm:"type:varchar(20);not null" json:"address"`
}
