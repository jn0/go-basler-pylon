package pylon

// #include <string.h>
// #include <stdlib.h>
// #include "capture.h"
import "C"

import (
	"fmt"
	"sync"
	"unsafe"
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

func (cam *Camera) SetHardwareTriggerConfiguration() error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return fmt.Errorf("Camera is not open.")
	}
	C.setHardwareTriggerConfiguration()
	return nil
}

func (cam *Camera) RetrieveAndSave(batch, timeout int, outputPath string) error {
	cOutputPath := C.CString(outputPath)
	defer C.free(unsafe.Pointer(cOutputPath))
	s := C.retrieveAndSave(C.int(batch), C.int(timeout), cOutputPath)
	if errMsg := C.GoString(s); errMsg != "" {
		return fmt.Errorf(errMsg)
	}
	return nil
}

func (cam *Camera) setNodeMapEnum(name IEnumerationT, value interface{}) error {
	cName := C.CString(string(name))
	defer C.free(unsafe.Pointer(cName))
	if v, ok := value.(string); !ok {
		return fmt.Errorf("Expected string value, got %T", value)
	} else {
		cValue := C.CString(v)
		defer C.free(unsafe.Pointer(cValue))
		C.setNodeMapEnumParam(cName, cValue)
	}
	return nil
}

func (cam *Camera) setNodeMapFloat(name IEnumerationT, value interface{}) error {
	cName := C.CString(string(name))
	defer C.free(unsafe.Pointer(cName))
	if v, ok := value.(float64); !ok {
		return fmt.Errorf("Expected float64 value, got %T", value)
	} else {
		cValue := C.double(v)
		C.setNodeMapFloatParam(cName, cValue)
	}
	return nil
}

func (cam *Camera) setNodeMapInt(name IInteger, value interface{}) error {
	cName := C.CString(string(name))
	defer C.free(unsafe.Pointer(cName))
	if v, ok := value.(int); !ok {
		return fmt.Errorf("Expected int value, got %T", value)
	} else {
		cValue := C.int(v)
		C.setNodeMapIntParam(cName, cValue)
	}
	return nil
}
