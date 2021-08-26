# OsBaseInstallConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "os.Install"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "os.Install"]
**AdditionalParameters** | Pointer to [**[]OsPlaceHolder**](OsPlaceHolder.md) |  | [optional] 
**Answers** | Pointer to [**NullableOsAnswers**](OsAnswers.md) |  | [optional] 
**Description** | Pointer to **string** | User provided description about the OS install configuration. | [optional] 
**ErrorMsg** | Pointer to **string** | The failure message of the API. | [optional] [readonly] 
**InstallMethod** | Pointer to **string** | The install method to be used for OS installation - vMedia, iPXE.  Only vMedia is supported as of now. * &#x60;vMedia&#x60; - OS image is mounted as vMedia in target server for OS installation. | [optional] [default to "vMedia"]
**InstallTarget** | Pointer to [**NullableOsInstallTarget**](OsInstallTarget.md) |  | [optional] 
**OperState** | Pointer to **string** | Denotes API operating status as pending, in_progress, completed_ok, completed_error based on the execution status. * &#x60;Pending&#x60; - The initial value of the OperStatus. * &#x60;InProgress&#x60; - The OperStatus value will be InProgress during execution. * &#x60;CompletedOk&#x60; - The API is successful with operation then OperStatus will be marked as CompletedOk. * &#x60;CompletedError&#x60; - The API is failed with operation then OperStatus will be marked as CompletedError. * &#x60;CompletedWarning&#x60; - The API is completed with some warning then OperStatus will be CompletedWarning. | [optional] [readonly] [default to "Pending"]
**OperatingSystemParameters** | Pointer to [**NullableOsOperatingSystemParameters**](OsOperatingSystemParameters.md) |  | [optional] 

## Methods

### NewOsBaseInstallConfig

`func NewOsBaseInstallConfig(classId string, objectType string, ) *OsBaseInstallConfig`

NewOsBaseInstallConfig instantiates a new OsBaseInstallConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsBaseInstallConfigWithDefaults

`func NewOsBaseInstallConfigWithDefaults() *OsBaseInstallConfig`

NewOsBaseInstallConfigWithDefaults instantiates a new OsBaseInstallConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsBaseInstallConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsBaseInstallConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsBaseInstallConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsBaseInstallConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsBaseInstallConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsBaseInstallConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalParameters

`func (o *OsBaseInstallConfig) GetAdditionalParameters() []OsPlaceHolder`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *OsBaseInstallConfig) GetAdditionalParametersOk() (*[]OsPlaceHolder, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *OsBaseInstallConfig) SetAdditionalParameters(v []OsPlaceHolder)`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *OsBaseInstallConfig) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.

### SetAdditionalParametersNil

`func (o *OsBaseInstallConfig) SetAdditionalParametersNil(b bool)`

 SetAdditionalParametersNil sets the value for AdditionalParameters to be an explicit nil

### UnsetAdditionalParameters
`func (o *OsBaseInstallConfig) UnsetAdditionalParameters()`

UnsetAdditionalParameters ensures that no value is present for AdditionalParameters, not even an explicit nil
### GetAnswers

`func (o *OsBaseInstallConfig) GetAnswers() OsAnswers`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *OsBaseInstallConfig) GetAnswersOk() (*OsAnswers, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *OsBaseInstallConfig) SetAnswers(v OsAnswers)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *OsBaseInstallConfig) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### SetAnswersNil

`func (o *OsBaseInstallConfig) SetAnswersNil(b bool)`

 SetAnswersNil sets the value for Answers to be an explicit nil

### UnsetAnswers
`func (o *OsBaseInstallConfig) UnsetAnswers()`

UnsetAnswers ensures that no value is present for Answers, not even an explicit nil
### GetDescription

`func (o *OsBaseInstallConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OsBaseInstallConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OsBaseInstallConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OsBaseInstallConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrorMsg

`func (o *OsBaseInstallConfig) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *OsBaseInstallConfig) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *OsBaseInstallConfig) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *OsBaseInstallConfig) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetInstallMethod

`func (o *OsBaseInstallConfig) GetInstallMethod() string`

GetInstallMethod returns the InstallMethod field if non-nil, zero value otherwise.

### GetInstallMethodOk

`func (o *OsBaseInstallConfig) GetInstallMethodOk() (*string, bool)`

GetInstallMethodOk returns a tuple with the InstallMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallMethod

`func (o *OsBaseInstallConfig) SetInstallMethod(v string)`

SetInstallMethod sets InstallMethod field to given value.

### HasInstallMethod

`func (o *OsBaseInstallConfig) HasInstallMethod() bool`

HasInstallMethod returns a boolean if a field has been set.

### GetInstallTarget

`func (o *OsBaseInstallConfig) GetInstallTarget() OsInstallTarget`

GetInstallTarget returns the InstallTarget field if non-nil, zero value otherwise.

### GetInstallTargetOk

`func (o *OsBaseInstallConfig) GetInstallTargetOk() (*OsInstallTarget, bool)`

GetInstallTargetOk returns a tuple with the InstallTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallTarget

`func (o *OsBaseInstallConfig) SetInstallTarget(v OsInstallTarget)`

SetInstallTarget sets InstallTarget field to given value.

### HasInstallTarget

`func (o *OsBaseInstallConfig) HasInstallTarget() bool`

HasInstallTarget returns a boolean if a field has been set.

### SetInstallTargetNil

`func (o *OsBaseInstallConfig) SetInstallTargetNil(b bool)`

 SetInstallTargetNil sets the value for InstallTarget to be an explicit nil

### UnsetInstallTarget
`func (o *OsBaseInstallConfig) UnsetInstallTarget()`

UnsetInstallTarget ensures that no value is present for InstallTarget, not even an explicit nil
### GetOperState

`func (o *OsBaseInstallConfig) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *OsBaseInstallConfig) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *OsBaseInstallConfig) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *OsBaseInstallConfig) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperatingSystemParameters

`func (o *OsBaseInstallConfig) GetOperatingSystemParameters() OsOperatingSystemParameters`

GetOperatingSystemParameters returns the OperatingSystemParameters field if non-nil, zero value otherwise.

### GetOperatingSystemParametersOk

`func (o *OsBaseInstallConfig) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool)`

GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemParameters

`func (o *OsBaseInstallConfig) SetOperatingSystemParameters(v OsOperatingSystemParameters)`

SetOperatingSystemParameters sets OperatingSystemParameters field to given value.

### HasOperatingSystemParameters

`func (o *OsBaseInstallConfig) HasOperatingSystemParameters() bool`

HasOperatingSystemParameters returns a boolean if a field has been set.

### SetOperatingSystemParametersNil

`func (o *OsBaseInstallConfig) SetOperatingSystemParametersNil(b bool)`

 SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil

### UnsetOperatingSystemParameters
`func (o *OsBaseInstallConfig) UnsetOperatingSystemParameters()`

UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


