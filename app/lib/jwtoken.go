package lib

import (
	"errors"
	"log"
	"net/http"
	_ "strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("DreamDropDistance")

func GenerateJWT(usercode string) (string, error) {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))

	claims := sign.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["usercode"] = usercode // semisal id user atau code user
	claims["expired"] = time.Now().Add(time.Minute * 60 * 24).Unix()

	token, err := sign.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return token, nil
} // end func

func AuthJWT(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	/* jika menggunakan request header */
	// if r.Header["Token"] == nil {
	// 	return "", errors.New("Missing Header")
	// }
	// http_token := r.Header["Token"][0]

	/* jika menggunakan go session */
	http_token, _ := GetSession("jwt_token", w, r)
	if http_token == nil {
		return "", errors.New("JWT : Missing Header Request Token")
	}

	/* jika menggunakan cookie */
	// http_token, err := r.Cookie("jwt_token")
	// if err != nil {
	// 	return "", err
	// }

	token, err := jwt.Parse(http_token.(string), func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return "", errors.New("JWT : Unexpected signing method")
		}

		return mySigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			limit := claims["expired"]
			if time.Now().Unix() > int64(limit.(float64)) {
				log.Printf("%+v - %+v", time.Now().Unix(), int64(limit.(float64)))
				return "", errors.New("JWT : Token Expired")
			}
			return claims, nil
		} else {
			return "", errors.New("JWT : Invalid Params")
		}
	} else {
		return "", errors.New("JWT : Token Invalid")
	}
}
