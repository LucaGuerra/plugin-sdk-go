/*
Copyright (C) 2021 The Falco Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sdk

import "io"

// PluginState represents the state of a plugin returned by plugin_init().
type PluginState interface {
}

type ExtractRequests interface {
	ExtractRequests() ExtractRequestPool
	SetExtractRequests(ExtractRequestPool)
}

// LastError is a compasable interface wrapping the basic LastError and
// SetLastError methods. This is meant to be used as a standard
// container for the last error catched during the execution of a plugin.
type LastError interface {
	// LastError returns the last error occurred in the plugin.
	LastError() error
	//
	// SetLastError sets the last error occurred in the plugin.
	SetLastError(err error)
}

// Destroyer is an interface wrapping the basic Destroy method.
// Destroy deinitializes the resources opened or allocated by a plugin.
// This is meant to be used in plugin_destroy() to release the plugin's
// resources. The behavior of Destroy after the first call is undefined.
type Destroyer interface {
	Destroy()
}

// Stringer is an interface wrapping the basic String method.
// String takes a io.ReadSeeker byte input and produces a string representation
// describing its internal data. This is meant to be used in
// plugin_event_to_string(), where the byte input represents event data
// provided by the framework and previouly produced by the plugin_next_batch().
type Stringer interface {
	String(in io.ReadSeeker) (string, error)
}

// Extractor is an interface wrapping the basic Extract method.
// Extract is meant to be used in plugin_extract_fields() to extract the value
// of a single field from a given event data.
type Extractor interface {
	Extract(req ExtractRequest, evt EventReader) error
}
