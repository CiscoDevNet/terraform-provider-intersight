# ConvergedinfraPod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.Pod"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.Pod"]
**DeploymentType** | Pointer to **string** | The deployment type for this integrated system. * &#x60;FlexPodInfra&#x60; - The deployment type for a pod is of Infrastructure. * &#x60;FlexPodNG&#x60; - The deployment type for a pod is of Nextgen type. | [optional] [readonly] [default to "FlexPodInfra"]
**InteropStatus** | Pointer to **string** | The interoperability status for this solution pod. * &#x60;NotEvaluated&#x60; - The interoperability compliance for the component has not be checked. * &#x60;Approved&#x60; - The component is valid as per the interoperability compliance check. * &#x60;NotApproved&#x60; - The component is not valid as per the interoperability compliance check. * &#x60;Incomplete&#x60; - The interoperability compliance check could not be completed for the component due to incomplete data. | [optional] [readonly] [default to "NotEvaluated"]
**Summary** | Pointer to [**ConvergedinfraPodSummary**](ConvergedinfraPodSummary.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**PodCompliance** | Pointer to [**NullableConvergedinfraPodComplianceInfoRelationship**](ConvergedinfraPodComplianceInfoRelationship.md) |  | [optional] 
**PodResourceGroup** | Pointer to [**NullableResourceGroupRelationship**](ResourceGroupRelationship.md) |  | [optional] 
**ServiceItemInstance** | Pointer to [**NullableWorkflowServiceItemInstanceRelationship**](WorkflowServiceItemInstanceRelationship.md) |  | [optional] 

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

### GetInteropStatus

`func (o *ConvergedinfraPod) GetInteropStatus() string`

GetInteropStatus returns the InteropStatus field if non-nil, zero value otherwise.

### GetInteropStatusOk

`func (o *ConvergedinfraPod) GetInteropStatusOk() (*string, bool)`

GetInteropStatusOk returns a tuple with the InteropStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteropStatus

`func (o *ConvergedinfraPod) SetInteropStatus(v string)`

SetInteropStatus sets InteropStatus field to given value.

### HasInteropStatus

`func (o *ConvergedinfraPod) HasInteropStatus() bool`

HasInteropStatus returns a boolean if a field has been set.

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

### SetOrganizationNil

`func (o *ConvergedinfraPod) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ConvergedinfraPod) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetPodCompliance

`func (o *ConvergedinfraPod) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship`

GetPodCompliance returns the PodCompliance field if non-nil, zero value otherwise.

### GetPodComplianceOk

`func (o *ConvergedinfraPod) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool)`

GetPodComplianceOk returns a tuple with the PodCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCompliance

`func (o *ConvergedinfraPod) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship)`

SetPodCompliance sets PodCompliance field to given value.

### HasPodCompliance

`func (o *ConvergedinfraPod) HasPodCompliance() bool`

HasPodCompliance returns a boolean if a field has been set.

### SetPodComplianceNil

`func (o *ConvergedinfraPod) SetPodComplianceNil(b bool)`

 SetPodComplianceNil sets the value for PodCompliance to be an explicit nil

### UnsetPodCompliance
`func (o *ConvergedinfraPod) UnsetPodCompliance()`

UnsetPodCompliance ensures that no value is present for PodCompliance, not even an explicit nil
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

### SetPodResourceGroupNil

`func (o *ConvergedinfraPod) SetPodResourceGroupNil(b bool)`

 SetPodResourceGroupNil sets the value for PodResourceGroup to be an explicit nil

### UnsetPodResourceGroup
`func (o *ConvergedinfraPod) UnsetPodResourceGroup()`

UnsetPodResourceGroup ensures that no value is present for PodResourceGroup, not even an explicit nil
### GetServiceItemInstance

`func (o *ConvergedinfraPod) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *ConvergedinfraPod) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *ConvergedinfraPod) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *ConvergedinfraPod) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.

### SetServiceItemInstanceNil

`func (o *ConvergedinfraPod) SetServiceItemInstanceNil(b bool)`

 SetServiceItemInstanceNil sets the value for ServiceItemInstance to be an explicit nil

### UnsetServiceItemInstance
`func (o *ConvergedinfraPod) UnsetServiceItemInstance()`

UnsetServiceItemInstance ensures that no value is present for ServiceItemInstance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


