package entity

type HelloRequest struct {
	Name string `json:"name" binding:"required"`
}
