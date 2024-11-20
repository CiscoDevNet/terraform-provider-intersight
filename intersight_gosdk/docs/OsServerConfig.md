# OsServerConfig

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
**OperatingSystemParameters** | Pointer to [**NullableMoBaseComplexType**](MoBaseComplexType.md) | Installation parameters specific to selected OS. | [optional] 
**ProcessedInstallTarget** | Pointer to [**NullableMoBaseComplexType**](MoBaseComplexType.md) | Install Target upon which Operating System is installed. | [optional] 
**SerialNumber** | Pointer to **string** | The Serial Number of the server. | [optional] [readonly] 
**TargetIqn** | Pointer to **string** | IQN (iSCSI qualified name) of Storage iSCSI target. Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique_name. | [optional] [readonly] 
**TargetWwpn** | Pointer to **string** | The WWPN Address of the underlying fibre channel interface at the target used by the storage. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 51:4F:0C:50:14:1F:AF:01. | [optional] [readonly] 
**VnicMac** | Pointer to **string** | MAC address of the VNIC used as iSCSI initiator interface. | [optional] [readonly] 

## Methods

### NewOsServerConfig

`func NewOsServerConfig(classId string, objectType string, ) *OsServerConfig`

NewOsServerConfig instantiates a new OsServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsServerConfigWithDefaults

`func NewOsServerConfigWithDefaults() *OsServerConfig`

NewOsServerConfigWithDefaults instantiates a new OsServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsServerConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsServerConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsServerConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsServerConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsServerConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsServerConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalParameters

`func (o *OsServerConfig) GetAdditionalParameters() []OsPlaceHolder`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *OsServerConfig) GetAdditionalParametersOk() (*[]OsPlaceHolder, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *OsServerConfig) SetAdditionalParameters(v []OsPlaceHolder)`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *OsServerConfig) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.

### SetAdditionalParametersNil

`func (o *OsServerConfig) SetAdditionalParametersNil(b bool)`

 SetAdditionalParametersNil sets the value for AdditionalParameters to be an explicit nil

### UnsetAdditionalParameters
`func (o *OsServerConfig) UnsetAdditionalParameters()`

UnsetAdditionalParameters ensures that no value is present for AdditionalParameters, not even an explicit nil
### GetAnswers

`func (o *OsServerConfig) GetAnswers() OsAnswers`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *OsServerConfig) GetAnswersOk() (*OsAnswers, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *OsServerConfig) SetAnswers(v OsAnswers)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *OsServerConfig) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### SetAnswersNil

`func (o *OsServerConfig) SetAnswersNil(b bool)`

 SetAnswersNil sets the value for Answers to be an explicit nil

### UnsetAnswers
`func (o *OsServerConfig) UnsetAnswers()`

UnsetAnswers ensures that no value is present for Answers, not even an explicit nil
### GetErrorMsgs

`func (o *OsServerConfig) GetErrorMsgs() []string`

GetErrorMsgs returns the ErrorMsgs field if non-nil, zero value otherwise.

### GetErrorMsgsOk

`func (o *OsServerConfig) GetErrorMsgsOk() (*[]string, bool)`

GetErrorMsgsOk returns a tuple with the ErrorMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsgs

`func (o *OsServerConfig) SetErrorMsgs(v []string)`

SetErrorMsgs sets ErrorMsgs field to given value.

### HasErrorMsgs

`func (o *OsServerConfig) HasErrorMsgs() bool`

HasErrorMsgs returns a boolean if a field has been set.

### SetErrorMsgsNil

`func (o *OsServerConfig) SetErrorMsgsNil(b bool)`

 SetErrorMsgsNil sets the value for ErrorMsgs to be an explicit nil

### UnsetErrorMsgs
`func (o *OsServerConfig) UnsetErrorMsgs()`

UnsetErrorMsgs ensures that no value is present for ErrorMsgs, not even an explicit nil
### GetInitiatorWwpn

`func (o *OsServerConfig) GetInitiatorWwpn() string`

GetInitiatorWwpn returns the InitiatorWwpn field if non-nil, zero value otherwise.

### GetInitiatorWwpnOk

`func (o *OsServerConfig) GetInitiatorWwpnOk() (*string, bool)`

GetInitiatorWwpnOk returns a tuple with the InitiatorWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorWwpn

`func (o *OsServerConfig) SetInitiatorWwpn(v string)`

SetInitiatorWwpn sets InitiatorWwpn field to given value.

### HasInitiatorWwpn

`func (o *OsServerConfig) HasInitiatorWwpn() bool`

HasInitiatorWwpn returns a boolean if a field has been set.

### GetInstallTarget

`func (o *OsServerConfig) GetInstallTarget() string`

GetInstallTarget returns the InstallTarget field if non-nil, zero value otherwise.

### GetInstallTargetOk

`func (o *OsServerConfig) GetInstallTargetOk() (*string, bool)`

GetInstallTargetOk returns a tuple with the InstallTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallTarget

`func (o *OsServerConfig) SetInstallTarget(v string)`

SetInstallTarget sets InstallTarget field to given value.

### HasInstallTarget

`func (o *OsServerConfig) HasInstallTarget() bool`

HasInstallTarget returns a boolean if a field has been set.

### GetLunId

`func (o *OsServerConfig) GetLunId() int64`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *OsServerConfig) GetLunIdOk() (*int64, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *OsServerConfig) SetLunId(v int64)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *OsServerConfig) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### GetOperatingSystemParameters

`func (o *OsServerConfig) GetOperatingSystemParameters() MoBaseComplexType`

GetOperatingSystemParameters returns the OperatingSystemParameters field if non-nil, zero value otherwise.

### GetOperatingSystemParametersOk

`func (o *OsServerConfig) GetOperatingSystemParametersOk() (*MoBaseComplexType, bool)`

GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemParameters

`func (o *OsServerConfig) SetOperatingSystemParameters(v MoBaseComplexType)`

SetOperatingSystemParameters sets OperatingSystemParameters field to given value.

### HasOperatingSystemParameters

`func (o *OsServerConfig) HasOperatingSystemParameters() bool`

HasOperatingSystemParameters returns a boolean if a field has been set.

### SetOperatingSystemParametersNil

`func (o *OsServerConfig) SetOperatingSystemParametersNil(b bool)`

 SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil

### UnsetOperatingSystemParameters
`func (o *OsServerConfig) UnsetOperatingSystemParameters()`

UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil
### GetProcessedInstallTarget

`func (o *OsServerConfig) GetProcessedInstallTarget() MoBaseComplexType`

GetProcessedInstallTarget returns the ProcessedInstallTarget field if non-nil, zero value otherwise.

### GetProcessedInstallTargetOk

`func (o *OsServerConfig) GetProcessedInstallTargetOk() (*MoBaseComplexType, bool)`

GetProcessedInstallTargetOk returns a tuple with the ProcessedInstallTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedInstallTarget

`func (o *OsServerConfig) SetProcessedInstallTarget(v MoBaseComplexType)`

SetProcessedInstallTarget sets ProcessedInstallTarget field to given value.

### HasProcessedInstallTarget

`func (o *OsServerConfig) HasProcessedInstallTarget() bool`

HasProcessedInstallTarget returns a boolean if a field has been set.

### SetProcessedInstallTargetNil

`func (o *OsServerConfig) SetProcessedInstallTargetNil(b bool)`

 SetProcessedInstallTargetNil sets the value for ProcessedInstallTarget to be an explicit nil

### UnsetProcessedInstallTarget
`func (o *OsServerConfig) UnsetProcessedInstallTarget()`

UnsetProcessedInstallTarget ensures that no value is present for ProcessedInstallTarget, not even an explicit nil
### GetSerialNumber

`func (o *OsServerConfig) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *OsServerConfig) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *OsServerConfig) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *OsServerConfig) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetTargetIqn

`func (o *OsServerConfig) GetTargetIqn() string`

GetTargetIqn returns the TargetIqn field if non-nil, zero value otherwise.

### GetTargetIqnOk

`func (o *OsServerConfig) GetTargetIqnOk() (*string, bool)`

GetTargetIqnOk returns a tuple with the TargetIqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIqn

`func (o *OsServerConfig) SetTargetIqn(v string)`

SetTargetIqn sets TargetIqn field to given value.

### HasTargetIqn

`func (o *OsServerConfig) HasTargetIqn() bool`

HasTargetIqn returns a boolean if a field has been set.

### GetTargetWwpn

`func (o *OsServerConfig) GetTargetWwpn() string`

GetTargetWwpn returns the TargetWwpn field if non-nil, zero value otherwise.

### GetTargetWwpnOk

`func (o *OsServerConfig) GetTargetWwpnOk() (*string, bool)`

GetTargetWwpnOk returns a tuple with the TargetWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetWwpn

`func (o *OsServerConfig) SetTargetWwpn(v string)`

SetTargetWwpn sets TargetWwpn field to given value.

### HasTargetWwpn

`func (o *OsServerConfig) HasTargetWwpn() bool`

HasTargetWwpn returns a boolean if a field has been set.

### GetVnicMac

`func (o *OsServerConfig) GetVnicMac() string`

GetVnicMac returns the VnicMac field if non-nil, zero value otherwise.

### GetVnicMacOk

`func (o *OsServerConfig) GetVnicMacOk() (*string, bool)`

GetVnicMacOk returns a tuple with the VnicMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicMac

`func (o *OsServerConfig) SetVnicMac(v string)`

SetVnicMac sets VnicMac field to given value.

### HasVnicMac

`func (o *OsServerConfig) HasVnicMac() bool`

HasVnicMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


