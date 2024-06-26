package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jumpaku/cyamli/name"
	"github.com/Jumpaku/cyamli/schema"

	"github.com/Jumpaku/go-assert"
)

type Option struct {
	Name    name.Path
	Short   name.Path
	Type    schema.Type
	Default string
}

func (d Option) DefaultLiteral() string {
	switch d.Type {
	default:
		return assert.Unexpected1[string]("unexpected type: %s", d.Type)
	case schema.TypeBoolean:
		if d.Default == "" {
			return "false"
		}
		val, err := strconv.ParseBool(d.Default)
		assert.State(err == nil, "default value must be parsed as bool: %w", err)
		return strconv.FormatBool(val)
	case schema.TypeFloat:
		if d.Default == "" {
			return "0.0"
		}
		val, err := strconv.ParseFloat(d.Default, 64)
		assert.State(err == nil, "default value must be parsed as float64: %w", err)
		return strconv.FormatFloat(val, 'f', -1, 64)
	case schema.TypeInteger:
		if d.Default == "" {
			return "0"
		}
		val, err := strconv.ParseInt(d.Default, 0, 64)
		assert.State(err == nil, "default value must be parsed as int64: %w", err)
		return strconv.FormatInt(val, 10)
	case schema.TypeString, schema.TypeUnspecified:
		return strings.ReplaceAll(fmt.Sprintf(`%q`, d.Default), "$", "\\$")
	}
}

func (d Option) NameLiteral() string {
	return fmt.Sprintf("%q", d.Name.Join("-", "-", ""))
}

func (d Option) ShortNameLiteral() string {
	if len(d.Short) == 0 {
		return ""
	}
	return fmt.Sprintf("%q", d.Short.Join("", "-", ""))
}

func (d Option) InputFieldType() string {
	return DartType(d.Type, false)
}

func (d Option) InputFieldName() string {
	return d.Name.Map(name.Title).Join("", "opt", "")
}
