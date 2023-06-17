package types

import (
	"strings"

	"github.com/jwijenbergh/puregotk/internal/gir/util"
)

func ConvertInterface(currns string, ins string, inter Interface, implemented map[string]bool, kinds KindMap) InterfaceTemplate {
	var methods []InterfaceFuncTemplate

	for _, m := range inter.Methods {
		name := util.SnakeToCamel(m.Name)
		if implemented != nil {
			v, ok := implemented[name]
			if v && ok {
				continue
			}
		}
		var newns string
		if ins != "" {
			newns = ins + "."
		}
		methods = append(methods, InterfaceFuncTemplate{
			Namespace: newns,
			FullName:  util.SnakeToCamel(m.CIdentifier),
			FuncTemplate: FuncTemplate{
				Doc:   m.Doc.StringSafe(),
				CName: m.CIdentifier,
				Name:  name,
				Args:  m.Parameters.Template(currns, ins, kinds),
				Ret:   m.ReturnValue.Template(currns, ins, kinds),
			},
		})
	}
	name := util.SnakeToCamel(inter.Name)
	return InterfaceTemplate{
		Name:    name,
		Doc:     inter.Doc.StringSafe(),
		Methods: methods,
	}
}

func GetInterfaceFuncs(ns string, name string, implemented map[string]bool, kinds KindMap) InterfaceTemplate {
	inter := kinds.MustInterface(ns, name)
	parts := strings.Split(name, ".")
	var ins string
	if len(parts) > 1 {
		ins = strings.ToLower(parts[0])
	}
	return ConvertInterface(ns, ins, inter, implemented, kinds)
}
