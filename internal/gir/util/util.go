// package util implements some utility functions for parsing/converting gir files
// TODO: Maybe some of this can more easily be done with regexes?
//
//	But using regexes introduces 2 problems :^)
package util

import (
	"path/filepath"
	"strings"
)

var (
	// Variable names that should not be dereferenced when using ConvertPtr() in handlePtr mode
	// TODO: This was mostly discovered via trial and error, and might point towards issues in
	// the GIR files
	specialConvertPtrVars = []string{
		"ModelVar",
		"TreeModelVar",
		"OutChildVar",
		"ChildVar",
	}
)

// delimToCamel to camel converts a string with parts separated by `delim` to CamelCase
func delimToCamel(s string, delim string) string {
	var sb strings.Builder
	words := strings.Split(s, delim)
	for _, w := range words {
		sb.WriteString(strings.Title(w))
	}
	return sb.String()
}

// StarsInFront adds pointer characters (*, stars) in front of the type
// if there is a slice in front
// we need to add the slice and then afterwards the stars
// e.g. [2]foo becomes [2]*foo with n=1
func StarsInFront(str string, n int) string {
	b := strings.Index(str, "[")
	e := strings.Index(str, "]")
	stars := strings.Repeat("*", n)
	if b == 0 && e != -1 {
		return str[b:e+1] + stars + str[e+1:]
	}
	return stars + str
}

// SnakeToCamel converts hello_world to HelloWorld
func SnakeToCamel(s string) string {
	return delimToCamel(s, "_")
}

// DashToCamel converts hello-world to HelloWorld
func DashToCamel(s string) string {
	return delimToCamel(s, "-")
}

// RemoveSnakePrefix removes `prefix` from string `s` if that prefix ise separated with a _
// it removes lowercase or all u
func RemoveSnakePrefix(s string, prefix string) string {
	parts := strings.Split(s, "_")
	if len(parts) <= 1 {
		return s
	}
	if strings.EqualFold(parts[0], prefix) {
		parts = parts[1:]
	}
	return strings.Join(parts, "_")
}

// ReplaceExtension replaces an extension from filename with ext
// the extension is found by splitting on "." and taking the last part
func ReplaceExtension(filename string, ext string) string {
	splt := strings.Split(filename, ".")
	if len(splt) == 1 {
		return filename
	}
	splt[len(splt)-1] = ext
	return strings.Join(splt, ".")
}

func PrefixValue(val, prefix string) string {
	// if it's a slice, it has to come first
	b := strings.Index(val, "[")
	e := strings.Index(val, "]")
	if b == 0 && e != -1 {
		return val[b:e+1] + prefix + val[e+1:]
	}
	return prefix + val
}

func AddNamespace(val, ns string) string {
	if ns == "" || strings.Count(val, ".") >= 1 {
		return val
	}
	return PrefixValue(val, ns+".")
}

// NormalizeNamespace converts a type to one that always includes a lowercase namespace
// if no namespace is found, it adds `ns`, unless if strip is True then namespaces always equaling `ns` will be removed
func NormalizeNamespace(ns string, gotype string, strip bool) string {
	if ns == "" {
		return ""
	}
	gotype = strings.Trim(gotype, "*")
	splt := strings.Split(gotype, ".")
	if len(splt) <= 1 {
		splt = append([]string{ns}, splt...)
	}
	splt[0] = strings.ToLower(splt[0])
	if strip && splt[0] == strings.ToLower(ns) {
		splt = splt[1:]
	}
	return strings.Join(splt, ".")
}

// TranslateFilename translates a file path by renaming the file to a go suitable file
func TranslateFilename(filename string) string {
	if filename == "" {
		return "main.go"
	}

	b := filepath.Base(filename)
	return ReplaceExtension(b, "go")
}

func ConvertArgs(a []string) string {
	return strings.Join(a, ", ")
}

func ConvertArgsComma(a []string) string {
	if len(a) == 0 {
		return ""
	}
	return ", " + strings.Join(a, ", ")
}

func convertCallbackArgs(a []string, prependComma, skipEmpty, skipErr, handlePtr bool) string {
	var validArgs []string
	for _, arg := range a {
		if skipEmpty && arg == "" {
			continue
		}
		if skipErr && arg == "&cerr" {
			continue
		}

		if strings.Contains(arg, "{Ptr:") {
			if !handlePtr {
				// For ConvertCallbackArgs: remove * prefix and add &
				arg = strings.TrimPrefix(arg, "*")
			}
			validArgs = append(validArgs, "&"+arg)
		} else if strings.Contains(arg, "ConvertPtr(") && handlePtr {
			isSpecialVar := false
			for _, specialVar := range specialConvertPtrVars {
				if strings.Contains(arg, specialVar) {
					isSpecialVar = true

					break
				}
			}

			if isSpecialVar {
				validArgs = append(validArgs, arg)
			} else {
				validArgs = append(validArgs, "*"+arg)
			}
		} else {
			validArgs = append(validArgs, arg)
		}
	}

	if len(validArgs) == 0 {
		return ""
	}

	result := strings.Join(validArgs, ", ")
	if prependComma {
		return ", " + result
	}

	return result
}

func ConvertCallbackArgs(a []string) string {
	return convertCallbackArgs(a, false, false, false, false)
}

func ConvertArgsCommaDeref(a []string) string {
	return convertCallbackArgs(a, true, true, false, true)
}

func ConvertArgsDeref(a []string) string {
	return convertCallbackArgs(a, false, true, false, true)
}

func ConvertCallbackArgsNoErr(a []string) string {
	return convertCallbackArgs(a, false, true, true, true)
}

// ConstructorName returns a Go friendly constructor name given the raw constructor name `name` and the class/record name `outer`
func ConstructorName(name string, outer string) string {
	cname := SnakeToCamel(name)
	// construct the final constructor name
	// for example if we have gtk_builder
	// gtk_builder_new_from_file
	// cname will be NewFromFile
	// we convert it to NewBuilderFromFile
	if strings.HasPrefix(cname, "New") {
		return "New" + outer + cname[3:]
	}
	// the default is just a concatenation if the constructor doesn't start with New
	return outer + cname
}

// PropertyScalarSet returns the code for setting a scalar property value
func PropertyScalarSet(notGObject bool, gvalueType, setMethod string) string {
	prefix := ""
	if notGObject {
		prefix = "gobject."
	}
	return prefix + gvalueType + ")\n\tv." + setMethod + "(value"
}

// PropertyScalarGet returns the code for getting a scalar property value
func PropertyScalarGet(getMethod string) string {
	return "v." + getMethod + "()"
}

// PropertyVectorSet returns the code for setting a vector property value
func PropertyVectorSet(notGLib bool) string {
	prefix := ""
	if notGLib {
		prefix = "glib."
	}

	return prefix + `StrvGetType())

	cStrBytes := make([][]byte, len(value))
	cStrings := make([]uintptr, len(value)+1)
	for i, s := range value {
		cStrBytes[i] = make([]byte, len(s)+1)
		copy(cStrBytes[i], s)
		cStrBytes[i][len(s)] = 0
		cStrings[i] = uintptr(unsafe.Pointer(&cStrBytes[i][0]))
	}
	cStrings[len(value)] = 0

	v.SetBoxed(uintptr(unsafe.Pointer(&cStrings[0])))`
}

// PropertyVectorGet returns the code for getting a vector property value
func PropertyVectorGet() string {
	return `cStrvPtr := v.GetBoxed()
	if cStrvPtr == 0 {
		return nil
	}

	var result []string
	ptr := cStrvPtr
	for {
		strPtr := *(*uintptr)(unsafe.Pointer(ptr))
		if strPtr == 0 {
			break
		}

		var length int
		for i := 0; ; i++ {
			if *(*byte)(unsafe.Pointer(strPtr + uintptr(i))) == 0 {
				length = i
				break
			}
		}

		bytes := unsafe.Slice((*byte)(unsafe.Pointer(strPtr)), length)
		result = append(result, string(bytes))
		ptr += unsafe.Sizeof(uintptr(0))
	}
	return result`
}
