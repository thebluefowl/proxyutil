package proxyutil

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// CloneRequest clones the request.  This should not be used for
// very large bodies.  The body is first read and the reader is
// replaced in both the source and destination requests.
func CloneRequest(src *http.Request) (*http.Request, error) {
	ctx := src.Context()
	dst := src.Clone(ctx)

	b, err := io.ReadAll(src.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}
	src.Body.Close()

	src.Body = io.NopCloser(bytes.NewReader(b))
	dst.Body = io.NopCloser(bytes.NewReader(b))

	dst.ContentLength = int64(len(b))

	return dst, nil
}

func GetProto(r *http.Request) string {
	if r.TLS == nil {
		return "http"
	}
	return "https"
}
