// Package cache defines interfaces for using Key-Value cache and Redis
package cache

type KV interface {
	Get(key string) (val interface{}, err error)
	Set(key string, val interface{}) error
	//KeyCount()
	//Size()
	ClearCache() error
	SyncCache() error
}
