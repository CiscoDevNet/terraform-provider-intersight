# IqnpoolReservationAllOf

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

### NewIqnpoolReservationAllOf

`func NewIqnpoolReservationAllOf(classId string, objectType string, ) *IqnpoolReservationAllOf`

NewIqnpoolReservationAllOf instantiates a new IqnpoolReservationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqnpoolReservationAllOfWithDefaults

`func NewIqnpoolReservationAllOfWithDefaults() *IqnpoolReservationAllOf`

NewIqnpoolReservationAllOfWithDefaults instantiates a new IqnpoolReservationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IqnpoolReservationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IqnpoolReservationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IqnpoolReservationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IqnpoolReservationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IqnpoolReservationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IqnpoolReservationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *IqnpoolReservationAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *IqnpoolReservationAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *IqnpoolReservationAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *IqnpoolReservationAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIqnNumber

`func (o *IqnpoolReservationAllOf) GetIqnNumber() int64`

GetIqnNumber returns the IqnNumber field if non-nil, zero value otherwise.

### GetIqnNumberOk

`func (o *IqnpoolReservationAllOf) GetIqnNumberOk() (*int64, bool)`

GetIqnNumberOk returns a tuple with the IqnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnNumber

`func (o *IqnpoolReservationAllOf) SetIqnNumber(v int64)`

SetIqnNumber sets IqnNumber field to given value.

### HasIqnNumber

`func (o *IqnpoolReservationAllOf) HasIqnNumber() bool`

HasIqnNumber returns a boolean if a field has been set.

### GetIqnPrefix

`func (o *IqnpoolReservationAllOf) GetIqnPrefix() string`

GetIqnPrefix returns the IqnPrefix field if non-nil, zero value otherwise.

### GetIqnPrefixOk

`func (o *IqnpoolReservationAllOf) GetIqnPrefixOk() (*string, bool)`

GetIqnPrefixOk returns a tuple with the IqnPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnPrefix

`func (o *IqnpoolReservationAllOf) SetIqnPrefix(v string)`

SetIqnPrefix sets IqnPrefix field to given value.

### HasIqnPrefix

`func (o *IqnpoolReservationAllOf) HasIqnPrefix() bool`

HasIqnPrefix returns a boolean if a field has been set.

### GetIqnSuffix

`func (o *IqnpoolReservationAllOf) GetIqnSuffix() string`

GetIqnSuffix returns the IqnSuffix field if non-nil, zero value otherwise.

### GetIqnSuffixOk

`func (o *IqnpoolReservationAllOf) GetIqnSuffixOk() (*string, bool)`

GetIqnSuffixOk returns a tuple with the IqnSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnSuffix

`func (o *IqnpoolReservationAllOf) SetIqnSuffix(v string)`

SetIqnSuffix sets IqnSuffix field to given value.

### HasIqnSuffix

`func (o *IqnpoolReservationAllOf) HasIqnSuffix() bool`

HasIqnSuffix returns a boolean if a field has been set.

### GetBlock

`func (o *IqnpoolReservationAllOf) GetBlock() IqnpoolBlockRelationship`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *IqnpoolReservationAllOf) GetBlockOk() (*IqnpoolBlockRelationship, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *IqnpoolReservationAllOf) SetBlock(v IqnpoolBlockRelationship)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *IqnpoolReservationAllOf) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetOrganization

`func (o *IqnpoolReservationAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IqnpoolReservationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IqnpoolReservationAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IqnpoolReservationAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPool

`func (o *IqnpoolReservationAllOf) GetPool() IqnpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IqnpoolReservationAllOf) GetPoolOk() (*IqnpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IqnpoolReservationAllOf) SetPool(v IqnpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IqnpoolReservationAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *IqnpoolReservationAllOf) GetPoolMember() IqnpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *IqnpoolReservationAllOf) GetPoolMemberOk() (*IqnpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *IqnpoolReservationAllOf) SetPoolMember(v IqnpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *IqnpoolReservationAllOf) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *IqnpoolReservationAllOf) GetUniverse() IqnpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *IqnpoolReservationAllOf) GetUniverseOk() (*IqnpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *IqnpoolReservationAllOf) SetUniverse(v IqnpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *IqnpoolReservationAllOf) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


