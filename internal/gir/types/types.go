// package types implements types as found in gir XML files
// NOTE: This was copied largely from (the type definitions) from https://github.com/diamondburned/gotk4
package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"github.com/jwijenbergh/puregotk/internal/gir/util"
)

// convList maps the given GIR primitive type to a Go builtin type.
// See https://github.com/diamondburned/gotk4/blob/fd960d20b525a07580938d10a214336bafb47d12/gir/girgen/types/types.go#LL483C1-L512C2
var convList = map[string]string{
	"none":     "",
	"gboolean": "bool",
	"gfloat":   "float32",
	"gdouble":  "float64",
	"gint":     "int",
	"gssize":   "int",
	"gint8":    "int8",
	"gint16":   "int16",
	"gshort":   "int16",
	"gint32":   "int32",
	"glong":    "int32",
	"int32":    "int32",
	"gint64":   "int64",
	"guint":    "uint",
	"gsize":    "uint",
	"guchar":   "byte",
	"gchar":    "byte",
	"guint8":   "byte", // some weird cases
	"guint16":  "uint16",
	"gushort":  "uint16",
	"guint32":  "uint32",
	"gulong":   "uint32",
	"gunichar": "uint32",
	"guint64":  "uint64",
	"guintptr": "uintptr",
	"utf8":     "string",
	"filename": "string",
	"GType":    "types.GType",

	// these are probably not correct but needed to compile
	"_Value__data__union": "uint64",
	"long double":         "float64",
	"va_list":             "[]interface{}",
	"gpointer":            "uintptr",
	"const char*":         "string",
	"const char**":        "[]string",
	"char**":              "[]string",
	// not exported
	"HarfBuzz.feature_t": "uintptr",
	"HarfBuzz.font_t":    "uintptr",
	//"GLib.Quark": "byte",
	//"Allocation": "uintptr",
}

// Repository represents a GObject Introspection Repository, which contains the
// includes, C includes and namespaces of a single gir file.
type Repository struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 repository"`

	Version             string `xml:"version,attr"`
	CIdentifierPrefixes string `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`

	Includes   []Include   `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	CIncludes  []CInclude  `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Packages   []Package   `xml:"http://www.gtk.org/introspection/core/1.0 package"`
	Namespaces []Namespace `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`
}

// https://gitlab.gnome.org/GNOME/gobject-introspection/-/blob/master/docs/gir-1.2.rnc

type Alias struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 alias"`

	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	InfoAttrs
	InfoElements

	Type Type `xml:"http://www.gtk.org/introspection/core/1.0 type"`
}

type AnyType struct {
	// Possible variants.
	Type  *Type  `xml:"http://www.gtk.org/introspection/core/1.0 type"`
	Array *Array `xml:"http://www.gtk.org/introspection/core/1.0 array"`
}

func (a *AnyType) Translate(ns string, kinds KindMap) string {
	if a == nil {
		panic("unknown type found")
	}
	t := a.Type
	if t == nil {
		if a.Array == nil {
			// no type information available
			return ""
		}
		return a.Array.Template(ns, kinds)
	}
	return a.Type.Template(ns, kinds, false)
}

type Annotation struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Array struct {
	XMLName        xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 array"`
	Name           string   `xml:"name,attr"`
	CType          string   `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Length         *int     `xml:"length,attr"` // ix of .Parameters
	ZeroTerminated *bool    `xml:"zero-terminated,attr"`
	FixedSize      int      `xml:"fixed-size,attr"`
	Introspectable bool     `xml:"introspectable,attr"`
	// Type is the array's inner type.
	Type *Type `xml:"http://www.gtk.org/introspection/core/1.0 type"`
}

func (a *Array) Size() int {
	// there is a Length property
	// but it seems that GTK doesn't really use this
	// see e.g. gtk_css_provider_load_from_data
	//if a.Length != nil {
	//	return *a.Length
	//}
	return a.FixedSize
}

func ReplaceSize(str string, n int) string {
	trimmed := strings.TrimPrefix(str, "*")
	if trimmed == str {
		trimmed = strings.TrimPrefix(str, "[]")
	}
	return fmt.Sprintf("[%d]%s", n, trimmed)
}

func (a *Array) Template(ns string, kinds KindMap) string {
	// Does the c type exist in our little hacky conversion rule?
	if v, ok := convList[a.CType]; ok {
		return v
	}
	// default case of an array is just a pointer
	ret := "uintptr"

	if a.Type != nil {
		t := a.Type.Template(ns, kinds, true)
		if strings.Count(t, "*") > 0 {
			return "uintptr"
		}
		ret = "[]" + t
	}

	// Check if there is a fixed size
	n := a.Size()
	if n > 0 {
		return ReplaceSize(ret, n)
	}
	return ret
}

// IsZeroTerminated returns true if the Array is zero-terminated. It accounts
// for edge cases of the structure.
func (a Array) IsZeroTerminated() bool {
	return a.Name == "" && (a.ZeroTerminated == nil || *a.ZeroTerminated)
}

type Bitfield struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 bitfield"`

	Name         string `xml:"name,attr"` // Go case
	CType        string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GLibTypeName string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType  string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`

	Members   []Member   `xml:"http://www.gtk.org/introspection/core/1.0 member"`
	Functions []Function `xml:"http://www.gtk.org/introspection/core/1.0 function"`

	InfoAttrs
	InfoElements
}

func (b *Bitfield) Template(ns string) EnumTemplate {
	els := make([]enumValues, len(b.Members))
	for i, m := range b.Members {
		v, err := strconv.Atoi(m.Value)
		if err != nil {
			panic(err)
		}
		// + Value needed to get rid of duplicates
		els[i] = enumValues{
			Doc:   m.Doc.StringSafe(),
			Name:  util.SnakeToCamel(util.RemoveSnakePrefix(strings.ToLower(m.CIdentifier), ns)) + "Value",
			Value: v,
		}
	}
	return EnumTemplate{
		Name:   util.SnakeToCamel(b.Name),
		Doc:    b.Doc.StringSafe(),
		Values: els,
	}
}

type Signal struct {
	XMLName   xml.Name   `xml:"http://www.gtk.org/introspection/glib/1.0 signal"`
	Name      string     `xml:"name,attr"`
	Detailed  bool       `xml:"detailed,attr"`
	When      SignalWhen `xml:"when,attr"`
	Action    bool       `xml:"action,attr"`
	NoHooks   bool       `xml:"no-hooks,attr"`
	NoRecurse bool       `xml:"no-recurse,attr"`
	InfoElements
	Parameters  *Parameters  `xml:"http://www.gtk.org/introspection/core/1.0 parameters"`
	ReturnValue *ReturnValue `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
}

type SignalWhen string

const (
	SignalWhenFirst   = "first"
	SignalWhenLast    = "last"
	SignalWhenCleanup = "cleanup"
)

type SourcePosition struct {
	XMLName  xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 source-position"`
	Filename string   `xml:"filename,attr"`
	Line     int      `xml:"line,attr"`
	Column   int      `xml:"column,attr"`
}

type TransferOwnership struct {
	TransferOwnership string `xml:"transfer-ownership,attr"`
}

type Type struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 type"`

	Name           string `xml:"name,attr"`
	CType          string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Introspectable *bool  `xml:"introspectable,attr"`

	DocElements
	// Types is the type's inner types.
	Types []Type `xml:"http://www.gtk.org/introspection/core/1.0 type"`
}

func (typ *Type) IsIntrospectable() bool {
	return typ.Introspectable == nil || *typ.Introspectable
}

func (t *Type) Template(ns string, kinds KindMap, array bool) string {
	if v, ok := convList[t.Name]; ok {
		return v
	}

	_type := util.NormalizeNamespace(ns, t.Name, true)

	kind := kinds.Kind(ns, _type)
	// find the total pointer count
	count := strings.Count(t.CType, "*")
	w := strings.Trim(t.CType, "*")
	if w == "gpointer" {
		count += 1
	}
	// set the same pointer types on the name
	if count > 0 && kind != CallbackType {
		_type = strings.Repeat("*", count) + _type
	} else if kind == RecordsType {
		if array {
			return _type
		}
		// if it's not a pointer
		// then purego doesn't support it (struct by value is not supported
		// so return uintptr
		_type = "uintptr"
	}
	return _type
}

type Union struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 union"`

	Name          string `xml:"name,attr"` // Go case
	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	GLibTypeName  string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType   string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`

	InfoAttrs
	InfoElements

	Fields       []Field       `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Constructors []Constructor `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Methods      []Method      `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Functions    []Function    `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Records      []Record      `xml:"http://www.gtk.org/introspection/core/1.0 record"`
}

type VarArgs struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 varargs"`
}

type VirtualMethod struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 virtual-method"`

	Invoker string `xml:"invoker,attr"`
	CallableAttrs
}

type Boxed struct{}

type CInclude struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Name    string   `xml:"name,attr"`
}

type CallableAttrs struct {
	Name        string       `xml:"name,attr"`
	CIdentifier string       `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	ShadowedBy  string       `xml:"shadowed-by,attr"`
	Shadows     string       `xml:"shadows,attr"`
	Throws      bool         `xml:"throws,attr"`
	MovedTo     string       `xml:"moved-to,attr"`
	Parameters  *Parameters  `xml:"http://www.gtk.org/introspection/core/1.0 parameters"`
	ReturnValue *ReturnValue `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
	InfoAttrs
	InfoElements
}

type Callback struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	CallableAttrs
}

type Class struct {
	XMLName  xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 class"`
	Name     string   `xml:"name,attr"`
	Parent   string   `xml:"parent,attr"`
	Abstract bool     `xml:"abstract,attr"`

	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`

	GLibTypeName   string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType    string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GLibTypeStruct string `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`

	InfoAttrs
	InfoElements

	Functions      []Function      `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Implements     []Implements    `xml:"http://www.gtk.org/introspection/core/1.0 implements"`
	Constructors   []Constructor   `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Methods        []Method        `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	VirtualMethods []VirtualMethod `xml:"http://www.gtk.org/introspection/core/1.0 virtual-method"`
	Fields         []Field         `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Signals        []Signal        `xml:"http://www.gtk.org/introspection/glib/1.0 signal"`
}

type Constant struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 constant"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
	CType   string   `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	Type Type `xml:"http://www.gtk.org/introspection/core/1.0 type"`

	InfoAttrs
	InfoElements
}

func (c *Constant) Template(ns string, kinds KindMap) ConstantTemplate {
	t := c.Type.Template(ns, kinds, false)
	v := c.Value

	switch t {
	case "string":
		v = fmt.Sprintf(`"%s"`, v)
	}

	return ConstantTemplate{
		Name:  c.Name,
		Doc:   c.Doc.StringSafe(),
		Type:  t,
		Value: v,
	}
}

type Constructor struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	CallableAttrs
}

type Doc struct {
	XMLName  xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 doc"`
	Filename string   `xml:"filename,attr"`
	String   string   `xml:",innerxml"`
	Line     int      `xml:"line,attr"`
}

func (d *Doc) StringSafe() string {
	if d == nil {
		return ""
	}
	lines := strings.Split(d.String, "\n")
	for i, l := range lines {
		lines[i] = "// " + l
	}
	return strings.Join(lines, "\n")
}

type DocDeprecated struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 doc-deprecated"`
	String  string   `xml:",innerxml"`
}

type DocElements struct {
	Doc            *Doc
	DocDeprecated  *DocDeprecated
	SourcePosition *SourcePosition
}

func (d *DocElements) FilenameSafe() string {
	if d == nil || (d.SourcePosition == nil && d.Doc == nil) {
		return "main.go"
	}
	if d.SourcePosition != nil {
		return util.TranslateFilename(d.SourcePosition.Filename)
	}
	return util.TranslateFilename(d.Doc.Filename)
}

type Enum struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`

	Name            string `xml:"name,attr"` // Go case
	CType           string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GLibTypeName    string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType     string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GLibErrorDomain string `xml:"http://www.gtk.org/introspection/glib/1.0 error-domain,attr"`

	Members   []Member   `xml:"http://www.gtk.org/introspection/core/1.0 member"`
	Functions []Function `xml:"http://www.gtk.org/introspection/core/1.0 function"`

	InfoAttrs
	InfoElements
}

func (e *Enum) Template(ns string) EnumTemplate {
	els := make([]enumValues, len(e.Members))

	for i, m := range e.Members {
		v, err := strconv.Atoi(m.Value)
		if err != nil {
			panic(err)
		}
		id := strings.ToLower(m.CIdentifier)
		sID := util.RemoveSnakePrefix(id, ns)
		els[i] = enumValues{
			Doc: m.Doc.StringSafe(),
			// + Value needed to get rid of duplicates
			Name:  util.SnakeToCamel(sID) + "Value",
			Value: v,
		}
	}
	return EnumTemplate{
		Name:   util.SnakeToCamel(e.Name),
		Doc:    e.Doc.StringSafe(),
		Values: els,
	}
}

type Field struct {
	XMLName  xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Name     string   `xml:"name,attr"`
	Bits     int      `xml:"bits,attr"`
	Private  bool     `xml:"private,attr"`
	Writable bool     `xml:"writable,attr"` // default false
	Readable *bool    `xml:"readable,attr"` // default true
	AnyType
	Callback *Callback
	Doc      *Doc
}

// IsReadable returns true if the field is readable.
func (f Field) IsReadable() bool {
	return f.Readable == nil || *f.Readable
}

type Function struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	CallableAttrs
}

type Implements struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 implements"`
	Name    string   `xml:"name,attr"`
}

type Include struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	Name    string   `xml:"name,attr"`
	Version string   `xml:"version,attr"`
}

type InfoAttrs struct {
	Introspectable    *bool  `xml:"introspectable,attr"` // default true
	Deprecated        bool   `xml:"deprecated,attr"`
	DeprecatedVersion string `xml:"deprecated-version,attr"`
	Version           string `xml:"version,attr"`
	Stability         string `xml:"stability,attr"`
}

// IsIntrospectable returns true if the InfoAttrs indicates that the type is
// introspectable.
func (inf InfoAttrs) IsIntrospectable() bool {
	return inf.Introspectable == nil || *inf.Introspectable
}

type InfoElements struct {
	DocElements
	Annotations []Annotation `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
}

type InstanceParameter struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 instance-parameter"`
	ParameterAttrs
}

type Interface struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 interface"`
	Name    string   `xml:"name,attr"`

	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`

	GLibTypeName   string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType    string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GLibTypeStruct string `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`

	Functions      []Function      `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Methods        []Method        `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	VirtualMethods []VirtualMethod `xml:"http://www.gtk.org/introspection/core/1.0 virtual-method"`
	Prerequisites  []Prerequisite  `xml:"http://www.gtk.org/introspection/core/1.0 prerequisite"`
	Signals        []Signal        `xml:"http://www.gtk.org/introspection/glib/1.0 signal"`

	InfoAttrs
	InfoElements
}

type Member struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 member"`

	Names       []xml.Attr `xml:"name,attr"`
	Value       string     `xml:"value,attr"`
	CIdentifier string     `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	GLibNick    string     `xml:"http://www.gtk.org/introspection/glib/1.0 nick,attr"`

	InfoAttrs
	InfoElements
}

func (m Member) Name() string {
	return m.nameAttr(xml.Name{Local: "name"})
}

func (m Member) GLibName() string {
	return m.nameAttr(xml.Name{Space: "http://www.gtk.org/introspection/glib/1.0", Local: "name"})
}

func (m Member) nameAttr(name xml.Name) string {
	for _, attr := range m.Names {
		if attr.Name == name {
			return attr.Value
		}
	}
	return ""
}

type Method struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	CallableAttrs
}

type Namespace struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`

	Name                string `xml:"name,attr"`
	Version             string `xml:"version,attr"`
	CIdentifierPrefixes string `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`
	Prefix              string `xml:"http://www.gtk.org/introspection/c/1.0 prefix,attr"`
	SharedLibrary       string `xml:"shared-library,attr"`

	Aliases     []Alias      `xml:"http://www.gtk.org/introspection/core/1.0 alias"`
	Classes     []Class      `xml:"http://www.gtk.org/introspection/core/1.0 class"`
	Interfaces  []Interface  `xml:"http://www.gtk.org/introspection/core/1.0 interface"`
	Records     []Record     `xml:"http://www.gtk.org/introspection/core/1.0 record"`
	Enums       []Enum       `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`
	Functions   []Function   `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Unions      []Union      `xml:"http://www.gtk.org/introspection/core/1.0 union"`
	Bitfields   []Bitfield   `xml:"http://www.gtk.org/introspection/core/1.0 bitfield"`
	Callbacks   []Callback   `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	Constants   []Constant   `xml:"http://www.gtk.org/introspection/core/1.0 constant"`
	Annotations []Annotation `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
	Boxeds      []Boxed      `xml:"http://www.gtk.org/introspection/core/1.0 boxed"`
}

type Package struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 package"`
	Name    string   `xml:"name,attr"`
}

type Parameter struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 parameter"`
	ParameterAttrs
}

func (p *Parameter) Translate(ns string, kinds KindMap) string {
	if p.Name == "..." {
		return "...interface{}"
	}
	return p.AnyType.Translate(ns, kinds)
}

func (p *Parameter) VarName() string {
	snaked := util.SnakeToCamel(p.Name)
	if snaked == "..." {
		return "varArgs"
	}
	// avoid clash between same name of type and var
	// could also do first letter to lower?
	return snaked + "Var"
}

func (p *Parameters) Template(ns string, ifacens string, kinds KindMap, throws bool) funcArgsTemplate {
	if p == nil {
		return funcArgsTemplate{}
	}
	params := p.Parameters
	if len(params) == 0 {
		return funcArgsTemplate{}
	}
	args := funcArgsTemplate{}
	for _, par := range params {
		args.Add(par, ifacens, ns, kinds)
	}
	if throws {
		args.AddThrows(ns)
	}
	return args
}

type ParameterAttrs struct {
	Name            string `xml:"name,attr"`
	Direction       string `xml:"direction,attr"`
	Scope           string `xml:"scope,attr"`
	Closure         *int   `xml:"closure,attr"`
	Destroy         *int   `xml:"destroy,attr"`
	CallerAllocates bool   `xml:"caller-allocates,attr"`
	Skip            bool   `xml:"skip,attr"`
	Optional        bool   `xml:"optional,attr"`
	Nullable        bool   `xml:"nullable,attr"`

	TransferOwnership
	AnyType
	Doc *Doc
}

type Parameters struct {
	XMLName           xml.Name           `xml:"http://www.gtk.org/introspection/core/1.0 parameters"`
	InstanceParameter *InstanceParameter `xml:"http://www.gtk.org/introspection/core/1.0 instance-parameter"`
	Parameters        []Parameter        `xml:"http://www.gtk.org/introspection/core/1.0 parameter"`
}

type Prerequisite struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 prerequisite"`
	Name    string   `xml:"name,attr"`
}

type Property struct{}

type Record struct {
	XMLName              xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 record"`
	Name                 string   `xml:"name,attr"`
	CType                string   `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GLibTypeName         string   `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType          string   `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	CSymbolPrefix        string   `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	GLibIsGTypeStructFor string   `xml:"http://www.gtk.org/introspection/glib/1.0 is-gtype-struct-for,attr"`
	Disguised            bool     `xml:"disguised,attr"`
	Foreign              bool     `xml:"foreign,attr"`

	Fields       []Field       `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Functions    []Function    `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Unions       []Union       `xml:"http://www.gtk.org/introspection/core/1.0 union"`
	Methods      []Method      `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Constructors []Constructor `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Properties   []Property    `xml:"http://www.gtk.org/introspection/core/1.0 property"`

	InfoAttrs
	InfoElements
}

type ReturnValue struct {
	XMLName        xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
	Scope          string   `xml:"scope,attr"`
	Closure        *int     `xml:"closure,attr"`
	Destroy        *int     `xml:"destroy,attr"`
	Introspectable bool     `xml:"introspectable,attr"`
	Nullable       bool     `xml:"nullable,attr"`
	Skip           bool     `xml:"skip,attr"`
	AllowNone      bool     `xml:"allow-none,attr"`
	TransferOwnership
	DocElements
	AnyType
}

func (r *ReturnValue) Template(ns string, ins string, kinds KindMap, throws bool) funcRetTemplate {
	val := r.AnyType.Translate(ns, kinds)
	raw := val
	class := false
	lns := ns
	if ins != "" {
		lns = ins
	}
	kind := kinds.Kind(lns, raw)
	stars := strings.Count(val, "*")

	if kind != OtherType && kind != UnknownType && kind != SliceType {
		val = util.AddNamespace(val, ins)
	}
	if stars > 0 {
		val = util.StarsInFront(strings.ReplaceAll(val, "*", ""), stars)
	}
	switch kind {
	case ClassesType:
		raw = "uintptr"
		class = true
		// TODO: why does this happen?
		if strings.HasPrefix(val, "[]") {
			class = false
			val = "uintptr"
		}
	case InterfacesType:
		raw = "uintptr"
		val += "Base"
		class = true
	// callback returns should always be uintptr
	// I needed this for glib.LogSetDefaultHandler
	// Otherwise I got 'panic: reflect.MakeFunc: value of type *glib.LogFunc is not assignable to type glib.LogFunc'
	case CallbackType:
		raw = "uintptr"
		val = "uintptr"
	}
	return funcRetTemplate{
		Raw:     raw,
		Value:   val,
		Class:   class,
		RefSink: r.TransferOwnership.TransferOwnership == "none",
		Throws:  throws,
	}
}
