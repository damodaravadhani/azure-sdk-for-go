//go:build go1.18 && (linux || darwin || dragonfly || freebsd || openbsd || netbsd || solaris)
// +build go1.18
// +build linux darwin dragonfly freebsd openbsd netbsd solaris

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package blockblob

import (
	"fmt"
	"os"
	"syscall"
)

// mmb is a memory mapped buffer
type mmb []byte

// newMMB creates a new memory mapped buffer with the specified size
func newMMB(size int64) (mmb, error) {
	prot, flags := syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE
	addr, err := syscall.Mmap(-1, 0, int(size), prot, flags)
	if err != nil {
		return nil, os.NewSyscallError("Mmap", err)
	}
	return mmb(addr), nil
}

// delete cleans up the memory mapped buffer
func (m *mmb) delete() {
	err := syscall.Munmap(*m)
	*m = nil
	if err != nil {
		// if we get here, there is likely memory corruption.
		// please open an issue https://github.com/Azure/azure-sdk-for-go/issues
		panic(fmt.Sprintf("Munmap error: %v", err))
	}
}
