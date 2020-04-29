package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const logfile = "example.log"

func openFile(file string) error {
	fd, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fd.Close()

	rd := bufio.NewReader(fd)
	for {
		if line, err := rd.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			}
		} else {
			// do something here with eadh line
			fmt.Printf(">> %s", line)
		}
	}

	return nil
}

func main() {
	openFile(logfile)
}
