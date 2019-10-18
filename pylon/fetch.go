// Copyright (C) jno, 2019; License: MIT

// As the `RetrieveAndSave` does not work for me.
// So, I follow the code from $PYLON_HOME/Samples/C++/Grab/Grab.cpp - it works.
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

	if false {
		fmt.Printf("\ncallback(w=%#v, h=%#v, pxt=%#v, size=%#v, buf=%#v)\n",
			   W, H, P, Z, B[0])
	}
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

/* EOF */
