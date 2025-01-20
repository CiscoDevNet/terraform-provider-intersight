# ResourcepoolPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.Pool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.Pool"]
**Action** | Pointer to **string** | The pool is evaluated for resources with associated policies based on action. This action will help users to re-sync the resources for a pool. * &#x60;None&#x60; - The pool will not be considered for evaluation. * &#x60;ReEvaluate&#x60; - The resources in the pool will be re-evaluated against the server pool qualification associated with it. | [optional] [default to "None"]
**PoolType** | Pointer to **string** | The resource management type in the pool, it can be either static or dynamic. * &#x60;Static&#x60; - The resources in the pool will not be changed until user manually update it. * &#x60;Dynamic&#x60; - The resources in the pool will be updated dynamically based on the condition. * &#x60;Hybrid&#x60; - The resources in the pool can be added by the user statically or dynamically, based on the matching conditions of the qualification policy. If the pool contains both statically added resources and resources added based on the qualification policy, the pool type can be classified as hybrid. | [optional] [default to "Static"]
**ResourceEvaluationStatus** | Pointer to [**NullableResourcepoolResourceEvaluationStatus**](ResourcepoolResourceEvaluationStatus.md) |  | [optional] 
**ResourcePoolParameters** | Pointer to [**NullableMoBaseComplexType**](MoBaseComplexType.md) | The resource pool can hold different type of resources, each resources can have some specific parameters and functionality, those details are captured as part of this. | [optional] 
**ResourceType** | Pointer to **string** | The type of the resource present in the pool, example &#39;server&#39; its combination of RackUnit and Blade. * &#x60;Server&#x60; - Resource Pool holds the server kind of resources, example - RackServer, Blade. * &#x60;None&#x60; - The resource cannot consider for Resource Pool. | [optional] [default to "Server"]
**Selectors** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**QualificationPolicies** | Pointer to [**[]ResourcepoolQualificationPolicyRelationship**](ResourcepoolQualificationPolicyRelationship.md) | An array of relationships to resourcepoolQualificationPolicy resources. | [optional] 

## Methods

### NewResourcepoolPool

`func NewResourcepoolPool(classId string, objectType string, ) *ResourcepoolPool`

NewResourcepoolPool instantiates a new ResourcepoolPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolPoolWithDefaults

`func NewResourcepoolPoolWithDefaults() *ResourcepoolPool`

NewResourcepoolPoolWithDefaults instantiates a new ResourcepoolPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolPool) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolPool) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolPool) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolPool) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolPool) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolPool) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *ResourcepoolPool) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ResourcepoolPool) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ResourcepoolPool) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ResourcepoolPool) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPoolType

`func (o *ResourcepoolPool) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *ResourcepoolPool) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *ResourcepoolPool) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *ResourcepoolPool) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetResourceEvaluationStatus

`func (o *ResourcepoolPool) GetResourceEvaluationStatus() ResourcepoolResourceEvaluationStatus`

GetResourceEvaluationStatus returns the ResourceEvaluationStatus field if non-nil, zero value otherwise.

### GetResourceEvaluationStatusOk

`func (o *ResourcepoolPool) GetResourceEvaluationStatusOk() (*ResourcepoolResourceEvaluationStatus, bool)`

GetResourceEvaluationStatusOk returns a tuple with the ResourceEvaluationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceEvaluationStatus

`func (o *ResourcepoolPool) SetResourceEvaluationStatus(v ResourcepoolResourceEvaluationStatus)`

SetResourceEvaluationStatus sets ResourceEvaluationStatus field to given value.

### HasResourceEvaluationStatus

`func (o *ResourcepoolPool) HasResourceEvaluationStatus() bool`

HasResourceEvaluationStatus returns a boolean if a field has been set.

### SetResourceEvaluationStatusNil

`func (o *ResourcepoolPool) SetResourceEvaluationStatusNil(b bool)`

 SetResourceEvaluationStatusNil sets the value for ResourceEvaluationStatus to be an explicit nil

### UnsetResourceEvaluationStatus
`func (o *ResourcepoolPool) UnsetResourceEvaluationStatus()`

UnsetResourceEvaluationStatus ensures that no value is present for ResourceEvaluationStatus, not even an explicit nil
### GetResourcePoolParameters

`func (o *ResourcepoolPool) GetResourcePoolParameters() MoBaseComplexType`

GetResourcePoolParameters returns the ResourcePoolParameters field if non-nil, zero value otherwise.

### GetResourcePoolParametersOk

`func (o *ResourcepoolPool) GetResourcePoolParametersOk() (*MoBaseComplexType, bool)`

GetResourcePoolParametersOk returns a tuple with the ResourcePoolParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolParameters

`func (o *ResourcepoolPool) SetResourcePoolParameters(v MoBaseComplexType)`

SetResourcePoolParameters sets ResourcePoolParameters field to given value.

### HasResourcePoolParameters

`func (o *ResourcepoolPool) HasResourcePoolParameters() bool`

HasResourcePoolParameters returns a boolean if a field has been set.

### SetResourcePoolParametersNil

`func (o *ResourcepoolPool) SetResourcePoolParametersNil(b bool)`

 SetResourcePoolParametersNil sets the value for ResourcePoolParameters to be an explicit nil

### UnsetResourcePoolParameters
`func (o *ResourcepoolPool) UnsetResourcePoolParameters()`

UnsetResourcePoolParameters ensures that no value is present for ResourcePoolParameters, not even an explicit nil
### GetResourceType

`func (o *ResourcepoolPool) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourcepoolPool) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourcepoolPool) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourcepoolPool) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSelectors

`func (o *ResourcepoolPool) GetSelectors() []ResourceSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *ResourcepoolPool) GetSelectorsOk() (*[]ResourceSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *ResourcepoolPool) SetSelectors(v []ResourceSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *ResourcepoolPool) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### SetSelectorsNil

`func (o *ResourcepoolPool) SetSelectorsNil(b bool)`

 SetSelectorsNil sets the value for Selectors to be an explicit nil

### UnsetSelectors
`func (o *ResourcepoolPool) UnsetSelectors()`

UnsetSelectors ensures that no value is present for Selectors, not even an explicit nil
### GetOrganization

`func (o *ResourcepoolPool) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResourcepoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResourcepoolPool) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ResourcepoolPool) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ResourcepoolPool) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ResourcepoolPool) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetQualificationPolicies

`func (o *ResourcepoolPool) GetQualificationPolicies() []ResourcepoolQualificationPolicyRelationship`

GetQualificationPolicies returns the QualificationPolicies field if non-nil, zero value otherwise.

### GetQualificationPoliciesOk

`func (o *ResourcepoolPool) GetQualificationPoliciesOk() (*[]ResourcepoolQualificationPolicyRelationship, bool)`

GetQualificationPoliciesOk returns a tuple with the QualificationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualificationPolicies

`func (o *ResourcepoolPool) SetQualificationPolicies(v []ResourcepoolQualificationPolicyRelationship)`

SetQualificationPolicies sets QualificationPolicies field to given value.

### HasQualificationPolicies

`func (o *ResourcepoolPool) HasQualificationPolicies() bool`

HasQualificationPolicies returns a boolean if a field has been set.

### SetQualificationPoliciesNil

`func (o *ResourcepoolPool) SetQualificationPoliciesNil(b bool)`

 SetQualificationPoliciesNil sets the value for QualificationPolicies to be an explicit nil

### UnsetQualificationPolicies
`func (o *ResourcepoolPool) UnsetQualificationPolicies()`

UnsetQualificationPolicies ensures that no value is present for QualificationPolicies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


