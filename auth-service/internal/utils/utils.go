package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"github.com/golang-jwt/jwt/v5"
		"github.com/gin-gonic/gin"
		"strings"
)


const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"


func IntToBase62(num int64) string {
	var result string
	for num > 0 {
		remainder := num % 62
		result = string(base62Alphabet[remainder]) + result
		num /= 62
	}
	return result
}

func Base62ToInt(base62 string) (int64, error) {
	var result int64
	for _, char := range base62 {
		var idx int
		for k, c := range base62Alphabet {
			if c == char {
				idx = k
				break
			}
		}
		result = result*62 + int64(idx)
	}
	return result, nil
}

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetSharedMetadataTableName(id int64) string {
    return fmt.Sprintf("timesheetmetadata%d", id%4)
}



var jwtSecret = []byte("abcd") // Replace with your secret

func GenerateJWT(userID uint, expiryHours int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(expiryHours) * time.Hour).Unix(),
		"secret":  string(jwtSecret),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims")
	}
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found in token")
	}
	return int64(userIDFloat), nil
}




func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Assuming the format is: "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		userID, err := ValidateJWT(tokenParts[1])
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		// Store userID in context for further handlers to use
		c.Set("userID", userID)

		c.Next()
	}
}