package model

type pageQuery struct {
	Current int64 `json:"current"`
	Size    int64 `json:"size"`
}
