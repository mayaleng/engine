package version

import (
	"time"
)

// BuildNumber represents the version of the app
var BuildNumber string = "1.0.0"

// BuildTime represents when this version was built
var BuildTime string = time.Now().String()
