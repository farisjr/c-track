package middlewares

import (
	"app/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//create token with adding limit time
func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenPatientId(e echo.Context) int {
	patient := e.Get("user").(*jwt.Token)

	if patient.Valid {
		claims := patient.Claims.(jwt.MapClaims)
		patientId := int(claims["userId"].(float64))
		return patientId
	}
	return 0
}
