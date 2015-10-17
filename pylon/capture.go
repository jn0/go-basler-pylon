package pylon

// #include <string.h>
// #include "capture.h"
import "C"

import (
	"fmt"
	"sync"
)

type Camera struct {
	startMutex, attachedMutex, openMutex sync.Mutex
}

func (c *Camera) OpenCamera() {
	c.openMutex.Lock()
	defer c.openMutex.Unlock()
	if !C.isOpen() {
		C.openCamera()
	}
}

func (c *Camera) CloseCamera() {
	c.openMutex.Lock()
	defer c.openMutex.Unlock()
	if C.isOpen() {
		C.closeCamera()
	}
}

func (c *Camera) StartCapture() error {
	c.startMutex.Lock()
	defer c.startMutex.Unlock()
	if !C.isCameraGrabbing() {
		s := C.startCapture()
		if errMsg := C.GoString(s); errMsg != "" {
			return fmt.Errorf(errMsg)
		}
	}
	return nil
}

func (c *Camera) StopCapture() {
	c.startMutex.Lock()
	defer c.startMutex.Unlock()
	if C.isCameraGrabbing() {
		C.stopCapture()
	}
}

func (c *Camera) AttachDevice() {
	c.attachedMutex.Lock()
	defer c.attachedMutex.Unlock()
	if !C.isAttached() {
		C.attachDevice()
	}
}

func (c *Camera) ConfigureCamera() {
	C.configureCamera()
}

func (c *Camera) Grab(batch, timeout int, outputPath string) error {
	s := C.grab(C.int(batch), C.int(timeout), C.CString(outputPath))
	if errMsg := C.GoString(s); errMsg != "" {
		return fmt.Errorf(errMsg)
	}
	return nil
}
