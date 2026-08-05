# ResourcepoolReservationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.ReservationReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.ReservationReference"]
**PoolMoid** | Pointer to **string** | The moid of the pool object, if applicable. | [optional] 
**ResourceSerial** | Pointer to **string** | The serial number of the resource that is being reserved. | [optional] 
**ResourceType** | Pointer to **string** | The resource type that is being reserved. | [optional] 

## Methods

### NewResourcepoolReservationReference

`func NewResourcepoolReservationReference(classId string, objectType string, ) *ResourcepoolReservationReference`

NewResourcepoolReservationReference instantiates a new ResourcepoolReservationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolReservationReferenceWithDefaults

`func NewResourcepoolReservationReferenceWithDefaults() *ResourcepoolReservationReference`

NewResourcepoolReservationReferenceWithDefaults instantiates a new ResourcepoolReservationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolReservationReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolReservationReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolReservationReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolReservationReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolReservationReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolReservationReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPoolMoid

`func (o *ResourcepoolReservationReference) GetPoolMoid() string`

GetPoolMoid returns the PoolMoid field if non-nil, zero value otherwise.

### GetPoolMoidOk

`func (o *ResourcepoolReservationReference) GetPoolMoidOk() (*string, bool)`

GetPoolMoidOk returns a tuple with the PoolMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMoid

`func (o *ResourcepoolReservationReference) SetPoolMoid(v string)`

SetPoolMoid sets PoolMoid field to given value.

### HasPoolMoid

`func (o *ResourcepoolReservationReference) HasPoolMoid() bool`

HasPoolMoid returns a boolean if a field has been set.

### GetResourceSerial

`func (o *ResourcepoolReservationReference) GetResourceSerial() string`

GetResourceSerial returns the ResourceSerial field if non-nil, zero value otherwise.

### GetResourceSerialOk

`func (o *ResourcepoolReservationReference) GetResourceSerialOk() (*string, bool)`

GetResourceSerialOk returns a tuple with the ResourceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSerial

`func (o *ResourcepoolReservationReference) SetResourceSerial(v string)`

SetResourceSerial sets ResourceSerial field to given value.

### HasResourceSerial

`func (o *ResourcepoolReservationReference) HasResourceSerial() bool`

HasResourceSerial returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourcepoolReservationReference) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourcepoolReservationReference) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourcepoolReservationReference) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourcepoolReservationReference) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


