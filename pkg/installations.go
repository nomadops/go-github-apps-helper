package installations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
)

// RemoveRepoFromInstallation removes a GitHub repository from a GitHub App installation.
func RemoveRepoFromInstallation(ctx context.Context, appID int64, installationID int64, repoID int64, itr *ghinstallation.AppsTransport) error {

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
