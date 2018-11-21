package core

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

func GetMockJWTToken(gameId int) *jwt.Token {
	return &jwt.Token{
		Claims: jwt.MapClaims{
			"v1": map[string]interface{}{
				"claims": map[string]interface{}{
					"gameIds": []interface{}{
						strconv.Itoa(gameId),
					},
				},
			},
		},
	}
}
