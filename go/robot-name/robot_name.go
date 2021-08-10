package robotname

import (
	"errors"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var storename = map[string]string{}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(storename) == 676000 { // The number of possibilities for unique names
		return "", errors.New("not new name available")
	}
	newName := randomName()
	for {
		if newName == storename[newName] {
			newName = randomName()
		} else {
			break
		}
	}

	storename[newName] = newName
	r.name = newName
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func randomName() string {
	rand.Seed(time.Now().UnixNano())
	randName := string(65 + rand.Intn(90-65))
	randName += string(65 + rand.Intn(90-65))
	randName += string(48 + rand.Intn(57-48))
	randName += string(48 + rand.Intn(57-48))
	randName += string(48 + rand.Intn(57-48))
	return randName
}
