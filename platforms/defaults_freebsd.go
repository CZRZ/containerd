/*
   Copyright The containerd Authors.

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

package platforms

import (
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"runtime"
)

// DefaultSpec returns the current platform's default platform specification.
func DefaultSpec() specs.Platform {
	return specs.Platform{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		// The Variant field will be empty if arch != ARM.
		Variant: cpuVariant(),
	}
}

// Default returns the default matcher for the platform.
func Default() MatchComparer {
	return Ordered(DefaultSpec(), specs.Platform{
		OS:           "linux",
		Architecture: runtime.GOARCH,
		// The Variant field will be empty if arch != ARM.
		Variant: cpuVariant(),
	})
}
