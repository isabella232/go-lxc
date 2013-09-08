/*
 * console.go
 *
 * Copyright © 2013, S.Çağlar Onur
 *
 * Authors:
 * S.Çağlar Onur <caglar@10ur.org>
 *
 * This library is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 2, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package main

import (
	"flag"
	"fmt"
	"github.com/caglar10ur/lxc"
)

var (
	name string
)

func init() {
	flag.StringVar(&name, "name", "rubik", "Name of the container")
	flag.Parse()
}

func main() {
	c := lxc.NewContainer(name)
	defer lxc.PutContainer(c)

	if c.Defined() {
		if c.Running() {
			fmt.Printf("Attaching to container's console...\n")
			c.Console(-1, 0, 1, 2, 1)
		} else {
			fmt.Printf("Container is not running...\n")
		}
	} else {
		fmt.Printf("No such container...\n")
	}
}