# PII Masking example

Protect sensitive customer information by limiting access and visibility.

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

  // @gotags: mask:"url"
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

## Filter Sensitive Data

### By specified field

#### Default

```go
type myRecord struct {
    ID    string
    Phone string
}
record := myRecord{
    ID:    "userId",
    Phone: "090-1234-5678",
}

maskTool := mask.New(mask.FieldFilter("Phone"))
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId","Phone":"[filtered]"}
```

#### Custom

```go
type myRecord struct {
    ID         string
    Phone      string
    Url        string
    Email      string
    Name       string
    Address    string
    CreditCard string
}
record := myRecord{
    ID:         "userId",
    Phone:      "090-1234-5678",
    Url:        "http://admin:mysecretpassword@localhost:1234/uri",
    Email:      "dummy@dummy.com",
    Name:       "John Doe",
    Address:    "1 AB Road, Paradise",
    CreditCard: "4444-4444-4444-4444",
}

maskTool := mask.New(
    mask.FieldFilter("Phone", masker.MMobile),
    mask.FieldFilter("Email", masker.MEmail),
    mask.FieldFilter("Url", masker.MURL),
    mask.FieldFilter("Name", masker.MName),
    mask.FieldFilter("ID", masker.MID),
    mask.FieldFilter("Address", masker.MAddress),
    mask.FieldFilter("CreditCard", masker.MCreditCard),
)
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId****","Phone":"090-***4-5678","Url":"http://admin:xxxxx@localhost:1234/uri","Email":"dum****@dummy.com","Name":"J**n D**e","Address":"1 AB R******","CreditCard":"4444-4******44-4444"}
```

### By specified field-prefix

#### Default

```go
type myRecord struct {
    ID          string
    SecurePhone string
}
record := myRecord{
    ID:          "userId",
    SecurePhone: "090-1234-5678",
}

maskTool := mask.New(mask.FieldPrefixFilter("Secure"))
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId","SecurePhone":"[filtered]"}
```

#### Custom

```go
type myRecord struct {
    ID          string
    SecurePhone string
}
record := myRecord{
    ID:          "userId",
    SecurePhone: "090-1234-5678",
}

maskTool := mask.New(mask.FieldPrefixFilter("Secure", masker.MMobile))
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId","SecurePhone":"090-***4-5678"}
```

### By specified value

#### Default

```go
const issuedToken = "abcd123456"
maskTool := mask.New(mask.ValueFilter(issuedToken))

record := "Authorization: Bearer " + issuedToken
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
"Authorization: Bearer [filtered]"
```

#### Custom

```go
const issuedToken = "abcd123456"
maskTool := mask.New(mask.ValueFilter(issuedToken, masker.MPassword))

record := "Authorization: Bearer " + issuedToken
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
"Authorization: Bearer ************"
```

### By custom type

#### Default

```go
type password string
type apps []int32
type myRecord struct {
    ID       string
    Password password
    Apps     apps
}
record := myRecord{
    ID:       "userId",
    Password: "abcd1234",
    Apps:     apps{123, 456},
}
maskTool := mask.New(mask.TypeFilter(password("")), mask.TypeFilter(apps{}))
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId","Password":"[filtered]","Apps":null}
```

#### Custom

```go
type password string
type apps []int32
type myRecord struct {
    ID       string
    Password password
    Apps     apps
}
record := myRecord{
    ID:       "userId",
    Password: "abcd1234",
    Apps:     apps{123, 456},
}
maskTool := mask.New(mask.TypeFilter(password(""), masker.MPassword), mask.TypeFilter(apps{}))
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId","Password":"************","Apps":null}
```

### By struct tag

#### Default

```go
type myRecord struct {
    ID    string
    EMail string `mask:"secret"`
}
record := myRecord{
    ID:    "userId",
    EMail: "dummy@dummy.com",
}
maskTool := mask.New(mask.TagFilter())
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId","EMail":"[filtered]"}
```

#### Custom

```go
type myRecord struct {
    ID    string
    EMail string `mask:"secret"`
    Phone string `mask:"mobile"`
}
record := myRecord{
    ID:    "userId",
    EMail: "dummy@dummy.com",
    Phone: "9191919191",
}
maskTool := mask.New(mask.TagFilter(masker.MEmail, masker.MMobile))
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId","EMail":"dummy@dummy.com","Phone":"9191***191"}
```

### By regex filter

#### Default

```go
customRegex := "^https:\\/\\/(dummy-backend.)[0-9a-z]*.com\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)$"
type myRecord struct {
    ID   string
    Link string
}
record := myRecord{
    ID:   "userId",
    Link: "https://dummy-backend.dummy.com/v2/random",
}
maskTool := mask.New(mask.RegexFilter(customRegex))
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId","Link":"[filtered]"}
```

#### Custom

```go
customRegex := "^https:\\/\\/(dummy-backend.)[0-9a-z]*.com\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)$"
type myRecord struct {
    ID   string
    Link string
}
record := myRecord{
    ID:   "userId",
    Link: "https://dummy-backend.dummy.com/v2/random",
}
maskTool := mask.New(mask.RegexFilter(customRegex, masker.MPassword))
filteredData := maskTool.Apply(record)

fmt.Println(filteredData)
```

```json
{"ID":"userId","Link":"************"}
```

## Customise Masker

```go
mask.SetMasker(masker.NewMasker(
    masker.WithMarkTypes("custom1", "custom2", "custom3"),
    masker.WithStringFunc(func(t masker.MType, s string) (bool, string) {
        return false, ""
    }),
    masker.WithMaskingCharacter(string(masker.PCross)),
    masker.WithFilteredLabel("[removed]"),
))
```
