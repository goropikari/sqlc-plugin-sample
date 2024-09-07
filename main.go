package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/goropikari/sqlc_plugin_sample/proto/sqlc/plugin"
	"google.golang.org/protobuf/proto"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error generating sample: %s", err)
		os.Exit(2)
	}
}

func run() error {
	req, err := ParseRequest(os.Stdin)
	if err != nil {
		return err
	}

	resp, err := Generate(req)
	if err != nil {
		return err
	}

	respBlob, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(os.Stdout)
	if _, err := w.Write(respBlob); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}

func ParseRequest(r io.Reader) (*plugin.GenerateRequest, error) {
	var req plugin.GenerateRequest
	reqBlob, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(reqBlob, &req); err != nil {
		return nil, err
	}

	return &req, nil
}

func Generate(req *plugin.GenerateRequest) (*plugin.GenerateResponse, error) {
	w := new(bytes.Buffer)

	opts := struct {
		Filename string
		Foo      string
		Bar      string
	}{}

	if err := json.Unmarshal(req.GetPluginOptions(), &opts); err != nil {
		return nil, err
	}

	w.WriteString(opts.Foo + "\n")
	w.WriteString(opts.Bar + "\n")
	for _, schema := range req.GetCatalog().GetSchemas() {
		for _, tb := range schema.GetTables() {
			w.WriteString(tb.GetRel().GetName() + "\n")
			for _, col := range tb.GetColumns() {
				w.WriteString("\t" + col.GetName() + "\n")
			}
		}
	}

	return &plugin.GenerateResponse{
		Files: []*plugin.File{
			{
				Name:     opts.Filename,
				Contents: w.Bytes(),
			},
		},
	}, nil
}
