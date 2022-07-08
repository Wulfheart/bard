package schema

// A protocol to force use of when creating a repository URL for the `source` value of the
// package metadata. One of `git` or `http`. By default, Composer will generate a git URL
// for private repositories and http one for public repos.
type GitlabProtocol string
const (
	GitlabProtocolGit GitlabProtocol = "git"
	HTTP GitlabProtocol = "http"
	HTTPS GitlabProtocol = "https"
)
