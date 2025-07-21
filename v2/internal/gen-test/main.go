package main

import (
	_ "embed"
	"fmt"
	"github.com/Jumpaku/cyamli/v2/generate/cpp"
	"github.com/Jumpaku/cyamli/v2/generate/csharp"
	"github.com/Jumpaku/cyamli/v2/generate/dart3"
	"github.com/Jumpaku/cyamli/v2/generate/golang"
	"github.com/Jumpaku/cyamli/v2/generate/kotlin"
	"github.com/Jumpaku/cyamli/v2/generate/php"
	"github.com/Jumpaku/cyamli/v2/generate/python3"
	"github.com/Jumpaku/cyamli/v2/generate/typescript"
	"github.com/Jumpaku/cyamli/v2/schema"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s [target] [outDir] < [schema.yaml]\n", os.Args[0])
		os.Exit(1)
	}
	target := os.Args[1]
	outDir := os.Args[2]

	s, err := schema.Load(os.Stdin)
	if err != nil {
		panic(fmt.Sprintf("fail to load schema: %+v", err))
	}

	switch target {
	default:
		fmt.Fprintf(os.Stderr, "Unknown target: %s\n", target)
		os.Exit(1)
	case "golang":
		{
			if len(os.Args) != 4 {
				fmt.Fprintf(os.Stderr, "Usage: %s [target] [outDir] [package] < [schema.yaml]\n", os.Args[0])
				os.Exit(1)
			}
			packageName := os.Args[3]
			t, err := os.Create(filepath.Join(outDir, "cli.gen_test.go"))
			if err != nil {
				panic(fmt.Sprintf("fail to create output file: %+v", err))
			}
			defer t.Close()
			if err = golang.GenerateTest(s, "cyamli/test/golang", packageName, "cyamli", t); err != nil {
				panic(fmt.Sprintf("fail to generate test code: %+v", err))
			}
		}
	case "cpp":
		{
			if len(os.Args) != 4 {
				fmt.Fprintf(os.Stderr, "Usage: %s [target] [outDir] [package] < [schema.yaml]\n", os.Args[0])
				os.Exit(1)
			}
			namespace := os.Args[3]
			t, err := os.Create(filepath.Join(outDir, "cli_test.gen.cpp"))
			if err != nil {
				panic(fmt.Sprintf("fail to create output file: %+v", err))
			}
			defer t.Close()
			if err = cpp.GenerateTestCpp(s, "cli.gen.h", namespace, "cyamli", t); err != nil {
				panic(err)
			}
		}
	case "csharp":
		{
			if len(os.Args) != 4 {
				fmt.Fprintf(os.Stderr, "Usage: %s [target] [outDir] [package] < [schema.yaml]\n", os.Args[0])
				os.Exit(1)
			}
			namespace := os.Args[3]
			t, err := os.Create(filepath.Join(outDir, "Cli_GenTest.cs"))
			if err != nil {
				panic(fmt.Sprintf("fail to create output file: %+v", err))
			}
			defer t.Close()
			if err = csharp.GenerateTest(s, namespace, "cyamli", t); err != nil {
				panic(err)
			}
		}
	case "dart3":
		{
			if len(os.Args) != 4 {
				fmt.Fprintf(os.Stderr, "Usage: %s [target] [outDir] [package] < [schema.yaml]\n", os.Args[0])
				os.Exit(1)
			}
			packageName := os.Args[3]
			t, err := os.Create(filepath.Join(outDir, "cli.g_test.dart"))
			if err != nil {
				panic(fmt.Sprintf("fail to create output file: %+v", err))
			}
			defer t.Close()
			if err = dart3.GenerateTest(s, packageName, "cyamli", t); err != nil {
				panic(err)
			}
		}
	case "kotlin":
		{
			if len(os.Args) != 4 {
				fmt.Fprintf(os.Stderr, "Usage: %s [target] [outDir] [package] < [schema.yaml]\n", os.Args[0])
				os.Exit(1)
			}
			packageName := os.Args[3]
			t, err := os.Create(filepath.Join(outDir, "CliTest.kt"))
			if err != nil {
				panic(fmt.Sprintf("fail to create output file: %+v", err))
			}
			defer t.Close()
			if err = kotlin.GenerateTest(s, packageName, "cyamli", t); err != nil {
				panic(err)
			}
		}
	case "php":
		{
			if len(os.Args) != 5 {
				fmt.Fprintf(os.Stderr, "Usage: %s [target] [outDir] [namespace] [testNamespace] < [schema.yaml]\n", os.Args[0])
				os.Exit(1)
			}
			namespace := os.Args[3]
			testNamespace := os.Args[4]
			if err = php.GenerateTest(s, namespace, testNamespace, "cyamli", namedWriter{Directory: outDir}); err != nil {
				panic(err)
			}
		}
	case "python3":
		{
			if len(os.Args) != 4 {
				fmt.Fprintf(os.Stderr, "Usage: %s [target] [outDir] [module] < [schema.yaml]\n", os.Args[0])
				os.Exit(1)
			}
			module := os.Args[3]
			t, err := os.Create(filepath.Join(outDir, "test_cli_gen.py"))
			if err != nil {
				panic(fmt.Sprintf("fail to create output file: %+v", err))
			}
			defer t.Close()
			if err = python3.GenerateTest(s, module, "cyamli", t); err != nil {
				panic(err)
			}
		}
	case "typescript":
		{
			if len(os.Args) != 4 {
				fmt.Fprintf(os.Stderr, "Usage: %s [target] [outDir] [module] < [schema.yaml]\n", os.Args[0])
				os.Exit(1)
			}
			module := os.Args[3]
			t, err := os.Create(filepath.Join(outDir, "cli_gen.test.ts"))
			if err != nil {
				panic(fmt.Sprintf("fail to create output file: %+v", err))
			}
			defer t.Close()
			if err = typescript.GenerateTest(s, module, "cyamli", t); err != nil {
				panic(err)
			}
		}
	}
	/*
		if err = typescript.GenerateTest(s, "./cli.gen.mjs", "cyamli", t); err != nil {
			panic(err)
		}
	*/
}

type namedWriter struct {
	Directory string
}

func (w namedWriter) Write(name string, b []byte) (int, error) {
	f, err := os.Create(filepath.Join(w.Directory, name))
	if err != nil {
		return 0, fmt.Errorf("fail to create file %q: %w", name, err)
	}
	defer f.Close()
	return f.Write(b)
}
