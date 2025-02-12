// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewMassaGetAddressesParams creates a new MassaGetAddressesParams object
//
// There are no default values defined in the spec.
func NewMassaGetAddressesParams() MassaGetAddressesParams {

	return MassaGetAddressesParams{}
}

// MassaGetAddressesParams contains all the bound params for the massa get addresses operation
// typically these are obtained from a http.Request
//
// swagger:parameters massaGetAddresses
type MassaGetAddressesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*list of wanted addresses
	  Required: true
	  In: query
	  Collection Format: multi
	*/
	Addresses []string
	/*Specifies the attributes to return. If no attributes are provided, they are all returned.
	Possible values:

	| Attribute | Content |
	| ----------- | -----------|
	| balance | the pending balances (takes into account pending/non-final operations) and the final balances (takes into account only final operations). |

	  In: query
	  Collection Format: multi
	*/
	Attributes []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewMassaGetAddressesParams() beforehand.
func (o *MassaGetAddressesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAddresses, qhkAddresses, _ := qs.GetOK("addresses")
	if err := o.bindAddresses(qAddresses, qhkAddresses, route.Formats); err != nil {
		res = append(res, err)
	}

	qAttributes, qhkAttributes, _ := qs.GetOK("attributes")
	if err := o.bindAttributes(qAttributes, qhkAttributes, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAddresses binds and validates array parameter Addresses from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *MassaGetAddressesParams) bindAddresses(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("addresses", "query", rawData)
	}
	// CollectionFormat: multi
	addressesIC := rawData
	if len(addressesIC) == 0 {
		return errors.Required("addresses", "query", addressesIC)
	}

	var addressesIR []string
	for _, addressesIV := range addressesIC {
		addressesI := addressesIV

		addressesIR = append(addressesIR, addressesI)
	}

	o.Addresses = addressesIR

	return nil
}

// bindAttributes binds and validates array parameter Attributes from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *MassaGetAddressesParams) bindAttributes(rawData []string, hasKey bool, formats strfmt.Registry) error {
	// CollectionFormat: multi
	attributesIC := rawData
	if len(attributesIC) == 0 {
		return nil
	}

	var attributesIR []string
	for i, attributesIV := range attributesIC {
		attributesI := attributesIV

		if err := validate.EnumCase(fmt.Sprintf("%s.%v", "attributes", i), "query", attributesI, []interface{}{"balance"}, true); err != nil {
			return err
		}

		attributesIR = append(attributesIR, attributesI)
	}

	o.Attributes = attributesIR

	return nil
}
