// Include files to use the PYLON API.
#include <pylon/PylonIncludes.h>
#include <pylon/gige/BaslerGigEInstantCamera.h>
#include <mutex>

// Namespace for using pylon objects.
// using namespace Pylon;

const int CAPTURE_STATUS_IDLE=0;
const int CAPTURE_STATUS_GRABBING=1;
const int CAPTURE_STATUS_STOP_GRABBING=2;

const int ATTACH_STATUS_NOT_ATTACHED=0;
const int ATTACH_STATUS_ATTACHED=1;

class CameraWrapper {
    public:
        static CameraWrapper& getInstance()
        {
            static CameraWrapper    instance;
            return instance;
        }
        void startCapture(int batch, std::string outputPath);
        void stopCapture();
        void attachDevice();
        void configureCamera();
        int batchCaptured();
        int totalCaptured();

    private:
        CameraWrapper() {}
        CameraWrapper(CameraWrapper const&);              // Don't Implement.
        void operator=(CameraWrapper const&); // Don't implement

        std::mutex captureMutex;
        int captureStatus;
        int batchCapturedQty;
        int totalCapturedQty;

        std::mutex captureStopMutex;

        std::mutex attachMutex;
        int attachStatus;

        Pylon::CBaslerGigEInstantCamera camera;
        Pylon::PylonAutoInitTerm autoInitTerm;

};

void stopCaptureCPP() {
    CameraWrapper::getInstance().stopCapture();
}

void attachDeviceCPP() {
    CameraWrapper::getInstance().attachDevice();
}

void configureCameraCPP() {
    CameraWrapper::getInstance().configureCamera();
}

int batchCapturedCPP() {
    return CameraWrapper::getInstance().batchCaptured();
}

int totalCapturedCPP() {
    return CameraWrapper::getInstance().totalCaptured();
}

void startCaptureCPP(int batch, char* outputPath) {
    std::string op;
    op.assign(outputPath);
    CameraWrapper::getInstance().startCapture(batch,op);
}

void CameraWrapper::stopCapture() {
    if(this->captureStatus==CAPTURE_STATUS_GRABBING) {
        this->captureMutex.lock();
        this->captureStatus=CAPTURE_STATUS_STOP_GRABBING;
        this->captureMutex.unlock();

        // HACK: Two calls to this mutex will block and depend on the start loop to complete and unlock
        this->captureStopMutex.lock();
        this->captureStopMutex.lock();
        this->captureStopMutex.unlock();
    }
}

void CameraWrapper::attachDevice() {
    this->attachMutex.lock();
    if(this->attachStatus==ATTACH_STATUS_NOT_ATTACHED) {
        this->camera.Attach(Pylon::CTlFactory::GetInstance().CreateFirstDevice());
        this->camera.Open();
        this->attachStatus = ATTACH_STATUS_ATTACHED;
    }
    this->attachMutex.unlock();
}

void CameraWrapper::startCapture(int batch, std::string outputPath) {
    this->captureMutex.lock();
    if(this->captureStatus!=CAPTURE_STATUS_IDLE) {
        this->captureMutex.unlock();
        return;
    }
    this->captureStatus=CAPTURE_STATUS_GRABBING;
    this->batchCapturedQty=0;
    this->captureMutex.unlock();
    this->attachDevice();
    
    try {

        // Start the grabbing of c_countOfImagesToGrab images.
        // The camera device is parameterized with a default configuration which
        // sets up free-running continuous acquisition.
        // GrabLoop_ProvidedByUser for testing
        this->camera.StartGrabbing(Pylon::GrabStrategy_OneByOne, Pylon::GrabLoop_ProvidedByInstantCamera);
        // this->camera.StartGrabbing(Pylon::GrabStrategy_OneByOne, Pylon::GrabLoop_ProvidedByUser);

        // This smart pointer will receive the grab result data.
        Pylon::CGrabResultPtr ptrGrabResult;

        while (this->camera.IsGrabbing()) {
            // Wait for an image and then retrieve it. A timeout of 5000 ms is used.
            bool retrieved = this->camera.RetrieveResult(5000, ptrGrabResult, Pylon::TimeoutHandling_ThrowException);
            if(!retrieved) {
                std::cerr << "RetrieveResult false" << std::endl;
                continue;
            }

            // Image grabbed successfully?
            if (retrieved && ptrGrabResult->GrabSucceeded()) {
                Pylon::EPixelType pixelType = ptrGrabResult->GetPixelType();
                uint32_t width = ptrGrabResult->GetWidth();
                uint32_t height = ptrGrabResult->GetHeight();
                size_t paddingX = ptrGrabResult->GetPaddingX();
                size_t bufferSize = ptrGrabResult->GetImageSize();
                void* buffer = ptrGrabResult->GetBuffer();
                
                std::ostringstream nameStream;
                
                nameStream << outputPath << batch << "_" << ptrGrabResult->GetTimeStamp() << ".tiff";
                
                Pylon::CImagePersistence::Save(
                        Pylon::ImageFileFormat_Tiff,
                        nameStream.str().c_str(),
                        buffer,
                        bufferSize,
                        pixelType,
                        width,
                        height,
                        paddingX,
                        Pylon::ImageOrientation_TopDown,NULL);

            this->totalCapturedQty++;
            this->batchCapturedQty++;
            } else {
                std::cout << "Error: " << ptrGrabResult->GetErrorCode() << " " << ptrGrabResult->GetErrorDescription() << std::endl;
            }

            this->captureMutex.lock();
            if(this->captureStatus==CAPTURE_STATUS_STOP_GRABBING) {
                this->captureStatus = CAPTURE_STATUS_IDLE;
                this->captureMutex.unlock();

                // HACK: Unblock stopCapture call
                this->captureStopMutex.unlock();
                camera.StopGrabbing();
                break;
            }
            this->captureMutex.unlock();
        }
    } catch (GenICam::GenericException &e) {
        camera.Close();
        // Error handling.
        std::cerr << "An exception occurred." << std::endl << e.GetDescription() << std::endl;
        this->captureMutex.lock();
        this->captureStatus = CAPTURE_STATUS_IDLE;
        this->captureMutex.unlock();
    }

}

void CameraWrapper::configureCamera() {
    this->camera.GainAuto.SetValue(Basler_GigECamera::GainAuto_Continuous);
    this->camera.GainRaw.SetValue(1);
    this->camera.BlackLevelRaw.SetValue(90);
    this->camera.DigitalShift.SetValue(1);
    this->camera.ExposureAuto.SetValue(Basler_GigECamera::ExposureAuto_Continuous);
}

int CameraWrapper::batchCaptured() {
    return this->batchCapturedQty;
}

int CameraWrapper::totalCaptured() {
    return this->totalCapturedQty;
}