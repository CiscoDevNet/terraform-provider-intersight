# PartnerintegrationInventoryAllOf

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

### NewPartnerintegrationInventoryAllOf

`func NewPartnerintegrationInventoryAllOf(classId string, objectType string, ) *PartnerintegrationInventoryAllOf`

NewPartnerintegrationInventoryAllOf instantiates a new PartnerintegrationInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationInventoryAllOfWithDefaults

`func NewPartnerintegrationInventoryAllOfWithDefaults() *PartnerintegrationInventoryAllOf`

NewPartnerintegrationInventoryAllOfWithDefaults instantiates a new PartnerintegrationInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *PartnerintegrationInventoryAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PartnerintegrationInventoryAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PartnerintegrationInventoryAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PartnerintegrationInventoryAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBuildFlags

`func (o *PartnerintegrationInventoryAllOf) GetBuildFlags() string`

GetBuildFlags returns the BuildFlags field if non-nil, zero value otherwise.

### GetBuildFlagsOk

`func (o *PartnerintegrationInventoryAllOf) GetBuildFlagsOk() (*string, bool)`

GetBuildFlagsOk returns a tuple with the BuildFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildFlags

`func (o *PartnerintegrationInventoryAllOf) SetBuildFlags(v string)`

SetBuildFlags sets BuildFlags field to given value.

### HasBuildFlags

`func (o *PartnerintegrationInventoryAllOf) HasBuildFlags() bool`

HasBuildFlags returns a boolean if a field has been set.

### GetBuildStartTime

`func (o *PartnerintegrationInventoryAllOf) GetBuildStartTime() string`

GetBuildStartTime returns the BuildStartTime field if non-nil, zero value otherwise.

### GetBuildStartTimeOk

`func (o *PartnerintegrationInventoryAllOf) GetBuildStartTimeOk() (*string, bool)`

GetBuildStartTimeOk returns a tuple with the BuildStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStartTime

`func (o *PartnerintegrationInventoryAllOf) SetBuildStartTime(v string)`

SetBuildStartTime sets BuildStartTime field to given value.

### HasBuildStartTime

`func (o *PartnerintegrationInventoryAllOf) HasBuildStartTime() bool`

HasBuildStartTime returns a boolean if a field has been set.

### GetBuildStatus

`func (o *PartnerintegrationInventoryAllOf) GetBuildStatus() string`

GetBuildStatus returns the BuildStatus field if non-nil, zero value otherwise.

### GetBuildStatusOk

`func (o *PartnerintegrationInventoryAllOf) GetBuildStatusOk() (*string, bool)`

GetBuildStatusOk returns a tuple with the BuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStatus

`func (o *PartnerintegrationInventoryAllOf) SetBuildStatus(v string)`

SetBuildStatus sets BuildStatus field to given value.

### HasBuildStatus

`func (o *PartnerintegrationInventoryAllOf) HasBuildStatus() bool`

HasBuildStatus returns a boolean if a field has been set.

### GetDeployStartTime

`func (o *PartnerintegrationInventoryAllOf) GetDeployStartTime() string`

GetDeployStartTime returns the DeployStartTime field if non-nil, zero value otherwise.

### GetDeployStartTimeOk

`func (o *PartnerintegrationInventoryAllOf) GetDeployStartTimeOk() (*string, bool)`

GetDeployStartTimeOk returns a tuple with the DeployStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStartTime

`func (o *PartnerintegrationInventoryAllOf) SetDeployStartTime(v string)`

SetDeployStartTime sets DeployStartTime field to given value.

### HasDeployStartTime

`func (o *PartnerintegrationInventoryAllOf) HasDeployStartTime() bool`

HasDeployStartTime returns a boolean if a field has been set.

### GetDeployStatus

`func (o *PartnerintegrationInventoryAllOf) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *PartnerintegrationInventoryAllOf) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *PartnerintegrationInventoryAllOf) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.

### HasDeployStatus

`func (o *PartnerintegrationInventoryAllOf) HasDeployStatus() bool`

HasDeployStatus returns a boolean if a field has been set.

### GetImageName

`func (o *PartnerintegrationInventoryAllOf) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *PartnerintegrationInventoryAllOf) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *PartnerintegrationInventoryAllOf) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *PartnerintegrationInventoryAllOf) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetName

`func (o *PartnerintegrationInventoryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerintegrationInventoryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerintegrationInventoryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartnerintegrationInventoryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPythonSdkUrl

`func (o *PartnerintegrationInventoryAllOf) GetPythonSdkUrl() string`

GetPythonSdkUrl returns the PythonSdkUrl field if non-nil, zero value otherwise.

### GetPythonSdkUrlOk

`func (o *PartnerintegrationInventoryAllOf) GetPythonSdkUrlOk() (*string, bool)`

GetPythonSdkUrlOk returns a tuple with the PythonSdkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonSdkUrl

`func (o *PartnerintegrationInventoryAllOf) SetPythonSdkUrl(v string)`

SetPythonSdkUrl sets PythonSdkUrl field to given value.

### HasPythonSdkUrl

`func (o *PartnerintegrationInventoryAllOf) HasPythonSdkUrl() bool`

HasPythonSdkUrl returns a boolean if a field has been set.

### GetDocIssues

`func (o *PartnerintegrationInventoryAllOf) GetDocIssues() []PartnerintegrationDocIssuesRelationship`

GetDocIssues returns the DocIssues field if non-nil, zero value otherwise.

### GetDocIssuesOk

`func (o *PartnerintegrationInventoryAllOf) GetDocIssuesOk() (*[]PartnerintegrationDocIssuesRelationship, bool)`

GetDocIssuesOk returns a tuple with the DocIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocIssues

`func (o *PartnerintegrationInventoryAllOf) SetDocIssues(v []PartnerintegrationDocIssuesRelationship)`

SetDocIssues sets DocIssues field to given value.

### HasDocIssues

`func (o *PartnerintegrationInventoryAllOf) HasDocIssues() bool`

HasDocIssues returns a boolean if a field has been set.

### SetDocIssuesNil

`func (o *PartnerintegrationInventoryAllOf) SetDocIssuesNil(b bool)`

 SetDocIssuesNil sets the value for DocIssues to be an explicit nil

### UnsetDocIssues
`func (o *PartnerintegrationInventoryAllOf) UnsetDocIssues()`

UnsetDocIssues ensures that no value is present for DocIssues, not even an explicit nil
### GetEtls

`func (o *PartnerintegrationInventoryAllOf) GetEtls() []PartnerintegrationEtlRelationship`

GetEtls returns the Etls field if non-nil, zero value otherwise.

### GetEtlsOk

`func (o *PartnerintegrationInventoryAllOf) GetEtlsOk() (*[]PartnerintegrationEtlRelationship, bool)`

GetEtlsOk returns a tuple with the Etls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtls

`func (o *PartnerintegrationInventoryAllOf) SetEtls(v []PartnerintegrationEtlRelationship)`

SetEtls sets Etls field to given value.

### HasEtls

`func (o *PartnerintegrationInventoryAllOf) HasEtls() bool`

HasEtls returns a boolean if a field has been set.

### SetEtlsNil

`func (o *PartnerintegrationInventoryAllOf) SetEtlsNil(b bool)`

 SetEtlsNil sets the value for Etls to be an explicit nil

### UnsetEtls
`func (o *PartnerintegrationInventoryAllOf) UnsetEtls()`

UnsetEtls ensures that no value is present for Etls, not even an explicit nil
### GetLogs

`func (o *PartnerintegrationInventoryAllOf) GetLogs() []PartnerintegrationLogsRelationship`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PartnerintegrationInventoryAllOf) GetLogsOk() (*[]PartnerintegrationLogsRelationship, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PartnerintegrationInventoryAllOf) SetLogs(v []PartnerintegrationLogsRelationship)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PartnerintegrationInventoryAllOf) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogsNil

`func (o *PartnerintegrationInventoryAllOf) SetLogsNil(b bool)`

 SetLogsNil sets the value for Logs to be an explicit nil

### UnsetLogs
`func (o *PartnerintegrationInventoryAllOf) UnsetLogs()`

UnsetLogs ensures that no value is present for Logs, not even an explicit nil
### GetModels

`func (o *PartnerintegrationInventoryAllOf) GetModels() []PartnerintegrationModelRelationship`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *PartnerintegrationInventoryAllOf) GetModelsOk() (*[]PartnerintegrationModelRelationship, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *PartnerintegrationInventoryAllOf) SetModels(v []PartnerintegrationModelRelationship)`

SetModels sets Models field to given value.

### HasModels

`func (o *PartnerintegrationInventoryAllOf) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *PartnerintegrationInventoryAllOf) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *PartnerintegrationInventoryAllOf) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetOrganization

`func (o *PartnerintegrationInventoryAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerintegrationInventoryAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerintegrationInventoryAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PartnerintegrationInventoryAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


