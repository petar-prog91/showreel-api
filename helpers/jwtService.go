package helpers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var tokenSigningKey = []byte("y42jh9824j9h82j49h82j40g9im240h9240h94p2hjk0249h0924")

type MyJWTClaims struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Sgroup   int    `json:"sGroup"`
	jwt.StandardClaims
}

func GenerateNewToken(userId int, userName string, sGroup int) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userId,
		"userName": userName,
		"sGroup":   sGroup,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(tokenSigningKey)

	return tokenString, err
}

func ParseToken(myToken string) (bool, *MyJWTClaims, error) {
	token, err := jwt.ParseWithClaims(myToken, &MyJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSigningKey, nil
	})

	if claims, ok := token.Claims.(*MyJWTClaims); ok && token.Valid {
		return true, claims, nil
	} else {
		fmt.Println(err)
		return false, nil, err
	}
}

func JwtAuth(h httprouter.Handle, reqUserRole int) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var jwtToken = r.Header["Auth_jwt_token"]

		if len(jwtToken) > 0 {
			var validToken, claims, err = ParseToken(jwtToken[0])

			if err != nil {
				StatusUnauthorized(w)
			}

			var userID = claims.Id
			var jwtUserRole = claims.Sgroup
			var paramID = ps.ByName("id")

			var isAdmin = checkIfAdmin(jwtUserRole)
			var canAccessRegularUser = checkUserId(paramID, string(userID))

			if validToken && isAdmin {
				h(w, r, ps)
			} else if validToken && canAccessRegularUser {
				h(w, r, ps)
			} else {
				StatusUnauthorized(w)
			}
		} else {
			// Request Basic Authentication otherwise
			StatusUnauthorized(w)
		}
	}
}

func checkIfAdmin(jwtUserRole int) bool {
	if jwtUserRole == 0 {
		return true
	}

	return false
}

func checkUserId(requestId string, userId string) bool {
	if userId == requestId {
		return true
	}

	return false
}
