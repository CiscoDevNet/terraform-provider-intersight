# ResourcepoolReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.Reservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.Reservation"]
**Identity** | Pointer to **string** | The type and serial number of the resource for display. | [optional] [readonly] 
**ResourceSerial** | Pointer to **string** | The serial number of the resource. | [optional] 
**ResourceType** | Pointer to **string** | The MO type of the resource being reserved. | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**NullableResourcepoolPoolMemberRelationship**](ResourcepoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**NullableResourcepoolUniverseRelationship**](ResourcepoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewResourcepoolReservation

`func NewResourcepoolReservation(classId string, objectType string, ) *ResourcepoolReservation`

NewResourcepoolReservation instantiates a new ResourcepoolReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolReservationWithDefaults

`func NewResourcepoolReservationWithDefaults() *ResourcepoolReservation`

NewResourcepoolReservationWithDefaults instantiates a new ResourcepoolReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolReservation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolReservation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolReservation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolReservation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolReservation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolReservation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *ResourcepoolReservation) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ResourcepoolReservation) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ResourcepoolReservation) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ResourcepoolReservation) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetResourceSerial

`func (o *ResourcepoolReservation) GetResourceSerial() string`

GetResourceSerial returns the ResourceSerial field if non-nil, zero value otherwise.

### GetResourceSerialOk

`func (o *ResourcepoolReservation) GetResourceSerialOk() (*string, bool)`

GetResourceSerialOk returns a tuple with the ResourceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSerial

`func (o *ResourcepoolReservation) SetResourceSerial(v string)`

SetResourceSerial sets ResourceSerial field to given value.

### HasResourceSerial

`func (o *ResourcepoolReservation) HasResourceSerial() bool`

HasResourceSerial returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourcepoolReservation) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourcepoolReservation) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourcepoolReservation) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourcepoolReservation) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetOrganization

`func (o *ResourcepoolReservation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResourcepoolReservation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResourcepoolReservation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ResourcepoolReservation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ResourcepoolReservation) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ResourcepoolReservation) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetPool

`func (o *ResourcepoolReservation) GetPool() ResourcepoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ResourcepoolReservation) GetPoolOk() (*ResourcepoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ResourcepoolReservation) SetPool(v ResourcepoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ResourcepoolReservation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *ResourcepoolReservation) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *ResourcepoolReservation) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetPoolMember

`func (o *ResourcepoolReservation) GetPoolMember() ResourcepoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *ResourcepoolReservation) GetPoolMemberOk() (*ResourcepoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *ResourcepoolReservation) SetPoolMember(v ResourcepoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *ResourcepoolReservation) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### SetPoolMemberNil

`func (o *ResourcepoolReservation) SetPoolMemberNil(b bool)`

 SetPoolMemberNil sets the value for PoolMember to be an explicit nil

### UnsetPoolMember
`func (o *ResourcepoolReservation) UnsetPoolMember()`

UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
### GetUniverse

`func (o *ResourcepoolReservation) GetUniverse() ResourcepoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *ResourcepoolReservation) GetUniverseOk() (*ResourcepoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *ResourcepoolReservation) SetUniverse(v ResourcepoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *ResourcepoolReservation) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### SetUniverseNil

`func (o *ResourcepoolReservation) SetUniverseNil(b bool)`

 SetUniverseNil sets the value for Universe to be an explicit nil

### UnsetUniverse
`func (o *ResourcepoolReservation) UnsetUniverse()`

UnsetUniverse ensures that no value is present for Universe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


