package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/santihill/crud_go/initializers"
	"github.com/santihill/crud_go/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash",
		})
		return
	}
	//Create the user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	//Respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	//Get the email and pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email) // SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//Generate a jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // Create a new token object, specifying signing method and the claims
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) // Sign and get the complete encoded token as a string using the secret
	fmt.Println(tokenString, err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	//send ir back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	//user.(models.User).ID     //Para utilizar el usuario y sus datos

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})

}
