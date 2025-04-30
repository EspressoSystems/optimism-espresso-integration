package environment

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	tagged_base64 "github.com/EspressoSystems/espresso-network-go/tagged-base64"
	types "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
)

type InterceptBehavior uint

const (
	BehaviorProxy                        InterceptBehavior = 0
	BehaviorSubmitTxnSuccessWhileDropped InterceptBehavior = 1 << iota
)

// HasBehavior is a method that checks whether the passed InterceptBehavior is
// contained within the bitmap of the host InterceptBehavior.
func (i InterceptBehavior) HasBehavior(b InterceptBehavior) bool {
	return i&b == b
}

// requestMatchesPath checks if the HTTP request matches the specified method
func requestMatchesPath(r *http.Request, method string, pathMatcher func(string) bool) bool {
	return r.Method == method && r.URL != nil && pathMatcher(r.URL.Path)
}

// stringEquals is a helper function that returns a function that checks if
// a given path string equals the specified string.
func stringEquals(s string) func(string) bool {
	return func(path string) bool {
		return path == s
	}
}

// isSubmitTransactionRequest represents the different variations of the submit
// transaction endpoint that we can utilize or support.
func isSubmitTransactionRequest(r *http.Request) bool {
	return requestMatchesPath(r, http.MethodPost, stringEquals("/submit/submit")) ||
		requestMatchesPath(r, http.MethodPost, stringEquals("/v0/submit/submit"))
}

// EspressoDevNodeIntercept is a struct that is a Proxy to the Espresso Dev Node.
// It is used to intercept request to the Espresso Dev Node and make decisions
// about handling the requests.  This is useful for simulating failures or
// bad behaviors for Espresso.
type EspressoDevNodeIntercept struct {
	u      url.URL
	b      InterceptBehavior
	client *http.Client
}

// performProxy performs the actual proxying of the request to the stored URL.
func (e *EspressoDevNodeIntercept) performProxy(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r.Body); err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(r.Method, e.u.JoinPath(r.URL.Path).String(), buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := e.client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	w.WriteHeader(res.StatusCode)
	for k, v := range res.Header {
		w.Header().Set(k, v[0])
	}
	w.Header().Set("Origin", e.u.Host)
	w.Header().Set("Host", e.u.Host)

	// Write the proxy response contents
	if _, err := io.Copy(w, res.Body); err != nil {
		// If we encounter an error here, it will be difficult to actually
		// handle it at this point, as we've already sent the response headers.
		//
		// The best we can do at this point, is log the error.
		_ = err
		return
	}
}

func (e *EspressoDevNodeIntercept) generateCommitForSubmitTransaction(r *http.Request) (*types.TaggedBase64, error) {
	defer r.Body.Close()

	var txn types.Transaction
	if err := json.NewDecoder(r.Body).Decode(&txn); err != nil {
		// Unable to decode, this is a problem?
		var emptyHash [32]byte
		return tagged_base64.New("FAKE", emptyHash[:])
	}

	commit := txn.Commit()
	return tagged_base64.New("FAKE", commit[:])
}

func (e *EspressoDevNodeIntercept) simulateSuccessfulSubmitTransaction(w http.ResponseWriter, r *http.Request) {
	// We could do a lot of effort to validate the request, and return an
	// hash that is actually representative of the transaction that was
	// just submitted. In some cases we may actually want this sort of
	// validated behavior, but it's very simple to just return any hash
	// instead.

	// We should probably validate the request contents and format here, but
	// we will just assume the settings.
	defer r.Body.Close()
	hash, err := e.generateCommitForSubmitTransaction(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contents, err := json.Marshal(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(contents)
}

func (e *EspressoDevNodeIntercept) handleBehavior(w http.ResponseWriter, r *http.Request) {
	if e.b.HasBehavior(BehaviorSubmitTxnSuccessWhileDropped) && isSubmitTransactionRequest(r) {
		e.simulateSuccessfulSubmitTransaction(w, r)
		return
	}

	// If we don't have any other behavior to perform that we've detected, then
	// we'll just default to proxying the request.
	e.performProxy(w, r)
}

// ServerHTTP implements http.Handler
func (e *EspressoDevNodeIntercept) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Close the request body to prevent resource leaks
	defer r.Body.Close()

	if e.u == (url.URL{}) {
		// we don't have a URL to redirect to, so we can do nothing
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// Perform the proxy request.
	e.handleBehavior(w, r)
}

// createEspressoProxyOption will return a Batch CLIConfig option that will
// replace the Espresso URL with the URL of the proxy server.
func createEspressoProxyOption(ctx *DevNetLauncherContext, proxy *EspressoDevNodeIntercept, server *httptest.Server) func(*batcher.CLIConfig) {
	return func(cfg *batcher.CLIConfig) {
		if ctx.Error != nil {
			return
		}

		if cfg.EspressoUrl == "" {
			// This should be being called after the Espresso
			// Dev Node is Already Live.
			// Without an Espresso URL, we cannot proceed.
			return
		}

		u, err := url.Parse(cfg.EspressoUrl)
		if err != nil || u == nil {
			// We encountered an error
			ctx.Error = err
			return
		}

		// Set the proxy
		proxy.u = *u
		// Replace the Espresso URL with the proxy URL
		cfg.EspressoUrl = server.URL
	}
}

// SetupQueryServiceIntercept sets up an intercept traffic headed for the
// Query Service for the Espresso Dev Node
func SetupQueryServiceIntercept() (*EspressoDevNodeIntercept, *httptest.Server, DevNetLauncherOption) {
	// Start a Server to proxy requests to Espresso
	proxy := &EspressoDevNodeIntercept{
		client: http.DefaultClient,
	}

	// Start up a local http server to handle the requests
	server := httptest.NewServer(proxy)

	return proxy, server, func(ctx *DevNetLauncherContext) E2eSystemOption {
		return E2eSystemOption{
			StartOptions: []e2esys.StartOption{
				{
					Key:        "espresso-proxy",
					BatcherMod: createEspressoProxyOption(ctx, proxy, server),
				},
			},
		}
	}
}
