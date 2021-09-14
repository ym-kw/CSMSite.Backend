package entities

// お知らせ
type Notification struct {
	Id             int    `column:"id" gorm:"primary_key"`
	Title          string `column:"title"`
	Text           string `column:"text"`
	TargetId       *int   `column:"targetId"`
	ContuributorId int    `column:"contributorId"`
	DisableFlag    bool   `column:"disableFlag"`
	CreatedAt      string `column:"createdAt"`
	UpdatedAt      string `column:"updatedAt"`
}
