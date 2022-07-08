package schema

type Support struct {
	Chat   *string `json:"chat,omitempty"`  // URL to the support chat.
	Docs   *string `json:"docs,omitempty"`  // URL to the documentation.
	Email  *string `json:"email,omitempty"` // Email address for support.
	Forum  *string `json:"forum,omitempty"` // URL to the forum.
	IRC    *string `json:"irc,omitempty"`   // IRC channel for support, as irc://server/channel.
	Issues *string `json:"issues,omitempty"`// URL to the issue tracker.
	RSS    *string `json:"rss,omitempty"`   // URL to the RSS feed.
	Source *string `json:"source,omitempty"`// URL to browse or download the sources.
	Wiki   *string `json:"wiki,omitempty"`  // URL to the wiki.
}
