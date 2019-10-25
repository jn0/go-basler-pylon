// Copyright (C) jno, 2019; License: MIT

// Save to .JPEG using vanilla Go library
package pylon

import (
	"os"
	"fmt"
	"bytes"
	"image"
	"image/jpeg"
	"time"
	"github.com/dsoprea/go-exif"
)

// save to JPEG using Go's builtin facility
var GoJpegOptions = jpeg.Options{ Quality: jpeg.DefaultQuality } // sorta const, yeah

func writeTo(path string, data []byte) error {
	of, e := os.Create(path)
	if e != nil {
		return fmt.Errorf("Cannot Create(%#v): %v", path, e)
	}
	defer of.Close()
	_, e = of.Write(data)
	if e != nil {
		return fmt.Errorf("Cannot Write(...%d...) to %q: %v",
				len(data), path, e)
	}
	return nil
}

// makeExif(jpegData) --> (exifData, error)
func makeExif(jpg []byte) ([]byte, error) {
	var tagList = ExifTagList{
		&ExifTag{
			Path: "IFD0",
			Name: "DateTime",
			Value: exif.ExifFullTimestampString(time.Now().UTC()),
		},
		&ExifTag{
			Path: "IFD/Exif",
			Name: "UserComment",
			Value: exif.TagUnknownType_9298_UserComment{
				EncodingType: exif.TagUnknownType_9298_UserComment_Encoding_ASCII,
				EncodingBytes: []byte("TEST COMMENT"),
			},
		},
	}

	if injector, e := NewExifInjector(jpg); e != nil {
		return nil, fmt.Errorf("NewExifInjector: %v", e)
	} else {
		injector.AddTagList(&tagList)

		if e = injector.Inject(); e != nil {
			return nil, fmt.Errorf("Cannot Inject: %v", e)
		}

		return injector.Jpeg, nil
	}

	return nil, nil
}

// go_save2jpeg: native Go implementation
func go_save2jpeg(path string, width, height int, data []byte) error {
	var e error
	a := image.NewGray(image.Rect(0, 0, width, height))
	a.Pix = data

	/*
	e = jpeg.Encode(of, a, &GoJpegOptions)
	if e != nil {
		return fmt.Errorf("Cannot Encode() to %#v: %v", path, e)
	}
	*/

	var tmp bytes.Buffer
	e = jpeg.Encode(&tmp, a, &GoJpegOptions)
	if e != nil {
		return fmt.Errorf("Cannot Encode(): %v", e)
	}
	jpg := tmp.Bytes()

	exif, e := makeExif(jpg)
	if e != nil {
		return fmt.Errorf("Cannot makeExif(): %v", e)
	}
	fmt.Printf("exif(%d): %v\n", len(exif), exif[:32])

	e = writeTo(path, jpg)
	if e != nil {
		return fmt.Errorf("Cannot writeTo(%q): %v", path, e)
	}

	e = writeTo(path + ".jpg", exif)
	if e != nil {
		return fmt.Errorf("Cannot writeTo(%q): %v", path + ".jpg", e)
	}

	fmt.Printf("Written to %#v\n", path)
	return nil
}

/* EOF */
