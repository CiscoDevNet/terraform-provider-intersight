# OsVirtualDriveAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.VirtualDrive"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.VirtualDrive"]
**Id** | Pointer to **string** | Virtual Drive ID to be used as Install Target. | [optional] 
**Name** | Pointer to **string** | The Virtual Drive Name to be used as Install Target. | [optional] 
**StorageControllerSlotId** | Pointer to **string** | The SlotID of the Storage Controller associated to the virtual drive. | [optional] 

## Methods

### NewOsVirtualDriveAllOf

`func NewOsVirtualDriveAllOf(classId string, objectType string, ) *OsVirtualDriveAllOf`

NewOsVirtualDriveAllOf instantiates a new OsVirtualDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsVirtualDriveAllOfWithDefaults

`func NewOsVirtualDriveAllOfWithDefaults() *OsVirtualDriveAllOf`

NewOsVirtualDriveAllOfWithDefaults instantiates a new OsVirtualDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsVirtualDriveAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsVirtualDriveAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsVirtualDriveAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsVirtualDriveAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsVirtualDriveAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsVirtualDriveAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *OsVirtualDriveAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OsVirtualDriveAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OsVirtualDriveAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OsVirtualDriveAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OsVirtualDriveAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsVirtualDriveAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsVirtualDriveAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsVirtualDriveAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageControllerSlotId

`func (o *OsVirtualDriveAllOf) GetStorageControllerSlotId() string`

GetStorageControllerSlotId returns the StorageControllerSlotId field if non-nil, zero value otherwise.

### GetStorageControllerSlotIdOk

`func (o *OsVirtualDriveAllOf) GetStorageControllerSlotIdOk() (*string, bool)`

GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllerSlotId

`func (o *OsVirtualDriveAllOf) SetStorageControllerSlotId(v string)`

SetStorageControllerSlotId sets StorageControllerSlotId field to given value.

### HasStorageControllerSlotId

`func (o *OsVirtualDriveAllOf) HasStorageControllerSlotId() bool`

HasStorageControllerSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


