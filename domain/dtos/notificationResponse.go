package Dtos

type NotificationResponse struct {
	Id             int    `column:"id"`
	Title          string `column:"title"`
	Text           string `column:"text"`
	ContuributorId int    `column:"contributorId"`
	DisableFlag    bool   `column:"disableFlag"`
	CreatedAt      string `column:"createdAt"`
	UpdatedAt      string `column:"updatedAt"`
}
