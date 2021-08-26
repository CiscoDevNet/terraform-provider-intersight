# OsGlobalConfig

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

### NewOsGlobalConfig

`func NewOsGlobalConfig(classId string, objectType string, ) *OsGlobalConfig`

NewOsGlobalConfig instantiates a new OsGlobalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsGlobalConfigWithDefaults

`func NewOsGlobalConfigWithDefaults() *OsGlobalConfig`

NewOsGlobalConfigWithDefaults instantiates a new OsGlobalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsGlobalConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsGlobalConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsGlobalConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsGlobalConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsGlobalConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsGlobalConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigurationFileName

`func (o *OsGlobalConfig) GetConfigurationFileName() string`

GetConfigurationFileName returns the ConfigurationFileName field if non-nil, zero value otherwise.

### GetConfigurationFileNameOk

`func (o *OsGlobalConfig) GetConfigurationFileNameOk() (*string, bool)`

GetConfigurationFileNameOk returns a tuple with the ConfigurationFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFileName

`func (o *OsGlobalConfig) SetConfigurationFileName(v string)`

SetConfigurationFileName sets ConfigurationFileName field to given value.

### HasConfigurationFileName

`func (o *OsGlobalConfig) HasConfigurationFileName() bool`

HasConfigurationFileName returns a boolean if a field has been set.

### GetConfigurationSource

`func (o *OsGlobalConfig) GetConfigurationSource() string`

GetConfigurationSource returns the ConfigurationSource field if non-nil, zero value otherwise.

### GetConfigurationSourceOk

`func (o *OsGlobalConfig) GetConfigurationSourceOk() (*string, bool)`

GetConfigurationSourceOk returns a tuple with the ConfigurationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSource

`func (o *OsGlobalConfig) SetConfigurationSource(v string)`

SetConfigurationSource sets ConfigurationSource field to given value.

### HasConfigurationSource

`func (o *OsGlobalConfig) HasConfigurationSource() bool`

HasConfigurationSource returns a boolean if a field has been set.

### GetInstallMethod

`func (o *OsGlobalConfig) GetInstallMethod() string`

GetInstallMethod returns the InstallMethod field if non-nil, zero value otherwise.

### GetInstallMethodOk

`func (o *OsGlobalConfig) GetInstallMethodOk() (*string, bool)`

GetInstallMethodOk returns a tuple with the InstallMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallMethod

`func (o *OsGlobalConfig) SetInstallMethod(v string)`

SetInstallMethod sets InstallMethod field to given value.

### HasInstallMethod

`func (o *OsGlobalConfig) HasInstallMethod() bool`

HasInstallMethod returns a boolean if a field has been set.

### GetInstallTargetType

`func (o *OsGlobalConfig) GetInstallTargetType() string`

GetInstallTargetType returns the InstallTargetType field if non-nil, zero value otherwise.

### GetInstallTargetTypeOk

`func (o *OsGlobalConfig) GetInstallTargetTypeOk() (*string, bool)`

GetInstallTargetTypeOk returns a tuple with the InstallTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallTargetType

`func (o *OsGlobalConfig) SetInstallTargetType(v string)`

SetInstallTargetType sets InstallTargetType field to given value.

### HasInstallTargetType

`func (o *OsGlobalConfig) HasInstallTargetType() bool`

HasInstallTargetType returns a boolean if a field has been set.

### GetOperatingSystemParameters

`func (o *OsGlobalConfig) GetOperatingSystemParameters() OsOperatingSystemParameters`

GetOperatingSystemParameters returns the OperatingSystemParameters field if non-nil, zero value otherwise.

### GetOperatingSystemParametersOk

`func (o *OsGlobalConfig) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool)`

GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemParameters

`func (o *OsGlobalConfig) SetOperatingSystemParameters(v OsOperatingSystemParameters)`

SetOperatingSystemParameters sets OperatingSystemParameters field to given value.

### HasOperatingSystemParameters

`func (o *OsGlobalConfig) HasOperatingSystemParameters() bool`

HasOperatingSystemParameters returns a boolean if a field has been set.

### SetOperatingSystemParametersNil

`func (o *OsGlobalConfig) SetOperatingSystemParametersNil(b bool)`

 SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil

### UnsetOperatingSystemParameters
`func (o *OsGlobalConfig) UnsetOperatingSystemParameters()`

UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil
### GetOsImageName

`func (o *OsGlobalConfig) GetOsImageName() string`

GetOsImageName returns the OsImageName field if non-nil, zero value otherwise.

### GetOsImageNameOk

`func (o *OsGlobalConfig) GetOsImageNameOk() (*string, bool)`

GetOsImageNameOk returns a tuple with the OsImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImageName

`func (o *OsGlobalConfig) SetOsImageName(v string)`

SetOsImageName sets OsImageName field to given value.

### HasOsImageName

`func (o *OsGlobalConfig) HasOsImageName() bool`

HasOsImageName returns a boolean if a field has been set.

### GetScuImageName

`func (o *OsGlobalConfig) GetScuImageName() string`

GetScuImageName returns the ScuImageName field if non-nil, zero value otherwise.

### GetScuImageNameOk

`func (o *OsGlobalConfig) GetScuImageNameOk() (*string, bool)`

GetScuImageNameOk returns a tuple with the ScuImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScuImageName

`func (o *OsGlobalConfig) SetScuImageName(v string)`

SetScuImageName sets ScuImageName field to given value.

### HasScuImageName

`func (o *OsGlobalConfig) HasScuImageName() bool`

HasScuImageName returns a boolean if a field has been set.

### GetWindowsEdition

`func (o *OsGlobalConfig) GetWindowsEdition() string`

GetWindowsEdition returns the WindowsEdition field if non-nil, zero value otherwise.

### GetWindowsEditionOk

`func (o *OsGlobalConfig) GetWindowsEditionOk() (*string, bool)`

GetWindowsEditionOk returns a tuple with the WindowsEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsEdition

`func (o *OsGlobalConfig) SetWindowsEdition(v string)`

SetWindowsEdition sets WindowsEdition field to given value.

### HasWindowsEdition

`func (o *OsGlobalConfig) HasWindowsEdition() bool`

HasWindowsEdition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


