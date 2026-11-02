package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type StructTemp struct {
	name  string
	attrs []string
}

var (
	pkgName   = "ast"
	importPkg = []string{"golox/lox/token"}
)

func main() {
	args := os.Args
	var errs []error
	if len(args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: gen_ast <output directory>")
		os.Exit(64)
	}
	outputDir := args[1]
	if !strings.HasSuffix(outputDir, "/") {
		outputDir += "/"
	}
	{
		path := outputDir + "expr.go"
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
		if _, err := f.WriteString("package "); err != nil {
			errs = append(errs, err)
		}
		if _, err := f.WriteString(pkgName); err != nil {
			errs = append(errs, err)
		}
		if _, err := f.WriteString("\n"); err != nil {
			errs = append(errs, err)
		}

		if _, err := f.WriteString("import(\n "); err != nil {
			errs = append(errs, err)
		}
		for _, pkg := range importPkg {
			if _, err := f.WriteString("\""); err != nil {
				errs = append(errs, err)
			}
			if _, err := f.WriteString(pkg); err != nil {
				errs = append(errs, err)
			}
			if _, err := f.WriteString("\"\n"); err != nil {
				errs = append(errs, err)
			}
		}
		if _, err := f.WriteString(")\n"); err != nil {
			errs = append(errs, err)
		}

		if _, err := f.WriteString(comment()); err != nil {
			errs = append(errs, err)
		}
		if _, err := f.WriteString(buildExpr()); err != nil {
			errs = append(errs, err)
		}

		_ = f.Close()
		cmd := exec.Command("gofmt", "-w", path)
		_ = cmd.Run()
	}
	{
		path := outputDir + "stmt.go"
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
		if _, err := f.WriteString("package "); err != nil {
			errs = append(errs, err)
		}
		if _, err := f.WriteString(pkgName); err != nil {
			errs = append(errs, err)
		}
		if _, err := f.WriteString("\n"); err != nil {
			errs = append(errs, err)
		}

		if _, err := f.WriteString("import(\n "); err != nil {
			errs = append(errs, err)
		}
		for _, pkg := range importPkg {
			if _, err := f.WriteString("\""); err != nil {
				errs = append(errs, err)
			}
			if _, err := f.WriteString(pkg); err != nil {
				errs = append(errs, err)
			}
			if _, err := f.WriteString("\"\n"); err != nil {
				errs = append(errs, err)
			}
		}
		if _, err := f.WriteString(")\n"); err != nil {
			errs = append(errs, err)
		}

		if _, err := f.WriteString(buildStmt()); err != nil {
			errs = append(errs, err)
		}
		_ = f.Close()
		cmd := exec.Command("gofmt", "-w", path)
		_ = cmd.Run()
	}
	for _, err := range errs {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
func comment() string {
	return `
// literal        → NUMBER | STRING | "true" | "false" | "nil" ;
// grouping       → "(" expression ")" ;
// unary          → ( "-" | "!" ) expression ;
// binary         → expression operator expression ;
// ternary        → expression operator expression operator expression ;
// operator       → "==" | "!=" | "<" | "<=" | ">" | ">="

`
}

// Expr
func buildExpr() string {
	var b strings.Builder
	b.WriteString(defineParent("Expr"))
	child := []StructTemp{}
	child = append(child, StructTemp{
		name:  "Assign",
		attrs: []string{"Name token.Token", "Value Expr"},
	})
	child = append(child, StructTemp{
		name:  "Binary",
		attrs: []string{"Left Expr", "Operator token.Token", "Right Expr"},
	})
	child = append(child, StructTemp{
		name:  "Grouping",
		attrs: []string{"Expression Expr"},
	})
	child = append(child, StructTemp{
		name:  "Literal",
		attrs: []string{"Value any"},
	})
	child = append(child, StructTemp{
		name:  "Logical",
		attrs: []string{"Left Expr", "Operator token.Token", "Right Expr"},
	})
	child = append(child, StructTemp{
		name:  "Unary",
		attrs: []string{"Operator token.Token", "Right Expr"},
	})
	child = append(child, StructTemp{
		name:  "Ternary",
		attrs: []string{"Left Expr", "OperatorL token.Token", "Middle Expr", "OperatorR token.Token", "Right Expr"},
	})
	child = append(child, StructTemp{
		name:  "Variable",
		attrs: []string{"Name token.Token"},
	})
	var names []string
	for _, c := range child {
		names = append(names, c.name)
	}
	b.WriteString(defineVisitor("ExprVisitor", names))
	b.WriteString(defineChild(child))
	b.WriteString(defineAccept("ExprVisitor", names))
	return b.String()
}

// Stmt
func buildStmt() string {
	var b strings.Builder
	b.WriteString(defineParent("Stmt"))

	child := []StructTemp{}
	child = append(child, StructTemp{
		name:  "IfStmt",
		attrs: []string{"Condition Expr", "ThenBranch Stmt", "ElseBranch Stmt"},
	})
	child = append(child, StructTemp{
		name:  "PrintStmt",
		attrs: []string{"Expression Expr"},
	})
	child = append(child, StructTemp{
		name:  "Block",
		attrs: []string{"Statements []Stmt"},
	})
	child = append(child, StructTemp{
		name:  "Expression",
		attrs: []string{"Expression Expr"},
	})
	child = append(child, StructTemp{
		name:  "VarStmt",
		attrs: []string{"Name token.Token", "Initializer Expr"},
	})
	var names []string
	for _, c := range child {
		names = append(names, c.name)
	}

	b.WriteString(defineVisitor("StmtVisitor", names))
	b.WriteString(defineChild(child))
	b.WriteString(defineAccept("StmtVisitor", names))

	return b.String()
}

func defineParent(name string) string {
	var b strings.Builder
	b.WriteString("type ")
	b.WriteString(name)
	b.WriteString(" interface{\n")
	b.WriteString("Accept(visitor ")
	b.WriteString(name)
	b.WriteString("Visitor)")
	b.WriteString("any\n}\n")
	return b.String()
}
func defineChild(exprs []StructTemp) string {
	var b strings.Builder
	for _, s := range exprs {
		b.WriteString("type ")
		b.WriteString(s.name)
		b.WriteString(" struct{\n")
		for _, attr := range s.attrs {
			b.WriteString(attr)
			b.WriteString("\n")
		}
		b.WriteString("}\n")
	}
	return b.String()
}
func defineVisitor(vname string, names []string) string {
	var b strings.Builder
	b.WriteString("type ")
	b.WriteString(vname)
	b.WriteString(" interface{\n")
	for _, name := range names {
		b.WriteString("Visit")
		b.WriteString(name)
		b.WriteString("(")
		b.WriteByte(strings.ToLower(name)[0])
		b.WriteString(" *")
		b.WriteString(name)
		b.WriteString(") any\n")
	}
	b.WriteString("}\n")
	return b.String()
}
func defineAccept(visitorName string, names []string) string {
	var b strings.Builder
	for _, name := range names {
		b.WriteString("func(")
		b.WriteByte(strings.ToLower(name)[0])
		b.WriteString(" ")
		b.WriteString(name)
		b.WriteString(")Accept(visitor ")
		b.WriteString(visitorName)
		b.WriteString(")any{\n")
		b.WriteString("return visitor.Visit")
		b.WriteString(name)
		b.WriteString("(&")
		b.WriteByte(strings.ToLower(name)[0])
		b.WriteString(")\n}\n")
	}
	return b.String()
}
