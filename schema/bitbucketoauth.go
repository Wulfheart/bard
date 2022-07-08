package schema

type BitbucketOauth struct {
	AccessToken           *string `json:"access-token,omitempty"`           // The OAuth token retrieved from Bitbucket's API, this is written by Composer and you; should not set it nor modify it.
	AccessTokenExpiration *int64  `json:"access-token-expiration,omitempty"`// The generated token's expiration timestamp, this is written by Composer and you should; not set it nor modify it.
	ConsumerKey           string  `json:"consumer-key"`                     // The consumer-key used for OAuth authentication
	ConsumerSecret        string  `json:"consumer-secret"`                  // The consumer-secret used for OAuth authentication
}
