package pylon

import (
	"os"
	"path"
	"testing"
)

func TestStart(t *testing.T) {
	cam := &Camera{}
	imgPath := path.Join(os.TempDir(), "go-basler-pylon-test")
	os.RemoveAll(imgPath)

	if e := os.MkdirAll(imgPath, 0777); e != nil {
		t.Fatal(e)
	}

	cam.AttachDevice()
	cam.OpenCamera()
	defer cam.CloseCamera()
	cam.ConfigureCamera()

	if e := cam.StartCapture(); e != nil {
		t.Fatal(e)
	}
	defer cam.StopCapture()

	if e := cam.RetrieveAndSave(1, 5000, imgPath+"/"); e != nil {
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

func TestHardwareTrigger(t *testing.T) {
	cam := &Camera{}
	imgPath := path.Join(os.TempDir(), "go-basler-pylon-test")
	os.RemoveAll(imgPath)

	if e := os.MkdirAll(imgPath, 0777); e != nil {
		t.Fatal(e)
	}

	cam.AttachDevice()
	cam.OpenCamera()
	defer cam.CloseCamera()
	if e := cam.SetHardwareTriggerConfiguration(); e != nil {
		t.Fatal(e)
	}
	if e := cam.ConfigureCamera(); e != nil {
		t.Fatal(e)
	}

	if e := cam.StartCapture(); e != nil {
		t.Fatal(e)
	}
	defer cam.StopCapture()

	if e := cam.RetrieveAndSave(1, 1000, imgPath+"/"); e == nil {
		t.Fatalf("Expected timeout.")
	}
}

func TestParamGroup(t *testing.T) {
	expected := "This category includes items used to conduct file operations"
	if s := FileAccessControl.Description(); s != expected {
		t.Fatalf("Expected description '%s' and go '%s'.", expected, s)
	}

	expected = "FileAccessControl"
	if s := FileAccessControl.String(); s != expected {
		t.Fatalf("Expected description '%s' and go '%s'.", expected, s)
	}

}
