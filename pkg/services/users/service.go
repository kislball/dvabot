package users

import (
	"encoding/json"
	"errors"
	"github.com/kislball/dva/pkg/communication"
	"time"
)

type User struct {
	ID string `json:"id"`
}

type Service struct {
	nats *communication.Manager
}

func (u *Service) OAuthExchange(code string, redirectUri string) (r OAuthExchangeResponse, err error) {
	d, err := json.Marshal(OAuthExchangeRequest{
		Code:        code,
		RedirectURI: redirectUri,
	})
	if err != nil {
		return
	}

	resp, err := u.nats.Request("users.oauth_exchange", d, time.Second*5)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &r)

	if !r.Ok {
		err = errors.New(r.FailReason)
	}

	return
}

func (u *Service) GetUser(id string) (usr GetUserResponse, err error) {
	d, err := json.Marshal(GetUserRequest{
		ID: id,
	})
	if err != nil {
		return
	}

	resp, err := u.nats.Request("users.get_user", d, time.Second*5)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &usr)

	if !usr.Ok {
		err = errors.New(usr.FailReason)
	}

	return
}
