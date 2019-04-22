package handle

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/flosch/pongo2"
	"go.etcd.io/etcd/client"
)

var dir string
var kapi client.KeysAPI
var cfg = client.Config{
	Endpoints:               []string{"http://127.0.0.1:2379"},
	Transport:               client.DefaultTransport,
	HeaderTimeoutPerRequest: time.Second,
}

func get(k string) string {
	resp, err := kapi.Get(context.Background(), dir+k, nil)
	if err != nil {
		fmt.Println("Error: check etcd.")
		os.Exit(0)
	}

	v := string(resp.Node.Value)
	lastIndex := len(v) - 1
	if lastIndex >= 0 && v[lastIndex] == '\n' {
		v = v[0:lastIndex]
	}

	return v
}

func init() {
	c, err := client.New(cfg)
	if err != nil {
		fmt.Println("Error: client", err)
		os.Exit(0)
	}
	kapi = client.NewKeysAPI(c)

}

func stream(buf []byte, sub string, d string) []byte {
	dir = d
	t, err := pongo2.FromString(string(buf))
	if err != nil {
		fmt.Println("Error: check file's format")
		os.Exit(0)
	}

	out, err := t.Execute(pongo2.Context{sub: get})
	if err != nil {
		fmt.Println("Error: exec", err)
		os.Exit(0)
	}

	return []byte(out)
}
