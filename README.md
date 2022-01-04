go-fsm-example
==============

An example FSM implementation in Go using State functions (or what I understood from [Lexical Scanning in Go](https://talks.golang.org/2011/lex.slide#16))

The idea is simple, each state function returns the next state.
```go
// stateFn represents the state of the machine
// as a function that returns the next state.
type stateFn func(*fsm) stateFn
```

Then each state is defined as a function that receives a `*fsm` and return a `stateFn`

```go
func idleState(f *fsm) stateFn {
	f.code = ""
	fmt.Printf("Enter code: ")
	return readState
}
```
In this example I'm creating a fsm that verify some secret code.

To represent the `done` state a `stateFn` must return `nil` and the machine stops, the main loop looks like:

```go
func (self *fsm) run() {
	self.reader = bufio.NewReader(os.Stdin)
	for s := idleState; s != nil; {
		s = s(self)
	}
}
```

And that's it.
