package commands

const (
	RunHelp        = "run server on port"
	DBBuildHelp    = "create all tables"
	DBExportHelp   = "print all create table sql queries"
	DBLoadHelp     = "load all sql file from path (--path=database/fixtures)"
	DBFixturesHelp = "load all fixtures from directory (--path=relative_path_to_fixtures)"

	UserUpdateRolesHelp    = "set new user roles list (--by=username/email --roles=firstrole&secondrole...)"
	UserAddHelp            = "add new user (--userproperty=value...)"
	UserUpdatePasswordHelp = "set new password for user (--by=username/email --password=somePassword)"
	UserConnectAddHelp     = "add new user connection to remote service (--service=serviceName --remote=remoteID --local=username/email)"

	EnvArg    = "set app environment (env=prod|test|dev, prod by default)"
	LoglvlArg = "set app log level (loglvl=prod|test|dev, prod by default)"
	HostArg   = "set server host (host=:5555, :5555 by default)"

	// DefaultFixtureDir is default path to fixtures directory
	DefaultFixtureDir = "database/fixtures"
)
