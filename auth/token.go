package auth

import (
	"context"
	"github.com/carlmjohnson/requests"
	"github.com/ddshuai/go-qywx/api"
)

type AccessTokenResponse struct {
	api.BaseResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// CreateToken 创建token/*
func CreateToken(url, corpId, secret string) (string, error) {
	var res AccessTokenResponse
	err := requests.URL(url).Path(api.Prefix+api.GetToken).Param("corpid", corpId).Param("corpsecret", secret).ToJSON(&res).Fetch(context.Background())
	if err != nil {
		return "", err
	}

	return res.AccessToken, nil
}
