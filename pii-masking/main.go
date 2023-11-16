package main

import (
	"fmt"

	v1 "github.com/devexps/go-examples/pii-masking/pb"

	"github.com/devexps/go-utils/mask"
	"github.com/devexps/go-utils/mask/masker"
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
