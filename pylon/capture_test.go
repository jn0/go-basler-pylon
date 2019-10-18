package pylon

import (
	"os"
	"fmt"
	"path"
	"path/filepath"
	"time"
	"image"
	"image/jpeg"
	"testing"
)

func try(t *testing.T, e error, args ...string) {
	if len(args) > 1 { panic("try(*testing.T, error[, \"format\"])"); }
	format := "Failed with %v"
	if len(args) > 0 { format = args[0]; }
	if e != nil { t.Fatalf(format, e); }
}

func TestStart(t *testing.T) {
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

	FrameCallback := func(w, h, pxt, size int, buffer []byte) int {
		t1 := time.Now()
		pt := EPixelType(pxt)
		fmt.Printf("FrameCallback(w=%#v, h=%#v, pt=%08x=%s, size=%#v, buffer=%#v...)\n",
			   w, h, pxt, pt.String(), size, buffer[0])
		// DO STUFF HERE
		/*
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				i := y * w + x
				p := buffer[i]
			}
		}
		*/

		r := image.Rect(0, 0, w, h)
		a := image.NewGray(r)
		a.Pix = buffer

		nsecs := t1.UTC().Unix() * 1000000000 + int64(t1.UTC().Nanosecond())
		imgName := fmt.Sprintf("%x.jpg", nsecs / 1000)
		path := filepath.Join(imgPath, imgName)
		if of, e := os.Create(path); e != nil {
			t.Fatalf("Cannot Create(%#v): %v", path, e)
		} else {
			jpeg.Encode(of, a, nil)
			of.Close()
			fmt.Printf("Written to %#v\n", path)
		}

		t2 := time.Now()
		dt := t2.Sub(t1)
		fmt.Printf("FrameCallback taken %v\n", dt)
		return 0
	}

	cam.SetFetchTimeout(5000) // ms
	cam.SetFetchCount(10)
	cb := FrameCallbackType(FrameCallback)
	try(t, cam.Fetch(cb), "Fetch failed: %v")

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
