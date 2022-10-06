package jwt_test

import (
	"fmt"
	"go-template/internal/jwt"
	"go-template/models"
	"strings"
	"testing"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

func TestNew(t *testing.T) {
     cases := map[string]struct {
          algo string
          secret string
          minSecretLen int
          req models.Author
          wantResp jwt.Service
          wantErr bool
          error string
     }{
          "invalid_algo": {
               algo: "invalid",
               wantErr: true,
               minSecretLen: 1,
               secret: "g0r$kt3$t1ng",
               error: "invalid jwt signing method: invalid",
          },
          "invalid secret length": {
               algo: "HS256",
               secret: "123",
               wantErr: true,
               error: "jwt secret length is 3, which is less than required 128",
          },
          "invalid secret length with min defined": {
               algo: "HS256",
               minSecretLen: 4,
               secret: "123",
               wantErr: true,
               error: "jwt secret length is 3, which is less than required 4",
          },
          "success": {
               algo: "HS256",
               secret: "g0r$kt3$t1ng",
               minSecretLen: 1,
               req: models.Author{
                    Username: null.StringFrom("johndoe"),
               },
               wantResp: jwt.Service{},
          },
     }
     for name, tt := range cases {
          t.Run(name, func(t *testing.T){
               _, err := jwt.New(tt.algo, tt.secret, 60, tt.minSecretLen)
               assert.Equal(t, tt.wantErr, err != nil)
               if err != nil {
                    assert.Equal(t, tt.error, err.Error())
               }
          })
     }
}

func TestGenerateToken(t *testing.T){
     cases := map[string]struct {
          algo string
          secret string
          minSecretLen int
          req models.Author
          wantResp string
          wantErr bool
     }{
          "invalid algo": {
               algo: "invalid",
               wantErr: true,
          },
          "secret not set": {
               algo: "HS256",
               wantErr: true,
          },
          "invalid secret length": {
               algo: "HS256",
               secret: "123",
               wantErr: true,
          },
          "invalid secret length with min defined": {
               algo: "HS256",
               minSecretLen: 4,
               secret: "123",
               wantErr: true,
          },
          "success": {
               algo: "HS256",
               secret: "g0r$kt3$t1ng",
               minSecretLen: 1,
               req: models.Author{
                    Username: null.StringFrom("johndoe"),
               },
               wantResp: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
          },
     }
     for name, tt := range cases {
          t.Run(name, func(t *testing.T){
               jwtService, err := jwt.New(tt.algo, tt.secret, 60, tt.minSecretLen)
               assert.Equal(t, tt.wantErr, err != nil)
               if err == nil && !tt.wantErr {
                    jwtToken, _ := jwtService.GenerateToken(&tt.req)
                    assert.Equal(t, tt.wantResp, strings.Split(jwtToken, ".")[0])
               }
          })
     }
}

func TestParseToken(t *testing.T){
     cases := map[string]struct {
          algo string
          authHeader string
          error string
     }{
          "Failure_InvalidToken": {
               authHeader: "",
               algo: "HS256",
               error: "Invalid token",
          },
          "Failure_NoAuth": {
               authHeader: "",
               error: "Invalid token",
               algo: "HS256",
          },
          "Failure_MismatchTokenMethod": {
               authHeader: "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIi" +
               "wibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
               algo: "HS256",
               error: "signature is invalid",
          },
          "Success": {
               authHeader: "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIi" +
               "wibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
               algo: "HS256",
          },
     }
     for name, tt := range cases {
          t.Run(name, func(t *testing.T){
               fmt.Println("Name:", name)
               jwtService, err := jwt.New(tt.algo, "g0r$kt3$t1ng", 60, 1)
               if err != nil {
                    assert.Equal(t, tt.error, err.Error())
               }
               token, err := jwtService.ParseToken(tt.authHeader)
               if len(tt.error) != 0 {
                    assert.Equal(t, tt.error, err.Error())
               } else {
                    assert.Equal(t, "John Doe", token.Claims.(jwtgo.MapClaims)["name"])
               }
          })
     }
}