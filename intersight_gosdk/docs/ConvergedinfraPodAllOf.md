# ConvergedinfraPodAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.Pod"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.Pod"]
**DeploymentType** | Pointer to **string** | The deployment type for this integrated system. * &#x60;FlexPodInfra&#x60; - The deployment type for a pod is of Infrastructure. * &#x60;FlexPodNG&#x60; - The deployment type for a pod is of Nextgen type. | [optional] [readonly] [default to "FlexPodInfra"]
**InteropStatus** | Pointer to **string** | The interoperability status for this solution pod. * &#x60;NotEvaluated&#x60; - The interoperability compliance for the component has not be checked. * &#x60;Approved&#x60; - The component is valid as per the interoperability compliance check. * &#x60;NotApproved&#x60; - The component is not valid as per the interoperability compliance check. * &#x60;Incomplete&#x60; - The interoperability compliance check could not be completed for the component due to incomplete data. | [optional] [readonly] [default to "NotEvaluated"]
**Summary** | Pointer to [**ConvergedinfraPodSummary**](ConvergedinfraPodSummary.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**PodCompliance** | Pointer to [**ConvergedinfraPodComplianceInfoRelationship**](ConvergedinfraPodComplianceInfoRelationship.md) |  | [optional] 
**PodResourceGroup** | Pointer to [**ResourceGroupRelationship**](ResourceGroupRelationship.md) |  | [optional] 
**ServiceItemInstance** | Pointer to [**WorkflowServiceItemInstanceRelationship**](WorkflowServiceItemInstanceRelationship.md) |  | [optional] 

## Methods

### NewConvergedinfraPodAllOf

`func NewConvergedinfraPodAllOf(classId string, objectType string, ) *ConvergedinfraPodAllOf`

NewConvergedinfraPodAllOf instantiates a new ConvergedinfraPodAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraPodAllOfWithDefaults

`func NewConvergedinfraPodAllOfWithDefaults() *ConvergedinfraPodAllOf`

NewConvergedinfraPodAllOfWithDefaults instantiates a new ConvergedinfraPodAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraPodAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraPodAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraPodAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraPodAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraPodAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraPodAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeploymentType

`func (o *ConvergedinfraPodAllOf) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *ConvergedinfraPodAllOf) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *ConvergedinfraPodAllOf) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *ConvergedinfraPodAllOf) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetInteropStatus

`func (o *ConvergedinfraPodAllOf) GetInteropStatus() string`

GetInteropStatus returns the InteropStatus field if non-nil, zero value otherwise.

### GetInteropStatusOk

`func (o *ConvergedinfraPodAllOf) GetInteropStatusOk() (*string, bool)`

GetInteropStatusOk returns a tuple with the InteropStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteropStatus

`func (o *ConvergedinfraPodAllOf) SetInteropStatus(v string)`

SetInteropStatus sets InteropStatus field to given value.

### HasInteropStatus

`func (o *ConvergedinfraPodAllOf) HasInteropStatus() bool`

HasInteropStatus returns a boolean if a field has been set.

### GetSummary

`func (o *ConvergedinfraPodAllOf) GetSummary() ConvergedinfraPodSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ConvergedinfraPodAllOf) GetSummaryOk() (*ConvergedinfraPodSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ConvergedinfraPodAllOf) SetSummary(v ConvergedinfraPodSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ConvergedinfraPodAllOf) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetOrganization

`func (o *ConvergedinfraPodAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ConvergedinfraPodAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ConvergedinfraPodAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ConvergedinfraPodAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPodCompliance

`func (o *ConvergedinfraPodAllOf) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship`

GetPodCompliance returns the PodCompliance field if non-nil, zero value otherwise.

### GetPodComplianceOk

`func (o *ConvergedinfraPodAllOf) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool)`

GetPodComplianceOk returns a tuple with the PodCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCompliance

`func (o *ConvergedinfraPodAllOf) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship)`

SetPodCompliance sets PodCompliance field to given value.

### HasPodCompliance

`func (o *ConvergedinfraPodAllOf) HasPodCompliance() bool`

HasPodCompliance returns a boolean if a field has been set.

### GetPodResourceGroup

`func (o *ConvergedinfraPodAllOf) GetPodResourceGroup() ResourceGroupRelationship`

GetPodResourceGroup returns the PodResourceGroup field if non-nil, zero value otherwise.

### GetPodResourceGroupOk

`func (o *ConvergedinfraPodAllOf) GetPodResourceGroupOk() (*ResourceGroupRelationship, bool)`

GetPodResourceGroupOk returns a tuple with the PodResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodResourceGroup

`func (o *ConvergedinfraPodAllOf) SetPodResourceGroup(v ResourceGroupRelationship)`

SetPodResourceGroup sets PodResourceGroup field to given value.

### HasPodResourceGroup

`func (o *ConvergedinfraPodAllOf) HasPodResourceGroup() bool`

HasPodResourceGroup returns a boolean if a field has been set.

### GetServiceItemInstance

`func (o *ConvergedinfraPodAllOf) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *ConvergedinfraPodAllOf) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *ConvergedinfraPodAllOf) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *ConvergedinfraPodAllOf) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


