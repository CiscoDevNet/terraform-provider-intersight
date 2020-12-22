# VnicUsnicSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.UsnicSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.UsnicSettings"]
**Cos** | Pointer to **int64** | Class of Service to be used for traffic on the usNIC. | [optional] [default to 5]
**Count** | Pointer to **int64** | Number of usNIC interfaces to be created. | [optional] 
**UsnicAdapterPolicy** | Pointer to **string** | Ethernet Adapter policy to be associated with the usNICs. | [optional] 

## Methods

### NewVnicUsnicSettingsAllOf

`func NewVnicUsnicSettingsAllOf(classId string, objectType string, ) *VnicUsnicSettingsAllOf`

NewVnicUsnicSettingsAllOf instantiates a new VnicUsnicSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicUsnicSettingsAllOfWithDefaults

`func NewVnicUsnicSettingsAllOfWithDefaults() *VnicUsnicSettingsAllOf`

NewVnicUsnicSettingsAllOfWithDefaults instantiates a new VnicUsnicSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicUsnicSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicUsnicSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicUsnicSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicUsnicSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicUsnicSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicUsnicSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCos

`func (o *VnicUsnicSettingsAllOf) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *VnicUsnicSettingsAllOf) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *VnicUsnicSettingsAllOf) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *VnicUsnicSettingsAllOf) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetCount

`func (o *VnicUsnicSettingsAllOf) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VnicUsnicSettingsAllOf) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VnicUsnicSettingsAllOf) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VnicUsnicSettingsAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetUsnicAdapterPolicy

`func (o *VnicUsnicSettingsAllOf) GetUsnicAdapterPolicy() string`

GetUsnicAdapterPolicy returns the UsnicAdapterPolicy field if non-nil, zero value otherwise.

### GetUsnicAdapterPolicyOk

`func (o *VnicUsnicSettingsAllOf) GetUsnicAdapterPolicyOk() (*string, bool)`

GetUsnicAdapterPolicyOk returns a tuple with the UsnicAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicAdapterPolicy

`func (o *VnicUsnicSettingsAllOf) SetUsnicAdapterPolicy(v string)`

SetUsnicAdapterPolicy sets UsnicAdapterPolicy field to given value.

### HasUsnicAdapterPolicy

`func (o *VnicUsnicSettingsAllOf) HasUsnicAdapterPolicy() bool`

HasUsnicAdapterPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


