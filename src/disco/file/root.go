// Copyright (C) 2013 Mark Stahl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package file

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

func CreateDiscoRoot() {
	if rootDoesNotExist() {
		user, _ := user.Current()

		rootDir := path.Join(user.HomeDir, "/.disco.root")
		err := os.Mkdir(rootDir, 0700)
		if err != nil {
			fmt.Printf("error creating ~/.disco.root: %s", err)
		}
	}
}

func rootDoesNotExist() bool {
	user, _ := user.Current()
	rootDir := path.Join(user.HomeDir, "/.disco.root")

	if _, err := os.Stat(rootDir); err != nil {
		if os.IsNotExist(err) {
			return true
		}
	}

	return false
}
