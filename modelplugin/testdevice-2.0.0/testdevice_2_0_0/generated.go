/*
Package testdevice_2_0_0 is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was false
in this case).

This package was generated by /home/scondon/go/pkg/mod/github.com/openconfig/ygot@v0.6.1-0.20200103195725-e3c44fa43926/genutil/names.go
using the following YANG input files:
	- test1@2019-06-10.yang
Imported modules were sourced from:
	- yang/...
*/
package testdevice_2_0_0

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ytypes"
)

// Binary is a type that is used for fields that have a YANG type of
// binary. It is used such that binary fields can be distinguished from
// leaf-lists of uint8s (which are mapped to []uint8, equivalent to
// []byte in reflection).
type Binary []byte

// YANGEmpty is a type that is used for fields that have a YANG type of
// empty. It is used such that empty fields can be distinguished from boolean fields
// in the generated code.
type YANGEmpty bool

var (
	SchemaTree map[string]*yang.Entry
)

func init() {
	var err error
	if SchemaTree, err = UnzipSchema(); err != nil {
		panic("schema error: " +  err.Error())
	}
}

// Schema returns the details of the generated schema.
func Schema() (*ytypes.Schema, error) {
	uzp, err := UnzipSchema()
	if err != nil {
		return nil, fmt.Errorf("cannot unzip schema, %v", err)
	}

	return &ytypes.Schema{
		Root: &Device{},
		SchemaTree: uzp,
		Unmarshal: Unmarshal,
	}, nil
}

// UnzipSchema unzips the zipped schema and returns a map of yang.Entry nodes,
// keyed by the name of the struct that the yang.Entry describes the schema for.
func UnzipSchema() (map[string]*yang.Entry, error) {
	var schemaTree map[string]*yang.Entry
	var err error
	if schemaTree, err = ygot.GzipToSchema(ySchema); err != nil {
		return nil, fmt.Errorf("could not unzip the schema; %v", err)
	}
	return schemaTree, nil
}

// Unmarshal unmarshals data, which must be RFC7951 JSON format, into
// destStruct, which must be non-nil and the correct GoStruct type. It returns
// an error if the destStruct is not found in the schema or the data cannot be
// unmarshaled. The supplied options (opts) are used to control the behaviour
// of the unmarshal function - for example, determining whether errors are
// thrown for unknown fields in the input JSON.
func Unmarshal(data []byte, destStruct ygot.GoStruct, opts ...ytypes.UnmarshalOpt) error {
	tn := reflect.TypeOf(destStruct).Elem().Name()
	schema, ok := SchemaTree[tn]
	if !ok {
		return fmt.Errorf("could not find schema for type %s", tn )
	}
	var jsonTree interface{}
	if err := json.Unmarshal([]byte(data), &jsonTree); err != nil {
		return err
	}
	return ytypes.Unmarshal(schema, destStruct, jsonTree, opts...)
}

// Device represents the /device YANG schema element.
type Device struct {
	Cont1A	*Test1_Cont1A	`path:"cont1a" module:"test1"`
	Cont1BState	*Test1_Cont1BState	`path:"cont1b-state" module:"test1"`
	LeafAtTopLevel	*string	`path:"leafAtTopLevel" module:"test1"`
}

// IsYANGGoStruct ensures that Device implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Device) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Device"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Device) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A represents the /test1/cont1a YANG schema element.
type Test1_Cont1A struct {
	Cont2A	*Test1_Cont1A_Cont2A	`path:"cont2a" module:"test1"`
	Leaf1A	*string	`path:"leaf1a" module:"test1"`
	List2A	map[string]*Test1_Cont1A_List2A	`path:"list2a" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A) IsYANGGoStruct() {}

// NewList2A creates a new entry in the List2A list of the
// Test1_Cont1A struct. The keys of the list are populated from the input
// arguments.
func (t *Test1_Cont1A) NewList2A(Name string) (*Test1_Cont1A_List2A, error){

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.List2A == nil {
		t.List2A = make(map[string]*Test1_Cont1A_List2A)
	}

	key := Name

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.List2A[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list List2A", key)
	}

	t.List2A[key] = &Test1_Cont1A_List2A{
		Name: &Name,
	}

	return t.List2A[key], nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A_Cont2A represents the /test1/cont1a/cont2a YANG schema element.
type Test1_Cont1A_Cont2A struct {
	Leaf2A	*uint8	`path:"leaf2a" module:"test1"`
	Leaf2B	*float64	`path:"leaf2b" module:"test1"`
	Leaf2C	*string	`path:"leaf2c" module:"test1"`
	Leaf2D	*float64	`path:"leaf2d" module:"test1"`
	Leaf2E	[]int16	`path:"leaf2e" module:"test1"`
	Leaf2F	Binary	`path:"leaf2f" module:"test1"`
	Leaf2G	*bool	`path:"leaf2g" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A_Cont2A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A_Cont2A) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A_Cont2A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A_Cont2A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A_Cont2A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A_List2A represents the /test1/cont1a/list2a YANG schema element.
type Test1_Cont1A_List2A struct {
	Name	*string	`path:"name" module:"test1"`
	RxPower	*uint16	`path:"rx-power" module:"test1"`
	TxPower	*uint16	`path:"tx-power" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A_List2A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A_List2A) IsYANGGoStruct() {}

// ΛListKeyMap returns the keys of the Test1_Cont1A_List2A struct, which is a YANG list entry.
func (t *Test1_Cont1A_List2A) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Name == nil {
		return nil, fmt.Errorf("nil value for key Name")
	}

	return map[string]interface{}{
		"name": *t.Name,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A_List2A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A_List2A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A_List2A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1BState represents the /test1/cont1b-state YANG schema element.
type Test1_Cont1BState struct {
	Cont2C	*Test1_Cont1BState_Cont2C	`path:"cont2c" module:"test1"`
	Leaf2D	*uint16	`path:"leaf2d" module:"test1"`
	List2B	map[uint8]*Test1_Cont1BState_List2B	`path:"list2b" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1BState implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1BState) IsYANGGoStruct() {}

// NewList2B creates a new entry in the List2B list of the
// Test1_Cont1BState struct. The keys of the list are populated from the input
// arguments.
func (t *Test1_Cont1BState) NewList2B(Index uint8) (*Test1_Cont1BState_List2B, error){

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.List2B == nil {
		t.List2B = make(map[uint8]*Test1_Cont1BState_List2B)
	}

	key := Index

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.List2B[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list List2B", key)
	}

	t.List2B[key] = &Test1_Cont1BState_List2B{
		Index: &Index,
	}

	return t.List2B[key], nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1BState) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1BState"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1BState) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1BState_Cont2C represents the /test1/cont1b-state/cont2c YANG schema element.
type Test1_Cont1BState_Cont2C struct {
	Leaf3A	*bool	`path:"leaf3a" module:"test1"`
	Leaf3B	*string	`path:"leaf3b" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1BState_Cont2C implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1BState_Cont2C) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1BState_Cont2C) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1BState_Cont2C"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1BState_Cont2C) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1BState_List2B represents the /test1/cont1b-state/list2b YANG schema element.
type Test1_Cont1BState_List2B struct {
	Index	*uint8	`path:"index" module:"test1"`
	Leaf3C	*string	`path:"leaf3c" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1BState_List2B implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1BState_List2B) IsYANGGoStruct() {}

// ΛListKeyMap returns the keys of the Test1_Cont1BState_List2B struct, which is a YANG list entry.
func (t *Test1_Cont1BState_List2B) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Index == nil {
		return nil, fmt.Errorf("nil value for key Index")
	}

	return map[string]interface{}{
		"index": *t.Index,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1BState_List2B) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1BState_List2B"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1BState_List2B) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }



var (
	// ySchema is a byte slice contain a gzip compressed representation of the
	// YANG schema from which the Go code was generated. When uncompressed the
	// contents of the byte slice is a JSON document containing an object, keyed
	// on the name of the generated struct, and containing the JSON marshalled
	// contents of a goyang yang.Entry struct, which defines the schema for the
	// fields within the struct.
	ySchema = []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x5d, 0x5f, 0x6f, 0xe2, 0x38,
		0x10, 0x7f, 0xe7, 0x53, 0x44, 0x79, 0x3c, 0x15, 0x35, 0x09, 0xa5, 0xea, 0xf2, 0x46, 0xdb, 0x5b,
		0x9d, 0xb4, 0xed, 0xdd, 0x6a, 0xb7, 0xba, 0x87, 0xab, 0xd0, 0xc9, 0x04, 0xc3, 0x5a, 0x17, 0x1c,
		0x94, 0x98, 0x5d, 0x50, 0xc5, 0x77, 0x3f, 0xe5, 0x0f, 0x94, 0x84, 0x24, 0x9e, 0x09, 0x4d, 0x81,
		0x32, 0x4f, 0xab, 0x0d, 0xe3, 0xd8, 0x1e, 0x8f, 0x7f, 0xf3, 0xc7, 0xbf, 0xb8, 0x2f, 0x2d, 0xc3,
		0x30, 0x0c, 0xf3, 0x4f, 0x36, 0xe5, 0x66, 0xcf, 0x30, 0x47, 0xfc, 0xa7, 0x70, 0xb9, 0x79, 0x91,
		0x3c, 0xfd, 0x22, 0xe4, 0xc8, 0xec, 0x19, 0x76, 0xfa, 0xdf, 0x3b, 0x5f, 0x8e, 0xc5, 0xc4, 0xec,
		0x19, 0x56, 0xfa, 0xe0, 0x5e, 0x04, 0x66, 0xcf, 0x48, 0x5e, 0x11, 0x3f, 0x70, 0x7d, 0xa9, 0x6c,
		0x96, 0x79, 0x96, 0x79, 0x7d, 0xfa, 0xfb, 0x45, 0xf6, 0xd7, 0x6c, 0x37, 0x9b, 0xc7, 0xf9, 0xee,
		0x36, 0x3f, 0x7c, 0x0d, 0xf8, 0x58, 0x2c, 0x76, 0x7a, 0xc9, 0xf4, 0xa4, 0xec, 0x5c, 0x2f, 0xf1,
		0xaf, 0xdf, 0xfd, 0x79, 0xe0, 0xf2, 0xc2, 0x96, 0xc9, 0x48, 0xf8, 0xf2, 0x97, 0x1f, 0x44, 0x83,
		0x31, 0x67, 0x49, 0x27, 0x17, 0xc5, 0x82, 0x7f, 0xb0, 0xb0, 0x1f, 0x4c, 0xe6, 0x53, 0x2e, 0x95,
		0xd9, 0x33, 0x54, 0x30, 0xe7, 0x25, 0x82, 0x5b, 0x52, 0xd1, 0x98, 0x76, 0x84, 0x56, 0x99, 0x27,
		0xab, 0xdc, 0x4c, 0xf3, 0x0a, 0xce, 0x28, 0xda, 0x61, 0xe5, 0x13, 0xd9, 0x56, 0xb8, 0xc3, 0xca,
		0x66, 0x51, 0xac, 0x78, 0xed, 0x02, 0x40, 0x16, 0x02, 0xb6, 0x20, 0xd0, 0x85, 0x41, 0x2f, 0x10,
		0x7a, 0xa1, 0xc0, 0x0b, 0x56, 0xbc, 0x70, 0x25, 0x0b, 0xa8, 0x5d, 0xc8, 0x8d, 0x80, 0xc7, 0xd9,
		0xb8, 0x62, 0x41, 0x77, 0xd4, 0x99, 0xca, 0x6b, 0x26, 0x73, 0xcf, 0xc7, 0x6c, 0xee, 0xc5, 0x73,
		0x71, 0x74, 0xb2, 0xa9, 0x31, 0x58, 0x1a, 0x31, 0x9d, 0x51, 0x60, 0x8c, 0x03, 0x67, 0x24, 0x58,
		0x63, 0xa9, 0x6d, 0x34, 0xb5, 0x8d, 0x07, 0x6d, 0x44, 0xd5, 0xc6, 0xa4, 0x31, 0xaa, 0x4d, 0x6f,
		0x4f, 0xcb, 0x19, 0xc7, 0xe9, 0x79, 0x2e, 0xa4, 0xba, 0x81, 0xa8, 0x3a, 0x35, 0x8a, 0x2e, 0x40,
		0xf4, 0x1b, 0x93, 0x93, 0xe8, 0xe5, 0xcf, 0x20, 0x15, 0xc1, 0x96, 0x2e, 0x7e, 0xf1, 0xa3, 0x90,
		0xe0, 0xb5, 0x46, 0x5a, 0xf3, 0x4e, 0xb3, 0xbf, 0x99, 0x37, 0xe7, 0xe5, 0x90, 0x58, 0xda, 0xee,
		0x73, 0xc0, 0x5c, 0x25, 0x7c, 0x79, 0x2f, 0x26, 0x42, 0x85, 0x51, 0xc7, 0xe0, 0xf6, 0xab, 0x0b,
		0x84, 0x2a, 0xd8, 0xe2, 0xdd, 0x55, 0xd1, 0x79, 0x47, 0x55, 0xb4, 0xde, 0x50, 0x61, 0xc7, 0x6e,
		0x61, 0x64, 0x62, 0xaf, 0xba, 0x38, 0x3e, 0x1b, 0xd3, 0x4a, 0x0d, 0x5a, 0xf5, 0xda, 0x57, 0xac,
		0x45, 0xe2, 0xd4, 0x87, 0xc8, 0x20, 0x60, 0xf8, 0xd6, 0x8e, 0xdd, 0x26, 0xc7, 0x7e, 0xb2, 0x8e,
		0x7d, 0xc4, 0x5d, 0x31, 0x65, 0xde, 0xf5, 0x15, 0xc2, 0xb9, 0xdb, 0x0e, 0x40, 0x76, 0x67, 0xc7,
		0x75, 0x28, 0x24, 0xc0, 0x6a, 0xec, 0xc3, 0xe0, 0xb5, 0x63, 0x59, 0xd6, 0x3b, 0x6a, 0xe3, 0xd8,
		0x11, 0xdb, 0x45, 0x22, 0xb6, 0xfb, 0xd6, 0x88, 0xed, 0x10, 0x62, 0x9f, 0x2c, 0x62, 0x87, 0x2a,
		0x10, 0x72, 0x82, 0x81, 0xeb, 0x9b, 0xa6, 0xec, 0x78, 0x84, 0xb4, 0xe3, 0x11, 0x95, 0x14, 0xc8,
		0x8e, 0x29, 0xf2, 0xa0, 0xc8, 0x83, 0x22, 0x8f, 0x03, 0x45, 0x1e, 0x1c, 0x89, 0xd8, 0x9c, 0x10,
		0x9b, 0x10, 0x7b, 0xad, 0x67, 0x21, 0x95, 0x7d, 0x8d, 0x40, 0x6b, 0xe7, 0x54, 0x71, 0xd7, 0xae,
		0x8b, 0xbb, 0xfb, 0x63, 0x8d, 0xf5, 0x91, 0x90, 0xf7, 0xc3, 0x15, 0xe9, 0x34, 0x5b, 0xea, 0x41,
		0x84, 0xaa, 0xaf, 0x54, 0x00, 0xdb, 0x56, 0x8f, 0x42, 0xfe, 0xee, 0xf1, 0x68, 0xc3, 0x47, 0x73,
		0x95, 0x73, 0xcf, 0x03, 0xec, 0x97, 0x47, 0xb6, 0xc0, 0x37, 0xfa, 0x2b, 0x18, 0xf1, 0x80, 0x8f,
		0x6e, 0x97, 0x69, 0x93, 0xa6, 0xbc, 0xcb, 0x18, 0xe9, 0x5d, 0xc6, 0xe4, 0x5d, 0xc8, 0xbb, 0xac,
		0xf5, 0x3c, 0x14, 0x92, 0x05, 0x4b, 0x84, 0x7b, 0xf9, 0x04, 0x10, 0x7d, 0xe0, 0x72, 0xa2, 0x7e,
		0x7c, 0x94, 0xb8, 0xde, 0x21, 0xf7, 0x72, 0x18, 0x5d, 0x1c, 0x7b, 0x58, 0x3f, 0x41, 0x02, 0xef,
		0x84, 0x80, 0x97, 0x80, 0x77, 0x03, 0xbc, 0xbe, 0xef, 0x71, 0x26, 0x31, 0x65, 0x18, 0xbb, 0xae,
		0x21, 0xa3, 0x38, 0x50, 0x7d, 0x29, 0x7d, 0xc5, 0xa2, 0x3d, 0x5b, 0x4d, 0x85, 0x0a, 0xdd, 0x1f,
		0x7c, 0xca, 0x66, 0x2c, 0x06, 0x7a, 0xf3, 0x52, 0xf1, 0x50, 0xd9, 0x97, 0x09, 0x73, 0xf0, 0xb2,
		0x92, 0xcf, 0x96, 0xb4, 0x56, 0xc1, 0xdc, 0x55, 0x32, 0x55, 0xc6, 0x53, 0xd4, 0xf8, 0xdf, 0xbb,
		0xa8, 0x71, 0x3f, 0xfe, 0xc7, 0xe9, 0x17, 0xaf, 0xd8, 0xee, 0x54, 0x0a, 0xa6, 0x11, 0xef, 0x36,
		0x1b, 0x40, 0xbb, 0x4b, 0xe5, 0xaa, 0x69, 0x77, 0x16, 0xd1, 0xee, 0xf0, 0xbb, 0x0a, 0x67, 0x72,
		0xda, 0xdd, 0x03, 0x2f, 0xc3, 0xbf, 0x96, 0xdf, 0x2b, 0x64, 0x80, 0xf1, 0x09, 0x2c, 0x9b, 0x40,
		0x80, 0x20, 0xca, 0xf7, 0x6e, 0x7c, 0x6e, 0x17, 0x28, 0x5f, 0xc3, 0xd5, 0xae, 0x60, 0xb9, 0x4f,
		0xe3, 0x53, 0xb4, 0xad, 0x06, 0xe7, 0x58, 0x13, 0x35, 0x07, 0xfb, 0x00, 0x90, 0x08, 0x41, 0xbc,
		0xdf, 0x54, 0x8e, 0x78, 0xbf, 0xc7, 0xce, 0xfb, 0x4d, 0x1d, 0x15, 0x30, 0xda, 0x8b, 0xa5, 0x29,
		0xd6, 0xa3, 0x58, 0xaf, 0xfe, 0xe1, 0xf1, 0xd9, 0x65, 0xd9, 0x57, 0x94, 0x64, 0xaf, 0x55, 0x71,
		0x43, 0x39, 0x76, 0x34, 0xad, 0x60, 0xd1, 0x9e, 0xf9, 0xbf, 0x78, 0x00, 0xc7, 0xdd, 0x4d, 0x0b,
		0xc2, 0x5e, 0xc2, 0xde, 0xad, 0x6f, 0x28, 0x50, 0xe7, 0x67, 0xd7, 0xe7, 0xc6, 0x5b, 0xa0, 0xfa,
		0xe6, 0xd6, 0x57, 0x14, 0x54, 0xdf, 0x8c, 0xa7, 0xa5, 0xd0, 0xd8, 0xab, 0x08, 0x7b, 0x09, 0x7b,
		0x09, 0x7b, 0x91, 0x78, 0x43, 0x5f, 0x17, 0x9d, 0xe7, 0xd1, 0x12, 0xaa, 0x3a, 0xf1, 0x85, 0x2f,
		0x35, 0x55, 0x05, 0x18, 0x07, 0x02, 0xc5, 0x7d, 0xc8, 0x71, 0x1e, 0x80, 0x5e, 0x40, 0x47, 0x2b,
		0xc5, 0x80, 0xea, 0x36, 0xa0, 0x4e, 0xd9, 0xa2, 0xcd, 0xd7, 0xa3, 0x01, 0x80, 0x44, 0x1d, 0x48,
		0xcd, 0xc0, 0xe9, 0x95, 0xd9, 0x80, 0x57, 0x05, 0x31, 0x42, 0x9a, 0x3c, 0x97, 0xa9, 0xac, 0x37,
		0x1a, 0x9a, 0x73, 0x99, 0xc8, 0xc6, 0x10, 0xe7, 0x32, 0x95, 0x37, 0x26, 0x68, 0x26, 0x53, 0x35,
		0x89, 0xa2, 0x0b, 0x23, 0xca, 0x47, 0x9d, 0x1d, 0xee, 0xeb, 0xa0, 0xb6, 0x06, 0x94, 0x5c, 0x77,
		0x31, 0x6c, 0x87, 0x8a, 0x29, 0xae, 0xb9, 0x14, 0x63, 0x2d, 0x85, 0xbc, 0x1a, 0xc3, 0xa1, 0xab,
		0x31, 0xf2, 0xaf, 0x8d, 0x8f, 0x08, 0x5d, 0xe0, 0xd5, 0x18, 0x2e, 0x95, 0xc8, 0x4f, 0xe2, 0x6a,
		0x8c, 0x0e, 0xf2, 0x6a, 0x8c, 0x0e, 0xa3, 0x74, 0x81, 0xd2, 0x85, 0x43, 0x50, 0x22, 0x34, 0xdc,
		0x9e, 0x0e, 0xf2, 0xf3, 0xee, 0xce, 0x90, 0x0c, 0x99, 0x0c, 0x99, 0xce, 0x7b, 0xe0, 0xe9, 0x5e,
		0x97, 0x32, 0x5f, 0x2c, 0xfb, 0xe1, 0xfc, 0x32, 0xdf, 0xfd, 0x73, 0x9e, 0x34, 0x60, 0xbf, 0xac,
		0x0c, 0x23, 0xab, 0x73, 0x88, 0xdb, 0xef, 0xd1, 0x1b, 0x12, 0x5a, 0xda, 0xdd, 0xbe, 0xb4, 0xb4,
		0x8a, 0xaf, 0x77, 0x61, 0x5f, 0xed, 0x82, 0x69, 0x69, 0x0e, 0x85, 0xbc, 0x7b, 0x5a, 0x1f, 0x9c,
		0x96, 0xa6, 0x2d, 0x74, 0x02, 0x0a, 0x9c, 0xc0, 0xc2, 0xe6, 0x91, 0x90, 0xd2, 0x6c, 0x0b, 0xfc,
		0x89, 0xe7, 0x29, 0x13, 0xd3, 0x9c, 0x86, 0xa7, 0x79, 0x28, 0x6e, 0xda, 0x10, 0xc8, 0x4d, 0x1b,
		0xee, 0x9b, 0x78, 0x13, 0x0a, 0x35, 0x9e, 0x78, 0x0b, 0x39, 0xe2, 0x0b, 0x78, 0xba, 0x92, 0x88,
		0x53, 0xb6, 0x42, 0xd9, 0x0a, 0xdd, 0x32, 0x09, 0x73, 0x02, 0xc4, 0x8f, 0x78, 0xf5, 0x87, 0xdd,
		0x2e, 0x11, 0x24, 0x36, 0x45, 0x22, 0xe4, 0x8d, 0x52, 0x1d, 0xba, 0x51, 0x8a, 0x60, 0x97, 0x8a,
		0x44, 0x98, 0x2c, 0x83, 0x90, 0x97, 0xe8, 0x11, 0xba, 0x00, 0x39, 0xa5, 0x47, 0x54, 0xc5, 0xb5,
		0xcd, 0xf3, 0x23, 0x34, 0xc2, 0x07, 0x3e, 0xf9, 0x5f, 0x57, 0xc1, 0x2a, 0x73, 0x3a, 0x50, 0x15,
		0x2c, 0x26, 0x01, 0xdc, 0x1e, 0x9a, 0x04, 0x50, 0x7c, 0x0c, 0x0f, 0x99, 0x00, 0x84, 0x0f, 0x10,
		0xf9, 0xea, 0xbe, 0x7a, 0xf2, 0x67, 0x0f, 0xfc, 0x27, 0xf7, 0xca, 0x19, 0x01, 0x39, 0xb9, 0x62,
		0x4e, 0x80, 0x45, 0x7f, 0x2e, 0x23, 0xbf, 0xee, 0xa5, 0xbe, 0x52, 0xef, 0x1b, 0xab, 0x7c, 0xa1,
		0xf9, 0x95, 0x29, 0xc5, 0x03, 0x59, 0xea, 0xfc, 0xcc, 0xe7, 0x7e, 0xfb, 0x9f, 0xc1, 0x4b, 0x67,
		0xd5, 0x7e, 0xb6, 0xda, 0x9f, 0x06, 0xbf, 0xed, 0x8e, 0x7b, 0x50, 0x66, 0x1d, 0xad, 0xad, 0x79,
		0x94, 0xd9, 0xad, 0x29, 0xc2, 0xcf, 0xec, 0x3f, 0xfe, 0xcd, 0xf7, 0x77, 0xb5, 0x97, 0xb7, 0x65,
		0x73, 0xfb, 0xa7, 0x8c, 0xc5, 0xde, 0x27, 0x7f, 0xda, 0x25, 0xe9, 0xb0, 0xb5, 0xfa, 0x1f, 0x00,
		0x00, 0xff, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x02, 0x6d, 0x7a, 0xa4, 0xf9, 0x65, 0x00, 0x00,
	}
)


// ΛEnumTypes is a map, keyed by a YANG schema path, of the enumerated types that
// correspond with the leaf. The type is represented as a reflect.Type. The naming
// of the map ensures that there are no clashes with valid YANG identifiers.
var ΛEnumTypes = map[string][]reflect.Type{
}
