# OsServerConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.ServerConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.ServerConfig"]
**AdditionalParameters** | Pointer to [**[]OsPlaceHolder**](OsPlaceHolder.md) |  | [optional] 
**Answers** | Pointer to [**NullableOsAnswers**](OsAnswers.md) |  | [optional] 
**ErrorMsgs** | Pointer to **[]string** |  | [optional] 
**InitiatorWwpn** | Pointer to **string** | The WWPN Address of the underlying fibre channel interface at the host side used for SAN accesss. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 20:00:D4:C9:3C:35:02:01. | [optional] [readonly] 
**InstallTarget** | Pointer to **string** | The target in which OS installation triggered, the value represented is StorageControllerID follwed by PhysicalDisk SerialNumber in case of Physcial disk or VirtualDriveId for virtual drive. | [optional] [readonly] 
**LunId** | Pointer to **int64** | The Logical Unit Number (LUN) of the install target. | [optional] [readonly] [default to 0]
**OperatingSystemParameters** | Pointer to [**NullableOsOperatingSystemParameters**](OsOperatingSystemParameters.md) |  | [optional] 
**ProcessedInstallTarget** | Pointer to [**NullableOsInstallTarget**](OsInstallTarget.md) |  | [optional] 
**SerialNumber** | Pointer to **string** | The Serial Number of the server. | [optional] [readonly] 
**TargetIqn** | Pointer to **string** | IQN (iSCSI qualified name) of Storage iSCSI target. Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique_name. | [optional] [readonly] 
**TargetWwpn** | Pointer to **string** | The WWPN Address of the underlying fibre channel interface at the target used by the storage. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 51:4F:0C:50:14:1F:AF:01. | [optional] [readonly] 
**VnicMac** | Pointer to **string** | MAC address of the VNIC used as iSCSI initiator interface. | [optional] [readonly] 

## Methods

### NewOsServerConfigAllOf

`func NewOsServerConfigAllOf(classId string, objectType string, ) *OsServerConfigAllOf`

NewOsServerConfigAllOf instantiates a new OsServerConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsServerConfigAllOfWithDefaults

`func NewOsServerConfigAllOfWithDefaults() *OsServerConfigAllOf`

NewOsServerConfigAllOfWithDefaults instantiates a new OsServerConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsServerConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsServerConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsServerConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsServerConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsServerConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsServerConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalParameters

`func (o *OsServerConfigAllOf) GetAdditionalParameters() []OsPlaceHolder`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *OsServerConfigAllOf) GetAdditionalParametersOk() (*[]OsPlaceHolder, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *OsServerConfigAllOf) SetAdditionalParameters(v []OsPlaceHolder)`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *OsServerConfigAllOf) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.

### SetAdditionalParametersNil

`func (o *OsServerConfigAllOf) SetAdditionalParametersNil(b bool)`

 SetAdditionalParametersNil sets the value for AdditionalParameters to be an explicit nil

### UnsetAdditionalParameters
`func (o *OsServerConfigAllOf) UnsetAdditionalParameters()`

UnsetAdditionalParameters ensures that no value is present for AdditionalParameters, not even an explicit nil
### GetAnswers

`func (o *OsServerConfigAllOf) GetAnswers() OsAnswers`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *OsServerConfigAllOf) GetAnswersOk() (*OsAnswers, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *OsServerConfigAllOf) SetAnswers(v OsAnswers)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *OsServerConfigAllOf) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### SetAnswersNil

`func (o *OsServerConfigAllOf) SetAnswersNil(b bool)`

 SetAnswersNil sets the value for Answers to be an explicit nil

### UnsetAnswers
`func (o *OsServerConfigAllOf) UnsetAnswers()`

UnsetAnswers ensures that no value is present for Answers, not even an explicit nil
### GetErrorMsgs

`func (o *OsServerConfigAllOf) GetErrorMsgs() []string`

GetErrorMsgs returns the ErrorMsgs field if non-nil, zero value otherwise.

### GetErrorMsgsOk

`func (o *OsServerConfigAllOf) GetErrorMsgsOk() (*[]string, bool)`

GetErrorMsgsOk returns a tuple with the ErrorMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsgs

`func (o *OsServerConfigAllOf) SetErrorMsgs(v []string)`

SetErrorMsgs sets ErrorMsgs field to given value.

### HasErrorMsgs

`func (o *OsServerConfigAllOf) HasErrorMsgs() bool`

HasErrorMsgs returns a boolean if a field has been set.

### SetErrorMsgsNil

`func (o *OsServerConfigAllOf) SetErrorMsgsNil(b bool)`

 SetErrorMsgsNil sets the value for ErrorMsgs to be an explicit nil

### UnsetErrorMsgs
`func (o *OsServerConfigAllOf) UnsetErrorMsgs()`

UnsetErrorMsgs ensures that no value is present for ErrorMsgs, not even an explicit nil
### GetInitiatorWwpn

`func (o *OsServerConfigAllOf) GetInitiatorWwpn() string`

GetInitiatorWwpn returns the InitiatorWwpn field if non-nil, zero value otherwise.

### GetInitiatorWwpnOk

`func (o *OsServerConfigAllOf) GetInitiatorWwpnOk() (*string, bool)`

GetInitiatorWwpnOk returns a tuple with the InitiatorWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorWwpn

`func (o *OsServerConfigAllOf) SetInitiatorWwpn(v string)`

SetInitiatorWwpn sets InitiatorWwpn field to given value.

### HasInitiatorWwpn

`func (o *OsServerConfigAllOf) HasInitiatorWwpn() bool`

HasInitiatorWwpn returns a boolean if a field has been set.

### GetInstallTarget

`func (o *OsServerConfigAllOf) GetInstallTarget() string`

GetInstallTarget returns the InstallTarget field if non-nil, zero value otherwise.

### GetInstallTargetOk

`func (o *OsServerConfigAllOf) GetInstallTargetOk() (*string, bool)`

GetInstallTargetOk returns a tuple with the InstallTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallTarget

`func (o *OsServerConfigAllOf) SetInstallTarget(v string)`

SetInstallTarget sets InstallTarget field to given value.

### HasInstallTarget

`func (o *OsServerConfigAllOf) HasInstallTarget() bool`

HasInstallTarget returns a boolean if a field has been set.

### GetLunId

`func (o *OsServerConfigAllOf) GetLunId() int64`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *OsServerConfigAllOf) GetLunIdOk() (*int64, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *OsServerConfigAllOf) SetLunId(v int64)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *OsServerConfigAllOf) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### GetOperatingSystemParameters

`func (o *OsServerConfigAllOf) GetOperatingSystemParameters() OsOperatingSystemParameters`

GetOperatingSystemParameters returns the OperatingSystemParameters field if non-nil, zero value otherwise.

### GetOperatingSystemParametersOk

`func (o *OsServerConfigAllOf) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool)`

GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemParameters

`func (o *OsServerConfigAllOf) SetOperatingSystemParameters(v OsOperatingSystemParameters)`

SetOperatingSystemParameters sets OperatingSystemParameters field to given value.

### HasOperatingSystemParameters

`func (o *OsServerConfigAllOf) HasOperatingSystemParameters() bool`

HasOperatingSystemParameters returns a boolean if a field has been set.

### SetOperatingSystemParametersNil

`func (o *OsServerConfigAllOf) SetOperatingSystemParametersNil(b bool)`

 SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil

### UnsetOperatingSystemParameters
`func (o *OsServerConfigAllOf) UnsetOperatingSystemParameters()`

UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil
### GetProcessedInstallTarget

`func (o *OsServerConfigAllOf) GetProcessedInstallTarget() OsInstallTarget`

GetProcessedInstallTarget returns the ProcessedInstallTarget field if non-nil, zero value otherwise.

### GetProcessedInstallTargetOk

`func (o *OsServerConfigAllOf) GetProcessedInstallTargetOk() (*OsInstallTarget, bool)`

GetProcessedInstallTargetOk returns a tuple with the ProcessedInstallTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedInstallTarget

`func (o *OsServerConfigAllOf) SetProcessedInstallTarget(v OsInstallTarget)`

SetProcessedInstallTarget sets ProcessedInstallTarget field to given value.

### HasProcessedInstallTarget

`func (o *OsServerConfigAllOf) HasProcessedInstallTarget() bool`

HasProcessedInstallTarget returns a boolean if a field has been set.

### SetProcessedInstallTargetNil

`func (o *OsServerConfigAllOf) SetProcessedInstallTargetNil(b bool)`

 SetProcessedInstallTargetNil sets the value for ProcessedInstallTarget to be an explicit nil

### UnsetProcessedInstallTarget
`func (o *OsServerConfigAllOf) UnsetProcessedInstallTarget()`

UnsetProcessedInstallTarget ensures that no value is present for ProcessedInstallTarget, not even an explicit nil
### GetSerialNumber

`func (o *OsServerConfigAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *OsServerConfigAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *OsServerConfigAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *OsServerConfigAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetTargetIqn

`func (o *OsServerConfigAllOf) GetTargetIqn() string`

GetTargetIqn returns the TargetIqn field if non-nil, zero value otherwise.

### GetTargetIqnOk

`func (o *OsServerConfigAllOf) GetTargetIqnOk() (*string, bool)`

GetTargetIqnOk returns a tuple with the TargetIqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIqn

`func (o *OsServerConfigAllOf) SetTargetIqn(v string)`

SetTargetIqn sets TargetIqn field to given value.

### HasTargetIqn

`func (o *OsServerConfigAllOf) HasTargetIqn() bool`

HasTargetIqn returns a boolean if a field has been set.

### GetTargetWwpn

`func (o *OsServerConfigAllOf) GetTargetWwpn() string`

GetTargetWwpn returns the TargetWwpn field if non-nil, zero value otherwise.

### GetTargetWwpnOk

`func (o *OsServerConfigAllOf) GetTargetWwpnOk() (*string, bool)`

GetTargetWwpnOk returns a tuple with the TargetWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetWwpn

`func (o *OsServerConfigAllOf) SetTargetWwpn(v string)`

SetTargetWwpn sets TargetWwpn field to given value.

### HasTargetWwpn

`func (o *OsServerConfigAllOf) HasTargetWwpn() bool`

HasTargetWwpn returns a boolean if a field has been set.

### GetVnicMac

`func (o *OsServerConfigAllOf) GetVnicMac() string`

GetVnicMac returns the VnicMac field if non-nil, zero value otherwise.

### GetVnicMacOk

`func (o *OsServerConfigAllOf) GetVnicMacOk() (*string, bool)`

GetVnicMacOk returns a tuple with the VnicMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicMac

`func (o *OsServerConfigAllOf) SetVnicMac(v string)`

SetVnicMac sets VnicMac field to given value.

### HasVnicMac

`func (o *OsServerConfigAllOf) HasVnicMac() bool`

HasVnicMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


