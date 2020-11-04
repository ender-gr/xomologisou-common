package xomologisou

import (
	"fmt"
	"testing"
)

func TestPublishable_BuildFormats(t *testing.T) {

	c := &Carrier{
		Posting: &CarrierPosting{
			Format:  `{n}{n}{m}{n}{n}[Από: ]{set-3}`,
			Hashtag: "test",
			Id:      1,
		},
		Form: &CarrierForm{
			OptionSets: map[string]CarrierOptions{
				"set-1": {
					AllowCustom: true,
				},
				"set-2": {
					AllowCustom: true,
				},
				"set-3": {
					AllowCustom: true,
				},
				"set-4": {
					AllowCustom: true,
				},
			},
		},
	}
	p := &Publishable{
		Content: "Hello world!",
		Options: map[string]string{
			"set-1": "University",
			"set-2": "Eclipse",
			"set-3": "IntelliJ",
		},
	}
	p.BuildFormats(c, 0)
	fmt.Println(p.FinalForm)
}
