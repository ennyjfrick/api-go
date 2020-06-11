// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enums/v1/common.proto

package enums

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EncodingType int32

const (
	ENCODING_TYPE_UNSPECIFIED EncodingType = 0
	ENCODING_TYPE_PROTO3      EncodingType = 1
	ENCODING_TYPE_JSON        EncodingType = 2
)

var EncodingType_name = map[int32]string{
	0: "Unspecified",
	1: "Proto3",
	2: "Json",
}

var EncodingType_value = map[string]int32{
	"Unspecified": 0,
	"Proto3":      1,
	"Json":        2,
}

func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94ec60e502d91644, []int{0}
}

type IndexedValueType int32

const (
	INDEXED_VALUE_TYPE_UNSPECIFIED IndexedValueType = 0
	INDEXED_VALUE_TYPE_STRING      IndexedValueType = 1
	INDEXED_VALUE_TYPE_KEYWORD     IndexedValueType = 2
	INDEXED_VALUE_TYPE_INT         IndexedValueType = 3
	INDEXED_VALUE_TYPE_DOUBLE      IndexedValueType = 4
	INDEXED_VALUE_TYPE_BOOL        IndexedValueType = 5
	INDEXED_VALUE_TYPE_DATETIME    IndexedValueType = 6
)

var IndexedValueType_name = map[int32]string{
	0: "Unspecified",
	1: "String",
	2: "Keyword",
	3: "Int",
	4: "Double",
	5: "Bool",
	6: "Datetime",
}

var IndexedValueType_value = map[string]int32{
	"Unspecified": 0,
	"String":      1,
	"Keyword":     2,
	"Int":         3,
	"Double":      4,
	"Bool":        5,
	"Datetime":    6,
}

func (IndexedValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94ec60e502d91644, []int{1}
}

func init() {
	proto.RegisterEnum("enums.v1.EncodingType", EncodingType_name, EncodingType_value)
	proto.RegisterEnum("enums.v1.IndexedValueType", IndexedValueType_name, IndexedValueType_value)
}

func init() { proto.RegisterFile("enums/v1/common.proto", fileDescriptor_94ec60e502d91644) }

var fileDescriptor_94ec60e502d91644 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0x32, 0x41,
	0x18, 0xc7, 0x77, 0x7c, 0xdf, 0x24, 0xa6, 0x0e, 0xc3, 0x50, 0x96, 0x8a, 0x4f, 0xd0, 0x51, 0x6a,
	0x17, 0xf1, 0xd8, 0xc9, 0x75, 0x27, 0x99, 0xb2, 0x9d, 0x45, 0x57, 0xcb, 0x2e, 0x8b, 0xe9, 0x22,
	0x82, 0xbb, 0x23, 0xa6, 0x52, 0xb7, 0xbe, 0x40, 0xd0, 0xc7, 0xe8, 0xa3, 0x74, 0xf4, 0xe8, 0x31,
	0xc7, 0x4b, 0xa7, 0xf0, 0x23, 0x44, 0x1b, 0x42, 0xc2, 0x76, 0x7b, 0xf8, 0xff, 0x9e, 0xe7, 0xf9,
	0x1f, 0x7e, 0x78, 0xdf, 0x0f, 0x27, 0xc1, 0xbd, 0x31, 0x2d, 0x18, 0x1d, 0x19, 0x04, 0x32, 0xd4,
	0x87, 0x23, 0x39, 0x96, 0x74, 0x3b, 0x8a, 0xf5, 0x69, 0x21, 0xef, 0xe1, 0x5d, 0x16, 0x76, 0x64,
	0xb7, 0x1f, 0xf6, 0xdc, 0xc7, 0xa1, 0x4f, 0x73, 0x38, 0xcd, 0xec, 0xb2, 0xb0, 0xb8, 0x5d, 0xf1,
	0xdc, 0x96, 0xc3, 0xbc, 0x86, 0x5d, 0x77, 0x58, 0x99, 0x9f, 0x73, 0x66, 0x11, 0x8d, 0x1e, 0xe2,
	0xbd, 0x4d, 0xec, 0xd4, 0x84, 0x2b, 0x8a, 0x04, 0xd1, 0x14, 0xa6, 0x9b, 0xe4, 0xa2, 0x2e, 0x6c,
	0x92, 0xc8, 0x7f, 0x22, 0x4c, 0x78, 0xd8, 0xf5, 0x1f, 0xfc, 0x6e, 0xb3, 0x3d, 0x98, 0xf8, 0x51,
	0xcb, 0x31, 0x06, 0x6e, 0x5b, 0xec, 0x86, 0x59, 0x5e, 0xb3, 0x54, 0x6d, 0xb0, 0xb8, 0xaa, 0x1c,
	0x4e, 0xc7, 0xec, 0xd4, 0xdd, 0x1a, 0xb7, 0x2b, 0x04, 0x51, 0xc0, 0x99, 0x18, 0x7c, 0xc9, 0x5a,
	0xd7, 0xa2, 0x66, 0x91, 0x04, 0xcd, 0xe0, 0x54, 0x0c, 0xe7, 0xb6, 0x4b, 0xfe, 0xfd, 0xf1, 0xda,
	0x12, 0x0d, 0xb3, 0xca, 0xc8, 0x7f, 0x9a, 0xc5, 0x07, 0x31, 0xd8, 0x14, 0xa2, 0x4a, 0xb6, 0xe8,
	0x11, 0xce, 0xc6, 0xdd, 0x96, 0x5c, 0xe6, 0xf2, 0x2b, 0x46, 0x92, 0xe6, 0x33, 0x9a, 0x2d, 0x40,
	0x9b, 0x2f, 0x40, 0x5b, 0x2d, 0x00, 0x3d, 0x29, 0x40, 0xaf, 0x0a, 0xd0, 0x9b, 0x02, 0x34, 0x53,
	0x80, 0xde, 0x15, 0xa0, 0x0f, 0x05, 0xda, 0x4a, 0x01, 0x7a, 0x59, 0x82, 0x36, 0x5b, 0x82, 0x36,
	0x5f, 0x82, 0x86, 0x33, 0x7d, 0xa9, 0x8f, 0xfd, 0x60, 0x28, 0x47, 0xed, 0xc1, 0x8f, 0x25, 0x7d,
	0x2d, 0xc9, 0xdc, 0x29, 0x47, 0xf2, 0x9c, 0xef, 0xd4, 0x41, 0xb7, 0x27, 0xbd, 0x5f, 0xab, 0x7d,
	0x69, 0xac, 0xe7, 0xd3, 0xe8, 0xcc, 0x58, 0x2b, 0x3f, 0x8b, 0x86, 0xbb, 0x64, 0x94, 0x16, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xce, 0x24, 0xa4, 0x0b, 0x02, 0x00, 0x00,
}

func (x EncodingType) String() string {
	s, ok := EncodingType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x IndexedValueType) String() string {
	s, ok := IndexedValueType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
