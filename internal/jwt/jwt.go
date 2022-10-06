package jwt

import (
	"fmt"
	"go-template/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service struct {
     key []byte
     ttl time.Duration
     algo jwt.SigningMethod
}

func New(algo, secret string, ttlMinutes, minSecretLength int) (Service, error){
     
     var minSecretLen = 128

     if minSecretLength > 0 {
          minSecretLen = minSecretLength
     }
     if len(secret) < minSecretLen {
          return Service{}, fmt.Errorf("jwt secret length is %v, which is less than required %v", len(secret), minSecretLen)
     }
     signingMethod := jwt.GetSigningMethod(algo)
     if signingMethod == nil {
          return Service{}, fmt.Errorf("invalid jwt signing method: %s", algo)
     }
     return Service{
          key: []byte(secret),
          ttl: time.Duration(ttlMinutes) * time.Minute,
          algo: signingMethod,
     }, nil
}

func (s Service) ParseToken(authHeader string) (*jwt.Token, error) {
     parts := strings.SplitN(authHeader, " ", 2)
     if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
          return nil, fmt.Errorf("Invalid token")
     }      
     return jwt.Parse(parts[1], func(token *jwt.Token)(interface{}, error){
          if s.algo != token.Method {
               return nil, fmt.Errorf("Invalid token method")
          }
          return s.key, nil
     })
}

func (s Service) GenerateToken(a *models.Author) (string, error) {
     return jwt.NewWithClaims(s.algo, jwt.MapClaims{
          "id": a.ID,
          "u": a.Username,
          "exp": time.Now().Add(s.ttl).Unix(),
     }).SignedString(s.key)
}