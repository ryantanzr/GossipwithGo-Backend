package models

type AccountRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateRequest struct {
	ID          int    `json:"id"`
	Newusername string `json:"newusername" binding:"required"`
	Newpassword string `json:"newpassword" binding:"required"`
}
