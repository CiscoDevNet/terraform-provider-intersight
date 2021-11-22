# FirmwareIncludeComponentListTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.IncludeComponentListType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.IncludeComponentListType"]
**IncludeLocalDiskList** | Pointer to **[]string** |  | [optional] 
**IncludeStorageControllerList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFirmwareIncludeComponentListTypeAllOf

`func NewFirmwareIncludeComponentListTypeAllOf(classId string, objectType string, ) *FirmwareIncludeComponentListTypeAllOf`

NewFirmwareIncludeComponentListTypeAllOf instantiates a new FirmwareIncludeComponentListTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareIncludeComponentListTypeAllOfWithDefaults

`func NewFirmwareIncludeComponentListTypeAllOfWithDefaults() *FirmwareIncludeComponentListTypeAllOf`

NewFirmwareIncludeComponentListTypeAllOfWithDefaults instantiates a new FirmwareIncludeComponentListTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareIncludeComponentListTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareIncludeComponentListTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareIncludeComponentListTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareIncludeComponentListTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareIncludeComponentListTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareIncludeComponentListTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIncludeLocalDiskList

`func (o *FirmwareIncludeComponentListTypeAllOf) GetIncludeLocalDiskList() []string`

GetIncludeLocalDiskList returns the IncludeLocalDiskList field if non-nil, zero value otherwise.

### GetIncludeLocalDiskListOk

`func (o *FirmwareIncludeComponentListTypeAllOf) GetIncludeLocalDiskListOk() (*[]string, bool)`

GetIncludeLocalDiskListOk returns a tuple with the IncludeLocalDiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLocalDiskList

`func (o *FirmwareIncludeComponentListTypeAllOf) SetIncludeLocalDiskList(v []string)`

SetIncludeLocalDiskList sets IncludeLocalDiskList field to given value.

### HasIncludeLocalDiskList

`func (o *FirmwareIncludeComponentListTypeAllOf) HasIncludeLocalDiskList() bool`

HasIncludeLocalDiskList returns a boolean if a field has been set.

### SetIncludeLocalDiskListNil

`func (o *FirmwareIncludeComponentListTypeAllOf) SetIncludeLocalDiskListNil(b bool)`

 SetIncludeLocalDiskListNil sets the value for IncludeLocalDiskList to be an explicit nil

### UnsetIncludeLocalDiskList
`func (o *FirmwareIncludeComponentListTypeAllOf) UnsetIncludeLocalDiskList()`

UnsetIncludeLocalDiskList ensures that no value is present for IncludeLocalDiskList, not even an explicit nil
### GetIncludeStorageControllerList

`func (o *FirmwareIncludeComponentListTypeAllOf) GetIncludeStorageControllerList() []string`

GetIncludeStorageControllerList returns the IncludeStorageControllerList field if non-nil, zero value otherwise.

### GetIncludeStorageControllerListOk

`func (o *FirmwareIncludeComponentListTypeAllOf) GetIncludeStorageControllerListOk() (*[]string, bool)`

GetIncludeStorageControllerListOk returns a tuple with the IncludeStorageControllerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStorageControllerList

`func (o *FirmwareIncludeComponentListTypeAllOf) SetIncludeStorageControllerList(v []string)`

SetIncludeStorageControllerList sets IncludeStorageControllerList field to given value.

### HasIncludeStorageControllerList

`func (o *FirmwareIncludeComponentListTypeAllOf) HasIncludeStorageControllerList() bool`

HasIncludeStorageControllerList returns a boolean if a field has been set.

### SetIncludeStorageControllerListNil

`func (o *FirmwareIncludeComponentListTypeAllOf) SetIncludeStorageControllerListNil(b bool)`

 SetIncludeStorageControllerListNil sets the value for IncludeStorageControllerList to be an explicit nil

### UnsetIncludeStorageControllerList
`func (o *FirmwareIncludeComponentListTypeAllOf) UnsetIncludeStorageControllerList()`

UnsetIncludeStorageControllerList ensures that no value is present for IncludeStorageControllerList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


