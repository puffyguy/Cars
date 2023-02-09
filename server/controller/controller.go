package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/puffyguy/Cars/connection"
	"github.com/puffyguy/Cars/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var JWTKey = []byte("secret_key")
var DAO *mongo.Database

func init() {
	DAO = connection.GetDB()
}

func RegisterUser(c *gin.Context) {
	var user model.User
	bindErr := c.Bind(&user)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"bind error": bindErr.Error()})
		return
	}
	u1 := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: HashAndSalt([]byte(user.Password)),
	}
	_, insertErr := DAO.Collection("userInfo").InsertOne(context.TODO(), u1)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"insert error": insertErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": u1})
}

func LoginUser(c *gin.Context) {
	var user model.User
	bindErr := c.Bind(&user)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"bind error": bindErr.Error()})
		return
	}
	u1 := model.User{
		Email:    user.Email,
		Password: user.Password,
	}
	filter := bson.M{"email": u1.Email}

	res := model.User{}
	findOneErr := DAO.Collection("userInfo").FindOne(context.TODO(), filter).Decode(&res)
	if findOneErr != nil {
		if findOneErr == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": findOneErr.Error()})
			return
		}
	}
	// fmt.Println("Value of user fetched from DB", u1)
	// c.JSON(http.StatusOK, res)
	validUserErr := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(u1.Password))
	if validUserErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	//JWT Implementation
	expirationTime := time.Now().Add(5 * time.Minute).Unix()
	claims := &model.Claims{
		Email: u1.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Create JWT string
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in", "token": tokenString, "tokenExp": claims.StandardClaims.ExpiresAt})
}

func CreateCar(c *gin.Context) {
	fmt.Println("Creating Car")
	var car model.Car
	bindErr := c.Bind(&car)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"bind error": bindErr.Error()})
		return
	}
	c1 := model.Car{
		Manufacturer: car.Manufacturer,
		Name:         car.Name,
		Count:        car.Count,
	}
	_, insertErr := DAO.Collection("carInfo").InsertOne(context.TODO(), c1)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"insert error": insertErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": c1})
}

func GetAllCars(c *gin.Context) {
	fmt.Println("Fetching all cars")
	var cars []model.Car
	filter := bson.D{}
	fetchRes, fetchErr := DAO.Collection("carInfo").Find(context.TODO(), filter)
	if fetchErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"fetch error": fetchErr.Error()})
		return
	}
	for fetchRes.Next(context.TODO()) {
		car := model.Car{}
		if err := fetchRes.Decode(&car); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"decode error": err.Error()})
			return
		}
		cars = append(cars, car)
	}
	c.JSON(http.StatusOK, gin.H{"message": cars})
}

func UpdateCar(c *gin.Context) {
	n := c.Param("name")
	// fmt.Println("id param:", n)
	var car model.Car
	bindErr := c.Bind(&car)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"bind error": bindErr.Error()})
		return
	}
	// c1 := model.Car{
	// 	Name:  car.Name,
	// 	Count: car.Count,
	// }
	filter := bson.M{"name": n}
	update := bson.M{"name": car.Name, "count": car.Count}
	updateRes, updateErr := DAO.Collection("carInfo").UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if updateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"update error": updateErr.Error()})
		return
	}

	var updatedCar model.Car
	if updateRes.MatchedCount == 1 {
		err := DAO.Collection("carInfo").FindOne(context.TODO(), bson.M{"name": car.Name}).Decode(&updatedCar)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"find error": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": updatedCar})
}

func DeleteCar(c *gin.Context) {
	n := c.Param("name")
	filter := bson.M{"name": n}
	deletedRes, deleteErr := DAO.Collection("carInfo").DeleteOne(context.TODO(), filter)
	if deleteErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"delete error": deleteErr.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": deletedRes.DeletedCount})
	}
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		fmt.Println("Error grnerating password hash:", err)
		return ""
	}
	return string(hash)
}
