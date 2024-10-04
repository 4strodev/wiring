package mocks

import (
	"errors"
	"fmt"
	"io"
)

const DEFAULT_MESSAGE = "Hello world"
const TESTING_TOKEN = "token"

type Abstraction interface {
	Greet()
}

type Implementation struct {
	Message string
}

func (i *Implementation) Greet() {
	fmt.Println(i.Message)
}

func Resolver() Abstraction {
	return &Implementation{
		DEFAULT_MESSAGE,
	}
}

func TokenResolver() string {
	return DEFAULT_MESSAGE
}

func ResolverWithMessage(message string) func() Abstraction {
	return func() Abstraction {
		return &Implementation{
			message,
		}
	}
}

type InvalidStruct struct {
	TokenResolved string      `wire:"token"` // This should be resolved by token
	TypeResolved  Abstraction // This should be resolved by their type
	Reader        io.Reader
}

type FillableStruct struct {
	message       string      // This value should be ignored by the contianer
	Message       string      `wire:",ignore"` // This also should be ignored by the container
	TokenResolved string      `wire:"token"`   // This should be resolved by token
	TypeResolved  Abstraction // This should be resolved by their type
}

func (s FillableStruct) CheckResolvedFields() error {
	if s.message != "" {
		return errors.New("'message' field has no default value")
	}

	if s.Message != "" {
		return errors.New("'Message' field has no default value")
	}

	if s.TokenResolved == "" {
		return errors.New("'TokenResolved' was not resolved")
	}

	if s.TypeResolved == nil {
		return errors.New("'TypeResolved' was not resolved")
	}

	return nil
}
