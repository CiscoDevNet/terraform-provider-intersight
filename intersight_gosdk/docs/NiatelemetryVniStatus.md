# NiatelemetryVniStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.VniStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.VniStatus"]
**Vni** | Pointer to **string** | Returns the vni id of the vni. | [optional] 
**VniState** | Pointer to **string** | Returns the vni state of the vni. | [optional] 
**VniType** | Pointer to **string** | Returns the vni type of the vni. | [optional] 

## Methods

### NewNiatelemetryVniStatus

`func NewNiatelemetryVniStatus(classId string, objectType string, ) *NiatelemetryVniStatus`

NewNiatelemetryVniStatus instantiates a new NiatelemetryVniStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryVniStatusWithDefaults

`func NewNiatelemetryVniStatusWithDefaults() *NiatelemetryVniStatus`

NewNiatelemetryVniStatusWithDefaults instantiates a new NiatelemetryVniStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryVniStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryVniStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryVniStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryVniStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryVniStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryVniStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVni

`func (o *NiatelemetryVniStatus) GetVni() string`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *NiatelemetryVniStatus) GetVniOk() (*string, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *NiatelemetryVniStatus) SetVni(v string)`

SetVni sets Vni field to given value.

### HasVni

`func (o *NiatelemetryVniStatus) HasVni() bool`

HasVni returns a boolean if a field has been set.

### GetVniState

`func (o *NiatelemetryVniStatus) GetVniState() string`

GetVniState returns the VniState field if non-nil, zero value otherwise.

### GetVniStateOk

`func (o *NiatelemetryVniStatus) GetVniStateOk() (*string, bool)`

GetVniStateOk returns a tuple with the VniState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniState

`func (o *NiatelemetryVniStatus) SetVniState(v string)`

SetVniState sets VniState field to given value.

### HasVniState

`func (o *NiatelemetryVniStatus) HasVniState() bool`

HasVniState returns a boolean if a field has been set.

### GetVniType

`func (o *NiatelemetryVniStatus) GetVniType() string`

GetVniType returns the VniType field if non-nil, zero value otherwise.

### GetVniTypeOk

`func (o *NiatelemetryVniStatus) GetVniTypeOk() (*string, bool)`

GetVniTypeOk returns a tuple with the VniType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniType

`func (o *NiatelemetryVniStatus) SetVniType(v string)`

SetVniType sets VniType field to given value.

### HasVniType

`func (o *NiatelemetryVniStatus) HasVniType() bool`

HasVniType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


