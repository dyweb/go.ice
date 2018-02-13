package common

// TODO: this should be handled by go.ice ...
var (
	version   string
	gitCommit string
	buildTime string
	buildUser string
)

func Version() string {
	return version
}

func GitCommit() string {
	return gitCommit
}

func BuildTime() string {
	return buildTime
}

func BuildUser() string {
	return buildUser
}
