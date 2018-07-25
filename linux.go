// Copyright (c) 2018 ChrisOboe
//
// This file is part of dirs
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// +build linux

package dirs

import (
	"os"
)

func Get(appname string) Dirs {
	var d Dirs

	d.Home = os.Getenv("HOME") + "/"
	d.Temp = "/tmp/" + appname + "/"

	if os.Getenv("XDG_CONFIG_HOME") != "" {
		d.Config = os.Getenv("XDG_CONFIG_HOME")
	} else {
		d.Config = d.Home + ".config"
	}
	d.Config = d.Config + "/" + appname + "/"

	if os.Getenv("XDG_CACHE_HOME") != "" {
		d.Cache = os.Getenv("XDG_CACHE_HOME")
	} else {
		d.Cache = d.Home + ".cache"
	}
	d.Cache = d.Cache + "/" + appname + "/"

	if os.Getenv("XDG_DATA_HOME") != "" {
		d.Data = os.Getenv("XDG_DATA_HOME")
	} else {
		d.Data = d.Home + ".local/share"
	}
	d.Data = d.Data + "/" + appname + "/"

	if os.Getenv("XDG_RUNTIME_DIR") != "" {
		d.Data = os.Getenv("XDG_RUNTIME_DIR") + "/" + appname + "/"
	} else {
		d.Data = d.Temp
	}

	return d
}
