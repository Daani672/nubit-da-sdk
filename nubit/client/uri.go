package client

import "sync"

type Uri struct {
	mu  *sync.Mutex
	ctx Context
}

func NewUri(ctx Context) *Uri {
	return &Uri{mu: &sync.Mutex{}, ctx: ctx}
}

func (u *Uri) GetUri(name string) string {
	u.mu.Lock()
	defer u.mu.Unlock()
	if v, ok := u.ctx.UriMap[name]; ok {
		return v
	}
	return ""
}
