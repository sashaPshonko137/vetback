package user

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *userService) GenerateToken(info model.Claims) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := model.Claims{
		UserId:         info.UserId,
		FullName:       info.FullName,
		Address:        info.Address,
		PhoneNumber:    info.PhoneNumber,
		Role:           info.Role,
		Specialization: info.Specialization,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(s.config.SecretKey)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		s.logger.Error(err)
		return "", customError.InternalServerError
	}

	return signedToken, nil
}

func (s *userService) getInfoFromToken(token string) (*model.Claims, error) {
	secretKey := []byte(s.config.SecretKey)

	parsedToken, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, customError.InvalidToken
	}

	if !parsedToken.Valid {
		return nil, customError.InvalidToken
	}

	claims, ok := parsedToken.Claims.(*model.Claims)
	if !ok {
		return nil, customError.InvalidToken
	}

	return claims, nil
}
