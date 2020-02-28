package middlewares

import (
	"ChGo/db"
	"ChGo/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	Type         string
	Token        string
	RefreshToken string
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

type RefreshClaims struct {
	ID uuid.UUID
	jwt.StandardClaims
}

func Sign(c *gin.Context) {
	var login Login
	var user models.User

	db := db.GetDB()

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status: "Failure",
			Error:  err,
		})
		return
	}

	if err := db.Where("username = ?", login.Username).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
			Status: "Failure",
			Error:  err,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
			Status: "Failure",
			Error:  err,
		})
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	refreshExpirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	refreshClaims := &RefreshClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	// Create the JWT string
	tokenString, _ := token.SignedString(hmacSampleSecret)
	refreshTokenString, _ := refreshToken.SignedString(hmacSampleSecret)

	c.JSON(http.StatusOK, models.Response{
		Status: "Success",
		Data: Token{
			Type:         "Bearer",
			Token:        tokenString,
			RefreshToken: refreshTokenString,
		},
	})
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
}

// func AuthRequired(c *gin.Context) {

// }
