// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DepositResourceHandlerFunc turns a function with the right signature into a deposit resource handler
type DepositResourceHandlerFunc func(DepositResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DepositResourceHandlerFunc) Handle(params DepositResourceParams) middleware.Responder {
	return fn(params)
}

// DepositResourceHandler interface for that can handle valid deposit resource params
type DepositResourceHandler interface {
	Handle(DepositResourceParams) middleware.Responder
}

// NewDepositResource creates a new http.Handler for the deposit resource operation
func NewDepositResource(ctx *middleware.Context, handler DepositResourceHandler) *DepositResource {
	return &DepositResource{Context: ctx, Handler: handler}
}

/*DepositResource swagger:route POST /resource depositResource

Deposit New TACO Resource.

Deposits a new resource (Collection, Digital Repository Object, File [metadata only] or subclass of those) into SDR. Will return the SDR identifier for the resource.

*/
type DepositResource struct {
	Context *middleware.Context
	Handler DepositResourceHandler
}

func (o *DepositResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDepositResourceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}