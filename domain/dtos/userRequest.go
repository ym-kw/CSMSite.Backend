package dtos

type UserRequest struct {
	UserName string `column:"userName"`
	Password string `column:"password"`
	Email    string `column:"email"`
}
