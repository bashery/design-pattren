package main

import (
	"fmt"
	"sync"
)

type pool struct {
    idle []iPoolObject
    active []iPoolObject
    capacity int
    mulock *sync.Mutex
}

// InitPool Initalize the pool
func initPool(poolObjects []iPoolObject) (*pool, error) {
    if len(poolObjects) == 0 {
        return nil, fmt.Errorf("Cannot create a pool of 0 length")
    }
    active := make([]iPoolObject, 0)
    pool := &pool {
        idle: poolObjects,
        active: active,
        capacity: len(poolObjects),
        mulock: new(sync.Mutex),
    }
    return pool, nil
} 

func (p *pool) loan() (iPoolObject, error) {
    p.mulock.Lock()
    defer p.mulock.Unlock()
    if len(p.idle) == 0 {
        return nil, fmt.Errorf("No pool object free. Please request later")
    }
    obj := p.idle[0]
    p.idle = p.idle[1:]
    p.active = append(p.active, obj)
    fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())
    return obj, nil
}
