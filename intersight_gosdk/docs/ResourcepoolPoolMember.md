# ResourcepoolPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.PoolMember"]
**Features** | Pointer to **[]string** |  | [optional] 
**QualificationType** | Pointer to [**NullableResourcepoolQualificationType**](ResourcepoolQualificationType.md) |  | [optional] 
**AssignedToEntity** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Peer** | Pointer to [**NullableResourcepoolLeaseRelationship**](ResourcepoolLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) |  | [optional] 
**Resource** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewResourcepoolPoolMember

`func NewResourcepoolPoolMember(classId string, objectType string, ) *ResourcepoolPoolMember`

NewResourcepoolPoolMember instantiates a new ResourcepoolPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolPoolMemberWithDefaults

`func NewResourcepoolPoolMemberWithDefaults() *ResourcepoolPoolMember`

NewResourcepoolPoolMemberWithDefaults instantiates a new ResourcepoolPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolPoolMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolPoolMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolPoolMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolPoolMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolPoolMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolPoolMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeatures

`func (o *ResourcepoolPoolMember) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ResourcepoolPoolMember) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ResourcepoolPoolMember) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ResourcepoolPoolMember) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *ResourcepoolPoolMember) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *ResourcepoolPoolMember) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetQualificationType

`func (o *ResourcepoolPoolMember) GetQualificationType() ResourcepoolQualificationType`

GetQualificationType returns the QualificationType field if non-nil, zero value otherwise.

### GetQualificationTypeOk

`func (o *ResourcepoolPoolMember) GetQualificationTypeOk() (*ResourcepoolQualificationType, bool)`

GetQualificationTypeOk returns a tuple with the QualificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualificationType

`func (o *ResourcepoolPoolMember) SetQualificationType(v ResourcepoolQualificationType)`

SetQualificationType sets QualificationType field to given value.

### HasQualificationType

`func (o *ResourcepoolPoolMember) HasQualificationType() bool`

HasQualificationType returns a boolean if a field has been set.

### SetQualificationTypeNil

`func (o *ResourcepoolPoolMember) SetQualificationTypeNil(b bool)`

 SetQualificationTypeNil sets the value for QualificationType to be an explicit nil

### UnsetQualificationType
`func (o *ResourcepoolPoolMember) UnsetQualificationType()`

UnsetQualificationType ensures that no value is present for QualificationType, not even an explicit nil
### GetAssignedToEntity

`func (o *ResourcepoolPoolMember) GetAssignedToEntity() []MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *ResourcepoolPoolMember) GetAssignedToEntityOk() (*[]MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *ResourcepoolPoolMember) SetAssignedToEntity(v []MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *ResourcepoolPoolMember) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### SetAssignedToEntityNil

`func (o *ResourcepoolPoolMember) SetAssignedToEntityNil(b bool)`

 SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil

### UnsetAssignedToEntity
`func (o *ResourcepoolPoolMember) UnsetAssignedToEntity()`

UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
### GetPeer

`func (o *ResourcepoolPoolMember) GetPeer() ResourcepoolLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *ResourcepoolPoolMember) GetPeerOk() (*ResourcepoolLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *ResourcepoolPoolMember) SetPeer(v ResourcepoolLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *ResourcepoolPoolMember) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### SetPeerNil

`func (o *ResourcepoolPoolMember) SetPeerNil(b bool)`

 SetPeerNil sets the value for Peer to be an explicit nil

### UnsetPeer
`func (o *ResourcepoolPoolMember) UnsetPeer()`

UnsetPeer ensures that no value is present for Peer, not even an explicit nil
### GetPool

`func (o *ResourcepoolPoolMember) GetPool() ResourcepoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ResourcepoolPoolMember) GetPoolOk() (*ResourcepoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ResourcepoolPoolMember) SetPool(v ResourcepoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ResourcepoolPoolMember) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *ResourcepoolPoolMember) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *ResourcepoolPoolMember) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetResource

`func (o *ResourcepoolPoolMember) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourcepoolPoolMember) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourcepoolPoolMember) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourcepoolPoolMember) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ResourcepoolPoolMember) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ResourcepoolPoolMember) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


