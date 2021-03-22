package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("Golang-STARTUP-s3cretK3y")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	//set payload
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	//set algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
