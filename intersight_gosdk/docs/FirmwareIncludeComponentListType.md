# FirmwareIncludeComponentListType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.IncludeComponentListType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.IncludeComponentListType"]
**IncludeLocalDiskList** | Pointer to **[]string** |  | [optional] 
**IncludeStorageControllerList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFirmwareIncludeComponentListType

`func NewFirmwareIncludeComponentListType(classId string, objectType string, ) *FirmwareIncludeComponentListType`

NewFirmwareIncludeComponentListType instantiates a new FirmwareIncludeComponentListType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareIncludeComponentListTypeWithDefaults

`func NewFirmwareIncludeComponentListTypeWithDefaults() *FirmwareIncludeComponentListType`

NewFirmwareIncludeComponentListTypeWithDefaults instantiates a new FirmwareIncludeComponentListType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareIncludeComponentListType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareIncludeComponentListType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareIncludeComponentListType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareIncludeComponentListType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareIncludeComponentListType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareIncludeComponentListType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIncludeLocalDiskList

`func (o *FirmwareIncludeComponentListType) GetIncludeLocalDiskList() []string`

GetIncludeLocalDiskList returns the IncludeLocalDiskList field if non-nil, zero value otherwise.

### GetIncludeLocalDiskListOk

`func (o *FirmwareIncludeComponentListType) GetIncludeLocalDiskListOk() (*[]string, bool)`

GetIncludeLocalDiskListOk returns a tuple with the IncludeLocalDiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLocalDiskList

`func (o *FirmwareIncludeComponentListType) SetIncludeLocalDiskList(v []string)`

SetIncludeLocalDiskList sets IncludeLocalDiskList field to given value.

### HasIncludeLocalDiskList

`func (o *FirmwareIncludeComponentListType) HasIncludeLocalDiskList() bool`

HasIncludeLocalDiskList returns a boolean if a field has been set.

### SetIncludeLocalDiskListNil

`func (o *FirmwareIncludeComponentListType) SetIncludeLocalDiskListNil(b bool)`

 SetIncludeLocalDiskListNil sets the value for IncludeLocalDiskList to be an explicit nil

### UnsetIncludeLocalDiskList
`func (o *FirmwareIncludeComponentListType) UnsetIncludeLocalDiskList()`

UnsetIncludeLocalDiskList ensures that no value is present for IncludeLocalDiskList, not even an explicit nil
### GetIncludeStorageControllerList

`func (o *FirmwareIncludeComponentListType) GetIncludeStorageControllerList() []string`

GetIncludeStorageControllerList returns the IncludeStorageControllerList field if non-nil, zero value otherwise.

### GetIncludeStorageControllerListOk

`func (o *FirmwareIncludeComponentListType) GetIncludeStorageControllerListOk() (*[]string, bool)`

GetIncludeStorageControllerListOk returns a tuple with the IncludeStorageControllerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStorageControllerList

`func (o *FirmwareIncludeComponentListType) SetIncludeStorageControllerList(v []string)`

SetIncludeStorageControllerList sets IncludeStorageControllerList field to given value.

### HasIncludeStorageControllerList

`func (o *FirmwareIncludeComponentListType) HasIncludeStorageControllerList() bool`

HasIncludeStorageControllerList returns a boolean if a field has been set.

### SetIncludeStorageControllerListNil

`func (o *FirmwareIncludeComponentListType) SetIncludeStorageControllerListNil(b bool)`

 SetIncludeStorageControllerListNil sets the value for IncludeStorageControllerList to be an explicit nil

### UnsetIncludeStorageControllerList
`func (o *FirmwareIncludeComponentListType) UnsetIncludeStorageControllerList()`

UnsetIncludeStorageControllerList ensures that no value is present for IncludeStorageControllerList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


