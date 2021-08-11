// It is a creational design pattern that lets you create copies of objects. In this pattern, the responsibility of creating the clone objects is delegated to the actual object to clone.
package main

type iNode interface {
    print(string)
    clone() iNode
}
