package httpclient

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyrillemad/csmt/types"
)

func TestClientGet(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if "/check" != r.URL.Path {
				t.Errorf("Bad url path, got %q, want %q", r.URL.Path, "/check")
			}
			if r.Method != http.MethodGet {
				t.Errorf("Bad method, got %q, want %q", r.Method, http.MethodGet)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`"Hello World"`))
		}))
	defer server.Close()

	client := NewClient()
	var response string

	err := client.Get(
		context.Background(),
		server.URL+"/check",
		types.Authorize{},
		&response)

	if err != nil {
		t.Fatal(err)
	}
	if response != "Hello World" {
		t.Errorf("Bad response, got %q, want %q", response, "Hello World")
	}
}

func TestClientTimeout(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`"Hello World"`))
		}))
	defer server.Close()

	client := NewClient(
		WithTimeout(25*time.Millisecond),
		WithRetryCount(0))
	var response string

	err := client.Get(
		context.Background(),
		server.URL,
		types.Authorize{},
		&response)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Error("Expected timeout error")
	}
}

func TestClientAuthorize(t *testing.T) {
	token := "TOP-SECRET-KEY"
	auth := types.Authorize{
		Key:    token,
		Header: "password",
	}

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("password") != token {
				t.Errorf(
					"Bad response, got %q, want %q",
					r.Header.Get("password"), token)
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`"Unauthorized"`))
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`"Secret data"`))
			}
		}))
	defer server.Close()

	client := NewClient()
	var response string

	_ = client.Get(
		context.Background(),
		server.URL,
		auth,
		&response)
}
