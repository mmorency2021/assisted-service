// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/restapi (interfaces: InstallerAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	middleware "github.com/go-openapi/runtime/middleware"
	gomock "github.com/golang/mock/gomock"
	installer "github.com/openshift/assisted-service/restapi/operations/installer"
)

// MockInstallerAPI is a mock of InstallerAPI interface.
type MockInstallerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerAPIMockRecorder
}

// MockInstallerAPIMockRecorder is the mock recorder for MockInstallerAPI.
type MockInstallerAPIMockRecorder struct {
	mock *MockInstallerAPI
}

// NewMockInstallerAPI creates a new mock instance.
func NewMockInstallerAPI(ctrl *gomock.Controller) *MockInstallerAPI {
	mock := &MockInstallerAPI{ctrl: ctrl}
	mock.recorder = &MockInstallerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstallerAPI) EXPECT() *MockInstallerAPIMockRecorder {
	return m.recorder
}

// BindHost mocks base method.
func (m *MockInstallerAPI) BindHost(arg0 context.Context, arg1 installer.BindHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// BindHost indicates an expected call of BindHost.
func (mr *MockInstallerAPIMockRecorder) BindHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindHost", reflect.TypeOf((*MockInstallerAPI)(nil).BindHost), arg0, arg1)
}

// DeregisterInfraEnv mocks base method.
func (m *MockInstallerAPI) DeregisterInfraEnv(arg0 context.Context, arg1 installer.DeregisterInfraEnvParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterInfraEnv", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DeregisterInfraEnv indicates an expected call of DeregisterInfraEnv.
func (mr *MockInstallerAPIMockRecorder) DeregisterInfraEnv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterInfraEnv", reflect.TypeOf((*MockInstallerAPI)(nil).DeregisterInfraEnv), arg0, arg1)
}

// DownloadMinimalInitrd mocks base method.
func (m *MockInstallerAPI) DownloadMinimalInitrd(arg0 context.Context, arg1 installer.DownloadMinimalInitrdParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadMinimalInitrd", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DownloadMinimalInitrd indicates an expected call of DownloadMinimalInitrd.
func (mr *MockInstallerAPIMockRecorder) DownloadMinimalInitrd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadMinimalInitrd", reflect.TypeOf((*MockInstallerAPI)(nil).DownloadMinimalInitrd), arg0, arg1)
}

// GetClusterSupportedPlatforms mocks base method.
func (m *MockInstallerAPI) GetClusterSupportedPlatforms(arg0 context.Context, arg1 installer.GetClusterSupportedPlatformsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterSupportedPlatforms", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetClusterSupportedPlatforms indicates an expected call of GetClusterSupportedPlatforms.
func (mr *MockInstallerAPIMockRecorder) GetClusterSupportedPlatforms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterSupportedPlatforms", reflect.TypeOf((*MockInstallerAPI)(nil).GetClusterSupportedPlatforms), arg0, arg1)
}

// GetInfraEnv mocks base method.
func (m *MockInstallerAPI) GetInfraEnv(arg0 context.Context, arg1 installer.GetInfraEnvParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfraEnv", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetInfraEnv indicates an expected call of GetInfraEnv.
func (mr *MockInstallerAPIMockRecorder) GetInfraEnv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfraEnv", reflect.TypeOf((*MockInstallerAPI)(nil).GetInfraEnv), arg0, arg1)
}

// GetInfraEnvDownloadURL mocks base method.
func (m *MockInstallerAPI) GetInfraEnvDownloadURL(arg0 context.Context, arg1 installer.GetInfraEnvDownloadURLParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfraEnvDownloadURL", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetInfraEnvDownloadURL indicates an expected call of GetInfraEnvDownloadURL.
func (mr *MockInstallerAPIMockRecorder) GetInfraEnvDownloadURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfraEnvDownloadURL", reflect.TypeOf((*MockInstallerAPI)(nil).GetInfraEnvDownloadURL), arg0, arg1)
}

// GetInfraEnvPresignedFileURL mocks base method.
func (m *MockInstallerAPI) GetInfraEnvPresignedFileURL(arg0 context.Context, arg1 installer.GetInfraEnvPresignedFileURLParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfraEnvPresignedFileURL", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// GetInfraEnvPresignedFileURL indicates an expected call of GetInfraEnvPresignedFileURL.
func (mr *MockInstallerAPIMockRecorder) GetInfraEnvPresignedFileURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfraEnvPresignedFileURL", reflect.TypeOf((*MockInstallerAPI)(nil).GetInfraEnvPresignedFileURL), arg0, arg1)
}

// ListInfraEnvs mocks base method.
func (m *MockInstallerAPI) ListInfraEnvs(arg0 context.Context, arg1 installer.ListInfraEnvsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInfraEnvs", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// ListInfraEnvs indicates an expected call of ListInfraEnvs.
func (mr *MockInstallerAPIMockRecorder) ListInfraEnvs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInfraEnvs", reflect.TypeOf((*MockInstallerAPI)(nil).ListInfraEnvs), arg0, arg1)
}

// RegenerateInfraEnvSigningKey mocks base method.
func (m *MockInstallerAPI) RegenerateInfraEnvSigningKey(arg0 context.Context, arg1 installer.RegenerateInfraEnvSigningKeyParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegenerateInfraEnvSigningKey", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// RegenerateInfraEnvSigningKey indicates an expected call of RegenerateInfraEnvSigningKey.
func (mr *MockInstallerAPIMockRecorder) RegenerateInfraEnvSigningKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegenerateInfraEnvSigningKey", reflect.TypeOf((*MockInstallerAPI)(nil).RegenerateInfraEnvSigningKey), arg0, arg1)
}

// RegisterInfraEnv mocks base method.
func (m *MockInstallerAPI) RegisterInfraEnv(arg0 context.Context, arg1 installer.RegisterInfraEnvParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInfraEnv", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// RegisterInfraEnv indicates an expected call of RegisterInfraEnv.
func (mr *MockInstallerAPIMockRecorder) RegisterInfraEnv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInfraEnv", reflect.TypeOf((*MockInstallerAPI)(nil).RegisterInfraEnv), arg0, arg1)
}

// TransformClusterToDay2 mocks base method.
func (m *MockInstallerAPI) TransformClusterToDay2(arg0 context.Context, arg1 installer.TransformClusterToDay2Params) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransformClusterToDay2", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// TransformClusterToDay2 indicates an expected call of TransformClusterToDay2.
func (mr *MockInstallerAPIMockRecorder) TransformClusterToDay2(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransformClusterToDay2", reflect.TypeOf((*MockInstallerAPI)(nil).TransformClusterToDay2), arg0, arg1)
}

// UnbindHost mocks base method.
func (m *MockInstallerAPI) UnbindHost(arg0 context.Context, arg1 installer.UnbindHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbindHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// UnbindHost indicates an expected call of UnbindHost.
func (mr *MockInstallerAPIMockRecorder) UnbindHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbindHost", reflect.TypeOf((*MockInstallerAPI)(nil).UnbindHost), arg0, arg1)
}

// UpdateInfraEnv mocks base method.
func (m *MockInstallerAPI) UpdateInfraEnv(arg0 context.Context, arg1 installer.UpdateInfraEnvParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfraEnv", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// UpdateInfraEnv indicates an expected call of UpdateInfraEnv.
func (mr *MockInstallerAPIMockRecorder) UpdateInfraEnv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfraEnv", reflect.TypeOf((*MockInstallerAPI)(nil).UpdateInfraEnv), arg0, arg1)
}

// V2CancelInstallation mocks base method.
func (m *MockInstallerAPI) V2CancelInstallation(arg0 context.Context, arg1 installer.V2CancelInstallationParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2CancelInstallation", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2CancelInstallation indicates an expected call of V2CancelInstallation.
func (mr *MockInstallerAPIMockRecorder) V2CancelInstallation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2CancelInstallation", reflect.TypeOf((*MockInstallerAPI)(nil).V2CancelInstallation), arg0, arg1)
}

// V2CompleteInstallation mocks base method.
func (m *MockInstallerAPI) V2CompleteInstallation(arg0 context.Context, arg1 installer.V2CompleteInstallationParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2CompleteInstallation", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2CompleteInstallation indicates an expected call of V2CompleteInstallation.
func (mr *MockInstallerAPIMockRecorder) V2CompleteInstallation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2CompleteInstallation", reflect.TypeOf((*MockInstallerAPI)(nil).V2CompleteInstallation), arg0, arg1)
}

// V2DeregisterCluster mocks base method.
func (m *MockInstallerAPI) V2DeregisterCluster(arg0 context.Context, arg1 installer.V2DeregisterClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DeregisterCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2DeregisterCluster indicates an expected call of V2DeregisterCluster.
func (mr *MockInstallerAPIMockRecorder) V2DeregisterCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DeregisterCluster", reflect.TypeOf((*MockInstallerAPI)(nil).V2DeregisterCluster), arg0, arg1)
}

// V2DeregisterHost mocks base method.
func (m *MockInstallerAPI) V2DeregisterHost(arg0 context.Context, arg1 installer.V2DeregisterHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DeregisterHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2DeregisterHost indicates an expected call of V2DeregisterHost.
func (mr *MockInstallerAPIMockRecorder) V2DeregisterHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DeregisterHost", reflect.TypeOf((*MockInstallerAPI)(nil).V2DeregisterHost), arg0, arg1)
}

// V2DownloadClusterCredentials mocks base method.
func (m *MockInstallerAPI) V2DownloadClusterCredentials(arg0 context.Context, arg1 installer.V2DownloadClusterCredentialsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DownloadClusterCredentials", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2DownloadClusterCredentials indicates an expected call of V2DownloadClusterCredentials.
func (mr *MockInstallerAPIMockRecorder) V2DownloadClusterCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DownloadClusterCredentials", reflect.TypeOf((*MockInstallerAPI)(nil).V2DownloadClusterCredentials), arg0, arg1)
}

// V2DownloadClusterFiles mocks base method.
func (m *MockInstallerAPI) V2DownloadClusterFiles(arg0 context.Context, arg1 installer.V2DownloadClusterFilesParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DownloadClusterFiles", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2DownloadClusterFiles indicates an expected call of V2DownloadClusterFiles.
func (mr *MockInstallerAPIMockRecorder) V2DownloadClusterFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DownloadClusterFiles", reflect.TypeOf((*MockInstallerAPI)(nil).V2DownloadClusterFiles), arg0, arg1)
}

// V2DownloadClusterLogs mocks base method.
func (m *MockInstallerAPI) V2DownloadClusterLogs(arg0 context.Context, arg1 installer.V2DownloadClusterLogsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DownloadClusterLogs", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2DownloadClusterLogs indicates an expected call of V2DownloadClusterLogs.
func (mr *MockInstallerAPIMockRecorder) V2DownloadClusterLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DownloadClusterLogs", reflect.TypeOf((*MockInstallerAPI)(nil).V2DownloadClusterLogs), arg0, arg1)
}

// V2DownloadHostIgnition mocks base method.
func (m *MockInstallerAPI) V2DownloadHostIgnition(arg0 context.Context, arg1 installer.V2DownloadHostIgnitionParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DownloadHostIgnition", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2DownloadHostIgnition indicates an expected call of V2DownloadHostIgnition.
func (mr *MockInstallerAPIMockRecorder) V2DownloadHostIgnition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DownloadHostIgnition", reflect.TypeOf((*MockInstallerAPI)(nil).V2DownloadHostIgnition), arg0, arg1)
}

// V2DownloadInfraEnvFiles mocks base method.
func (m *MockInstallerAPI) V2DownloadInfraEnvFiles(arg0 context.Context, arg1 installer.V2DownloadInfraEnvFilesParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DownloadInfraEnvFiles", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2DownloadInfraEnvFiles indicates an expected call of V2DownloadInfraEnvFiles.
func (mr *MockInstallerAPIMockRecorder) V2DownloadInfraEnvFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DownloadInfraEnvFiles", reflect.TypeOf((*MockInstallerAPI)(nil).V2DownloadInfraEnvFiles), arg0, arg1)
}

// V2GetCluster mocks base method.
func (m *MockInstallerAPI) V2GetCluster(arg0 context.Context, arg1 installer.V2GetClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetCluster indicates an expected call of V2GetCluster.
func (mr *MockInstallerAPIMockRecorder) V2GetCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetCluster", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetCluster), arg0, arg1)
}

// V2GetClusterDefaultConfig mocks base method.
func (m *MockInstallerAPI) V2GetClusterDefaultConfig(arg0 context.Context, arg1 installer.V2GetClusterDefaultConfigParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetClusterDefaultConfig", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetClusterDefaultConfig indicates an expected call of V2GetClusterDefaultConfig.
func (mr *MockInstallerAPIMockRecorder) V2GetClusterDefaultConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetClusterDefaultConfig", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetClusterDefaultConfig), arg0, arg1)
}

// V2GetClusterInstallConfig mocks base method.
func (m *MockInstallerAPI) V2GetClusterInstallConfig(arg0 context.Context, arg1 installer.V2GetClusterInstallConfigParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetClusterInstallConfig", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetClusterInstallConfig indicates an expected call of V2GetClusterInstallConfig.
func (mr *MockInstallerAPIMockRecorder) V2GetClusterInstallConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetClusterInstallConfig", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetClusterInstallConfig), arg0, arg1)
}

// V2GetCredentials mocks base method.
func (m *MockInstallerAPI) V2GetCredentials(arg0 context.Context, arg1 installer.V2GetCredentialsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetCredentials", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetCredentials indicates an expected call of V2GetCredentials.
func (mr *MockInstallerAPIMockRecorder) V2GetCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetCredentials", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetCredentials), arg0, arg1)
}

// V2GetHost mocks base method.
func (m *MockInstallerAPI) V2GetHost(arg0 context.Context, arg1 installer.V2GetHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetHost indicates an expected call of V2GetHost.
func (mr *MockInstallerAPIMockRecorder) V2GetHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetHost", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetHost), arg0, arg1)
}

// V2GetHostIgnition mocks base method.
func (m *MockInstallerAPI) V2GetHostIgnition(arg0 context.Context, arg1 installer.V2GetHostIgnitionParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetHostIgnition", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetHostIgnition indicates an expected call of V2GetHostIgnition.
func (mr *MockInstallerAPIMockRecorder) V2GetHostIgnition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetHostIgnition", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetHostIgnition), arg0, arg1)
}

// V2GetNextSteps mocks base method.
func (m *MockInstallerAPI) V2GetNextSteps(arg0 context.Context, arg1 installer.V2GetNextStepsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetNextSteps", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetNextSteps indicates an expected call of V2GetNextSteps.
func (mr *MockInstallerAPIMockRecorder) V2GetNextSteps(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetNextSteps", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetNextSteps), arg0, arg1)
}

// V2GetPreflightRequirements mocks base method.
func (m *MockInstallerAPI) V2GetPreflightRequirements(arg0 context.Context, arg1 installer.V2GetPreflightRequirementsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetPreflightRequirements", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetPreflightRequirements indicates an expected call of V2GetPreflightRequirements.
func (mr *MockInstallerAPIMockRecorder) V2GetPreflightRequirements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetPreflightRequirements", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetPreflightRequirements), arg0, arg1)
}

// V2GetPresignedForClusterCredentials mocks base method.
func (m *MockInstallerAPI) V2GetPresignedForClusterCredentials(arg0 context.Context, arg1 installer.V2GetPresignedForClusterCredentialsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetPresignedForClusterCredentials", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetPresignedForClusterCredentials indicates an expected call of V2GetPresignedForClusterCredentials.
func (mr *MockInstallerAPIMockRecorder) V2GetPresignedForClusterCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetPresignedForClusterCredentials", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetPresignedForClusterCredentials), arg0, arg1)
}

// V2GetPresignedForClusterFiles mocks base method.
func (m *MockInstallerAPI) V2GetPresignedForClusterFiles(arg0 context.Context, arg1 installer.V2GetPresignedForClusterFilesParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2GetPresignedForClusterFiles", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2GetPresignedForClusterFiles indicates an expected call of V2GetPresignedForClusterFiles.
func (mr *MockInstallerAPIMockRecorder) V2GetPresignedForClusterFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2GetPresignedForClusterFiles", reflect.TypeOf((*MockInstallerAPI)(nil).V2GetPresignedForClusterFiles), arg0, arg1)
}

// V2ImportCluster mocks base method.
func (m *MockInstallerAPI) V2ImportCluster(arg0 context.Context, arg1 installer.V2ImportClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2ImportCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2ImportCluster indicates an expected call of V2ImportCluster.
func (mr *MockInstallerAPIMockRecorder) V2ImportCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2ImportCluster", reflect.TypeOf((*MockInstallerAPI)(nil).V2ImportCluster), arg0, arg1)
}

// V2InstallCluster mocks base method.
func (m *MockInstallerAPI) V2InstallCluster(arg0 context.Context, arg1 installer.V2InstallClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2InstallCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2InstallCluster indicates an expected call of V2InstallCluster.
func (mr *MockInstallerAPIMockRecorder) V2InstallCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2InstallCluster", reflect.TypeOf((*MockInstallerAPI)(nil).V2InstallCluster), arg0, arg1)
}

// V2InstallHost mocks base method.
func (m *MockInstallerAPI) V2InstallHost(arg0 context.Context, arg1 installer.V2InstallHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2InstallHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2InstallHost indicates an expected call of V2InstallHost.
func (mr *MockInstallerAPIMockRecorder) V2InstallHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2InstallHost", reflect.TypeOf((*MockInstallerAPI)(nil).V2InstallHost), arg0, arg1)
}

// V2ListClusters mocks base method.
func (m *MockInstallerAPI) V2ListClusters(arg0 context.Context, arg1 installer.V2ListClustersParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2ListClusters", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2ListClusters indicates an expected call of V2ListClusters.
func (mr *MockInstallerAPIMockRecorder) V2ListClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2ListClusters", reflect.TypeOf((*MockInstallerAPI)(nil).V2ListClusters), arg0, arg1)
}

// V2ListFeatureSupportLevels mocks base method.
func (m *MockInstallerAPI) V2ListFeatureSupportLevels(arg0 context.Context, arg1 installer.V2ListFeatureSupportLevelsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2ListFeatureSupportLevels", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2ListFeatureSupportLevels indicates an expected call of V2ListFeatureSupportLevels.
func (mr *MockInstallerAPIMockRecorder) V2ListFeatureSupportLevels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2ListFeatureSupportLevels", reflect.TypeOf((*MockInstallerAPI)(nil).V2ListFeatureSupportLevels), arg0, arg1)
}

// V2ListHosts mocks base method.
func (m *MockInstallerAPI) V2ListHosts(arg0 context.Context, arg1 installer.V2ListHostsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2ListHosts", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2ListHosts indicates an expected call of V2ListHosts.
func (mr *MockInstallerAPIMockRecorder) V2ListHosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2ListHosts", reflect.TypeOf((*MockInstallerAPI)(nil).V2ListHosts), arg0, arg1)
}

// V2PostStepReply mocks base method.
func (m *MockInstallerAPI) V2PostStepReply(arg0 context.Context, arg1 installer.V2PostStepReplyParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2PostStepReply", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2PostStepReply indicates an expected call of V2PostStepReply.
func (mr *MockInstallerAPIMockRecorder) V2PostStepReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2PostStepReply", reflect.TypeOf((*MockInstallerAPI)(nil).V2PostStepReply), arg0, arg1)
}

// V2RegisterCluster mocks base method.
func (m *MockInstallerAPI) V2RegisterCluster(arg0 context.Context, arg1 installer.V2RegisterClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2RegisterCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2RegisterCluster indicates an expected call of V2RegisterCluster.
func (mr *MockInstallerAPIMockRecorder) V2RegisterCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2RegisterCluster", reflect.TypeOf((*MockInstallerAPI)(nil).V2RegisterCluster), arg0, arg1)
}

// V2RegisterHost mocks base method.
func (m *MockInstallerAPI) V2RegisterHost(arg0 context.Context, arg1 installer.V2RegisterHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2RegisterHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2RegisterHost indicates an expected call of V2RegisterHost.
func (mr *MockInstallerAPIMockRecorder) V2RegisterHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2RegisterHost", reflect.TypeOf((*MockInstallerAPI)(nil).V2RegisterHost), arg0, arg1)
}

// V2ResetCluster mocks base method.
func (m *MockInstallerAPI) V2ResetCluster(arg0 context.Context, arg1 installer.V2ResetClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2ResetCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2ResetCluster indicates an expected call of V2ResetCluster.
func (mr *MockInstallerAPIMockRecorder) V2ResetCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2ResetCluster", reflect.TypeOf((*MockInstallerAPI)(nil).V2ResetCluster), arg0, arg1)
}

// V2ResetHost mocks base method.
func (m *MockInstallerAPI) V2ResetHost(arg0 context.Context, arg1 installer.V2ResetHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2ResetHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2ResetHost indicates an expected call of V2ResetHost.
func (mr *MockInstallerAPIMockRecorder) V2ResetHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2ResetHost", reflect.TypeOf((*MockInstallerAPI)(nil).V2ResetHost), arg0, arg1)
}

// V2ResetHostValidation mocks base method.
func (m *MockInstallerAPI) V2ResetHostValidation(arg0 context.Context, arg1 installer.V2ResetHostValidationParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2ResetHostValidation", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2ResetHostValidation indicates an expected call of V2ResetHostValidation.
func (mr *MockInstallerAPIMockRecorder) V2ResetHostValidation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2ResetHostValidation", reflect.TypeOf((*MockInstallerAPI)(nil).V2ResetHostValidation), arg0, arg1)
}

// V2UpdateCluster mocks base method.
func (m *MockInstallerAPI) V2UpdateCluster(arg0 context.Context, arg1 installer.V2UpdateClusterParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateCluster", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UpdateCluster indicates an expected call of V2UpdateCluster.
func (mr *MockInstallerAPIMockRecorder) V2UpdateCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateCluster", reflect.TypeOf((*MockInstallerAPI)(nil).V2UpdateCluster), arg0, arg1)
}

// V2UpdateClusterInstallConfig mocks base method.
func (m *MockInstallerAPI) V2UpdateClusterInstallConfig(arg0 context.Context, arg1 installer.V2UpdateClusterInstallConfigParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateClusterInstallConfig", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UpdateClusterInstallConfig indicates an expected call of V2UpdateClusterInstallConfig.
func (mr *MockInstallerAPIMockRecorder) V2UpdateClusterInstallConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateClusterInstallConfig", reflect.TypeOf((*MockInstallerAPI)(nil).V2UpdateClusterInstallConfig), arg0, arg1)
}

// V2UpdateClusterLogsProgress mocks base method.
func (m *MockInstallerAPI) V2UpdateClusterLogsProgress(arg0 context.Context, arg1 installer.V2UpdateClusterLogsProgressParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateClusterLogsProgress", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UpdateClusterLogsProgress indicates an expected call of V2UpdateClusterLogsProgress.
func (mr *MockInstallerAPIMockRecorder) V2UpdateClusterLogsProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateClusterLogsProgress", reflect.TypeOf((*MockInstallerAPI)(nil).V2UpdateClusterLogsProgress), arg0, arg1)
}

// V2UpdateHost mocks base method.
func (m *MockInstallerAPI) V2UpdateHost(arg0 context.Context, arg1 installer.V2UpdateHostParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateHost", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UpdateHost indicates an expected call of V2UpdateHost.
func (mr *MockInstallerAPIMockRecorder) V2UpdateHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateHost", reflect.TypeOf((*MockInstallerAPI)(nil).V2UpdateHost), arg0, arg1)
}

// V2UpdateHostIgnition mocks base method.
func (m *MockInstallerAPI) V2UpdateHostIgnition(arg0 context.Context, arg1 installer.V2UpdateHostIgnitionParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateHostIgnition", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UpdateHostIgnition indicates an expected call of V2UpdateHostIgnition.
func (mr *MockInstallerAPIMockRecorder) V2UpdateHostIgnition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateHostIgnition", reflect.TypeOf((*MockInstallerAPI)(nil).V2UpdateHostIgnition), arg0, arg1)
}

// V2UpdateHostInstallProgress mocks base method.
func (m *MockInstallerAPI) V2UpdateHostInstallProgress(arg0 context.Context, arg1 installer.V2UpdateHostInstallProgressParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateHostInstallProgress", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UpdateHostInstallProgress indicates an expected call of V2UpdateHostInstallProgress.
func (mr *MockInstallerAPIMockRecorder) V2UpdateHostInstallProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateHostInstallProgress", reflect.TypeOf((*MockInstallerAPI)(nil).V2UpdateHostInstallProgress), arg0, arg1)
}

// V2UpdateHostInstallerArgs mocks base method.
func (m *MockInstallerAPI) V2UpdateHostInstallerArgs(arg0 context.Context, arg1 installer.V2UpdateHostInstallerArgsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateHostInstallerArgs", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UpdateHostInstallerArgs indicates an expected call of V2UpdateHostInstallerArgs.
func (mr *MockInstallerAPIMockRecorder) V2UpdateHostInstallerArgs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateHostInstallerArgs", reflect.TypeOf((*MockInstallerAPI)(nil).V2UpdateHostInstallerArgs), arg0, arg1)
}

// V2UpdateHostLogsProgress mocks base method.
func (m *MockInstallerAPI) V2UpdateHostLogsProgress(arg0 context.Context, arg1 installer.V2UpdateHostLogsProgressParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateHostLogsProgress", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UpdateHostLogsProgress indicates an expected call of V2UpdateHostLogsProgress.
func (mr *MockInstallerAPIMockRecorder) V2UpdateHostLogsProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateHostLogsProgress", reflect.TypeOf((*MockInstallerAPI)(nil).V2UpdateHostLogsProgress), arg0, arg1)
}

// V2UploadClusterIngressCert mocks base method.
func (m *MockInstallerAPI) V2UploadClusterIngressCert(arg0 context.Context, arg1 installer.V2UploadClusterIngressCertParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UploadClusterIngressCert", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UploadClusterIngressCert indicates an expected call of V2UploadClusterIngressCert.
func (mr *MockInstallerAPIMockRecorder) V2UploadClusterIngressCert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UploadClusterIngressCert", reflect.TypeOf((*MockInstallerAPI)(nil).V2UploadClusterIngressCert), arg0, arg1)
}

// V2UploadLogs mocks base method.
func (m *MockInstallerAPI) V2UploadLogs(arg0 context.Context, arg1 installer.V2UploadLogsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UploadLogs", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2UploadLogs indicates an expected call of V2UploadLogs.
func (mr *MockInstallerAPIMockRecorder) V2UploadLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UploadLogs", reflect.TypeOf((*MockInstallerAPI)(nil).V2UploadLogs), arg0, arg1)
}
