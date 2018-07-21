/*
 * // Copyright (c) 2018 Cisco and/or its affiliates.
 * //
 * // Licensed under the Apache License, Version 2.0 (the "License");
 * // you may not use this file except in compliance with the License.
 * // You may obtain a copy of the License at:
 * //
 * //     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing, software
 * // distributed under the License is distributed on an "AS IS" BASIS,
 * // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * // See the License for the specific language governing permissions and
 * // limitations under the License.
 */

package cache

import "github.com/ligato/cn-infra/datasync"

// NodeTelemetryCacheAPI defines the API of NodeTelemetryCache used for
// a non-persistent storage of K8s State data with fast lookups.
// The cache processes K8s State data updates and RESYNC events through Update()
// and Resync() APIs, respectively.
// The cache allows to get notified about changes via convenient callbacks.
type NodeTelemetryCacheAPI interface {
	// Update processes a datasync change event associated with K8s State data.
	// The change is applied into the cache.
	Update(dataChngEv datasync.ChangeEvent) error

	// Resync processes a datasync resync event associated with K8s State data.
	// The cache content is full replaced with the received data and all
	// subscribed watchers are notified.
	// The function will forward any error returned by a watcher.
	Resync(resyncEv datasync.ResyncEvent) error

	// LookupNodeInfo returns data of a given node or list of nodes.
	// LookupNode() Node

	// LookupNodeInfo returns data of all nodes.
	// ListAllNodes() []Node

	// LookupNodeInfo returns data of a given pod or list of nodes.
	// LookupPod() Pod

	// LookupNodeInfo returns data of a given node or list of nodes.
	// LookupKsrInfo() KsrInfo

	// LookupNodeInfo returns data of a given node or list of nodes.
	// LookupEtcdInfo() EtcdInfo
}