# NP Go Assignment: Shopping Cart

------

Assignment for NP Go course

## Getting started

```
go build
./assignment-shopping-cart-BlockByBlock
```

Note: Default items are loaded in `storeinfo.go` for convenience



## Folder structure

- selections: Processing logic for options on menu. The switch condition for the logic selection, `select.go` is in this folder.
- types: StoreInfo and ItemInfo
  - StoreInfo: Item categories and map of item name to item info
  - ItemInfo: category, quantity and unit cost
- utils: Welcome message and menu input, viper to parse json config file
- config: JSON config file for loading test/default data



## Config

`config.go` reads from the path set by `configPath` (hardcoded as ./config). It takes only config name of `default` and type `json`.

`Viper` is used to parse the JSON file. Viper does not come with default parsing of items (slices of object) - [here](https://github.com/spf13/viper#getting-values-from-viper). Reflect is used to parse the unmarshal interface from Viper to the expected type. E.g. Viper interpret number as Float64 instead of Int. Reflect is also useful in converting `interface{}` to iterable type through the use of available methods e.g. Len() when cast into reflect type.



## Reference

[Viper](https://github.com/spf13/viper)



