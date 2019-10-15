#ifndef CAPTURE_H
#define CAPTURE_H

#ifdef __cplusplus 
extern "C" {
#else
#include <stdbool.h>
#endif

// prototypes
const char* stopCapture();
const char* attachDevice();
const char* configureCamera();
const char* retrieveAndSave(int batch, int timeout, char* outputPath);
const char* startCapture(int max);
const char* openCamera();
const char* closeCamera();
bool isCameraGrabbing();
bool isAttached();
bool isOpen();
const char* setHardwareTriggerConfiguration();
const char* setNodeMapIntParam(char* name, int value);
const char* setNodeMapFloatParam(char* name, double value);
const char* setNodeMapEnumParam(char* name, char* value);

const char* fullName();
const char* vendorName();
const char* modelName();
const char* serialNumber();
const char* deviceVersion();
int productId();
int vendorId();

int width();
int height();

#ifdef __cplusplus
}
#endif

#endif // CAPTURE_H
