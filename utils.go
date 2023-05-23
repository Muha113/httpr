package httpr

import "net/url"

func joinPathPanic(path ...string) string {
	p, err := url.JoinPath("/", path...)
	if err != nil {
		panic(err)
	}
	return p
}

func wrapMiddleware(handle Handle, mws ...Middleware) Handle {
	h := handle
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
