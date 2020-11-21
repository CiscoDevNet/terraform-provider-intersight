# OsVirtualDriveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.VirtualDriveResponse"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.VirtualDriveResponse"]
**Bootable** | Pointer to **string** | Bootable field of the Virtual Drive target. | [optional] 
**Id** | Pointer to **string** | Virtual Drive ID to be used as Install Target. | [optional] 
**Name** | Pointer to **string** | The Virtual Drive Name to be used as Install Target. | [optional] 
**StorageControllerSlotId** | Pointer to **string** | The Storage Controller associated to the virtual drive. | [optional] 

## Methods

### NewOsVirtualDriveResponse

`func NewOsVirtualDriveResponse(classId string, objectType string, ) *OsVirtualDriveResponse`

NewOsVirtualDriveResponse instantiates a new OsVirtualDriveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsVirtualDriveResponseWithDefaults

`func NewOsVirtualDriveResponseWithDefaults() *OsVirtualDriveResponse`

NewOsVirtualDriveResponseWithDefaults instantiates a new OsVirtualDriveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsVirtualDriveResponse) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsVirtualDriveResponse) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsVirtualDriveResponse) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsVirtualDriveResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsVirtualDriveResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsVirtualDriveResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootable

`func (o *OsVirtualDriveResponse) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *OsVirtualDriveResponse) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *OsVirtualDriveResponse) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *OsVirtualDriveResponse) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetId

`func (o *OsVirtualDriveResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OsVirtualDriveResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OsVirtualDriveResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OsVirtualDriveResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OsVirtualDriveResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsVirtualDriveResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsVirtualDriveResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsVirtualDriveResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageControllerSlotId

`func (o *OsVirtualDriveResponse) GetStorageControllerSlotId() string`

GetStorageControllerSlotId returns the StorageControllerSlotId field if non-nil, zero value otherwise.

### GetStorageControllerSlotIdOk

`func (o *OsVirtualDriveResponse) GetStorageControllerSlotIdOk() (*string, bool)`

GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllerSlotId

`func (o *OsVirtualDriveResponse) SetStorageControllerSlotId(v string)`

SetStorageControllerSlotId sets StorageControllerSlotId field to given value.

### HasStorageControllerSlotId

`func (o *OsVirtualDriveResponse) HasStorageControllerSlotId() bool`

HasStorageControllerSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


