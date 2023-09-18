package proxyutil

import (
	"net/url"
	"strings"
)

func JoinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	apath := a.EscapedPath()
	bpath := b.EscapedPath()
	return adjustPaths(a.Path, b.Path, apath, bpath)
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func adjustPaths(aPath, bPath, aEscaped, bEscaped string) (path, rawpath string) {
	aslash := strings.HasSuffix(aEscaped, "/")
	bslash := strings.HasPrefix(bEscaped, "/")

	switch {
	case aslash && bslash:
		return aPath + bPath[1:], aEscaped + bEscaped[1:]
	case !aslash && !bslash:
		return aPath + "/" + bPath, aEscaped + "/" + bEscaped
	}
	return aPath + bPath, aEscaped + bEscaped
}
