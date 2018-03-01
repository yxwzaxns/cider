package server

type CreateProjectReq struct {
	ProjectURL string `json:"projectURL" binding:"required"`
}
