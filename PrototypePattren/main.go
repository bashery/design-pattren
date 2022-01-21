package main

import "fmt"

func main() {
    file1 := &file{name:"File1"}
    file2 := &file{name:"File2"}
    file3 := &file{name:"File3"}
    file4 := &file{name:"File4"}
    file5 := &file{name:"File5"}

    folder1 := &folder{
        childrens: []iNode{file1, file5},
        name: "Folder1",
    }
    
    folder2 := &folder{
        childrens: []iNode{folder1, file2,file3, file4},
        name: "Folder1",
    }

    fmt.Println("\nPrinting hierarchy for Folder2")
    folder2.print("    ")

    cloneFolder := folder2.clone()

    fmt.Println("\nPrinting hierarchy for clone Folder1")
    cloneFolder.print("    ")
}
