package route

import (
	"fmt"

	"github.com/transcom/mymove/pkg/models"
)

// ErrorCode contains error codes for the route package
type ErrorCode string

const (
	// UnroutableRoute happens when a valid route can't be calculated between two locations
	UnroutableRoute ErrorCode = "UNROUTABLE_ROUTE"
	// AddressLookupError happens when doing a LatLong lookup of an address
	AddressLookupError ErrorCode = "ADDRESS_LOOKUP_ERROR"
	// GeocodeResponseDecodingError happens when attempting to decode a geocode response
	GeocodeResponseDecodingError ErrorCode = "GEOCODE_RESPONSE_DECODE_ERROR"
	// RoutingResponseDecodingError happens when attempting to decode a routing response
	RoutingResponseDecodingError ErrorCode = "ROUTING_RESPONSE_DECODE_ERROR"
	// ShortHaulError - no moves under 50 miles
	ShortHaulError ErrorCode = "SHORT_HAUL_ERROR"
	// UnknownError is for when the cause of the error can't be ascertained
	UnknownError ErrorCode = "UNKNOWN_ERROR"
)

// Error is used for handling errors from the Route package
type Error interface {
	error
	Code() ErrorCode
}

// baseError contains basic route error functionality
type baseError struct {
	code ErrorCode
}

// Code returns the error code enum
func (b *baseError) Code() ErrorCode {
	return b.code
}

type responseError struct {
	baseError
	statusCode  int
	routingInfo string
}

func (e *responseError) Error() string {
	return fmt.Sprintf("Error when communicating with HERE server: (error_code: (%s), status_code: %d, routing_info: %s)", e.code, e.statusCode, e.routingInfo)
}

// NewUnroutableRouteError creates a new responseError error.
func NewUnroutableRouteError(statusCode int, source LatLong, dest LatLong) Error {
	return &responseError{
		baseError{UnroutableRoute},
		statusCode,
		fmt.Sprintf("source: (%s), dest: (%s", source.Coords(), dest.Coords()),
	}
}

// NewUnknownRoutingError returns an error for failed postal code lookups
func NewUnknownRoutingError(statusCode int, source LatLong, dest LatLong) Error {
	return &responseError{
		baseError{UnknownError},
		statusCode,
		fmt.Sprintf("source: (%s), dest: (%s", source.Coords(), dest.Coords()),
	}
}

// NewAddressLookupError returns a known error for failed address lookups
func NewAddressLookupError(statusCode int, a *models.Address) Error {
	return &responseError{
		baseError{AddressLookupError},
		statusCode,
		fmt.Sprint(a.LineFormat()),
	}
}

// NewUnknownAddressLookupError returns an unknown error for failed address lookups
func NewUnknownAddressLookupError(statusCode int, a *models.Address) Error {
	return &responseError{
		baseError{UnknownError},
		statusCode,
		fmt.Sprint(a.LineFormat()),
	}
}

type geocodeResponseDecodingError struct {
	baseError
	response GeocodeResponseBody
}

func (e *geocodeResponseDecodingError) Error() string {
	return fmt.Sprintf("Error trying to decode GeocodeResponse: %+v", e.response)
}

// NewGeocodeResponseDecodingError creates a new geocodeResponseDecodingError error.
func NewGeocodeResponseDecodingError(r GeocodeResponseBody) Error {
	return &geocodeResponseDecodingError{
		baseError{GeocodeResponseDecodingError},
		r,
	}
}

type routingResponseDecodingError struct {
	baseError
	response RoutingResponseBody
}

func (e *routingResponseDecodingError) Error() string {
	return fmt.Sprintf("Error trying to decode RoutingResponseBody: %+v", e.response)
}

// NewRoutingResponseDecodingError creates a new routingResponseDecodingError error.
func NewRoutingResponseDecodingError(r RoutingResponseBody) Error {
	return &routingResponseDecodingError{
		baseError{RoutingResponseDecodingError},
		r,
	}
}

type shortHaulError struct {
	baseError
	routingInfo string
}

// NewShortHaulError is the new short haul error
func NewShortHaulError(source LatLong, dest LatLong, moveDistance int) Error {
	return &shortHaulError{
		baseError{ShortHaulError},
		fmt.Sprintf("source: (%s), dest: (%s), distance: (%d)", source.Coords(), dest.Coords(), moveDistance),
	}
}

func (e *shortHaulError) Error() string {
	return fmt.Sprintf("Unsupported short haul move distance: (error_code: (%s), routing_info: %s)", e.code, e.routingInfo)
}
