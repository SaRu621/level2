package pkg

import (
	"fmt"
	"testing"
)

var Tests = []string{
	"a4bc2d5e",
	"abcd",
	"45",
	"",
	"\\45",
	"ab2c\\34d",
}

var Ans = []string{
	"aaaabccddddde",
	"abcd",
	"",
	"",
	"44444",
	"abbc3333d",
}

func TestPkg(t *testing.T) {
	for i := range Tests {
		if Ans[i] != string(Extraction(Tests[i])) {
			t.Errorf("Не пройден [%d]-й тест.\n", i)
			fmt.Println(string(Extraction(Tests[i])))
		} else {
			fmt.Printf("Пройден [%d]-й.\n", i)
		}
	}
}
