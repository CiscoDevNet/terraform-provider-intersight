# VnicUsnicSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.UsnicSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.UsnicSettings"]
**Cos** | Pointer to **int64** | Class of Service to be used for traffic on the usNIC. | [optional] [default to 5]
**Count** | Pointer to **int64** | Number of usNIC interfaces to be created. | [optional] 
**UsnicAdapterPolicy** | Pointer to **string** | Ethernet Adapter policy to be associated with the usNICs. | [optional] 

## Methods

### NewVnicUsnicSettings

`func NewVnicUsnicSettings(classId string, objectType string, ) *VnicUsnicSettings`

NewVnicUsnicSettings instantiates a new VnicUsnicSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicUsnicSettingsWithDefaults

`func NewVnicUsnicSettingsWithDefaults() *VnicUsnicSettings`

NewVnicUsnicSettingsWithDefaults instantiates a new VnicUsnicSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicUsnicSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicUsnicSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicUsnicSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicUsnicSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicUsnicSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicUsnicSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCos

`func (o *VnicUsnicSettings) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *VnicUsnicSettings) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *VnicUsnicSettings) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *VnicUsnicSettings) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetCount

`func (o *VnicUsnicSettings) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VnicUsnicSettings) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VnicUsnicSettings) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VnicUsnicSettings) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetUsnicAdapterPolicy

`func (o *VnicUsnicSettings) GetUsnicAdapterPolicy() string`

GetUsnicAdapterPolicy returns the UsnicAdapterPolicy field if non-nil, zero value otherwise.

### GetUsnicAdapterPolicyOk

`func (o *VnicUsnicSettings) GetUsnicAdapterPolicyOk() (*string, bool)`

GetUsnicAdapterPolicyOk returns a tuple with the UsnicAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicAdapterPolicy

`func (o *VnicUsnicSettings) SetUsnicAdapterPolicy(v string)`

SetUsnicAdapterPolicy sets UsnicAdapterPolicy field to given value.

### HasUsnicAdapterPolicy

`func (o *VnicUsnicSettings) HasUsnicAdapterPolicy() bool`

HasUsnicAdapterPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


