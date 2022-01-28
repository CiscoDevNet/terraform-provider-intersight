# FirmwareExcludeComponentPidListType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.ExcludeComponentPidListType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.ExcludeComponentPidListType"]
**ExcludeLocalDiskList** | Pointer to **[]string** |  | [optional] 
**ExcludeStorageControllerList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFirmwareExcludeComponentPidListType

`func NewFirmwareExcludeComponentPidListType(classId string, objectType string, ) *FirmwareExcludeComponentPidListType`

NewFirmwareExcludeComponentPidListType instantiates a new FirmwareExcludeComponentPidListType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareExcludeComponentPidListTypeWithDefaults

`func NewFirmwareExcludeComponentPidListTypeWithDefaults() *FirmwareExcludeComponentPidListType`

NewFirmwareExcludeComponentPidListTypeWithDefaults instantiates a new FirmwareExcludeComponentPidListType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareExcludeComponentPidListType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareExcludeComponentPidListType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareExcludeComponentPidListType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareExcludeComponentPidListType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareExcludeComponentPidListType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareExcludeComponentPidListType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeLocalDiskList

`func (o *FirmwareExcludeComponentPidListType) GetExcludeLocalDiskList() []string`

GetExcludeLocalDiskList returns the ExcludeLocalDiskList field if non-nil, zero value otherwise.

### GetExcludeLocalDiskListOk

`func (o *FirmwareExcludeComponentPidListType) GetExcludeLocalDiskListOk() (*[]string, bool)`

GetExcludeLocalDiskListOk returns a tuple with the ExcludeLocalDiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLocalDiskList

`func (o *FirmwareExcludeComponentPidListType) SetExcludeLocalDiskList(v []string)`

SetExcludeLocalDiskList sets ExcludeLocalDiskList field to given value.

### HasExcludeLocalDiskList

`func (o *FirmwareExcludeComponentPidListType) HasExcludeLocalDiskList() bool`

HasExcludeLocalDiskList returns a boolean if a field has been set.

### SetExcludeLocalDiskListNil

`func (o *FirmwareExcludeComponentPidListType) SetExcludeLocalDiskListNil(b bool)`

 SetExcludeLocalDiskListNil sets the value for ExcludeLocalDiskList to be an explicit nil

### UnsetExcludeLocalDiskList
`func (o *FirmwareExcludeComponentPidListType) UnsetExcludeLocalDiskList()`

UnsetExcludeLocalDiskList ensures that no value is present for ExcludeLocalDiskList, not even an explicit nil
### GetExcludeStorageControllerList

`func (o *FirmwareExcludeComponentPidListType) GetExcludeStorageControllerList() []string`

GetExcludeStorageControllerList returns the ExcludeStorageControllerList field if non-nil, zero value otherwise.

### GetExcludeStorageControllerListOk

`func (o *FirmwareExcludeComponentPidListType) GetExcludeStorageControllerListOk() (*[]string, bool)`

GetExcludeStorageControllerListOk returns a tuple with the ExcludeStorageControllerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeStorageControllerList

`func (o *FirmwareExcludeComponentPidListType) SetExcludeStorageControllerList(v []string)`

SetExcludeStorageControllerList sets ExcludeStorageControllerList field to given value.

### HasExcludeStorageControllerList

`func (o *FirmwareExcludeComponentPidListType) HasExcludeStorageControllerList() bool`

HasExcludeStorageControllerList returns a boolean if a field has been set.

### SetExcludeStorageControllerListNil

`func (o *FirmwareExcludeComponentPidListType) SetExcludeStorageControllerListNil(b bool)`

 SetExcludeStorageControllerListNil sets the value for ExcludeStorageControllerList to be an explicit nil

### UnsetExcludeStorageControllerList
`func (o *FirmwareExcludeComponentPidListType) UnsetExcludeStorageControllerList()`

UnsetExcludeStorageControllerList ensures that no value is present for ExcludeStorageControllerList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


