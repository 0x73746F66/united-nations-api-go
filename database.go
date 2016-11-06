package unapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unapi/config"
)

func FetchDatabases(organization Organization) (collection []Database, results int) {
	var res []map[string]string
	url := fmt.Sprintf("%s/%s/databases?format=json&app_id=%s&app_key=%s", unapi.Domain, organization.name, unapi.ApplicationID, unapi.ApplicationKey)
	r, _ := http.Get(url)
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Println(err)
	}

	for _, v := range res {
		db := Database{}
		for k, vv := range v {
			switch k {
			case "name":
				db.SetName(vv)
			}
		}
		collection = append(collection, db)
	}
	results = len(collection)
	return
}

type Database struct {
	name string
}

func (org *Database) SetName(name string) {
	org.name = name
}
