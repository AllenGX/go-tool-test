package windows

import (
	"fmt"
	"os"
	"os/exec"
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
	err := os.Remove(path + "\\gengokit\\template\\NAME-service\\svc\\transport_http.gotemplate")
	if err != nil {
		fmt.Println("no file name is transport_http.gotemplate")
	}

	err = os.Remove(path + "\\gengokit\\template\\template.go")
	if err != nil {
		fmt.Println("no file name is template.go")
	}
}

func removedir(path string) {
	err := os.RemoveAll(path + "\\gengokit\\template\\NAME-service\\svc\\client\\http")
	if err != nil {
		fmt.Println("no dir name is http")
	}

	err = os.RemoveAll(path + "\\gengokit\\template\\NAME-service\\cmd")
	if err != nil {
		fmt.Println("no dir name is cmd")
	}
}

func updatefile(path string) {
	clifile, err := os.OpenFile(path+"\\gengokit\\template\\NAME-service\\svc\\server\\cli\\cli.gotemplate", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("cil not found")
	}
	_, err = clifile.WriteString(cliContext)
	if err != nil {
		fmt.Println("cil write fail")
	}
	runfile, err := os.OpenFile(path+"\\gengokit\\template\\NAME-service\\svc\\server\\run.gotemplate", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("run not found")
	}
	_, err = runfile.WriteString(runContext)
	if err != nil {
		fmt.Println("run write fail")
	}
}

func Option(path string) {
	updatefile(path)
	removedir(path)
	removefile(path)
}

func Windows(path string) {
	//windows run
	wininstall := exec.Command(path + "\\wininstall.bat")
	err := wininstall.Run()
	if err != nil {
		fmt.Println("wininstall run fail in windows")
	}
	time.Sleep(time.Second * 2)
}
