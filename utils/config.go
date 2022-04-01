package utils

import (
	"fmt"
	"reflect"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
	"github.com/spf13/viper"
)

var (
	configPath = "./config/"
)

type Config struct{}

func BuildDefaultConfig(storeInfo *types.StoreInfo) {
	viper.SetConfigName("Default")
	viper.SetConfigType("json")
	setupStore(storeInfo)
}

func setupStore(storeInfo *types.StoreInfo) {
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// load categories
	storeInfo.Categories = viper.GetStringSlice("categories")

	// load item info
	var items []map[string]interface{}
	marshalErr := viper.UnmarshalKey("items", &items)
	if marshalErr != nil {
		fmt.Println(marshalErr)
		return
	}

	// reflect used to handle viper unmarshal interface to correct type
	for _, contents := range items {
		var (
			name     string
			category string
			qty      int
			unitCost float64
		)
		for k, v := range contents {
			if k == "name" {
				name = reflect.ValueOf(v).String()
			}

			if k == "category" {
				category = reflect.ValueOf(v).String()
			}

			if k == "qty" {
				qty = int(reflect.ValueOf(v).Float())
			}

			if k == "unitCost" {
				unitCost = reflect.ValueOf(v).Float()
			}
		}

		categoryIndex := storeInfo.GetCategoryIndex(&category)

		storeInfo.AllItemInfo[name] = types.ItemInfo{Category: categoryIndex, Qty: qty, UnitCost: unitCost}
	}
}
