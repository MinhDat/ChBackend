package middlewares

import (
	"ChGo/db"
	"ChGo/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Error message code
const (
	FORBIDDEN    = "FORBIDDEN"
	UNAUTHORIZED = "UNAUTHORIZED"
)

// Declare the expiration time of the token
// here, we have kept it as 5 minutes
var expirationTime = time.Now().Add(60 * time.Minute)
var refreshExpirationTime = time.Now().Add(24 * time.Hour)

var UnAuthorizedResponse = models.Response{
	Status: "Failure",
	Error:  UNAUTHORIZED,
}

// Token "Object for responding"
type Token struct {
	Type         string
	Token        string
	RefreshToken string
}

// Login "Object for Sign request"
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
		c.AbortWithStatusJSON(http.StatusForbidden, models.Response{
			Status: "Failure",
			Error:  FORBIDDEN,
		})
		return
	}

	if err := db.Where("username = ?", login.Username).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnAuthorizedResponse)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnAuthorizedResponse)
		return
	}

	// Create the JWT token
	tokenString, refreshTokenString := generateToken(user)

	c.JSON(http.StatusOK, models.Response{
		Status: "Success",
		Data: Token{
			Type:         "Bearer",
			Token:        tokenString,
			RefreshToken: refreshTokenString,
		},
	})
}

func Auth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if len(tokenString) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnAuthorizedResponse)
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err == nil && token.Valid {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnAuthorizedResponse)
	}

}

func Refresh(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if len(tokenString) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
			Status: "Failure",
			Error:  UNAUTHORIZED,
		})
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(*RefreshClaims); ok && err == nil && token.Valid {
		var user models.User

		db := db.GetDB()
		if err := db.Where("id = ?", claims.ID).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, UnAuthorizedResponse)
			return
		}

		// Create the JWT token
		tokenString, refreshTokenString := generateToken(user)

		c.JSON(http.StatusOK, models.Response{
			Status: "Success",
			Data: Token{
				Type:         "Bearer",
				Token:        tokenString,
				RefreshToken: refreshTokenString,
			},
		})

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnAuthorizedResponse)
	}
}

func generateToken(user models.User) (tokenString string, refreshTokenString string) {
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
	tokenString, _ = token.SignedString(hmacSampleSecret)
	refreshTokenString, _ = refreshToken.SignedString(hmacSampleSecret)
	return
}
