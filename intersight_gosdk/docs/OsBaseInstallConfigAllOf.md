# OsBaseInstallConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "os.Install"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "os.Install"]
**AdditionalParameters** | Pointer to [**[]OsPlaceHolder**](os.PlaceHolder.md) |  | [optional] 
**Answers** | Pointer to [**NullableOsAnswers**](os.Answers.md) |  | [optional] 
**Description** | Pointer to **string** | User provided description about the OS install configuration. | [optional] 
**InstallMethod** | Pointer to **string** | The install method to be used for OS installation - vMedia, iPXE.  Only vMedia is supported as of now. * &#x60;vMedia&#x60; - OS image is mounted as vMedia in target server for OS installation. | [optional] [default to "vMedia"]
**InstallTarget** | Pointer to [**NullableOsInstallTarget**](os.InstallTarget.md) |  | [optional] 
**OperatingSystemParameters** | Pointer to [**NullableOsOperatingSystemParameters**](os.OperatingSystemParameters.md) |  | [optional] 

## Methods

### NewOsBaseInstallConfigAllOf

`func NewOsBaseInstallConfigAllOf(classId string, objectType string, ) *OsBaseInstallConfigAllOf`

NewOsBaseInstallConfigAllOf instantiates a new OsBaseInstallConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsBaseInstallConfigAllOfWithDefaults

`func NewOsBaseInstallConfigAllOfWithDefaults() *OsBaseInstallConfigAllOf`

NewOsBaseInstallConfigAllOfWithDefaults instantiates a new OsBaseInstallConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsBaseInstallConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsBaseInstallConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsBaseInstallConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsBaseInstallConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsBaseInstallConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsBaseInstallConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalParameters

`func (o *OsBaseInstallConfigAllOf) GetAdditionalParameters() []OsPlaceHolder`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *OsBaseInstallConfigAllOf) GetAdditionalParametersOk() (*[]OsPlaceHolder, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *OsBaseInstallConfigAllOf) SetAdditionalParameters(v []OsPlaceHolder)`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *OsBaseInstallConfigAllOf) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.

### SetAdditionalParametersNil

`func (o *OsBaseInstallConfigAllOf) SetAdditionalParametersNil(b bool)`

 SetAdditionalParametersNil sets the value for AdditionalParameters to be an explicit nil

### UnsetAdditionalParameters
`func (o *OsBaseInstallConfigAllOf) UnsetAdditionalParameters()`

UnsetAdditionalParameters ensures that no value is present for AdditionalParameters, not even an explicit nil
### GetAnswers

`func (o *OsBaseInstallConfigAllOf) GetAnswers() OsAnswers`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *OsBaseInstallConfigAllOf) GetAnswersOk() (*OsAnswers, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *OsBaseInstallConfigAllOf) SetAnswers(v OsAnswers)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *OsBaseInstallConfigAllOf) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### SetAnswersNil

`func (o *OsBaseInstallConfigAllOf) SetAnswersNil(b bool)`

 SetAnswersNil sets the value for Answers to be an explicit nil

### UnsetAnswers
`func (o *OsBaseInstallConfigAllOf) UnsetAnswers()`

UnsetAnswers ensures that no value is present for Answers, not even an explicit nil
### GetDescription

`func (o *OsBaseInstallConfigAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OsBaseInstallConfigAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OsBaseInstallConfigAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OsBaseInstallConfigAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstallMethod

`func (o *OsBaseInstallConfigAllOf) GetInstallMethod() string`

GetInstallMethod returns the InstallMethod field if non-nil, zero value otherwise.

### GetInstallMethodOk

`func (o *OsBaseInstallConfigAllOf) GetInstallMethodOk() (*string, bool)`

GetInstallMethodOk returns a tuple with the InstallMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallMethod

`func (o *OsBaseInstallConfigAllOf) SetInstallMethod(v string)`

SetInstallMethod sets InstallMethod field to given value.

### HasInstallMethod

`func (o *OsBaseInstallConfigAllOf) HasInstallMethod() bool`

HasInstallMethod returns a boolean if a field has been set.

### GetInstallTarget

`func (o *OsBaseInstallConfigAllOf) GetInstallTarget() OsInstallTarget`

GetInstallTarget returns the InstallTarget field if non-nil, zero value otherwise.

### GetInstallTargetOk

`func (o *OsBaseInstallConfigAllOf) GetInstallTargetOk() (*OsInstallTarget, bool)`

GetInstallTargetOk returns a tuple with the InstallTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallTarget

`func (o *OsBaseInstallConfigAllOf) SetInstallTarget(v OsInstallTarget)`

SetInstallTarget sets InstallTarget field to given value.

### HasInstallTarget

`func (o *OsBaseInstallConfigAllOf) HasInstallTarget() bool`

HasInstallTarget returns a boolean if a field has been set.

### SetInstallTargetNil

`func (o *OsBaseInstallConfigAllOf) SetInstallTargetNil(b bool)`

 SetInstallTargetNil sets the value for InstallTarget to be an explicit nil

### UnsetInstallTarget
`func (o *OsBaseInstallConfigAllOf) UnsetInstallTarget()`

UnsetInstallTarget ensures that no value is present for InstallTarget, not even an explicit nil
### GetOperatingSystemParameters

`func (o *OsBaseInstallConfigAllOf) GetOperatingSystemParameters() OsOperatingSystemParameters`

GetOperatingSystemParameters returns the OperatingSystemParameters field if non-nil, zero value otherwise.

### GetOperatingSystemParametersOk

`func (o *OsBaseInstallConfigAllOf) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool)`

GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemParameters

`func (o *OsBaseInstallConfigAllOf) SetOperatingSystemParameters(v OsOperatingSystemParameters)`

SetOperatingSystemParameters sets OperatingSystemParameters field to given value.

### HasOperatingSystemParameters

`func (o *OsBaseInstallConfigAllOf) HasOperatingSystemParameters() bool`

HasOperatingSystemParameters returns a boolean if a field has been set.

### SetOperatingSystemParametersNil

`func (o *OsBaseInstallConfigAllOf) SetOperatingSystemParametersNil(b bool)`

 SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil

### UnsetOperatingSystemParameters
`func (o *OsBaseInstallConfigAllOf) UnsetOperatingSystemParameters()`

UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


