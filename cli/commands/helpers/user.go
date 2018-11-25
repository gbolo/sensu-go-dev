package helpers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sensu/sensu-go/cli/client/config"
	types "github.com/sensu/sensu-go/api/core/v2"
)

// GetCurrentUsername retrieves the username from the active JWT
func GetCurrentUsername(cfg config.Config) string {
	accessToken := cfg.Tokens().Access
	token, _ := jwt.ParseWithClaims(accessToken, &types.Claims{}, nil)
	claims := token.Claims.(*types.Claims)
	return claims.StandardClaims.Subject
}
