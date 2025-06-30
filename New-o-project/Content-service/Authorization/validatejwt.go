package Authorization

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
// 			return
// 		}

// 		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("Bahubali"), nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			return
// 		}

// 		c.Next()
// 	}
// }
