# FirmwareExcludeComponentPidListTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.ExcludeComponentPidListType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.ExcludeComponentPidListType"]
**ExcludeLocalDiskList** | Pointer to **[]string** |  | [optional] 
**ExcludeStorageControllerList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFirmwareExcludeComponentPidListTypeAllOf

`func NewFirmwareExcludeComponentPidListTypeAllOf(classId string, objectType string, ) *FirmwareExcludeComponentPidListTypeAllOf`

NewFirmwareExcludeComponentPidListTypeAllOf instantiates a new FirmwareExcludeComponentPidListTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareExcludeComponentPidListTypeAllOfWithDefaults

`func NewFirmwareExcludeComponentPidListTypeAllOfWithDefaults() *FirmwareExcludeComponentPidListTypeAllOf`

NewFirmwareExcludeComponentPidListTypeAllOfWithDefaults instantiates a new FirmwareExcludeComponentPidListTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareExcludeComponentPidListTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareExcludeComponentPidListTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareExcludeComponentPidListTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareExcludeComponentPidListTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareExcludeComponentPidListTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareExcludeComponentPidListTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeLocalDiskList

`func (o *FirmwareExcludeComponentPidListTypeAllOf) GetExcludeLocalDiskList() []string`

GetExcludeLocalDiskList returns the ExcludeLocalDiskList field if non-nil, zero value otherwise.

### GetExcludeLocalDiskListOk

`func (o *FirmwareExcludeComponentPidListTypeAllOf) GetExcludeLocalDiskListOk() (*[]string, bool)`

GetExcludeLocalDiskListOk returns a tuple with the ExcludeLocalDiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLocalDiskList

`func (o *FirmwareExcludeComponentPidListTypeAllOf) SetExcludeLocalDiskList(v []string)`

SetExcludeLocalDiskList sets ExcludeLocalDiskList field to given value.

### HasExcludeLocalDiskList

`func (o *FirmwareExcludeComponentPidListTypeAllOf) HasExcludeLocalDiskList() bool`

HasExcludeLocalDiskList returns a boolean if a field has been set.

### SetExcludeLocalDiskListNil

`func (o *FirmwareExcludeComponentPidListTypeAllOf) SetExcludeLocalDiskListNil(b bool)`

 SetExcludeLocalDiskListNil sets the value for ExcludeLocalDiskList to be an explicit nil

### UnsetExcludeLocalDiskList
`func (o *FirmwareExcludeComponentPidListTypeAllOf) UnsetExcludeLocalDiskList()`

UnsetExcludeLocalDiskList ensures that no value is present for ExcludeLocalDiskList, not even an explicit nil
### GetExcludeStorageControllerList

`func (o *FirmwareExcludeComponentPidListTypeAllOf) GetExcludeStorageControllerList() []string`

GetExcludeStorageControllerList returns the ExcludeStorageControllerList field if non-nil, zero value otherwise.

### GetExcludeStorageControllerListOk

`func (o *FirmwareExcludeComponentPidListTypeAllOf) GetExcludeStorageControllerListOk() (*[]string, bool)`

GetExcludeStorageControllerListOk returns a tuple with the ExcludeStorageControllerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeStorageControllerList

`func (o *FirmwareExcludeComponentPidListTypeAllOf) SetExcludeStorageControllerList(v []string)`

SetExcludeStorageControllerList sets ExcludeStorageControllerList field to given value.

### HasExcludeStorageControllerList

`func (o *FirmwareExcludeComponentPidListTypeAllOf) HasExcludeStorageControllerList() bool`

HasExcludeStorageControllerList returns a boolean if a field has been set.

### SetExcludeStorageControllerListNil

`func (o *FirmwareExcludeComponentPidListTypeAllOf) SetExcludeStorageControllerListNil(b bool)`

 SetExcludeStorageControllerListNil sets the value for ExcludeStorageControllerList to be an explicit nil

### UnsetExcludeStorageControllerList
`func (o *FirmwareExcludeComponentPidListTypeAllOf) UnsetExcludeStorageControllerList()`

UnsetExcludeStorageControllerList ensures that no value is present for ExcludeStorageControllerList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


