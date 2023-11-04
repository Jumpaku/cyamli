package data_test

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/cliautor/golang/data"
	"github.com/Jumpaku/cliautor/name"
	"github.com/Jumpaku/cliautor/schema"

	"github.com/stretchr/testify/assert"
)

func TestOption_InputFieldName(t *testing.T) {
	testcases := []struct {
		sut  data.Option
		want string
	}{
		{
			sut: data.Option{
				Name:    name.Path{"opt", "name", "123"},
				Type:    schema.TypeInteger,
				Short:   name.Path{"o"},
				Default: "-123",
			},
			want: `Opt_OptName123`,
		},
		{
			sut: data.Option{
				Name:  name.Path{"optname"},
				Type:  schema.TypeInteger,
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `Opt_Optname`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.InputFieldName()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestOption_InputFieldType(t *testing.T) {
	testcases := []struct {
		sut  data.Option
		want string
	}{
		{
			sut: data.Option{
				Type:  schema.TypeInteger,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `int64`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeBoolean,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `bool`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeFloat,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `float64`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeString,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `string`,
		},
		{
			sut: data.Option{
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `string`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: type=%#v", number, testcase.sut.Type), func(t *testing.T) {
			got := testcase.sut.InputFieldType()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestOption_NameLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Option
		want string
	}{
		{
			sut: data.Option{
				Name:  name.Path{"opt", "name", "123"},
				Type:  schema.TypeInteger,
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `"-opt-name-123"`,
		},
		{
			sut: data.Option{
				Name:  name.Path{"optname"},
				Type:  schema.TypeInteger,
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `"-optname"`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Name), func(t *testing.T) {
			got := testcase.sut.NameLiteral()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestOption_ShortNameLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Option
		want string
	}{
		{
			sut: data.Option{
				Name:  name.Path{"opt", "name", "123"},
				Type:  schema.TypeInteger,
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `"-o"`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Short), func(t *testing.T) {
			got := testcase.sut.ShortNameLiteral()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestOption_DefaultLiteral(t *testing.T) {
	testcases := []struct {
		sut  data.Option
		want string
	}{
		{
			sut: data.Option{
				Type:  schema.TypeInteger,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "-123",
			},
			want: `int64(-123)`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeInteger,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},
			},
			want: `int64(0)`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeBoolean,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "true",
			},
			want: `true`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeBoolean,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "false",
			},
			want: `false`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeBoolean,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},
			},
			want: `false`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeFloat,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "-123.456",
			},
			want: `float64(-123.456)`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeFloat,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},
			},
			want: `float64(0.0)`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeString,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "abc",
			},
			want: `"abc"`,
		},
		{
			sut: data.Option{
				Type:  schema.TypeString,
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},
			},
			want: `""`,
		},
		{
			sut: data.Option{
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},

				Default: "abc",
			},
			want: `"abc"`,
		},
		{
			sut: data.Option{
				Name:  name.Path{"opt", "name", "123"},
				Short: name.Path{"o"},
			},
			want: `""`,
		},
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf("%03d: %#v", number, testcase.sut.Default), func(t *testing.T) {
			got := testcase.sut.DefaultLiteral()
			assert.Equal(t, testcase.want, got)
		})
	}
}
