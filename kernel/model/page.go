package model

type Page struct {
	Data  any   `json:"data"`  // data object
	Page  int64 `json:"page"`  // current page
	Size  int64 `json:"size"`  // page size
	Total int64 `json:"total"` // total
}
