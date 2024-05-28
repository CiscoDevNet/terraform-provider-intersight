# IppoolShadowBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.ShadowBlock"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.ShadowBlock"]
**IpType** | Pointer to **string** | Type of this IP addresses blocks. * &#x60;IPv4&#x60; - IP V4 address type requested. * &#x60;IPv6&#x60; - IP V6 address type requested. | [optional] [default to "IPv4"]
**IpV4Block** | Pointer to [**IppoolIpV4Block**](IppoolIpV4Block.md) |  | [optional] 
**IpV6Block** | Pointer to [**IppoolIpV6Block**](IppoolIpV6Block.md) |  | [optional] 
**Pool** | Pointer to [**NullableIppoolShadowPoolRelationship**](IppoolShadowPoolRelationship.md) |  | [optional] 
**Reservations** | Pointer to [**[]IppoolReservationRelationship**](IppoolReservationRelationship.md) | An array of relationships to ippoolReservation resources. | [optional] [readonly] 
**Vrf** | Pointer to [**NullableVrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 

## Methods

### NewIppoolShadowBlock

`func NewIppoolShadowBlock(classId string, objectType string, ) *IppoolShadowBlock`

NewIppoolShadowBlock instantiates a new IppoolShadowBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolShadowBlockWithDefaults

`func NewIppoolShadowBlockWithDefaults() *IppoolShadowBlock`

NewIppoolShadowBlockWithDefaults instantiates a new IppoolShadowBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolShadowBlock) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolShadowBlock) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolShadowBlock) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolShadowBlock) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolShadowBlock) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolShadowBlock) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpType

`func (o *IppoolShadowBlock) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IppoolShadowBlock) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IppoolShadowBlock) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IppoolShadowBlock) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetIpV4Block

`func (o *IppoolShadowBlock) GetIpV4Block() IppoolIpV4Block`

GetIpV4Block returns the IpV4Block field if non-nil, zero value otherwise.

### GetIpV4BlockOk

`func (o *IppoolShadowBlock) GetIpV4BlockOk() (*IppoolIpV4Block, bool)`

GetIpV4BlockOk returns a tuple with the IpV4Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Block

`func (o *IppoolShadowBlock) SetIpV4Block(v IppoolIpV4Block)`

SetIpV4Block sets IpV4Block field to given value.

### HasIpV4Block

`func (o *IppoolShadowBlock) HasIpV4Block() bool`

HasIpV4Block returns a boolean if a field has been set.

### GetIpV6Block

`func (o *IppoolShadowBlock) GetIpV6Block() IppoolIpV6Block`

GetIpV6Block returns the IpV6Block field if non-nil, zero value otherwise.

### GetIpV6BlockOk

`func (o *IppoolShadowBlock) GetIpV6BlockOk() (*IppoolIpV6Block, bool)`

GetIpV6BlockOk returns a tuple with the IpV6Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Block

`func (o *IppoolShadowBlock) SetIpV6Block(v IppoolIpV6Block)`

SetIpV6Block sets IpV6Block field to given value.

### HasIpV6Block

`func (o *IppoolShadowBlock) HasIpV6Block() bool`

HasIpV6Block returns a boolean if a field has been set.

### GetPool

`func (o *IppoolShadowBlock) GetPool() IppoolShadowPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolShadowBlock) GetPoolOk() (*IppoolShadowPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolShadowBlock) SetPool(v IppoolShadowPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolShadowBlock) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *IppoolShadowBlock) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *IppoolShadowBlock) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetReservations

`func (o *IppoolShadowBlock) GetReservations() []IppoolReservationRelationship`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *IppoolShadowBlock) GetReservationsOk() (*[]IppoolReservationRelationship, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *IppoolShadowBlock) SetReservations(v []IppoolReservationRelationship)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *IppoolShadowBlock) HasReservations() bool`

HasReservations returns a boolean if a field has been set.

### SetReservationsNil

`func (o *IppoolShadowBlock) SetReservationsNil(b bool)`

 SetReservationsNil sets the value for Reservations to be an explicit nil

### UnsetReservations
`func (o *IppoolShadowBlock) UnsetReservations()`

UnsetReservations ensures that no value is present for Reservations, not even an explicit nil
### GetVrf

`func (o *IppoolShadowBlock) GetVrf() VrfVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IppoolShadowBlock) GetVrfOk() (*VrfVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IppoolShadowBlock) SetVrf(v VrfVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IppoolShadowBlock) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *IppoolShadowBlock) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *IppoolShadowBlock) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


