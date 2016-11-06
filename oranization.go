package unapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unapi/config"
)

func FetchOranizations() (collection []Organization, results int) {
	var res []map[string]string
	url := fmt.Sprintf("%s/organizations?format=json&app_id=%s&app_key=%s", unapi.Domain, unapi.ApplicationID, unapi.ApplicationKey)
	r, _ := http.Get(url)
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Println(err)
	}
	for _, v := range res {
		org := Organization{}
		for k, vv := range v {
			switch k {
			case "full_name":
				org.SetFullName(vv)
			case "name":
				org.SetName(vv)
			}
		}
		collection = append(collection, org)
	}
	results = len(collection)
	return
}

type Organization struct {
	fullName, name string
}

func (org *Organization) SetName(name string) {
	org.name = name
}
func (org *Organization) SetFullName(fullName string) {
	org.fullName = fullName
}
