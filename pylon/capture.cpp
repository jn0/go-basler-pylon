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
        std::string retrieveAndSave(int, int, std::string);
        std::string startCapture();
        void configureCamera();
        void setHardwareTriggerConfiguration();
        void setNodeMapIntParam(const GenICam::gcstring,int64_t);
        void setNodeMapFloatParam(const GenICam::gcstring, double);
        void setNodeMapEnumParam(const GenICam::gcstring, const GenICam::gcstring);

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

void setHardwareTriggerConfiguration() {
    CameraWrapper::getInstance().setHardwareTriggerConfiguration();
}

const char* retrieveAndSave(int batch, int timeout, char* outputPath) {
    std::string op;
    op.assign(outputPath);
    return CameraWrapper::getInstance().retrieveAndSave(batch,timeout,op).c_str();
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

void setNodeMapIntParam(char* name, int value) {
    GenICam::gcstring op;
    op.assign(name);
    CameraWrapper::getInstance().setNodeMapIntParam(op,(int64_t)value);

}
void setNodeMapFloatParam(char* name, double value) {
    GenICam::gcstring op;
    op.assign(name);
    CameraWrapper::getInstance().setNodeMapFloatParam(op,value);

}
void setNodeMapEnumParam(char* name, char* value) {
    GenICam::gcstring op;
    op.assign(name);

    GenICam::gcstring v;
    v.assign(value);
    CameraWrapper::getInstance().setNodeMapEnumParam(op,v);
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

std::string CameraWrapper::retrieveAndSave(int batch, int timeout, std::string outputPath) {
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

void CameraWrapper::setNodeMapIntParam(const GenICam::gcstring name, int64_t value) {
    GenApi::CIntegerPtr(this->camera.GetNodeMap().GetNode(name))->SetValue(value);
}

void CameraWrapper::setNodeMapFloatParam(const GenICam::gcstring name, double value) {
    GenApi::CFloatPtr(this->camera.GetNodeMap().GetNode(name))->SetValue(value);
}

void CameraWrapper::setNodeMapEnumParam(const GenICam::gcstring name, const GenICam::gcstring value) {
    GenApi::CEnumerationPtr(this->camera.GetNodeMap().GetNode(name))->FromString(value);
}

void CameraWrapper::setHardwareTriggerConfiguration() {
    GenApi::INodeMap& nodemap = this->camera.GetNodeMap();

    //Disable all trigger types.
    // Get required enumerations.
    GenApi::CEnumerationPtr triggerSelector( nodemap.GetNode("TriggerSelector"));

    // Check the available camera trigger mode(s) to select the appropriate one: acquisition start trigger mode 
    // (used by previous cameras, i.e. for cameras supporting only the legacy image acquisition control mode;
    // do not confuse with acquisition start command) or frame start trigger mode
    // (used by newer cameras, i.e.for cameras using the standard image acquisition control mode;
    // equivalent to the acquisition start trigger mode in the legacy image acquisition control mode).
    Pylon::String_t triggerName( "FrameStart");
    if ( !IsAvailable( triggerSelector->GetEntryByName(triggerName)))
    {
        triggerName = "AcquisitionStart";
        if ( !IsAvailable( triggerSelector->GetEntryByName(triggerName)))
        {
            throw RUNTIME_EXCEPTION( "Could not select trigger. Neither FrameStart nor AcquisitionStart is available.");
        }
    }

    // Get all enumeration entries of Trigger Selector.
    GenApi::NodeList_t triggerSelectorEntries;
    triggerSelector->GetEntries( triggerSelectorEntries );

    // Turn Trigger Mode off For all Trigger Selector entries.
    for ( GenApi::NodeList_t::iterator it = triggerSelectorEntries.begin(); it != triggerSelectorEntries.end(); ++it) 
    {
        // Set Trigger Mode to off if the trigger is available.
        GenApi::CEnumEntryPtr pEntry(*it);
        if ( IsAvailable( pEntry)) 
        {
            Pylon::String_t triggerNameOfEntry( pEntry->GetSymbolic());
            triggerSelector->FromString( triggerNameOfEntry);
            if ( triggerName == triggerNameOfEntry)
            {
                // Activate trigger.
                this->setNodeMapEnumParam("TriggerMode","On");

                // The trigger source must be set to the trigger input, e.g. 'Line1'.
                // GenApi::CEnumerationPtr(nodemap.GetNode("TriggerSource"))->FromString("Line1");
                this->setNodeMapEnumParam("TriggerSource","Line1");

                //The trigger activation must be set to e.g. 'RisingEdge'.
                //GenApi::CEnumerationPtr(nodemap.GetNode("TriggerActivation"))->FromString("RisingEdge");
                this->setNodeMapEnumParam("TriggerActivation","RisingEdge");
            }
            else
            {
                this->setNodeMapEnumParam("TriggerMode","Off");
            }
        }
    }

    //Set acquisition mode.
    this->setNodeMapEnumParam("AcquisitionMode","Continuous");
}