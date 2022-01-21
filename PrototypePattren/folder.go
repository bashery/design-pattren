package main

import "fmt"

type folder struct {
    childrens []iNode
    name string
}

func (f *folder) print(indentation string) {
    fmt.Println(indentation+f.name)
    for _, i := range f.childrens {
        i.print(indentation+indentation)
    }
}

func (f *folder) clone() iNode {

    cloneFolder := &folder{name: f.name+ "_clone"}
    var tempChildrens []iNode
    for _, i := range f.childrens {
        copy := i.clone()
        tempChildrens = append(tempChildrens, copy)
    }
    cloneFolder.childrens = tempChildrens
    return cloneFolder
}

