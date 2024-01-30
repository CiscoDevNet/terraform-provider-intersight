# ResourcepoolLease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.Lease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.Lease"]
**Condition** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**Feature** | Pointer to **string** | Lease opertion applied for the feature. | [optional] 
**IsExclusiveAtAssignedEntity** | Pointer to **bool** | Indicates whether a lease allocation is exclusive based on the Assigned Entity, if the AssignedEntity holds any lease then not allowed to create new lease later. | [optional] [default to false]
**LeaseParameters** | Pointer to [**NullableResourcepoolLeaseParameters**](ResourcepoolLeaseParameters.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMo**](MoBaseMo.md) |  | [optional] 
**ResourceType** | Pointer to **string** | The type of resource present in the pool, such as &#39;server&#39; can be a RackUnit or Blade. * &#x60;None&#x60; - The resource cannot consider for Resource Pool. * &#x60;Server&#x60; - Resource Pool holds the server kind of resources, example - RackServer, Blade. | [optional] [default to "None"]
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**LeasedResource** | Pointer to [**ResourcepoolLeaseResourceRelationship**](ResourcepoolLeaseResourceRelationship.md) |  | [optional] 
**Pool** | Pointer to [**ResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**ResourcepoolPoolMemberRelationship**](ResourcepoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**ResourcepoolUniverseRelationship**](ResourcepoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewResourcepoolLease

`func NewResourcepoolLease(classId string, objectType string, ) *ResourcepoolLease`

NewResourcepoolLease instantiates a new ResourcepoolLease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolLeaseWithDefaults

`func NewResourcepoolLeaseWithDefaults() *ResourcepoolLease`

NewResourcepoolLeaseWithDefaults instantiates a new ResourcepoolLease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolLease) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolLease) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolLease) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolLease) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolLease) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolLease) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCondition

`func (o *ResourcepoolLease) GetCondition() []ResourceSelector`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ResourcepoolLease) GetConditionOk() (*[]ResourceSelector, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ResourcepoolLease) SetCondition(v []ResourceSelector)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ResourcepoolLease) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ResourcepoolLease) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ResourcepoolLease) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetFeature

`func (o *ResourcepoolLease) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *ResourcepoolLease) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *ResourcepoolLease) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *ResourcepoolLease) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetIsExclusiveAtAssignedEntity

`func (o *ResourcepoolLease) GetIsExclusiveAtAssignedEntity() bool`

GetIsExclusiveAtAssignedEntity returns the IsExclusiveAtAssignedEntity field if non-nil, zero value otherwise.

### GetIsExclusiveAtAssignedEntityOk

`func (o *ResourcepoolLease) GetIsExclusiveAtAssignedEntityOk() (*bool, bool)`

GetIsExclusiveAtAssignedEntityOk returns a tuple with the IsExclusiveAtAssignedEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusiveAtAssignedEntity

`func (o *ResourcepoolLease) SetIsExclusiveAtAssignedEntity(v bool)`

SetIsExclusiveAtAssignedEntity sets IsExclusiveAtAssignedEntity field to given value.

### HasIsExclusiveAtAssignedEntity

`func (o *ResourcepoolLease) HasIsExclusiveAtAssignedEntity() bool`

HasIsExclusiveAtAssignedEntity returns a boolean if a field has been set.

### GetLeaseParameters

`func (o *ResourcepoolLease) GetLeaseParameters() ResourcepoolLeaseParameters`

GetLeaseParameters returns the LeaseParameters field if non-nil, zero value otherwise.

### GetLeaseParametersOk

`func (o *ResourcepoolLease) GetLeaseParametersOk() (*ResourcepoolLeaseParameters, bool)`

GetLeaseParametersOk returns a tuple with the LeaseParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseParameters

`func (o *ResourcepoolLease) SetLeaseParameters(v ResourcepoolLeaseParameters)`

SetLeaseParameters sets LeaseParameters field to given value.

### HasLeaseParameters

`func (o *ResourcepoolLease) HasLeaseParameters() bool`

HasLeaseParameters returns a boolean if a field has been set.

### SetLeaseParametersNil

`func (o *ResourcepoolLease) SetLeaseParametersNil(b bool)`

 SetLeaseParametersNil sets the value for LeaseParameters to be an explicit nil

### UnsetLeaseParameters
`func (o *ResourcepoolLease) UnsetLeaseParameters()`

UnsetLeaseParameters ensures that no value is present for LeaseParameters, not even an explicit nil
### GetResource

`func (o *ResourcepoolLease) GetResource() MoBaseMo`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourcepoolLease) GetResourceOk() (*MoBaseMo, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourcepoolLease) SetResource(v MoBaseMo)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourcepoolLease) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourcepoolLease) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourcepoolLease) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourcepoolLease) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourcepoolLease) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *ResourcepoolLease) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *ResourcepoolLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *ResourcepoolLease) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *ResourcepoolLease) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetLeasedResource

`func (o *ResourcepoolLease) GetLeasedResource() ResourcepoolLeaseResourceRelationship`

GetLeasedResource returns the LeasedResource field if non-nil, zero value otherwise.

### GetLeasedResourceOk

`func (o *ResourcepoolLease) GetLeasedResourceOk() (*ResourcepoolLeaseResourceRelationship, bool)`

GetLeasedResourceOk returns a tuple with the LeasedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasedResource

`func (o *ResourcepoolLease) SetLeasedResource(v ResourcepoolLeaseResourceRelationship)`

SetLeasedResource sets LeasedResource field to given value.

### HasLeasedResource

`func (o *ResourcepoolLease) HasLeasedResource() bool`

HasLeasedResource returns a boolean if a field has been set.

### GetPool

`func (o *ResourcepoolLease) GetPool() ResourcepoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ResourcepoolLease) GetPoolOk() (*ResourcepoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ResourcepoolLease) SetPool(v ResourcepoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ResourcepoolLease) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *ResourcepoolLease) GetPoolMember() ResourcepoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *ResourcepoolLease) GetPoolMemberOk() (*ResourcepoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *ResourcepoolLease) SetPoolMember(v ResourcepoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *ResourcepoolLease) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *ResourcepoolLease) GetUniverse() ResourcepoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *ResourcepoolLease) GetUniverseOk() (*ResourcepoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *ResourcepoolLease) SetUniverse(v ResourcepoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *ResourcepoolLease) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


