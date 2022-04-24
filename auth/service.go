package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("KITABANTU_S3CR3T_K3Y")

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}