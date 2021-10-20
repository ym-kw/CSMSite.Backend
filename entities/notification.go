package Entities

// お知らせ
type Notification struct {
	Id             int    `column:"id" gorm:"primary_key"`
	Title          string `column:"title"`
	Text           string `column:"text"`
	ContuributorId int    `column:"contributorId"`
	DisableFlag    bool   `column:"disableFlag"`
	CreatedAt      string `column:"createdAt"`
	UpdatedAt      string `column:"updatedAt"`
}
