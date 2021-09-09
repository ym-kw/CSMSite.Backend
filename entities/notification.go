package entities

// お知らせ
type Notification struct {
	Id             int    `column:"id"`
	Title          string `column:"title"`
	Text           string `column:"text"`
	Type           int    `column:"type"`
	Users          *Users
	ContuributorId int    `column:"contributorId"`
	DisableFlag    bool   `column:"disableFlag"`
	CreatedAt      string `column:"createdAt"`
	UpdatedAt      string `column:"updatedAt"`
}
