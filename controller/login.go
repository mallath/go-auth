package controller

import (
	"net/http"

	"github.com/youtube/google-oauth/config"
	"github.com/youtube/google-oauth/utils"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	oauthState := utils.GenerateStateOauthCookie(w)

	u := config.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}
