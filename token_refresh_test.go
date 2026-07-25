package chainstream_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	chainstream "github.com/chainstream-io/chainstream-go-sdk/v2"
)

// counterTokenProvider returns a different token on every GetToken() call,
// simulating a real provider that rotates/refreshes short-lived JWTs.
type counterTokenProvider struct {
	mu sync.Mutex
	n  int
}

func (p *counterTokenProvider) GetToken() (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.n++
	return fmt.Sprintf("token-%d", p.n), nil
}

// captureServer records the Authorization header of every incoming request.
func captureServer(t *testing.T) (*httptest.Server, func() []string) {
	t.Helper()
	var mu sync.Mutex
	var auths []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		auths = append(auths, r.Header.Get("Authorization"))
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	}))
	return srv, func() []string {
		mu.Lock()
		defer mu.Unlock()
		out := make([]string, len(auths))
		copy(out, auths)
		return out
	}
}

// TestTokenProviderRefreshesPerRequest is the regression test for the H1 bug:
// with a TokenProvider, EVERY REST request must fetch a fresh token instead of
// reusing the one captured at client-creation time.
func TestTokenProviderRefreshesPerRequest(t *testing.T) {
	srv, getAuths := captureServer(t)
	defer srv.Close()

	provider := &counterTokenProvider{}
	client, err := chainstream.NewChainStreamClientWithTokenProvider(provider, &chainstream.ChainStreamClientOptions{
		ServerURL: srv.URL,
	})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	for i := 0; i < 2; i++ {
		if _, err := client.Blockchain.GetSupportedBlockchainsWithResponse(ctx); err != nil {
			t.Fatalf("request %d failed: %v", i, err)
		}
	}

	auths := getAuths()
	if len(auths) != 2 {
		t.Fatalf("expected 2 captured requests, got %d: %v", len(auths), auths)
	}
	for i, a := range auths {
		if !strings.HasPrefix(a, "Bearer token-") {
			t.Fatalf("request %d: expected Bearer token, got %q", i, a)
		}
	}
	if auths[0] == auths[1] {
		t.Fatalf("token was NOT refreshed between requests: both %q (H1 regression)", auths[0])
	}
}

// TestStaticAccessTokenStaysConstant documents that a static-string accessToken
// is (intentionally) NOT refreshed: the same value is sent on every request.
func TestStaticAccessTokenStaysConstant(t *testing.T) {
	srv, getAuths := captureServer(t)
	defer srv.Close()

	client, err := chainstream.NewChainStreamClient("static-token", &chainstream.ChainStreamClientOptions{
		ServerURL: srv.URL,
	})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	for i := 0; i < 2; i++ {
		if _, err := client.Blockchain.GetSupportedBlockchainsWithResponse(ctx); err != nil {
			t.Fatalf("request %d failed: %v", i, err)
		}
	}

	auths := getAuths()
	if len(auths) != 2 {
		t.Fatalf("expected 2 captured requests, got %d: %v", len(auths), auths)
	}
	want := "Bearer static-token"
	if auths[0] != want || auths[1] != want {
		t.Fatalf("expected both requests to send %q, got %v", want, auths)
	}
}
