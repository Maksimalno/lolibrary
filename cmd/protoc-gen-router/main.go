package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	routerproto "github.com/lolibrary/lolibrary/cmd/protoc-gen-router/proto"
)

// TemplateService is a representation of a service proto.
type TemplateService struct {
	// Filename is the name of the file we're generating.
	Filename string

	// Package is the go package name.
	Package string

	// Name is the name of this service verbatim, snake_case
	Name string

	// Host is the name of this service, as a fully qualified host (service.foo-bar)
	Host string

	// ShortHost is the "short" version of Host, e.g. s-foo-bar
	ShortHost string

	// RPCs is a list of all RPCs in this service, with options attached to discover the path.
	RPCs []*TemplateRPC
}

// TemplateRPC is a field representing an RPC call in a service.
type TemplateRPC struct {
	// Method is the HTTP method for this RPC call, derived from the Request name.
	Method string

	// Path is the path to this RPC, starting with /
	Path string

	// Request is the Go type of the request.
	Request string

	// Response is the Go type of the response.
	Response string

	// Future is the "Request" name, with the suffix replaced with Future.
	// It is a Go type.
	Future string
}

func main() {
	g := generator.New()

	t, err := template.New("protoc-gen-router").Parse(tmpl)
	if err != nil {
		g.Error(err, "failed to create go template")
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "failed to read data from stdin")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "failed to unmarshal protobuf input")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("there is no file to generate")
	}

	files := g.Request.GetProtoFile()

	for _, file := range files {
		if len(file.Service) == 0 || len(file.Service) > 1 {
			// doesn't apply to us.
			continue
		}

		service := &TemplateService{
			Filename: file.GetName(),
			Package:  file.GetPackage(),
		}

		for _, svc := range file.Service {
			service.RPCs = make([]*TemplateRPC, 0, len(svc.Method))
			service.Name = svc.GetName()

			if svc.Options != nil {
				opts, err := proto.GetExtension(svc.Options, routerproto.E_Router)
				if err != nil {
					g.Error(err, "failed to get router extension")
				}

				if o, ok := opts.(*routerproto.ServiceRouter); ok && o != nil {
					service.Host = o.Name
					service.ShortHost = o.ShortName
				}
			}

			for _, rpc := range svc.Method {
				// get options for method.
				if rpc.Options == nil {
					g.Fail("Unable to find path. Have you used option (handler).path?")
				}

				opts, err := proto.GetExtension(rpc.Options, routerproto.E_Handler)
				if err != nil {
					g.Error(err, "failed to get handler extension")
				}

				r, ok := opts.(*routerproto.RPCHandler)
				if !ok || r == nil {
					g.Fail("option (handler) is not an RPC handler. Did you import the right .proto file?")
					return // to satisfy goland
				}

				service.RPCs = append(service.RPCs, rpcFromProto(rpc, r.Path, file.GetPackage()))
			}
		}

		// generate file using template
		var buf bytes.Buffer
		if err := t.Execute(&buf, service); err != nil {
			g.Error(err, "Failed to generate service. Check the template is correct?")
		}

		b, err := format.Source(buf.Bytes())
		if err != nil {
			g.Error(err, "Failed to format go source.")
		}

		name := strings.Replace(file.GetName(), ".proto", ".router.go", 1)
		err = ioutil.WriteFile(path.Join(os.Getenv("GOPATH"), "src", name), b, 0644)
		if err != nil {
			g.Error(err, "Failed to write file")
		}
	}
}

func rpcFromProto(rpc *descriptor.MethodDescriptorProto, path, pkg string) *TemplateRPC {
	prefix := fmt.Sprintf(".%s.", pkg)

	request := strings.TrimPrefix(rpc.GetInputType(), prefix)
	response := strings.TrimPrefix(rpc.GetOutputType(), prefix)
	future := fmt.Sprintf("%sFuture", strings.TrimSuffix(request, "Request"))
	method := getMethod(request)

	return &TemplateRPC{
		Path:     path,
		Method:   method,
		Request:  request,
		Response: response,
		Future:   future,
	}
}

func getMethod(req string) string {
	methods := [...]string{
		"GET",
		"PUT",
		"POST",
		"PATCH",
		"DELETE",
	}

	for _, method := range methods {
		if strings.HasPrefix(req, method) {
			return method
		}
	}

	return ""
}

const tmpl = `// Code generated by protoc-gen-gotemplate; DO NOT EDIT.

package {{ .Package }}

import (
	"context"

    "github.com/monzo/typhon"
)

{{- range .RPCs }}

    // -------------------------
    // {{ .Method }} /{{ $.Host }}{{ .Path }}
    // -------------------------

    // Method is the HTTP method used for this request.
    // It is inferred from the name of the Request using a prefix match.
    func (body {{ .Request}}) Method() string {
        return "{{ .Method }}"
    }

    // Path is the HTTP path to this endpoint
    func (body {{ .Request }}) Path() string {
        return "{{ .Path }}"
    }

    // ServiceName is the long-form service name, e.g. service.brand.
    func (body {{ .Request }}) ServiceName() string {
        return "{{ $.Host }}"
    }

    // Host is the short-form service name, e.g. s-brand.
    func (body {{ .Request }}) Host() string {
        return "{{ $.ShortHost }}"
    }

    // FullPath is the full routable URL to this service.
    func (body {{ .Request }}) FullPath() string {
        return "http://{{ $.ShortHost }}{{ .Path }}"
    }

    // Request returns a typhon request for this type.
    func (body {{ .Request }}) Request(ctx context.Context) typhon.Request {
        return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
    }

    // Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
    // This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
    func (body {{ .Request }}) Response(ctx context.Context) ({{ .Response }}, error) {
        return body.Send(ctx).DecodeResponse()
    }

    // Send creates a typhon future and immediately returns it.
    // To wait for the request to complete and return the response, use DecodeResponse on the returned future.
    func (body {{ .Request }}) Send(ctx context.Context) ({{ .Future }}, error) {
        return &{{ .Future }}{Future: body.Request(ctx).Send()}
    }

    // {{ .Future }} is an intermediate future used for parallel requests with {{ .Request }}
    type {{ .Future }} struct {
        Future   *typhon.ResponseFuture
        Response *typhon.Response
    }

    // Done waits for a response from a typhon future, and is safe to call multiple times in a row.
    func (f *{{ .Future }}) Done() {
        if f.Response == nil {
            rsp := f.Future.Response()
            f.Response = &rsp
        }
    }

    // DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
    func (f *{{ .Future }}) DecodeResponse() ({{ .Response }}, error) {
        f.Done()

        body := &{{ .Response }}{}
        if err := f.Response.Decode(body); err != nil {
            return nil, err
        }
    }

{{- end }}
`
