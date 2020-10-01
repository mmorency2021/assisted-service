// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/restapi (interfaces: InstallerAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	middleware "github.com/go-openapi/runtime/middleware"
	gomock "github.com/golang/mock/gomock"
	installer "github.com/openshift/assisted-service/restapi/operations/installer"
	reflect "reflect"
)

// MockInstallerAPI is a mock of InstallerAPI interface
type MockInstallerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerAPIMockRecorder
}

// MockInstallerAPIMockRecorder is the mock recorder for MockInstallerAPI
type MockInstallerAPIMockRecorder struct {
	mock *MockInstallerAPI
}

// NewMockInstallerAPI creates a new mock instance
func NewMockInstallerAPI(ctrl *gomock.Controller) *MockInstallerAPI {
	mock := &MockInstallerAPI{ctrl: ctrl}
	mock.recorder = &MockInstallerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstallerAPI) EXPECT() *MockInstallerAPIMockRecorder {
	return m.recorder
}

// CancelInstallation mocks base method
func (m *MockInstallerAPI) CancelInstallation(arg0 context.Context, arg1 installer.CancelInstallationParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelInstallation", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// CancelInstallation indicates an expected call of CancelInstallation
func (mr *MockInstallerAPIMockRecorder) CancelInstallation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelInstallation", reflect.TypeOf((*MockInstallerAPI)(nil).CancelInstallation), arg0, arg1)
}

// CompleteInstallation mocks base method
func (m *MockInstallerAPI) CompleteInstallation(arg0 context.Context, arg1 installer.CompleteInstallationParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteInstallation", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// CompleteInstallation indicates an expected call of CompleteInstallation
func (mr *MockInstallerAPIMockRecorder) CompleteInstallation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteInstallation", reflect.TypeOf((*MockInstallerAPI)(nil).CompleteInstallation), arg0, arg1)
}

// DeregisterCluster mocks base method
func (m *MockInstallerAPI) DeregisterCluster(arg0 context.Context, arg1 installer.DeregisterClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DeregisterCluster indicates an expected call of DeregisterCluster
func (mr *MockInstallerAPIMockRecorder) DeregisterCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterCluster", reflect.TypeOf((*MockInstallerAPI)(nil).DeregisterCluster), arg0, arg1)
}

// DeregisterHost mocks base method
func (m *MockInstallerAPI) DeregisterHost(arg0 context.Context, arg1 installer.DeregisterHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DeregisterHost indicates an expected call of DeregisterHost
func (mr *MockInstallerAPIMockRecorder) DeregisterHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterHost", reflect.TypeOf((*MockInstallerAPI)(nil).DeregisterHost), arg0, arg1)
}

// DisableHost mocks base method
func (m *MockInstallerAPI) DisableHost(arg0 context.Context, arg1 installer.DisableHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DisableHost indicates an expected call of DisableHost
func (mr *MockInstallerAPIMockRecorder) DisableHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableHost", reflect.TypeOf((*MockInstallerAPI)(nil).DisableHost), arg0, arg1)
}

// DownloadClusterFiles mocks base method
func (m *MockInstallerAPI) DownloadClusterFiles(arg0 context.Context, arg1 installer.DownloadClusterFilesParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadClusterFiles", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DownloadClusterFiles indicates an expected call of DownloadClusterFiles
func (mr *MockInstallerAPIMockRecorder) DownloadClusterFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadClusterFiles", reflect.TypeOf((*MockInstallerAPI)(nil).DownloadClusterFiles), arg0, arg1)
}

// DownloadClusterISO mocks base method
func (m *MockInstallerAPI) DownloadClusterISO(arg0 context.Context, arg1 installer.DownloadClusterISOParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadClusterISO", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DownloadClusterISO indicates an expected call of DownloadClusterISO
func (mr *MockInstallerAPIMockRecorder) DownloadClusterISO(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadClusterISO", reflect.TypeOf((*MockInstallerAPI)(nil).DownloadClusterISO), arg0, arg1)
}

// DownloadClusterKubeconfig mocks base method
func (m *MockInstallerAPI) DownloadClusterKubeconfig(arg0 context.Context, arg1 installer.DownloadClusterKubeconfigParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadClusterKubeconfig", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DownloadClusterKubeconfig indicates an expected call of DownloadClusterKubeconfig
func (mr *MockInstallerAPIMockRecorder) DownloadClusterKubeconfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadClusterKubeconfig", reflect.TypeOf((*MockInstallerAPI)(nil).DownloadClusterKubeconfig), arg0, arg1)
}

// DownloadClusterLogs mocks base method
func (m *MockInstallerAPI) DownloadClusterLogs(arg0 context.Context, arg1 installer.DownloadClusterLogsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadClusterLogs", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DownloadClusterLogs indicates an expected call of DownloadClusterLogs
func (mr *MockInstallerAPIMockRecorder) DownloadClusterLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadClusterLogs", reflect.TypeOf((*MockInstallerAPI)(nil).DownloadClusterLogs), arg0, arg1)
}

// DownloadHostLogs mocks base method
func (m *MockInstallerAPI) DownloadHostLogs(arg0 context.Context, arg1 installer.DownloadHostLogsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadHostLogs", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DownloadHostLogs indicates an expected call of DownloadHostLogs
func (mr *MockInstallerAPIMockRecorder) DownloadHostLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadHostLogs", reflect.TypeOf((*MockInstallerAPI)(nil).DownloadHostLogs), arg0, arg1)
}

// EnableHost mocks base method
func (m *MockInstallerAPI) EnableHost(arg0 context.Context, arg1 installer.EnableHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// EnableHost indicates an expected call of EnableHost
func (mr *MockInstallerAPIMockRecorder) EnableHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableHost", reflect.TypeOf((*MockInstallerAPI)(nil).EnableHost), arg0, arg1)
}

// GenerateClusterISO mocks base method
func (m *MockInstallerAPI) GenerateClusterISO(arg0 context.Context, arg1 installer.GenerateClusterISOParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateClusterISO", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GenerateClusterISO indicates an expected call of GenerateClusterISO
func (mr *MockInstallerAPIMockRecorder) GenerateClusterISO(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateClusterISO", reflect.TypeOf((*MockInstallerAPI)(nil).GenerateClusterISO), arg0, arg1)
}

// GetCluster mocks base method
func (m *MockInstallerAPI) GetCluster(arg0 context.Context, arg1 installer.GetClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockInstallerAPIMockRecorder) GetCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockInstallerAPI)(nil).GetCluster), arg0, arg1)
}

// GetClusterInstallConfig mocks base method
func (m *MockInstallerAPI) GetClusterInstallConfig(arg0 context.Context, arg1 installer.GetClusterInstallConfigParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInstallConfig", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetClusterInstallConfig indicates an expected call of GetClusterInstallConfig
func (mr *MockInstallerAPIMockRecorder) GetClusterInstallConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInstallConfig", reflect.TypeOf((*MockInstallerAPI)(nil).GetClusterInstallConfig), arg0, arg1)
}

// GetCredentials mocks base method
func (m *MockInstallerAPI) GetCredentials(arg0 context.Context, arg1 installer.GetCredentialsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentials", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetCredentials indicates an expected call of GetCredentials
func (mr *MockInstallerAPIMockRecorder) GetCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentials", reflect.TypeOf((*MockInstallerAPI)(nil).GetCredentials), arg0, arg1)
}

// GetFreeAddresses mocks base method
func (m *MockInstallerAPI) GetFreeAddresses(arg0 context.Context, arg1 installer.GetFreeAddressesParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFreeAddresses", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetFreeAddresses indicates an expected call of GetFreeAddresses
func (mr *MockInstallerAPIMockRecorder) GetFreeAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFreeAddresses", reflect.TypeOf((*MockInstallerAPI)(nil).GetFreeAddresses), arg0, arg1)
}

// GetHost mocks base method
func (m *MockInstallerAPI) GetHost(arg0 context.Context, arg1 installer.GetHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetHost indicates an expected call of GetHost
func (mr *MockInstallerAPIMockRecorder) GetHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHost", reflect.TypeOf((*MockInstallerAPI)(nil).GetHost), arg0, arg1)
}

// GetHostRequirements mocks base method
func (m *MockInstallerAPI) GetHostRequirements(arg0 context.Context, arg1 installer.GetHostRequirementsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostRequirements", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetHostRequirements indicates an expected call of GetHostRequirements
func (mr *MockInstallerAPIMockRecorder) GetHostRequirements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostRequirements", reflect.TypeOf((*MockInstallerAPI)(nil).GetHostRequirements), arg0, arg1)
}

// GetNextSteps mocks base method
func (m *MockInstallerAPI) GetNextSteps(arg0 context.Context, arg1 installer.GetNextStepsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSteps", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetNextSteps indicates an expected call of GetNextSteps
func (mr *MockInstallerAPIMockRecorder) GetNextSteps(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSteps", reflect.TypeOf((*MockInstallerAPI)(nil).GetNextSteps), arg0, arg1)
}

// GetPresignedForClusterFiles mocks base method
func (m *MockInstallerAPI) GetPresignedForClusterFiles(arg0 context.Context, arg1 installer.GetPresignedForClusterFilesParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPresignedForClusterFiles", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetPresignedForClusterFiles indicates an expected call of GetPresignedForClusterFiles
func (mr *MockInstallerAPIMockRecorder) GetPresignedForClusterFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPresignedForClusterFiles", reflect.TypeOf((*MockInstallerAPI)(nil).GetPresignedForClusterFiles), arg0, arg1)
}

// InstallCluster mocks base method
func (m *MockInstallerAPI) InstallCluster(arg0 context.Context, arg1 installer.InstallClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// InstallCluster indicates an expected call of InstallCluster
func (mr *MockInstallerAPIMockRecorder) InstallCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallCluster", reflect.TypeOf((*MockInstallerAPI)(nil).InstallCluster), arg0, arg1)
}

// ListClusters mocks base method
func (m *MockInstallerAPI) ListClusters(arg0 context.Context, arg1 installer.ListClustersParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// ListClusters indicates an expected call of ListClusters
func (mr *MockInstallerAPIMockRecorder) ListClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockInstallerAPI)(nil).ListClusters), arg0, arg1)
}

// ListHosts mocks base method
func (m *MockInstallerAPI) ListHosts(arg0 context.Context, arg1 installer.ListHostsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHosts", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// ListHosts indicates an expected call of ListHosts
func (mr *MockInstallerAPIMockRecorder) ListHosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHosts", reflect.TypeOf((*MockInstallerAPI)(nil).ListHosts), arg0, arg1)
}

// PostStepReply mocks base method
func (m *MockInstallerAPI) PostStepReply(arg0 context.Context, arg1 installer.PostStepReplyParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostStepReply", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// PostStepReply indicates an expected call of PostStepReply
func (mr *MockInstallerAPIMockRecorder) PostStepReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostStepReply", reflect.TypeOf((*MockInstallerAPI)(nil).PostStepReply), arg0, arg1)
}

// RegisterCluster mocks base method
func (m *MockInstallerAPI) RegisterCluster(arg0 context.Context, arg1 installer.RegisterClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// RegisterCluster indicates an expected call of RegisterCluster
func (mr *MockInstallerAPIMockRecorder) RegisterCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCluster", reflect.TypeOf((*MockInstallerAPI)(nil).RegisterCluster), arg0, arg1)
}

// RegisterHost mocks base method
func (m *MockInstallerAPI) RegisterHost(arg0 context.Context, arg1 installer.RegisterHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// RegisterHost indicates an expected call of RegisterHost
func (mr *MockInstallerAPIMockRecorder) RegisterHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterHost", reflect.TypeOf((*MockInstallerAPI)(nil).RegisterHost), arg0, arg1)
}

// ResetCluster mocks base method
func (m *MockInstallerAPI) ResetCluster(arg0 context.Context, arg1 installer.ResetClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// ResetCluster indicates an expected call of ResetCluster
func (mr *MockInstallerAPIMockRecorder) ResetCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetCluster", reflect.TypeOf((*MockInstallerAPI)(nil).ResetCluster), arg0, arg1)
}

// UpdateCluster mocks base method
func (m *MockInstallerAPI) UpdateCluster(arg0 context.Context, arg1 installer.UpdateClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// UpdateCluster indicates an expected call of UpdateCluster
func (mr *MockInstallerAPIMockRecorder) UpdateCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockInstallerAPI)(nil).UpdateCluster), arg0, arg1)
}

// UpdateClusterInstallConfig mocks base method
func (m *MockInstallerAPI) UpdateClusterInstallConfig(arg0 context.Context, arg1 installer.UpdateClusterInstallConfigParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClusterInstallConfig", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// UpdateClusterInstallConfig indicates an expected call of UpdateClusterInstallConfig
func (mr *MockInstallerAPIMockRecorder) UpdateClusterInstallConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterInstallConfig", reflect.TypeOf((*MockInstallerAPI)(nil).UpdateClusterInstallConfig), arg0, arg1)
}

// UpdateHostInstallProgress mocks base method
func (m *MockInstallerAPI) UpdateHostInstallProgress(arg0 context.Context, arg1 installer.UpdateHostInstallProgressParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostInstallProgress", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// UpdateHostInstallProgress indicates an expected call of UpdateHostInstallProgress
func (mr *MockInstallerAPIMockRecorder) UpdateHostInstallProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostInstallProgress", reflect.TypeOf((*MockInstallerAPI)(nil).UpdateHostInstallProgress), arg0, arg1)
}

// UploadClusterIngressCert mocks base method
func (m *MockInstallerAPI) UploadClusterIngressCert(arg0 context.Context, arg1 installer.UploadClusterIngressCertParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadClusterIngressCert", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// UploadClusterIngressCert indicates an expected call of UploadClusterIngressCert
func (mr *MockInstallerAPIMockRecorder) UploadClusterIngressCert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadClusterIngressCert", reflect.TypeOf((*MockInstallerAPI)(nil).UploadClusterIngressCert), arg0, arg1)
}

// UploadHostLogs mocks base method
func (m *MockInstallerAPI) UploadHostLogs(arg0 context.Context, arg1 installer.UploadHostLogsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadHostLogs", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// UploadHostLogs indicates an expected call of UploadHostLogs
func (mr *MockInstallerAPIMockRecorder) UploadHostLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadHostLogs", reflect.TypeOf((*MockInstallerAPI)(nil).UploadHostLogs), arg0, arg1)
}
