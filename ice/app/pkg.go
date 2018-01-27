// Package app define common interface for both client and server application
package app

type App interface {
	IsClient() bool
	IsServer() bool
	Verbose() bool
	Version() string
}
