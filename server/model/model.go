package model

import "github.com/dgrijalva/jwt-go"

type Car struct {
	Manufacturer string `json:"manufacturer" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Count        int    `json:"count" binding:"required"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
