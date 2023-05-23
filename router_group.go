package httpr

// not thread-safe
type Group struct {
	prefix string
	mws    []Middleware
	r      *Router
}

func (g *Group) Group(path string, mws ...Middleware) *Group {
	return &Group{
		r:      g.r,
		mws:    append(g.mws, mws...),
		prefix: joinPathPanic(g.prefix, path),
	}
}

// GET is a shortcut for router.Handle(http.MethodGet, path, handle) with handle wrapped by middlewares
func (g *Group) GET(path string, handle Handle, mws ...Middleware) {
	g.r.GET(joinPathPanic(g.prefix, path), handle, append(g.mws, mws...)...)
}

// HEAD is a shortcut for router.Handle(http.MethodHead, path, handle) with handle wrapped by middlewares
func (g *Group) HEAD(path string, handle Handle, mws ...Middleware) {
	g.r.HEAD(joinPathPanic(g.prefix, path), handle, append(g.mws, mws...)...)
}

// OPTIONS is a shortcut for router.Handle(http.MethodOptions, path, handle) with handle wrapped by middlewares
func (g *Group) OPTIONS(path string, handle Handle, mws ...Middleware) {
	g.r.OPTIONS(joinPathPanic(g.prefix, path), handle, append(g.mws, mws...)...)
}

// POST is a shortcut for router.Handle(http.MethodPost, path, handle) with handle wrapped by middlewares
func (g *Group) POST(path string, handle Handle, mws ...Middleware) {
	g.r.POST(joinPathPanic(g.prefix, path), handle, append(g.mws, mws...)...)
}

// PUT is a shortcut for router.Handle(http.MethodPut, path, handle) with handle wrapped by middlewares
func (g *Group) PUT(path string, handle Handle, mws ...Middleware) {
	g.r.PUT(joinPathPanic(g.prefix, path), handle, append(g.mws, mws...)...)
}

// PATCH is a shortcut for router.Handle(http.MethodPatch, path, handle) with handle wrapped by middlewares
func (g *Group) PATCH(path string, handle Handle, mws ...Middleware) {
	g.r.PATCH(joinPathPanic(g.prefix, path), handle, append(g.mws, mws...)...)
}

// DELETE is a shortcut for router.Handle(http.MethodDelete, path, handle) with handle wrapped by middlewares
func (g *Group) DELETE(path string, handle Handle, mws ...Middleware) {
	g.r.DELETE(joinPathPanic(g.prefix, path), handle, append(g.mws, mws...)...)
}
