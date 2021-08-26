# IqnpoolLease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iqnpool.Lease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iqnpool.Lease"]
**IqnAddress** | Pointer to **string** | IQN address allocated for pool-based allocation \&quot;prefix+suffix+number\&quot;. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**Pool** | Pointer to [**IqnpoolPoolRelationship**](IqnpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**IqnpoolPoolMemberRelationship**](IqnpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**IqnpoolUniverseRelationship**](IqnpoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewIqnpoolLease

`func NewIqnpoolLease(classId string, objectType string, ) *IqnpoolLease`

NewIqnpoolLease instantiates a new IqnpoolLease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqnpoolLeaseWithDefaults

`func NewIqnpoolLeaseWithDefaults() *IqnpoolLease`

NewIqnpoolLeaseWithDefaults instantiates a new IqnpoolLease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IqnpoolLease) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IqnpoolLease) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IqnpoolLease) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IqnpoolLease) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IqnpoolLease) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IqnpoolLease) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIqnAddress

`func (o *IqnpoolLease) GetIqnAddress() string`

GetIqnAddress returns the IqnAddress field if non-nil, zero value otherwise.

### GetIqnAddressOk

`func (o *IqnpoolLease) GetIqnAddressOk() (*string, bool)`

GetIqnAddressOk returns a tuple with the IqnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnAddress

`func (o *IqnpoolLease) SetIqnAddress(v string)`

SetIqnAddress sets IqnAddress field to given value.

### HasIqnAddress

`func (o *IqnpoolLease) HasIqnAddress() bool`

HasIqnAddress returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IqnpoolLease) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IqnpoolLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IqnpoolLease) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IqnpoolLease) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetPool

`func (o *IqnpoolLease) GetPool() IqnpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IqnpoolLease) GetPoolOk() (*IqnpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IqnpoolLease) SetPool(v IqnpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IqnpoolLease) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *IqnpoolLease) GetPoolMember() IqnpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *IqnpoolLease) GetPoolMemberOk() (*IqnpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *IqnpoolLease) SetPoolMember(v IqnpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *IqnpoolLease) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *IqnpoolLease) GetUniverse() IqnpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *IqnpoolLease) GetUniverseOk() (*IqnpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *IqnpoolLease) SetUniverse(v IqnpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *IqnpoolLease) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


