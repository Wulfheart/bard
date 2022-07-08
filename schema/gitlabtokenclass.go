package schema

type GitlabTokenClass struct {
	Token    string `json:"token"`   // The token used for GitLab authentication
	Username string `json:"username"`// The username used for GitLab authentication
}
