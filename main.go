package main

import (
	"bufio"
	"fmt"
	"os"
)

type ExquisiteApp struct {
	reader    *bufio.Reader
	delimeter byte
}

func (x *ExquisiteApp) ReadString(delim byte) (line string, err error) {

	return x.reader.ReadString(delim)
}

func (x *ExquisiteApp) run() {
	println("Running ExquisiteCorpse...")
	text, err := x.ReadString(x.delimeter)
	if err != nil {
		println("we received an error.")
		println(err)
	} else {
		fmt.Printf("Received text: %s", text)
	}
}

func NewExquisiteApp() *ExquisiteApp {
	return &ExquisiteApp{reader: bufio.NewReader(os.Stdin), delimeter: '\n'}
}
func main() {

	var app = *NewExquisiteApp()
	app.run()

}

func getInput() {

}
