// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/biblia/pkg/biblehub (interfaces: Downloader)
//
// Generated by this command:
//
//	mockgen -destination mocks_test.go -package biblehub . Downloader
//

// Package biblehub is a generated GoMock package.
package biblehub

import (
	gomock "go.uber.org/mock/gomock"
)

// MockDownloader is a mock of Downloader interface.
type MockDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockDownloaderMockRecorder
	isgomock struct{}
}

// MockDownloaderMockRecorder is the mock recorder for MockDownloader.
type MockDownloaderMockRecorder struct {
	mock *MockDownloader
}

// NewMockDownloader creates a new mock instance.
func NewMockDownloader(ctrl *gomock.Controller) *MockDownloader {
	mock := &MockDownloader{ctrl: ctrl}
	mock.recorder = &MockDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloader) EXPECT() *MockDownloaderMockRecorder {
	return m.recorder
}