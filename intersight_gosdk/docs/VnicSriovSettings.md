# VnicSriovSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.SriovSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.SriovSettings"]
**CompCountPerVf** | Pointer to **int64** | Completion Queue resources per Virtual Function (VF). | [optional] [default to 5]
**Enabled** | Pointer to **bool** | If enabled, sets Single Root Input Output Virtualization (SR-IOV) on this vNIC. | [optional] 
**IntCountPerVf** | Pointer to **int64** | Interrupt Count resources per Virtual Function (VF). | [optional] [default to 8]
**RxCountPerVf** | Pointer to **int64** | Receive Queue resources per Virtual Function (VF). | [optional] [default to 4]
**TxCountPerVf** | Pointer to **int64** | Transmit Queue resources per Virtual Function (VF). | [optional] [default to 1]
**VfCount** | Pointer to **int64** | Number of Virtual Functions (VF) to be created for this vNIC. Valid values are 1 to 64 when SR-IOV is enabled. | [optional] [default to 64]

## Methods

### NewVnicSriovSettings

`func NewVnicSriovSettings(classId string, objectType string, ) *VnicSriovSettings`

NewVnicSriovSettings instantiates a new VnicSriovSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicSriovSettingsWithDefaults

`func NewVnicSriovSettingsWithDefaults() *VnicSriovSettings`

NewVnicSriovSettingsWithDefaults instantiates a new VnicSriovSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicSriovSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicSriovSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicSriovSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicSriovSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicSriovSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicSriovSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompCountPerVf

`func (o *VnicSriovSettings) GetCompCountPerVf() int64`

GetCompCountPerVf returns the CompCountPerVf field if non-nil, zero value otherwise.

### GetCompCountPerVfOk

`func (o *VnicSriovSettings) GetCompCountPerVfOk() (*int64, bool)`

GetCompCountPerVfOk returns a tuple with the CompCountPerVf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompCountPerVf

`func (o *VnicSriovSettings) SetCompCountPerVf(v int64)`

SetCompCountPerVf sets CompCountPerVf field to given value.

### HasCompCountPerVf

`func (o *VnicSriovSettings) HasCompCountPerVf() bool`

HasCompCountPerVf returns a boolean if a field has been set.

### GetEnabled

`func (o *VnicSriovSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VnicSriovSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VnicSriovSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VnicSriovSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntCountPerVf

`func (o *VnicSriovSettings) GetIntCountPerVf() int64`

GetIntCountPerVf returns the IntCountPerVf field if non-nil, zero value otherwise.

### GetIntCountPerVfOk

`func (o *VnicSriovSettings) GetIntCountPerVfOk() (*int64, bool)`

GetIntCountPerVfOk returns a tuple with the IntCountPerVf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntCountPerVf

`func (o *VnicSriovSettings) SetIntCountPerVf(v int64)`

SetIntCountPerVf sets IntCountPerVf field to given value.

### HasIntCountPerVf

`func (o *VnicSriovSettings) HasIntCountPerVf() bool`

HasIntCountPerVf returns a boolean if a field has been set.

### GetRxCountPerVf

`func (o *VnicSriovSettings) GetRxCountPerVf() int64`

GetRxCountPerVf returns the RxCountPerVf field if non-nil, zero value otherwise.

### GetRxCountPerVfOk

`func (o *VnicSriovSettings) GetRxCountPerVfOk() (*int64, bool)`

GetRxCountPerVfOk returns a tuple with the RxCountPerVf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxCountPerVf

`func (o *VnicSriovSettings) SetRxCountPerVf(v int64)`

SetRxCountPerVf sets RxCountPerVf field to given value.

### HasRxCountPerVf

`func (o *VnicSriovSettings) HasRxCountPerVf() bool`

HasRxCountPerVf returns a boolean if a field has been set.

### GetTxCountPerVf

`func (o *VnicSriovSettings) GetTxCountPerVf() int64`

GetTxCountPerVf returns the TxCountPerVf field if non-nil, zero value otherwise.

### GetTxCountPerVfOk

`func (o *VnicSriovSettings) GetTxCountPerVfOk() (*int64, bool)`

GetTxCountPerVfOk returns a tuple with the TxCountPerVf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCountPerVf

`func (o *VnicSriovSettings) SetTxCountPerVf(v int64)`

SetTxCountPerVf sets TxCountPerVf field to given value.

### HasTxCountPerVf

`func (o *VnicSriovSettings) HasTxCountPerVf() bool`

HasTxCountPerVf returns a boolean if a field has been set.

### GetVfCount

`func (o *VnicSriovSettings) GetVfCount() int64`

GetVfCount returns the VfCount field if non-nil, zero value otherwise.

### GetVfCountOk

`func (o *VnicSriovSettings) GetVfCountOk() (*int64, bool)`

GetVfCountOk returns a tuple with the VfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVfCount

`func (o *VnicSriovSettings) SetVfCount(v int64)`

SetVfCount sets VfCount field to given value.

### HasVfCount

`func (o *VnicSriovSettings) HasVfCount() bool`

HasVfCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


