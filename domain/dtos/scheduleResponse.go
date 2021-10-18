package Dtos

type ScheduleResponse struct {
	Id            int          `column:"id"`
	StartDateTime string       `column:"startDateTime"`
	EndDateTime   string       `column:"endDateTime"`
	UserId        int          `column:"userId"`
	User          UserResponse `column:"user"`
	DisableFlag   bool         `column:"disableFlag"`
	CreatedAt     string       `column:"createdAt"`
	UpdatedAt     string       `column:"updatedAt"`
}
