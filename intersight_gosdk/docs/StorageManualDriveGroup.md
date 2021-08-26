# StorageManualDriveGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.ManualDriveGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.ManualDriveGroup"]
**DedicatedHotSpares** | Pointer to **string** | A collection of drives to be used as hot spares for this Drive Group. | [optional] 
**SpanGroups** | Pointer to [**[]StorageSpanDrives**](StorageSpanDrives.md) |  | [optional] 

## Methods

### NewStorageManualDriveGroup

`func NewStorageManualDriveGroup(classId string, objectType string, ) *StorageManualDriveGroup`

NewStorageManualDriveGroup instantiates a new StorageManualDriveGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageManualDriveGroupWithDefaults

`func NewStorageManualDriveGroupWithDefaults() *StorageManualDriveGroup`

NewStorageManualDriveGroupWithDefaults instantiates a new StorageManualDriveGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageManualDriveGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageManualDriveGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageManualDriveGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageManualDriveGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageManualDriveGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageManualDriveGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDedicatedHotSpares

`func (o *StorageManualDriveGroup) GetDedicatedHotSpares() string`

GetDedicatedHotSpares returns the DedicatedHotSpares field if non-nil, zero value otherwise.

### GetDedicatedHotSparesOk

`func (o *StorageManualDriveGroup) GetDedicatedHotSparesOk() (*string, bool)`

GetDedicatedHotSparesOk returns a tuple with the DedicatedHotSpares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedHotSpares

`func (o *StorageManualDriveGroup) SetDedicatedHotSpares(v string)`

SetDedicatedHotSpares sets DedicatedHotSpares field to given value.

### HasDedicatedHotSpares

`func (o *StorageManualDriveGroup) HasDedicatedHotSpares() bool`

HasDedicatedHotSpares returns a boolean if a field has been set.

### GetSpanGroups

`func (o *StorageManualDriveGroup) GetSpanGroups() []StorageSpanDrives`

GetSpanGroups returns the SpanGroups field if non-nil, zero value otherwise.

### GetSpanGroupsOk

`func (o *StorageManualDriveGroup) GetSpanGroupsOk() (*[]StorageSpanDrives, bool)`

GetSpanGroupsOk returns a tuple with the SpanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanGroups

`func (o *StorageManualDriveGroup) SetSpanGroups(v []StorageSpanDrives)`

SetSpanGroups sets SpanGroups field to given value.

### HasSpanGroups

`func (o *StorageManualDriveGroup) HasSpanGroups() bool`

HasSpanGroups returns a boolean if a field has been set.

### SetSpanGroupsNil

`func (o *StorageManualDriveGroup) SetSpanGroupsNil(b bool)`

 SetSpanGroupsNil sets the value for SpanGroups to be an explicit nil

### UnsetSpanGroups
`func (o *StorageManualDriveGroup) UnsetSpanGroups()`

UnsetSpanGroups ensures that no value is present for SpanGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


