/*
 * Copyright 2016 Dgraph Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 		http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package query

import (
	"time"

	"github.com/dgraph-io/dgraph/protos/graphp"
	"github.com/dgraph-io/dgraph/types"
	"github.com/dgraph-io/dgraph/x"
	geom "github.com/twpayne/go-geom"
)

// This file contains helper functions for converting scalar types to
// protobuf values.

func toProtoValue(v types.Val) *graphp.Value {
	switch v.Tid {
	case types.StringID:
		return &graphp.Value{&graphp.Value_StrVal{v.Value.(string)}}

	case types.Int32ID:
		return &graphp.Value{&graphp.Value_IntVal{v.Value.(int32)}}

	case types.FloatID:
		return &graphp.Value{&graphp.Value_DoubleVal{v.Value.(float64)}}

	case types.BoolID:
		return &graphp.Value{&graphp.Value_BoolVal{v.Value.(bool)}}

	case types.DateID:
		val := v.Value.(time.Time)
		return &graphp.Value{&graphp.Value_StrVal{val.Format(time.RFC3339)}}

	case types.DateTimeID:
		val := v.Value.(time.Time)
		return &graphp.Value{&graphp.Value_StrVal{val.Format(time.RFC3339)}}

	case types.GeoID:
		b := types.ValueForType(types.BinaryID)
		src := types.ValueForType(types.GeoID)
		src.Value = v.Value.(geom.T)
		x.Check(types.Marshal(src, &b))
		return &graphp.Value{&graphp.Value_GeoVal{b.Value.([]byte)}}

	case types.PasswordID:
		return &graphp.Value{&graphp.Value_PasswordVal{v.Value.(string)}}

	case types.DefaultID:
		return &graphp.Value{&graphp.Value_DefaultVal{v.Value.(string)}}

	default:
		// A type that isn't supported in the proto
		return nil
	}
}
