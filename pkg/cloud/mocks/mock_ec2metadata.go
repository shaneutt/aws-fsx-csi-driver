// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/aws-fsx-csi-driver/pkg/cloud (interfaces: EC2Metadata)

// Package mocks is a generated GoMock package.
package mocks

import (
	ec2metadata "github.com/aws/aws-sdk-go/aws/ec2metadata"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEC2Metadata is a mock of EC2Metadata interface
type MockEC2Metadata struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MetadataMockRecorder
}

// MockEC2MetadataMockRecorder is the mock recorder for MockEC2Metadata
type MockEC2MetadataMockRecorder struct {
	mock *MockEC2Metadata
}

// NewMockEC2Metadata creates a new mock instance
func NewMockEC2Metadata(ctrl *gomock.Controller) *MockEC2Metadata {
	mock := &MockEC2Metadata{ctrl: ctrl}
	mock.recorder = &MockEC2MetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEC2Metadata) EXPECT() *MockEC2MetadataMockRecorder {
	return m.recorder
}

// Available mocks base method
func (m *MockEC2Metadata) Available() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Available")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Available indicates an expected call of Available
func (mr *MockEC2MetadataMockRecorder) Available() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Available", reflect.TypeOf((*MockEC2Metadata)(nil).Available))
}

// GetInstanceIdentityDocument mocks base method
func (m *MockEC2Metadata) GetInstanceIdentityDocument() (ec2metadata.EC2InstanceIdentityDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceIdentityDocument")
	ret0, _ := ret[0].(ec2metadata.EC2InstanceIdentityDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceIdentityDocument indicates an expected call of GetInstanceIdentityDocument
func (mr *MockEC2MetadataMockRecorder) GetInstanceIdentityDocument() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceIdentityDocument", reflect.TypeOf((*MockEC2Metadata)(nil).GetInstanceIdentityDocument))
}
