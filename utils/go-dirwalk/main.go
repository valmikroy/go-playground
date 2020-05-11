package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func DirWalk(dir string) error {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {

		var path string

		if dir == "/" {
			path = dir + f.Name()
		} else {
			path = dir + "/" + f.Name()
		}

		if f.IsDir() {
			Walk(path)
		}

		// Do something here with each file
		fmt.Println(path)
	}

	return nil
}

func main() {

	err := DirWalk("/")
	if err != nil {
		log.Fatalln(err)
	}

}
