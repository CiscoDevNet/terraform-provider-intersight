# HciExternalStorageArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ExternalStorageArray"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ExternalStorageArray"]
**ArrayExtId** | Pointer to **string** | The external identifier of the Pure Storage FlashArray. | [optional] [readonly] 
**ArrayIpAddress** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 

## Methods

### NewHciExternalStorageArray

`func NewHciExternalStorageArray(classId string, objectType string, ) *HciExternalStorageArray`

NewHciExternalStorageArray instantiates a new HciExternalStorageArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciExternalStorageArrayWithDefaults

`func NewHciExternalStorageArrayWithDefaults() *HciExternalStorageArray`

NewHciExternalStorageArrayWithDefaults instantiates a new HciExternalStorageArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciExternalStorageArray) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciExternalStorageArray) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciExternalStorageArray) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciExternalStorageArray) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciExternalStorageArray) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciExternalStorageArray) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArrayExtId

`func (o *HciExternalStorageArray) GetArrayExtId() string`

GetArrayExtId returns the ArrayExtId field if non-nil, zero value otherwise.

### GetArrayExtIdOk

`func (o *HciExternalStorageArray) GetArrayExtIdOk() (*string, bool)`

GetArrayExtIdOk returns a tuple with the ArrayExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayExtId

`func (o *HciExternalStorageArray) SetArrayExtId(v string)`

SetArrayExtId sets ArrayExtId field to given value.

### HasArrayExtId

`func (o *HciExternalStorageArray) HasArrayExtId() bool`

HasArrayExtId returns a boolean if a field has been set.

### GetArrayIpAddress

`func (o *HciExternalStorageArray) GetArrayIpAddress() HciIpAddress`

GetArrayIpAddress returns the ArrayIpAddress field if non-nil, zero value otherwise.

### GetArrayIpAddressOk

`func (o *HciExternalStorageArray) GetArrayIpAddressOk() (*HciIpAddress, bool)`

GetArrayIpAddressOk returns a tuple with the ArrayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayIpAddress

`func (o *HciExternalStorageArray) SetArrayIpAddress(v HciIpAddress)`

SetArrayIpAddress sets ArrayIpAddress field to given value.

### HasArrayIpAddress

`func (o *HciExternalStorageArray) HasArrayIpAddress() bool`

HasArrayIpAddress returns a boolean if a field has been set.

### SetArrayIpAddressNil

`func (o *HciExternalStorageArray) SetArrayIpAddressNil(b bool)`

 SetArrayIpAddressNil sets the value for ArrayIpAddress to be an explicit nil

### UnsetArrayIpAddress
`func (o *HciExternalStorageArray) UnsetArrayIpAddress()`

UnsetArrayIpAddress ensures that no value is present for ArrayIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


