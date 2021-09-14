package entities

// ユーザー
type User struct {
	Id           int             `column:"id" gorm:"primary_key"`
	UserName     string          `column:"userName"`
	Password     string          `column:"password"`
	Email        string          `column:"email"`
	Admin        bool            `column:"admin"`
	DisableFlag  bool            `column:"disableFlag"`
	CreatedAt    string          `column:"createdAt"`
	UpdatedAt    string          `column:"updatedAt"`
	Notification *[]Notification `column:"notification" gorm:"foreignKey:TargetId"`
	Schedule     *[]Schedule     `column:"schedule" gorm:"foreingKey:UserId"`
}
