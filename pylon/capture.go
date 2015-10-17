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

func (cam *Camera) OpenCamera() {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		C.openCamera()
	}
}

func (cam *Camera) CloseCamera() {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if C.isOpen() {
		C.closeCamera()
	}
}

func (cam *Camera) StartCapture() error {
	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if !C.isCameraGrabbing() {
		s := C.startCapture()
		if errMsg := C.GoString(s); errMsg != "" {
			return fmt.Errorf(errMsg)
		}
	}
	return nil
}

func (cam *Camera) IsGrabbing() bool {
	if C.isCameraGrabbing() {
		return true
	}
	return false
}

func (cam *Camera) StopCapture() {
	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		C.stopCapture()
	}
}

func (cam *Camera) AttachDevice() {
	cam.attachedMutex.Lock()
	defer cam.attachedMutex.Unlock()
	if !C.isAttached() {
		C.attachDevice()
	}
}

func (cam *Camera) ConfigureCamera() error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return fmt.Errorf("Camera is not open.")
	}
	C.configureCamera()
	return nil
}

func (cam *Camera) SetHardwareTriggerConfiguration() {
	C.setHardwareTriggerConfiguration()
}

func (cam *Camera) RetrieveAndSave(batch, timeout int, outputPath string) error {
	s := C.retrieveAndSave(C.int(batch), C.int(timeout), C.CString(outputPath))
	if errMsg := C.GoString(s); errMsg != "" {
		return fmt.Errorf(errMsg)
	}
	return nil
}
