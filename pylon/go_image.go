// Copyright (C) jno, 2019; License: MIT

// Save to .JPEG using vanilla Go library
package pylon

import (
	"os"
	"fmt"
	"image"
	"image/jpeg"
)

// save to JPEG using Go's builtin facility
var GoJpegOptions = jpeg.Options{ Quality: jpeg.DefaultQuality } // sorta const, yeah

// go_save2jpeg: native Go implementation
func go_save2jpeg(path string, width, height int, data []byte) error {
	a := image.NewGray(image.Rect(0, 0, width, height))
	a.Pix = data

	if of, e := os.Create(path); e != nil {
		return fmt.Errorf("Cannot Create(%#v): %v", path, e)
	} else {
		e := jpeg.Encode(of, a, &GoJpegOptions)
		of.Close()
		if e != nil {
			return fmt.Errorf("Cannot Encode() to %#v: %v", path, e)
		}
		fmt.Printf("Written to %#v\n", path)
	}
	return nil
}

/* EOF */
