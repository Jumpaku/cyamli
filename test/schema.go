package test

import (
	"cliautor/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertMatchSchema(t *testing.T, want, got *schema.Schema) {
	wp, gp := want.Program, got.Program
	assert.Equal(t, wp.Name, gp.Name)
	assert.Equal(t, wp.Version, gp.Version)
	AssertMatchCommand(t, wp.Command(), gp.Command())
}

func AssertMatchCommand(t *testing.T, want, got *schema.Command) {
	assert.Equal(t, want.Description, got.Description)

	wOpt, gOpt := want.Options, got.Options
	for name, w := range wOpt {
		g, ok := gOpt[name]
		assert.Truef(t, ok, "option %q not found", name)
		assert.Equal(t, w.Short, g.Short)
		assert.Equal(t, w.Default, g.Default)
		assert.Equal(t, w.Description, g.Description)
		assert.Equal(t, w.Type, g.Type)
	}

	wArg, gArg := want.Arguments, got.Arguments
	for idx, w := range wArg {
		assert.Truef(t, idx < len(gArg), "argument %q not found at %d", w.Name, idx)
		g := gArg[idx]
		assert.Equal(t, w.Name, g.Name)
		assert.Equal(t, w.Variadic, g.Variadic)
		assert.Equal(t, w.Description, g.Description)
		assert.Equal(t, w.Type, g.Type)
	}

	wSub, gSub := want.Subcommands, got.Subcommands
	for name, w := range wSub {
		g, ok := gSub[name]
		assert.Truef(t, ok, "subcommand %q not found", name)
		AssertMatchCommand(t, w, g)
	}
}
