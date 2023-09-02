package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/puffyguy/Cars/controller"
	"github.com/puffyguy/Cars/model"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	//User Routes
	router.POST("/register", controller.RegisterUser)
	router.POST("/login", controller.LoginUser)

	// router.GET("/admin", isAuthenticated, func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Admin section"})
	// })

	//Car Routes
	router.POST("/createcar", controller.CreateCar)
	router.GET("/getcars", isAuthenticated, controller.GetAllCars)
	router.PUT("/updatecar/:name", controller.UpdateCar)
	router.DELETE("/deletecar/:name", controller.DeleteCar)

	router.Run("localhost:8000")
	// log.Println("Server listening at:")
}

func isAuthenticated(c *gin.Context) {
	fmt.Println("Is Authenticated method")
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		fmt.Println("Empty token")
		c.JSON(http.StatusUnauthorized, gin.H{"status": "token not present"})
		c.Abort()
	}
	claims := &model.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return controller.JWTKey, nil
	})
	if err != nil {
		fmt.Println("Invalid token while checking")
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			c.Abort()
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": "cannot parse token with claims"})
		c.Abort()
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid token"})
		return
	}
}
