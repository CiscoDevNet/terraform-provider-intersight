# OsBaseInstallConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalParameters** | Pointer to [**[]OsPlaceHolder**](os.PlaceHolder.md) |  | [optional] 
**Answers** | Pointer to [**OsAnswers**](os.Answers.md) |  | [optional] 
**Description** | Pointer to **string** | User provided description about the OS install configuration. | [optional] 
**InstallMethod** | Pointer to **string** | The install method to be used for OS installation - vMedia, iPXE.  Only vMedia is supported as of now. * &#x60;vMedia&#x60; - OS image is mounted as vMedia in target server for OS installation. | [optional] [default to "vMedia"]
**OperatingSystemParameters** | Pointer to [**OsOperatingSystemParameters**](os.OperatingSystemParameters.md) |  | [optional] 

## Methods

### NewOsBaseInstallConfigAllOf

`func NewOsBaseInstallConfigAllOf() *OsBaseInstallConfigAllOf`

NewOsBaseInstallConfigAllOf instantiates a new OsBaseInstallConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsBaseInstallConfigAllOfWithDefaults

`func NewOsBaseInstallConfigAllOfWithDefaults() *OsBaseInstallConfigAllOf`

NewOsBaseInstallConfigAllOfWithDefaults instantiates a new OsBaseInstallConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


