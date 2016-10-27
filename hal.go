// Package hal separate depends on on real hardware and os, such as time, file,
// and other os operations. Use hal package mainly for easier to do unit test.
//
// Nearly all members in this package is unprotected alias of standard go
// functions, do not alter them unless inside unit tests.
package hal

import (
	"os"
	"time"
)

var (
	// Exit is alias of os.Exit
	Exit = os.Exit

	// Now is alias of time.Now
	Now = time.Now

	// Getenv is alias of os.Getenv
	Getenv = os.Getenv
)
