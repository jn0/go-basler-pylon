// Copyright (C) jno, 2019; License: MIT

/* Save to .JPEG using MagicWand from ImageMagick
 * Use as:
 *
 *	mw := im_setup(); defer im_cleanup(mw)
 *	...
 *	if e := im_save2jpeg(mw, ...); if e != nil { t.Fatalf("save: %v", e); }
 *	...
 *
 */
package pylon

import (
	"os"
	"fmt"
	"gopkg.in/gographics/imagick.v2/imagick"
)

const imageJpegQuality = 75

// im_save2jpeg: using ImageMagick
func im_save2jpeg(mw *imagick.MagickWand, path string, width, height int, data []byte) error {
	pw := imagick.NewPixelWand(); defer pw.Destroy()
	/* t.Logf("SetColor(gray): %v", */  pw.SetColor("gray") //)

	mw.ResetIterator()
	if e := mw.NewImage(uint(width), uint(height), pw); e != nil {
		return fmt.Errorf("NewImage: %v", e); }
	if e := mw.ImportImagePixels(0, 0, // start
				     uint(width), uint(height), // size
				     "I", imagick.PIXEL_CHAR, // type
				     data); // data
	  e != nil { return fmt.Errorf("ImportImagePixels: %v", e); }
	if e := mw.SetImageFormat("JPEG"); e != nil {
		return fmt.Errorf("SetImageFormat: %v", e); }
	if e := mw.WriteImage(path); e != nil {
		return fmt.Errorf("WriteImage: %v", e); }
	mw.Clear()
	return nil
}

// ImageMagick setup routine
func im_setup() *imagick.MagickWand {
	imagick.Initialize()
	imName, imVer := imagick.GetVersion()
	fmt.Printf("About to use ImageMagick %#v (%v)\n\n", imName, imVer)
	return imagick.NewMagickWand()
}
// ImageMagick cleanup routine
func im_cleanup(mw *imagick.MagickWand) {
	mw.Destroy()
	imagick.Terminate()
	fmt.Printf("Done with ImageMagick\n")
}

// im_show: utility function to display an image
func im_show(mw *imagick.MagickWand, path string) error {
	if e := mw.ReadImage(path); e != nil {
		return fmt.Errorf("ReadImage(%#v): %v", path, e)
	}
	w := mw.GetImageWidth()
	h := mw.GetImageHeight()
	if e := mw.ResizeImage(w/2, h/2, imagick.FILTER_LANCZOS, 1); e != nil {
		return fmt.Errorf("ResizeImage: %v", e)
	}
	if e := mw.SetImageCompressionQuality(imageJpegQuality); e != nil {
		return fmt.Errorf("SetImageCompressionQuality: %v", e)
	}
	if e := mw.DisplayImage(os.Getenv("DISPLAY")); e != nil {
		return fmt.Errorf("DisplayImage(%#v): %v", os.Getenv("DISPLAY"), e)
	}
	return nil
}

/* EOF */
