#include <cpp_capture.h>
void startCapture(int batch, char* outputPath) {
	startCaptureCPP(batch,outputPath);
}

void stopCapture() {
	stopCaptureCPP();
}
void attachDevice() {
	attachDeviceCPP();
}
void configureCamera() {
	configureCameraCPP();
}
