// Copyright 2018 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package paths

import "runtime"

type PathConfig struct {
	// Whether to create the symlink in the new PATH for this tool.
	Symlink bool

	// Whether to log about usages of this tool to the soong.log
	Log bool

	// Whether to exit with an error instead of invoking the underlying tool.
	Error bool

	// Whether we use a linux-specific prebuilt for this tool. On Darwin,
	// we'll allow the host executable instead.
	LinuxOnlyPrebuilt bool
}

var Allowed = PathConfig{
	Symlink: true,
	Log:     false,
	Error:   false,
}

var Forbidden = PathConfig{
	Symlink: false,
	Log:     true,
	Error:   true,
}

var Log = PathConfig{
	Symlink: true,
	Log:     true,
	Error:   false,
}

// The configuration used if the tool is not listed in the config below.
// Currently this will create the symlink, but log and error when it's used. In
// the future, I expect the symlink to be removed, and this will be equivalent
// to Forbidden.
var Missing = PathConfig{
	Symlink: true,
	Log:     true,
	Error:   true,
}

var LinuxOnlyPrebuilt = PathConfig{
	Symlink:           false,
	Log:               true,
	Error:             true,
	LinuxOnlyPrebuilt: true,
}

func GetConfig(name string) PathConfig {
	if config, ok := Configuration[name]; ok {
		return config
	}
	return Missing
}

var Configuration = map[string]PathConfig{
	"bash":     Allowed,
	"bc":       Allowed,
	"bzip2":    Allowed,
	"date":     Allowed,
	"dd":       Allowed,
	"diff":     Allowed,
	"egrep":    Allowed,
	"expr":     Allowed,
	"find":     Allowed,
	"fuser":    Allowed,
	"getopt":   Allowed,
	"git":      Allowed,
	"grep":     Allowed,
	"gzip":     Allowed,
	"hexdump":  Allowed,
	"jar":      Allowed,
	"java":     Allowed,
	"javap":    Allowed,
	"lsof":     Allowed,
	"m4":       Allowed,
	"nproc":    Allowed,
	"openssl":  Allowed,
	"patch":    Allowed,
	"pstree":   Allowed,
	"python3":  Allowed,
	"realpath": Allowed,
	"rsync":    Allowed,
	"sed":      Allowed,
	"sh":       Allowed,
	"tar":      Allowed,
	"timeout":  Allowed,
	"tr":       Allowed,
	"unzip":    Allowed,
	"xz":       Allowed,
	"zip":      Allowed,
	"zipinfo":  Allowed,

	"aarch64-linux-android-addr2line":    Allowed,
	"aarch64-linux-android-ar":           Allowed,
	"aarch64-linux-android-as":           Allowed,
	"aarch64-linux-android-c++filt":      Allowed,
	"aarch64-linux-android-dwp":          Allowed,
	"aarch64-linux-android-elfedit":      Allowed,
	"aarch64-linux-android-gcc":          Allowed,
	"aarch64-linux-android-gcc-ar":       Allowed,
	"aarch64-linux-android-gcc-nm":       Allowed,
	"aarch64-linux-android-gcc-ranlib":   Allowed,
	"aarch64-linux-android-gcov":         Allowed,
	"aarch64-linux-android-gcov-tool":    Allowed,
	"aarch64-linux-android-gprof":        Allowed,
	"aarch64-linux-android-ld":           Allowed,
	"aarch64-linux-android-ld.bfd":       Allowed,
	"aarch64-linux-android-ld.gold":      Allowed,
	"aarch64-linux-android-nm":           Allowed,
	"aarch64-linux-android-objcopy":      Allowed,
	"aarch64-linux-android-objdump":      Allowed,
	"aarch64-linux-android-ranlib":       Allowed,
	"aarch64-linux-android-readelf":      Allowed,
	"aarch64-linux-android-size":         Allowed,
	"aarch64-linux-android-strings":      Allowed,
	"aarch64-linux-android-strip":        Allowed,
	"aarch64-linux-gnu-as":               Allowed,
	"arm-linux-androideabi-addr2line":    Allowed,
	"arm-linux-androideabi-ar":           Allowed,
	"arm-linux-androideabi-as":           Allowed,
	"arm-linux-androideabi-c++filt":      Allowed,
	"arm-linux-androideabi-cpp":          Allowed,
	"arm-linux-androideabi-dwp":          Allowed,
	"arm-linux-androideabi-elfedit":      Allowed,
	"arm-linux-androideabi-gcc":          Allowed,
	"arm-linux-androideabi-gcc-ar":       Allowed,
	"arm-linux-androideabi-gcc-nm":       Allowed,
	"arm-linux-androideabi-gcc-ranlib":   Allowed,
	"arm-linux-androideabi-gcov":         Allowed,
	"arm-linux-androideabi-gcov-tool":    Allowed,
	"arm-linux-androideabi-gprof":        Allowed,
	"arm-linux-androideabi-ld":           Allowed,
	"arm-linux-androideabi-ld.bfd":       Allowed,
	"arm-linux-androideabi-ld.gold":      Allowed,
	"arm-linux-androideabi-nm":           Allowed,
	"arm-linux-androideabi-objcopy":      Allowed,
	"arm-linux-androideabi-objdump":      Allowed,
	"arm-linux-androideabi-ranlib":       Allowed,
	"arm-linux-androideabi-readelf":      Allowed,
	"arm-linux-androideabi-size":         Allowed,
	"arm-linux-androideabi-strings":      Allowed,
	"arm-linux-androideabi-strip":        Allowed,
	"arm-linux-androidkernel-addr2line":  Allowed,
	"arm-linux-androidkernel-ar":         Allowed,
	"arm-linux-androidkernel-as":         Allowed,
	"arm-linux-androidkernel-c++filt":    Allowed,
	"arm-linux-androidkernel-cpp":        Allowed,
	"arm-linux-androidkernel-dwp":        Allowed,
	"arm-linux-androidkernel-elfedit":    Allowed,
	"arm-linux-androidkernel-gcc":        Allowed,
	"arm-linux-androidkernel-gcc-ar":     Allowed,
	"arm-linux-androidkernel-gcc-nm":     Allowed,
	"arm-linux-androidkernel-gcc-ranlib": Allowed,
	"arm-linux-androidkernel-gcov":       Allowed,
	"arm-linux-androidkernel-gcov-tool":  Allowed,
	"arm-linux-androidkernel-gprof":      Allowed,
	"arm-linux-androidkernel-ld":         Allowed,
	"arm-linux-androidkernel-ld.bfd":     Allowed,
	"arm-linux-androidkernel-ld.gold":    Allowed,
	"arm-linux-androidkernel-nm":         Allowed,
	"arm-linux-androidkernel-objcopy":    Allowed,
	"arm-linux-androidkernel-objdump":    Allowed,
	"arm-linux-androidkernel-ranlib":     Allowed,
	"arm-linux-androidkernel-readelf":    Allowed,
	"arm-linux-androidkernel-size":       Allowed,
	"arm-linux-androidkernel-strings":    Allowed,
	"arm-linux-androidkernel-strip":      Allowed,

	// Host toolchain is removed. In-tree toolchain should be used instead.
	// GCC also can't find cc1 with this implementation.
	"ar":         Forbidden,
	"as":         Forbidden,
	"cc":         Forbidden,
	"clang":      Forbidden,
	"clang++":    Forbidden,
	"gcc":        Forbidden,
	"g++":        Forbidden,
	"ld":         Forbidden,
	"ld.bfd":     Forbidden,
	"ld.gold":    Forbidden,
	"pkg-config": Forbidden,

	// On Linux we'll use the toybox versions of these instead.
	"basename":  LinuxOnlyPrebuilt,
	"cat":       LinuxOnlyPrebuilt,
	"chmod":     LinuxOnlyPrebuilt,
	"cmp":       LinuxOnlyPrebuilt,
	"cp":        LinuxOnlyPrebuilt,
	"comm":      LinuxOnlyPrebuilt,
	"cut":       LinuxOnlyPrebuilt,
	"dirname":   LinuxOnlyPrebuilt,
	"du":        LinuxOnlyPrebuilt,
	"echo":      LinuxOnlyPrebuilt,
	"env":       LinuxOnlyPrebuilt,
	"head":      LinuxOnlyPrebuilt,
	"getconf":   LinuxOnlyPrebuilt,
	"hostname":  LinuxOnlyPrebuilt,
	"id":        LinuxOnlyPrebuilt,
	"ln":        LinuxOnlyPrebuilt,
	"ls":        LinuxOnlyPrebuilt,
	"md5sum":    LinuxOnlyPrebuilt,
	"mkdir":     LinuxOnlyPrebuilt,
	"mktemp":    LinuxOnlyPrebuilt,
	"mv":        LinuxOnlyPrebuilt,
	"od":        LinuxOnlyPrebuilt,
	"paste":     LinuxOnlyPrebuilt,
	"pgrep":     LinuxOnlyPrebuilt,
	"pkill":     LinuxOnlyPrebuilt,
	"ps":        LinuxOnlyPrebuilt,
	"pwd":       LinuxOnlyPrebuilt,
	"readlink":  LinuxOnlyPrebuilt,
	"rm":        LinuxOnlyPrebuilt,
	"rmdir":     LinuxOnlyPrebuilt,
	"seq":       LinuxOnlyPrebuilt,
	"setsid":    LinuxOnlyPrebuilt,
	"sha1sum":   LinuxOnlyPrebuilt,
	"sha256sum": LinuxOnlyPrebuilt,
	"sha512sum": LinuxOnlyPrebuilt,
	"sleep":     LinuxOnlyPrebuilt,
	"sort":      LinuxOnlyPrebuilt,
	"stat":      LinuxOnlyPrebuilt,
	"tail":      LinuxOnlyPrebuilt,
	"tee":       LinuxOnlyPrebuilt,
	"touch":     LinuxOnlyPrebuilt,
	"true":      LinuxOnlyPrebuilt,
	"uname":     LinuxOnlyPrebuilt,
	"uniq":      LinuxOnlyPrebuilt,
	"unix2dos":  LinuxOnlyPrebuilt,
	"wc":        LinuxOnlyPrebuilt,
	"whoami":    LinuxOnlyPrebuilt,
	"which":     LinuxOnlyPrebuilt,
	"xargs":     LinuxOnlyPrebuilt,
	"xxd":       LinuxOnlyPrebuilt,
}

func init() {
	if runtime.GOOS == "darwin" {
		Configuration["md5"] = Allowed
		Configuration["sw_vers"] = Allowed
		Configuration["xcrun"] = Allowed

		// We don't have darwin prebuilts for some tools (like toybox),
		// so allow the host versions.
		for name, config := range Configuration {
			if config.LinuxOnlyPrebuilt {
				Configuration[name] = Allowed
			}
		}
	}
}
