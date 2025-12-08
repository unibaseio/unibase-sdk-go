package types

type MemeStruct struct {
	Owner   string `json:"owner" binding:"required"`
	Bucket  string `json:"bucket" binding:"-"`
	ID      string `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type MemeMeta struct {
	File  string
	Start uint64
	Size  uint64
}
