# Interfaces

Interfaces in Go (and most languages that offer interfaces) are:

- A set of method signatures

```go
package example

type Runner interface {
	SetStatement(stmt string)
	Run() error
}

```

Types implement interfaces by **implicitly** implementing it's methods.
No need to declare that your Type implements an interface. The type
implements an interface if it simple implements it's methods.

Pointers and functions implements interfaces since they are Types
themselves.

```go
package example

type CassandraRunner struct {
	Config string
}

func (cr *CassandraRunner) SetStatement(stmt string) {
    // do something
}

func (cr *CassandraRunner) Run

```