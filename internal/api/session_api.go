package api

type CreateSessionReqBody struct {
	Email string `json:"email" binging:"required"`
	Code  string `json:"code" binging:"required"`
}

type CreateSessionResBody struct {
	JWT    string `json:"jwt"`
	UserId int32  `json:"userId"`
}
