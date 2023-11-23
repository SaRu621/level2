package pkg

import (
	"fmt"
	"io/ioutil"
	"task3/pkg/structs"
	"testing"
)

var testFolder = "/home/rus/GoFolder/level2/develop/task3/test/"

var testConfig = []structs.Config{
	{
		SortedColumn: 2,
		NumSort:      false,
		Reversed:     false,
		Unique:       false,

		InputFile:  testFolder + "task/test1.txt",
		OutputFile: testFolder + "ans/test1.txt",
	},
	{
		SortedColumn: 2,
		NumSort:      true,
		Reversed:     false,
		Unique:       false,

		InputFile:  testFolder + "task/test2.txt",
		OutputFile: testFolder + "ans/test2.txt",
	},
	{
		SortedColumn: 2,
		NumSort:      true,
		Reversed:     false,
		Unique:       true,

		InputFile:  testFolder + "task/test3.txt",
		OutputFile: testFolder + "ans/test3.txt",
	},
	{
		SortedColumn: 2,
		NumSort:      true,
		Reversed:     true,
		Unique:       true,

		InputFile:  testFolder + "task/test4.txt",
		OutputFile: testFolder + "ans/test4.txt",
	},
}

var testExp = []string{
	testFolder + "exp/test1.txt",
	testFolder + "exp/test2.txt",
	testFolder + "exp/test3.txt",
	testFolder + "exp/test4.txt",
}

func compareFiles(file1 string, file2 string) (bool, error) {
	content1, err := ioutil.ReadFile(file1)
	if err != nil {
		return false, err
	}

	content2, err := ioutil.ReadFile(file2)
	if err != nil {
		return false, err
	}

	return string(content1) == string(content2), nil
}

func TestPkg(t *testing.T) {
	for i := range testExp {
		ManSort(testConfig[i])
		equal, err := compareFiles(testExp[i], testConfig[i].OutputFile)

		if err != nil {
			t.Errorf("Test not passed: %s\n", err)
		}

		if !equal {
			fmt.Println("Test not passed: files not equals")
		}

		fmt.Printf("Test [%d] passed.\n", i)
	}
}
