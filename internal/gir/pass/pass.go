// package pass implements the first and second pass to go from gir files to go files
// the first pass collects basic type information
// the second pass uses the basic type information to go over the gir files again and convert it to go files
package pass

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/jwijenbergh/puregotk/internal/gir/types"
	"github.com/jwijenbergh/puregotk/internal/gir/util"
)

type Pass struct {
	Parsed []types.Repository
	Types  types.KindMap
}

// New creates a new pass struct by parsing gir files in the string slice
// This pass object will then be used to go over these files multiple times up until we have the full info to convert it to go files
func New(files []string) (*Pass, error) {
	p := Pass{
		Parsed: make([]types.Repository, len(files)),
		Types:  make(types.KindMap),
	}
	for i, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		var r types.Repository
		err = xml.Unmarshal(b, &r)
		if err != nil {
			return nil, err
		}
		p.Parsed[i] = r
	}
	return &p, nil
}

func (p *Pass) collectTypes(r types.Repository) {
	ns := r.Namespaces[0]
	for _, cls := range ns.Classes {
		p.Types.Add(ns.Name, cls.Name, types.ClassesType, cls)
	}
	for _, rec := range ns.Records {
		p.Types.Add(ns.Name, rec.Name, types.RecordsType, rec)
	}
	for _, en := range ns.Enums {
		// TODO: This probably shouldn't be aliastype, but we should make dedicated types
		p.Types.Add(ns.Name, en.Name, types.AliasType, en)
	}
	for _, cb := range ns.Callbacks {
		p.Types.Add(ns.Name, cb.Name, types.CallbackType, cb)
	}
	for _, b := range ns.Bitfields {
		// TODO: This probably shouldn't be aliastype, but we should make dedicated types
		p.Types.Add(ns.Name, b.Name, types.AliasType, b)
	}
	for _, inter := range ns.Interfaces {
		p.Types.Add(ns.Name, inter.Name, types.InterfacesType, inter)
	}
}

// First does a "first pass" meaning it collects basic type information for all the repositories
func (p *Pass) First() {
	for _, r := range p.Parsed {
		p.collectTypes(r)
	}
}

func (p *Pass) writeGo(r types.Repository, gotemp *template.Template, dir string) {
	ns := r.Namespaces[0]

	aliases := make(map[string][]types.AliasTemplate)
	enums := make(map[string][]types.EnumTemplate)
	var files []string
	for _, el := range ns.Bitfields {
		temp := el.Template(ns.Name)
		fn := el.FilenameSafe()
		files = append(files, fn)
		enums[fn] = append(enums[fn], temp)
	}

	for _, el := range ns.Enums {
		temp := el.Template(ns.Name)
		fn := el.FilenameSafe()
		files = append(files, fn)
		enums[fn] = append(enums[fn], temp)
	}

	records := make(map[string][]types.RecordTemplate)
	recordLookup := make(map[string]bool)
	for _, rec := range ns.Records {
		name := util.SnakeToCamel(rec.Name)
		fields := make([]types.RecordField, len(rec.Fields))
		fn := rec.FilenameSafe()
		files = append(files, fn)
		for _, f := range rec.Fields {
			_type := f.Translate(ns.Name, p.Types)
			if _type == "" {
				continue
			}
			// HACK: Handle the specific case where a gint is converted to an int
			// But for structs this needs to be an int32 as purego just gets the pointer to the struct
			// Instead of converting each field separately
			if f.AnyType.Type != nil && f.AnyType.Type.CType == "gint" {
				_type = "int32"
			}

			// HACK: in structs the strings should be uintptr as we convert it ourselves
			if _type == "string" {
				_type = "uintptr"
			}
			fields = append(fields, types.RecordField{
				Name: util.SnakeToCamel(f.Name),
				Type: _type,
			})
		}
		records[fn] = append(records[fn], types.RecordTemplate{
			Name:   name,
			Doc:    rec.Doc.StringSafe(),
			Fields: fields,
		})
		recordLookup[name] = true
	}

	callbacks := make(map[string][]types.CallbackTemplate)
	// set every callback equal to uintptr as well
	for _, cb := range ns.Callbacks {
		fn := cb.FilenameSafe()
		files = append(files, fn)
		cbT := types.CallbackTemplate{
			Doc:  cb.Doc.StringSafe(),
			Name: cb.Name,
			Args: cb.Parameters.Template(ns.Name, "", p.Types, cb.Throws),
			Ret:  cb.ReturnValue.Template(ns.Name, "", p.Types, cb.Throws),
		}
		callbacks[fn] = append(callbacks[fn], cbT)
	}

	interfaces := make(map[string][]types.InterfaceTemplate)
	for _, inter := range ns.Interfaces {
		fn := inter.FilenameSafe()
		files = append(files, fn)
		interfaces[fn] = append(interfaces[fn], types.ConvertInterface(ns.Name, "", inter, nil, p.Types))
	}

	for _, union := range ns.Unions {
		fn := union.FilenameSafe()
		files = append(files, fn)
		name := util.SnakeToCamel(union.Name)
		interT := types.AliasTemplate{
			Doc:  union.Doc.StringSafe(),
			Name: name,
			// structs are not yet supported in CGO
			Value: "uintptr",
		}
		aliases[fn] = append(aliases[fn], interT)
	}

	for _, alias := range ns.Aliases {
		fn := alias.FilenameSafe()
		files = append(files, fn)
		typeName := alias.Type.Template(ns.Name, p.Types)
		if typeName == "" {
			typeName = "uintptr"
		}
		name := util.SnakeToCamel(alias.Name)
		aliasT := types.AliasTemplate{
			Doc:  alias.Doc.StringSafe(),
			Name: name,
			// structs are not yet supported in CGO
			Value: typeName,
		}
		aliases[fn] = append(aliases[fn], aliasT)
	}

	functions := make(map[string][]types.FuncTemplate)
	for _, f := range ns.Functions {
		name := util.SnakeToCamel(f.Name)
		if p.Types.Kind(ns.Name, name) != types.UnknownType {
			name = "New" + name
		}
		fn := f.FilenameSafe()
		files = append(files, fn)
		functions[fn] = append(functions[fn], types.FuncTemplate{
			Name:  name,
			CName: f.CIdentifier,
			Doc:   f.Doc.StringSafe(),
			Args:  f.Parameters.Template(ns.Name, "", p.Types, f.Throws),
			Ret:   f.ReturnValue.Template(ns.Name, "", p.Types, f.Throws),
		})
	}

	classes := make(map[string][]types.ClassTemplate)
	for _, cls := range ns.Classes {
		implemented := make(map[string]bool)
		constructors := make([]types.FuncTemplate, len(cls.Constructors))
		fn := cls.FilenameSafe()
		files = append(files, fn)

		for i, c := range cls.Constructors {
			name := util.SnakeToCamel(c.Name)
			constructors[i] = types.FuncTemplate{
				Name:  name,
				CName: c.CIdentifier,
				Doc:   c.Doc.StringSafe(),
				Args:  c.Parameters.Template(ns.Name, "", p.Types, c.Throws),
				Ret:   c.ReturnValue.Template(ns.Name, "", p.Types, c.Throws),
			}
		}
		signals := make([]types.SignalsTemplate, len(cls.Signals))
		for i, s := range cls.Signals {
			signals[i] = types.SignalsTemplate{
				Doc:   s.Doc.StringSafe(),
				Name:  util.DashToCamel(s.Name),
				CName: s.Name,
				Args:  s.Parameters.Template(ns.Name, "", p.Types, false),
				Ret:   s.ReturnValue.Template(ns.Name, "", p.Types, false),
			}
		}
		receivers := make([]types.FuncTemplate, len(cls.Methods))
		for i, f := range cls.Methods {
			name := util.SnakeToCamel(f.Name)
			implemented[name] = true
			receivers[i] = types.FuncTemplate{
				Doc:   f.Doc.StringSafe(),
				Name:  name,
				CName: f.CIdentifier,
				Args:  f.Parameters.Template(ns.Name, "", p.Types, f.Throws),
				Ret:   f.ReturnValue.Template(ns.Name, "", p.Types, f.Throws),
			}
		}
		var interfaces []types.InterfaceTemplate
		for _, impl := range cls.Implements {
			interfaces = append(interfaces, types.GetInterfaceFuncs(ns.Name, impl.Name, implemented, p.Types))
		}
		classes[fn] = append(classes[fn], types.ClassTemplate{
			Doc:          cls.Doc.StringSafe(),
			Name:         cls.Name,
			Parent:       util.NormalizeNamespace(ns.Name, cls.Parent, true),
			Constructors: constructors,
			Receivers:    receivers,
			Interfaces:   interfaces,
			Signals:      signals,
		})
	}

	pkgName := strings.ToLower(ns.Name)
	for _, fn := range files {
		var includesSlice []string
		//includesSlice := make([]string, len(includes[fn]))
		//i := 0
		//for k, v := range includes[fn] {
		//	if v {
		//		includesSlice[i] = k
		//	}
		//	i++
		//}
		methods := 0
		for _, i := range interfaces[fn] {
			for range i.Methods {
				methods += 1
			}
		}
		for _, i := range classes[fn] {
			for range i.Constructors {
				methods += 1
			}
			for range i.Receivers {
				methods += 1
			}
		}
		// we do not need to add the length of interfaces in here
		// as they should only be loaded when there are classes
		needsInit := (len(functions[fn]) + methods) > 0

		args := types.TemplateArg{
			PkgName:    pkgName,
			PkgEnv:     strings.ToUpper(pkgName),
			NeedsInit:  needsInit,
			Imports:    includesSlice,
			Aliases:    aliases[fn],
			Callbacks:  callbacks[fn],
			Records:    records[fn],
			Enums:      enums[fn],
			Functions:  functions[fn],
			Interfaces: interfaces[fn],
			Classes:    classes[fn],
		}

		os.MkdirAll(fmt.Sprintf(dir+"/%s", pkgName), 0o755)

		f, err := os.Create(fmt.Sprintf(dir+"/%s/%s", pkgName, fn))
		if err != nil {
			panic(err)
		}
		err = gotemp.Execute(f, args)
		if err != nil {
			panic(err)
		}

	}
}

func (p *Pass) Second(dir string, gotemp *template.Template) {
	for _, r := range p.Parsed {
		p.writeGo(r, gotemp, dir)
	}
}
