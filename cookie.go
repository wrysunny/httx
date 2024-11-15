package httx

import (
	"net/http/cookiejar"
)

var jar, _ = cookiejar.New(nil)
