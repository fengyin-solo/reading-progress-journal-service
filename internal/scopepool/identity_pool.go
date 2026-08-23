package scopepool

import (
	"sync"

	"readinglog/internal/model"
)

type IdentityPool struct {
	pool sync.Pool
}

func New() *IdentityPool {
	p := &IdentityPool{}
	p.pool.New = func() interface{} { return &model.RequestIdentity{} }
	return p
}

func (p *IdentityPool) Acquire(reader string, books []string) *model.RequestIdentity {
	identity := p.pool.Get().(*model.RequestIdentity)
	identity.ReaderID = reader
	identity.BookIDs = books
	return identity
}

func (p *IdentityPool) Release(identity *model.RequestIdentity) {
	p.pool.Put(identity)
}
