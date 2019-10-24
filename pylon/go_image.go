// Copyright (C) jno, 2019; License: MIT

// Save to .JPEG using vanilla Go library
package pylon

import (
	"os"
	"fmt"
	"bytes"
	"time"
	"image"
	"image/jpeg"
	"github.com/dsoprea/go-exif"
	"github.com/dsoprea/go-jpeg-image-structure"
)

// save to JPEG using Go's builtin facility
var GoJpegOptions = jpeg.Options{ Quality: jpeg.DefaultQuality } // sorta const, yeah

func blankIfdBuilder() (*exif.IfdBuilder, error) {
	im := exif.NewIfdMapping()
	if e := exif.LoadStandardIfds(im); e != nil {
		return nil, fmt.Errorf("exif.LoadStandardIfds: %v", e)
	}
	ib := exif.NewIfdBuilder(
		im,
		exif.NewTagIndex(),
		exif.IfdPathStandard,
		exif.TestDefaultByteOrder,
	)
	return ib, nil
}

func makeExif(jpg []byte) ([]byte, error) {
	jmp := jpegstructure.NewJpegMediaParser()

	sl, e := jmp.ParseBytes(jpg)
	if e != nil {
		return nil, fmt.Errorf("Cannot ParseBytes: %v", e)
	}
	sl.Print()
	rootIb, e := sl.ConstructExifBuilder()
	if e != nil {
		if e.Error() == "no exif data" {
			rootIb, e = blankIfdBuilder()
			if e != nil {
				return nil, fmt.Errorf("blankIfdBuilder: %v", e)
			}
			e = sl.SetExif(rootIb)
			if e != nil {
				return nil, fmt.Errorf("sl.SetExif: %v", e)
			}
		} else {
			return nil, fmt.Errorf("Cannot ConstructExifBuilder: %v", e)
		}
	}
	ifdPath := "IFD0"
	ifdIb, e := exif.GetOrCreateIbFromRootIb(rootIb, ifdPath)
	if e != nil {
		return nil, fmt.Errorf("Cannot GetOrCreateIbFromRootIb: %v", e)
	}
	//
	now := time.Now().UTC()
	e = ifdIb.SetStandardWithName("DateTime",
					exif.ExifFullTimestampString(now))
	if e != nil {
		return nil, fmt.Errorf("Cannot SetStandardWithName: %v", e)
	}
	//
	e = sl.SetExif(rootIb)
	if e != nil {
		return nil, fmt.Errorf("Cannot SetExif: %v", e)
	}
	buf := new(bytes.Buffer)
	e = sl.Write(buf)
	if e != nil {
		return nil, fmt.Errorf("Cannot sl.Write: %v", e)
	}
	return buf.Bytes(), nil
}

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
