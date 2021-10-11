package Dtos

type UserResponse struct {
	Id          int    `column:"id"`
	UserName    string `column:"userName"`
	Password    string `column:"password"`
	Email       string `column:"email"`
	Admin       bool   `column:"admin"`
	DisableFlag bool   `column:"disableFlag"`
	CreatedAt   string `column:"createdAt"`
	UpdatedAt   string `column:"updatedAt"`
}
