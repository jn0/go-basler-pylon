package pylon

// #include <string.h>
// #include "capture.h"
import "C"

import (
	"sync"
)

type Camera struct {
	start      bool
	startMutex sync.Mutex
}

func (c *Camera) StartCapture(batchId int, outputPath string) (e error) {
	c.startMutex.Lock()
	if !c.start {
		c.start = true
		c.startMutex.Unlock()
		_, e = C.startCapture(C.int(batchId), C.CString(outputPath))
	} else {
		c.startMutex.Unlock()
	}
	return
}

func (c *Camera) StopCapture() (e error) {
	c.startMutex.Lock()
	defer c.startMutex.Unlock()
	if c.start {
		_, e = C.stopCapture()

		// If error assume stop unsuccessful
		c.start = e != nil
	}
	return
}

func (c *Camera) AttachDevice() error {
	_, e := C.attachDevice()
	return e
}

func (c *Camera) ConfigureCamera() error {
	_, e := C.configureCamera()
	return e
}

func (c *Camera) batchCaptured() (int, error) {
	i, e := C.batchCaptured()
	return int(i), e
}
func (c *Camera) totalCaptured() (int, error) {
	i, e := C.totalCaptured()
	return int(i), e
}

func (c *Camera) Started() bool {
	return c.start
}
