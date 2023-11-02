/*
Copyright 2023 Avi Zimmerman <avi.zimmerman@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// tsutil-gen generates Typescript helper methods alongside the generated protobuf code.
package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
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
	outPath := flag.String("out", "-", "output file")
	flag.Parse()
	outFile := os.Stdout
	if *outPath != "-" {
		var err error
		outFile, err = os.Create(*outPath)
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
	}
	t := template.Must(template.New("templates").Funcs(funcMap).ParseFS(templateFS, "templates/*.tmpl"))
	genspec := Spec{
		TypesPath: "../v1",
		Interfaces: []Interface{
			{
				Type:      &v1.MeshNode{},
				JSImport:  "node_pb",
				QueryType: "PEERS",
			},
			{
				Type:      &v1.MeshEdge{},
				JSImport:  "mesh_pb",
				QueryType: "EDGES",
			},
			{
				Type:      &v1.Role{},
				JSImport:  "rbac_pb",
				QueryType: "ROLES",
			},
			{
				Type:      &v1.RoleBinding{},
				JSImport:  "rbac_pb",
				QueryType: "ROLEBINDINGS",
			},
			{
				Type:      &v1.Group{},
				JSImport:  "rbac_pb",
				QueryType: "GROUPS",
			},
			{
				Type:      &v1.NetworkACL{},
				JSImport:  "network_acls_pb",
				QueryType: "ACLS",
			},
			{
				Type:      &v1.Route{},
				JSImport:  "network_acls_pb",
				QueryType: "ROUTES",
			},
		},
	}
	var buf strings.Builder
	err := t.ExecuteTemplate(&buf, "ts-rpcdb.ts.tmpl", map[string]any{
		"Spec":         &genspec,
		"TemplateName": "ts-rpcdb.ts.tmpl",
	})
	if err != nil {
		panic(err)
	}
	// Replace all tabs with 2 spaces.
	out := strings.ReplaceAll(buf.String(), "\t", "  ")
	fmt.Fprintln(outFile, out)
}
