package main

import(
	"bufio"
	"fmt"
	"os"
)

const (
	SECRET = "gchaincl"
)

type fsm struct {
	reader *bufio.Reader
	code string
}

type stateFn func(*fsm) stateFn

func idleState(f *fsm) stateFn {
	f.code = ""
	fmt.Printf("Enter code: ")
	return readState
}

func readState(f *fsm) stateFn {
	input, _ := f.reader.ReadByte()

	// verify when key is a new line
	if (input == 10) {
		return verifyState
	}

	f.code = fmt.Sprintf("%s%c", f.code, input)
	return readState
}

func verifyState(f *fsm) stateFn {
	fmt.Printf("Verifying code...")
	if f.code != SECRET {
		fmt.Println("Invalid!")
		return idleState
	}
	return finishState
}

func finishState(f *fsm) stateFn {
	fmt.Println("Success!")
	return nil
}

func (self *fsm) run() {
	self.reader = bufio.NewReader(os.Stdin)
	for s := idleState; s != nil; {
		s = s(self)
	}
}

func main() {
	f := fsm{}
	f.run()
}