package protein

import ()

var ErrStop error

var ErrInvalidBase error

func FromCodon(s string) (string, error) {
	return tcodon[s], nil
}

func FromRNA(s string) ([]string, error) {
	list := []string{}
	for i := 0; i < len(s); i += 3 {
		code := tcodon[s[i:i+3]]
		if code == tErrStop[code] {
			break
		}
		list = append(list, tcodon[s[i:i+3]])
	}
	return list, nil
}

var tcodon = map[string]string{
	//Codon 	Protein
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
}

var tErrStop = map[string]string{
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}
