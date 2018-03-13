// Copyright 2017 Kozyrev Yury
// MIT license.
package main

import (
	"fmt"

	"github.com/SteveRuble/go-json-schema-generator"
)

type NestedItem struct {
	NestedItemValue string `json:"nestedItemValue" description:"Some nested value"`
}

type Domain struct {
	DataNoOmitEmpty     string           `json:"dataNoOmitEmpty" required:"true"`
	DataOmitEmpty       string           `json:"dataOmitEmpty,omitempty"`
	NullableData        *string          `json:"nullableData,omitempty"`
	RequiredPointerData *string          `json:"requiredPointerData,omitempty" required:"true"`
	NestedItem          NestedItem       `json:"nestedItem,omitempty"`
	NestedItemPointer   *NestedItem      `json:"nestedItemPointer,omitempty"`
	ArrayNoPointers     []NestedItem     `json:"arrayNoPointers,omitempty"`
	ArrayPointers       []*NestedItem    `json:"arrayPointers,omitempty"`
	OtherDefinedType    OtherDefinedType `json:"otherDefinedType,omitempty"`
}

type OtherDefinedType struct {
	Data string `json:"data"`
}

func main() {
	config := generator.Config{
		Root: Domain{},
		Definitions: map[string]interface{}{
			"nestedItem": NestedItem{},
			"other":      OtherDefinedType{},
		},
	}
	fmt.Println(generator.GenerateAll(config))
}
