package main

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/google/go-cmp/cmp"
)

type issue struct {
	actual   []string       // the actual order of fields
	expected []string       // the expected order of fields
	position token.Position // file, line, and column of the issue
}

func (self issue) String() string {
	return fmt.Sprintf(
		"%s:%d:%d unsorted fields, expected order:\n\n%s\n\nDiff (-got +want)\n\n%s\n\n",
		self.position.Filename,
		self.position.Line,
		self.position.Column,
		strings.Join(self.expected, "\n"),
		cmp.Diff(self.actual, self.expected),
	)
}

type Issues []issue

func (self Issues) String() string {
	result := strings.Builder{}
	for _, issue := range self {
		result.WriteString(issue.String())
	}
	return result.String()
}
