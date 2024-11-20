# IqnpoolLease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iqnpool.Lease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iqnpool.Lease"]
**IqnAddress** | Pointer to **string** | IQN address allocated for pool-based allocation. It is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] 
**IqnNumber** | Pointer to **int64** | Number of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnPrefix** | Pointer to **string** | Prefix of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnSuffix** | Pointer to **string** | Suffix of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**Reservation** | Pointer to [**NullablePoolReservationReference**](PoolReservationReference.md) | The reference to the reservation object. | [optional] 
**AssignedToEntity** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableIqnpoolPoolRelationship**](IqnpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**NullableIqnpoolPoolMemberRelationship**](IqnpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**NullableIqnpoolUniverseRelationship**](IqnpoolUniverseRelationship.md) |  | [optional] 

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

### GetIqnNumber

`func (o *IqnpoolLease) GetIqnNumber() int64`

GetIqnNumber returns the IqnNumber field if non-nil, zero value otherwise.

### GetIqnNumberOk

`func (o *IqnpoolLease) GetIqnNumberOk() (*int64, bool)`

GetIqnNumberOk returns a tuple with the IqnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnNumber

`func (o *IqnpoolLease) SetIqnNumber(v int64)`

SetIqnNumber sets IqnNumber field to given value.

### HasIqnNumber

`func (o *IqnpoolLease) HasIqnNumber() bool`

HasIqnNumber returns a boolean if a field has been set.

### GetIqnPrefix

`func (o *IqnpoolLease) GetIqnPrefix() string`

GetIqnPrefix returns the IqnPrefix field if non-nil, zero value otherwise.

### GetIqnPrefixOk

`func (o *IqnpoolLease) GetIqnPrefixOk() (*string, bool)`

GetIqnPrefixOk returns a tuple with the IqnPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnPrefix

`func (o *IqnpoolLease) SetIqnPrefix(v string)`

SetIqnPrefix sets IqnPrefix field to given value.

### HasIqnPrefix

`func (o *IqnpoolLease) HasIqnPrefix() bool`

HasIqnPrefix returns a boolean if a field has been set.

### GetIqnSuffix

`func (o *IqnpoolLease) GetIqnSuffix() string`

GetIqnSuffix returns the IqnSuffix field if non-nil, zero value otherwise.

### GetIqnSuffixOk

`func (o *IqnpoolLease) GetIqnSuffixOk() (*string, bool)`

GetIqnSuffixOk returns a tuple with the IqnSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnSuffix

`func (o *IqnpoolLease) SetIqnSuffix(v string)`

SetIqnSuffix sets IqnSuffix field to given value.

### HasIqnSuffix

`func (o *IqnpoolLease) HasIqnSuffix() bool`

HasIqnSuffix returns a boolean if a field has been set.

### GetReservation

`func (o *IqnpoolLease) GetReservation() PoolReservationReference`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *IqnpoolLease) GetReservationOk() (*PoolReservationReference, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *IqnpoolLease) SetReservation(v PoolReservationReference)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *IqnpoolLease) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### SetReservationNil

`func (o *IqnpoolLease) SetReservationNil(b bool)`

 SetReservationNil sets the value for Reservation to be an explicit nil

### UnsetReservation
`func (o *IqnpoolLease) UnsetReservation()`

UnsetReservation ensures that no value is present for Reservation, not even an explicit nil
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

### SetAssignedToEntityNil

`func (o *IqnpoolLease) SetAssignedToEntityNil(b bool)`

 SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil

### UnsetAssignedToEntity
`func (o *IqnpoolLease) UnsetAssignedToEntity()`

UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
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

### SetPoolNil

`func (o *IqnpoolLease) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *IqnpoolLease) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
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

### SetPoolMemberNil

`func (o *IqnpoolLease) SetPoolMemberNil(b bool)`

 SetPoolMemberNil sets the value for PoolMember to be an explicit nil

### UnsetPoolMember
`func (o *IqnpoolLease) UnsetPoolMember()`

UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
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

### SetUniverseNil

`func (o *IqnpoolLease) SetUniverseNil(b bool)`

 SetUniverseNil sets the value for Universe to be an explicit nil

### UnsetUniverse
`func (o *IqnpoolLease) UnsetUniverse()`

UnsetUniverse ensures that no value is present for Universe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


