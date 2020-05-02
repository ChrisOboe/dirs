// Copyright (c) 2020 ChrisOboe
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

// +build windows

package dirs

import (
	"os"
)

func Get(appname string) Dirs {
	var d Dirs

	d.Home = os.Getenv("USERPROFILE") + "/"
	d.Temp = os.Getenv("TEMP") + "/" + appname + "/"
	d.Config = os.Getenv("APPDATA")+ "/" + appname + "/"
	d.Cache = d.Config + "/cache/"
	d.Data = d.Config + "/data/"
    d.Runtime = d.Data

	return d
}
