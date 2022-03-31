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
- utils: Welcome message and menu input

