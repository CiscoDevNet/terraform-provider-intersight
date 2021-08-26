# HyperflexSnapshotInfoBriefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SnapshotInfoBrief"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SnapshotInfoBrief"]
**ReplicationStatus** | Pointer to [**NullableHyperflexReplicationStatus**](HyperflexReplicationStatus.md) |  | [optional] 
**Site** | Pointer to **string** | Cluster site for this snapshot. * &#x60;PRIMARY&#x60; - The cluster site for this backup is primary. * &#x60;SECONDARY&#x60; - The cluster site for this backup is secondary. | [optional] [readonly] [default to "PRIMARY"]
**SnapshotStatus** | Pointer to [**NullableHyperflexSnapshotStatus**](HyperflexSnapshotStatus.md) |  | [optional] 
**VmSnapshotEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 

## Methods

### NewHyperflexSnapshotInfoBriefAllOf

`func NewHyperflexSnapshotInfoBriefAllOf(classId string, objectType string, ) *HyperflexSnapshotInfoBriefAllOf`

NewHyperflexSnapshotInfoBriefAllOf instantiates a new HyperflexSnapshotInfoBriefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSnapshotInfoBriefAllOfWithDefaults

`func NewHyperflexSnapshotInfoBriefAllOfWithDefaults() *HyperflexSnapshotInfoBriefAllOf`

NewHyperflexSnapshotInfoBriefAllOfWithDefaults instantiates a new HyperflexSnapshotInfoBriefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSnapshotInfoBriefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSnapshotInfoBriefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSnapshotInfoBriefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSnapshotInfoBriefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSnapshotInfoBriefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSnapshotInfoBriefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReplicationStatus

`func (o *HyperflexSnapshotInfoBriefAllOf) GetReplicationStatus() HyperflexReplicationStatus`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *HyperflexSnapshotInfoBriefAllOf) GetReplicationStatusOk() (*HyperflexReplicationStatus, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *HyperflexSnapshotInfoBriefAllOf) SetReplicationStatus(v HyperflexReplicationStatus)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *HyperflexSnapshotInfoBriefAllOf) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### SetReplicationStatusNil

`func (o *HyperflexSnapshotInfoBriefAllOf) SetReplicationStatusNil(b bool)`

 SetReplicationStatusNil sets the value for ReplicationStatus to be an explicit nil

### UnsetReplicationStatus
`func (o *HyperflexSnapshotInfoBriefAllOf) UnsetReplicationStatus()`

UnsetReplicationStatus ensures that no value is present for ReplicationStatus, not even an explicit nil
### GetSite

`func (o *HyperflexSnapshotInfoBriefAllOf) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *HyperflexSnapshotInfoBriefAllOf) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *HyperflexSnapshotInfoBriefAllOf) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *HyperflexSnapshotInfoBriefAllOf) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSnapshotStatus

`func (o *HyperflexSnapshotInfoBriefAllOf) GetSnapshotStatus() HyperflexSnapshotStatus`

GetSnapshotStatus returns the SnapshotStatus field if non-nil, zero value otherwise.

### GetSnapshotStatusOk

`func (o *HyperflexSnapshotInfoBriefAllOf) GetSnapshotStatusOk() (*HyperflexSnapshotStatus, bool)`

GetSnapshotStatusOk returns a tuple with the SnapshotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotStatus

`func (o *HyperflexSnapshotInfoBriefAllOf) SetSnapshotStatus(v HyperflexSnapshotStatus)`

SetSnapshotStatus sets SnapshotStatus field to given value.

### HasSnapshotStatus

`func (o *HyperflexSnapshotInfoBriefAllOf) HasSnapshotStatus() bool`

HasSnapshotStatus returns a boolean if a field has been set.

### SetSnapshotStatusNil

`func (o *HyperflexSnapshotInfoBriefAllOf) SetSnapshotStatusNil(b bool)`

 SetSnapshotStatusNil sets the value for SnapshotStatus to be an explicit nil

### UnsetSnapshotStatus
`func (o *HyperflexSnapshotInfoBriefAllOf) UnsetSnapshotStatus()`

UnsetSnapshotStatus ensures that no value is present for SnapshotStatus, not even an explicit nil
### GetVmSnapshotEntityReference

`func (o *HyperflexSnapshotInfoBriefAllOf) GetVmSnapshotEntityReference() HyperflexEntityReference`

GetVmSnapshotEntityReference returns the VmSnapshotEntityReference field if non-nil, zero value otherwise.

### GetVmSnapshotEntityReferenceOk

`func (o *HyperflexSnapshotInfoBriefAllOf) GetVmSnapshotEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetVmSnapshotEntityReferenceOk returns a tuple with the VmSnapshotEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSnapshotEntityReference

`func (o *HyperflexSnapshotInfoBriefAllOf) SetVmSnapshotEntityReference(v HyperflexEntityReference)`

SetVmSnapshotEntityReference sets VmSnapshotEntityReference field to given value.

### HasVmSnapshotEntityReference

`func (o *HyperflexSnapshotInfoBriefAllOf) HasVmSnapshotEntityReference() bool`

HasVmSnapshotEntityReference returns a boolean if a field has been set.

### SetVmSnapshotEntityReferenceNil

`func (o *HyperflexSnapshotInfoBriefAllOf) SetVmSnapshotEntityReferenceNil(b bool)`

 SetVmSnapshotEntityReferenceNil sets the value for VmSnapshotEntityReference to be an explicit nil

### UnsetVmSnapshotEntityReference
`func (o *HyperflexSnapshotInfoBriefAllOf) UnsetVmSnapshotEntityReference()`

UnsetVmSnapshotEntityReference ensures that no value is present for VmSnapshotEntityReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


