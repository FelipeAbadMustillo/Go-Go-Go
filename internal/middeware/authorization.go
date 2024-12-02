package middeware

import (
	"errors"
	"time"

	"github.com/Go-Go-Go/api"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

var jwtKey = []byte("secret_key") // Cambiar esto por una clave segura

// Claim structure
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var UnAuthorizedError = errors.New("Invalid username or token.") // Fijarse de cambiar esto con los errores

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(api.JWT_EXPIRATION_MINUTES * time.Minute)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func JWTAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" { // "Missing token"
			context.Abort()
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(context.Writer, UnAuthorizedError) //Agregar otro tipo de error de sesion?
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid { // "Invalid token"
			context.Abort()
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(context.Writer, UnAuthorizedError) //Agregar otro tipo de error de sesion?
			return
		}

		// Save user data in context for later usage
		context.Set("username", claims.Username)
		context.Next()
	}
}
