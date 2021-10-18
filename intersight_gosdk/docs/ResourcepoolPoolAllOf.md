# ResourcepoolPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.Pool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.Pool"]
**PoolType** | Pointer to **string** | The resource management type in the pool, it can be either static or dynamic. * &#x60;Static&#x60; - The resources in the pool will not be changed until user manually update it. * &#x60;Dynamic&#x60; - The resources in the pool will be updated dynamically based on the condition. | [optional] [default to "Static"]
**ResourcePoolParameters** | Pointer to [**NullableResourcepoolResourcePoolParameters**](ResourcepoolResourcePoolParameters.md) |  | [optional] 
**ResourceType** | Pointer to **string** | The type of the resource present in the pool, example &#39;server&#39; its combination of RackUnit and Blade. * &#x60;None&#x60; - The resource cannot consider for Resource Pool. * &#x60;Server&#x60; - Resource Pool holds the server kind of resources, example - RackServer, Blade. | [optional] [default to "None"]
**Selectors** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewResourcepoolPoolAllOf

`func NewResourcepoolPoolAllOf(classId string, objectType string, ) *ResourcepoolPoolAllOf`

NewResourcepoolPoolAllOf instantiates a new ResourcepoolPoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolPoolAllOfWithDefaults

`func NewResourcepoolPoolAllOfWithDefaults() *ResourcepoolPoolAllOf`

NewResourcepoolPoolAllOfWithDefaults instantiates a new ResourcepoolPoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolPoolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolPoolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolPoolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolPoolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolPoolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolPoolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPoolType

`func (o *ResourcepoolPoolAllOf) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *ResourcepoolPoolAllOf) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *ResourcepoolPoolAllOf) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *ResourcepoolPoolAllOf) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetResourcePoolParameters

`func (o *ResourcepoolPoolAllOf) GetResourcePoolParameters() ResourcepoolResourcePoolParameters`

GetResourcePoolParameters returns the ResourcePoolParameters field if non-nil, zero value otherwise.

### GetResourcePoolParametersOk

`func (o *ResourcepoolPoolAllOf) GetResourcePoolParametersOk() (*ResourcepoolResourcePoolParameters, bool)`

GetResourcePoolParametersOk returns a tuple with the ResourcePoolParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolParameters

`func (o *ResourcepoolPoolAllOf) SetResourcePoolParameters(v ResourcepoolResourcePoolParameters)`

SetResourcePoolParameters sets ResourcePoolParameters field to given value.

### HasResourcePoolParameters

`func (o *ResourcepoolPoolAllOf) HasResourcePoolParameters() bool`

HasResourcePoolParameters returns a boolean if a field has been set.

### SetResourcePoolParametersNil

`func (o *ResourcepoolPoolAllOf) SetResourcePoolParametersNil(b bool)`

 SetResourcePoolParametersNil sets the value for ResourcePoolParameters to be an explicit nil

### UnsetResourcePoolParameters
`func (o *ResourcepoolPoolAllOf) UnsetResourcePoolParameters()`

UnsetResourcePoolParameters ensures that no value is present for ResourcePoolParameters, not even an explicit nil
### GetResourceType

`func (o *ResourcepoolPoolAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourcepoolPoolAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourcepoolPoolAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourcepoolPoolAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSelectors

`func (o *ResourcepoolPoolAllOf) GetSelectors() []ResourceSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *ResourcepoolPoolAllOf) GetSelectorsOk() (*[]ResourceSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *ResourcepoolPoolAllOf) SetSelectors(v []ResourceSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *ResourcepoolPoolAllOf) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### SetSelectorsNil

`func (o *ResourcepoolPoolAllOf) SetSelectorsNil(b bool)`

 SetSelectorsNil sets the value for Selectors to be an explicit nil

### UnsetSelectors
`func (o *ResourcepoolPoolAllOf) UnsetSelectors()`

UnsetSelectors ensures that no value is present for Selectors, not even an explicit nil
### GetOrganization

`func (o *ResourcepoolPoolAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResourcepoolPoolAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResourcepoolPoolAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ResourcepoolPoolAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


