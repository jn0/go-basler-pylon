// Include files to use the PYLON API.
#include <pylon/PylonIncludes.h>
#include <pylon/gige/BaslerGigEInstantCamera.h>

class CameraWrapper {
    public:
        static CameraWrapper& getInstance()
        {
            static CameraWrapper    instance;
            return instance;
        }
        void stopCapture();
        void attachDevice();
        void openCamera();
        void closeCamera();
        bool isGrabbing();
        bool isAttached();
        bool isOpen();
        std::string grab(int batch, int timeout, std::string outputPath);
        std::string startCapture();
        void configureCamera();

    private:
        CameraWrapper() {}
        CameraWrapper(CameraWrapper const&);              // Don't Implement.
        void operator=(CameraWrapper const&); // Don't implement

        Pylon::CBaslerGigEInstantCamera camera;
        Pylon::PylonAutoInitTerm autoInitTerm;
};

void stopCapture() {
    CameraWrapper::getInstance().stopCapture();
}

void attachDevice() {
    CameraWrapper::getInstance().attachDevice();
}

void configureCamera() {
    CameraWrapper::getInstance().configureCamera();
}

const char* grab(int batch, int timeout, char* outputPath) {
    std::string op;
    op.assign(outputPath);
    return CameraWrapper::getInstance().grab(batch,timeout,op).c_str();
}
const char* startCapture() {
    return CameraWrapper::getInstance().startCapture().c_str();
}

void openCamera() {
    CameraWrapper::getInstance().openCamera();
}

void closeCamera() {
    CameraWrapper::getInstance().closeCamera();
}

bool isCameraGrabbing() {
    return CameraWrapper::getInstance().isGrabbing();
}

bool isAttached() {
    return CameraWrapper::getInstance().isAttached();
}

bool isOpen() {
    return CameraWrapper::getInstance().isOpen();
}

void CameraWrapper::stopCapture() {
    camera.Close();
}

void CameraWrapper::attachDevice() {
    this->camera.Attach(Pylon::CTlFactory::GetInstance().CreateFirstDevice());
}

void CameraWrapper::openCamera() {
    this->camera.Open();
}

void CameraWrapper::closeCamera() {
    this->camera.Close();
}

bool CameraWrapper::isGrabbing() {
    return this->camera.IsGrabbing();
}

bool CameraWrapper::isAttached() {
    return this->camera.IsPylonDeviceAttached();
}

bool CameraWrapper::isOpen() {
    return this->camera.IsOpen();
}

std::string CameraWrapper::grab(int batch, int timeout, std::string outputPath) {
    Pylon::CGrabResultPtr ptrGrabResult;

    try {
        if(this->camera.RetrieveResult(timeout, ptrGrabResult, Pylon::TimeoutHandling_ThrowException) && ptrGrabResult->GrabSucceeded()) {
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
        } else {
            std::ostringstream result;
            result << ptrGrabResult->GetErrorCode() << " " << ptrGrabResult->GetErrorDescription();
            return result.str();
            
        }
    } catch (GenICam::GenericException &e) {
        return e.GetDescription();
    }
    return "";
}

std::string CameraWrapper::startCapture() {
    try {
        // Pylon::GrabLoop_ProvidedByInstantCamera
        this->camera.StartGrabbing(Pylon::GrabStrategy_OneByOne, Pylon::GrabLoop_ProvidedByUser);
    } catch (GenICam::GenericException &e) {
        camera.Close();
        // Error handling.
        return e.GetDescription();
    }
    return "";
}

void CameraWrapper::configureCamera() {
    this->camera.GainAuto.SetValue(Basler_GigECamera::GainAuto_Continuous);
    this->camera.GainRaw.SetValue(1);
    this->camera.BlackLevelRaw.SetValue(90);
    this->camera.DigitalShift.SetValue(1);
    this->camera.ExposureAuto.SetValue(Basler_GigECamera::ExposureAuto_Continuous);
}