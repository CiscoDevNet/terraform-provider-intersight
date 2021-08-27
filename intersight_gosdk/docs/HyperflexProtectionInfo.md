# HyperflexProtectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ProtectionInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ProtectionInfo"]
**VmCurrentProtectionInfo** | Pointer to [**NullableHyperflexSnapshotInfoBrief**](HyperflexSnapshotInfoBrief.md) |  | [optional] 
**VmLastSuccessfulProtectionInfo** | Pointer to [**NullableHyperflexSnapshotInfoBrief**](HyperflexSnapshotInfoBrief.md) |  | [optional] 
**VmSpaceUsage** | Pointer to [**NullableHyperflexVmProtectionSpaceUsage**](HyperflexVmProtectionSpaceUsage.md) |  | [optional] 

## Methods

### NewHyperflexProtectionInfo

`func NewHyperflexProtectionInfo(classId string, objectType string, ) *HyperflexProtectionInfo`

NewHyperflexProtectionInfo instantiates a new HyperflexProtectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexProtectionInfoWithDefaults

`func NewHyperflexProtectionInfoWithDefaults() *HyperflexProtectionInfo`

NewHyperflexProtectionInfoWithDefaults instantiates a new HyperflexProtectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexProtectionInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexProtectionInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexProtectionInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexProtectionInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexProtectionInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexProtectionInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVmCurrentProtectionInfo

`func (o *HyperflexProtectionInfo) GetVmCurrentProtectionInfo() HyperflexSnapshotInfoBrief`

GetVmCurrentProtectionInfo returns the VmCurrentProtectionInfo field if non-nil, zero value otherwise.

### GetVmCurrentProtectionInfoOk

`func (o *HyperflexProtectionInfo) GetVmCurrentProtectionInfoOk() (*HyperflexSnapshotInfoBrief, bool)`

GetVmCurrentProtectionInfoOk returns a tuple with the VmCurrentProtectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCurrentProtectionInfo

`func (o *HyperflexProtectionInfo) SetVmCurrentProtectionInfo(v HyperflexSnapshotInfoBrief)`

SetVmCurrentProtectionInfo sets VmCurrentProtectionInfo field to given value.

### HasVmCurrentProtectionInfo

`func (o *HyperflexProtectionInfo) HasVmCurrentProtectionInfo() bool`

HasVmCurrentProtectionInfo returns a boolean if a field has been set.

### SetVmCurrentProtectionInfoNil

`func (o *HyperflexProtectionInfo) SetVmCurrentProtectionInfoNil(b bool)`

 SetVmCurrentProtectionInfoNil sets the value for VmCurrentProtectionInfo to be an explicit nil

### UnsetVmCurrentProtectionInfo
`func (o *HyperflexProtectionInfo) UnsetVmCurrentProtectionInfo()`

UnsetVmCurrentProtectionInfo ensures that no value is present for VmCurrentProtectionInfo, not even an explicit nil
### GetVmLastSuccessfulProtectionInfo

`func (o *HyperflexProtectionInfo) GetVmLastSuccessfulProtectionInfo() HyperflexSnapshotInfoBrief`

GetVmLastSuccessfulProtectionInfo returns the VmLastSuccessfulProtectionInfo field if non-nil, zero value otherwise.

### GetVmLastSuccessfulProtectionInfoOk

`func (o *HyperflexProtectionInfo) GetVmLastSuccessfulProtectionInfoOk() (*HyperflexSnapshotInfoBrief, bool)`

GetVmLastSuccessfulProtectionInfoOk returns a tuple with the VmLastSuccessfulProtectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmLastSuccessfulProtectionInfo

`func (o *HyperflexProtectionInfo) SetVmLastSuccessfulProtectionInfo(v HyperflexSnapshotInfoBrief)`

SetVmLastSuccessfulProtectionInfo sets VmLastSuccessfulProtectionInfo field to given value.

### HasVmLastSuccessfulProtectionInfo

`func (o *HyperflexProtectionInfo) HasVmLastSuccessfulProtectionInfo() bool`

HasVmLastSuccessfulProtectionInfo returns a boolean if a field has been set.

### SetVmLastSuccessfulProtectionInfoNil

`func (o *HyperflexProtectionInfo) SetVmLastSuccessfulProtectionInfoNil(b bool)`

 SetVmLastSuccessfulProtectionInfoNil sets the value for VmLastSuccessfulProtectionInfo to be an explicit nil

### UnsetVmLastSuccessfulProtectionInfo
`func (o *HyperflexProtectionInfo) UnsetVmLastSuccessfulProtectionInfo()`

UnsetVmLastSuccessfulProtectionInfo ensures that no value is present for VmLastSuccessfulProtectionInfo, not even an explicit nil
### GetVmSpaceUsage

`func (o *HyperflexProtectionInfo) GetVmSpaceUsage() HyperflexVmProtectionSpaceUsage`

GetVmSpaceUsage returns the VmSpaceUsage field if non-nil, zero value otherwise.

### GetVmSpaceUsageOk

`func (o *HyperflexProtectionInfo) GetVmSpaceUsageOk() (*HyperflexVmProtectionSpaceUsage, bool)`

GetVmSpaceUsageOk returns a tuple with the VmSpaceUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSpaceUsage

`func (o *HyperflexProtectionInfo) SetVmSpaceUsage(v HyperflexVmProtectionSpaceUsage)`

SetVmSpaceUsage sets VmSpaceUsage field to given value.

### HasVmSpaceUsage

`func (o *HyperflexProtectionInfo) HasVmSpaceUsage() bool`

HasVmSpaceUsage returns a boolean if a field has been set.

### SetVmSpaceUsageNil

`func (o *HyperflexProtectionInfo) SetVmSpaceUsageNil(b bool)`

 SetVmSpaceUsageNil sets the value for VmSpaceUsage to be an explicit nil

### UnsetVmSpaceUsage
`func (o *HyperflexProtectionInfo) UnsetVmSpaceUsage()`

UnsetVmSpaceUsage ensures that no value is present for VmSpaceUsage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


