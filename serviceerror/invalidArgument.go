// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package serviceerror

import (
	"fmt"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

type (
	// InvalidArgument represents invalid argument error.
	InvalidArgument struct {
		Message string
		st *status.Status
	}
)

// NewInvalidArgument returns new InvalidArgument error.
func NewInvalidArgument(message string) *InvalidArgument {
	return &InvalidArgument{
		Message: message,
	}
}

// Error returns string message.
func (e *InvalidArgument) MessageArgs(a ...interface{}) *InvalidArgument {
	return NewInvalidArgument(fmt.Sprintf(e.Message, a))
}

// Error returns string message.
func (e *InvalidArgument) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *InvalidArgument) GRPCStatus() *status.Status {
	if e.st != nil{
		return e.st
	}

	return status.New(codes.InvalidArgument, e.Message)
}

func invalidArgument(st *status.Status) (*InvalidArgument, bool) {
	if st == nil || st.Code() != codes.InvalidArgument {
		return nil, false
	}

	return &InvalidArgument{
		Message: st.Message(),
		st: st,
	}, true
}
