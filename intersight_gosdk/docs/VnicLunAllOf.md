# VnicLunAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.Lun"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.Lun"]
**Bootable** | Pointer to **bool** | Specifies LUN is bootable. | [optional] 
**LunId** | Pointer to **int64** | The Identifier of the LUN. | [optional] 

## Methods

### NewVnicLunAllOf

`func NewVnicLunAllOf(classId string, objectType string, ) *VnicLunAllOf`

NewVnicLunAllOf instantiates a new VnicLunAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicLunAllOfWithDefaults

`func NewVnicLunAllOfWithDefaults() *VnicLunAllOf`

NewVnicLunAllOfWithDefaults instantiates a new VnicLunAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicLunAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicLunAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicLunAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicLunAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicLunAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicLunAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootable

`func (o *VnicLunAllOf) GetBootable() bool`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *VnicLunAllOf) GetBootableOk() (*bool, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *VnicLunAllOf) SetBootable(v bool)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *VnicLunAllOf) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetLunId

`func (o *VnicLunAllOf) GetLunId() int64`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *VnicLunAllOf) GetLunIdOk() (*int64, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *VnicLunAllOf) SetLunId(v int64)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *VnicLunAllOf) HasLunId() bool`

HasLunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


