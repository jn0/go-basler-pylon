package pylon

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func TestStart(t *testing.T) {
	cam := &Camera{}
	imgPath := path.Join(os.TempDir(), "go-basler-pylon-test")
	os.Remove(imgPath)

	if e := os.MkdirAll(imgPath, 0777); e != nil {
		t.Fatal(e)
	}

	t.Log(cam.AttachDevice())
	t.Log(cam.OpenCamera())
	t.Log(cam.StartCapture())
	t.Log(cam.StopCapture())
	if e := cam.Grab(1, 50000, imgPath+"/"); e != nil {
		t.Fatal(e)
	}

	if e := cam.AttachDevice(); e != nil {
		t.Fatal(e)
	}

	if e := cam.OpenCamera(); e != nil {
		t.Fatal(e)
	}

	if e := cam.StartCapture(); e != nil {
		t.Fatal(e)
	}

	defer func() {
		if e := cam.StopCapture(); e != nil {
			t.Fatal(e)
		}
	}()
	if e := cam.Grab(1, 5000, imgPath+"/"); e != nil {
		t.Fatal(e)
	}

	f, e := os.Open(imgPath)
	if e != nil {
		t.Fatal(e)
	} else if n, e := f.Readdirnames(-1); e != nil {
		t.Fatal(e)
	} else if len(n) == 0 {
		t.Fatalf("No scans taken.")
	} else {
		t.Logf("Scanned %s", n[0])
	}
}
