# VnicIscsiAuthProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.IscsiAuthProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.IscsiAuthProfile"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | Password of Initiator/Target Interface. Enter between 12 and 16 characters, including special characters except spaces, tabs, line breaks. | [optional] 
**UserId** | Pointer to **string** | User Id of Initiator/Target Interface. Enter between 1 and 128 characters, spaces, or special characters. | [optional] 

## Methods

### NewVnicIscsiAuthProfile

`func NewVnicIscsiAuthProfile(classId string, objectType string, ) *VnicIscsiAuthProfile`

NewVnicIscsiAuthProfile instantiates a new VnicIscsiAuthProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicIscsiAuthProfileWithDefaults

`func NewVnicIscsiAuthProfileWithDefaults() *VnicIscsiAuthProfile`

NewVnicIscsiAuthProfileWithDefaults instantiates a new VnicIscsiAuthProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicIscsiAuthProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicIscsiAuthProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicIscsiAuthProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicIscsiAuthProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicIscsiAuthProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicIscsiAuthProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *VnicIscsiAuthProfile) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *VnicIscsiAuthProfile) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *VnicIscsiAuthProfile) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *VnicIscsiAuthProfile) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *VnicIscsiAuthProfile) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VnicIscsiAuthProfile) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VnicIscsiAuthProfile) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VnicIscsiAuthProfile) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserId

`func (o *VnicIscsiAuthProfile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VnicIscsiAuthProfile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VnicIscsiAuthProfile) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *VnicIscsiAuthProfile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


