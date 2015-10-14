#ifdef __cplusplus 
extern "C" {
#endif

// prototypes
void startCapture(int batch, char* outputPath);
void stopCapture();
void attachDevice();
void configureCamera();

#ifdef __cplusplus
}
#endif
