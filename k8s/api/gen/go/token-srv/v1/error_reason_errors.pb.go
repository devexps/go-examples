// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/devexps/go-micro/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the micro package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsCreateTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CREATE_TOKEN_FAILED.String() && e.Code == 101
}

func ErrorCreateTokenFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(101, ErrorReason_CREATE_TOKEN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSaveTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SAVE_TOKEN_FAILED.String() && e.Code == 102
}

func ErrorSaveTokenFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(102, ErrorReason_SAVE_TOKEN_FAILED.String(), fmt.Sprintf(format, args...))
}