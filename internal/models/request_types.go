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

type PostRequest struct {
	ID      int    `json:"id"`
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Content string `json:"content" binding:"required"`
	Likes   int    `json:"likes"`
}

type UserActionRequest struct {
	ActorID    int    `json:"actorID" binding:"required"`
	ReceiverID int    `json:"receiverID" binding:"required"`
	ActionType string `json:"actionType" binding:"required"`
}
