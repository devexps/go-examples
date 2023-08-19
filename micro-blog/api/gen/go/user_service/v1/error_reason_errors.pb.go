// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/devexps/go-micro/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the micro package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 101
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(101, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}