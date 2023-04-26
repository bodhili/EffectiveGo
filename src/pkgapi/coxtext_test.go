package pkgapi

import (
	"context"
	"testing"
	"time"
)

const shortDuration = 1 * time.Millisecond

func TestContextDeadline(t *testing.T) {
	d := time.Now().Add(shortDuration)

	ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	defer cancelFunc()

	select {
	case <-time.After(1 * time.Second):
		t.Log("overslept")
	case <-ctx.Done():
		t.Log(ctx.Err())
	}
}
