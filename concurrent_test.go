package templ_test

import (
	"context"
	"io"
	"sync"
	"testing"

	"github.com/a-h/templ"
)

func TestWithChildrenConcurrentSafety(t *testing.T) {
	ctx := templ.InitializeContext(context.Background())
	child := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error { return nil })

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newCtx := templ.WithChildren(ctx, child)
			_ = templ.GetChildren(newCtx)
		}()
	}
	wg.Wait()
}
