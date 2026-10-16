package jwt_test

import (
	"http-server/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "dev@dev.com"
	jwtService := jwt.NewJWT("NduPTVlWkk6w97eSp1BlXywF31lgxiWhQtrjjOxOAju")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatalf("token is invalid")
	}
	if data.Email != email {
		t.Fatalf("email %s not equal %s", data.Email, email)
	}
}
