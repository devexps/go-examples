# PII Masking example

## Install

- [protobuf version 3](https://github.com/google/protobuf)
- go support for protobuf: `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`
- go support to inject custom tags to generated protobuf
  files: `go install github.com/favadi/protoc-go-inject-tag@latest`
- go lib to support masking PII sensitive data:
  [https://github.com/devexps/go-pkg/mask](https://github.com/devexps/go-pkg/tree/main/mask)

## Example

### Define protobuf file

```protobuf
// file: test.proto
syntax = "proto3";

package test_srv.v1;

option go_package = "github.com/devexps/go-examples/pii-masking/test-srv/v1;v1";

message ContactInfo {
  // @gotags: mask:"name"
  string name = 1;

  // title
  string title = 2;

  // @gotags: mask:"email"
  string email = 3;

  // @gotags: mask:"mobile"
  string mobile = 4;

  // @gotags: mask:"telephone"
  string telephone = 5;
}

message PresenterInfo {
  // @gotags: mask:"id"
  string id = 1;

  // @gotags: mask:"name"
  string name = 2;

  // title
  string title = 3;

  // @gotags: mask:"address"
  string address = 4;
}

message IntegrationInfo {
  // name
  string name = 1;

  // @gotags: mask:"URL"
  string hook_url = 2;

  // @gotags: mask:"secret"
  string token = 3;
}

message LoginInfo {
  // username
  string username = 1;

  // @gotags: mask:"password"
  string password = 2;
}

message UserInfo {
  // contact_info
  ContactInfo contact_info = 1;

  // presenter_info
  PresenterInfo presenter_info = 2;

  // integration_info list
  repeated IntegrationInfo integration_infos = 3;

  // login_info
  LoginInfo login_info = 4;

  // @gotags: mask:"custom1"
  string business_license_number = 5;

  // @gotags: mask:"custom2"
  string business_license_no = 6;
}
```

### Generate protobuf file

```shell
#!/bin/bash

set -eu

protoc --proto_path=. --go_out=paths=source_relative:. test.proto
protoc-go-inject-tag -input="*.pb.go"
```

The custom tags will be injected to `test.pb.go`, for example:

```go
type ContactInfo struct {
state         protoimpl.MessageState
sizeCache     protoimpl.SizeCache
unknownFields protoimpl.UnknownFields

// @gotags: mask:"name"
Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" mask:"name"`
// title
Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
// @gotags: mask:"email"
Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" mask:"email"`
// @gotags: mask:"mobile"
Mobile string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty" mask:"mobile"`
// @gotags: mask:"telephone"
Telephone string `protobuf:"bytes,5,opt,name=telephone,proto3" json:"telephone,omitempty" mask:"telephone"`
}
```

### Example Go file

```go
package main

import (
	"fmt"

	v1 "github.com/devexps/go-examples/pii-masking/pb"

	"github.com/devexps/go-pkg/v2/mask"
	"github.com/devexps/go-pkg/v2/mask/masker"
)

func main() {
	record := &v1.UserInfo{
		ContactInfo: &v1.ContactInfo{
			Name:      "Golden Name",
			Title:     "A Poor Developer",
			Email:     "poor.developer@gmail.com",
			Mobile:    "0923456789",
			Telephone: "0287634512",
		},
		PresenterInfo: &v1.PresenterInfo{
			Id:      "1209000231234",
			Name:    "Silver Name",
			Title:   "A Rich Developer",
			Address: "123 St4, xyz street, D7",
		},
		IntegrationInfos: []*v1.IntegrationInfo{{
			Name:    "In-App",
			HookUrl: "http://app:mysecretpassword@localhost:1234/uri",
			Token:   "123ABC456",
		}, {
			Name:    "App-to-App",
			HookUrl: "http://app2:mysecretpassword@localhost:1234/uri2",
			Token:   "ABC123DEF",
		}},
		LoginInfo: &v1.LoginInfo{
			Username: "demo",
			Password: "123456@.123",
		},
		BusinessLicenseNumber: "001-23245-33214",
		BusinessLicenseNo:     "BL:7701",
	}

	useCustomMasker()

	maskTool := mask.New(mask.TagsFilter())
	filteredRecord := maskTool.Apply(record)

	fmt.Println(filteredRecord)
}

func useCustomMasker() {
	mask.SetMasker(masker.NewMasker(
		masker.WithMarkTypes("custom1", "custom2"),
		masker.WithStringFunc(customMaskerString),
	))
}

func customMaskerString(t masker.MType, s string) (bool, string) {
	if t == "custom1" {
		return true, "---@@---"
	}
	return false, ""
}
```

The good result is gonna be:

```json
{
  "contact_info": {
    "name": "G**den N**e",
    "title": "A Poor Developer",
    "email": "poo****veloper@gmail.com",
    "mobile": "0923***789",
    "telephone": "(02)8763-****"
  },
  "presenter_info": {
    "id": "120900****",
    "name": "S**ver N**e",
    "title": "A Rich Developer",
    "address": "123 St******"
  },
  "integration_infos": [
    {
      "name": "In-App",
      "hook_url": "http://app:xxxxx@localhost:1234/uri",
      "token": "[filtered]"
    },
    {
      "name": "App-to-App",
      "hook_url": "http://app2:xxxxx@localhost:1234/uri2",
      "token": "[filtered]"
    }
  ],
  "login_info": {
    "username": "demo",
    "password": "************"
  },
  "business_license_number": "---@@---",
  "business_license_no": "[filtered]"
}
```
