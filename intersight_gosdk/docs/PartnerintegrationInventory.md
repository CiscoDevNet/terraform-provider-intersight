# PartnerintegrationInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.Inventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.Inventory"]
**Action** | Pointer to **string** | Action to be performed on the inventory collection. * &#x60;None&#x60; - Default Value of the action, i.e. do nothing. * &#x60;Build&#x60; - Build the inventory service image. * &#x60;Deploy&#x60; - Deploy the inventory service on the appliance. | [optional] [default to "None"]
**BuildFlags** | Pointer to **string** | Addtional flags to control build action. | [optional] 
**BuildStartTime** | Pointer to **string** | Time when build was triggered. | [optional] [readonly] 
**BuildStatus** | Pointer to **string** | Status of build for inventory collection. * &#x60;None&#x60; - Default value of the status. i.e. done nothing. * &#x60;BackendInProgress&#x60; - The backend build is in progress. * &#x60;BackendFailed&#x60; - The backend build has failed. * &#x60;DockerInProgress&#x60; - The docker build is in progress. * &#x60;DockerFailed&#x60; - The docker build has failed. * &#x60;UiInProgress&#x60; - The UI build is in progress. * &#x60;UiFailed&#x60; - The inventory UI build has failed. * &#x60;ApidocsInProgress&#x60; - The apidocs build is in progress. * &#x60;ApidocsFailed&#x60; - The apidocs build has failed. * &#x60;Completed&#x60; - The operation completed successfully. | [optional] [readonly] [default to "None"]
**DeployStartTime** | Pointer to **string** | Time when deploy was triggered. | [optional] [readonly] 
**DeployStatus** | Pointer to **string** | Status of deployment of the inventory microservice. * &#x60;None&#x60; - Default value of the status. i.e. done nothing. * &#x60;Completed&#x60; - The operation completed successfully. * &#x60;Failed&#x60; - The deploy operation failed. | [optional] [readonly] [default to "None"]
**ImageName** | Pointer to **string** | Name of the docker image that is built. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the inventory collection. | [optional] 
**PythonSdkUrl** | Pointer to **string** | Link to the generated v3 python SDK. | [optional] [readonly] 
**DocIssues** | Pointer to [**[]PartnerintegrationDocIssuesRelationship**](PartnerintegrationDocIssuesRelationship.md) | An array of relationships to partnerintegrationDocIssues resources. | [optional] [readonly] 
**Etls** | Pointer to [**[]PartnerintegrationEtlRelationship**](PartnerintegrationEtlRelationship.md) | An array of relationships to partnerintegrationEtl resources. | [optional] 
**Logs** | Pointer to [**[]PartnerintegrationLogsRelationship**](PartnerintegrationLogsRelationship.md) | An array of relationships to partnerintegrationLogs resources. | [optional] [readonly] 
**Models** | Pointer to [**[]PartnerintegrationModelRelationship**](PartnerintegrationModelRelationship.md) | An array of relationships to partnerintegrationModel resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationInventory

`func NewPartnerintegrationInventory(classId string, objectType string, ) *PartnerintegrationInventory`

NewPartnerintegrationInventory instantiates a new PartnerintegrationInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationInventoryWithDefaults

`func NewPartnerintegrationInventoryWithDefaults() *PartnerintegrationInventory`

NewPartnerintegrationInventoryWithDefaults instantiates a new PartnerintegrationInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *PartnerintegrationInventory) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PartnerintegrationInventory) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PartnerintegrationInventory) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PartnerintegrationInventory) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBuildFlags

`func (o *PartnerintegrationInventory) GetBuildFlags() string`

GetBuildFlags returns the BuildFlags field if non-nil, zero value otherwise.

### GetBuildFlagsOk

`func (o *PartnerintegrationInventory) GetBuildFlagsOk() (*string, bool)`

GetBuildFlagsOk returns a tuple with the BuildFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildFlags

`func (o *PartnerintegrationInventory) SetBuildFlags(v string)`

SetBuildFlags sets BuildFlags field to given value.

### HasBuildFlags

`func (o *PartnerintegrationInventory) HasBuildFlags() bool`

HasBuildFlags returns a boolean if a field has been set.

### GetBuildStartTime

`func (o *PartnerintegrationInventory) GetBuildStartTime() string`

GetBuildStartTime returns the BuildStartTime field if non-nil, zero value otherwise.

### GetBuildStartTimeOk

`func (o *PartnerintegrationInventory) GetBuildStartTimeOk() (*string, bool)`

GetBuildStartTimeOk returns a tuple with the BuildStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStartTime

`func (o *PartnerintegrationInventory) SetBuildStartTime(v string)`

SetBuildStartTime sets BuildStartTime field to given value.

### HasBuildStartTime

`func (o *PartnerintegrationInventory) HasBuildStartTime() bool`

HasBuildStartTime returns a boolean if a field has been set.

### GetBuildStatus

`func (o *PartnerintegrationInventory) GetBuildStatus() string`

GetBuildStatus returns the BuildStatus field if non-nil, zero value otherwise.

### GetBuildStatusOk

`func (o *PartnerintegrationInventory) GetBuildStatusOk() (*string, bool)`

GetBuildStatusOk returns a tuple with the BuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStatus

`func (o *PartnerintegrationInventory) SetBuildStatus(v string)`

SetBuildStatus sets BuildStatus field to given value.

### HasBuildStatus

`func (o *PartnerintegrationInventory) HasBuildStatus() bool`

HasBuildStatus returns a boolean if a field has been set.

### GetDeployStartTime

`func (o *PartnerintegrationInventory) GetDeployStartTime() string`

GetDeployStartTime returns the DeployStartTime field if non-nil, zero value otherwise.

### GetDeployStartTimeOk

`func (o *PartnerintegrationInventory) GetDeployStartTimeOk() (*string, bool)`

GetDeployStartTimeOk returns a tuple with the DeployStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStartTime

`func (o *PartnerintegrationInventory) SetDeployStartTime(v string)`

SetDeployStartTime sets DeployStartTime field to given value.

### HasDeployStartTime

`func (o *PartnerintegrationInventory) HasDeployStartTime() bool`

HasDeployStartTime returns a boolean if a field has been set.

### GetDeployStatus

`func (o *PartnerintegrationInventory) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *PartnerintegrationInventory) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *PartnerintegrationInventory) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.

### HasDeployStatus

`func (o *PartnerintegrationInventory) HasDeployStatus() bool`

HasDeployStatus returns a boolean if a field has been set.

### GetImageName

`func (o *PartnerintegrationInventory) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *PartnerintegrationInventory) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *PartnerintegrationInventory) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *PartnerintegrationInventory) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetName

`func (o *PartnerintegrationInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerintegrationInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerintegrationInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartnerintegrationInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPythonSdkUrl

`func (o *PartnerintegrationInventory) GetPythonSdkUrl() string`

GetPythonSdkUrl returns the PythonSdkUrl field if non-nil, zero value otherwise.

### GetPythonSdkUrlOk

`func (o *PartnerintegrationInventory) GetPythonSdkUrlOk() (*string, bool)`

GetPythonSdkUrlOk returns a tuple with the PythonSdkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonSdkUrl

`func (o *PartnerintegrationInventory) SetPythonSdkUrl(v string)`

SetPythonSdkUrl sets PythonSdkUrl field to given value.

### HasPythonSdkUrl

`func (o *PartnerintegrationInventory) HasPythonSdkUrl() bool`

HasPythonSdkUrl returns a boolean if a field has been set.

### GetDocIssues

`func (o *PartnerintegrationInventory) GetDocIssues() []PartnerintegrationDocIssuesRelationship`

GetDocIssues returns the DocIssues field if non-nil, zero value otherwise.

### GetDocIssuesOk

`func (o *PartnerintegrationInventory) GetDocIssuesOk() (*[]PartnerintegrationDocIssuesRelationship, bool)`

GetDocIssuesOk returns a tuple with the DocIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocIssues

`func (o *PartnerintegrationInventory) SetDocIssues(v []PartnerintegrationDocIssuesRelationship)`

SetDocIssues sets DocIssues field to given value.

### HasDocIssues

`func (o *PartnerintegrationInventory) HasDocIssues() bool`

HasDocIssues returns a boolean if a field has been set.

### SetDocIssuesNil

`func (o *PartnerintegrationInventory) SetDocIssuesNil(b bool)`

 SetDocIssuesNil sets the value for DocIssues to be an explicit nil

### UnsetDocIssues
`func (o *PartnerintegrationInventory) UnsetDocIssues()`

UnsetDocIssues ensures that no value is present for DocIssues, not even an explicit nil
### GetEtls

`func (o *PartnerintegrationInventory) GetEtls() []PartnerintegrationEtlRelationship`

GetEtls returns the Etls field if non-nil, zero value otherwise.

### GetEtlsOk

`func (o *PartnerintegrationInventory) GetEtlsOk() (*[]PartnerintegrationEtlRelationship, bool)`

GetEtlsOk returns a tuple with the Etls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtls

`func (o *PartnerintegrationInventory) SetEtls(v []PartnerintegrationEtlRelationship)`

SetEtls sets Etls field to given value.

### HasEtls

`func (o *PartnerintegrationInventory) HasEtls() bool`

HasEtls returns a boolean if a field has been set.

### SetEtlsNil

`func (o *PartnerintegrationInventory) SetEtlsNil(b bool)`

 SetEtlsNil sets the value for Etls to be an explicit nil

### UnsetEtls
`func (o *PartnerintegrationInventory) UnsetEtls()`

UnsetEtls ensures that no value is present for Etls, not even an explicit nil
### GetLogs

`func (o *PartnerintegrationInventory) GetLogs() []PartnerintegrationLogsRelationship`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PartnerintegrationInventory) GetLogsOk() (*[]PartnerintegrationLogsRelationship, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PartnerintegrationInventory) SetLogs(v []PartnerintegrationLogsRelationship)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PartnerintegrationInventory) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogsNil

`func (o *PartnerintegrationInventory) SetLogsNil(b bool)`

 SetLogsNil sets the value for Logs to be an explicit nil

### UnsetLogs
`func (o *PartnerintegrationInventory) UnsetLogs()`

UnsetLogs ensures that no value is present for Logs, not even an explicit nil
### GetModels

`func (o *PartnerintegrationInventory) GetModels() []PartnerintegrationModelRelationship`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *PartnerintegrationInventory) GetModelsOk() (*[]PartnerintegrationModelRelationship, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *PartnerintegrationInventory) SetModels(v []PartnerintegrationModelRelationship)`

SetModels sets Models field to given value.

### HasModels

`func (o *PartnerintegrationInventory) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *PartnerintegrationInventory) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *PartnerintegrationInventory) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetOrganization

`func (o *PartnerintegrationInventory) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerintegrationInventory) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerintegrationInventory) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PartnerintegrationInventory) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


