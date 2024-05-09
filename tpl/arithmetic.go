package tpl

import "text/template"

func ArithmeticFuncs() template.FuncMap {
	return template.FuncMap{
		"func_add": Add,
		"func_sub": Sub,
	}
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a + b
}
