package main

import "testing"
import "context"
import "io"
import "fmt"
import "net/http"
import "golang.org/x/sync/errgroup"

func TestMainFunc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	in := "message"
	rsp, err := http.Get("http://localhost:18080/" + in)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// HTTPサーバーの返り値を検証する
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// run関数に終了を通知する
	cancel()
	// run関数の返り値を検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
