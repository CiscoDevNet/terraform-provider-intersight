# FcpoolFcBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fcpool.FcBlock"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fcpool.FcBlock"]
**IdBlock** | Pointer to [**FcpoolBlock**](FcpoolBlock.md) |  | [optional] 
**Pool** | Pointer to [**NullableFcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 
**Reservations** | Pointer to [**[]FcpoolReservationRelationship**](FcpoolReservationRelationship.md) | An array of relationships to fcpoolReservation resources. | [optional] [readonly] 

## Methods

### NewFcpoolFcBlock

`func NewFcpoolFcBlock(classId string, objectType string, ) *FcpoolFcBlock`

NewFcpoolFcBlock instantiates a new FcpoolFcBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolFcBlockWithDefaults

`func NewFcpoolFcBlockWithDefaults() *FcpoolFcBlock`

NewFcpoolFcBlockWithDefaults instantiates a new FcpoolFcBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolFcBlock) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolFcBlock) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolFcBlock) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolFcBlock) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolFcBlock) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolFcBlock) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdBlock

`func (o *FcpoolFcBlock) GetIdBlock() FcpoolBlock`

GetIdBlock returns the IdBlock field if non-nil, zero value otherwise.

### GetIdBlockOk

`func (o *FcpoolFcBlock) GetIdBlockOk() (*FcpoolBlock, bool)`

GetIdBlockOk returns a tuple with the IdBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdBlock

`func (o *FcpoolFcBlock) SetIdBlock(v FcpoolBlock)`

SetIdBlock sets IdBlock field to given value.

### HasIdBlock

`func (o *FcpoolFcBlock) HasIdBlock() bool`

HasIdBlock returns a boolean if a field has been set.

### GetPool

`func (o *FcpoolFcBlock) GetPool() FcpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *FcpoolFcBlock) GetPoolOk() (*FcpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *FcpoolFcBlock) SetPool(v FcpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *FcpoolFcBlock) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *FcpoolFcBlock) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *FcpoolFcBlock) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetReservations

`func (o *FcpoolFcBlock) GetReservations() []FcpoolReservationRelationship`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *FcpoolFcBlock) GetReservationsOk() (*[]FcpoolReservationRelationship, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *FcpoolFcBlock) SetReservations(v []FcpoolReservationRelationship)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *FcpoolFcBlock) HasReservations() bool`

HasReservations returns a boolean if a field has been set.

### SetReservationsNil

`func (o *FcpoolFcBlock) SetReservationsNil(b bool)`

 SetReservationsNil sets the value for Reservations to be an explicit nil

### UnsetReservations
`func (o *FcpoolFcBlock) UnsetReservations()`

UnsetReservations ensures that no value is present for Reservations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


