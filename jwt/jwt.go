package jwt

import (
	"time"
	"twister/app/models"

	jwtoken "github.com/dgrijalva/jwt-go"
)

// GenerateJWT return a token
func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("CodigoCorrectoTwister")

	payload := jwtoken.MapClaims{
		"email":    user.Email,
		"name":     user.Name,
		"lastname": user.Lastname,
		"birthDay": user.BirthDay,
		"biografy": user.Biografy,
		"location": user.Location,
		"webSite":  user.WebSite,
		"_id":      user.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwtoken.NewWithClaims(jwtoken.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
