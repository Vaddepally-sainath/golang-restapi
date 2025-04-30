package Authentication

import ("github.com/dgrijalva/jwt-go"
"time"
)

var jwtKey = []byte("Bahubali")

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(), // token expires in 2 hours
	})
	return token.SignedString(jwtKey)
}