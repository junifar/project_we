package constant

// context
const (
	ContextUserID = "UserID"
)

// cookie
const (
	CookieName     = "x-Session"
	CookieSecret   = "mySecret" // TODO move into config
	CookieMaxAge   = 28800
	CookieHttpOnly = true
	CookieSameSite = "None"
)
