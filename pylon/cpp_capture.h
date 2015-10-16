#ifndef CPP_CAPTURE_H
#define CPP_CAPTURE_H

#ifdef __cplusplus
extern "C"
#endif
void startCaptureCPP(int batch, char* outputPath); 

#ifdef __cplusplus
extern "C"
#endif
void stopCaptureCPP();

#ifdef __cplusplus
extern "C"
#endif
void attachDeviceCPP();

#ifdef __cplusplus
extern "C"
#endif
void configureCameraCPP();

#ifdef __cplusplus
extern "C"
#endif
int batchCapturedCPP();

#ifdef __cplusplus
extern "C"
#endif
int totalCapturedCPP();

#endif // CPP_CAPTURE_H
