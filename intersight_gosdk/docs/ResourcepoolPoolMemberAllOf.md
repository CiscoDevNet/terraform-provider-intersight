# ResourcepoolPoolMemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.PoolMember"]
**Features** | Pointer to **[]string** |  | [optional] 
**AssignedToEntity** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] 
**Peer** | Pointer to [**ResourcepoolLeaseRelationship**](resourcepool.Lease.Relationship.md) |  | [optional] 
**Pool** | Pointer to [**ResourcepoolPoolRelationship**](resourcepool.Pool.Relationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 

## Methods

### NewResourcepoolPoolMemberAllOf

`func NewResourcepoolPoolMemberAllOf(classId string, objectType string, ) *ResourcepoolPoolMemberAllOf`

NewResourcepoolPoolMemberAllOf instantiates a new ResourcepoolPoolMemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolPoolMemberAllOfWithDefaults

`func NewResourcepoolPoolMemberAllOfWithDefaults() *ResourcepoolPoolMemberAllOf`

NewResourcepoolPoolMemberAllOfWithDefaults instantiates a new ResourcepoolPoolMemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolPoolMemberAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolPoolMemberAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolPoolMemberAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolPoolMemberAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolPoolMemberAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolPoolMemberAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeatures

`func (o *ResourcepoolPoolMemberAllOf) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ResourcepoolPoolMemberAllOf) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ResourcepoolPoolMemberAllOf) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ResourcepoolPoolMemberAllOf) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *ResourcepoolPoolMemberAllOf) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *ResourcepoolPoolMemberAllOf) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetAssignedToEntity

`func (o *ResourcepoolPoolMemberAllOf) GetAssignedToEntity() []MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *ResourcepoolPoolMemberAllOf) GetAssignedToEntityOk() (*[]MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *ResourcepoolPoolMemberAllOf) SetAssignedToEntity(v []MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *ResourcepoolPoolMemberAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### SetAssignedToEntityNil

`func (o *ResourcepoolPoolMemberAllOf) SetAssignedToEntityNil(b bool)`

 SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil

### UnsetAssignedToEntity
`func (o *ResourcepoolPoolMemberAllOf) UnsetAssignedToEntity()`

UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
### GetPeer

`func (o *ResourcepoolPoolMemberAllOf) GetPeer() ResourcepoolLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *ResourcepoolPoolMemberAllOf) GetPeerOk() (*ResourcepoolLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *ResourcepoolPoolMemberAllOf) SetPeer(v ResourcepoolLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *ResourcepoolPoolMemberAllOf) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPool

`func (o *ResourcepoolPoolMemberAllOf) GetPool() ResourcepoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ResourcepoolPoolMemberAllOf) GetPoolOk() (*ResourcepoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ResourcepoolPoolMemberAllOf) SetPool(v ResourcepoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ResourcepoolPoolMemberAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetResource

`func (o *ResourcepoolPoolMemberAllOf) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourcepoolPoolMemberAllOf) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourcepoolPoolMemberAllOf) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourcepoolPoolMemberAllOf) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


