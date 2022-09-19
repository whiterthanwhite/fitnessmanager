package clientapi

import (
	"context"
	"fmt"
	"time"
)

func EmptyFunction() {
	fmt.Println(time.Now())
}

func SyncEntities(ctx context.Context) error {
	return nil
}
