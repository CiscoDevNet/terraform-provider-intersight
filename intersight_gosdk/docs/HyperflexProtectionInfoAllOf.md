# HyperflexProtectionInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ProtectionInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ProtectionInfo"]
**VmCurrentProtectionInfo** | Pointer to [**NullableHyperflexSnapshotInfoBrief**](HyperflexSnapshotInfoBrief.md) |  | [optional] 
**VmLastSuccessfulProtectionInfo** | Pointer to [**NullableHyperflexSnapshotInfoBrief**](HyperflexSnapshotInfoBrief.md) |  | [optional] 
**VmSpaceUsage** | Pointer to [**NullableHyperflexVmProtectionSpaceUsage**](HyperflexVmProtectionSpaceUsage.md) |  | [optional] 

## Methods

### NewHyperflexProtectionInfoAllOf

`func NewHyperflexProtectionInfoAllOf(classId string, objectType string, ) *HyperflexProtectionInfoAllOf`

NewHyperflexProtectionInfoAllOf instantiates a new HyperflexProtectionInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexProtectionInfoAllOfWithDefaults

`func NewHyperflexProtectionInfoAllOfWithDefaults() *HyperflexProtectionInfoAllOf`

NewHyperflexProtectionInfoAllOfWithDefaults instantiates a new HyperflexProtectionInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexProtectionInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexProtectionInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexProtectionInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexProtectionInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexProtectionInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexProtectionInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVmCurrentProtectionInfo

`func (o *HyperflexProtectionInfoAllOf) GetVmCurrentProtectionInfo() HyperflexSnapshotInfoBrief`

GetVmCurrentProtectionInfo returns the VmCurrentProtectionInfo field if non-nil, zero value otherwise.

### GetVmCurrentProtectionInfoOk

`func (o *HyperflexProtectionInfoAllOf) GetVmCurrentProtectionInfoOk() (*HyperflexSnapshotInfoBrief, bool)`

GetVmCurrentProtectionInfoOk returns a tuple with the VmCurrentProtectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCurrentProtectionInfo

`func (o *HyperflexProtectionInfoAllOf) SetVmCurrentProtectionInfo(v HyperflexSnapshotInfoBrief)`

SetVmCurrentProtectionInfo sets VmCurrentProtectionInfo field to given value.

### HasVmCurrentProtectionInfo

`func (o *HyperflexProtectionInfoAllOf) HasVmCurrentProtectionInfo() bool`

HasVmCurrentProtectionInfo returns a boolean if a field has been set.

### SetVmCurrentProtectionInfoNil

`func (o *HyperflexProtectionInfoAllOf) SetVmCurrentProtectionInfoNil(b bool)`

 SetVmCurrentProtectionInfoNil sets the value for VmCurrentProtectionInfo to be an explicit nil

### UnsetVmCurrentProtectionInfo
`func (o *HyperflexProtectionInfoAllOf) UnsetVmCurrentProtectionInfo()`

UnsetVmCurrentProtectionInfo ensures that no value is present for VmCurrentProtectionInfo, not even an explicit nil
### GetVmLastSuccessfulProtectionInfo

`func (o *HyperflexProtectionInfoAllOf) GetVmLastSuccessfulProtectionInfo() HyperflexSnapshotInfoBrief`

GetVmLastSuccessfulProtectionInfo returns the VmLastSuccessfulProtectionInfo field if non-nil, zero value otherwise.

### GetVmLastSuccessfulProtectionInfoOk

`func (o *HyperflexProtectionInfoAllOf) GetVmLastSuccessfulProtectionInfoOk() (*HyperflexSnapshotInfoBrief, bool)`

GetVmLastSuccessfulProtectionInfoOk returns a tuple with the VmLastSuccessfulProtectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmLastSuccessfulProtectionInfo

`func (o *HyperflexProtectionInfoAllOf) SetVmLastSuccessfulProtectionInfo(v HyperflexSnapshotInfoBrief)`

SetVmLastSuccessfulProtectionInfo sets VmLastSuccessfulProtectionInfo field to given value.

### HasVmLastSuccessfulProtectionInfo

`func (o *HyperflexProtectionInfoAllOf) HasVmLastSuccessfulProtectionInfo() bool`

HasVmLastSuccessfulProtectionInfo returns a boolean if a field has been set.

### SetVmLastSuccessfulProtectionInfoNil

`func (o *HyperflexProtectionInfoAllOf) SetVmLastSuccessfulProtectionInfoNil(b bool)`

 SetVmLastSuccessfulProtectionInfoNil sets the value for VmLastSuccessfulProtectionInfo to be an explicit nil

### UnsetVmLastSuccessfulProtectionInfo
`func (o *HyperflexProtectionInfoAllOf) UnsetVmLastSuccessfulProtectionInfo()`

UnsetVmLastSuccessfulProtectionInfo ensures that no value is present for VmLastSuccessfulProtectionInfo, not even an explicit nil
### GetVmSpaceUsage

`func (o *HyperflexProtectionInfoAllOf) GetVmSpaceUsage() HyperflexVmProtectionSpaceUsage`

GetVmSpaceUsage returns the VmSpaceUsage field if non-nil, zero value otherwise.

### GetVmSpaceUsageOk

`func (o *HyperflexProtectionInfoAllOf) GetVmSpaceUsageOk() (*HyperflexVmProtectionSpaceUsage, bool)`

GetVmSpaceUsageOk returns a tuple with the VmSpaceUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSpaceUsage

`func (o *HyperflexProtectionInfoAllOf) SetVmSpaceUsage(v HyperflexVmProtectionSpaceUsage)`

SetVmSpaceUsage sets VmSpaceUsage field to given value.

### HasVmSpaceUsage

`func (o *HyperflexProtectionInfoAllOf) HasVmSpaceUsage() bool`

HasVmSpaceUsage returns a boolean if a field has been set.

### SetVmSpaceUsageNil

`func (o *HyperflexProtectionInfoAllOf) SetVmSpaceUsageNil(b bool)`

 SetVmSpaceUsageNil sets the value for VmSpaceUsage to be an explicit nil

### UnsetVmSpaceUsage
`func (o *HyperflexProtectionInfoAllOf) UnsetVmSpaceUsage()`

UnsetVmSpaceUsage ensures that no value is present for VmSpaceUsage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


