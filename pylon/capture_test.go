package pylon

import (
	"os"
	"path"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	cam := &Camera{}
	imgPath := path.Join(os.TempDir(), "go-basler-pylon-test")
	os.Remove(imgPath)

	if e := os.MkdirAll(imgPath, 0777); e != nil {
		t.Fatal(e)
	}

	go cam.StartCapture(1, imgPath+"/")

	for i := 0; !cam.Started(); i = i + 1 {
		time.Sleep(time.Second)
		if i > 10 {
			defer cam.StopCapture()
			t.Fatalf("Expected camera to start by now.")
		}
	}
	cam.StopCapture()
	f, e := os.Open(imgPath)
	if e != nil {
		t.Fatal(e)
	}
	if n, e := f.Readdirnames(-1); e != nil {
		t.Fatal(e)
	} else if len(n) == 0 {
		t.Fatalf("No scans taken.")
	}
}
