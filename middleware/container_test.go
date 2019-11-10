package middleware

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/biguatch/msservice"
)

type MockJwt struct {
	mock.Mock
}

func (m *MockJwt) Generate(id string, isAdmin bool) (string, error) {
	return "", nil
}
func (m *MockJwt) Validate(str string) (jwt.MapClaims, error) {
	return jwt.MapClaims{}, nil
}

func TestContainer(t *testing.T) {
	a := assert.New(t)
	c := NewContainer(&msservice.Service{}, new(MockJwt))
	a.True(c.service != nil)
}
