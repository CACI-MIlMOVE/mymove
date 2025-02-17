// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	appcontext "github.com/transcom/mymove/pkg/appcontext"

	uuid "github.com/gofrs/uuid"
)

// EvaluationReportDeleter is an autogenerated mock type for the EvaluationReportDeleter type
type EvaluationReportDeleter struct {
	mock.Mock
}

// DeleteEvaluationReport provides a mock function with given fields: appCtx, reportID
func (_m *EvaluationReportDeleter) DeleteEvaluationReport(appCtx appcontext.AppContext, reportID uuid.UUID) error {
	ret := _m.Called(appCtx, reportID)

	var r0 error
	if rf, ok := ret.Get(0).(func(appcontext.AppContext, uuid.UUID) error); ok {
		r0 = rf(appCtx, reportID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEvaluationReportDeleter creates a new instance of EvaluationReportDeleter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvaluationReportDeleter(t interface {
	mock.TestingT
	Cleanup(func())
}) *EvaluationReportDeleter {
	mock := &EvaluationReportDeleter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
