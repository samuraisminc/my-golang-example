package main

import "strings"

type Rule interface {
	Apply(value Value) Value
}

type SimpleRule struct {
	base int
	text string
}

func (sr *SimpleRule) Apply(value Value) Value {
	if value.Original%sr.base != 0 {
		return value
	}
	if value.Value == nil {
		return Value{
			Value:    &SimpleText{text: sr.text},
			Original: value.Original,
		}
	}
	var sb strings.Builder
	sb.WriteString(value.Value.Text())
	sb.WriteString(sr.text)
	return Value{
		Value:    &SimpleText{text: sb.String()},
		Original: value.Original,
	}
}

type Rules []Rule

func (rs Rules) Apply(value Value) Value {
	v := value
	for _, r := range rs {
		v = r.Apply(v)
	}
	return v
}
