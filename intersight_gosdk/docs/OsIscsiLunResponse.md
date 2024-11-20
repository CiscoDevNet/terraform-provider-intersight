# OsIscsiLunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.IscsiLunResponse"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.IscsiLunResponse"]
**LunId** | Pointer to **int64** | The Logical Unit Number (LUN) of the install target. | [optional] [default to 0]
**TargetIqn** | Pointer to **string** | IQN (iSCSI qualified name) of Storage iSCSI target. Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique_name. | [optional] 
**VnicMac** | Pointer to **string** | MAC address of the VNIC to be used as initiator iSCSI interface. | [optional] 

## Methods

### NewOsIscsiLunResponse

`func NewOsIscsiLunResponse(classId string, objectType string, ) *OsIscsiLunResponse`

NewOsIscsiLunResponse instantiates a new OsIscsiLunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsIscsiLunResponseWithDefaults

`func NewOsIscsiLunResponseWithDefaults() *OsIscsiLunResponse`

NewOsIscsiLunResponseWithDefaults instantiates a new OsIscsiLunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsIscsiLunResponse) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsIscsiLunResponse) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsIscsiLunResponse) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsIscsiLunResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsIscsiLunResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsIscsiLunResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLunId

`func (o *OsIscsiLunResponse) GetLunId() int64`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *OsIscsiLunResponse) GetLunIdOk() (*int64, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *OsIscsiLunResponse) SetLunId(v int64)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *OsIscsiLunResponse) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### GetTargetIqn

`func (o *OsIscsiLunResponse) GetTargetIqn() string`

GetTargetIqn returns the TargetIqn field if non-nil, zero value otherwise.

### GetTargetIqnOk

`func (o *OsIscsiLunResponse) GetTargetIqnOk() (*string, bool)`

GetTargetIqnOk returns a tuple with the TargetIqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIqn

`func (o *OsIscsiLunResponse) SetTargetIqn(v string)`

SetTargetIqn sets TargetIqn field to given value.

### HasTargetIqn

`func (o *OsIscsiLunResponse) HasTargetIqn() bool`

HasTargetIqn returns a boolean if a field has been set.

### GetVnicMac

`func (o *OsIscsiLunResponse) GetVnicMac() string`

GetVnicMac returns the VnicMac field if non-nil, zero value otherwise.

### GetVnicMacOk

`func (o *OsIscsiLunResponse) GetVnicMacOk() (*string, bool)`

GetVnicMacOk returns a tuple with the VnicMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicMac

`func (o *OsIscsiLunResponse) SetVnicMac(v string)`

SetVnicMac sets VnicMac field to given value.

### HasVnicMac

`func (o *OsIscsiLunResponse) HasVnicMac() bool`

HasVnicMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


