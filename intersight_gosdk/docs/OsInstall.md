# OsInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.Install"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.Install"]
**Name** | Pointer to **string** | The name of the OS install configuration. | [optional] 
**ConfigurationFile** | Pointer to [**NullableOsConfigurationFileRelationship**](OsConfigurationFileRelationship.md) |  | [optional] 
**Image** | Pointer to [**NullableSoftwarerepositoryOperatingSystemFileRelationship**](SoftwarerepositoryOperatingSystemFileRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**OsduImage** | Pointer to [**NullableFirmwareServerConfigurationUtilityDistributableRelationship**](FirmwareServerConfigurationUtilityDistributableRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewOsInstall

`func NewOsInstall(classId string, objectType string, ) *OsInstall`

NewOsInstall instantiates a new OsInstall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsInstallWithDefaults

`func NewOsInstallWithDefaults() *OsInstall`

NewOsInstallWithDefaults instantiates a new OsInstall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsInstall) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsInstall) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsInstall) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsInstall) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsInstall) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsInstall) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *OsInstall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsInstall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsInstall) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsInstall) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigurationFile

`func (o *OsInstall) GetConfigurationFile() OsConfigurationFileRelationship`

GetConfigurationFile returns the ConfigurationFile field if non-nil, zero value otherwise.

### GetConfigurationFileOk

`func (o *OsInstall) GetConfigurationFileOk() (*OsConfigurationFileRelationship, bool)`

GetConfigurationFileOk returns a tuple with the ConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFile

`func (o *OsInstall) SetConfigurationFile(v OsConfigurationFileRelationship)`

SetConfigurationFile sets ConfigurationFile field to given value.

### HasConfigurationFile

`func (o *OsInstall) HasConfigurationFile() bool`

HasConfigurationFile returns a boolean if a field has been set.

### SetConfigurationFileNil

`func (o *OsInstall) SetConfigurationFileNil(b bool)`

 SetConfigurationFileNil sets the value for ConfigurationFile to be an explicit nil

### UnsetConfigurationFile
`func (o *OsInstall) UnsetConfigurationFile()`

UnsetConfigurationFile ensures that no value is present for ConfigurationFile, not even an explicit nil
### GetImage

`func (o *OsInstall) GetImage() SoftwarerepositoryOperatingSystemFileRelationship`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OsInstall) GetImageOk() (*SoftwarerepositoryOperatingSystemFileRelationship, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OsInstall) SetImage(v SoftwarerepositoryOperatingSystemFileRelationship)`

SetImage sets Image field to given value.

### HasImage

`func (o *OsInstall) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *OsInstall) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *OsInstall) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetOrganization

`func (o *OsInstall) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OsInstall) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OsInstall) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OsInstall) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *OsInstall) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *OsInstall) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetOsduImage

`func (o *OsInstall) GetOsduImage() FirmwareServerConfigurationUtilityDistributableRelationship`

GetOsduImage returns the OsduImage field if non-nil, zero value otherwise.

### GetOsduImageOk

`func (o *OsInstall) GetOsduImageOk() (*FirmwareServerConfigurationUtilityDistributableRelationship, bool)`

GetOsduImageOk returns a tuple with the OsduImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsduImage

`func (o *OsInstall) SetOsduImage(v FirmwareServerConfigurationUtilityDistributableRelationship)`

SetOsduImage sets OsduImage field to given value.

### HasOsduImage

`func (o *OsInstall) HasOsduImage() bool`

HasOsduImage returns a boolean if a field has been set.

### SetOsduImageNil

`func (o *OsInstall) SetOsduImageNil(b bool)`

 SetOsduImageNil sets the value for OsduImage to be an explicit nil

### UnsetOsduImage
`func (o *OsInstall) UnsetOsduImage()`

UnsetOsduImage ensures that no value is present for OsduImage, not even an explicit nil
### GetServer

`func (o *OsInstall) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OsInstall) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OsInstall) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *OsInstall) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *OsInstall) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *OsInstall) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetWorkflowInfo

`func (o *OsInstall) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *OsInstall) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *OsInstall) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *OsInstall) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.

### SetWorkflowInfoNil

`func (o *OsInstall) SetWorkflowInfoNil(b bool)`

 SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil

### UnsetWorkflowInfo
`func (o *OsInstall) UnsetWorkflowInfo()`

UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


