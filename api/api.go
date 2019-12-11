package api

import (
	"fmt"
	"net/http"

	"github.com/gcla/go-bindata-assetfs"
)

func Run(host string, port int) {
	dashboardMux := http.NewServeMux()
	// base html
	dashboardMux.Handle("/", http.FileServer(
		&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "vue_src/dist/"},
	))

	fmt.Printf("Starting dashboard at %s:%d\n", host, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), dashboardMux); err != nil {
		panic(err)
	}
}
