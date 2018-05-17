package common

// TODO: might move to service?
// TODO: https://github.com/dyweb/go.ice/issues/12 error tracker
// common error structs

type ErrUserExists struct {
	name string
}
