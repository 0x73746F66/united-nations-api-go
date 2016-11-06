package unapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unapi/config"
)

func FetchDatasets(organization Organization) (collection []Dataset, results int) {
	var res []map[string]string
	url := fmt.Sprintf("%s/%s/datasets?format=json&app_id=%s&app_key=%s", unapi.Domain, organization.name, unapi.ApplicationID, unapi.ApplicationKey)

	r, _ := http.Get(url)
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Println(err)
	}

	for _, v := range res {
		dataset := Dataset{}
		for k, _ := range v {
			fmt.Printf("k:%+v\n", k)
			// switch k {
			// case "name":
			// 	dataset.SetName(vv)
			// }
		}
		collection = append(collection, dataset)
	}
	results = len(collection)
	return
}

type Dataset struct {
	name   string
	topics []string
}

func (org *Dataset) SetName(name string) {
	org.name = name
}

func (org *Dataset) AddTopic(topic string) {
	org.topics = append(org.topics, topic)
}
