# OsBulkInstallInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.BulkInstallInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.BulkInstallInfo"]
**FileContent** | Pointer to **string** | The content of the entire CSV file is stored as value. The content can hold complete OS install parameters in two sections. The first section holds generic information about the OS Install like OS Image, SCU Image etc. The second section holds parameters which are specific to each server level data. | [optional] 
**GlobalConfig** | Pointer to [**NullableOsGlobalConfig**](OsGlobalConfig.md) |  | [optional] 
**IsFileContentSet** | Pointer to **bool** | Indicates whether the value of the &#39;fileContent&#39; property has been set. | [optional] [readonly] [default to false]
**Name** | Pointer to **string** | The name of the CSV file, which holds the OS install parameters. | [optional] 
**OperState** | Pointer to **string** | Denotes if the operating is pending, in_progress, completed_ok, completed_error. * &#x60;Pending&#x60; - The initial value of the OperStatus. * &#x60;InProgress&#x60; - The OperStatus value will be InProgress during execution. * &#x60;CompletedOk&#x60; - The API is successful with operation then OperStatus will be marked as CompletedOk. * &#x60;CompletedError&#x60; - The API is failed with operation then OperStatus will be marked as CompletedError. * &#x60;CompletedWarning&#x60; - The API is completed with some warning then OperStatus will be CompletedWarning. | [optional] [readonly] [default to "Pending"]
**ServerConfigs** | Pointer to [**[]OsServerConfig**](OsServerConfig.md) |  | [optional] 
**ValidationInfos** | Pointer to [**[]OsValidationInformation**](OsValidationInformation.md) |  | [optional] 
**ConfigurationFile** | Pointer to [**OsConfigurationFileRelationship**](OsConfigurationFileRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**OsImage** | Pointer to [**SoftwarerepositoryOperatingSystemFileRelationship**](SoftwarerepositoryOperatingSystemFileRelationship.md) |  | [optional] 
**ScuImage** | Pointer to [**FirmwareServerConfigurationUtilityDistributableRelationship**](FirmwareServerConfigurationUtilityDistributableRelationship.md) |  | [optional] 
**Servers** | Pointer to [**[]ComputePhysicalRelationship**](ComputePhysicalRelationship.md) | An array of relationships to computePhysical resources. | [optional] 

## Methods

### NewOsBulkInstallInfo

`func NewOsBulkInstallInfo(classId string, objectType string, ) *OsBulkInstallInfo`

NewOsBulkInstallInfo instantiates a new OsBulkInstallInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsBulkInstallInfoWithDefaults

`func NewOsBulkInstallInfoWithDefaults() *OsBulkInstallInfo`

NewOsBulkInstallInfoWithDefaults instantiates a new OsBulkInstallInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsBulkInstallInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsBulkInstallInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsBulkInstallInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsBulkInstallInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsBulkInstallInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsBulkInstallInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileContent

`func (o *OsBulkInstallInfo) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *OsBulkInstallInfo) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *OsBulkInstallInfo) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.

### HasFileContent

`func (o *OsBulkInstallInfo) HasFileContent() bool`

HasFileContent returns a boolean if a field has been set.

### GetGlobalConfig

`func (o *OsBulkInstallInfo) GetGlobalConfig() OsGlobalConfig`

GetGlobalConfig returns the GlobalConfig field if non-nil, zero value otherwise.

### GetGlobalConfigOk

`func (o *OsBulkInstallInfo) GetGlobalConfigOk() (*OsGlobalConfig, bool)`

GetGlobalConfigOk returns a tuple with the GlobalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalConfig

`func (o *OsBulkInstallInfo) SetGlobalConfig(v OsGlobalConfig)`

SetGlobalConfig sets GlobalConfig field to given value.

### HasGlobalConfig

`func (o *OsBulkInstallInfo) HasGlobalConfig() bool`

HasGlobalConfig returns a boolean if a field has been set.

### SetGlobalConfigNil

`func (o *OsBulkInstallInfo) SetGlobalConfigNil(b bool)`

 SetGlobalConfigNil sets the value for GlobalConfig to be an explicit nil

### UnsetGlobalConfig
`func (o *OsBulkInstallInfo) UnsetGlobalConfig()`

UnsetGlobalConfig ensures that no value is present for GlobalConfig, not even an explicit nil
### GetIsFileContentSet

`func (o *OsBulkInstallInfo) GetIsFileContentSet() bool`

GetIsFileContentSet returns the IsFileContentSet field if non-nil, zero value otherwise.

### GetIsFileContentSetOk

`func (o *OsBulkInstallInfo) GetIsFileContentSetOk() (*bool, bool)`

GetIsFileContentSetOk returns a tuple with the IsFileContentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFileContentSet

`func (o *OsBulkInstallInfo) SetIsFileContentSet(v bool)`

SetIsFileContentSet sets IsFileContentSet field to given value.

### HasIsFileContentSet

`func (o *OsBulkInstallInfo) HasIsFileContentSet() bool`

HasIsFileContentSet returns a boolean if a field has been set.

### GetName

`func (o *OsBulkInstallInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsBulkInstallInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsBulkInstallInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsBulkInstallInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *OsBulkInstallInfo) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *OsBulkInstallInfo) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *OsBulkInstallInfo) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *OsBulkInstallInfo) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetServerConfigs

`func (o *OsBulkInstallInfo) GetServerConfigs() []OsServerConfig`

GetServerConfigs returns the ServerConfigs field if non-nil, zero value otherwise.

### GetServerConfigsOk

`func (o *OsBulkInstallInfo) GetServerConfigsOk() (*[]OsServerConfig, bool)`

GetServerConfigsOk returns a tuple with the ServerConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfigs

`func (o *OsBulkInstallInfo) SetServerConfigs(v []OsServerConfig)`

SetServerConfigs sets ServerConfigs field to given value.

### HasServerConfigs

`func (o *OsBulkInstallInfo) HasServerConfigs() bool`

HasServerConfigs returns a boolean if a field has been set.

### SetServerConfigsNil

`func (o *OsBulkInstallInfo) SetServerConfigsNil(b bool)`

 SetServerConfigsNil sets the value for ServerConfigs to be an explicit nil

### UnsetServerConfigs
`func (o *OsBulkInstallInfo) UnsetServerConfigs()`

UnsetServerConfigs ensures that no value is present for ServerConfigs, not even an explicit nil
### GetValidationInfos

`func (o *OsBulkInstallInfo) GetValidationInfos() []OsValidationInformation`

GetValidationInfos returns the ValidationInfos field if non-nil, zero value otherwise.

### GetValidationInfosOk

`func (o *OsBulkInstallInfo) GetValidationInfosOk() (*[]OsValidationInformation, bool)`

GetValidationInfosOk returns a tuple with the ValidationInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInfos

`func (o *OsBulkInstallInfo) SetValidationInfos(v []OsValidationInformation)`

SetValidationInfos sets ValidationInfos field to given value.

### HasValidationInfos

`func (o *OsBulkInstallInfo) HasValidationInfos() bool`

HasValidationInfos returns a boolean if a field has been set.

### SetValidationInfosNil

`func (o *OsBulkInstallInfo) SetValidationInfosNil(b bool)`

 SetValidationInfosNil sets the value for ValidationInfos to be an explicit nil

### UnsetValidationInfos
`func (o *OsBulkInstallInfo) UnsetValidationInfos()`

UnsetValidationInfos ensures that no value is present for ValidationInfos, not even an explicit nil
### GetConfigurationFile

`func (o *OsBulkInstallInfo) GetConfigurationFile() OsConfigurationFileRelationship`

GetConfigurationFile returns the ConfigurationFile field if non-nil, zero value otherwise.

### GetConfigurationFileOk

`func (o *OsBulkInstallInfo) GetConfigurationFileOk() (*OsConfigurationFileRelationship, bool)`

GetConfigurationFileOk returns a tuple with the ConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFile

`func (o *OsBulkInstallInfo) SetConfigurationFile(v OsConfigurationFileRelationship)`

SetConfigurationFile sets ConfigurationFile field to given value.

### HasConfigurationFile

`func (o *OsBulkInstallInfo) HasConfigurationFile() bool`

HasConfigurationFile returns a boolean if a field has been set.

### GetOrganization

`func (o *OsBulkInstallInfo) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OsBulkInstallInfo) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OsBulkInstallInfo) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OsBulkInstallInfo) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOsImage

`func (o *OsBulkInstallInfo) GetOsImage() SoftwarerepositoryOperatingSystemFileRelationship`

GetOsImage returns the OsImage field if non-nil, zero value otherwise.

### GetOsImageOk

`func (o *OsBulkInstallInfo) GetOsImageOk() (*SoftwarerepositoryOperatingSystemFileRelationship, bool)`

GetOsImageOk returns a tuple with the OsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImage

`func (o *OsBulkInstallInfo) SetOsImage(v SoftwarerepositoryOperatingSystemFileRelationship)`

SetOsImage sets OsImage field to given value.

### HasOsImage

`func (o *OsBulkInstallInfo) HasOsImage() bool`

HasOsImage returns a boolean if a field has been set.

### GetScuImage

`func (o *OsBulkInstallInfo) GetScuImage() FirmwareServerConfigurationUtilityDistributableRelationship`

GetScuImage returns the ScuImage field if non-nil, zero value otherwise.

### GetScuImageOk

`func (o *OsBulkInstallInfo) GetScuImageOk() (*FirmwareServerConfigurationUtilityDistributableRelationship, bool)`

GetScuImageOk returns a tuple with the ScuImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScuImage

`func (o *OsBulkInstallInfo) SetScuImage(v FirmwareServerConfigurationUtilityDistributableRelationship)`

SetScuImage sets ScuImage field to given value.

### HasScuImage

`func (o *OsBulkInstallInfo) HasScuImage() bool`

HasScuImage returns a boolean if a field has been set.

### GetServers

`func (o *OsBulkInstallInfo) GetServers() []ComputePhysicalRelationship`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OsBulkInstallInfo) GetServersOk() (*[]ComputePhysicalRelationship, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OsBulkInstallInfo) SetServers(v []ComputePhysicalRelationship)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OsBulkInstallInfo) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *OsBulkInstallInfo) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *OsBulkInstallInfo) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


