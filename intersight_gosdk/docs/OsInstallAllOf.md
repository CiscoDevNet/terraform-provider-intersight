# OsInstallAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the OS install configuration. | [optional] 
**ConfigurationFile** | Pointer to [**OsConfigurationFileRelationship**](os.ConfigurationFile.Relationship.md) |  | [optional] 
**Image** | Pointer to [**SoftwarerepositoryOperatingSystemFileRelationship**](softwarerepository.OperatingSystemFile.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**OsduImage** | Pointer to [**FirmwareServerConfigurationUtilityDistributableRelationship**](firmware.ServerConfigurationUtilityDistributable.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewOsInstallAllOf

`func NewOsInstallAllOf() *OsInstallAllOf`

NewOsInstallAllOf instantiates a new OsInstallAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsInstallAllOfWithDefaults

`func NewOsInstallAllOfWithDefaults() *OsInstallAllOf`

NewOsInstallAllOfWithDefaults instantiates a new OsInstallAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OsInstallAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsInstallAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsInstallAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsInstallAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigurationFile

`func (o *OsInstallAllOf) GetConfigurationFile() OsConfigurationFileRelationship`

GetConfigurationFile returns the ConfigurationFile field if non-nil, zero value otherwise.

### GetConfigurationFileOk

`func (o *OsInstallAllOf) GetConfigurationFileOk() (*OsConfigurationFileRelationship, bool)`

GetConfigurationFileOk returns a tuple with the ConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFile

`func (o *OsInstallAllOf) SetConfigurationFile(v OsConfigurationFileRelationship)`

SetConfigurationFile sets ConfigurationFile field to given value.

### HasConfigurationFile

`func (o *OsInstallAllOf) HasConfigurationFile() bool`

HasConfigurationFile returns a boolean if a field has been set.

### GetImage

`func (o *OsInstallAllOf) GetImage() SoftwarerepositoryOperatingSystemFileRelationship`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OsInstallAllOf) GetImageOk() (*SoftwarerepositoryOperatingSystemFileRelationship, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OsInstallAllOf) SetImage(v SoftwarerepositoryOperatingSystemFileRelationship)`

SetImage sets Image field to given value.

### HasImage

`func (o *OsInstallAllOf) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetOrganization

`func (o *OsInstallAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OsInstallAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OsInstallAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OsInstallAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOsduImage

`func (o *OsInstallAllOf) GetOsduImage() FirmwareServerConfigurationUtilityDistributableRelationship`

GetOsduImage returns the OsduImage field if non-nil, zero value otherwise.

### GetOsduImageOk

`func (o *OsInstallAllOf) GetOsduImageOk() (*FirmwareServerConfigurationUtilityDistributableRelationship, bool)`

GetOsduImageOk returns a tuple with the OsduImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsduImage

`func (o *OsInstallAllOf) SetOsduImage(v FirmwareServerConfigurationUtilityDistributableRelationship)`

SetOsduImage sets OsduImage field to given value.

### HasOsduImage

`func (o *OsInstallAllOf) HasOsduImage() bool`

HasOsduImage returns a boolean if a field has been set.

### GetServer

`func (o *OsInstallAllOf) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OsInstallAllOf) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OsInstallAllOf) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *OsInstallAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *OsInstallAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *OsInstallAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *OsInstallAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *OsInstallAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


