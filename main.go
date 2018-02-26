package main

import (
	"encoding/json"
	"fmt"
	"github.com/xaionaro-go/zoodb-ru_taxonomy-csv-exporter/types"
	"io/ioutil"
	"strings"
)

func getJson(filePath string, objPtr interface{}) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, objPtr)
	if err != nil {
		panic(err)
	}
}

func getCategoryMap() (categoryMap map[int]types.Category) {
	getJson("../zoodb-ru_taxonomy-csv-exporter/categoryMap.json", &categoryMap)
	return
}

func getItems() (items []types.Item) {
	getJson("../zoodb-ru_taxonomy-csv-exporter/items.json", &items)
	return
}

func main() {
	categoryMap := getCategoryMap()
	items := getItems()

	for _, item := range items {
		category := categoryMap[item.CategoryId]
		var fullCategoryName []string
		for ; category.Id != 0; category = categoryMap[category.ParentId] {
			fullCategoryName = append([]string{category.Name}, fullCategoryName...)
		}

		fmt.Println(strings.Join(fullCategoryName, " "), item.Name)
		for _, synonym := range item.Synonyms {
			fmt.Println(strings.Join(fullCategoryName, " "), synonym)
		}
	}
}
