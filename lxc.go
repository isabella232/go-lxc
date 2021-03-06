// Copyright © 2013, 2014, S.Çağlar Onur
// Use of this source code is governed by a LGPLv2.1
// license that can be found in the LICENSE file.
//
// Authors:
// S.Çağlar Onur <caglar@10ur.org>

// +build linux

// Package lxc provides Go Bindings for LXC (Linux Containers) C API.
package lxc

// #cgo pkg-config: lxc
// #include <lxc/lxccontainer.h>
// #include "lxc.h"
import "C"

import "unsafe"

// NewContainer returns a new container struct.
func NewContainer(name string, lxcpath ...string) (*Container, error) {
	var container *C.struct_lxc_container

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	if lxcpath != nil && len(lxcpath) == 1 {
		clxcpath := C.CString(lxcpath[0])
		defer C.free(unsafe.Pointer(clxcpath))

		container = C.lxc_container_new(cname, clxcpath)
	} else {
		container = C.lxc_container_new(cname, nil)
	}

	if container == nil {
		return nil, ErrNewFailed
	}
	return &Container{container: container, verbosity: Quiet}, nil
}

// GetContainer increments the reference counter of the container object.
func GetContainer(c *Container) bool {
	return C.lxc_container_get(c.container) == 1
}

// PutContainer decrements the reference counter of the container object.
func PutContainer(c *Container) bool {
	return C.lxc_container_put(c.container) == 1
}

// Version returns the LXC version.
func Version() string {
	return C.GoString(C.lxc_get_version())
}

// GlobalConfigItem returns the value of the given global config key.
func GlobalConfigItem(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return C.GoString(C.lxc_get_global_config_item(cname))
}

// DefaultConfigPath returns default config path.
func DefaultConfigPath() string {
	return GlobalConfigItem("lxc.lxcpath")
}

// DefaultLvmVg returns the name of the default LVM volume group.
func DefaultLvmVg() string {
	return GlobalConfigItem("lxc.bdev.lvm.vg")
}

// DefaultZfsRoot returns the name of the default ZFS root.
func DefaultZfsRoot() string {
	return GlobalConfigItem("lxc.bdec.zfs.root")
}

// ContainerNames returns the names of defined and active containers on the system.
func ContainerNames(lxcpath ...string) []string {
	var size int
	var cnames **C.char

	if lxcpath != nil && len(lxcpath) == 1 {
		clxcpath := C.CString(lxcpath[0])
		defer C.free(unsafe.Pointer(clxcpath))

		size = int(C.list_all_containers(clxcpath, &cnames, nil))
	} else {

		size = int(C.list_all_containers(nil, &cnames, nil))
	}

	if size < 1 {
		return nil
	}
	return convertNArgs(cnames, size)
}

// Containers returns the defined and active containers on the system. Only
// containers that could retrieved successfully are returned.
func Containers(lxcpath ...string) []Container {
	var containers []Container

	for _, v := range ContainerNames(lxcpath...) {
		if container, err := NewContainer(v, lxcpath...); err == nil {
			containers = append(containers, *container)
		}
	}

	return containers
}

// DefinedContainerNames returns the names of the defined containers on the system.
func DefinedContainerNames(lxcpath ...string) []string {
	var size int
	var cnames **C.char

	if lxcpath != nil && len(lxcpath) == 1 {
		clxcpath := C.CString(lxcpath[0])
		defer C.free(unsafe.Pointer(clxcpath))

		size = int(C.list_defined_containers(clxcpath, &cnames, nil))
	} else {

		size = int(C.list_defined_containers(nil, &cnames, nil))
	}

	if size < 1 {
		return nil
	}
	return convertNArgs(cnames, size)
}

// DefinedContainers returns the defined containers on the system.  Only
// containers that could retrieved successfully are returned.
func DefinedContainers(lxcpath ...string) []Container {
	var containers []Container

	for _, v := range DefinedContainerNames(lxcpath...) {
		if container, err := NewContainer(v, lxcpath...); err == nil {
			containers = append(containers, *container)
		}
	}

	return containers
}

// ActiveContainerNames returns the names of the active containers on the system.
func ActiveContainerNames(lxcpath ...string) []string {
	var size int
	var cnames **C.char

	if lxcpath != nil && len(lxcpath) == 1 {
		clxcpath := C.CString(lxcpath[0])
		defer C.free(unsafe.Pointer(clxcpath))

		size = int(C.list_active_containers(clxcpath, &cnames, nil))
	} else {

		size = int(C.list_active_containers(nil, &cnames, nil))
	}

	if size < 1 {
		return nil
	}
	return convertNArgs(cnames, size)
}

// ActiveContainers returns the active containers on the system. Only
// containers that could retrieved successfully are returned.
func ActiveContainers(lxcpath ...string) []Container {
	var containers []Container

	for _, v := range ActiveContainerNames(lxcpath...) {
		if container, err := NewContainer(v, lxcpath...); err == nil {
			containers = append(containers, *container)
		}
	}

	return containers
}
