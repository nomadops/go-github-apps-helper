package installations

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bradleyfalzon/ghinstallation/v2"
	jwt "github.com/golang-jwt/jwt"
	"github.com/google/go-github/github"
)

// Token returns the complete, signed Github app JWT token
func Token(itr *ghinstallation.Transport, appID int64, key []byte) (string, error) {
	claims := &jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute).Unix(),
		Issuer:    strconv.FormatInt(appID, 10),
	}
	bearer := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return bearer.SignedString(key)
}

// Token returns the complete, signed Github app JWT token.
func AppToken(itr *ghinstallation.AppsTransport, appID int64, key []byte) ([]byte, error) {

	log.Printf("AppToken itr: %#v", itr)
	log.Printf("key: %s", key)

	claims := &jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute).Unix(),
		Issuer:    strconv.FormatInt(appID, 10),
	}

	log.Printf("claims: %#v", claims)
	bearer := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	log.Printf("bearer: %#v", bearer)

	bearerString, err := bearer.SignedString(key)
	return []byte(bearerString), err
}

// AppRemoveRepoFromInstallation removes a GitHub repository from a GitHub App installation.
// func AppRemoveRepoFromInstallation(ctx context.Context, appID int64, installationID int64, repoID int64, itr *ghinstallation.AppsTransport, token []byte) error {
// func AppRemoveRepoFromInstallation(appID int64, installationID int64, repoID int64, itr *ghinstallation.AppsTransport, token []byte) error {
// func AppRemoveRepoFromInstallation(appID int64, installationID int64, repoID int64, itr *ghinstallation.AppsTransport, key []byte) error {
func AppRemoveRepoFromInstallation(appID int64, installationID int64, repoID int64, itr *ghinstallation.AppsTransport) error {

	client := github.NewClient(&http.Client{Transport: itr})

	req, err := client.NewRequest("DELETE", fmt.Sprintf("installations/%d/repositories/%d", installationID, repoID), nil)
	if err != nil {
		return fmt.Errorf("failed to create the request: %w", err)
	}

	resp, err := client.Do(context.Background(), req, nil)
	if err != nil {
		return fmt.Errorf("failed to remove the repository from the installation: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Printf("Repository with ID %d has been removed from installation %d\n", repoID, installationID)
	return nil
}
