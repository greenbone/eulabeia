package feedservice

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/models"
)

var vt = models.VT{
	OID:                "test",
	Name:               "name",
	FileName:           "filename.nasl",
	RequiredKeys:       "required_key1, required_key2",
	MandatoryKeys:      "mandatory_key1, mandatory_key2",
	ExcludedKeys:       "excluded_key1, excluded_key2",
	RequiredPorts:      "42, 400",
	RequiredUDPPorts:   "22, 80",
	Category:           "1",
	Family:             "test_family",
	VTCreationTime:     "0",
	VTModificationTime: "0",
	Summary:            "test summary",
	Solution:           "none",
	SolutionType:       "noFix",
	SolutionMethod:     "foo",
	Impact:             "no impact",
	Insight:            "this is only for testing",
	Affected:           "this test",
	Vuldetect:          "test vulnerability",
	QoDType:            "test",
	QoDValue:           "",
	References: []models.RefType{
		{
			Type: "CVE",
			ID:   "CVE-2000-0001",
		},
		{
			Type: "CVE",
			ID:   "CVE-2000-0002",
		},
		{
			Type: "URL",
			ID:   "https://example.org",
		},
		{
			Type: "Advisory-ID",
			ID:   "testID",
		},
	},
	VTParameters: []models.VTParamType{
		{
			ID:           1,
			Name:         "Test1:",
			Value:        "",
			Type:         "entry",
			Description:  "Description",
			DefaultValue: "",
		},
		{
			ID:           3,
			Name:         "Test3:",
			Value:        "",
			Type:         "password",
			Description:  "Description",
			DefaultValue: "",
		},
		{
			ID:           2,
			Name:         "Test2:",
			Value:        "",
			Type:         "entry",
			Description:  "Description",
			DefaultValue: "default",
		},
	},
	VTDependencies: []string{
		"dependency1.nasl",
		"dependency2.nasl",
	},
	Severity: models.SeverityType{
		Vector:  "AV:N/AC:L/Au:N/C:P/I:P/A:P",
		Version: "cvss_base_v2",
		Date:    "0",
		Origin:  "",
	},
}

type MockPubSub struct{}

func (mps MockPubSub) Close() error {
	return nil
}

func (mps MockPubSub) Connect() error {
	return nil
}

func (mps MockPubSub) Publish(topic string, message interface{}) error {
	return nil
}

func (mps MockPubSub) Preprocess(topic string, message []byte) ([]connection.TopicData, bool) {
	return nil, false
}

func (mps MockPubSub) Subscribe(handler map[string]connection.OnMessage) error {
	return nil
}

type RedisMock struct {
}

func (rm RedisMock) Close() error {
	return nil
}

func (rm RedisMock) GetList(db int, key string, start int, end int) ([]string, error) {
	switch key {
	case "oid:test:prefs":
		return []string{
			"1|||Test1:|||entry|||",
			"3|||Test3:|||password|||",
			"2|||Test2:|||entry|||default",
		}, nil
	case "nvt:test":
		return []string{
			"filename.nasl",
			"required_key1, required_key2",
			"mandatory_key1, mandatory_key2",
			"excluded_key1, excluded_key2",
			"22, 80",
			"42, 400",
			"dependency1.nasl, dependency2.nasl",
			"cvss_base_vector=AV:N/AC:L/Au:N/C:P/I:P/A:P|" +
				"last_modification=0|" +
				"creation_date=0|" +
				"summary=test summary|" +
				"vuldetect=test vulnerability|" +
				"impact=no impact|" +
				"insight=this is only for testing|" +
				"affected=this test|" +
				"solution=none|" +
				"solution_method=foo|" +
				"solution_type=noFix|" +
				"qod_type=test",
			"CVE-2000-0001, CVE-2000-0002",
			"",
			"URL:https://example.org, Advisory-ID:testID",
			"1",
			"test_family",
			"name",
		}, nil
	default:
		return nil, fmt.Errorf("unknown key")
	}
}

func TestGetVT(t *testing.T) {
	fs := feed{
		mqtt:    MockPubSub{},
		context: "",
		rc:      &RedisMock{},
	}

	vtTest, err := fs.GetVt("test")
	if err != nil {
		t.Fatalf("Unable to get VT: %s\n", err)
	}

	vtTestJSON, err := json.Marshal(vtTest)
	if err != nil {
		t.Fatalf("Unable to get JSON of TestVT: %s\n", err)
	}

	vtJSON, err := json.Marshal(vt)
	if err != nil {
		t.Fatalf("Unable to get JSON of VT: %s\n", err)
	}

	if string(vtTestJSON) != string(vtJSON) {
		t.Fatalf("vtTestJSON != vtJSON\n\nvtTestJSON:\n %s\n\n vtJSON:\n%s\n", vtTestJSON, vtJSON)
	}
}
