package middlewares

import (
	"ChGo/db"
	"ChGo/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Token struct {
	Type    string
	Token   string
	Refresh string
}

// Login "Object
type Login struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

var hmacSampleSecret = []byte("ChStore")

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string
	jwt.StandardClaims
}

func Sign(c *gin.Context) {
	var login Login
	var user models.User

	c.BindJSON(&login)
	db := db.GetDB()
	if err := db.Where("Username = ?", login.Username).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: login.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, _ := token.SignedString(hmacSampleSecret)

	c.JSON(http.StatusOK, models.Response{
		Status: "Success",
		Data: Token{
			Type:  "Bearer",
			Token: tokenString,
		},
	})
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
}

func AuthRequired(c *gin.Context) {

}
