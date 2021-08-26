# StorageManualDriveGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.ManualDriveGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.ManualDriveGroup"]
**DedicatedHotSpares** | Pointer to **string** | A collection of drives to be used as hot spares for this Drive Group. | [optional] 
**SpanGroups** | Pointer to [**[]StorageSpanDrives**](StorageSpanDrives.md) |  | [optional] 

## Methods

### NewStorageManualDriveGroupAllOf

`func NewStorageManualDriveGroupAllOf(classId string, objectType string, ) *StorageManualDriveGroupAllOf`

NewStorageManualDriveGroupAllOf instantiates a new StorageManualDriveGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageManualDriveGroupAllOfWithDefaults

`func NewStorageManualDriveGroupAllOfWithDefaults() *StorageManualDriveGroupAllOf`

NewStorageManualDriveGroupAllOfWithDefaults instantiates a new StorageManualDriveGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageManualDriveGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageManualDriveGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageManualDriveGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageManualDriveGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageManualDriveGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageManualDriveGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDedicatedHotSpares

`func (o *StorageManualDriveGroupAllOf) GetDedicatedHotSpares() string`

GetDedicatedHotSpares returns the DedicatedHotSpares field if non-nil, zero value otherwise.

### GetDedicatedHotSparesOk

`func (o *StorageManualDriveGroupAllOf) GetDedicatedHotSparesOk() (*string, bool)`

GetDedicatedHotSparesOk returns a tuple with the DedicatedHotSpares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedHotSpares

`func (o *StorageManualDriveGroupAllOf) SetDedicatedHotSpares(v string)`

SetDedicatedHotSpares sets DedicatedHotSpares field to given value.

### HasDedicatedHotSpares

`func (o *StorageManualDriveGroupAllOf) HasDedicatedHotSpares() bool`

HasDedicatedHotSpares returns a boolean if a field has been set.

### GetSpanGroups

`func (o *StorageManualDriveGroupAllOf) GetSpanGroups() []StorageSpanDrives`

GetSpanGroups returns the SpanGroups field if non-nil, zero value otherwise.

### GetSpanGroupsOk

`func (o *StorageManualDriveGroupAllOf) GetSpanGroupsOk() (*[]StorageSpanDrives, bool)`

GetSpanGroupsOk returns a tuple with the SpanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanGroups

`func (o *StorageManualDriveGroupAllOf) SetSpanGroups(v []StorageSpanDrives)`

SetSpanGroups sets SpanGroups field to given value.

### HasSpanGroups

`func (o *StorageManualDriveGroupAllOf) HasSpanGroups() bool`

HasSpanGroups returns a boolean if a field has been set.

### SetSpanGroupsNil

`func (o *StorageManualDriveGroupAllOf) SetSpanGroupsNil(b bool)`

 SetSpanGroupsNil sets the value for SpanGroups to be an explicit nil

### UnsetSpanGroups
`func (o *StorageManualDriveGroupAllOf) UnsetSpanGroups()`

UnsetSpanGroups ensures that no value is present for SpanGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


