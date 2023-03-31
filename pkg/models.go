package installations

import "time"

type InstallationReposEventPayload struct {
	Action       string `json:"action"`
	Installation struct {
		ID      int `json:"id"`
		Account struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"account"`
		RepositorySelection string `json:"repository_selection"`
		AccessTokensURL     string `json:"access_tokens_url"`
		RepositoriesURL     string `json:"repositories_url"`
		HTMLURL             string `json:"html_url"`
		AppID               int    `json:"app_id"`
		AppSlug             string `json:"app_slug"`
		TargetID            int    `json:"target_id"`
		TargetType          string `json:"target_type"`
		Permissions         struct {
			Contents string `json:"contents"`
			Metadata string `json:"metadata"`
		} `json:"permissions"`
		Events                 []string      `json:"events"`
		CreatedAt              time.Time     `json:"created_at"`
		UpdatedAt              time.Time     `json:"updated_at"`
		SingleFileName         interface{}   `json:"single_file_name"`
		HasMultipleSingleFiles bool          `json:"has_multiple_single_files"`
		SingleFilePaths        []interface{} `json:"single_file_paths"`
		SuspendedBy            interface{}   `json:"suspended_by"`
		SuspendedAt            interface{}   `json:"suspended_at"`
	} `json:"installation"`
	Repositories []struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
	} `json:"repositories"`
	Requester interface{} `json:"requester"`
	Sender    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
}
