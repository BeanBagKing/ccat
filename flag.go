package main

import (
	"fmt"
	"strings"
)

type mapValue map[string]string

func (m mapValue) Type() string {
    return "map"
}

func (m mapValue) Set(val string) error {
	v := strings.SplitN(val, "=", 2)
	if len(v) != 2 {
		return fmt.Errorf("invalid flag format: expected '<name>=<value>', got '%s'", val)
	}

	m[v[0]] = v[1]

	return nil
}

func (m mapValue) String() string {
	s := make([]string, 0, len(m))
	for k, v := range m {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}

	var sb strings.Builder
	for i, pair := range s {
	    if i > 0 {
	        sb.WriteString(",")
	    }
	    sb.WriteString(pair)
	}
	return sb.String()
}

func (m *mapValue) UnmarshalText(text []byte) error {
    return m.Set(string(text))
}

