package models

import "github.com/golang-jwt/jwt/v4"

type Token struct {
	GivenName string `json:"GivenName"`
	Surname   string `json:"Surname"`
	Email     string `json:"Email"`
	Role      string `json:"Role"`
	Team      string `json:"Team"`
	UUID      string `json:"UUID"`
	jwt.RegisteredClaims
}
