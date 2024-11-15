package httx

import (
	"context"
	"time"
)

var ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
