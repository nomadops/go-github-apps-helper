package installations

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/golang-jwt/jwt"
)

// RemoveRepoFromInstallation removes a GitHub repository from a GitHub App installation.
func RemoveRepoFromInstallation(ctx context.Context, appID int64, installationID int64, repoID int64, itr *ghinstallation.Transport) error {

	// Create a new HTTP client using the installation transport
	client := &http.Client{
		Transport: itr,
	}

	// Create the API request to remove the repository from the installation
	requestUrl := fmt.Sprintf("https://api.github.com/user/installations/%d/repositories/%d", installationID, repoID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create API request: %w", err)
	}

	// Send the API request using the HTTP client
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to remove repository from installation: %w", err)
	}
	defer resp.Body.Close()

	// Check the API response status code
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected API response status code: %d", resp.StatusCode)
	}

	return nil
}

// Token returns the complete, signed Github app JWT token
func Token(itr *ghinstallation.AppsTransport, appID int64, key string) (string, error) {
	claims := &jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute).Unix(),
		Issuer:    strconv.FormatInt(appID, 10),
	}
	bearer := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return bearer.SignedString(key)
}

// AppRemoveRepoFromInstallation removes a GitHub repository from a GitHub App installation.
func AppRemoveRepoFromInstallation(ctx context.Context, appID int64, installationID int64, repoID int64, itr *ghinstallation.AppsTransport, key string) error {

	token, err := Token(itr, appID, key)
	if err != nil {
		return fmt.Errorf("failed to create token: %w", err)
	}

	// Create a new HTTP client using the installation transport
	client := &http.Client{
		Transport: itr,
	}

	// Create the API request to remove the repository from the installation
	requestUrl := fmt.Sprintf("https://api.github.com/user/installations/%d/repositories/%d", installationID, repoID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create API request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Send the API request using the HTTP client
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to remove repository from installation: %w", err)
	}
	defer resp.Body.Close()

	// Check the API response status code
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected API response status code: %d", resp.StatusCode)
	}

	return nil
}
