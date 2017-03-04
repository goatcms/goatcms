package services

const (
	CreateRequestScope = 1000

	// RequestTagName is a tag name used for a request injection
	RequestTagName = "request"
	// RequestTagName is a tag name used for a request injection
	FormTagName = "form"

	// SessionCookieID is default name of session cookie
	SessionCookieID = "session"
	// SessionCookieLength is default length of session id (storaged by cookie)
	SessionIDLength = 128
	// SessionCookieLifetime is a lifetime of cookie
	SessionLifetime = 365 * 24
	// SessionExpire is key to read expire time from session
	SessionExpireKey = "session.expire"

	// DefaultTemplatePath is a default path for temapates
	DefaultTemplatePath = "./cmsapp/templates"

	// DefaultDatabaseEngine is default engine for database
	DefaultDatabaseEngine = "sqlite3"
	// DefaultDatabaseUrl is default url/path for database
	DefaultDatabaseUrl = "./database/sqlite3_database.db"
)
