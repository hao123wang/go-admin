package jwt

import (
	"errors"
	"go-admin-server/api/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenExpireDuration = 24 * time.Hour
	TokenSecret         = "this-is-go-admin-server"
)

var (
	ErrAbsent  = errors.New("token absent")  // token不存在
	ErrInvalid = errors.New("token invalid") // token无效
)

type CustomClaims struct {
	entity.JwtAdmin
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.SysAdmin) (string, error) {
	claims := &CustomClaims{
		JwtAdmin: entity.JwtAdmin{
			ID:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Icon:     user.Icon,
			Email:    user.Email,
			Phone:    user.Phone,
			Note:     user.Note,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(TokenSecret))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(TokenSecret), nil
	})
	if err != nil {
		return nil, ErrInvalid
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalid
}
