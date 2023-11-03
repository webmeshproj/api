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

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

// Identifiers are a map of identifiers to their description.
type Identifiers map[string]string

// Interface represents an interface for performing CRUD operations
// on a resource.
type Interface struct {
	Type                 proto.Message
	JSImport             string
	QueryType            string
	Identifiers          Identifiers
	ExtraGetIdentifiers  Identifiers
	ExtraListIdentifiers Identifiers
}

// Name returns the name of the interface.
func (i *Interface) Name() string {
	return string(proto.MessageName(i.Type).Name())
}

//go:embed templates/*.tmpl
var templateFS embed.FS

func title(a string) string {
	if a == "id" {
		return "ID"
	}
	if a == "nodeid" {
		return "NodeID"
	}
	out := make([]byte, len(a))
	_, _, _ = cases.Title(language.English).Transform(out, []byte(a), true)
	return string(out)
}

var funcMap = template.FuncMap{
	"join": func(a []string) string {
		return strings.Join(a, ", ")
	},
	"andJoin": func(a Identifiers) string {
		var ids []string
		for id := range a {
			ids = append(ids, title(id))
		}
		return strings.Join(ids, " and ")
	},
	"title": func(a string) string {
		return title(a)
	},
	"stringParams": func(a any) string {
		switch v := a.(type) {
		case Identifiers:
			var out []string
			for id := range v {
				out = append(out, fmt.Sprintf("%s: string", id))
			}
			return strings.Join(out, ", ")
		case string:
			return fmt.Sprintf("%s: string", v)
		default:
			panic(fmt.Sprintf("unknown type %T", a))
		}
	},
	"queryString": func(a any) string {
		switch v := a.(type) {
		case Identifiers:
			var out []string
			for id := range v {
				out = append(out, fmt.Sprintf("%s=${%s}", id, id))
			}
			return strings.Join(out, ",")
		case string:
			return fmt.Sprintf("%s=${%s}", v, v)
		default:
			panic(fmt.Sprintf("unknown type %T", a))
		}
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
				Identifiers: map[string]string{
					"id": "The ID of the node",
				},
				ExtraGetIdentifiers: map[string]string{
					"pubkey": "The base64 encoded public key of the node",
				},
			},
			{
				Type:      &v1.MeshEdge{},
				JSImport:  "mesh_pb",
				QueryType: "EDGES",
				Identifiers: map[string]string{
					"sourceid": "The ID of the source node",
					"targetid": "The ID of the target node",
				},
			},
			{
				Type:        &v1.Role{},
				JSImport:    "rbac_pb",
				QueryType:   "ROLES",
				Identifiers: map[string]string{"id": "The name of the role"},
			},
			{
				Type:        &v1.RoleBinding{},
				JSImport:    "rbac_pb",
				QueryType:   "ROLEBINDINGS",
				Identifiers: map[string]string{"id": "The name of the rolebinding"},
			},
			{
				Type:        &v1.Group{},
				JSImport:    "rbac_pb",
				QueryType:   "GROUPS",
				Identifiers: map[string]string{"id": "The name of the group"},
			},
			{
				Type:        &v1.NetworkACL{},
				JSImport:    "network_acls_pb",
				QueryType:   "ACLS",
				Identifiers: map[string]string{"id": "The name of the network ACL"},
			},
			{
				Type:        &v1.Route{},
				JSImport:    "network_acls_pb",
				QueryType:   "ROUTES",
				Identifiers: map[string]string{"id": "The name of the route"},
				ExtraListIdentifiers: map[string]string{
					"nodeid": "The ID of the node",
					"cidr":   "The CIDR of the route",
				},
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
