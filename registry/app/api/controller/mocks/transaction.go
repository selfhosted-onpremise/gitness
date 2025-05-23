// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// WithTx provides a mock function with given fields: ctx, fn, options
func (m *Transaction) WithTx(ctx context.Context, fn func(context.Context) error, options ...interface{}) error {
	var args []interface{}
	args = append(args, ctx, fn)
	for _, opt := range options {
		args = append(args, opt)
	}
	ret := m.Called(args...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error, ...interface{}) error); ok {
		r0 = rf(ctx, fn, options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
