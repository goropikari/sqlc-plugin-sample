package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/goropikari/sqlc_plugin_sample/proto/sqlc/plugin"
	"google.golang.org/protobuf/proto"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, "error generating dump")
		os.Exit(2)
	}
}

func run() error {
	reqBlob, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(os.Stdout)
	resp := &plugin.GenerateResponse{
		Files: []*plugin.File{
			{
				Name:     "dump_proto",
				Contents: reqBlob,
			},
		},
	}
	respBlob, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	if _, err := w.Write(respBlob); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}
