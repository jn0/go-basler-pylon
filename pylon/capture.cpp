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
        void setHardwareTriggerConfiguration();

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

/*

Configuration classes

*/

class CHardwareTriggerConfiguration : public Pylon::CConfigurationEventHandler {
public:
    static void ApplyConfiguration( GenApi::INodeMap& nodemap)
    {
        //Disable all trigger types.
        // Get required enumerations.
        GenApi::CEnumerationPtr triggerSelector( nodemap.GetNode("TriggerSelector"));
        GenApi::CEnumerationPtr triggerMode( nodemap.GetNode("TriggerMode"));

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
                    triggerMode->FromString( "On");

                    //// The trigger source must be set to the trigger input, e.g. 'Line1'.
                    GenApi::CEnumerationPtr(nodemap.GetNode("TriggerSource"))->FromString("Line1");

                    ////The trigger activation must be set to e.g. 'RisingEdge'.
                    GenApi::CEnumerationPtr(nodemap.GetNode("TriggerActivation"))->FromString("RisingEdge");
                }
                else
                {
                    triggerMode->FromString( "Off");
                }
            }
        }
        

        //Set acquisition mode.
        GenApi::CEnumerationPtr(nodemap.GetNode("AcquisitionMode"))->FromString("Continuous");
    }

    //Set basic camera settings
    virtual void OnOpened( Pylon::CInstantCamera& camera)
    {
        try
        {
            ApplyConfiguration( camera.GetNodeMap());
        }
        catch (GenICam::GenericException& e)
        {
            throw RUNTIME_EXCEPTION( "Could not apply configuration. GenICam::GenericException caught in OnOpened method msg=%hs", e.what());
        }
        catch (std::exception& e)
        {
            throw RUNTIME_EXCEPTION( "Could not apply configuration. std::exception caught in OnOpened method msg=%hs", e.what());
        }
        catch (...)
        {
            throw RUNTIME_EXCEPTION( "Could not apply configuration. Unknown exception caught in OnOpened method.");
        }
    }
};

void CameraWrapper::setHardwareTriggerConfiguration() {
    this->camera.RegisterConfiguration(new CHardwareTriggerConfiguration, Pylon::RegistrationMode_ReplaceAll, Pylon::Cleanup_Delete);
}