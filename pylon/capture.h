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
const char* configureCamera();
const char* retrieveAndSave(int batch, int timeout, char* outputPath);
const char* startCapture();
void openCamera();
void closeCamera();
bool isCameraGrabbing();
bool isAttached();
bool isOpen();
void setHardwareTriggerConfiguration();
void setNodeMapIntParam(char* name, int value);
void setNodeMapFloatParam(char* name, double value);
void setNodeMapEnumParam(char* name, char* value);

#ifdef __cplusplus
}
#endif

#endif // CAPTURE_H
