// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ytypes

import (
	"fmt"

	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/util"
)

// UnmarshalOpt is an interface used for any option to be supplied
// to the Unmarshal function. Types implementing it can be used to
// control the behaviour of JSON unmarshalling.
type UnmarshalOpt interface {
	IsUnmarshalOpt()
}

// IgnoreExtraFields is an unmarshal option that controls the
// behaviour of the Unmarshal function when additional fields are
// found in the input JSON. By default, an error will be returned,
// by specifying the IgnoreExtraFields option to Unmarshal, extra
// fields will be discarded.
type IgnoreExtraFields struct{}

// IsUnmarshalOpt marks IgnoreExtraFields as a valid UnmarshalOpt.
func (*IgnoreExtraFields) IsUnmarshalOpt() {}

// Unmarshal recursively unmarshals JSON data tree in value into the given
// parent, using the given schema. Any values already in the parent that are
// not present in value are preserved.
func Unmarshal(schema *yang.Entry, parent interface{}, value interface{}, opts ...UnmarshalOpt) error {
	util.Indent()
	defer util.Dedent()

	// Nil value means the field is unset.
	if util.IsValueNil(value) {
		return nil
	}
	if schema == nil {
		return fmt.Errorf("nil schema for parent type %T, value %v (%T)", parent, value, value)
	}
	util.DbgPrint("Unmarshal value %v, type %T, into parent type %T, schema name %s", util.ValueStrDebug(value), value, parent, schema.Name)

	switch {
	case schema.IsLeaf():
		return unmarshalLeaf(schema, parent, value)
	case schema.IsLeafList():
		return unmarshalLeafList(schema, parent, value)
	case schema.IsList():
		return unmarshalList(schema, parent, value, opts...)
	case schema.IsChoice():
		return fmt.Errorf("cannot pass choice schema %s to Unmarshal", schema.Name)
	case schema.IsContainer():
		return unmarshalContainer(schema, parent, value, opts...)
	}
	return fmt.Errorf("unknown schema type for type %T, value %v", value, value)
}

// hasIgnoreExtraFields determines whether the supplied slice of UnmarshalOpts contains
// the IgnoreExtraFields option.
func hasIgnoreExtraFields(opts []UnmarshalOpt) bool {
	for _, o := range opts {
		if _, ok := o.(*IgnoreExtraFields); ok {
			return true
		}
	}
	return false
}
