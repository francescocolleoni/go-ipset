package commands

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/francescocolleoni/go-ipset/set"
	"github.com/francescocolleoni/go-ipset/utilities"
)

func TestCreateSetTranslateToCommandLine(t *testing.T) {
	type test struct {
		command *CreateSet
		expects [][]string
	}

	tests := []test{
		{
			NewCreateSet("testset", set.SetTypeBitmapIP, 0, 0, 0, false, false, false, false, ProtocolFamilyDefault),
			[][]string{{"create", "testset", "bitmap:ip"}},
		},
		{
			NewCreateSet("testset", set.SetTypeHashMAC, 0, 0, 0, false, false, false, false, ProtocolFamilyDefault),
			[][]string{{"create", "testset", "hash:mac"}},
		},
		{
			NewCreateSet("testset", set.SetTypeHashIP, 0, 0, 0, false, false, false, false, ProtocolFamilyDefault),
			[][]string{{"create", "testset", "hash:ip"}},
		},
		{
			NewCreateSet("testset", set.SetTypeHashIP, 10, 20, 30, true, true, true, true, ProtocolFamilyINet6),
			[][]string{
				{
					"create", "testset", "hash:ip",
					"timeout", "10", "hashsize", "20", "maxelem", "30",
					"counters", "skbinfo", "forceadd",
					"family", "inet6", "comment",
				},
			},
		},
	}

	for _, test := range tests {
		rawResult := test.command.TranslateToCommandLine()
		result := fmt.Sprintf("%v", rawResult)
		expects := fmt.Sprintf("%v", test.expects)

		if result != expects {
			t.Errorf("expectation failed: %s != %s (expected)", result, expects)
		}
	}
}

func TestCreateSet(t *testing.T) {
	// Create a test set, then check if it exists.
	utilities.RunIPSet("destroy", "testset")
	defer utilities.RunIPSet("destroy", "testset")

	//createSet := NewCreateSet("testset", set.SetTypeHashIP, 10, 20, 30, true, true, true, true, ProtocolFamilyINet6)
	createSet := NewCreateSet("testset", set.SetTypeHashIP, 10, 20, 30, true, true, true, true, ProtocolFamilyINet6)
	if err := createSet.Run(); err != nil {
		t.Error(err)
		t.FailNow()
	}

	ipsetOut, err := utilities.RunIPSet("list", "testset")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	regexs := []string{
		"Name: testset", "Type: hash:ip",
		"Header:.+timeout 10", "Header:.+counters", "Header:.+skbinfo",
		"Header:.+family inet6", "Header:.+hashsize 64", "Header:.+maxelem 30",
	}

	for _, pattern := range regexs {
		regex := regexp.MustCompile(pattern)
		if !regex.Match(ipsetOut) {
			t.Errorf("ipset list output expectation failed: no match for pattern %s", pattern)
		}
	}
}
