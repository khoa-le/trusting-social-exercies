package file

import (
	"fmt"
	"testing"
)

func TestOpenAndReadFile(t *testing.T){
	data := OpenAndReadFile("../data.csv")
	if len(data) == 0{
		t.Errorf("can not parse data on csv file or empty file")
	}
	fmt.Printf("%v",data)
}
