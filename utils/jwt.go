package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	Uid    uint
	Type   string
	TypeID uint
}

func GenerateToken(claim JWTClaim) (string, error) {
	token_lifetime, _ := strconv.Atoi(GetEnv("TOKEN_LIFETIME"))

	claims := jwt.MapClaims{}
	claims["uid"] = uint(claim.Uid)
	claims["type"] = claim.Type
	claims["type_id"] = claim.TypeID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(token_lifetime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetEnv("JWT_KEY")))
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(GetEnv("JWT_KEY")), nil
	})
}

func ValidateToken(tokenString string) error {
	token, err := ParseToken(tokenString)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return fmt.Errorf("Unauthenticated")
}

func ExtractToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func ExtractTokenClaim(tokenString string) (JWTClaim, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return JWTClaim{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var claim JWTClaim
		uid, _ := claims["uid"].(uint)
		utype, _ := claims["type"].(string)
		typeID, _ := claims["type_id"].(uint)
		claim.Uid = uid
		claim.Type = utype
		claim.TypeID = typeID
		return claim, nil
	}
	return JWTClaim{}, nil
}