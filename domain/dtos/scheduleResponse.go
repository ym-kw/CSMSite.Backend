package dtos

type ScheduleResponse struct {
	Id            int    `column:"id"`
	StartDateTime string `column:"startDateTime"`
	EndDateTime   string `column:"endDateTime"`
	UserId        int    `column:"userId"`
	DisableFlag   bool   `column:"disableFlag"`
	CreatedAt     string `column:"createdAt"`
	UpdatedAt     string `column:"updatedAt"`
}
