package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m        sync.Mutex
	resource chan io.Closer
	factory  func() (io.Closer, error)
	closed   bool
}

var ErrPoolClosed = errors.New("pool has been closed")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}
	return &Pool{
		factory:  fn,
		resource: make(chan io.Closer, size),
	}, nil
}

//Acquire 获取资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resource:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release 释放资源
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		err := r.Close()
		if err != nil {
			return
		}
		return
	}

	select {
	case p.resource <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release:", "Closing")
		err := r.Close()
		if err != nil {
			return
		}
	}
}

// Close 关闭资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.resource)

	for r := range p.resource {
		err := r.Close()
		if err != nil {
			return
		}
	}
}
