package models

type UserAction struct {
	ActorID    int    `json:"actorID"`
	ReceiverID int    `json:"receiverID"`
	ActionType string `json:"actionType"`
}
