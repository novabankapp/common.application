package errors

import (
	"context"
	"reflect"
)

func IsType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

type ErrorMapping struct {
	FromErrors     []error
	toStatusCode   int
	ToResponseFunc func(ctx context.Context, err error)
}

func (r *ErrorMapping) ToStatusCode(statusCode int) *ErrorMapping {
	r.toStatusCode = statusCode
	r.ToResponseFunc = func(ctx context.Context, err error) {

		context.WithValue(ctx, "StatusCode", statusCode)
	}
	return r
}

func (r *ErrorMapping) ToResponse(response func(ctx context.Context, err error)) *ErrorMapping {
	r.ToResponseFunc = response
	return r
}

func Map(err ...error) *ErrorMapping {
	return &ErrorMapping{
		FromErrors: err,
	}
}
