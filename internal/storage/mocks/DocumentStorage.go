// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "course/internal/service/dto"

	mock "github.com/stretchr/testify/mock"

	model "course/internal/model"
)

// DocumentStorage is an autogenerated mock type for the DocumentStorage type
type DocumentStorage struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *DocumentStorage) Create(ctx context.Context, request *dto.CreateDocumentRequest) (*model.Document, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.Document
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateDocumentRequest) (*model.Document, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateDocumentRequest) *model.Document); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Document)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.CreateDocumentRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, request
func (_m *DocumentStorage) Delete(ctx context.Context, request *dto.DeleteDocumentRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.DeleteDocumentRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, request
func (_m *DocumentStorage) GetByID(ctx context.Context, request *dto.GetDocumentByIDRequest) (*model.Document, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *model.Document
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetDocumentByIDRequest) (*model.Document, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetDocumentByIDRequest) *model.Document); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Document)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetDocumentByIDRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByInfoCardID provides a mock function with given fields: ctx, request
func (_m *DocumentStorage) GetByInfoCardID(ctx context.Context, request *dto.GetDocumentByInfoCardIDRequest) (*model.Document, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetByInfoCardID")
	}

	var r0 *model.Document
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetDocumentByInfoCardIDRequest) (*model.Document, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetDocumentByInfoCardIDRequest) *model.Document); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Document)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetDocumentByInfoCardIDRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDocumentStorage creates a new instance of DocumentStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDocumentStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *DocumentStorage {
	mock := &DocumentStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}