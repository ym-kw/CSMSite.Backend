package Dtos

type ScheduleRequest struct {
	StartDateTime string `column:"startDateTime"`
	EndDateTime   string `column:"endDateTime"`
	UserId        int    `column:"userId"`
}
