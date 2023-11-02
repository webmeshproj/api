// tsutil-gen generates Typescript helper methods alongside the generated protobuf code.
package main

import (
	"embed"
	"fmt"
	"strings"
	"text/template"

	"google.golang.org/protobuf/proto"

	v1 "github.com/webmeshproj/api/go/v1"
)

// Spec is the spec for the generated Typescript code.
type Spec struct {
	// TypesPath is the path to the generated Typescript types.
	TypesPath string
	// Interfaces is a list of interfaces to generate.
	Interfaces []Interface
}

// Imports returns a map of JS import path to the list of types that
// should be imported from that path.
func (s *Spec) Imports() map[string][]string {
	out := map[string][]string{}
	for _, iface := range s.Interfaces {
		out[iface.JSImport] = append(out[iface.JSImport], iface.Name())
	}
	return out
}

// Interface represents an interface for performing CRUD operations
// on a resource.
type Interface struct {
	Type      proto.Message
	JSImport  string
	QueryType string
}

// Name returns the name of the interface.
func (i *Interface) Name() string {
	return string(proto.MessageName(i.Type).Name())
}

//go:embed templates/*.tmpl
var templateFS embed.FS

var funcMap = template.FuncMap{
	"join": func(a []string, sep string) string {
		return strings.Join(a, sep)
	},
}

func main() {
	t := template.Must(template.New("templates").Funcs(funcMap).ParseFS(templateFS, "templates/*.tmpl"))
	genspec := Spec{
		TypesPath: "../v1",
		Interfaces: []Interface{
			{
				Type:      &v1.MeshNode{},
				JSImport:  "node_pb",
				QueryType: "PEERS",
			},
		},
	}
	var buf strings.Builder
	err := t.ExecuteTemplate(&buf, "ts-rpcdb.ts.tmpl", map[string]any{
		"Spec": &genspec,
	})
	if err != nil {
		panic(err)
	}
	// Replace all tabs with 2 spaces.
	out := strings.ReplaceAll(buf.String(), "\t", "  ")
	fmt.Println(out)
}
