package dtos

type NotificationRequest struct {
	Title          string `column:"title"`
	Text           string `column:"text"`
	Target         *int   `column:"target"`
	ContuributorId int    `column:"contributorId"`
}
