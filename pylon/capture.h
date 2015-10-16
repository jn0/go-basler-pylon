#ifndef CAPTURE_H
#define CAPTURE_H

#ifdef __cplusplus 
extern "C" {
#endif

// prototypes
void startCapture(int batch, char* outputPath);
void stopCapture();
void attachDevice();
void configureCamera();
int batchCaptured();
int totalCaptured();

#ifdef __cplusplus
}
#endif

#endif // CAPTURE_H