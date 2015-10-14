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

func (c *Camera) StartCapture(batchId int, outputPath string) {
	c.startMutex.Lock()
	c.start = true
	c.startMutex.Unlock()
	C.startCapture(C.int(batchId), C.CString(outputPath))
}

func (c *Camera) StopCapture() {
	C.stopCapture()
	c.startMutex.Lock()
	defer c.startMutex.Unlock()
	c.start = false
}

func (c *Camera) AttachDevice() {
	C.attachDevice()
}

func (c *Camera) ConfigureCamera() {
	C.configureCamera()
}

func (c *Camera) Started() bool {
	c.startMutex.Lock()
	defer c.startMutex.Unlock()
	return c.start
}
