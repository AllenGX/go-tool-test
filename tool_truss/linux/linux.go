package linux

import (
	"fmt"
	"os"
	"os/exec"
	"test/tool_truss/windows"
	"time"
)

var (
	cliContext = `
package cli

    import (
        "{{.ImportPath -}} /svc/server"
    )

    var Config server.Config

    func init() {
        Config.GRPCAddr = "127.0.0.1:8888"
	}
	`
	runContext = `
package server

    import (
        "log"
        "net"
        // 3d Party
        "google.golang.org/grpc"

        // This Service
        pb "{{.PBImportPath -}}"
        "{{.ImportPath -}} /svc"
        "{{.ImportPath -}} /handlers"
    )

    // Config contains the required fields for running a server
    type Config struct {
        GRPCAddr string
    }

    func NewEndpoints() svc.Endpoints {
        // Business domain.
        var service pb.{{.Service.Name}}Server
        {
            service = handlers.NewService()
            // Wrap Service with middlewares. See handlers/middlewares.go
            service = handlers.WrapService(service)
        }

        // Endpoint domain.
        var (
        {{range $i := .Service.Methods -}}
            {{ToLower $i.Name}}Endpoint = svc.Make{{$i.Name}}Endpoint(service)
        {{end}}
        )

        endpoints := svc.Endpoints{
        {{range $i := .Service.Methods -}}
            {{$i.Name}}Endpoint:    {{ToLower $i.Name}}Endpoint,
        {{end}}
        }

        // Wrap selected Endpoints with middlewares. See handlers/middlewares.go
        endpoints = handlers.WrapEndpoints(endpoints)

        return endpoints
    }

    // Run starts a new http server, gRPC server, and a debug server with the
    // passed config and logger
    func Run(cfg Config) {
        endpoints := NewEndpoints()

        // Mechanical domain.
        errc := make(chan error)

        // Interrupt handler.
        go handlers.InterruptHandler(errc)

        // gRPC transport.
        go func() {
            log.Println("transport", "gRPC","addr", cfg.GRPCAddr)
            ln, err := net.Listen("tcp", cfg.GRPCAddr)
            if err != nil {
                errc <- err
                return
            }

            srv := svc.MakeGRPCServer(endpoints)
            s := grpc.NewServer()
            pb.Register{{.Service.Name}}Server(s, srv)

            errc <- s.Serve(ln)
        }()

        // Run!
        log.Println("exit", <-errc)
	}
	`
)

func removefile(path string) {
	err := os.Remove(path + "/gengokit/template/NAME-service/svc/transport_http.gotemplate")
	if err != nil {
		fmt.Println("no file name is transport_http.gotemplate")
	}

	err = os.Remove(path + "/gengokit/template/template.go")
	if err != nil {
		fmt.Println("no file name is template.go")
	}
}

func removedir(path string) {
	err := os.RemoveAll(path + "/gengokit/template/NAME-service/svc/client/http")
	if err != nil {
		fmt.Println("no dir name is http")
	}

	err = os.RemoveAll(path + "/gengokit/template/NAME-service/cmd")
	if err != nil {
		fmt.Println("no dir name is cmd")
	}
}

func updatefile(path string) {
	clifile, err := os.OpenFile(path+"/gengokit/template/NAME-service/svc/server/cli/cli.gotemplate", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("cil not found")
	}
	_, err = clifile.WriteString(cliContext)
	if err != nil {
		fmt.Println("cil write fail")
	}
	runfile, err := os.OpenFile(path+"/gengokit/template/NAME-service/svc/server/run.gotemplate", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("run not found")
	}
	_, err = runfile.WriteString(runContext)
	if err != nil {
		fmt.Println("run write fail")
	}
}

func Option(path string) {
	windows.Option(path)

}

func Linux(path string) {
	cli := exec.Command(`
go get github.com/pauln/go-datefmt
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/golang/protobuf/proto
go get github.com/jteeuwen/go-bindata/...
go generate github.com/tuneinc/truss/gengokit/template
	`)
	err := cli.Run()
	if err != nil {
		fmt.Println("go get fail")
	}

	linux := exec.Command(`go install -ldflags "-X 'main.Version=7dc4d5d85c' -X 'main.VersionDate=Mon May 28 22:12:59 UTC 2018'" github.com/tuneinc/truss/cmd/truss`)
	err = linux.Run()
	if err != nil {
		fmt.Println("fail to install linux")
	}
	time.Sleep(time.Second * 2)
}
