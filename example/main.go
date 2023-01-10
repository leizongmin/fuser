package main

import (
	"fmt"

	"github.com/leizongmin/fuser"
)

func main() {
	data, err := fuser.BuildMap(nil)
	if err != nil {
		panic(err)
	}
	for k, v := range data {
		fmt.Println(k, v)
	}

	err = fuser.Update(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(fuser.GetPath("/dev/null"))

	fuser.Update(&fuser.Options{
		Filter: func(s string) bool {
			return true
		},
	})
}
