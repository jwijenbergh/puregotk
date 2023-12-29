package types

import (
	"fmt"
	"strings"

	"github.com/jwijenbergh/puregotk/internal/gir/util"
)

type argsTemplate struct {
	// Names are the variables but just the names
	Names []string

	// Types are the variables but just the types
	Types []string

	// Call are the variables as given in a function call
	Call []string

	Full []string
}

type funcArgsTemplate struct {
	// Pure are the arguments as passed directly to PureGo
	// The pure Call is a special case that contains the arguments for a callback call
	Pure argsTemplate

	// API are the arguments as suitable for a Go API
	API argsTemplate
}

func (f *funcArgsTemplate) AddAPI(t string, n string, k Kind, ns string) {
	c := n
	stars := strings.Count(t, "*")
	gobjectNs := "gobject."
	if strings.ToLower(ns) == "gobject" {
		gobjectNs = ""
	}
	switch k {
	case CallbackType:
		c = fmt.Sprintf("purego.NewCallback(%s)", n)
	case ClassesType:
		if stars == 0 {
			c = n
			t = "uintptr"
		} else if stars > 1 {
			c = fmt.Sprintf("%sConvertPtr(%s)", gobjectNs, n)
		} else if stars == 1 {
			c = n + ".GoPointer()"
		}
	case InterfacesType:
		t = strings.TrimPrefix(t, "*")
		if stars == 0 {
			c = n
			t = "uintptr"
		} else if stars > 1 {
			c = fmt.Sprintf("%sConvertPtr(%s)", gobjectNs, n)
		} else if stars == 1 {
			c = n + ".GoPointer()"
		}
	}

	// special case for varargs
	if n == "varArgs" {
		c = n + "..."
	}
	f.API.Names = append(f.API.Names, n)
	f.API.Types = append(f.API.Types, t)
	f.API.Call = append(f.API.Call, c)
	f.API.Full = append(f.API.Full, n+" "+t)
}

func (f *funcArgsTemplate) AddPure(t string, n string, k Kind) {
	n += "p"
	c := n
	stars := strings.Count(t, "*")
	switch k {
	case RecordsType:
		if stars == 0 {
			t = "uintptr"
		}
	case CallbackType:
		t = "uintptr"
	case ClassesType:
		if stars == 0 {
			c = n
			t = "uintptr"
		} else {
			c = fmt.Sprintf("%sNewFromInternalPtr(%s)", strings.TrimPrefix(t, "*"), n)
			t = strings.Repeat("*", stars-1) + "uintptr"
		}
	case InterfacesType:
		if stars == 0 {
			c = n
			t = "uintptr"
		} else {
			c = fmt.Sprintf("%s{Ptr: %s}", t+"Base", n)
			t = strings.Repeat("*", stars-1) + "uintptr"
		}
	}
	f.Pure.Names = append(f.Pure.Names, n)
	f.Pure.Types = append(f.Pure.Types, t)
	f.Pure.Call = append(f.Pure.Call, c)
	f.Pure.Full = append(f.Pure.Full, n+" "+t)
}

func (f *funcArgsTemplate) Add(p Parameter, ins string, ns string, kinds KindMap) {
	// get the lookup namespace
	// as if the interface namespace is non-empty
	// means we can also lookup in the namespace of the interface
	lns := ns
	if ins != "" {
		lns = ins
	}
	goType := p.Translate(lns, kinds)
	kind := kinds.Kind(lns, goType)
	stars := strings.Count(goType, "*")
	goType = util.NormalizeNamespace(ns, goType, true)
	if kind != OtherType && kind != UnknownType {
		if ins != "" && strings.Count(goType, ".") < 1 {
			goType = ins + "." + goType
		}
	}
	if stars > 0 {
		goType = strings.Repeat("*", stars) + strings.ReplaceAll(goType, "*", "")
	}

	// Get a suitable variable name
	varName := p.VarName()

	f.AddAPI(goType, varName, kind, ns)
	f.AddPure(goType, varName, kind)
}

func (f *funcArgsTemplate) AddThrows(ns string) {
	f.API.Call = append(f.API.Call, "&cerr")
	if strings.ToLower(ns) != "glib" {
		f.Pure.Types = append(f.Pure.Types, "**glib.Error")
	} else {
		f.Pure.Types = append(f.Pure.Types, "**Error")
	}
}

type CallbackTemplate struct {
	Doc  string
	Name string
	Args funcArgsTemplate
	Ret  funcRetTemplate
}

type AliasTemplate struct {
	// Name is the name of the alias given to the Go type declaration
	Name string

	// Doc is the documentation of the alias
	Doc string

	// Value is the value for the alias as a Go type
	Value string
}

type RecordField struct {
	// Name is the Go name of the field
	Name string

	// Type is the Go type of the field
	Type string
}

type RecordTemplate struct {
	// Name is the name of the record given to the Go type declaration
	Name string

	// Doc is the documentation of the alias
	Doc string

	// Constructors is the slice of functions that create the class struct
	Constructors []FuncTemplate

	// Fields is the list of record fields
	Fields []RecordField
}

type enumValues struct {
	// Doc is the documentation for the value
	Doc string
	// Name is the name of the enumeration value
	Name string
	// Value is the actual underlying value
	Value int
}

type EnumTemplate struct {
	// Name is the name of the enumeration declared as the Go type for the int
	Name string
	// Doc is the documentation for the enumeration
	Doc string
	// Values are the list of values for the enumeration
	Values []enumValues
}

type funcRetTemplate struct {
	// Raw is the raw value for the underlying purego function
	Raw string
	// Value is the underlying return value as a Go type
	Value string
	// Class indicates whether or not the return value is a class
	Class bool
	// RefSink indicates whether or not we should increase the reference count using obj.RefSink()
	RefSink bool
	// Throws indicates whether or not this function throws
	Throws bool
}

func (fr *funcRetTemplate) Instance() string {
	val := fr.Value + "{}"
	if strings.HasPrefix(fr.Value, "*") {
		return "&" + val[1:]
	}
	return val
}

func (fr *funcRetTemplate) Return() string {
	if fr.Throws {
		if fr.Value == "" {
			return "error"
		}
		return fmt.Sprintf("(%s, error)", fr.Value)
	}
	return fr.Value
}

func (fr *funcRetTemplate) HasReturn() bool {
	return fr.Value != "" || fr.Throws
}

func (fr *funcRetTemplate) Preamble(nglib bool) string {
	preamb := strings.Builder{}
	if fr.Class {
		preamb.WriteString("var cls ")
		preamb.WriteString(fr.Value)
		preamb.WriteString("\n")
	}
	if fr.Throws {
		preamb.WriteString("var cerr *")
		if nglib {
			preamb.WriteString("glib.")
		}
		preamb.WriteString("Error\n")
	}
	return preamb.String()
}

func (fr *funcRetTemplate) Fmt(ngo bool) string {
	if !fr.HasReturn() {
		return ""
	}
	after := strings.Builder{}
	val := "cret"
	if fr.Class {
		if fr.Throws {
			after.WriteString(`
    if cret == 0 {
        return nil, cerr
    }
`)
		} else {
			after.WriteString(`
    if cret == 0 {
        return nil
    }
`)
		}
		if fr.RefSink {
			if ngo {
				after.WriteString("gobject.")
			}
			after.WriteString("IncreaseRef(cret)\n")
		}
		after.WriteString("cls = ")
		after.WriteString(fr.Instance())
		after.WriteString("\n")
		after.WriteString("cls.Ptr = cret\n")
		val = "cls"
	}
	if fr.Throws {
		after.WriteString("if cerr == nil {\n")
		after.WriteString("return ")
		if fr.Value != "" {
			after.WriteString(val)
			after.WriteString(",")
		}
		after.WriteString("nil\n")
		after.WriteString("}\n")
		after.WriteString("return ")
		if fr.Value != "" {
			after.WriteString(val)
			after.WriteString(",")
		}
		after.WriteString("cerr\n")
		return after.String()
	}
	after.WriteString("return ")
	after.WriteString(val)
	return after.String()
}

type FuncTemplate struct {
	// Name is the name of the function declared as the Go function variable and public exposed API
	Name string
	// CName is the raw c name to be passed to purego register
	CName string
	// Doc is the documentation for the function
	Doc string
	// Args are the arguments
	Args funcArgsTemplate
	// Ret is the return argument
	Ret funcRetTemplate
}

type InterfaceFuncTemplate struct {
	Namespace string
	FullName  string
	FuncTemplate
}

type SignalsTemplate struct {
	Doc   string
	Name  string
	CName string
	Args  funcArgsTemplate
	Ret   funcRetTemplate
}

type ClassTemplate struct {
	// Doc is the documentation for the class
	Doc string
	// Name is the name of the class that is given to the Go struct
	Name string
	// Parent is a non-empty string of the embedded parent struct
	Parent string
	// Constructors is the slice of functions that create the class struct
	Constructors []FuncTemplate
	// Receivers is the slice of functions that have value receivers to the struct
	Receivers []FuncTemplate
	// Interfaces are receiver methods that are implemented because it needs to satisfy a certain interface
	Interfaces []InterfaceTemplate
	// Functions are the Go function declarations
	Functions []FuncTemplate
	// Signals are helpers for ConnectX receivers
	Signals []SignalsTemplate
}

type InterfaceTemplate struct {
	Doc  string
	Name string
	// Methods is the methods that this interface defines
	Methods []InterfaceFuncTemplate
}

type TemplateArg struct {
	// PkgName is the name of the package, declared at the top-level
	PkgName string
	// PkgEnv is the name of the package in the load environment variable
	PkgEnv string
	// NeedsInit declares whether or not this file needs an init code to register functions with purego
	NeedsInit bool
	// Imports defines the package imports that we need
	// This does not include purego
	// As the template already includes that if `NeedsInit` is set to true
	Imports []string
	// Aliases are type aliases declared as type ... = ...
	Aliases []AliasTemplate
	// Aliases are structs that are not classes
	Records []RecordTemplate
	// Callbacks are functions that will be converted with purego to uintptr
	Callbacks []CallbackTemplate
	// Enums are enumerations declared as const ... .... = ....
	Enums []EnumTemplate
	// Functions are the Go function declarations
	Functions []FuncTemplate
	// Interfaces is the list of interfaces that this package implements
	Interfaces []InterfaceTemplate
	// Classes are the Go struct with receiver declarations
	Classes []ClassTemplate
}
