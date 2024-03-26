# IqnpoolReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iqnpool.Reservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iqnpool.Reservation"]
**Identity** | Pointer to **string** | IQN identity to be reserved. | [optional] 
**IqnNumber** | Pointer to **int64** | Number of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnPrefix** | Pointer to **string** | Prefix of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnSuffix** | Pointer to **string** | Suffix of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**Block** | Pointer to [**IqnpoolBlockRelationship**](IqnpoolBlockRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Pool** | Pointer to [**IqnpoolPoolRelationship**](IqnpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**IqnpoolPoolMemberRelationship**](IqnpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**IqnpoolUniverseRelationship**](IqnpoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewIqnpoolReservation

`func NewIqnpoolReservation(classId string, objectType string, ) *IqnpoolReservation`

NewIqnpoolReservation instantiates a new IqnpoolReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqnpoolReservationWithDefaults

`func NewIqnpoolReservationWithDefaults() *IqnpoolReservation`

NewIqnpoolReservationWithDefaults instantiates a new IqnpoolReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IqnpoolReservation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IqnpoolReservation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IqnpoolReservation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IqnpoolReservation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IqnpoolReservation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IqnpoolReservation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *IqnpoolReservation) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *IqnpoolReservation) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *IqnpoolReservation) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *IqnpoolReservation) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIqnNumber

`func (o *IqnpoolReservation) GetIqnNumber() int64`

GetIqnNumber returns the IqnNumber field if non-nil, zero value otherwise.

### GetIqnNumberOk

`func (o *IqnpoolReservation) GetIqnNumberOk() (*int64, bool)`

GetIqnNumberOk returns a tuple with the IqnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnNumber

`func (o *IqnpoolReservation) SetIqnNumber(v int64)`

SetIqnNumber sets IqnNumber field to given value.

### HasIqnNumber

`func (o *IqnpoolReservation) HasIqnNumber() bool`

HasIqnNumber returns a boolean if a field has been set.

### GetIqnPrefix

`func (o *IqnpoolReservation) GetIqnPrefix() string`

GetIqnPrefix returns the IqnPrefix field if non-nil, zero value otherwise.

### GetIqnPrefixOk

`func (o *IqnpoolReservation) GetIqnPrefixOk() (*string, bool)`

GetIqnPrefixOk returns a tuple with the IqnPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnPrefix

`func (o *IqnpoolReservation) SetIqnPrefix(v string)`

SetIqnPrefix sets IqnPrefix field to given value.

### HasIqnPrefix

`func (o *IqnpoolReservation) HasIqnPrefix() bool`

HasIqnPrefix returns a boolean if a field has been set.

### GetIqnSuffix

`func (o *IqnpoolReservation) GetIqnSuffix() string`

GetIqnSuffix returns the IqnSuffix field if non-nil, zero value otherwise.

### GetIqnSuffixOk

`func (o *IqnpoolReservation) GetIqnSuffixOk() (*string, bool)`

GetIqnSuffixOk returns a tuple with the IqnSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnSuffix

`func (o *IqnpoolReservation) SetIqnSuffix(v string)`

SetIqnSuffix sets IqnSuffix field to given value.

### HasIqnSuffix

`func (o *IqnpoolReservation) HasIqnSuffix() bool`

HasIqnSuffix returns a boolean if a field has been set.

### GetBlock

`func (o *IqnpoolReservation) GetBlock() IqnpoolBlockRelationship`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *IqnpoolReservation) GetBlockOk() (*IqnpoolBlockRelationship, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *IqnpoolReservation) SetBlock(v IqnpoolBlockRelationship)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *IqnpoolReservation) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetOrganization

`func (o *IqnpoolReservation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IqnpoolReservation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IqnpoolReservation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IqnpoolReservation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPool

`func (o *IqnpoolReservation) GetPool() IqnpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IqnpoolReservation) GetPoolOk() (*IqnpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IqnpoolReservation) SetPool(v IqnpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IqnpoolReservation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *IqnpoolReservation) GetPoolMember() IqnpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *IqnpoolReservation) GetPoolMemberOk() (*IqnpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *IqnpoolReservation) SetPoolMember(v IqnpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *IqnpoolReservation) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *IqnpoolReservation) GetUniverse() IqnpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *IqnpoolReservation) GetUniverseOk() (*IqnpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *IqnpoolReservation) SetUniverse(v IqnpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *IqnpoolReservation) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


