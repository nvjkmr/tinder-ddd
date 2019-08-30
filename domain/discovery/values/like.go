package values // TODO: figure out if we can have same package name across directories

import (
	"time"
)

type Like struct {
	liker, likee int64
	// like_type    likeType
	expires_on time.Time
}
