package slog

import (
	"bytes"
	"fmt"

	"github.com/goadesign/goa"
)

type goaAdapter struct {
	keyvals []interface{}
}

var AdapterGOA goaAdapter

func (a *goaAdapter) Info(msg string, keyvals ...interface{}) {

	n := (len(keyvals) + 1) / 2
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, goa.ErrMissingLogValue)
	}
	m := (len(a.keyvals) + 1) / 2
	n += m
	var fm bytes.Buffer

	fm.WriteString(fmt.Sprintf("%s", msg))
	vals := make([]interface{}, n)
	offset := len(a.keyvals)
	for i := 0; i < offset; i += 2 {
		k := a.keyvals[i]
		v := a.keyvals[i+1]
		vals[i/2] = v
		fm.WriteString(fmt.Sprintf(" %s=%%+v", k))
	}
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		v := keyvals[i+1]
		vals[i/2+offset/2] = v
		fm.WriteString(fmt.Sprintf(" %s=%%+v", k))
	}

	Info(fm.String(), vals...)
}

// Error logs an error.
func (a *goaAdapter) Error(msg string, keyvals ...interface{}) {
	Error(msg, keyvals)
}

// New appends to the logger context and returns the updated logger logger.
func (a *goaAdapter) New(keyvals ...interface{}) goa.LogAdapter {
	if len(keyvals) == 0 {
		return a
	}
	kvs := append(a.keyvals, keyvals...)
	if len(kvs)%2 != 0 {
		kvs = append(kvs, goa.ErrMissingLogValue)
	}
	return &goaAdapter{
		// Limiting the capacity of the stored keyvals ensures that a new
		// backing array is created if the slice must grow.
		keyvals: kvs[:len(kvs):len(kvs)],
	}
}
