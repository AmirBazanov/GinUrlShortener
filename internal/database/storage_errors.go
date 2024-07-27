package database

import "errors"

var (
	UrlAlreadyExist = errors.New("URL_ALREADY_EXIST")
	AliasNotFound   = errors.New("Alias_NOT_FOUND")
	TokenNotFound   = errors.New("Token_NOT_FOUND")
)
