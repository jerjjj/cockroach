// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tests

import (
	"math/rand"

	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/asim/gen"
	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/asim/state"
)

// randomClusterInfoGen returns a randomly picked predefined configuration.
func (f randTestingFramework) randomClusterInfoGen(randSource *rand.Rand) gen.LoadedCluster {
	chosenIndex := randSource.Intn(len(state.ClusterOptions))
	chosenType := state.ClusterOptions[chosenIndex]
	return loadClusterInfo(chosenType)
}
