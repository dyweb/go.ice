package httpclient

import (
	"strings"
)

// JoinPath does not sanitize path like path.Join, which would change https:// to https:/, it only remove duplicated
// slashes to avoid // in url i.e. http://myapi.com/api/v1//comments/1
func JoinPath(base string, sub string) string {
	return strings.TrimRight(base, "/") + "/" + strings.TrimLeft(sub, "/")
}
