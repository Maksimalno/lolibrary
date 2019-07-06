// Code generated by protoc-gen-gotemplate; DO NOT EDIT.

package brandproto

import (
	"context"

	"github.com/monzo/typhon"
)

// -------------------------
// GET /service.brand/read
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body GETReadBrandRequest) Method() string {
	return "GET"
}

// Path is the HTTP path to this endpoint
func (body GETReadBrandRequest) Path() string {
	return "/read"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body GETReadBrandRequest) ServiceName() string {
	return "service.brand"
}

// Host is the short-form service name, e.g. s-brand.
func (body GETReadBrandRequest) Host() string {
	return "s-brand"
}

// FullPath is the full routable URL to this service.
func (body GETReadBrandRequest) FullPath() string {
	return "http://s-brand/read"
}

// Request returns a typhon request for this type.
func (body GETReadBrandRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body GETReadBrandRequest) Response(ctx context.Context) (*GETReadBrandResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body GETReadBrandRequest) Send(ctx context.Context) *GETReadBrandFuture {
	return &GETReadBrandFuture{Future: body.Request(ctx).Send()}
}

// SendVia creates a typhon future and immediately returns it, passing the request through svc.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body GETReadBrandRequest) SendVia(ctx context.Context, svc typhon.Service) *GETReadBrandFuture {
	return &GETReadBrandFuture{Future: body.Request(ctx).SendVia(svc)}
}

// GETReadBrandFuture is an intermediate future used for parallel requests with GETReadBrandRequest
type GETReadBrandFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *GETReadBrandFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *GETReadBrandFuture) DecodeResponse() (*GETReadBrandResponse, error) {
	f.Done()

	body := &GETReadBrandResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}

// -------------------------
// PUT /service.brand/update
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body PUTUpdateBrandRequest) Method() string {
	return "PUT"
}

// Path is the HTTP path to this endpoint
func (body PUTUpdateBrandRequest) Path() string {
	return "/update"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body PUTUpdateBrandRequest) ServiceName() string {
	return "service.brand"
}

// Host is the short-form service name, e.g. s-brand.
func (body PUTUpdateBrandRequest) Host() string {
	return "s-brand"
}

// FullPath is the full routable URL to this service.
func (body PUTUpdateBrandRequest) FullPath() string {
	return "http://s-brand/update"
}

// Request returns a typhon request for this type.
func (body PUTUpdateBrandRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body PUTUpdateBrandRequest) Response(ctx context.Context) (*PUTUpdateBrandResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body PUTUpdateBrandRequest) Send(ctx context.Context) *PUTUpdateBrandFuture {
	return &PUTUpdateBrandFuture{Future: body.Request(ctx).Send()}
}

// SendVia creates a typhon future and immediately returns it, passing the request through svc.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body PUTUpdateBrandRequest) SendVia(ctx context.Context, svc typhon.Service) *PUTUpdateBrandFuture {
	return &PUTUpdateBrandFuture{Future: body.Request(ctx).SendVia(svc)}
}

// PUTUpdateBrandFuture is an intermediate future used for parallel requests with PUTUpdateBrandRequest
type PUTUpdateBrandFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *PUTUpdateBrandFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *PUTUpdateBrandFuture) DecodeResponse() (*PUTUpdateBrandResponse, error) {
	f.Done()

	body := &PUTUpdateBrandResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}

// -------------------------
// DELETE /service.brand/delete
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body DELETERemoveBrandRequest) Method() string {
	return "DELETE"
}

// Path is the HTTP path to this endpoint
func (body DELETERemoveBrandRequest) Path() string {
	return "/delete"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body DELETERemoveBrandRequest) ServiceName() string {
	return "service.brand"
}

// Host is the short-form service name, e.g. s-brand.
func (body DELETERemoveBrandRequest) Host() string {
	return "s-brand"
}

// FullPath is the full routable URL to this service.
func (body DELETERemoveBrandRequest) FullPath() string {
	return "http://s-brand/delete"
}

// Request returns a typhon request for this type.
func (body DELETERemoveBrandRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body DELETERemoveBrandRequest) Response(ctx context.Context) (*DELETERemoveBrandResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body DELETERemoveBrandRequest) Send(ctx context.Context) *DELETERemoveBrandFuture {
	return &DELETERemoveBrandFuture{Future: body.Request(ctx).Send()}
}

// SendVia creates a typhon future and immediately returns it, passing the request through svc.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body DELETERemoveBrandRequest) SendVia(ctx context.Context, svc typhon.Service) *DELETERemoveBrandFuture {
	return &DELETERemoveBrandFuture{Future: body.Request(ctx).SendVia(svc)}
}

// DELETERemoveBrandFuture is an intermediate future used for parallel requests with DELETERemoveBrandRequest
type DELETERemoveBrandFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *DELETERemoveBrandFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *DELETERemoveBrandFuture) DecodeResponse() (*DELETERemoveBrandResponse, error) {
	f.Done()

	body := &DELETERemoveBrandResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}

// -------------------------
// POST /service.brand/create
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body POSTCreateBrandRequest) Method() string {
	return "POST"
}

// Path is the HTTP path to this endpoint
func (body POSTCreateBrandRequest) Path() string {
	return "/create"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body POSTCreateBrandRequest) ServiceName() string {
	return "service.brand"
}

// Host is the short-form service name, e.g. s-brand.
func (body POSTCreateBrandRequest) Host() string {
	return "s-brand"
}

// FullPath is the full routable URL to this service.
func (body POSTCreateBrandRequest) FullPath() string {
	return "http://s-brand/create"
}

// Request returns a typhon request for this type.
func (body POSTCreateBrandRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body POSTCreateBrandRequest) Response(ctx context.Context) (*POSTCreateBrandResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body POSTCreateBrandRequest) Send(ctx context.Context) *POSTCreateBrandFuture {
	return &POSTCreateBrandFuture{Future: body.Request(ctx).Send()}
}

// SendVia creates a typhon future and immediately returns it, passing the request through svc.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body POSTCreateBrandRequest) SendVia(ctx context.Context, svc typhon.Service) *POSTCreateBrandFuture {
	return &POSTCreateBrandFuture{Future: body.Request(ctx).SendVia(svc)}
}

// POSTCreateBrandFuture is an intermediate future used for parallel requests with POSTCreateBrandRequest
type POSTCreateBrandFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *POSTCreateBrandFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *POSTCreateBrandFuture) DecodeResponse() (*POSTCreateBrandResponse, error) {
	f.Done()

	body := &POSTCreateBrandResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}

// -------------------------
// GET /service.brand/list
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body GETListBrandsRequest) Method() string {
	return "GET"
}

// Path is the HTTP path to this endpoint
func (body GETListBrandsRequest) Path() string {
	return "/list"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body GETListBrandsRequest) ServiceName() string {
	return "service.brand"
}

// Host is the short-form service name, e.g. s-brand.
func (body GETListBrandsRequest) Host() string {
	return "s-brand"
}

// FullPath is the full routable URL to this service.
func (body GETListBrandsRequest) FullPath() string {
	return "http://s-brand/list"
}

// Request returns a typhon request for this type.
func (body GETListBrandsRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body GETListBrandsRequest) Response(ctx context.Context) (*GETListBrandsResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body GETListBrandsRequest) Send(ctx context.Context) *GETListBrandsFuture {
	return &GETListBrandsFuture{Future: body.Request(ctx).Send()}
}

// SendVia creates a typhon future and immediately returns it, passing the request through svc.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body GETListBrandsRequest) SendVia(ctx context.Context, svc typhon.Service) *GETListBrandsFuture {
	return &GETListBrandsFuture{Future: body.Request(ctx).SendVia(svc)}
}

// GETListBrandsFuture is an intermediate future used for parallel requests with GETListBrandsRequest
type GETListBrandsFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *GETListBrandsFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *GETListBrandsFuture) DecodeResponse() (*GETListBrandsResponse, error) {
	f.Done()

	body := &GETListBrandsResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}
