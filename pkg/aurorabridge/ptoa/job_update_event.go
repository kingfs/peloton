// Copyright (c) 2019 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ptoa

import (
	"fmt"
	"time"

	"github.com/uber/peloton/.gen/peloton/api/v1alpha/job/stateless"
	"github.com/uber/peloton/.gen/thrift/aurora/api"
	"github.com/uber/peloton/pkg/aurorabridge/opaquedata"

	"go.uber.org/thriftrw/ptr"
)

// NewJobUpdateEvent returns aurora job update event from
// peloton job update event
func NewJobUpdateEvent(
	e *stateless.WorkflowEvent,
	d *opaquedata.Data,
) (*api.JobUpdateEvent, error) {

	jobUpdateStatus, err := NewJobUpdateStatus(
		e.GetState(),
		d)
	if err != nil {
		return nil, fmt.Errorf("unable to parse job update event state %s", err)
	}

	t, err := time.Parse(time.RFC3339, e.GetTimestamp())
	if err != nil {
		return nil, fmt.Errorf("unable to parse job update event timestamp %s", err)
	}
	t64 := t.UnixNano() / int64(time.Millisecond)

	return &api.JobUpdateEvent{
		Status:      &jobUpdateStatus,
		TimestampMs: ptr.Int64(t64),
	}, nil
}
