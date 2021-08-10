package main

import "fmt"

func getGun(gunType string) (iGun, error) {
    if gunType == "aka47" {
        return newAka47(), nil
    }
    if gunType == "maverick" {
        return newMaverick(), nil
    }
    return nil, fmt.Errorf("Wrong gun type passed")
}


