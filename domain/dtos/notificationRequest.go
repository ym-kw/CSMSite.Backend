package Dtos

type NotificationRequest struct {
	Title          string `column:"title"`
	Text           string `column:"text"`
	ContuributorId int    `column:"contributorId"`
}
