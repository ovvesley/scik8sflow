// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/server/connector/connector.go
//
// Generated by this command:
//
//	mockgen --source=pkg/server/connector/connector.go --destination=test/mock/connector_mock.go --package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	connector_deployment_k8s "github.com/ovvesley/scik8sflow/pkg/server/connector/connector_deployment_k8s"
	connector_job_k8s "github.com/ovvesley/scik8sflow/pkg/server/connector/connector_job_k8s"
	connector_metrics_k8s "github.com/ovvesley/scik8sflow/pkg/server/connector/connector_metrics_k8s"
	connector_namespace_k8s "github.com/ovvesley/scik8sflow/pkg/server/connector/connector_namespace_k8s"
	connector_pod_k8s "github.com/ovvesley/scik8sflow/pkg/server/connector/connector_pod_k8s"
	connector_pvc_k8s "github.com/ovvesley/scik8sflow/pkg/server/connector/connector_pvc_k8s"
	gomock "go.uber.org/mock/gomock"
)

// MockIConnector is a mock of IConnector interface.
type MockIConnector struct {
	ctrl     *gomock.Controller
	recorder *MockIConnectorMockRecorder
}

// MockIConnectorMockRecorder is the mock recorder for MockIConnector.
type MockIConnectorMockRecorder struct {
	mock *MockIConnector
}

// NewMockIConnector creates a new mock instance.
func NewMockIConnector(ctrl *gomock.Controller) *MockIConnector {
	mock := &MockIConnector{ctrl: ctrl}
	mock.recorder = &MockIConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIConnector) EXPECT() *MockIConnectorMockRecorder {
	return m.recorder
}

// Deployment mocks base method.
func (m *MockIConnector) Deployment() connector_deployment_k8s.IConnectorDeployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployment")
	ret0, _ := ret[0].(connector_deployment_k8s.IConnectorDeployment)
	return ret0
}

// Deployment indicates an expected call of Deployment.
func (mr *MockIConnectorMockRecorder) Deployment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployment", reflect.TypeOf((*MockIConnector)(nil).Deployment))
}

// Job mocks base method.
func (m *MockIConnector) Job() connector_job_k8s.IConnectorJob {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Job")
	ret0, _ := ret[0].(connector_job_k8s.IConnectorJob)
	return ret0
}

// Job indicates an expected call of Job.
func (mr *MockIConnectorMockRecorder) Job() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Job", reflect.TypeOf((*MockIConnector)(nil).Job))
}

// Metrics mocks base method.
func (m *MockIConnector) Metrics() connector_metrics_k8s.IConnectorMetrics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metrics")
	ret0, _ := ret[0].(connector_metrics_k8s.IConnectorMetrics)
	return ret0
}

// Metrics indicates an expected call of Metrics.
func (mr *MockIConnectorMockRecorder) Metrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metrics", reflect.TypeOf((*MockIConnector)(nil).Metrics))
}

// Namespace mocks base method.
func (m *MockIConnector) Namespace() connector_namespace_k8s.IConnectorNamespace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(connector_namespace_k8s.IConnectorNamespace)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MockIConnectorMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockIConnector)(nil).Namespace))
}

// PersistentVolumeClain mocks base method.
func (m *MockIConnector) PersistentVolumeClain() connector_pvc_k8s.IConnectorPvc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersistentVolumeClain")
	ret0, _ := ret[0].(connector_pvc_k8s.IConnectorPvc)
	return ret0
}

// PersistentVolumeClain indicates an expected call of PersistentVolumeClain.
func (mr *MockIConnectorMockRecorder) PersistentVolumeClain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistentVolumeClain", reflect.TypeOf((*MockIConnector)(nil).PersistentVolumeClain))
}

// Pod mocks base method.
func (m *MockIConnector) Pod() connector_pod_k8s.IConnectorPod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pod")
	ret0, _ := ret[0].(connector_pod_k8s.IConnectorPod)
	return ret0
}

// Pod indicates an expected call of Pod.
func (mr *MockIConnectorMockRecorder) Pod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pod", reflect.TypeOf((*MockIConnector)(nil).Pod))
}
