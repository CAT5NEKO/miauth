package miauth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Session struct {
	ID string
}

type UserInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccessTokenResponse struct {
	Token string   `json:"token"`
	User  UserInfo `json:"user"`
}

func GenerateSessionID() string {
	return uuid.New().String()
}

func ConstructMiauthURL(sessionID, appName, callbackURL, permission string) string {
	return fmt.Sprintf("https://misskey.io/miauth/%s?name=%s&callback=%s&permission=%s", sessionID, appName, callbackURL, permission)
}

func PerformMiauthAuthentication(sessionID, host string) (*AccessTokenResponse, error) {
	checkURL := fmt.Sprintf("https://%s/api/miauth/%s/check", host, sessionID)

	resp, err := http.Post(checkURL, "application/json", nil)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	var accessTokenResponse AccessTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&accessTokenResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	return &accessTokenResponse, nil
}
