# OsIscsiTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.IscsiTarget"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.IscsiTarget"]
**LunId** | Pointer to **int64** | The Logical Unit Number (LUN) of the install target. | [optional] [default to 0]
**TargetIqn** | Pointer to **string** | IQN (iSCSI qualified name) of Storage iSCSI target.Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique_name. | [optional] 
**VnicMac** | Pointer to **string** | MAC address of the VNIC used as iSCSI initiator interface. | [optional] 

## Methods

### NewOsIscsiTarget

`func NewOsIscsiTarget(classId string, objectType string, ) *OsIscsiTarget`

NewOsIscsiTarget instantiates a new OsIscsiTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsIscsiTargetWithDefaults

`func NewOsIscsiTargetWithDefaults() *OsIscsiTarget`

NewOsIscsiTargetWithDefaults instantiates a new OsIscsiTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsIscsiTarget) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsIscsiTarget) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsIscsiTarget) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsIscsiTarget) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsIscsiTarget) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsIscsiTarget) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLunId

`func (o *OsIscsiTarget) GetLunId() int64`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *OsIscsiTarget) GetLunIdOk() (*int64, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *OsIscsiTarget) SetLunId(v int64)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *OsIscsiTarget) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### GetTargetIqn

`func (o *OsIscsiTarget) GetTargetIqn() string`

GetTargetIqn returns the TargetIqn field if non-nil, zero value otherwise.

### GetTargetIqnOk

`func (o *OsIscsiTarget) GetTargetIqnOk() (*string, bool)`

GetTargetIqnOk returns a tuple with the TargetIqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIqn

`func (o *OsIscsiTarget) SetTargetIqn(v string)`

SetTargetIqn sets TargetIqn field to given value.

### HasTargetIqn

`func (o *OsIscsiTarget) HasTargetIqn() bool`

HasTargetIqn returns a boolean if a field has been set.

### GetVnicMac

`func (o *OsIscsiTarget) GetVnicMac() string`

GetVnicMac returns the VnicMac field if non-nil, zero value otherwise.

### GetVnicMacOk

`func (o *OsIscsiTarget) GetVnicMacOk() (*string, bool)`

GetVnicMacOk returns a tuple with the VnicMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicMac

`func (o *OsIscsiTarget) SetVnicMac(v string)`

SetVnicMac sets VnicMac field to given value.

### HasVnicMac

`func (o *OsIscsiTarget) HasVnicMac() bool`

HasVnicMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


