// Include files to use the PYLON API.
#include <pylon/PylonIncludes.h>

#ifdef GIGE
#include <pylon/gige/BaslerGigEInstantCamera.h>
#define Basler_XCamera Basler_GigECamera
#define CBaslerXCamera Pylon::CBaslerGigEInstantCamera
#else
#include <pylon/usb/PylonUsbIncludes.h>
#define Basler_XCamera Basler_UsbCameraParams
#define CBaslerXCamera Pylon::CBaslerUsbInstantCamera
#endif
#include <capture.h>
#include <pylon/InstantCamera.h>

class CameraWrapper {
    public:
        static CameraWrapper& getInstance() {
            static CameraWrapper instance;
            return instance;
        }
        std::string stopCapture();
        std::string attachDevice();
        std::string openCamera();
        std::string closeCamera();
        bool isGrabbing();
        bool isAttached();
        bool isOpen();
        std::string retrieveAndSave(int, int, std::string);
        std::string startCapture(int max);
        std::string configureCamera();
        std::string setHardwareTriggerConfiguration();
        std::string setNodeMapIntParam(const GenICam::gcstring,int64_t);
        std::string setNodeMapFloatParam(const GenICam::gcstring, double);
        std::string setNodeMapEnumParam(const GenICam::gcstring,
					const GenICam::gcstring);

    private:
        CameraWrapper() {}
        CameraWrapper(CameraWrapper const&);  // Don't Implement.
        void operator=(CameraWrapper const&); // Don't implement

        CBaslerXCamera camera;
        Pylon::PylonAutoInitTerm autoInitTerm;
    	int width, height;
	Pylon::CInstantCamera::DeviceInfo_t info() { return this->camera.GetDeviceInfo(); }
	std::string fullName, vendorName, modelName;
    public:
	int Width() { return width; }
	int Height() { return height; }
	const char* FullName() { return fullName.c_str(); }
	const char* VendorName() { return vendorName.c_str(); }
	const char* ModelName() { return modelName.c_str(); }
};

/******************************************************************************
 * EXTERNALLY VISIBLE BEGIN
 */
const char* stopCapture() {
    std::string msg = CameraWrapper::getInstance().stopCapture();
    std::cerr << __func__ << "():" << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();
}

const char* attachDevice() {
    std::string msg = CameraWrapper::getInstance().attachDevice();
    std::cerr << __func__ << "():" << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();
}

const char* configureCamera() {
    assert(isOpen());
    std::string msg = CameraWrapper::getInstance().configureCamera();
    std::cerr << __func__ << "():" << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();
}

const char* setHardwareTriggerConfiguration() {
    std::string msg = CameraWrapper::getInstance().setHardwareTriggerConfiguration();
    std::cerr << __func__ << "():" << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();
}

const char* retrieveAndSave(int batch, int timeout, char* outputPath) {
    std::string op;
    op.assign(outputPath);
    std::string msg = CameraWrapper::getInstance().retrieveAndSave(batch,
								   timeout,
								   op);
    std::cerr << __func__ << "(" << batch << ", " << timeout << ", " << outputPath << "):"
	      << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();
}
const char* startCapture(int max) {
    std::string msg = CameraWrapper::getInstance().startCapture(max);
    std::cerr << __func__ << "():" << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();
}

const char* openCamera() {
    assert(isAttached());
    std::string msg = CameraWrapper::getInstance().openCamera();
    std::cerr << __func__ << "():" << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();
}

const char* closeCamera() {
    std::string msg = CameraWrapper::getInstance().closeCamera();
    std::cerr << __func__ << "():" << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();
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

const char* setNodeMapIntParam(char* name, int value) {
    GenICam::gcstring op;
    op.assign(name);
    std::string msg = CameraWrapper::getInstance().setNodeMapIntParam(op,
								      (int64_t)value);
    std::cerr << __func__ << "(" << name << ", " << value << "):"
	      << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();

}
const char* setNodeMapFloatParam(char* name, double value) {
    GenICam::gcstring op;
    op.assign(name);
    std::string msg = CameraWrapper::getInstance().setNodeMapFloatParam(op, value);
    std::cerr << __func__ << "(" << name << ", " << value << "):"
	      << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();

}
const char* setNodeMapEnumParam(char* name, char* value) {
    GenICam::gcstring op, v;
    op.assign(name);
    v.assign(value);
    std::string msg = CameraWrapper::getInstance().setNodeMapEnumParam(op, v);
    std::cerr << __func__ << "(" << name << ", " << value << "):"
	      << (msg == "" ? "ok" : msg) << std::endl;
    return msg.c_str();
}

const char* fullName() { return CameraWrapper::getInstance().FullName(); }
const char* vendorName() { return CameraWrapper::getInstance().VendorName(); }
const char* modelName() { return CameraWrapper::getInstance().ModelName(); }
int width() { return CameraWrapper::getInstance().Width(); }
int height() { return CameraWrapper::getInstance().Height(); }
/*
 * EXTERNALLY VISIBLE END
 ******************************************************************************/

std::string CameraWrapper::stopCapture() {
    try {
	camera.Close();
    } catch (GenICam::GenericException &e) {
        std::string msg = "stopCapture: "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }
    return "";
}

std::string CameraWrapper::attachDevice() {
    try {
	this->camera.Attach(Pylon::CTlFactory::GetInstance().CreateFirstDevice());
    } catch (GenICam::GenericException &e) {
        std::string msg = __func__; msg += ": "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }
    return "";
}

std::string CameraWrapper::openCamera() {
    try {
	this->camera.Open();
    } catch (GenICam::GenericException &e) {
        std::string msg = __func__; msg += ": "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }

    GenApi::CIntegerPtr width(camera.GetNodeMap().GetNode("Width"));
    GenApi::CIntegerPtr height(camera.GetNodeMap().GetNode("Height"));
    this->width = width->GetValue();
    this->height = height->GetValue();

    this->fullName.assign(this->info().GetFullName().c_str());
    this->vendorName.assign(this->info().GetVendorName().c_str());
    this->modelName.assign(this->info().GetModelName().c_str());

    std::cerr << "Camera [" << this->Width() << "×" << this->Height() << "]"
	      << " [" << this->fullName.c_str() << "]"
	      << " [" << this->vendorName.c_str() << "]"
	      << " [" << this->modelName.c_str() << "]"
	      << std::endl;
/*
    Pylon::CPylonUsbCameraT::DeviceInfo_t di = this->camera.GetDeviceInfo();
    std::cerr << "Camera ["
	      << di->GetFullName().c_str()
	      << "]" << std::endl;
    StringList_t props;
    di->GetPropertyNames(&props);
    String_t name = props.first();
    while (name) {
    	String_t value;
    	di->GetPropertyValue(name, &value);
	name = props.next();
	std::cerr << "[" << name << "]=[" << value << "]" << std::endl;
    }
*/
    return "";
}

std::string CameraWrapper::closeCamera() {
    try {
	this->camera.Close();
    } catch (GenICam::GenericException &e) {
        std::string msg = __func__; msg += ": "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }
    this->width = this->height = -1;
    return "";
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

std::string CameraWrapper::retrieveAndSave(int batch,
                                           int timeout,
					   std::string outputPath) {
    Pylon::CGrabResultPtr ptrGrabResult;

    try {
        if(this->camera.RetrieveResult(timeout,
	                               ptrGrabResult,
				       Pylon::TimeoutHandling_ThrowException) &&
           ptrGrabResult->GrabSucceeded()) {
            Pylon::EPixelType pixelType = ptrGrabResult->GetPixelType();
            uint32_t width = ptrGrabResult->GetWidth();
            uint32_t height = ptrGrabResult->GetHeight();
            size_t paddingX = ptrGrabResult->GetPaddingX();
            size_t bufferSize = ptrGrabResult->GetImageSize();
            void* buffer = ptrGrabResult->GetBuffer();
            
            std::ostringstream nameStream;
            nameStream << outputPath
	               << batch
		       << "_"
		       << ptrGrabResult->GetTimeStamp()
		       << ".tiff";
            
            Pylon::CImagePersistence::Save(
                    Pylon::ImageFileFormat_Tiff,
                    nameStream.str().c_str(),
                    buffer,
                    bufferSize,
                    pixelType,
                    width,
                    height,
                    paddingX,
                    Pylon::ImageOrientation_TopDown,
		    NULL);
        } else {
            std::ostringstream result;
            result << ptrGrabResult->GetErrorCode()
	           << " "
		   << ptrGrabResult->GetErrorDescription();
            return result.str();
        }
    } catch (GenICam::GenericException &e) {
        std::string msg = __func__; msg += ": "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
        return msg;
    }
    return "";
}

std::string CameraWrapper::startCapture(int max) {
    try {
	if (max <= 0) {
            // Pylon::GrabLoop_ProvidedByInstantCamera
	    this->camera.StartGrabbing(Pylon::GrabStrategy_OneByOne,
	                               Pylon::GrabLoop_ProvidedByUser);
	} else {
	    this->camera.StartGrabbing(max,
	    				Pylon::GrabStrategy_OneByOne,
					Pylon::GrabLoop_ProvidedByInstantCamera);
	}
    } catch (GenICam::GenericException &e) {
	std::string msg = __func__; msg += ":StartGrabbing: "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
        camera.Close();
        return msg;
    }
    return "";
}

std::string CameraWrapper::configureCamera() {
    std::string msg = __func__; msg += ": ";
    try {
        this->camera.GainAuto.SetValue(Basler_XCamera::GainAuto_Continuous);
    } catch (GenICam::GenericException &e) {
	msg += "GainAuto.Set: "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }
#ifdef GIGE
    this->camera.GainRaw.SetValue(1);
    this->camera.BlackLevelRaw.SetValue(90);
#else
    try {
        this->camera.Gain.SetValue(this->camera.Gain.GetMin());
    } catch (GenICam::GenericException &e) {
        msg += "Gain.Set: "; msg += e.GetDescription();
	msg += this->isOpen() ? "; cam is open" : "; cam is closed";
	std::cerr << msg << std::endl;
	return msg;
    }
    try {
        this->camera.BlackLevel.SetValue(90.0);
    } catch (GenICam::GenericException &e) {
        msg += "BlackLevel.Set: "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }
#endif
    try {
        this->camera.DigitalShift.SetValue(1);
    } catch (GenICam::GenericException &e) {
        msg += "DigitalShift.Set: "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }
    try {
        this->camera.ExposureAuto.SetValue(Basler_XCamera::ExposureAuto_Continuous);
    } catch (GenICam::GenericException &e) {
        msg += "ExposureAuto.Set: "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }
    return "";
}

std::string CameraWrapper::setNodeMapIntParam(const GenICam::gcstring name,
					      int64_t value) {
    try {
	GenApi::CIntegerPtr(this->camera.GetNodeMap().GetNode(name)
			   )->SetValue(value);
    } catch (GenICam::GenericException &e) {
        std::string msg = "setNodeMapIntParam: "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }
    return "";
}

std::string CameraWrapper::setNodeMapFloatParam(const GenICam::gcstring name,
						double value) {
    try {
	GenApi::CFloatPtr(this->camera.GetNodeMap().GetNode(name))->SetValue(value);
    } catch (GenICam::GenericException &e) {
        std::string msg = "setNodeMapFloatParam: "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg.c_str();
    }
    return "";
}

std::string CameraWrapper::setNodeMapEnumParam(const GenICam::gcstring name,
					       const GenICam::gcstring value) {
    try {
	GenApi::CEnumerationPtr(this->camera.GetNodeMap().GetNode(name)
				)->FromString(value);
    } catch (GenICam::GenericException &e) {
        std::string msg = "setNodeMapEnumParam: "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg.c_str();
    }
    return "";
}

std::string CameraWrapper::setHardwareTriggerConfiguration() {
    GenApi::INodeMap& nodemap = this->camera.GetNodeMap();

    //Disable all trigger types.
    // Get required enumerations.
    GenApi::CEnumerationPtr triggerSelector( nodemap.GetNode("TriggerSelector"));

    // Check the available camera trigger mode(s) to select the appropriate one:
    // acquisition start trigger mode (used by previous cameras, i.e. for
    // cameras supporting only the legacy image acquisition control mode;
    // do not confuse with acquisition start command) or frame start trigger mode
    // (used by newer cameras, i.e. for cameras using the standard image
    // acquisition control mode, equivalent to the acquisition start trigger
    // mode in the legacy image acquisition control mode).
    Pylon::String_t triggerName( "FrameStart");
    if (!IsAvailable( triggerSelector->GetEntryByName(triggerName))) {
        triggerName = "AcquisitionStart";
        if (!IsAvailable( triggerSelector->GetEntryByName(triggerName))) {
            throw RUNTIME_EXCEPTION("Could not select trigger."
	    	" Neither FrameStart nor AcquisitionStart is available.");
        }
    }

    // Get all enumeration entries of Trigger Selector.
    GenApi::NodeList_t triggerSelectorEntries;
    triggerSelector->GetEntries( triggerSelectorEntries );

    // Turn Trigger Mode off For all Trigger Selector entries.
    for (GenApi::NodeList_t::iterator it = triggerSelectorEntries.begin();
    	 it != triggerSelectorEntries.end();
	 ++it) {
        // Set Trigger Mode to off if the trigger is available.
        GenApi::CEnumEntryPtr pEntry(*it);
        if (IsAvailable(pEntry)) {
            Pylon::String_t triggerNameOfEntry(pEntry->GetSymbolic());
            triggerSelector->FromString( triggerNameOfEntry);
            if (triggerName == triggerNameOfEntry) {
                // Activate trigger.
                this->setNodeMapEnumParam("TriggerMode", "On");

                // The trigger source must be set to the trigger input, e.g.
		// 'Line1'.
                // GenApi::CEnumerationPtr(nodemap.GetNode("TriggerSource")
		//			  )->FromString("Line1");
                this->setNodeMapEnumParam("TriggerSource", "Line1");

                // The trigger activation must be set to e.g. 'RisingEdge'.
                // GenApi::CEnumerationPtr(nodemap.GetNode("TriggerActivation")
		//			  )->FromString("RisingEdge");
                this->setNodeMapEnumParam("TriggerActivation", "RisingEdge");
            } else {
                this->setNodeMapEnumParam("TriggerMode", "Off");
            }
        }
    }

    //Set acquisition mode.
    this->setNodeMapEnumParam("AcquisitionMode", "Continuous");

    return "";
}
