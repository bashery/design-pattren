package main

type iPoolObject interface {
   // this is any id which can be used to commpar to deferent objects
    getID() string
}
