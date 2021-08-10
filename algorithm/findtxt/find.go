package find

func Find(txt, word string) int {
    lenWord := len(word)
    lenTxt := len(txt)
    
    for i := 0; i<= lenTxt-lenWord; i++ {
        if txt[i:i+lenWord] == word {
            return i
        }
    }
    return -1
}

