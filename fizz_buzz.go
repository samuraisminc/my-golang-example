package main

import "fmt"

type FizzBuzzValue interface {
	Text() string
}

type SimpleText struct {
	text string
}

func (st *SimpleText) Text() string {
	return st.text
}

type Value struct {
	Value    FizzBuzzValue
	Original int
}

func ToValue(original int) Value {
	return Value{
		Value:    nil,
		Original: original,
	}
}

func (v Value) String() string {
	if v.Value == nil {
		return fmt.Sprintf("%d", v.Original)
	}
	return v.Value.Text()
}
