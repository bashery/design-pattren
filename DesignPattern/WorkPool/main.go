// The Object Pool Design Pattern is a creational design pattern
// in which a pool of objects is initialized and created beforehand
// and kept in a pool. As and when needed, a client can request an
// object from the pool, use it, and return it to the pool.
// The object in the pool is never destroyed.
package main

import (
	"log"
	"strconv"
)

func main() {
    connections := make([]iPoolObject, 0)
    
    for i := 0; i< 3; i++ {
        c := &connection{id:strconv.Itoa(i)}
        connections = append(connections, c)
    }
    pool, err := initPool(connections)
    if err != nil {
        log.Fatalf("Init Pool Error: %s", err)    
    }
    conn1, err := pool.loan()
    if err != nil {
        log.Fatalf("Pool Lean Error: %s", err)    
    }
    conn2, err := pool.loan()
    if err != nil {
        log.Fatalf("Pool Lean Error: %s", err)    
    }
    pool.receive(conn1)
    pool.receive(conn2)

}

