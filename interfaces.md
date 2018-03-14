[Golang Index](http://go.imti.co/)

# Interfaces

Interfaces in Go (and most languages that offer interfaces) are:

- A set of method signatures

```go
package main

type Runner interface {
	SetStatement(stmt string)
	Run() error
}

```

Types implement an interface by **implicitly** implementing it's methods.
No need to declare that your Type implements an interface. The type
implements an interface if it simply implements it's methods.

Pointers and functions implement interfaces as they are Types
themselves.

```go
package main

type Runner interface {
	SetStatement(stmt string)
	Run() error
}

type CassandraRunner struct {
	Config string
}

func (cr *CassandraRunner) SetStatement(stmt string) {
    // do something
}

func (cr *CassandraRunner) Run() {
    // do something
}

```

## Example Code

- [Go Playground](https://play.golang.org/p/s5e28KLzs1z)
- [Source: interfaces.go](examples/interfaces.go)

## Resources

- [Effective Go: Interfaces](https://golang.org/doc/effective_go.html#interfaces_and_types)
- [Go by Example: Interfaces](https://gobyexample.com/interfaces)
- Best Practices / Patterns
    - [Accept Interfaces and return concrete types](http://idiomaticgo.com/post/best-practice/accept-interfaces-return-structs/)