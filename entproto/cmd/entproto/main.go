// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"
	"os"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	var (
		schemaPath = flag.String("path", "", "path to schema directory")
		pbPath     = flag.String("pb_out", "", "path to output directory for generated pb.go files")
		pbName     = flag.String("pb_name", "", "name of the generated pb.go files")
	)
	flag.Parse()
	if *schemaPath == "" {
		log.Fatal("entproto: must specify schema path. use entproto -path ./ent/schema")
	}
	if *pbPath == "" {
		log.Fatal("entproto: must specify path to output directory for generated pb.go files. use entproto -pb_out ./entpb")
	}

	entproto.SetDefaultProtoServiceName(*pbName)

	graph, err := entc.LoadGraph(*schemaPath, &gen.Config{})
	if err != nil {
		log.Fatalf("entproto: failed loading ent graph: %v", err)
	}
	if err := entproto.Generate(graph, pbPath); err != nil {
		log.Fatalf("entproto: failed generating protos: %s", err)
	}
	if os.Rename(*pbPath+"pb/pb.proto", *pbPath+"pb/"+*pbName+".proto"); err != nil {
		log.Fatalf("entproto: failed renaming file: %s", err)
	}
}
