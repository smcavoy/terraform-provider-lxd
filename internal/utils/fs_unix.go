//go:build !windows

package utils

import (
	"log"

	"golang.org/x/sys/unix"
)

// IsSocketWritable returns true if user has write permissions for socket on the given path.
func IsSocketWritable(socketPath string) bool {
	err := unix.Access(socketPath, unix.W_OK)
	if err != nil {
		log.Printf("Unix socket %q: %v", socketPath, err)
		return false
	}

	return true
}
