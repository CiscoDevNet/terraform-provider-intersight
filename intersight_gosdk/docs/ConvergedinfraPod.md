# ConvergedinfraPod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.Pod"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.Pod"]
**DeploymentType** | Pointer to **string** | The deployment type for this solution pod. * &#x60;FlexPodInfra&#x60; - The deployment type for a pod is of Infrastructure. * &#x60;FlexPodNG&#x60; - The deployment type for a pod is of Nextgen type. | [optional] [readonly] [default to "FlexPodInfra"]
**Summary** | Pointer to [**ConvergedinfraPodSummary**](ConvergedinfraPodSummary.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**PodResourceGroup** | Pointer to [**ResourceGroupRelationship**](ResourceGroupRelationship.md) |  | [optional] 
**SolutionInstance** | Pointer to [**WorkflowSolutionInstanceRelationship**](WorkflowSolutionInstanceRelationship.md) |  | [optional] 

## Methods

### NewConvergedinfraPod

`func NewConvergedinfraPod(classId string, objectType string, ) *ConvergedinfraPod`

NewConvergedinfraPod instantiates a new ConvergedinfraPod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraPodWithDefaults

`func NewConvergedinfraPodWithDefaults() *ConvergedinfraPod`

NewConvergedinfraPodWithDefaults instantiates a new ConvergedinfraPod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraPod) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraPod) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraPod) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraPod) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraPod) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraPod) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeploymentType

`func (o *ConvergedinfraPod) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *ConvergedinfraPod) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *ConvergedinfraPod) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *ConvergedinfraPod) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetSummary

`func (o *ConvergedinfraPod) GetSummary() ConvergedinfraPodSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ConvergedinfraPod) GetSummaryOk() (*ConvergedinfraPodSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ConvergedinfraPod) SetSummary(v ConvergedinfraPodSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ConvergedinfraPod) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetOrganization

`func (o *ConvergedinfraPod) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ConvergedinfraPod) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ConvergedinfraPod) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ConvergedinfraPod) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPodResourceGroup

`func (o *ConvergedinfraPod) GetPodResourceGroup() ResourceGroupRelationship`

GetPodResourceGroup returns the PodResourceGroup field if non-nil, zero value otherwise.

### GetPodResourceGroupOk

`func (o *ConvergedinfraPod) GetPodResourceGroupOk() (*ResourceGroupRelationship, bool)`

GetPodResourceGroupOk returns a tuple with the PodResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodResourceGroup

`func (o *ConvergedinfraPod) SetPodResourceGroup(v ResourceGroupRelationship)`

SetPodResourceGroup sets PodResourceGroup field to given value.

### HasPodResourceGroup

`func (o *ConvergedinfraPod) HasPodResourceGroup() bool`

HasPodResourceGroup returns a boolean if a field has been set.

### GetSolutionInstance

`func (o *ConvergedinfraPod) GetSolutionInstance() WorkflowSolutionInstanceRelationship`

GetSolutionInstance returns the SolutionInstance field if non-nil, zero value otherwise.

### GetSolutionInstanceOk

`func (o *ConvergedinfraPod) GetSolutionInstanceOk() (*WorkflowSolutionInstanceRelationship, bool)`

GetSolutionInstanceOk returns a tuple with the SolutionInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionInstance

`func (o *ConvergedinfraPod) SetSolutionInstance(v WorkflowSolutionInstanceRelationship)`

SetSolutionInstance sets SolutionInstance field to given value.

### HasSolutionInstance

`func (o *ConvergedinfraPod) HasSolutionInstance() bool`

HasSolutionInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


