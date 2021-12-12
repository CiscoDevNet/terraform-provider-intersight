# NiatelemetryVniStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.VniStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.VniStatus"]
**Vni** | Pointer to **string** | Returns the vni id of the vni. | [optional] 
**VniState** | Pointer to **string** | Returns the vni state of the vni. | [optional] 
**VniType** | Pointer to **string** | Returns the vni type of the vni. | [optional] 

## Methods

### NewNiatelemetryVniStatusAllOf

`func NewNiatelemetryVniStatusAllOf(classId string, objectType string, ) *NiatelemetryVniStatusAllOf`

NewNiatelemetryVniStatusAllOf instantiates a new NiatelemetryVniStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryVniStatusAllOfWithDefaults

`func NewNiatelemetryVniStatusAllOfWithDefaults() *NiatelemetryVniStatusAllOf`

NewNiatelemetryVniStatusAllOfWithDefaults instantiates a new NiatelemetryVniStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryVniStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryVniStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryVniStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryVniStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryVniStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryVniStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVni

`func (o *NiatelemetryVniStatusAllOf) GetVni() string`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *NiatelemetryVniStatusAllOf) GetVniOk() (*string, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *NiatelemetryVniStatusAllOf) SetVni(v string)`

SetVni sets Vni field to given value.

### HasVni

`func (o *NiatelemetryVniStatusAllOf) HasVni() bool`

HasVni returns a boolean if a field has been set.

### GetVniState

`func (o *NiatelemetryVniStatusAllOf) GetVniState() string`

GetVniState returns the VniState field if non-nil, zero value otherwise.

### GetVniStateOk

`func (o *NiatelemetryVniStatusAllOf) GetVniStateOk() (*string, bool)`

GetVniStateOk returns a tuple with the VniState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniState

`func (o *NiatelemetryVniStatusAllOf) SetVniState(v string)`

SetVniState sets VniState field to given value.

### HasVniState

`func (o *NiatelemetryVniStatusAllOf) HasVniState() bool`

HasVniState returns a boolean if a field has been set.

### GetVniType

`func (o *NiatelemetryVniStatusAllOf) GetVniType() string`

GetVniType returns the VniType field if non-nil, zero value otherwise.

### GetVniTypeOk

`func (o *NiatelemetryVniStatusAllOf) GetVniTypeOk() (*string, bool)`

GetVniTypeOk returns a tuple with the VniType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniType

`func (o *NiatelemetryVniStatusAllOf) SetVniType(v string)`

SetVniType sets VniType field to given value.

### HasVniType

`func (o *NiatelemetryVniStatusAllOf) HasVniType() bool`

HasVniType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


