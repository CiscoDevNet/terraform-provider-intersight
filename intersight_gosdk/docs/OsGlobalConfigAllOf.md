# OsGlobalConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.GlobalConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.GlobalConfig"]
**ConfigurationFileName** | Pointer to **string** | Name of the Configuration file. | [optional] [readonly] 
**ConfigurationSource** | Pointer to **string** | Configuration source for the OS Installation. | [optional] [readonly] 
**InstallMethod** | Pointer to **string** | The install method to be used for OS installation - vMedia, iPXE. Only vMedia is supported as of now. | [optional] [readonly] 
**InstallTargetType** | Pointer to **string** | The Prefill install Target Name. | [optional] [readonly] 
**OperatingSystemParameters** | Pointer to [**NullableOsOperatingSystemParameters**](OsOperatingSystemParameters.md) |  | [optional] 
**OsImageName** | Pointer to **string** | The Operating System Image name. | [optional] [readonly] 
**ScuImageName** | Pointer to **string** | The name of the Server Configuration Utilities Image. | [optional] [readonly] 
**WindowsEdition** | Pointer to **string** | The Windows OS edition, this property required only for Windows server. | [optional] [readonly] 

## Methods

### NewOsGlobalConfigAllOf

`func NewOsGlobalConfigAllOf(classId string, objectType string, ) *OsGlobalConfigAllOf`

NewOsGlobalConfigAllOf instantiates a new OsGlobalConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsGlobalConfigAllOfWithDefaults

`func NewOsGlobalConfigAllOfWithDefaults() *OsGlobalConfigAllOf`

NewOsGlobalConfigAllOfWithDefaults instantiates a new OsGlobalConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsGlobalConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsGlobalConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsGlobalConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsGlobalConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsGlobalConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsGlobalConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigurationFileName

`func (o *OsGlobalConfigAllOf) GetConfigurationFileName() string`

GetConfigurationFileName returns the ConfigurationFileName field if non-nil, zero value otherwise.

### GetConfigurationFileNameOk

`func (o *OsGlobalConfigAllOf) GetConfigurationFileNameOk() (*string, bool)`

GetConfigurationFileNameOk returns a tuple with the ConfigurationFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFileName

`func (o *OsGlobalConfigAllOf) SetConfigurationFileName(v string)`

SetConfigurationFileName sets ConfigurationFileName field to given value.

### HasConfigurationFileName

`func (o *OsGlobalConfigAllOf) HasConfigurationFileName() bool`

HasConfigurationFileName returns a boolean if a field has been set.

### GetConfigurationSource

`func (o *OsGlobalConfigAllOf) GetConfigurationSource() string`

GetConfigurationSource returns the ConfigurationSource field if non-nil, zero value otherwise.

### GetConfigurationSourceOk

`func (o *OsGlobalConfigAllOf) GetConfigurationSourceOk() (*string, bool)`

GetConfigurationSourceOk returns a tuple with the ConfigurationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSource

`func (o *OsGlobalConfigAllOf) SetConfigurationSource(v string)`

SetConfigurationSource sets ConfigurationSource field to given value.

### HasConfigurationSource

`func (o *OsGlobalConfigAllOf) HasConfigurationSource() bool`

HasConfigurationSource returns a boolean if a field has been set.

### GetInstallMethod

`func (o *OsGlobalConfigAllOf) GetInstallMethod() string`

GetInstallMethod returns the InstallMethod field if non-nil, zero value otherwise.

### GetInstallMethodOk

`func (o *OsGlobalConfigAllOf) GetInstallMethodOk() (*string, bool)`

GetInstallMethodOk returns a tuple with the InstallMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallMethod

`func (o *OsGlobalConfigAllOf) SetInstallMethod(v string)`

SetInstallMethod sets InstallMethod field to given value.

### HasInstallMethod

`func (o *OsGlobalConfigAllOf) HasInstallMethod() bool`

HasInstallMethod returns a boolean if a field has been set.

### GetInstallTargetType

`func (o *OsGlobalConfigAllOf) GetInstallTargetType() string`

GetInstallTargetType returns the InstallTargetType field if non-nil, zero value otherwise.

### GetInstallTargetTypeOk

`func (o *OsGlobalConfigAllOf) GetInstallTargetTypeOk() (*string, bool)`

GetInstallTargetTypeOk returns a tuple with the InstallTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallTargetType

`func (o *OsGlobalConfigAllOf) SetInstallTargetType(v string)`

SetInstallTargetType sets InstallTargetType field to given value.

### HasInstallTargetType

`func (o *OsGlobalConfigAllOf) HasInstallTargetType() bool`

HasInstallTargetType returns a boolean if a field has been set.

### GetOperatingSystemParameters

`func (o *OsGlobalConfigAllOf) GetOperatingSystemParameters() OsOperatingSystemParameters`

GetOperatingSystemParameters returns the OperatingSystemParameters field if non-nil, zero value otherwise.

### GetOperatingSystemParametersOk

`func (o *OsGlobalConfigAllOf) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool)`

GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemParameters

`func (o *OsGlobalConfigAllOf) SetOperatingSystemParameters(v OsOperatingSystemParameters)`

SetOperatingSystemParameters sets OperatingSystemParameters field to given value.

### HasOperatingSystemParameters

`func (o *OsGlobalConfigAllOf) HasOperatingSystemParameters() bool`

HasOperatingSystemParameters returns a boolean if a field has been set.

### SetOperatingSystemParametersNil

`func (o *OsGlobalConfigAllOf) SetOperatingSystemParametersNil(b bool)`

 SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil

### UnsetOperatingSystemParameters
`func (o *OsGlobalConfigAllOf) UnsetOperatingSystemParameters()`

UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil
### GetOsImageName

`func (o *OsGlobalConfigAllOf) GetOsImageName() string`

GetOsImageName returns the OsImageName field if non-nil, zero value otherwise.

### GetOsImageNameOk

`func (o *OsGlobalConfigAllOf) GetOsImageNameOk() (*string, bool)`

GetOsImageNameOk returns a tuple with the OsImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImageName

`func (o *OsGlobalConfigAllOf) SetOsImageName(v string)`

SetOsImageName sets OsImageName field to given value.

### HasOsImageName

`func (o *OsGlobalConfigAllOf) HasOsImageName() bool`

HasOsImageName returns a boolean if a field has been set.

### GetScuImageName

`func (o *OsGlobalConfigAllOf) GetScuImageName() string`

GetScuImageName returns the ScuImageName field if non-nil, zero value otherwise.

### GetScuImageNameOk

`func (o *OsGlobalConfigAllOf) GetScuImageNameOk() (*string, bool)`

GetScuImageNameOk returns a tuple with the ScuImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScuImageName

`func (o *OsGlobalConfigAllOf) SetScuImageName(v string)`

SetScuImageName sets ScuImageName field to given value.

### HasScuImageName

`func (o *OsGlobalConfigAllOf) HasScuImageName() bool`

HasScuImageName returns a boolean if a field has been set.

### GetWindowsEdition

`func (o *OsGlobalConfigAllOf) GetWindowsEdition() string`

GetWindowsEdition returns the WindowsEdition field if non-nil, zero value otherwise.

### GetWindowsEditionOk

`func (o *OsGlobalConfigAllOf) GetWindowsEditionOk() (*string, bool)`

GetWindowsEditionOk returns a tuple with the WindowsEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsEdition

`func (o *OsGlobalConfigAllOf) SetWindowsEdition(v string)`

SetWindowsEdition sets WindowsEdition field to given value.

### HasWindowsEdition

`func (o *OsGlobalConfigAllOf) HasWindowsEdition() bool`

HasWindowsEdition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


