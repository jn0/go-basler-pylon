#ifndef CAPTURE_H
#define CAPTURE_H

#ifdef __cplusplus 
extern "C" {
#else
#include <stdbool.h>
#endif

// prototypes
void stopCapture();
void attachDevice();
void configureCamera();
const char* grab(int batch, int timeout, char* outputPath);
const char* startCapture();
void openCamera();
void closeCamera();
bool isCameraGrabbing();

#ifdef __cplusplus
}
#endif

#endif // CAPTURE_H