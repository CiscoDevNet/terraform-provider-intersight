# NiatelemetryModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Module"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Module"]
**AssetId** | Pointer to **string** | Asset identifier of the switch hardware module. | [optional] [readonly] 
**HardwareRevision** | Pointer to **string** | Hardware revision of the switch hardware module. | [optional] [readonly] 
**ModelName** | Pointer to **string** | Hardware model name of the switch hardware module. | [optional] [readonly] 
**ModuleType** | Pointer to **[]string** |  | [optional] 
**ModuleVersion** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Display name of the switch hardware module. | [optional] [readonly] 
**OperStatus** | Pointer to **string** | Operational state of the switch hardware module. * &#x60;unknown&#x60; - The switch hardware module operational state is not known. * &#x60;boot&#x60; - The switch hardware module is booting and is not yet operational. * &#x60;diagFailed&#x60; - The switch hardware module failed its diagnostic checks. * &#x60;disabled&#x60; - The switch hardware module is disabled and is not in service. * &#x60;dormant&#x60; - The switch hardware module is dormant and is not actively operating. * &#x60;down&#x60; - The switch hardware module is down and is not passing traffic. * &#x60;failed&#x60; - The switch hardware module has failed and is not operational. * &#x60;fwDownloadFailure&#x60; - Firmware download for the switch hardware module did not complete successfully. * &#x60;fwDownloadSuccess&#x60; - Firmware download for the switch hardware module completed successfully. * &#x60;fwMismatchFound&#x60; - The switch hardware module firmware does not match the expected version. * &#x60;goingDown&#x60; - The switch hardware module is transitioning to a down operational state. * &#x60;mdr&#x60; - The switch hardware module is performing a minimal disruptive restart. * &#x60;mismatchConfig&#x60; - The switch hardware module configuration does not match the expected configuration. * &#x60;mismatchWithParent&#x60; - The switch hardware module does not match its parent hardware configuration. * &#x60;missing&#x60; - The expected switch hardware module is missing from its slot. * &#x60;nonOperational&#x60; - The switch hardware module is present but is not operational. * &#x60;notApplicable&#x60; - An operational state does not apply to this switch hardware module. * &#x60;notOk&#x60; - The switch hardware module is not in an acceptable operational state. * &#x60;offAdmin&#x60; - The switch hardware module is powered off by an administrative action. * &#x60;offConnectorRating&#x60; - The switch hardware module is off because of a connector rating limit. * &#x60;offCooling&#x60; - The switch hardware module is off because of a cooling condition. * &#x60;offDenied&#x60; - The switch hardware module was denied power and remains off. * &#x60;offEnvFan&#x60; - The switch hardware module is off because of a fan environmental fault. * &#x60;offEnvOther&#x60; - The switch hardware module is off because of another environmental fault. * &#x60;offEnvPower&#x60; - The switch hardware module is off because of a power environmental fault. * &#x60;offEnvTemp&#x60; - The switch hardware module is off because of a temperature environmental fault. * &#x60;offline&#x60; - The switch hardware module is offline and is not available for service. * &#x60;ok&#x60; - The switch hardware module is operating in a normal state. * &#x60;okButAuthFailed&#x60; - The switch hardware module is up but its authentication check failed. * &#x60;okButDiagFailed&#x60; - The switch hardware module is up but a diagnostic check failed. * &#x60;okButPowerOverCritical&#x60; - The switch hardware module is up but power use is above the critical threshold. * &#x60;okButPowerOverWarning&#x60; - The switch hardware module is up but power use is above the warning threshold. * &#x60;on&#x60; - The switch hardware module is powered on. * &#x60;onButFanFail&#x60; - The switch hardware module is on but a fan has failed. * &#x60;onButInlinePowerFail&#x60; - The switch hardware module is on but inline power has failed. * &#x60;online&#x60; - The switch hardware module is online and available for service. * &#x60;outOfServiceAdmin&#x60; - The switch hardware module was taken out of service by an administrator. * &#x60;outOfServiceEnvTemp&#x60; - The switch hardware module is out of service because of a temperature fault. * &#x60;powerCycled&#x60; - The switch hardware module was power cycled and is recovering. * &#x60;powerDenied&#x60; - Power was denied to the switch hardware module. * &#x60;poweredDown&#x60; - The switch hardware module is powered down and is not operating. * &#x60;poweredUp&#x60; - The switch hardware module is powered up. * &#x60;selfTest&#x60; - The switch hardware module is running its self-test sequence. * &#x60;shut&#x60; - The switch hardware module has been shut down and is not in service. * &#x60;syncInProgress&#x60; - The switch hardware module is synchronizing and is not fully ready. * &#x60;unavailable&#x60; - The switch hardware module operational state is currently unavailable. * &#x60;up&#x60; - The switch hardware module is up and available for operation. * &#x60;upgrading&#x60; - The switch hardware module software or firmware upgrade is in progress. | [optional] [readonly] [default to "unknown"]
**SerialNumber** | Pointer to **string** | Serial number of the switch hardware module. | [optional] [readonly] 
**Slot** | Pointer to **string** | Slot identifier of the switch hardware module. | [optional] [readonly] 
**SoftwareRevision** | Pointer to **string** | Software revision of the switch hardware module. | [optional] [readonly] 
**Type** | Pointer to **string** | Hardware category of the switch hardware module. * &#x60;module&#x60; - The switch hardware module is a generic hardware component in the switch. * &#x60;lineCard&#x60; - The switch hardware module is a line card installed in the switch. * &#x60;powerSupply&#x60; - The switch hardware module is a power supply installed in the switch. * &#x60;fan&#x60; - The switch hardware module is a fan or fan tray installed in the switch. * &#x60;chassis&#x60; - The switch hardware module represents the switch chassis hardware. | [optional] [readonly] [default to "module"]

## Methods

### NewNiatelemetryModule

`func NewNiatelemetryModule(classId string, objectType string, ) *NiatelemetryModule`

NewNiatelemetryModule instantiates a new NiatelemetryModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryModuleWithDefaults

`func NewNiatelemetryModuleWithDefaults() *NiatelemetryModule`

NewNiatelemetryModuleWithDefaults instantiates a new NiatelemetryModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryModule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryModule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryModule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryModule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryModule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryModule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssetId

`func (o *NiatelemetryModule) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *NiatelemetryModule) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *NiatelemetryModule) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *NiatelemetryModule) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetHardwareRevision

`func (o *NiatelemetryModule) GetHardwareRevision() string`

GetHardwareRevision returns the HardwareRevision field if non-nil, zero value otherwise.

### GetHardwareRevisionOk

`func (o *NiatelemetryModule) GetHardwareRevisionOk() (*string, bool)`

GetHardwareRevisionOk returns a tuple with the HardwareRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareRevision

`func (o *NiatelemetryModule) SetHardwareRevision(v string)`

SetHardwareRevision sets HardwareRevision field to given value.

### HasHardwareRevision

`func (o *NiatelemetryModule) HasHardwareRevision() bool`

HasHardwareRevision returns a boolean if a field has been set.

### GetModelName

`func (o *NiatelemetryModule) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *NiatelemetryModule) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *NiatelemetryModule) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *NiatelemetryModule) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModuleType

`func (o *NiatelemetryModule) GetModuleType() []string`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *NiatelemetryModule) GetModuleTypeOk() (*[]string, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *NiatelemetryModule) SetModuleType(v []string)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *NiatelemetryModule) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *NiatelemetryModule) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *NiatelemetryModule) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetModuleVersion

`func (o *NiatelemetryModule) GetModuleVersion() []string`

GetModuleVersion returns the ModuleVersion field if non-nil, zero value otherwise.

### GetModuleVersionOk

`func (o *NiatelemetryModule) GetModuleVersionOk() (*[]string, bool)`

GetModuleVersionOk returns a tuple with the ModuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleVersion

`func (o *NiatelemetryModule) SetModuleVersion(v []string)`

SetModuleVersion sets ModuleVersion field to given value.

### HasModuleVersion

`func (o *NiatelemetryModule) HasModuleVersion() bool`

HasModuleVersion returns a boolean if a field has been set.

### SetModuleVersionNil

`func (o *NiatelemetryModule) SetModuleVersionNil(b bool)`

 SetModuleVersionNil sets the value for ModuleVersion to be an explicit nil

### UnsetModuleVersion
`func (o *NiatelemetryModule) UnsetModuleVersion()`

UnsetModuleVersion ensures that no value is present for ModuleVersion, not even an explicit nil
### GetName

`func (o *NiatelemetryModule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryModule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryModule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryModule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperStatus

`func (o *NiatelemetryModule) GetOperStatus() string`

GetOperStatus returns the OperStatus field if non-nil, zero value otherwise.

### GetOperStatusOk

`func (o *NiatelemetryModule) GetOperStatusOk() (*string, bool)`

GetOperStatusOk returns a tuple with the OperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStatus

`func (o *NiatelemetryModule) SetOperStatus(v string)`

SetOperStatus sets OperStatus field to given value.

### HasOperStatus

`func (o *NiatelemetryModule) HasOperStatus() bool`

HasOperStatus returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryModule) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryModule) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryModule) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryModule) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSlot

`func (o *NiatelemetryModule) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *NiatelemetryModule) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *NiatelemetryModule) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *NiatelemetryModule) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetSoftwareRevision

`func (o *NiatelemetryModule) GetSoftwareRevision() string`

GetSoftwareRevision returns the SoftwareRevision field if non-nil, zero value otherwise.

### GetSoftwareRevisionOk

`func (o *NiatelemetryModule) GetSoftwareRevisionOk() (*string, bool)`

GetSoftwareRevisionOk returns a tuple with the SoftwareRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRevision

`func (o *NiatelemetryModule) SetSoftwareRevision(v string)`

SetSoftwareRevision sets SoftwareRevision field to given value.

### HasSoftwareRevision

`func (o *NiatelemetryModule) HasSoftwareRevision() bool`

HasSoftwareRevision returns a boolean if a field has been set.

### GetType

`func (o *NiatelemetryModule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NiatelemetryModule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NiatelemetryModule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NiatelemetryModule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


