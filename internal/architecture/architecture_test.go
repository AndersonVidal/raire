package architecture_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"testing"
)

type listedPackage struct {
	ImportPath string
	Imports    []string
	Standard   bool
}

const module = "github.com/raire/raire/"

func loadPackages(t *testing.T) []listedPackage {
	t.Helper()
	cmd := exec.Command("go", "list", "-json", "./...")
	cmd.Dir = "../.."
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			t.Fatalf("go list failed: %s", ee.Stderr)
		}
		t.Fatal(err)
	}
	dec := json.NewDecoder(bytes.NewReader(out))
	var pkgs []listedPackage
	for {
		var p listedPackage
		if err := dec.Decode(&p); err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		pkgs = append(pkgs, p)
	}
	return pkgs
}

func contextOf(importPath string) string {
	rest := strings.TrimPrefix(importPath, module+"internal/")
	parts := strings.Split(rest, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}

func TestDependencyRules(t *testing.T) {
	pkgs := loadPackages(t)
	if len(pkgs) == 0 {
		t.Fatal("no packages listed")
	}
	businessContexts := map[string]bool{
		"identity": true, "catalog": true, "directory": true, "demand": true,
		"discovery": true, "fairness": true, "reputation": true, "moderation": true,
		"api": true, "app": true, "web": true,
	}
	var failures []string
	for _, p := range pkgs {
		if strings.HasPrefix(p.ImportPath, module+"internal/platform/") {
			for _, imp := range p.Imports {
				if !strings.HasPrefix(imp, module+"internal/") {
					continue
				}
				ctx := contextOf(imp)
				if businessContexts[ctx] {
					failures = append(failures, fmt.Sprintf("platform package %s imports business package %s", p.ImportPath, imp))
				}
			}
		}
		if strings.HasSuffix(p.ImportPath, "/domain") {
			for _, imp := range p.Imports {
				if imp == "net/http" || imp == "database/sql" ||
					strings.Contains(imp, "/adapters/") || strings.Contains(imp, "internal/platform/db") {
					failures = append(failures, fmt.Sprintf("domain package %s imports infrastructure %s", p.ImportPath, imp))
				}
			}
		}
		if p.ImportPath == module+"internal/contracts" {
			for _, imp := range p.Imports {
				if strings.HasPrefix(imp, module) {
					failures = append(failures, fmt.Sprintf("contracts imports project package %s", imp))
				}
			}
		}
		if strings.Contains(p.ImportPath, "/adapters/") {
			importerCtx := contextOf(p.ImportPath)
			for _, imp := range p.Imports {
				if !strings.Contains(imp, "/adapters/") {
					continue
				}
				adapterCtx := contextOf(imp)
				if p.ImportPath != module+"internal/app" && importerCtx != adapterCtx {
					failures = append(failures, fmt.Sprintf("package %s imports foreign adapter %s", p.ImportPath, imp))
				}
			}
		}
		if p.ImportPath == module+"internal/api" {
			for _, imp := range p.Imports {
				if strings.Contains(imp, "/adapters/") || strings.HasSuffix(imp, "/platform/db") {
					failures = append(failures, fmt.Sprintf("api imports persistence directly: %s", imp))
				}
			}
		}
	}
	if len(failures) > 0 {
		t.Fatalf("dependency rule violations:\n%s", strings.Join(failures, "\n"))
	}
}
