package entities

// 予定表
type Calendar struct {
	Id            int    `column:"id"`
	StartDateTime string `column:"startDateTime"`
	EndDateTime   string `column:"endDateTime"`
	DisableFlag   bool   `column:"disableFlag"`
	CreatedAt     string `column:"createdAt"`
	UpdatedAt     string `column:"updatedAt"`
	UserId        int    `column:"userId"`
	User          User
}

type Calenders []Calendar
