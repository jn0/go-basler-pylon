#ifndef CAPTURE_H
#define CAPTURE_H

#ifdef __cplusplus 
extern "C" {
#else
#include <stdbool.h>
#include <stdint.h>
#endif


// prototypes
void pylonInitialize();
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

// typedef int (*FrameCallbackType)(int w, int h, int pxt, int size, const char* buffer);
extern int Go_fetch_callback(int, int, int, int, int, char*);

void setFetchTimeout(uint32_t v);
void setFetchCount(uint32_t v);
const char* fetch(int idx);
#ifdef __cplusplus
}
#endif

#endif // CAPTURE_H
