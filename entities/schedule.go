package Entities

// 予定表
type Schedule struct {
	Id            int    `column:"id" gorm:"primary_key"`
	StartDateTime string `column:"startDateTime"`
	EndDateTime   string `column:"endDateTime"`
	UserId        int    `column:"userId"`
	DisableFlag   bool   `column:"disableFlag"`
	CreatedAt     string `column:"createdAt"`
	UpdatedAt     string `column:"updatedAt"`
}
