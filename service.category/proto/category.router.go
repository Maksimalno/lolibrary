// Code generated by protoc-gen-gotemplate; DO NOT EDIT.

package categoryproto

import (
	"context"

	"github.com/monzo/typhon"
)

// -------------------------
// GET /service.category/read
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body GETReadCategoryRequest) Method() string {
	return "GET"
}

// Path is the HTTP path to this endpoint
func (body GETReadCategoryRequest) Path() string {
	return "/read"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body GETReadCategoryRequest) ServiceName() string {
	return "service.category"
}

// Host is the short-form service name, e.g. s-brand.
func (body GETReadCategoryRequest) Host() string {
	return "s-category"
}

// FullPath is the full routable URL to this service.
func (body GETReadCategoryRequest) FullPath() string {
	return "http://s-category/read"
}

// Request returns a typhon request for this type.
func (body GETReadCategoryRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body GETReadCategoryRequest) Response(ctx context.Context) (*GETReadCategoryResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body GETReadCategoryRequest) Send(ctx context.Context) *GETReadCategoryFuture {
	return &GETReadCategoryFuture{Future: body.Request(ctx).Send()}
}

// GETReadCategoryFuture is an intermediate future used for parallel requests with GETReadCategoryRequest
type GETReadCategoryFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *GETReadCategoryFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *GETReadCategoryFuture) DecodeResponse() (*GETReadCategoryResponse, error) {
	f.Done()

	body := &GETReadCategoryResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}

// -------------------------
// PUT /service.category/update
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body PUTUpdateCategoryRequest) Method() string {
	return "PUT"
}

// Path is the HTTP path to this endpoint
func (body PUTUpdateCategoryRequest) Path() string {
	return "/update"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body PUTUpdateCategoryRequest) ServiceName() string {
	return "service.category"
}

// Host is the short-form service name, e.g. s-brand.
func (body PUTUpdateCategoryRequest) Host() string {
	return "s-category"
}

// FullPath is the full routable URL to this service.
func (body PUTUpdateCategoryRequest) FullPath() string {
	return "http://s-category/update"
}

// Request returns a typhon request for this type.
func (body PUTUpdateCategoryRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body PUTUpdateCategoryRequest) Response(ctx context.Context) (*PUTUpdateCategoryResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body PUTUpdateCategoryRequest) Send(ctx context.Context) *PUTUpdateCategoryFuture {
	return &PUTUpdateCategoryFuture{Future: body.Request(ctx).Send()}
}

// PUTUpdateCategoryFuture is an intermediate future used for parallel requests with PUTUpdateCategoryRequest
type PUTUpdateCategoryFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *PUTUpdateCategoryFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *PUTUpdateCategoryFuture) DecodeResponse() (*PUTUpdateCategoryResponse, error) {
	f.Done()

	body := &PUTUpdateCategoryResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}

// -------------------------
// DELETE /service.category/delete
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body DELETERemoveCategoryRequest) Method() string {
	return "DELETE"
}

// Path is the HTTP path to this endpoint
func (body DELETERemoveCategoryRequest) Path() string {
	return "/delete"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body DELETERemoveCategoryRequest) ServiceName() string {
	return "service.category"
}

// Host is the short-form service name, e.g. s-brand.
func (body DELETERemoveCategoryRequest) Host() string {
	return "s-category"
}

// FullPath is the full routable URL to this service.
func (body DELETERemoveCategoryRequest) FullPath() string {
	return "http://s-category/delete"
}

// Request returns a typhon request for this type.
func (body DELETERemoveCategoryRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body DELETERemoveCategoryRequest) Response(ctx context.Context) (*DELETERemoveCategoryResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body DELETERemoveCategoryRequest) Send(ctx context.Context) *DELETERemoveCategoryFuture {
	return &DELETERemoveCategoryFuture{Future: body.Request(ctx).Send()}
}

// DELETERemoveCategoryFuture is an intermediate future used for parallel requests with DELETERemoveCategoryRequest
type DELETERemoveCategoryFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *DELETERemoveCategoryFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *DELETERemoveCategoryFuture) DecodeResponse() (*DELETERemoveCategoryResponse, error) {
	f.Done()

	body := &DELETERemoveCategoryResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}

// -------------------------
// POST /service.category/create
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body POSTCreateCategoryRequest) Method() string {
	return "POST"
}

// Path is the HTTP path to this endpoint
func (body POSTCreateCategoryRequest) Path() string {
	return "/create"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body POSTCreateCategoryRequest) ServiceName() string {
	return "service.category"
}

// Host is the short-form service name, e.g. s-brand.
func (body POSTCreateCategoryRequest) Host() string {
	return "s-category"
}

// FullPath is the full routable URL to this service.
func (body POSTCreateCategoryRequest) FullPath() string {
	return "http://s-category/create"
}

// Request returns a typhon request for this type.
func (body POSTCreateCategoryRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body POSTCreateCategoryRequest) Response(ctx context.Context) (*POSTCreateCategoryResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body POSTCreateCategoryRequest) Send(ctx context.Context) *POSTCreateCategoryFuture {
	return &POSTCreateCategoryFuture{Future: body.Request(ctx).Send()}
}

// POSTCreateCategoryFuture is an intermediate future used for parallel requests with POSTCreateCategoryRequest
type POSTCreateCategoryFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *POSTCreateCategoryFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *POSTCreateCategoryFuture) DecodeResponse() (*POSTCreateCategoryResponse, error) {
	f.Done()

	body := &POSTCreateCategoryResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}

// -------------------------
// GET /service.category/list
// -------------------------

// Method is the HTTP method used for this request.
// It is inferred from the name of the Request using a prefix match.
func (body GETListCategoriesRequest) Method() string {
	return "GET"
}

// Path is the HTTP path to this endpoint
func (body GETListCategoriesRequest) Path() string {
	return "/list"
}

// ServiceName is the long-form service name, e.g. service.brand.
func (body GETListCategoriesRequest) ServiceName() string {
	return "service.category"
}

// Host is the short-form service name, e.g. s-brand.
func (body GETListCategoriesRequest) Host() string {
	return "s-category"
}

// FullPath is the full routable URL to this service.
func (body GETListCategoriesRequest) FullPath() string {
	return "http://s-category/list"
}

// Request returns a typhon request for this type.
func (body GETListCategoriesRequest) Request(ctx context.Context) typhon.Request {
	return typhon.NewRequest(ctx, body.Method(), body.FullPath(), body)
}

// Response is a shortcut for .Send(ctx).DecodeResponse(), for when you do not need a future.
// This saves on boilerplate throughout the codebase and you should use this method unless you need parallel requests.
func (body GETListCategoriesRequest) Response(ctx context.Context) (*GETListCategoriesResponse, error) {
	return body.Send(ctx).DecodeResponse()
}

// Send creates a typhon future and immediately returns it.
// To wait for the request to complete and return the response, use DecodeResponse on the returned future.
func (body GETListCategoriesRequest) Send(ctx context.Context) *GETListCategoriesFuture {
	return &GETListCategoriesFuture{Future: body.Request(ctx).Send()}
}

// GETListCategoriesFuture is an intermediate future used for parallel requests with GETListCategoriesRequest
type GETListCategoriesFuture struct {
	Future   *typhon.ResponseFuture
	Response *typhon.Response
}

// Done waits for a response from a typhon future, and is safe to call multiple times in a row.
func (f *GETListCategoriesFuture) Done() {
	if f.Response == nil {
		rsp := f.Future.Response()
		f.Response = &rsp
	}
}

// DecodeResponse waits for this future to be done and then decodes the response into a concrete type.
func (f *GETListCategoriesFuture) DecodeResponse() (*GETListCategoriesResponse, error) {
	f.Done()

	body := &GETListCategoriesResponse{}
	if err := f.Response.Decode(body); err != nil {
		return nil, err
	}

	return body, nil
}
