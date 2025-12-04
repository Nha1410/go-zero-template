package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// ZitadelConfig holds Zitadel OAuth2 configuration
type ZitadelConfig struct {
	Issuer       string
	ClientID     string
	ClientSecret string
	Scopes       []string
}

// ZitadelClient wraps Zitadel OAuth2 client
type ZitadelClient struct {
	config       *oauth2.Config
	clientConfig *clientcredentials.Config
	issuer       string
}

// UserInfo represents user information from Zitadel
type UserInfo struct {
	Sub           string   `json:"sub"`
	Email         string   `json:"email"`
	EmailVerified bool     `json:"email_verified"`
	Name          string   `json:"name"`
	GivenName     string   `json:"given_name"`
	FamilyName    string   `json:"family_name"`
	Picture       string   `json:"picture"`
	Roles         []string `json:"roles"`
}

// NewZitadelClient creates a new Zitadel OAuth2 client
func NewZitadelClient(config ZitadelConfig) (*ZitadelClient, error) {
	oauth2Config := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("%s/oauth/v2/authorize", config.Issuer),
			TokenURL: fmt.Sprintf("%s/oauth/v2/token", config.Issuer),
		},
	}

	clientConfig := &clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       config.Scopes,
		TokenURL:     fmt.Sprintf("%s/oauth/v2/token", config.Issuer),
	}

	return &ZitadelClient{
		config:       oauth2Config,
		clientConfig: clientConfig,
		issuer:       config.Issuer,
	}, nil
}

// ValidateToken validates an OAuth2 token
func (z *ZitadelClient) ValidateToken(ctx context.Context, token string) (*UserInfo, error) {
	// Extract token from "Bearer <token>" format
	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimSpace(token)

	// Create HTTP client with token
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
		TokenType:   "Bearer",
	}))

	// Call userinfo endpoint
	userInfoURL := fmt.Sprintf("%s/oidc/v1/userinfo", z.issuer)
	req, err := http.NewRequestWithContext(ctx, "GET", userInfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call userinfo endpoint: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid token: status code %d", resp.StatusCode)
	}

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode userinfo: %w", err)
	}

	return &userInfo, nil
}

// GetClientCredentialsToken gets a token using client credentials flow
func (z *ZitadelClient) GetClientCredentialsToken(ctx context.Context) (*oauth2.Token, error) {
	return z.clientConfig.Token(ctx)
}

// ExtractTokenFromRequest extracts token from HTTP request
func ExtractTokenFromRequest(r *http.Request) string {
	// Try to get token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
			return parts[1]
		}
	}

	// Try to get token from query parameter
	token := r.URL.Query().Get("token")
	if token != "" {
		return token
	}

	return ""
}

