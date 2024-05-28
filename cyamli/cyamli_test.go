package cyamli_test

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/Jumpaku/cyamli/cyamli"
	"github.com/kylelemons/godebug/diff"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var updateGolden = flag.Bool("update", false, "updates .golden files")

func mustOpen(t *testing.T, path string) *os.File {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	return f
}

func mustWrite(t *testing.T, path string, src string) {
	t.Helper()

	dst := mustCreate(t, path)
	defer dst.Close()

	_, err := dst.WriteString(src)
	if err != nil {
		t.Fatal(err)
	}
}

func mustRead(t *testing.T, path string) string {
	t.Helper()

	src := mustOpen(t, path)
	defer src.Close()

	buf := &bytes.Buffer{}

	_, err := io.Copy(buf, src)
	if err != nil {
		t.Fatal(err)
	}

	return buf.String()
}

func mustCreate(t *testing.T, path string) *os.File {
	t.Helper()
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	return f
}

func buildDiff(t *testing.T, want, got string) string {
	t.Helper()
	lines := strings.Split(diff.Diff(want, got), "\n")
	diffText := ""
	n := len(lines)
	r := 3
	added := true
	for i, line := range lines {
		var print bool
		for j := i - r; j <= i+r; j++ {
			if j >= 0 && j < n && (strings.HasPrefix(lines[j], "+") || strings.HasPrefix(lines[j], "-")) {
				print = true
			}
		}
		if print {
			diffText += line + "\n"
			added = true
		} else if added {
			diffText += "\n...\n\n"
			added = false
		}
	}
	return diffText
}

func TestExecute_stdio(t *testing.T) {
	tests := []struct {
		args       []string
		inFileName string
		wantCode   int
	}{
		// root
		{
			args:       []string{"cyamli"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "-version"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "-v"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "-help"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "-h"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "-unknown"},
			inFileName: "cyamli.yaml",
			wantCode:   1,
		},
		{
			args:       []string{"cyamli", "invalid"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},

		// list
		{
			args:       []string{"cyamli", "list"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "list"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "list"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "list", "-help"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},

		// generate
		{
			args:       []string{"cyamli", "generate"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},

		// generate golang
		{
			args:       []string{"cyamli", "generate", "golang"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "golang", "-help"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "golang", "-package=xyz"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "golang", "-package=xyz"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "golang", "-package=xyz"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},

		// generate python3
		{
			args:       []string{"cyamli", "generate", "python3"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "python3", "-help"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "python3"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "python3"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},

		// generate docs
		{
			args:       []string{"cyamli", "generate", "docs", "-help"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=text"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=html"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=markdown"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=text"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=html"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=markdown"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=text"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=html"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=markdown"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "generate"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "generate", "docs"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},

		// validate
		{
			args:       []string{"cyamli", "validate"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "validate"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "validate"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "validate", "-help"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
	}
	for number, tt := range tests {
		name := fmt.Sprintf(`stdio_%02d`, number)
		t.Run(name, func(t *testing.T) {
			stdin := mustOpen(t, filepath.Join("testdata", tt.inFileName))
			defer stdin.Close()

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}

			gotCode := cyamli.Execute(tt.args, stdin, stdout, stderr)

			if *updateGolden {
				mustWrite(t, filepath.Join("testdata", name+".golden.stdout"), stdout.String())
				mustWrite(t, filepath.Join("testdata", name+".golden.stderr"), stderr.String())
			}

			wantStdout := mustRead(t, filepath.Join("testdata", name+".golden.stdout"))
			wantStderr := mustRead(t, filepath.Join("testdata", name+".golden.stderr"))
			gotStdout := stdout.String()
			gotStderr := stderr.String()
			if wantStdout != gotStdout {
				t.Errorf("Execute() stdout mismatch\n%s", buildDiff(t, wantStdout, gotStdout))
			}
			if wantStderr != gotStderr {
				t.Errorf("Execute() stdout mismatch\n%s", buildDiff(t, wantStderr, gotStderr))
			}
			if tt.wantCode != gotCode {
				t.Errorf("Execute() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func TestExecute_fileio(t *testing.T) {
	tests := []struct {
		args       []string
		inFileName string
		wantCode   int
	}{
		// generate golang
		{
			args:       []string{"cyamli", "generate", "golang"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "golang", "-package=xyz"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "golang", "-package=xyz"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "golang", "-package=xyz"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},

		// generate python3
		{
			args:       []string{"cyamli", "generate", "python3"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "python3"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "python3"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},

		// generate docs
		{
			args:       []string{"cyamli", "generate", "docs", "-all"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=text"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=html"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=markdown"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=text"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=html"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=markdown"},
			inFileName: "empty.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=text"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=html"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "-all", "-format=markdown"},
			inFileName: "demo-app.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "generate"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
		{
			args:       []string{"cyamli", "generate", "docs", "generate", "docs"},
			inFileName: "cyamli.yaml",
			wantCode:   0,
		},
	}
	for number, tt := range tests {
		name := fmt.Sprintf(`fileio_%02d`, number)
		t.Run(name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}

			tmp := "fileio.tmp"
			gotCode := cyamli.Execute(
				append(tt.args,
					"-schema-path="+filepath.Join("testdata", tt.inFileName),
					"-out-path="+filepath.Join("testdata", tmp)),
				&bytes.Buffer{}, stdout, stderr)
			gotOutput := mustRead(t, filepath.Join("testdata", tmp))

			if *updateGolden {
				mustWrite(t, filepath.Join("testdata", name+".golden.output"), gotOutput)
				mustWrite(t, filepath.Join("testdata", name+".golden.stderr"), stderr.String())
			}

			wantOutput := mustRead(t, filepath.Join("testdata", name+".golden.output"))
			if wantOutput != gotOutput {
				t.Errorf("Execute() output mismatch\n%s", buildDiff(t, wantOutput, gotOutput))
			}
			wantStderr := mustRead(t, filepath.Join("testdata", name+".golden.stderr"))
			gotStdout := stdout.String()
			if "" != gotStdout {
				t.Errorf("Execute() stdout mismatch\n%s", buildDiff(t, "", gotStdout))
			}
			gotStderr := stderr.String()
			if wantStderr != gotStderr {
				t.Errorf("Execute() stderr mismatch\n%s", buildDiff(t, wantStderr, gotStderr))
			}
			if tt.wantCode != gotCode {
				t.Errorf("Execute() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
