package server

type CreateProjectReq struct {
	ProjectURL string `json:"projectURL" binding:"required"`
}

type AuthReq struct {
	Key string `json:"key" binding:"required"`
}
