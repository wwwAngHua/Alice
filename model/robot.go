package model

type Chat struct {
	Uid     string  `json:"uid"`
	Version float32 `json:"version"`
	Message string  `json:"message"`
}
