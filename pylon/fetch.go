// Copyright (C) jno, 2019; License: MIT

// As the `RetrieveAndSave` does not work for me.
// So, I follow the code from // $PYLON_HOME/Samples/C++/Grab/Grab.cpp - it works.
// Plus, I prefer to handle grabbed image on "application level" and not
// bury it hardcoded somewhere in C++ underpants.
// So, the callback implementation.
package pylon

// #include <string.h>
// #include <stdlib.h>
// #include "capture.h"
import "C"

import "unsafe"
import "fmt"

//export Go_fetch_callback
func Go_fetch_callback(i C.int,				// callback index
		       w C.int, h C.int, pxt C.int,	// image props
		       size C.int, buffer *C.char,	// image data
) C.int {
	W := int(w)
	H := int(h)
	P := int(pxt)
	Z := int(size)
	B := C.GoBytes(unsafe.Pointer(buffer), size)

	fmt.Printf("\ncallback(w=%#v, h=%#v, pxt=%#v, size=%#v, buf=%#v)\n",
		   W, H, P, Z, B[0])

	f := cbreg.Get(int(i))
	rc := C.int(f(W, H, P, Z, B))
	return rc
}

func (cam *Camera) SetFetchTimeout(v uint) {
	C.setFetchTimeout(C.uint(v))
}

func (cam *Camera) SetFetchCount(v uint) {
	C.setFetchCount(C.uint(v))
}

func (cam *Camera) Fetch(cb func(w, h, pxt, size int, buffer []byte) int) error {
	idx := cbreg.Add(cb)
	s := C.GoString(C.fetch(C.int(idx)));
	if s != "" {
		return fmt.Errorf("Fetch: %v", s)
	}
	return nil
}

// EPixelType is an `enum` (i.e. `int`).
// It is supposed to have 32 bits like 0xabCD1234
// where 0xab000000 is "pixel class" (mono/color/custom),
// 0x00CD0000 is "pixel width" (bits per pixel), and
// 0x00001234 is "pixel subtype" (just serial number).
// See $PYLON_HOME/include/pylon/PixelType.h for details
type EPixelType struct {
	value int
	IsValid, IsUndefined, IsCustom, IsMono, IsColor bool
	Bits int
	Type int
}

func (pt *EPixelType) String() string {
	if pt.IsValid {
		if pt.IsUndefined {
			return "<EPixelType:UNDEFINED>"
		}
		var class string = ""
		if pt.IsCustom {
			class = "CUSTOM"
		}
		if pt.IsColor {
			if class != "" {
				class += ","
			}
			class += "COLOR"
		}
		if pt.IsMono {
			if class != "" {
				class += ","
			}
			class += "MONO"
		}
		return fmt.Sprintf("<EPixelType:%s(%d)%04x>",
				   class, pt.Bits, pt.Type)
	}
	return fmt.Sprintf("<EPixelType(%08x)Invalid>", pt.value)
}

func (pt *EPixelType) Set(v int) {
	pt.value = v
	pt.IsValid = false
	pt.IsUndefined = false
	pt.IsCustom = false
	pt.IsMono = false
	pt.IsColor = false
	pt.Bits = 0
	pt.Type = 0
	if v == -1 {
		pt.IsValid = true
		pt.IsUndefined = true
		return
	}
	pt.Bits = (v >> 16) & 0x0ff
	pt.Type = v & 0x0ffff

	class := int((v >> 24) & 0x0ff)
	if class & 0x01 == 0x01 {
		pt.IsMono = true
	}
	if class & 0x02 == 0x02 {
		pt.IsColor = true
	}
	if class & 0x80 == 0x80 {
		pt.IsCustom = true
	}
	pt.IsValid = class & ^(0x01 | 0x02 | 0x80) == 0
}

/* EOF */
