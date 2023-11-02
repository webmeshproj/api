// tsutil-gen generates Typescript helper methods alongside the generated protobuf code.
package main

import (
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

var typescriptTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	"join": func(a []string, sep string) string {
		return strings.Join(a, sep)
	},
}).Parse(`
{{- $typesRoot := .Spec.TypesPath -}}
import { PromiseClient } from "@connectrpc/connect";
import { AppDaemon } from "{{ $typesRoot }}/app_connect.js";
{{ range $file, $types := .Spec.Imports -}}
import { {{ join $types "," }} } from "{{ $typesRoot }}/{{ $file }}.js";
{{ end -}}
import { 
	QueryRequest_QueryCommand,
	QueryRequest_QueryType,
} from "{{ $typesRoot }}/storage_query_pb.js";

{{ range .Spec.Interfaces }}
export class {{ .Name }}s {
	constructor(private readonly client: PromiseClient<AppDaemon>, private readonly connID: string) {}

	get(id: string): Promise<{{ .Name }}> {
		return new Promise((resolve, reject) => {
			const res = this.client.query({
				id: this.connID,
				query: {
					command: QueryRequest_QueryCommand.GET,
					type: QueryRequest_QueryType.{{ .QueryType }},
					query: 'id=' + id,
				}
			})
			return new {{ .Name }}.fromJson(res.items[0])
		});
	},

	list(): Promise<{{ .Name }}[]>;

	put(obj: {{ .Name }}): Promise<{{ .Name }}>;

	delete(id: string): Promise<void>;
}
{{- end }}
`))

func main() {
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
	err := typescriptTemplate.Execute(os.Stdout, map[string]any{
		"Spec": &genspec,
	})
	if err != nil {
		panic(err)
	}
}
