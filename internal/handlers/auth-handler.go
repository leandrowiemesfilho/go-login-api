package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leandrowiemesfilho/login-api/internal/db"
	"github.com/leandrowiemesfilho/login-api/internal/exceptions"
	"github.com/leandrowiemesfilho/login-api/internal/models"
	"github.com/leandrowiemesfilho/login-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const operationTimeout = 60
const userCollectionName = "user"

func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout*time.Second)
	defer cancel()

	var loginRequest models.LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var user models.User
	userCollection := db.GetCollection(userCollectionName)

	if err := userCollection.FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&user); err != nil {
		userNotFoundErr := exceptions.UserNotFoundErr{UserInfo: loginRequest.Email}
		c.JSON(http.StatusNotFound, gin.H{"error": userNotFoundErr.Error()})
		return
	}

	err := utils.VerifyPassword(user.Password, loginRequest.Password)
	if err != nil {
		invalidCredentialsErr := exceptions.InvalidCredentialsErr{}
		c.JSON(http.StatusInternalServerError, gin.H{"error": invalidCredentialsErr.Error()})
		return
	}

	c.Header("auth-token", utils.CreateJWT(&user))
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Signup(c *gin.Context) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), operationTimeout*time.Second)
	defer cancelFunc()

	var signUpRequest models.SignUpRequest

	if err := c.BindJSON(&signUpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userCollection := db.GetCollection(userCollectionName)

	userCounter, err := userCollection.CountDocuments(ctx, bson.M{"email": signUpRequest.Email})
	if err != nil || userCounter > 0 {
		userAlreadyExistsErr := exceptions.UserAlreadyExistsErr{UserInfo: signUpRequest.Email}
		c.JSON(http.StatusConflict, gin.H{"error": userAlreadyExistsErr.Error()})
	}

	user := &models.User{
		Id:           primitive.NewObjectID(),
		FirstName:    signUpRequest.FirstName,
		LastName:     signUpRequest.LastName,
		Email:        signUpRequest.Email,
		PhoneNumber:  signUpRequest.PhoneNumber,
		Password:     utils.HashPassword(signUpRequest.Password),
		CreationDate: time.Now(),
		UpdateDate:   time.Now(),
	}

	userCreated, err := userCollection.InsertOne(ctx, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"user": userCreated})
}
