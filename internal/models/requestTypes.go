package models

type AccountRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateRequest struct {
	Oldusername string `json:"oldusername" binding:"required"`
	Oldpassword string `json:"oldpassword" binding:"required"`
	Newusername string `json:"newusername" binding:"required"`
	Newpassword string `json:"newpassword" binding:"required"`
}
