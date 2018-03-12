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
	DataOmitEmpty       string        `json:"dataOmitEmpty,omitempty"`
	NullableData        *string       `json:"nullableData"`
	RequiredPointerData *string       `json:"requiredPointerData" required:"true"`
	NestedItem          NestedItem    `json:"nestedItem"`
	NestedItemPointer   *NestedItem   `json:"nestedItemPointer"`
	ArrayNoPointers     []NestedItem  `json:"arrayNoPointers"`
	ArrayPointers       []*NestedItem `json:"arrayPointers"`
}

type OtherRootType struct {
	Data string `json:"data"`
}

func main() {
	fmt.Println(generator.Generate(&Domain{}))
}
