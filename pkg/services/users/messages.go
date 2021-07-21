package users

import "github.com/kislball/dva/pkg/services"

type TokenType string

const (
	TokenTypeBearer TokenType = "Bearer"
)

type OAuthExchangeRequest struct {
	RedirectURI string `json:"redirect_uri"`
	Code        string `json:"code"`
}

type OAuthExchangeResponse struct {
	TokenType TokenType `json:"token_type"`
	Token     string    `json:"token"`
	Refresh   string    `json:"refresh"`
	Expiry    int64     `json:"expiry"`
	services.AtomicResponse
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	User
	services.AtomicResponse
}
