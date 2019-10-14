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

type CameraInfo struct {
	FullName, VendorName, ModelName string
	Width, Height int
}

type Camera struct {
	startMutex, attachedMutex, openMutex sync.Mutex
}

func (cam *Camera) Info() *CameraInfo {
	var i *CameraInfo = new(CameraInfo)
	i.FullName = C.GoString(C.fullName())
	i.VendorName = C.GoString(C.vendorName())
	i.ModelName = C.GoString(C.modelName())
	i.Width = int(C.width())
	i.Height = int(C.height())
	return i
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

func (cam *Camera) StartCapture(max int) error {
	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if !C.isCameraGrabbing() {
		s := C.startCapture(C.int(max))
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

	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return fmt.Errorf("Camera is grabbing.")
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

	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return fmt.Errorf("Camera is grabbing.")
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

func (cam *Camera) SetParam(p Param, value interface{}) error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return fmt.Errorf("Camera is not open.")
	}

	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return fmt.Errorf("Camera is grabbing.")
	}

	cName := C.CString(p.Name)
	defer C.free(unsafe.Pointer(cName))

	switch v := value.(type) {
	case string:
		// Multiple types for a string
		switch p.OriginalType {
		case OriginalTypeGenApiIEnumerationT:
			cValue := C.CString(v)
			defer C.free(unsafe.Pointer(cValue))
			C.setNodeMapEnumParam(cName, cValue)
		case OriginalTypeGenApiIString, OriginalTypeGenApiICommand:
			return fmt.Errorf("Original type %s not implemented.", p.OriginalType)
		default:
			return fmt.Errorf("Unexpected string for type %s", p.OriginalType)
		}

	case int64:
		if p.OriginalType != OriginalTypeGenApiIInteger {
			return fmt.Errorf("Unexpected int64 for type %s", p.OriginalType)
		}
		cValue := C.int(v)
		C.setNodeMapIntParam(cName, cValue)

	case float64:
		if p.OriginalType != OriginalTypeGenApiIFloat {
			return fmt.Errorf("Unexpected float64 for type %s", p.OriginalType)
		}

		cValue := C.double(v)
		C.setNodeMapFloatParam(cName, cValue)

	default:
		return fmt.Errorf("Value type %T of param %s not implemented.", value, p.Name)
	}
	return nil
}
