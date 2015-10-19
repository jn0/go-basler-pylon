package gige

import ".."

const (

	// SequenceControl - This category includes items that control the sequencer feature
	SequenceEnable         IBoolean = "SequenceEnable"         // GenApi::IBoolean || Enables the sequencer.
	SequenceAsyncRestart   ICommand = "SequenceAsyncRestart"   // GenApi::ICommand || Allows asynchronous restart of the sequence of sequence sets.
	SequenceAsyncAdvance   ICommand = "SequenceAsyncAdvance"   // GenApi::ICommand || Allows asynchronous advance from one sequence set to the next.
	SequenceCurrentSet     IInteger = "SequenceCurrentSet"     // GenApi::IInteger || Current sequence set.
	SequenceSetTotalNumber IInteger = "SequenceSetTotalNumber" // GenApi::IInteger || Total number of sequence sets.
	SequenceSetIndex       IInteger = "SequenceSetIndex"       // GenApi::IInteger || Selects the index number of a sequence set.
	SequenceSetLoad        ICommand = "SequenceSetLoad"        // GenApi::ICommand || Loads a sequence set.
	SequenceSetStore       ICommand = "SequenceSetStore"       // GenApi::ICommand || Stores the current sequence set.
	SequenceSetExecutions  IInteger = "SequenceSetExecutions"  // GenApi::IInteger || Sets the number of sequence set executions.

	// SequenceControlConfiguration - This category includes items that control the sequence set advance
	SequenceAdvanceMode        IEnumerationT = "SequenceAdvanceMode"        // GenApi::IEnumerationT<SequenceAdvanceModeEnums> || Selects the sequence set advance mode.
	SequenceControlSelector    IEnumerationT = "SequenceControlSelector"    // GenApi::IEnumerationT<SequenceControlSelectorEnums> || Selects between sequence restart or sequence set advance.
	SequenceControlSource      IEnumerationT = "SequenceControlSource"      // GenApi::IEnumerationT<SequenceControlSourceEnums> || Selects the source for sequence control.
	SequenceAddressBitSelector IEnumerationT = "SequenceAddressBitSelector" // GenApi::IEnumerationT<SequenceAddressBitSelectorEnums> || Selects a bit of the sequence set address.
	SequenceAddressBitSource   IEnumerationT = "SequenceAddressBitSource"   // GenApi::IEnumerationT<SequenceAddressBitSourceEnums> || Selects the source for the selected bit of the sequence set address.

	// AnalogControls - This category includes items that control the analog characteristics of the video signal
	GainAuto           IEnumerationT = "GainAuto"           // GenApi::IEnumerationT<GainAutoEnums> || Gain Auto is the 'automatic' counterpart of the manual gain feature.
	GainRaw            IInteger      = "GainRaw"            // GenApi::IInteger || This is an integer value that sets the selected gain control in device specific units.
	GainAbs            IFloat        = "GainAbs"            // GenApi::IFloat || This is a float value that sets the selected gain control in dB.
	GainSelector       IEnumerationT = "GainSelector"       // GenApi::IEnumerationT<GainSelectorEnums> || Selects the gain control to configure.
	BlackLevelRaw      IInteger      = "BlackLevelRaw"      // GenApi::IInteger || Sets the value of the selected black level control as an integer.
	BlackLevelAbs      IFloat        = "BlackLevelAbs"      // GenApi::IFloat || Sets the value of the selected black level control as a float.
	BlackLevelSelector IEnumerationT = "BlackLevelSelector" // GenApi::IEnumerationT<BlackLevelSelectorEnums> || Selcts a black level control to configure.
	GammaEnable        IBoolean      = "GammaEnable"        // GenApi::IBoolean || Enables the gamma correction.
	GammaSelector      IEnumerationT = "GammaSelector"      // GenApi::IEnumerationT<GammaSelectorEnums> || This enumeration selects the type of gamma to apply.
	Gamma              IFloat        = "Gamma"              // GenApi::IFloat || This feature is used to perform gamma correction of pixel intensity.
	DigitalShift       IInteger      = "DigitalShift"       // GenApi::IInteger || Sets the value of the selected digital shift control.
	SubstrateVoltage   IInteger      = "SubstrateVoltage"   // GenApi::IInteger || Sets the substrate voltage.

	// ImageFormat - This category includes items that control the size of the acquired image and the format of the transferred pixel data
	PixelFormat                   IEnumerationT = "PixelFormat"                   // GenApi::IEnumerationT<PixelFormatEnums> || Sets the format of the pixel data transmitted for acquired images.
	PixelCoding                   IEnumerationT = "PixelCoding"                   // GenApi::IEnumerationT<PixelCodingEnums> || Sets the color coding of the pixels in the acquired images.
	PixelSize                     IEnumerationT = "PixelSize"                     // GenApi::IEnumerationT<PixelSizeEnums> || Indicates the depth of the pixel values in the image in bits per pixel.
	PixelColorFilter              IEnumerationT = "PixelColorFilter"              // GenApi::IEnumerationT<PixelColorFilterEnums> || Indicates the alignment of the camera's Bayer filter to the pixels in the acquired images.
	ProcessedRawEnable            IBoolean      = "ProcessedRawEnable"            // GenApi::IBoolean || Enables color improved RGB raw output.
	SpatialCorrection             IInteger      = "SpatialCorrection"             // GenApi::IInteger || Specifies the direction of imaging and the separation (consecutive numbers) of related line captures.
	SpatialCorrectionAmount       IInteger      = "SpatialCorrectionAmount"       // GenApi::IInteger || <b>Visibility</b> = Invisible
	SpatialCorrectionStartingLine IEnumerationT = "SpatialCorrectionStartingLine" // GenApi::IEnumerationT<SpatialCorrectionStartingLineEnums> || <b>Visibility</b> = Invisible
	SensorBitDepth                IEnumerationT = "SensorBitDepth"                // GenApi::IEnumerationT<SensorBitDepthEnums> || This feature selects the amount of data bits the sensor produces for one sample.
	ReverseX                      IBoolean      = "ReverseX"                      // GenApi::IBoolean || Enables the horizontal flipping of the image.
	ReverseY                      IBoolean      = "ReverseY"                      // GenApi::IBoolean || Enables the vertical flipping of the image.
	PixelDynamicRangeMin          IInteger      = "PixelDynamicRangeMin"          // GenApi::IInteger || Indicates the minimum possible pixel value that could be transferred from the camera.
	PixelDynamicRangeMax          IInteger      = "PixelDynamicRangeMax"          // GenApi::IInteger || Indicates the maximum possible pixel value that could be transferred from the camera.
	FieldOutputMode               IEnumerationT = "FieldOutputMode"               // GenApi::IEnumerationT<FieldOutputModeEnums> || Selects the mode to output the fields.
	TestImageSelector             IEnumerationT = "TestImageSelector"             // GenApi::IEnumerationT<TestImageSelectorEnums> || Selecting a test image from the list will enable the test image.
	SensorDigitizationTaps        IEnumerationT = "SensorDigitizationTaps"        // GenApi::IEnumerationT<SensorDigitizationTapsEnums> || This feature represents the number of digitized samples outputted simultaneously by the camera A/D conversion stage.

	// DeviceInformation - This category includes items that describe the device and its sensor
	SensorWidth            IInteger      = "SensorWidth"            // GenApi::IInteger || Indicates the width of the camera's sensor in pixels.
	SensorHeight           IInteger      = "SensorHeight"           // GenApi::IInteger || Indicates the height of the camera's sensor in pixels.
	WidthMax               IInteger      = "WidthMax"               // GenApi::IInteger || Indicates the maximum allowed width of the image in pixels.
	HeightMax              IInteger      = "HeightMax"              // GenApi::IInteger || Indicates the maximum allowed height of the image in pixels.
	DeviceVendorName       IString       = "DeviceVendorName"       // GenApi::IString || Indicates the name of the device's vendor.
	DeviceModelName        IString       = "DeviceModelName"        // GenApi::IString || Indicates the model name of the device.
	DeviceManufacturerInfo IString       = "DeviceManufacturerInfo" // GenApi::IString || Provides additional information from the vendor about the device.
	DeviceVersion          IString       = "DeviceVersion"          // GenApi::IString || Indicates the version of the device.
	DeviceFirmwareVersion  IString       = "DeviceFirmwareVersion"  // GenApi::IString || Indicates the version of the device's firmware and software.
	DeviceID               IString       = "DeviceID"               // GenApi::IString || A unique identifier for the device such as a serial number or a GUID.
	DeviceUserID           IString       = "DeviceUserID"           // GenApi::IString || A device ID that is user programmable.
	DeviceReset            ICommand      = "DeviceReset"            // GenApi::ICommand || Immediately resets and reboots the device.
	DeviceScanType         IEnumerationT = "DeviceScanType"         // GenApi::IEnumerationT<DeviceScanTypeEnums> || Indicates the scan type of the device's sensor.
	LastError              IEnumerationT = "LastError"              // GenApi::IEnumerationT<LastErrorEnums> || Indicates the error that was detected last.
	ClearLastError         ICommand      = "ClearLastError"         // GenApi::ICommand || Erases the last error and possibly reveals a previous error.
	TemperatureSelector    IEnumerationT = "TemperatureSelector"    // GenApi::IEnumerationT<TemperatureSelectorEnums> || Lists the temperature sources available for readout.
	TemperatureAbs         IFloat        = "TemperatureAbs"         // GenApi::IFloat || Shows the current temperature of the selected target in degrees centigrade.
	TemperatureState       IEnumerationT = "TemperatureState"       // GenApi::IEnumerationT<TemperatureStateEnums> || Temperature State.
	CriticalTemperature    IBoolean      = "CriticalTemperature"    // GenApi::IBoolean || Shows the over temperature state of the selected target.
	OverTemperature        IBoolean      = "OverTemperature"        // GenApi::IBoolean || Shows the over temperature state of the selected target.

	// ColorImprovementsControl - This category includes items that control color improvement
	ColorTransformationSelector        IEnumerationT = "ColorTransformationSelector"        // GenApi::IEnumerationT<ColorTransformationSelectorEnums> || Selects the matrix color transformation between color spaces.
	LightSourceSelector                IEnumerationT = "LightSourceSelector"                // GenApi::IEnumerationT<LightSourceSelectorEnums> || Selects the type of light source to be considered for matrix color transformation.
	ColorTransformationMatrixFactor    IFloat        = "ColorTransformationMatrixFactor"    // GenApi::IFloat || Defines the extent to which the selected light source will be considered (float)
	ColorTransformationMatrixFactorRaw IInteger      = "ColorTransformationMatrixFactorRaw" // GenApi::IInteger || Defines the extent to which the selected light source will be considered (integer)
	ColorTransformationValueSelector   IEnumerationT = "ColorTransformationValueSelector"   // GenApi::IEnumerationT<ColorTransformationValueSelectorEnums> || Selects the element to be entered in the color transformation matrix.
	ColorTransformationValue           IFloat        = "ColorTransformationValue"           // GenApi::IFloat || Sets a floating point value for the selected element in the color transformation matrix.
	ColorTransformationValueRaw        IInteger      = "ColorTransformationValueRaw"        // GenApi::IInteger || Sets an integer value for the selected element in the color transformation matrix.
	ColorAdjustmentEnable              IBoolean      = "ColorAdjustmentEnable"              // GenApi::IBoolean || Enables color adjustment.
	ColorAdjustmentReset               ICommand      = "ColorAdjustmentReset"               // GenApi::ICommand || Allows returning to previous settings.
	ColorAdjustmentSelector            IEnumerationT = "ColorAdjustmentSelector"            // GenApi::IEnumerationT<ColorAdjustmentSelectorEnums> || Selects the color for color adjustment.
	ColorAdjustmentHue                 IFloat        = "ColorAdjustmentHue"                 // GenApi::IFloat || Adjustment of hue of the selected color (float)
	ColorAdjustmentHueRaw              IInteger      = "ColorAdjustmentHueRaw"              // GenApi::IInteger || Adjustment of hue of the selected color (integer)
	ColorAdjustmentSaturation          IFloat        = "ColorAdjustmentSaturation"          // GenApi::IFloat || Adjustment of saturation of the selected color (float)
	ColorAdjustmentSaturationRaw       IInteger      = "ColorAdjustmentSaturationRaw"       // GenApi::IInteger || Adjustment of saturation of the selected color (integer)
	BalanceWhiteReset                  ICommand      = "BalanceWhiteReset"                  // GenApi::ICommand || Allows returning to previous settings.
	BalanceWhiteAuto                   IEnumerationT = "BalanceWhiteAuto"                   // GenApi::IEnumerationT<BalanceWhiteAutoEnums> || Balance White Auto is the 'automatic' counterpart of the manual white balance feature.
	BalanceRatioSelector               IEnumerationT = "BalanceRatioSelector"               // GenApi::IEnumerationT<BalanceRatioSelectorEnums> || Selects a balance ratio to configure.
	BalanceRatioAbs                    IFloat        = "BalanceRatioAbs"                    // GenApi::IFloat || Sets the value of the selected balance ratio control as a float.
	BalanceRatioRaw                    IInteger      = "BalanceRatioRaw"                    // GenApi::IInteger || Sets the value of the selected balance ratio control as an integer.

	// AOI - This category includes items used to set the size and position of the area of interest
	Width                 IInteger      = "Width"                 // GenApi::IInteger || Sets the width of the area of interest in pixels.
	Height                IInteger      = "Height"                // GenApi::IInteger || Sets the height of the area of interest in pixels.
	OffsetX               IInteger      = "OffsetX"               // GenApi::IInteger || Sets the X offset (left offset) of the area of interest in pixels.
	OffsetY               IInteger      = "OffsetY"               // GenApi::IInteger || Sets the Y offset (top offset) for the area of interest in pixels.
	CenterX               IBoolean      = "CenterX"               // GenApi::IBoolean || Enables the horizontal centering of the image.
	CenterY               IBoolean      = "CenterY"               // GenApi::IBoolean || Enables the vertical centering of the image.
	BinningModeVertical   IEnumerationT = "BinningModeVertical"   // GenApi::IEnumerationT<BinningModeVerticalEnums> || Sets the vertical binning mode.
	BinningModeHorizontal IEnumerationT = "BinningModeHorizontal" // GenApi::IEnumerationT<BinningModeHorizontalEnums> || Sets the horizontal binning mode.
	BinningVertical       IInteger      = "BinningVertical"       // GenApi::IInteger || Sets the number of adjacent vertical pixes to be summed.
	BinningHorizontal     IInteger      = "BinningHorizontal"     // GenApi::IInteger || Sets the number of adjacent horizontal pixes to be summed.
	LegacyBinningVertical IEnumerationT = "LegacyBinningVertical" // GenApi::IEnumerationT<LegacyBinningVerticalEnums> || Sets the vertical binning feature.
	DecimationVertical    IInteger      = "DecimationVertical"    // GenApi::IInteger || Sets vertical sub-sampling.
	DecimationHorizontal  IInteger      = "DecimationHorizontal"  // GenApi::IInteger || Sets horizontal sub-sampling.

	// StackedZoneImaging -
	StackedZoneImagingEnable      IBoolean = "StackedZoneImagingEnable"      // GenApi::IBoolean || Enables the stacked zone imaging feature.
	StackedZoneImagingIndex       IInteger = "StackedZoneImagingIndex"       // GenApi::IInteger || This value sets the zone to access.
	StackedZoneImagingZoneEnable  IBoolean = "StackedZoneImagingZoneEnable"  // GenApi::IBoolean || Enables the selected zone.
	StackedZoneImagingZoneOffsetY IInteger = "StackedZoneImagingZoneOffsetY" // GenApi::IInteger || Sets the Y offset (top offset) for the selected zone.
	StackedZoneImagingZoneHeight  IInteger = "StackedZoneImagingZoneHeight"  // GenApi::IInteger || Sets the height for the selected zone.

	// AcquisitionTrigger - This category includes items used to set the image acquisition parameters and to start and stop acquisition
	EnableBurstAcquisition       IBoolean      = "EnableBurstAcquisition"       // GenApi::IBoolean || When enabled, the maximum frame rate does not depend on the image transfer rate out of the camera.
	AcquisitionMode              IEnumerationT = "AcquisitionMode"              // GenApi::IEnumerationT<AcquisitionModeEnums> || Sets the image acquisition mode.
	AcquisitionStart             ICommand      = "AcquisitionStart"             // GenApi::ICommand || Starts the acquisition of images.
	AcquisitionStop              ICommand      = "AcquisitionStop"              // GenApi::ICommand || Stops the acquisition of images.
	AcquisitionAbort             ICommand      = "AcquisitionAbort"             // GenApi::ICommand || Immediately aborts the acquisition of images.
	AcquisitionFrameCount        IInteger      = "AcquisitionFrameCount"        // GenApi::IInteger || Sets the number of frames acquired in the multiframe acquisition mode.
	TriggerControlImplementation IEnumerationT = "TriggerControlImplementation" // GenApi::IEnumerationT<TriggerControlImplementationEnums> || <b>Visibility</b> = Expert
	TriggerSelector              IEnumerationT = "TriggerSelector"              // GenApi::IEnumerationT<TriggerSelectorEnums> || Selects the trigger type to configure.
	TriggerMode                  IEnumerationT = "TriggerMode"                  // GenApi::IEnumerationT<TriggerModeEnums> || Sets the mode for the selected trigger.
	TriggerSoftware              ICommand      = "TriggerSoftware"              // GenApi::ICommand || Generates a software trigger signal that is used when the trigger source is set to 'software'.
	TriggerSource                IEnumerationT = "TriggerSource"                // GenApi::IEnumerationT<TriggerSourceEnums> || Sets the signal source for the selected trigger.
	TriggerActivation            IEnumerationT = "TriggerActivation"            // GenApi::IEnumerationT<TriggerActivationEnums> || Sets the signal transition needed to activate the selected trigger.
	TriggerPartialClosingFrame   IBoolean      = "TriggerPartialClosingFrame"   // GenApi::IBoolean || Determines whether a partial or complete frame is transmitted when the frame start trigger prematurely transitions.
	TriggerDelayAbs              IFloat        = "TriggerDelayAbs"              // GenApi::IFloat || Sets the trigger delay time in microseconds.
	TriggerDelaySource           IEnumerationT = "TriggerDelaySource"           // GenApi::IEnumerationT<TriggerDelaySourceEnums> || Selects the kind of trigger delay.
	TriggerDelayLineTriggerCount IInteger      = "TriggerDelayLineTriggerCount" // GenApi::IInteger || Sets the trigger delay expressed as number of line triggers.
	ExposureStartDelayRaw        IInteger      = "ExposureStartDelayRaw"        // GenApi::IInteger || <b>Visibility</b> = Beginner
	ExposureStartDelayAbs        IFloat        = "ExposureStartDelayAbs"        // GenApi::IFloat || <b>Visibility</b> = Beginner
	ExposureAuto                 IEnumerationT = "ExposureAuto"                 // GenApi::IEnumerationT<ExposureAutoEnums> || Exposure Auto is the 'automatic' counterpart to manually setting an 'absolute' exposure time.
	ExposureTimeRaw              IInteger      = "ExposureTimeRaw"              // GenApi::IInteger || Sets the 'raw' exposure time.
	ExposureTimeAbs              IFloat        = "ExposureTimeAbs"              // GenApi::IFloat || Directly sets the camera's exposure time in microseconds.
	ExposureTimeBaseAbs          IFloat        = "ExposureTimeBaseAbs"          // GenApi::IFloat || Sets the time base (in microseconds) that is used when the exposure time is set with the 'exposure time raw' setting.
	ExposureTimeBaseAbsEnable    IBoolean      = "ExposureTimeBaseAbsEnable"    // GenApi::IBoolean || Enables the use of the exposure time base.
	AcquisitionLineRateAbs       IFloat        = "AcquisitionLineRateAbs"       // GenApi::IFloat || Sets the camera's acquisition line rate in lines per second.
	ResultingLinePeriodAbs       IFloat        = "ResultingLinePeriodAbs"       // GenApi::IFloat || Indicates the minimum allowed line acquisition period (in microseconds) given the current settings for the area of interest, exposure time, and bandwidth.
	ResultingLineRateAbs         IFloat        = "ResultingLineRateAbs"         // GenApi::IFloat || Indicates the maximum allowed line acquisition rate (in lines per second) given the current settings for the area of interest, exposure time, and bandwidth.
	AcquisitionFrameRateAbs      IFloat        = "AcquisitionFrameRateAbs"      // GenApi::IFloat || If the acquisition frame rate feature is enabled, this value sets the camera's acquisition frame rate in frames per second.
	AcquisitionFrameRateEnable   IBoolean      = "AcquisitionFrameRateEnable"   // GenApi::IBoolean || Enables setting the camera's acquisition frame rate to a specified value.
	ResultingFramePeriodAbs      IFloat        = "ResultingFramePeriodAbs"      // GenApi::IFloat || Indicates the minimum allowed frame acquisition period (in microseconds) given the current settings for the area of interest, exposure time, and bandwidth.
	ResultingFrameRateAbs        IFloat        = "ResultingFrameRateAbs"        // GenApi::IFloat || Indicates the maximum allowed frame acquisition rate (in frames per second) given the current settings for the area of interest, exposure time, and bandwidth.
	ExposureMode                 IEnumerationT = "ExposureMode"                 // GenApi::IEnumerationT<ExposureModeEnums> || Sets the exposure mode.
	GlobalResetReleaseModeEnable IBoolean      = "GlobalResetReleaseModeEnable" // GenApi::IBoolean || Enable the Global Reset Release Mode.
	ShutterMode                  IEnumerationT = "ShutterMode"                  // GenApi::IEnumerationT<ShutterModeEnums> || Sets the shutter mode.
	ExposureOverlapTimeMaxRaw    IInteger      = "ExposureOverlapTimeMaxRaw"    // GenApi::IInteger || Sets the maximum overlap of the sensor exposure with the sensor readout in TriggerWidth exposure mode in raw units.
	InterlacedIntegrationMode    IEnumerationT = "InterlacedIntegrationMode"    // GenApi::IEnumerationT<InterlacedIntegrationModeEnums> || Selects the Interlaced Integration Mode.
	ExposureOverlapTimeMaxAbs    IFloat        = "ExposureOverlapTimeMaxAbs"    // GenApi::IFloat || Sets the maximum overlap of the sensor exposure with sensor readout in TriggerWidth exposure mode in microseconds.
	ReadoutTimeAbs               IFloat        = "ReadoutTimeAbs"               // GenApi::IFloat || Indicates the sensor readout time given the current settings.
	AcquisitionStatusSelector    IEnumerationT = "AcquisitionStatusSelector"    // GenApi::IEnumerationT<AcquisitionStatusSelectorEnums> || This enumeration is used to select which internal acquisition signal to read using AcquisitionStatus.
	AcquisitionStatus            IBoolean      = "AcquisitionStatus"            // GenApi::IBoolean || Reads the selected acquisition status.
	FrameTimeoutEnable           IBoolean      = "FrameTimeoutEnable"           // GenApi::IBoolean || Enables the frame timeout.
	FrameTimeoutAbs              IFloat        = "FrameTimeoutAbs"              // GenApi::IFloat || Sets the frame timeout in microseconds.

	// DigitalIO - This category includes items used to control the operation of the camera's digital I/O lines
	LineSelector           IEnumerationT = "LineSelector"           // GenApi::IEnumerationT<LineSelectorEnums> || Selects the I/O line to configure.
	LineInverter           IBoolean      = "LineInverter"           // GenApi::IBoolean || Enables the signal inverter function for the selected input or output line.
	LineTermination        IBoolean      = "LineTermination"        // GenApi::IBoolean || Enables the termination resistor for the selected input line.
	LineDebouncerTimeRaw   IInteger      = "LineDebouncerTimeRaw"   // GenApi::IInteger || Sets the raw value of the selected line debouncer time.
	LineDebouncerTimeAbs   IFloat        = "LineDebouncerTimeAbs"   // GenApi::IFloat || Sets the absolute value of the selected line debouncer time in microseconds.
	MinOutPulseWidthRaw    IInteger      = "MinOutPulseWidthRaw"    // GenApi::IInteger || Sets the raw value for the minimum signal width of an output signal.
	MinOutPulseWidthAbs    IFloat        = "MinOutPulseWidthAbs"    // GenApi::IFloat || Sets the absolute value (in microseconds) for the minimum signal width of an output signal.
	LineStatus             IBoolean      = "LineStatus"             // GenApi::IBoolean || Indicates the current logical state for the selected line.
	LineStatusAll          IInteger      = "LineStatusAll"          // GenApi::IInteger || A single bitfield indicating the current logical state of all available line signals at time of polling.
	UserOutputValue        IBoolean      = "UserOutputValue"        // GenApi::IBoolean || Sets the state of the selected user settable output signal.
	UserOutputValueAll     IInteger      = "UserOutputValueAll"     // GenApi::IInteger || A single bitfield that sets the state of all user settable output signals in one access.
	UserOutputValueAllMask IInteger      = "UserOutputValueAllMask" // GenApi::IInteger || Defines a mask that is used when the User Output Value All setting is used to set all of the user settable output signals in one access.
	SyncUserOutputValue    IBoolean      = "SyncUserOutputValue"    // GenApi::IBoolean || Sets the state of the selected user settable synchronous output signal.
	SyncUserOutputValueAll IInteger      = "SyncUserOutputValueAll" // GenApi::IInteger || A single bitfield that sets the state of all user settable synchronous output signals in one access.
	LineMode               IEnumerationT = "LineMode"               // GenApi::IEnumerationT<LineModeEnums> || Sets the mode for the selected line.
	LineSource             IEnumerationT = "LineSource"             // GenApi::IEnumerationT<LineSourceEnums> || Sets the source signal for the selected line (if the selected line is an output)
	LineLogic              IEnumerationT = "LineLogic"              // GenApi::IEnumerationT<LineLogicEnums> || <b>Visibility</b> = Beginner
	LineFormat             IEnumerationT = "LineFormat"             // GenApi::IEnumerationT<LineFormatEnums> || Sets the electrical configuration of the selected line.
	UserOutputSelector     IEnumerationT = "UserOutputSelector"     // GenApi::IEnumerationT<UserOutputSelectorEnums> || Selects the user settable output signal to configure.
	SyncUserOutputSelector IEnumerationT = "SyncUserOutputSelector" // GenApi::IEnumerationT<SyncUserOutputSelectorEnums> || <b>Visibility</b> = Beginner

	// VirtualInput - This category includes items used to control the operation of the camera’s virtual input I/O lines
	VInpSignalSource            IEnumerationT = "VInpSignalSource"            // GenApi::IEnumerationT<VInpSignalSourceEnums> || Sets the I/O line on which the camera receives the virtual input signal.
	VInpBitLength               IInteger      = "VInpBitLength"               // GenApi::IInteger || Sets the length of the input bit.
	VInpSamplingPoint           IInteger      = "VInpSamplingPoint"           // GenApi::IInteger || Time span between the beginning of the input bit and the time when the high/low status is evaluated.
	VInpSignalReadoutActivation IEnumerationT = "VInpSignalReadoutActivation" // GenApi::IEnumerationT<VInpSignalReadoutActivationEnums> || Selects when to start the signal evaluation.

	// ShaftEncoderModule - This category provides controls for operating the camera's shaft encoder module.
	ShaftEncoderModuleLineSelector        IEnumerationT = "ShaftEncoderModuleLineSelector"        // GenApi::IEnumerationT<ShaftEncoderModuleLineSelectorEnums> || Selects the phase of the shaft encoder.
	ShaftEncoderModuleLineSource          IEnumerationT = "ShaftEncoderModuleLineSource"          // GenApi::IEnumerationT<ShaftEncoderModuleLineSourceEnums> || Selects the input line as signal source for the shaft encoder module.
	ShaftEncoderModuleMode                IEnumerationT = "ShaftEncoderModuleMode"                // GenApi::IEnumerationT<ShaftEncoderModuleModeEnums> || Selects the circumstances for the shaft encoder module to output trigger signals.
	ShaftEncoderModuleCounterMode         IEnumerationT = "ShaftEncoderModuleCounterMode"         // GenApi::IEnumerationT<ShaftEncoderModuleCounterModeEnums> || Selects the counting mode of the tick counter.
	ShaftEncoderModuleCounter             IInteger      = "ShaftEncoderModuleCounter"             // GenApi::IInteger || Indicates the current value of the tick counter.
	ShaftEncoderModuleCounterMax          IInteger      = "ShaftEncoderModuleCounterMax"          // GenApi::IInteger || Sets the maximum value for the tick counter.
	ShaftEncoderModuleCounterReset        ICommand      = "ShaftEncoderModuleCounterReset"        // GenApi::ICommand || Resets the tick counter to 0.
	ShaftEncoderModuleReverseCounterMax   IInteger      = "ShaftEncoderModuleReverseCounterMax"   // GenApi::IInteger || Sets the maximum value for the reverse counter.
	ShaftEncoderModuleReverseCounterReset ICommand      = "ShaftEncoderModuleReverseCounterReset" // GenApi::ICommand || Resets the reverse counter to 0.

	// FrequencyConverter - This category includes items used to control the operation of the camera's frequency converter module
	FrequencyConverterInputSource        IEnumerationT = "FrequencyConverterInputSource"        // GenApi::IEnumerationT<FrequencyConverterInputSourceEnums> || Selects the input source.
	FrequencyConverterSignalAlignment    IEnumerationT = "FrequencyConverterSignalAlignment"    // GenApi::IEnumerationT<FrequencyConverterSignalAlignmentEnums> || Selects the signal transition relationships between received and generated signals.
	FrequencyConverterPreDivider         IInteger      = "FrequencyConverterPreDivider"         // GenApi::IInteger || Sets the pre-divider value for the pre-divider sub-module.
	FrequencyConverterMultiplier         IInteger      = "FrequencyConverterMultiplier"         // GenApi::IInteger || Sets the multiplier value for the multiplier sub-module.
	FrequencyConverterPostDivider        IInteger      = "FrequencyConverterPostDivider"        // GenApi::IInteger || Sets the post-divider value for the post-divider sub-module.
	FrequencyConverterPreventOvertrigger IBoolean      = "FrequencyConverterPreventOvertrigger" // GenApi::IBoolean || Enables overtriggering protection.

	// TimerControls - This category includes items used to control the operation of the camera's timers
	TimerDurationTimebaseAbs IFloat        = "TimerDurationTimebaseAbs" // GenApi::IFloat || Sets the time base (in microseconds) that is used when a timer duration is set with the 'timer duration raw' setting.
	TimerDelayTimebaseAbs    IFloat        = "TimerDelayTimebaseAbs"    // GenApi::IFloat || Sets the time base (in microseconds) that is used when a timer delay is set with the 'timer delay raw' setting.
	TimerSelector            IEnumerationT = "TimerSelector"            // GenApi::IEnumerationT<TimerSelectorEnums> || Selects the timer to configure.
	TimerDurationAbs         IFloat        = "TimerDurationAbs"         // GenApi::IFloat || Directly sets the duration for the selected timer in microseconds.
	TimerDurationRaw         IInteger      = "TimerDurationRaw"         // GenApi::IInteger || Sets the 'raw' duration for the selected timer.
	TimerDelayAbs            IFloat        = "TimerDelayAbs"            // GenApi::IFloat || Directly sets the delay for the selected timer in microseconds.
	TimerDelayRaw            IInteger      = "TimerDelayRaw"            // GenApi::IInteger || Sets the 'raw' delay for the selected timer.
	TimerTriggerSource       IEnumerationT = "TimerTriggerSource"       // GenApi::IEnumerationT<TimerTriggerSourceEnums> || Sets the internal camera signal used to trigger the selected timer.
	TimerTriggerActivation   IEnumerationT = "TimerTriggerActivation"   // GenApi::IEnumerationT<TimerTriggerActivationEnums> || Sets the type of signal transistion that will start the timer.
	CounterSelector          IEnumerationT = "CounterSelector"          // GenApi::IEnumerationT<CounterSelectorEnums> || Selects the counter to configure.
	CounterEventSource       IEnumerationT = "CounterEventSource"       // GenApi::IEnumerationT<CounterEventSourceEnums> || Selects the event that will be the source to increment the counter.
	CounterResetSource       IEnumerationT = "CounterResetSource"       // GenApi::IEnumerationT<CounterResetSourceEnums> || Selects the source of the reset for the selected counter.
	CounterReset             ICommand      = "CounterReset"             // GenApi::ICommand || Immediately resets the selected counter.

	// TimerSequence -
	TimerSequenceEnable            IBoolean      = "TimerSequenceEnable"            // GenApi::IBoolean || <b>Visibility</b> = Guru
	TimerSequenceLastEntryIndex    IInteger      = "TimerSequenceLastEntryIndex"    // GenApi::IInteger || <b>Visibility</b> = Guru
	TimerSequenceCurrentEntryIndex IInteger      = "TimerSequenceCurrentEntryIndex" // GenApi::IInteger || <b>Visibility</b> = Guru
	TimerSequenceEntrySelector     IEnumerationT = "TimerSequenceEntrySelector"     // GenApi::IEnumerationT<TimerSequenceEntrySelectorEnums> || <b>Visibility</b> = Guru
	TimerSequenceTimerSelector     IEnumerationT = "TimerSequenceTimerSelector"     // GenApi::IEnumerationT<TimerSequenceTimerSelectorEnums> || <b>Visibility</b> = Guru
	TimerSequenceTimerEnable       IBoolean      = "TimerSequenceTimerEnable"       // GenApi::IBoolean || <b>Visibility</b> = Guru
	TimerSequenceTimerInverter     IBoolean      = "TimerSequenceTimerInverter"     // GenApi::IBoolean || <b>Visibility</b> = Guru
	TimerSequenceTimerDelayRaw     IInteger      = "TimerSequenceTimerDelayRaw"     // GenApi::IInteger || <b>Visibility</b> = Guru
	TimerSequenceTimerDurationRaw  IInteger      = "TimerSequenceTimerDurationRaw"  // GenApi::IInteger || <b>Visibility</b> = Guru

	// LUTControls - This category includes items used to control the operation of the camera's lookup table (LUT)
	LUTEnable   IBoolean      = "LUTEnable"   // GenApi::IBoolean || Enables the selected LUT.
	LUTIndex    IInteger      = "LUTIndex"    // GenApi::IInteger || Sets the LUT element to access.
	LUTValue    IInteger      = "LUTValue"    // GenApi::IInteger || Sets the value of the LUT element at the LUT index.
	LUTValueAll IRegister     = "LUTValueAll" // GenApi::IRegister || Accesses the entire content of the selected LUT in one chunk access.
	LUTSelector IEnumerationT = "LUTSelector" // GenApi::IEnumerationT<LUTSelectorEnums> || Selects the lookup table (LUT) to configure.

	// TransportLayer - This category includes items related to the GigE Vision transport layer
	PixelFormatLegacy                         IBoolean      = "PixelFormatLegacy"                         // GenApi::IBoolean || Select legacy pixel format encoding.
	PayloadSize                               IInteger      = "PayloadSize"                               // GenApi::IInteger || Size of the payload in bytes.
	GevInterfaceSelector                      IEnumerationT = "GevInterfaceSelector"                      // GenApi::IEnumerationT<GevInterfaceSelectorEnums> || Selects the physical network interface to configure.
	GevVersionMajor                           IInteger      = "GevVersionMajor"                           // GenApi::IInteger || Indicates the major version number of the GigE Vision specification supported by this device.
	GevVersionMinor                           IInteger      = "GevVersionMinor"                           // GenApi::IInteger || Indicates the minor version number of the GigE Vision specification supported by this device.
	GevDeviceModeIsBigEndian                  IBoolean      = "GevDeviceModeIsBigEndian"                  // GenApi::IBoolean || Indicates the endianess of the bootstrap registers.
	GevDeviceModeCharacterSet                 IInteger      = "GevDeviceModeCharacterSet"                 // GenApi::IInteger || Indictes the character set.
	GevMACAddress                             IInteger      = "GevMACAddress"                             // GenApi::IInteger || Indicates the MAC address for the selected network interface.
	GevSupportedIPConfigurationLLA            IBoolean      = "GevSupportedIPConfigurationLLA"            // GenApi::IBoolean || Indicates whether the selected network interface supports auto IP addressing (also known as LLA)
	GevSupportedIPConfigurationDHCP           IBoolean      = "GevSupportedIPConfigurationDHCP"           // GenApi::IBoolean || Indicates whether the selected network interface supports DHCP IP addressing.
	GevSupportedIPConfigurationPersistentIP   IBoolean      = "GevSupportedIPConfigurationPersistentIP"   // GenApi::IBoolean || Indicates whether the selected network interface supports fixed IP addressing (also known as persistent IP addressing)
	GevCurrentIPConfiguration                 IInteger      = "GevCurrentIPConfiguration"                 // GenApi::IInteger || Sets the current IP configuration of the selected network interface.
	GevCurrentIPAddress                       IInteger      = "GevCurrentIPAddress"                       // GenApi::IInteger || Indicates the current IP address for the selected network interface.
	GevCurrentSubnetMask                      IInteger      = "GevCurrentSubnetMask"                      // GenApi::IInteger || Indicates the current subnet mask for the selected network interface.
	GevCurrentDefaultGateway                  IInteger      = "GevCurrentDefaultGateway"                  // GenApi::IInteger || Indicates the current default gateway for the selected network interface.
	GevPersistentIPAddress                    IInteger      = "GevPersistentIPAddress"                    // GenApi::IInteger || If fixed (persistent) IP addressing is supported by the device and enabled, sets the fixed IP address for the selected network interface.
	GevPersistentSubnetMask                   IInteger      = "GevPersistentSubnetMask"                   // GenApi::IInteger || If fixed (persistent) IP addressing is supported by the device and enabled, sets the fixed subnet mask for the selected network interface.
	GevPersistentDefaultGateway               IInteger      = "GevPersistentDefaultGateway"               // GenApi::IInteger || If fixed (persistent) IP addressing is supported by the device and enabled, sets the fixed default gateway for the selected network interface.
	GevLinkSpeed                              IInteger      = "GevLinkSpeed"                              // GenApi::IInteger || Indicates the connection speed in Mbps for the selected network interface.
	GevLinkMaster                             IBoolean      = "GevLinkMaster"                             // GenApi::IBoolean || Indicates whether the selected network interface is the clock master.
	GevLinkFullDuplex                         IBoolean      = "GevLinkFullDuplex"                         // GenApi::IBoolean || Indicates whether the selected network interface operates in full-duplex mode.
	GevLinkCrossover                          IBoolean      = "GevLinkCrossover"                          // GenApi::IBoolean || Indicates the state of medium-dependent interface crossover (MDIX) for the selected network interface.
	GevFirstURL                               IString       = "GevFirstURL"                               // GenApi::IString || Indicates the first URL to the XML device description file.
	GevSecondURL                              IString       = "GevSecondURL"                              // GenApi::IString || Indicates the second URL to the XML device description file.
	GevNumberOfInterfaces                     IInteger      = "GevNumberOfInterfaces"                     // GenApi::IInteger || Indicates the number of network interfaces on the device.
	GevMessageChannelCount                    IInteger      = "GevMessageChannelCount"                    // GenApi::IInteger || Indicates the number of message channels supported by the device.
	GevStreamChannelCount                     IInteger      = "GevStreamChannelCount"                     // GenApi::IInteger || Indicates the number of stream channels supported by the device.
	GevSupportedOptionalCommandsEVENTDATA     IBoolean      = "GevSupportedOptionalCommandsEVENTDATA"     // GenApi::IBoolean || Indicates whether EVENTDATA_CMD and EVENTDATA_ACK are supported.
	GevSupportedOptionalCommandsEVENT         IBoolean      = "GevSupportedOptionalCommandsEVENT"         // GenApi::IBoolean || Indicates whether EVENT_CMD and EVENT_ACK are supported.
	GevSupportedOptionalCommandsPACKETRESEND  IBoolean      = "GevSupportedOptionalCommandsPACKETRESEND"  // GenApi::IBoolean || Indicates whether PACKETRESEND_CMD is supported.
	GevSupportedOptionalCommandsWRITEMEM      IBoolean      = "GevSupportedOptionalCommandsWRITEMEM"      // GenApi::IBoolean || Indicates whether WRITEMEM_CMD and WRITEMEM_ACK are supported.
	GevSupportedOptionalCommandsConcatenation IBoolean      = "GevSupportedOptionalCommandsConcatenation" // GenApi::IBoolean || Indicates whether multiple operations in a single message are supported.
	GevHeartbeatTimeout                       IInteger      = "GevHeartbeatTimeout"                       // GenApi::IInteger || Sets the heartbeat timeout in milliseconds.
	GevTimestampTickFrequency                 IInteger      = "GevTimestampTickFrequency"                 // GenApi::IInteger || Indicates the number of timestamp clock ticks in 1 second.
	GevTimestampControlReset                  ICommand      = "GevTimestampControlReset"                  // GenApi::ICommand || Resets the timestamp value for the device.
	GevTimestampControlLatch                  ICommand      = "GevTimestampControlLatch"                  // GenApi::ICommand || Latches the current timestamp value of the device.
	GevTimestampControlLatchReset             ICommand      = "GevTimestampControlLatchReset"             // GenApi::ICommand || Resets the timestamp control latch.
	GevTimestampValue                         IInteger      = "GevTimestampValue"                         // GenApi::IInteger || Indicates the latched value of the timestamp.
	GevCCP                                    IEnumerationT = "GevCCP"                                    // GenApi::IEnumerationT<GevCCPEnums> || Sets the control channel privilege feature.
	GevStreamChannelSelector                  IEnumerationT = "GevStreamChannelSelector"                  // GenApi::IEnumerationT<GevStreamChannelSelectorEnums> || Selects the stream channel to configure.
	GevSCPInterfaceIndex                      IInteger      = "GevSCPInterfaceIndex"                      // GenApi::IInteger || Sets the index of the network interface to use.
	GevSCPHostPort                            IInteger      = "GevSCPHostPort"                            // GenApi::IInteger || Sets the port to which the device must send data streams.
	GevSCPSFireTestPacket                     ICommand      = "GevSCPSFireTestPacket"                     // GenApi::ICommand || <b>Visibility</b> = Guru
	GevSCPSDoNotFragment                      IBoolean      = "GevSCPSDoNotFragment"                      // GenApi::IBoolean || <b>Visibility</b> = Guru
	GevSCPSBigEndian                          IBoolean      = "GevSCPSBigEndian"                          // GenApi::IBoolean || <b>Visibility</b> = Guru
	GevSCPSPacketSize                         IInteger      = "GevSCPSPacketSize"                         // GenApi::IInteger || Sets the packet size in bytes for the selected stream channel.
	GevSCDA                                   IInteger      = "GevSCDA"                                   // GenApi::IInteger || Sets the stream channel destination IPv4 address for the selected stream channel.
	GevSCPD                                   IInteger      = "GevSCPD"                                   // GenApi::IInteger || Sets the inter-packet delay (in ticks) for the selected stream channel.
	GevSCFTD                                  IInteger      = "GevSCFTD"                                  // GenApi::IInteger || Sets the frame transfer start delay (in ticks) for the selected stream channel.
	GevSCBWR                                  IInteger      = "GevSCBWR"                                  // GenApi::IInteger || Sets a percentage of the Ethernet bandwidth assigned to the camera to be held in reserve.
	GevSCBWRA                                 IInteger      = "GevSCBWRA"                                 // GenApi::IInteger || Sets a multiplier for the Bandwidth Reserve parameter.
	GevSCBWA                                  IInteger      = "GevSCBWA"                                  // GenApi::IInteger || Indicates the bandwidth (in bytes per second) that will be used by the camera to transmit image and chunk feature data and to handle resends and control data transmissions.
	GevSCDMT                                  IInteger      = "GevSCDMT"                                  // GenApi::IInteger || Indicates the maximum amount of data (in bytes per second) that the camera could generate given its current settings and ideal conditions, i.e., unlimited bandwidth and no packet resends.
	GevSCDCT                                  IInteger      = "GevSCDCT"                                  // GenApi::IInteger || Indicates the actual bandwidth (in bytes per second) that the camera will use to transmit image data and chunk data given the current AOI settings, chunk feature settings, and the pixel format setting.
	GevSCFJM                                  IInteger      = "GevSCFJM"                                  // GenApi::IInteger || Indicates the maximum time (in ticks) that the next frame transmission could be delayed due to a burst of resends.
	TLParamsLocked                            IInteger      = "TLParamsLocked"                            // GenApi::IInteger || Indicates whether a live grab is under way.

	// ActionControl - This category includes items that control the action control feature
	ActionSelector     IInteger = "ActionSelector"     // GenApi::IInteger || Selects the action command to configure.
	ActionDeviceKey    IInteger = "ActionDeviceKey"    // GenApi::IInteger || Authorization key.
	ActionCommandCount IInteger = "ActionCommandCount" // GenApi::IInteger || Number of action command interfaces.
	ActionGroupKey     IInteger = "ActionGroupKey"     // GenApi::IInteger || Defines a group of devices.
	ActionGroupMask    IInteger = "ActionGroupMask"    // GenApi::IInteger || Filters out particular devices from its group.

	// DeviceControl -
	DeviceRegistersStreamingStart ICommand = "DeviceRegistersStreamingStart" // GenApi::ICommand || Prepare the device for registers streaming.
	DeviceRegistersStreamingEnd   ICommand = "DeviceRegistersStreamingEnd"   // GenApi::ICommand || Announce the end of registers streaming.

	// UserSets - This category includes items that control the configuration sets feature that is used to save sets of parameters in the camera
	UserSetLoad            ICommand      = "UserSetLoad"            // GenApi::ICommand || Loads the selected configuration into the camera's volatile memory and makes it the active configuration set.
	UserSetSave            ICommand      = "UserSetSave"            // GenApi::ICommand || Saves the current active configuration set into the selected user set.
	UserSetSelector        IEnumerationT = "UserSetSelector"        // GenApi::IEnumerationT<UserSetSelectorEnums> || Selects the configuration set to load, save, or configure.
	UserSetDefaultSelector IEnumerationT = "UserSetDefaultSelector" // GenApi::IEnumerationT<UserSetDefaultSelectorEnums> || Sets the configuration set to be used as the default startup set.
	DefaultSetSelector     IEnumerationT = "DefaultSetSelector"     // GenApi::IEnumerationT<DefaultSetSelectorEnums> || Selects the which factory setting will be used as default set.

	// AutoFunctions - This category includes items that parameterize the Auto Functions
	AutoTargetValue                  IInteger      = "AutoTargetValue"                  // GenApi::IInteger || Target average grey value for Gain Auto and Exposure Auto.
	GrayValueAdjustmentDampingAbs    IFloat        = "GrayValueAdjustmentDampingAbs"    // GenApi::IFloat || Gray value adjustment damping for Gain Auto and Exposure Auto.
	GrayValueAdjustmentDampingRaw    IInteger      = "GrayValueAdjustmentDampingRaw"    // GenApi::IInteger || Gray value adjustment damping for Gain Auto and Exposure Auto.
	BalanceWhiteAdjustmentDampingAbs IFloat        = "BalanceWhiteAdjustmentDampingAbs" // GenApi::IFloat || Balance White adjustment damping for Balance White Auto.
	BalanceWhiteAdjustmentDampingRaw IInteger      = "BalanceWhiteAdjustmentDampingRaw" // GenApi::IInteger || Balance White adjustment damping for Balance White Auto.
	AutoGainRawLowerLimit            IInteger      = "AutoGainRawLowerLimit"            // GenApi::IInteger || Lower limit of the Auto Gain (Raw) parameter.
	AutoGainRawUpperLimit            IInteger      = "AutoGainRawUpperLimit"            // GenApi::IInteger || Upper limit of the Auto Gain (Raw) parameter.
	AutoExposureTimeAbsLowerLimit    IFloat        = "AutoExposureTimeAbsLowerLimit"    // GenApi::IFloat || Lower limit of the Auto Exposure Time (Abs) [us] parameter.
	AutoExposureTimeAbsUpperLimit    IFloat        = "AutoExposureTimeAbsUpperLimit"    // GenApi::IFloat || Upper limit of the Auto Exposure Time (Abs) [us] parameter.
	AutoFunctionProfile              IEnumerationT = "AutoFunctionProfile"              // GenApi::IEnumerationT<AutoFunctionProfileEnums> || Selects the strategy for controlling gain and shutter simultaneously.

	// AutoFunctionAOIs - Portion of the sensor array used for auto function control
	AutoFunctionAOIWidth                   IInteger      = "AutoFunctionAOIWidth"                   // GenApi::IInteger || Sets the width of the auto function area of interest in pixels.
	AutoFunctionAOIHeight                  IInteger      = "AutoFunctionAOIHeight"                  // GenApi::IInteger || Sets the height of the auto function area of interest in pixels.
	AutoFunctionAOIOffsetX                 IInteger      = "AutoFunctionAOIOffsetX"                 // GenApi::IInteger || Sets the starting column of the auto function area of interest in pixels.
	AutoFunctionAOIOffsetY                 IInteger      = "AutoFunctionAOIOffsetY"                 // GenApi::IInteger || Sets the starting line of the auto function area of interest in pixels.
	AutoFunctionAOIUsageIntensity          IBoolean      = "AutoFunctionAOIUsageIntensity"          // GenApi::IBoolean || <b>Visibility</b> = Beginner
	AutoFunctionAOIUsageWhiteBalance       IBoolean      = "AutoFunctionAOIUsageWhiteBalance"       // GenApi::IBoolean || <b>Visibility</b> = Beginner
	AutoFunctionAOIUsageRedLightCorrection IBoolean      = "AutoFunctionAOIUsageRedLightCorrection" // GenApi::IBoolean || <b>Visibility</b> = Beginner
	AutoFunctionAOISelector                IEnumerationT = "AutoFunctionAOISelector"                // GenApi::IEnumerationT<AutoFunctionAOISelectorEnums> || Selects the Auto Function AOI.

	// ColorOverexposureCompensation - Compensates for deviations of hue resulting from overexposure
	ColorOverexposureCompensationAOISelector  IEnumerationT = "ColorOverexposureCompensationAOISelector"  // GenApi::IEnumerationT<ColorOverexposureCompensationAOISelectorEnums> || Selcts the AOI for color overexposure compensation.
	ColorOverexposureCompensationAOIEnable    IBoolean      = "ColorOverexposureCompensationAOIEnable"    // GenApi::IBoolean || Enables color overexposure compensation.
	ColorOverexposureCompensationAOIFactor    IFloat        = "ColorOverexposureCompensationAOIFactor"    // GenApi::IFloat || Sets the color overexposure compensation factor for the selected C.O.C.
	ColorOverexposureCompensationAOIFactorRaw IInteger      = "ColorOverexposureCompensationAOIFactorRaw" // GenApi::IInteger || Sets the raw value for the color overexposure compensation factor.
	ColorOverexposureCompensationAOIWidth     IInteger      = "ColorOverexposureCompensationAOIWidth"     // GenApi::IInteger || Sets the width for the selected C.O.C.
	ColorOverexposureCompensationAOIHeight    IInteger      = "ColorOverexposureCompensationAOIHeight"    // GenApi::IInteger || Sets the height for the selected C.O.C.
	ColorOverexposureCompensationAOIOffsetX   IInteger      = "ColorOverexposureCompensationAOIOffsetX"   // GenApi::IInteger || Sets the X offset for the selected C.O.C.
	ColorOverexposureCompensationAOIOffsetY   IInteger      = "ColorOverexposureCompensationAOIOffsetY"   // GenApi::IInteger || Sets the Y offset for the selected C.O.C.

	// Shading - Includes items used to control the operation of the camera's shading correction.
	ShadingSelector           IEnumerationT = "ShadingSelector"           // GenApi::IEnumerationT<ShadingSelectorEnums> || Selects the kind of shading correction.
	ShadingEnable             IBoolean      = "ShadingEnable"             // GenApi::IBoolean || Enables the selected kind of shading correction.
	ShadingStatus             IEnumerationT = "ShadingStatus"             // GenApi::IEnumerationT<ShadingStatusEnums> || Indicates error statuses related to shading correction.
	ShadingSetDefaultSelector IEnumerationT = "ShadingSetDefaultSelector" // GenApi::IEnumerationT<ShadingSetDefaultSelectorEnums> || Selects the bootup shading set.
	ShadingSetSelector        IEnumerationT = "ShadingSetSelector"        // GenApi::IEnumerationT<ShadingSetSelectorEnums> || Selects the shading set to which the activate command will be applied.
	ShadingSetActivate        ICommand      = "ShadingSetActivate"        // GenApi::ICommand || Activates the selected shading set.
	ShadingSetCreate          IEnumerationT = "ShadingSetCreate"          // GenApi::IEnumerationT<ShadingSetCreateEnums> || Creates a shading set.

	// UserDefinedValues -
	UserDefinedValueSelector IEnumerationT = "UserDefinedValueSelector" // GenApi::IEnumerationT<UserDefinedValueSelectorEnums> || <b>Visibility</b> = Guru
	UserDefinedValue         IInteger      = "UserDefinedValue"         // GenApi::IInteger || <b>Visibility</b> = Guru

	// FeatureSets -
	GenicamXmlFileDefault IInteger      = "GenicamXmlFileDefault" // GenApi::IInteger || Select default genicam XML file.
	FeatureSet            IEnumerationT = "FeatureSet"            // GenApi::IEnumerationT<FeatureSetEnums> || Select a camera description file.

	// RemoveParamLimits - This category includes items that allow removing the limits of camera parameters
	RemoveLimits      IBoolean      = "RemoveLimits"      // GenApi::IBoolean || Removes the factory-set limits of the selected parameter.
	ParameterSelector IEnumerationT = "ParameterSelector" // GenApi::IEnumerationT<ParameterSelectorEnums> || Selects the parameter to configure.
	Prelines          IInteger      = "Prelines"          // GenApi::IInteger || Sets the number of prelines.

	// ExpertFeatureAccess -
	ExpertFeatureAccessSelector IEnumerationT = "ExpertFeatureAccessSelector" // GenApi::IEnumerationT<ExpertFeatureAccessSelectorEnums> || Selects the feature to configure.
	ExpertFeatureAccessKey      IInteger      = "ExpertFeatureAccessKey"      // GenApi::IInteger || Sets the key to access the selected feature.
	ExpertFeatureEnable         IBoolean      = "ExpertFeatureEnable"         // GenApi::IBoolean || Enable the selected Feature.

	// ChunkDataStreams - This category includes items that control the chunk features available on the camera.
	ChunkModeActive IBoolean      = "ChunkModeActive" // GenApi::IBoolean || Enables the chunk mode.
	ChunkEnable     IBoolean      = "ChunkEnable"     // GenApi::IBoolean || Enables the inclusion of the selected chunk in the payload data.
	ChunkSelector   IEnumerationT = "ChunkSelector"   // GenApi::IEnumerationT<ChunkSelectorEnums> || Selects chunks for enabling.

	// ChunkData - This category includes items related to the chunk data that can be appended to the image data
	ChunkStride                              IInteger      = "ChunkStride"                              // GenApi::IInteger || Indicates the number of bytes of data between the beginning of one line in the acquired image and the beginning of the next line in the acquired image.
	ChunkSequenceSetIndex                    IInteger      = "ChunkSequenceSetIndex"                    // GenApi::IInteger || Indicates the sequence set index number related to the acquired image.
	ChunkOffsetX                             IInteger      = "ChunkOffsetX"                             // GenApi::IInteger || Indicates the X offset of the area of interest represented in the acquired image.
	ChunkOffsetY                             IInteger      = "ChunkOffsetY"                             // GenApi::IInteger || Indicates the Y offset of the area of interest represented in the acquired image.
	ChunkWidth                               IInteger      = "ChunkWidth"                               // GenApi::IInteger || Indicates the widtth of the area of interest represented in the acquired image.
	ChunkHeight                              IInteger      = "ChunkHeight"                              // GenApi::IInteger || Indicates the height of the area of interest represented in the acquired image.
	ChunkDynamicRangeMin                     IInteger      = "ChunkDynamicRangeMin"                     // GenApi::IInteger || Indicates the minimum possible pixel value in the acquired image.
	ChunkDynamicRangeMax                     IInteger      = "ChunkDynamicRangeMax"                     // GenApi::IInteger || Indicates the maximum possible pixel value in the acquired image.
	ChunkPixelFormat                         IEnumerationT = "ChunkPixelFormat"                         // GenApi::IEnumerationT<ChunkPixelFormatEnums> || Indicates the format of the pixel data in the acquired image.
	ChunkTimestamp                           IInteger      = "ChunkTimestamp"                           // GenApi::IInteger || Indicates the value of the timestamp when the image was acquired.
	ChunkFramecounter                        IInteger      = "ChunkFramecounter"                        // GenApi::IInteger || Indicates the value of the frame counter when the image was acquired.
	ChunkLineStatusAll                       IInteger      = "ChunkLineStatusAll"                       // GenApi::IInteger || A bit field that indicates the status of all of the camera's input and output lines when the image was acquired.
	ChunkVirtLineStatusAll                   IInteger      = "ChunkVirtLineStatusAll"                   // GenApi::IInteger || A bit field that indicates the status of all of the camera's virtual input and output lines when the image was acquired.
	ChunkTriggerinputcounter                 IInteger      = "ChunkTriggerinputcounter"                 // GenApi::IInteger || Indicates the value of the trigger input counter when the image was acquired.
	ChunkLineTriggerIgnoredCounter           IInteger      = "ChunkLineTriggerIgnoredCounter"           // GenApi::IInteger || <b>Visibility</b> = Beginner
	ChunkFrameTriggerIgnoredCounter          IInteger      = "ChunkFrameTriggerIgnoredCounter"          // GenApi::IInteger || <b>Visibility</b> = Beginner
	ChunkLineTriggerEndToEndCounter          IInteger      = "ChunkLineTriggerEndToEndCounter"          // GenApi::IInteger || <b>Visibility</b> = Beginner
	ChunkFrameTriggerCounter                 IInteger      = "ChunkFrameTriggerCounter"                 // GenApi::IInteger || <b>Visibility</b> = Beginner
	ChunkFramesPerTriggerCounter             IInteger      = "ChunkFramesPerTriggerCounter"             // GenApi::IInteger || <b>Visibility</b> = Beginner
	ChunkInputStatusAtLineTriggerBitsPerLine IInteger      = "ChunkInputStatusAtLineTriggerBitsPerLine" // GenApi::IInteger || Number of bits per status.
	ChunkInputStatusAtLineTriggerIndex       IInteger      = "ChunkInputStatusAtLineTriggerIndex"       // GenApi::IInteger || Used to select a certain status.
	ChunkInputStatusAtLineTriggerValue       IInteger      = "ChunkInputStatusAtLineTriggerValue"       // GenApi::IInteger || Value of the status selected by 'Index'.
	ChunkShaftEncoderCounter                 IInteger      = "ChunkShaftEncoderCounter"                 // GenApi::IInteger || Shaft encoder counter at frame trigger.
	ChunkPayloadCRC16                        IInteger      = "ChunkPayloadCRC16"                        // GenApi::IInteger || Indicates the value of CRC checksum.
	ChunkExposureTime                        IFloat        = "ChunkExposureTime"                        // GenApi::IFloat || <b>Visibility</b> = Beginner
	ChunkGainAll                             IInteger      = "ChunkGainAll"                             // GenApi::IInteger || <b>Visibility</b> = Beginner

	// EventsGeneration - This category includes items that control event generation by the camera.
	EventSelector     IEnumerationT = "EventSelector"     // GenApi::IEnumerationT<EventSelectorEnums> || Selects the type of event for enabling.
	EventNotification IEnumerationT = "EventNotification" // GenApi::IEnumerationT<EventNotificationEnums> || Sets the notification type that will be sent to the host application for the selected event.

	// ExposureEndEventData - This category includes items available for an exposure end event
	ExposureEndEventStreamChannelIndex IInteger = "ExposureEndEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an exposure end event.
	ExposureEndEventFrameID            IInteger = "ExposureEndEventFrameID"            // GenApi::IInteger || Indicates the frame ID for an exposure end event.
	ExposureEndEventTimestamp          IInteger = "ExposureEndEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for an exposure end event.

	// LineStartOvertriggerEventData - This category includes items available for an line start overtrigger event
	LineStartOvertriggerEventStreamChannelIndex IInteger = "LineStartOvertriggerEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an line start overtrigger event.
	LineStartOvertriggerEventTimestamp          IInteger = "LineStartOvertriggerEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for an line start overtrigger event.

	// FrameStartOvertriggerEventData - This category includes items available for an frame start overtrigger event
	FrameStartOvertriggerEventStreamChannelIndex IInteger = "FrameStartOvertriggerEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an frame start overtrigger event.
	FrameStartOvertriggerEventTimestamp          IInteger = "FrameStartOvertriggerEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for an frame start overtrigger event.

	// FrameStartEventData - This category includes items available for an frame start event
	FrameStartEventStreamChannelIndex IInteger = "FrameStartEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an frame start event.
	FrameStartEventTimestamp          IInteger = "FrameStartEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for an frame start event.

	// AcquisitionStartEventData - This category includes items available for an acquisition start event
	AcquisitionStartEventStreamChannelIndex IInteger = "AcquisitionStartEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an acquisition start event.
	AcquisitionStartEventTimestamp          IInteger = "AcquisitionStartEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for an acquisition start event.

	// AcquisitionStartOvertriggerEventData - This category includes items available for an acquisition start overtrigger event
	AcquisitionStartOvertriggerEventStreamChannelIndex IInteger = "AcquisitionStartOvertriggerEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an acquisition start overtrigger event.
	AcquisitionStartOvertriggerEventTimestamp          IInteger = "AcquisitionStartOvertriggerEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for an Acquisition start overtrigger event.

	// FrameTimeoutEventData - This category includes items available for an frame timeout event
	FrameTimeoutEventStreamChannelIndex IInteger = "FrameTimeoutEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an frame timeout event.
	FrameTimeoutEventTimestamp          IInteger = "FrameTimeoutEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for an frame timeout event.

	// EventOverrunEventData - This category includes items available for an event overrun event
	EventOverrunEventStreamChannelIndex IInteger = "EventOverrunEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an event overrun event.
	EventOverrunEventFrameID            IInteger = "EventOverrunEventFrameID"            // GenApi::IInteger || Indicates the frame ID for an event overrun event.
	EventOverrunEventTimestamp          IInteger = "EventOverrunEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for an event overrun event.

	// CriticalTemperatureEventData - This category includes items available for a critical temperature event
	CriticalTemperatureEventStreamChannelIndex IInteger = "CriticalTemperatureEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for a critical temperature event.
	CriticalTemperatureEventTimestamp          IInteger = "CriticalTemperatureEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for a critical temperature event.

	// OverTemperatureEventData - This category includes items available for an over temperature event
	OverTemperatureEventStreamChannelIndex IInteger = "OverTemperatureEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an over temperature event.
	OverTemperatureEventTimestamp          IInteger = "OverTemperatureEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for an over temperature event.

	// Line1RisingEdgeEventData - This category includes items available for an io line 1 rising edge event
	Line1RisingEdgeEventStreamChannelIndex IInteger = "Line1RisingEdgeEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an io line 1 rising edge event.
	Line1RisingEdgeEventTimestamp          IInteger = "Line1RisingEdgeEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for a line 1 rising edge event.

	// Line2RisingEdgeEventData - This category includes items available for an io line 2 rising edge event
	Line2RisingEdgeEventStreamChannelIndex IInteger = "Line2RisingEdgeEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an io line 2 rising edge event.
	Line2RisingEdgeEventTimestamp          IInteger = "Line2RisingEdgeEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for a line 2 rising edge event.

	// Line3RisingEdgeEventData - This category includes items available for an io line 3 rising edge event
	Line3RisingEdgeEventStreamChannelIndex IInteger = "Line3RisingEdgeEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an io line 3 rising edge event.
	Line3RisingEdgeEventTimestamp          IInteger = "Line3RisingEdgeEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for a line 3 rising edge event.

	// Line4RisingEdgeEventData - This category includes items available for an io line 4 rising edge event
	Line4RisingEdgeEventStreamChannelIndex IInteger = "Line4RisingEdgeEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an io line 4 rising edge event.
	Line4RisingEdgeEventTimestamp          IInteger = "Line4RisingEdgeEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for a line 4 rising edge event.

	// VirtualLine1RisingEdgeEventData - This category includes items available for an io virtual line 1 rising edge event
	VirtualLine1RisingEdgeEventStreamChannelIndex IInteger = "VirtualLine1RisingEdgeEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an io virtual line 1 rising edge event.
	VirtualLine1RisingEdgeEventTimestamp          IInteger = "VirtualLine1RisingEdgeEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for a virtual line 1 rising edge event.

	// VirtualLine2RisingEdgeEventData - This category includes items available for an io virtual line 2 rising edge event
	VirtualLine2RisingEdgeEventStreamChannelIndex IInteger = "VirtualLine2RisingEdgeEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an io virtual line 2 rising edge event.
	VirtualLine2RisingEdgeEventTimestamp          IInteger = "VirtualLine2RisingEdgeEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for a virtual line 2 rising edge event.

	// VirtualLine3RisingEdgeEventData - This category includes items available for an io virtual line 3 rising edge event
	VirtualLine3RisingEdgeEventStreamChannelIndex IInteger = "VirtualLine3RisingEdgeEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an io virtual line 3 rising edge event.
	VirtualLine3RisingEdgeEventTimestamp          IInteger = "VirtualLine3RisingEdgeEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for a virtual line 3 rising edge event.

	// VirtualLine4RisingEdgeEventData - This category includes items available for an io virtual line 4 rising edge event
	VirtualLine4RisingEdgeEventStreamChannelIndex IInteger = "VirtualLine4RisingEdgeEventStreamChannelIndex" // GenApi::IInteger || Indicates the stream channel index for an io virtual line 4 rising edge event.
	VirtualLine4RisingEdgeEventTimestamp          IInteger = "VirtualLine4RisingEdgeEventTimestamp"          // GenApi::IInteger || Indicates the time stamp for a virtual line 4 rising edge event.

	// FileAccessControl - This category includes items used to conduct file operations
	FileSelector          IEnumerationT = "FileSelector"          // GenApi::IEnumerationT<FileSelectorEnums> || This feature selects the target file in the device.
	FileOperationSelector IEnumerationT = "FileOperationSelector" // GenApi::IEnumerationT<FileOperationSelectorEnums> || Selects the target operation for the selected file.
	FileOpenMode          IEnumerationT = "FileOpenMode"          // GenApi::IEnumerationT<FileOpenModeEnums> || Selects the access mode in which a file is opened.
	FileAccessBuffer      IRegister     = "FileAccessBuffer"      // GenApi::IRegister || Defines the intermediate access buffer.
	FileAccessOffset      IInteger      = "FileAccessOffset"      // GenApi::IInteger || Controls the mapping between the device file storage and the FileAccessBuffer.
	FileAccessLength      IInteger      = "FileAccessLength"      // GenApi::IInteger || Controls the mapping between the device file storage and the FileAccessBuffer.
	FileOperationStatus   IEnumerationT = "FileOperationStatus"   // GenApi::IEnumerationT<FileOperationStatusEnums> || Represents the file operation execution status.
	FileOperationResult   IInteger      = "FileOperationResult"   // GenApi::IInteger || Represents the file operation result.
	FileSize              IInteger      = "FileSize"              // GenApi::IInteger || Represents the size of the selected file.
	FileOperationExecute  ICommand      = "FileOperationExecute"  // GenApi::ICommand || Executes the selected operation.
)
