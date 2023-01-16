# OsFibreChannelTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.FibreChannelTarget"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.FibreChannelTarget"]
**InitiatorWwpn** | Pointer to **string** | The WWPN Address of the underlying fibre channel interface at initator used for SAN boot. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 20:00:D4:C9:3C:35:02:01. | [optional] 
**LunId** | Pointer to **int64** | Logical Unit Number (LUN) of the install target. | [optional] [default to 0]
**TargetWwpn** | Pointer to **string** | The WWPN Address of the underlying fibre channel interface at target used by the SAN boot device. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 51:4F:0C:50:14:1F:AF:01. | [optional] 

## Methods

### NewOsFibreChannelTarget

`func NewOsFibreChannelTarget(classId string, objectType string, ) *OsFibreChannelTarget`

NewOsFibreChannelTarget instantiates a new OsFibreChannelTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsFibreChannelTargetWithDefaults

`func NewOsFibreChannelTargetWithDefaults() *OsFibreChannelTarget`

NewOsFibreChannelTargetWithDefaults instantiates a new OsFibreChannelTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsFibreChannelTarget) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsFibreChannelTarget) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsFibreChannelTarget) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsFibreChannelTarget) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsFibreChannelTarget) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsFibreChannelTarget) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInitiatorWwpn

`func (o *OsFibreChannelTarget) GetInitiatorWwpn() string`

GetInitiatorWwpn returns the InitiatorWwpn field if non-nil, zero value otherwise.

### GetInitiatorWwpnOk

`func (o *OsFibreChannelTarget) GetInitiatorWwpnOk() (*string, bool)`

GetInitiatorWwpnOk returns a tuple with the InitiatorWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorWwpn

`func (o *OsFibreChannelTarget) SetInitiatorWwpn(v string)`

SetInitiatorWwpn sets InitiatorWwpn field to given value.

### HasInitiatorWwpn

`func (o *OsFibreChannelTarget) HasInitiatorWwpn() bool`

HasInitiatorWwpn returns a boolean if a field has been set.

### GetLunId

`func (o *OsFibreChannelTarget) GetLunId() int64`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *OsFibreChannelTarget) GetLunIdOk() (*int64, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *OsFibreChannelTarget) SetLunId(v int64)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *OsFibreChannelTarget) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### GetTargetWwpn

`func (o *OsFibreChannelTarget) GetTargetWwpn() string`

GetTargetWwpn returns the TargetWwpn field if non-nil, zero value otherwise.

### GetTargetWwpnOk

`func (o *OsFibreChannelTarget) GetTargetWwpnOk() (*string, bool)`

GetTargetWwpnOk returns a tuple with the TargetWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetWwpn

`func (o *OsFibreChannelTarget) SetTargetWwpn(v string)`

SetTargetWwpn sets TargetWwpn field to given value.

### HasTargetWwpn

`func (o *OsFibreChannelTarget) HasTargetWwpn() bool`

HasTargetWwpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


