# OsServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.ServerConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.ServerConfig"]
**AdditionalParameters** | Pointer to [**[]OsPlaceHolder**](OsPlaceHolder.md) |  | [optional] 
**Answers** | Pointer to [**NullableOsAnswers**](OsAnswers.md) |  | [optional] 
**ErrorMsgs** | Pointer to **[]string** |  | [optional] 
**InstallTarget** | Pointer to **string** | The target in which OS installation triggered, the value represented is StorageControllerID follwed by PhysicalDisk SerialNumber in case of Physcial disk or VirtualDriveId for virtual drive. | [optional] [readonly] 
**ProcessedInstallTarget** | Pointer to [**NullableOsInstallTarget**](OsInstallTarget.md) |  | [optional] 
**SerialNumber** | Pointer to **string** | The Serial Number of the server. | [optional] [readonly] 

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

### GetProcessedInstallTarget

`func (o *OsServerConfig) GetProcessedInstallTarget() OsInstallTarget`

GetProcessedInstallTarget returns the ProcessedInstallTarget field if non-nil, zero value otherwise.

### GetProcessedInstallTargetOk

`func (o *OsServerConfig) GetProcessedInstallTargetOk() (*OsInstallTarget, bool)`

GetProcessedInstallTargetOk returns a tuple with the ProcessedInstallTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedInstallTarget

`func (o *OsServerConfig) SetProcessedInstallTarget(v OsInstallTarget)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


