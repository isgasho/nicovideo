package nicovideo

import (
	"strings"
)

// IsPostedByUser returns true if video ID starts with "sm", namely posted by
// user.
func IsPostedByUser(ID string) bool {
	return strings.HasPrefix(ID, "sm")
}
