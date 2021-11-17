package feedservice

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/greenbone/eulabeia/messages/cmds"
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
			Description:  "",
			DefaultValue: "",
		},
		{
			ID:           3,
			Name:         "Test3:",
			Value:        "",
			Type:         "password",
			Description:  "",
			DefaultValue: "",
		},
		{
			ID:           2,
			Name:         "Test2:",
			Value:        "",
			Type:         "entry",
			Description:  "",
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
	case "nvt:test2":
		return []string{
			"filename2.nasl",
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
				"solution_type=vendorfix|" +
				"qod_type=test",
			"CVE-2000-0003, CVE-2000-0004",
			"test_bid",
			"URL:https://example.org, Advisory-ID:testID",
			"1",
			"foo_bar_family",
			"name2",
		}, nil
	default:
		return nil, fmt.Errorf("unknown key")
	}
}

func (rm RedisMock) GetKeys(db int, filter string) ([]string, error) {
	return []string{
		"nvt:test",
		"nvt:test2",
	}, nil
}

func TestGetVT(t *testing.T) {
	fs := feed{
		context: "",
		rc:      &RedisMock{},
	}

	vtTest, f, err := fs.GetVT(cmds.NewGet("vt", "test", "", ""))
	if err != nil || f != nil {
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

func TestResolveFilter(t *testing.T) {
	fs := feed{
		context: "",
		rc:      &RedisMock{},
	}
	var oids []string
	var err error

	oids, err = fs.ResolveFilter([]models.VTFilter{
		{
			Key:   "family",
			Value: "foo_bar_family",
		},
	})
	if err != nil {
		t.Errorf("Unable to resolve family: %s", err)
	}
	if len(oids) != 1 {
		t.Fatalf("wrong number of oids: expected %d, got %d", 1, len(oids))
	}
	if oids[0] != "test2" {
		t.Errorf("Got wrong OID: expected %s, got %s", "test2", oids[0])
	}

	oids, err = fs.ResolveFilter([]models.VTFilter{
		{
			Key:   "category",
			Value: "1",
		},
	})
	if err != nil {
		t.Errorf("Unable to resolve category: %s", err)
	}
	if len(oids) != 2 {
		t.Fatalf("wrong number of oids: expected %d, got %d", 2, len(oids))
	}

	oids, err = fs.ResolveFilter([]models.VTFilter{
		{
			Key:   "tag",
			Value: "solution_type=noFix",
		},
	})
	if err != nil {
		t.Errorf("Unable to resolve tag: %s", err)
	}
	if len(oids) != 1 {
		t.Fatalf("wrong number of oids: expected %d, got %d", 1, len(oids))
	}
	if oids[0] != "test" {
		t.Errorf("Got wrong OID: expected %s, got %s", "test", oids[0])
	}

	oids, err = fs.ResolveFilter([]models.VTFilter{
		{
			Key:   "cve",
			Value: "CVE-2000-0003",
		},
	})
	if err != nil {
		t.Errorf("Unable to resolve cve: %s", err)
	}
	if len(oids) != 1 {
		t.Fatalf("wrong number of oids: expected %d, got %d", 1, len(oids))
	}
	if oids[0] != "test2" {
		t.Errorf("Got wrong OID: expected %s, got %s", "test2", oids[0])
	}

	oids, err = fs.ResolveFilter([]models.VTFilter{
		{
			Key:   "name",
			Value: "name2",
		},
	})
	if err != nil {
		t.Errorf("Unable to resolve name: %s", err)
	}
	if len(oids) != 1 {
		t.Fatalf("wrong number of oids: expected %d, got %d", 1, len(oids))
	}
	if oids[0] != "test2" {
		t.Errorf("Got wrong OID: expected %s, got %s", "test2", oids[0])
	}

	oids, err = fs.ResolveFilter([]models.VTFilter{
		{
			Key:   "filename",
			Value: "filename2.nasl",
		},
	})
	if err != nil {
		t.Errorf("Unable to resolve filename: %s", err)
	}
	if len(oids) != 1 {
		t.Fatalf("wrong number of oids: expected %d, got %d", 1, len(oids))
	}
	if oids[0] != "test2" {
		t.Errorf("Got wrong OID: expected %s, got %s", "test2", oids[0])
	}

	oids, err = fs.ResolveFilter([]models.VTFilter{
		{
			Key:   "bid",
			Value: "test_bid",
		},
	})
	if err != nil {
		t.Errorf("Unable to resolve bid: %s", err)
	}
	if len(oids) != 1 {
		t.Fatalf("wrong number of oids: expected %d, got %d", 1, len(oids))
	}
	if oids[0] != "test2" {
		t.Errorf("Got wrong OID: expected %s, got %s", "test2", oids[0])
	}

}
