package feedservice

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/greenbone/eulabeia/connection"
	"github.com/greenbone/eulabeia/feedservice/handler"
	"github.com/greenbone/eulabeia/feedservice/redis"
	"github.com/greenbone/eulabeia/models"
)

type DBConnecion interface {
	Close() error
	GetList(db int, key string, start int, end int) ([]string, error)
}

// feed is the struct representing the feedservice
type feed struct {
	mqtt    connection.PubSub
	context string
	rc      DBConnecion
	id      string
}

// getSeverity filters the necessary infomation from the tags of a nvti to
// create a models.SeverityType
func getSeverity(tags map[string]string) models.SeverityType {
	var severetyVector string
	var severetyVersion string
	var severetyDate string

	if v, ok := tags["severety_vector"]; ok {
		severetyVector = v
	} else {
		severetyVector = tags["cvss_base_vector"]
	}
	if strings.Contains(severetyVector, "CVSS:3") {
		severetyVersion = "cvss_base_v3"
	} else {
		severetyVersion = "cvss_base_v2"
	}
	if v, ok := tags["severety_date"]; ok {
		severetyDate = v
	} else {
		severetyDate = tags["creation_date"]
	}

	return models.SeverityType{
		SeverityVector:  severetyVector,
		SeverityVersion: severetyVersion,
		SeverityDate:    severetyDate,
		SeverityOrigin:  tags["severity_origin"],
	}
}

// getRefs expects a comma separated list of cves, bids and xrefs. The function
// splitts them and put them into a list of models.RefType
func getRefs(cves string, bids string, xrefs string) []models.RefType {
	var cveSlice []string
	var bidSlice []string
	var xrefSlice []string

	var l int

	if cves != "" {
		cveSlice = strings.Split(cves, ", ")
		l += len(cveSlice)
	}

	if bids != "" {
		bidSlice = strings.Split(bids, ", ")
		l += len(bidSlice)
	}

	if xrefs != "" {
		xrefSlice = strings.Split(xrefs, ", ")
		l += len(xrefSlice)
	}

	ret := make([]models.RefType, l)
	i := 0
	for _, v := range cveSlice {
		ret[i] = models.RefType{
			Type: "CVE",
			ID:   v,
		}
		i++
	}
	for _, v := range bidSlice {
		ret[i] = models.RefType{
			Type: "BID",
			ID:   v,
		}
		i++
	}
	for _, v := range xrefSlice {
		xref := strings.SplitN(v, ":", 2)
		if len(xref) != 2 {
			continue
		}
		ret[i] = models.RefType{
			Type: xref[0],
			ID:   xref[1],
		}
		i++
	}

	return ret
}

// getNvtPrefs expects an oid corresponding to a nvt. The function parses the
// preferences of a nvt into a list of models.VTParamType
func (f *feed) getNvtPrefs(oid string) []models.VTParamType {
	key := fmt.Sprintf("oid:%s:prefs", oid)
	prefs, err := f.rc.GetList(1, key, 0, -1)
	if err != nil {
		return nil
	}

	ret := make([]models.VTParamType, len(prefs))
	for i, v := range prefs {
		pref := strings.Split(v, "|||")
		id, err := strconv.Atoi(pref[0])
		if err != nil {
			return nil
		}
		def := ""
		if len(pref) > 3 {
			def = pref[3]
		}
		ret[i] = models.VTParamType{
			ParameterID:           id,
			ParameterName:         pref[1],
			ParameterValue:        "",
			ParameterType:         pref[2],
			ParameterDescription:  "Description",
			ParameterDefaultValue: def,
		}
	}
	return ret

}

// getVt expects a list of OIDs of VTs and returns a list of the corresponding
// VTs. If the oids list is empty all VTs are returned. If the VTs are currently
// loading nil will be returned.
func (f *feed) GetVt(oid string) (models.VT, error) {
	pref, err := f.rc.GetList(1, fmt.Sprintf("nvt:%s", oid), 0, -1)
	if err != nil {
		return models.VT{}, err
	}

	dependecies := strings.Split(pref[redis.NVT_DEPENDENCIES_POS], ", ")
	allTags := strings.Split(pref[redis.NVT_TAGS_POS], "|")
	tags := make(map[string]string)
	log.Println("test")

	for _, v := range allTags {
		log.Println("test")
		tag := strings.SplitN(v, "=", 2)
		tags[tag[0]] = tag[1]
	}
	log.Println("test")
	refs := getRefs(pref[redis.NVT_CVES_POS], pref[redis.NVT_BIDS_POS], pref[redis.NVT_XREFS_POS])

	var params []models.VTParamType

	timeout := pref[redis.NVT_TIMEOUT_POS]
	if timeout != "" {
		params = []models.VTParamType{
			{
				ParameterID:           0,
				ParameterName:         "timeout",
				ParameterType:         "entry",
				ParameterDescription:  "Script Timeout",
				ParameterDefaultValue: timeout,
			},
		}
		params = append(params, f.getNvtPrefs(oid)...)
	}
	vt := models.VT{
		OID:                oid,
		Name:               pref[redis.NVT_NAME_POS],
		FileName:           pref[redis.NVT_FILENAME_POS],
		RequiredKeys:       pref[redis.NVT_REQUIRED_KEYS_POS],
		MandatoryKeys:      pref[redis.NVT_MANDATORY_KEYS_POS],
		ExcludedKeys:       pref[redis.NVT_EXCLUDED_KEYS_POS],
		RequiredPorts:      pref[redis.NVT_REQUIRED_PORTS_POS],
		RequiredUDPPorts:   pref[redis.NVT_REQUIRED_UDP_PORTS_POS],
		Category:           pref[redis.NVT_CATEGORY_POS],
		Family:             pref[redis.NVT_FAMILY_POS],
		VTCreationTime:     tags["creation_date"],
		VTModificationTime: tags["last_modification"],
		Summary:            tags["summary"],
		Solution:           tags["solution"],
		SolutionType:       tags["solution_type"],
		SolutionMethod:     tags["solution_method"],
		Impact:             tags["impact"],
		Insight:            tags["insight"],
		Affected:           tags["affected"],
		Vuldetect:          tags["vuldetect"],
		QoDType:            tags["qod_type"],
		QoDValue:           tags["qod"],
		References:         refs,
		VTParameters:       params,
		VTDependencies:     dependecies,
		Severity:           getSeverity(tags),
	}

	return vt, err
}

// Start starts the feed service
func (f *feed) Start() {
	// MQTT Subscription Map

	f.mqtt.Subscribe(map[string]connection.OnMessage{
		fmt.Sprintf("%s/feed/cmd/%s", f.context, f.id): handler.FeedHandler{
			GetVt:   f.GetVt,
			Context: f.context,
		},
	})
}

func (f *feed) Close() error {
	return f.rc.Close()
}

// NewScheduler creates a new scheduler
func NewFeed(mqtt connection.PubSub, context string, id string) *feed {
	return &feed{
		mqtt:    mqtt,
		context: context,
		rc:      redis.NewRedisConnection("unix", "/run/redis-openvas/redis.sock"),
		id:      id,
	}
}
