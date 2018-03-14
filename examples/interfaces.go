package main

import (
	"fmt"
	"errors"
)

type Runner interface {
	SetStatement(stmt string)
	Run() error
}

type MySQLRunner struct {
	Statement string
}

type CassandraRunner struct {
	Statement string
}

func NewCassandraRunner() *CassandraRunner {
	return new(CassandraRunner)
}

func (cr *CassandraRunner) SetStatement(stmt string) {
	cr.Statement = stmt
}

func (cr *CassandraRunner) Run() error {
	if cr.Statement != "" {
		fmt.Printf("Running statement: %s\n", cr.Statement)
		return nil
	}

	return errors.New("no statement");
}

func (cr *CassandraRunner) SpecialOp() string {
	return "Hi"
}

func Process(runner Runner) error {
	err := runner.Run()
	return err
}

func main() {
	var rnr Runner
	rnr = NewCassandraRunner()
	err := Process(rnr)
	if err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}

	rnr.SetStatement("Hello World")
	err = Process(rnr)
	if err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}

	// check for implementation example
	if _, ok := rnr.(Runner); ok {
		fmt.Printf("Im a Runner\n")
	}

	// check for implementation example
	if cr, ok := rnr.(*CassandraRunner); ok {
		fmt.Printf("Im a Cassandra Runner: %s\n", cr.SpecialOp())
	}

}
