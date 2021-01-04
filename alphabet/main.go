package main

import (
	"fmt"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func main() {
	cities := []string{
		"Berlin",
		"Zurich",
		"Augsburg",
		"Bünde",
		"Budapest",
		"Ürkmez",
		"Rostock",
		"Ulm",
		"Lindau",
	}

	c := collate.New(language.German, collate.IgnoreCase)
	c.SortStrings(cities)
	fmt.Println(cities)

	mixedLanguagesCities := []string{
		"Ürkmez",
		"Budapest",
		"Бохольт",
		"Арнсберг",
		"Інцель",
		"Їндржихув-Градец",
		"Єна",
		"Шатору",
		"Ястшембя-Ґура",
		"Ґрудзьондз",
		"Атланта",
		"Zurich",
	}

	c = collate.New(language.Und, collate.IgnoreCase)
	c.SortStrings(mixedLanguagesCities)
	fmt.Println(mixedLanguagesCities)
}
