# OsInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.Install"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.Install"]
**Name** | Pointer to **string** | The name of the OS install configuration. | [optional] 
**ConfigurationFile** | Pointer to [**OsConfigurationFileRelationship**](OsConfigurationFileRelationship.md) |  | [optional] 
**Image** | Pointer to [**SoftwarerepositoryOperatingSystemFileRelationship**](SoftwarerepositoryOperatingSystemFileRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**OsduImage** | Pointer to [**FirmwareServerConfigurationUtilityDistributableRelationship**](FirmwareServerConfigurationUtilityDistributableRelationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


