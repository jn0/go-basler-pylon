package pylon

// #include <string.h>
// #include "capture.h"
import "C"

import (
	"fmt"
	"sync"
)

type Camera struct {
	startMutex sync.Mutex

	attached, opened         bool
	attachedMutex, openMutex sync.Mutex
}

func (c *Camera) OpenCamera() error {
	c.openMutex.Lock()
	defer c.openMutex.Unlock()
	if !c.opened {
		if _, e := C.openCamera(); e != nil {
			return e
		}
		c.opened = true
	}
	return nil
}

func (c *Camera) CloseCamera() error {
	c.openMutex.Lock()
	defer c.openMutex.Unlock()
	if !c.opened {
		if _, e := C.closeCamera(); e != nil {
			return e
		}
		c.opened = false
	}
	return nil
}

func (c *Camera) StartCapture() error {
	c.startMutex.Lock()
	defer c.startMutex.Unlock()
	if !C.isCameraGrabbing() {
		if s, e := C.startCapture(); e != nil {
			return e
		} else if errMsg := C.GoString(s); errMsg != "" {
			return fmt.Errorf(errMsg)
		}
	}
	return nil
}

func (c *Camera) StopCapture() error {
	c.startMutex.Lock()
	defer c.startMutex.Unlock()
	if C.isCameraGrabbing() {
		if _, e := C.stopCapture(); e != nil {
			return e
		}

	}
	return nil
}

func (c *Camera) AttachDevice() error {
	fmt.Println("Attach device")
	c.attachedMutex.Lock()
	defer c.attachedMutex.Unlock()
	if !c.attached {
		fmt.Println("Attach device do it")
		if x, e := C.attachDevice(); e != nil {
			fmt.Println("Attach device do it err", e)
			return e
		} else {
			fmt.Println(x)
		}
		c.attached = true
	}
	return nil
}

func (c *Camera) ConfigureCamera() error {
	_, e := C.configureCamera()
	return e
}

func (c *Camera) Grab(batch, timeout int, outputPath string) error {
	if s, e := C.grab(C.int(batch), C.int(timeout), C.CString(outputPath)); e != nil {
		return e
	} else if errMsg := C.GoString(s); errMsg != "" {
		return fmt.Errorf(errMsg)
	}
	return nil
}
