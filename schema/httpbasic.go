package schema

type HTTPBasic struct {
	Password string `json:"password"`// The password used for HTTP Basic authentication
	Username string `json:"username"`// The username used for HTTP Basic authentication
}
