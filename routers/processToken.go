package routers

import (
	"errors"
	"strings"

	"twister/app/db"
	"twister/app/models"

	jwtToken "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUser string

// checkig token sent
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	mykey := []byte("CodigoCorrectoTwister")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("token format invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwtToken.ParseWithClaims(token, claims, func(t *jwtToken.Token) (interface{}, error) {
		return mykey, nil
	})

	if err == nil {
		_, found, _ := db.CheckUserExists(claims.Email)
		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
