package param

import "mime/multipart"

type Base64EncodeParam struct {
	Data string                `form:"data"`
	File *multipart.FileHeader `form:"file"`
	Type string                `form:"type"`
}

type Base64DecodeParam struct {
	Data string `from:"data"`
}
