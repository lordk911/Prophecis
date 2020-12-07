// Code generated by go-swagger; DO NOT EDIT.

package logins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUMLoginParams creates a new UMLoginParams object
// no default values defined in spec.
func NewUMLoginParams() UMLoginParams {

	return UMLoginParams{}
}

// UMLoginParams contains all the bound params for the u m login operation
// typically these are obtained from a http.Request
//
// swagger:parameters UMLogin
type UMLoginParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the password
	  Required: true
	  In: query
	*/
	Password string
	/*the pwdEncoded
	  In: query
	*/
	PwdEncoded *bool
	/*the username
	  Required: true
	  In: query
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUMLoginParams() beforehand.
func (o *UMLoginParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qPassword, qhkPassword, _ := qs.GetOK("password")
	if err := o.bindPassword(qPassword, qhkPassword, route.Formats); err != nil {
		res = append(res, err)
	}

	qPwdEncoded, qhkPwdEncoded, _ := qs.GetOK("pwdEncoded")
	if err := o.bindPwdEncoded(qPwdEncoded, qhkPwdEncoded, route.Formats); err != nil {
		res = append(res, err)
	}

	qUsername, qhkUsername, _ := qs.GetOK("username")
	if err := o.bindUsername(qUsername, qhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPassword binds and validates parameter Password from query.
func (o *UMLoginParams) bindPassword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("password", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("password", "query", raw); err != nil {
		return err
	}

	o.Password = raw

	return nil
}

// bindPwdEncoded binds and validates parameter PwdEncoded from query.
func (o *UMLoginParams) bindPwdEncoded(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("pwdEncoded", "query", "bool", raw)
	}
	o.PwdEncoded = &value

	return nil
}

// bindUsername binds and validates parameter Username from query.
func (o *UMLoginParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("username", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("username", "query", raw); err != nil {
		return err
	}

	o.Username = raw

	return nil
}