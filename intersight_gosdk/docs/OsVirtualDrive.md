# OsVirtualDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.VirtualDrive"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.VirtualDrive"]
**Id** | Pointer to **string** | Virtual Drive ID to be used as Install Target. | [optional] 
**Name** | Pointer to **string** | The Virtual Drive Name to be used as Install Target. | [optional] 
**StorageControllerSlotId** | Pointer to **string** | The SlotID of the Storage Controller associated to the virtual drive. | [optional] 

## Methods

### NewOsVirtualDrive

`func NewOsVirtualDrive(classId string, objectType string, ) *OsVirtualDrive`

NewOsVirtualDrive instantiates a new OsVirtualDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsVirtualDriveWithDefaults

`func NewOsVirtualDriveWithDefaults() *OsVirtualDrive`

NewOsVirtualDriveWithDefaults instantiates a new OsVirtualDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsVirtualDrive) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsVirtualDrive) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsVirtualDrive) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsVirtualDrive) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsVirtualDrive) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsVirtualDrive) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *OsVirtualDrive) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OsVirtualDrive) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OsVirtualDrive) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OsVirtualDrive) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OsVirtualDrive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsVirtualDrive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsVirtualDrive) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsVirtualDrive) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageControllerSlotId

`func (o *OsVirtualDrive) GetStorageControllerSlotId() string`

GetStorageControllerSlotId returns the StorageControllerSlotId field if non-nil, zero value otherwise.

### GetStorageControllerSlotIdOk

`func (o *OsVirtualDrive) GetStorageControllerSlotIdOk() (*string, bool)`

GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllerSlotId

`func (o *OsVirtualDrive) SetStorageControllerSlotId(v string)`

SetStorageControllerSlotId sets StorageControllerSlotId field to given value.

### HasStorageControllerSlotId

`func (o *OsVirtualDrive) HasStorageControllerSlotId() bool`

HasStorageControllerSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


