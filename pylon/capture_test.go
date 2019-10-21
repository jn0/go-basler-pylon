package pylon

import (
	"os"
	"fmt"
	"path"
	"path/filepath"
	"time"
	"image"
	"image/jpeg"
	"gopkg.in/gographics/imagick.v2/imagick"
	"testing"
)

func try(t *testing.T, e error, args ...string) {
	if len(args) > 1 { panic("try(*testing.T, error[, \"format\"])"); }
	format := "Failed in %v"
	if len(args) > 0 { format = args[0]; }
	if e != nil { t.Fatalf(format, e); }
}

// save to JPEG using Go's builtin facility
var GoJpegOptions = jpeg.Options{ Quality: jpeg.DefaultQuality } // sorta const, yeah
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

func time2name(t time.Time, suffix string) string {
	nsecs := t.Unix() * 1000000000 + int64(t.Nanosecond())
	return fmt.Sprintf("%x%s", nsecs / 1000, suffix)
}

func TestStart(t *testing.T) {
	imagick.Initialize(); defer imagick.Terminate()
	imName, imVer := imagick.GetVersion()
	fmt.Printf("About to use ImageMagick %#v (%v)\n\n", imName, imVer)
	mw := imagick.NewMagickWand(); defer mw.Destroy()

	cam := &Camera{}
	imgPath := path.Join(os.TempDir(), "go-basler-pylon-test")
	os.RemoveAll(imgPath)

	try(t, os.MkdirAll(imgPath, 0777))

	try(t, cam.AttachDevice())
	try(t, cam.OpenCamera())
	defer cam.CloseCamera()
	try(t, cam.ConfigureCamera())

	i, e := cam.Info(); try(t, e)
	t.Logf("camera USB[%04x:%04x] is %d×%d %#v by %#v, %#v, serial %#v, version %#v",
		i.VendorId, i.ProductId,
		i.Width, i.Height, i.ModelName, i.VendorName,
		i.FullName, i.SerialNumber, i.DeviceVersion)

	var saved []string

	FrameCallback := func(w, h, pxt, size int, buffer []byte) int {
		pt := EPixelType(pxt)
		fmt.Printf("FrameCallback(w=%#v, h=%#v, pt=%08x=%s, size=%#v, buffer=%#v...)\n",
			   w, h, pxt, pt.String(), size, buffer[0])
		t1 := time.Now()
		// DO STUFF FROM HERE

		path := filepath.Join(imgPath, time2name(t1.UTC(), ".jpg"))
		// try(t, go_save2jpeg(path, w, h, buffer))
		try(t, mw.ImportImagePixels(0, 0, // start
					    uint(w), uint(h), // size
					    "I", imagick.PIXEL_CHAR, // type
					    buffer)) // data
		mw.ResetIterator()
		try(t, mw.SetImageFormat("JPEG"))
		try(t, mw.WriteImage(path))
		mw.Clear()
		saved = append(saved, path)

		// UNTIL HERE
		fmt.Printf("FrameCallback taken %v\n\n", time.Now().Sub(t1))
		return 0
	}

	cam.SetFetchTimeout(5000) // ms
	cam.SetFetchCount(10)
	cb := FrameCallbackType(FrameCallback)
	try(t, cam.Fetch(cb), "Fetch failed: %v")

	if len(saved) > 0 {
		i := len(saved) / 2
		try(t, mw.ReadImage(saved[i]))
		w := mw.GetImageWidth()
		h := mw.GetImageHeight()
		try(t, mw.ResizeImage(w/2, h/2, imagick.FILTER_LANCZOS, 1))
		try(t, mw.SetImageCompressionQuality(75))
		try(t, mw.DisplayImage(os.Getenv("DISPLAY")))
	}

	if f, e := os.Open(imgPath); e != nil {
		t.Fatalf("os.Open(%#v) failed: %v", imgPath, e)
	} else if n, e := f.Readdirnames(-1); e != nil {
		t.Fatalf("Readdirnames failed: %v", e)
	} else if len(n) == 0 {
		t.Fatalf("No scans taken.")
	} else {
		t.Logf("Scanned %s", n[0])
	}
}
/*
func TestHardwareTrigger(t *testing.T) {
	cam := &Camera{}
	imgPath := path.Join(os.TempDir(), "go-basler-pylon-test")
	os.RemoveAll(imgPath)

	if e := os.MkdirAll(imgPath, 0777); e != nil {
		t.Fatalf("Failed with %v", e)
	}

	cam.AttachDevice()
	cam.OpenCamera()
	defer cam.CloseCamera()
	if e := cam.SetHardwareTriggerConfiguration(); e != nil {
		t.Fatalf("Failed with %v", e)
	}
	if e := cam.ConfigureCamera(); e != nil {
		t.Fatalf("Failed with %v", e)
	}

	i := cam.Info()
	t.Logf("camera USB[%04x:%04x] is %d×%d %#v by %#v, %#v, serial %#v, version %#v",
		i.VendorId, i.ProductId,
		i.Width, i.Height, i.ModelName, i.VendorName,
		i.FullName, i.SerialNumber, i.DeviceVersion)

	if e := cam.StartCapture(500); e != nil {
		t.Fatalf("Failed with %v", e)
	}
	defer cam.StopCapture()

	if e := cam.RetrieveAndSave(1, 1000, imgPath+"/"); e == nil {
		t.Fatalf("Expected timeout.")
	}
}
*/
/*
func TestParamGroup(t *testing.T) {
	expected := "This category includes items used to conduct file operations"
	if s := FileAccessControl.Description(); s != expected {
		t.Fatalf("Expected description '%#v' and go '%#v'.", expected, s)
	}

	expected = "FileAccessControl"
	if s := FileAccessControl.String(); s != expected {
		t.Fatalf("Expected description '%#v' and go '%#v'.", expected, s)
	}

}
*/
/* EOF */
