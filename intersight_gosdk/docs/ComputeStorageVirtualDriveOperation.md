# ComputeStorageVirtualDriveOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.StorageVirtualDriveOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.StorageVirtualDriveOperation"]
**AdminAction** | Pointer to **string** | Administrative actions that can be performed on the Storage Virtual Drives. * &#x60;None&#x60; - No action on the selected Storage virtual Drives. * &#x60;Delete&#x60; - Delete action on the selected Storage Virtual Drives. | [optional] [default to "None"]
**ControllerId** | Pointer to **string** | Storage Controller Id of the storage Virtual Drives of the server. | [optional] 
**VirtualDrives** | Pointer to [**[]ComputeStorageVirtualDrive**](ComputeStorageVirtualDrive.md) |  | [optional] 

## Methods

### NewComputeStorageVirtualDriveOperation

`func NewComputeStorageVirtualDriveOperation(classId string, objectType string, ) *ComputeStorageVirtualDriveOperation`

NewComputeStorageVirtualDriveOperation instantiates a new ComputeStorageVirtualDriveOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeStorageVirtualDriveOperationWithDefaults

`func NewComputeStorageVirtualDriveOperationWithDefaults() *ComputeStorageVirtualDriveOperation`

NewComputeStorageVirtualDriveOperationWithDefaults instantiates a new ComputeStorageVirtualDriveOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeStorageVirtualDriveOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeStorageVirtualDriveOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeStorageVirtualDriveOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeStorageVirtualDriveOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeStorageVirtualDriveOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeStorageVirtualDriveOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *ComputeStorageVirtualDriveOperation) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *ComputeStorageVirtualDriveOperation) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *ComputeStorageVirtualDriveOperation) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *ComputeStorageVirtualDriveOperation) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetControllerId

`func (o *ComputeStorageVirtualDriveOperation) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *ComputeStorageVirtualDriveOperation) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *ComputeStorageVirtualDriveOperation) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *ComputeStorageVirtualDriveOperation) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetVirtualDrives

`func (o *ComputeStorageVirtualDriveOperation) GetVirtualDrives() []ComputeStorageVirtualDrive`

GetVirtualDrives returns the VirtualDrives field if non-nil, zero value otherwise.

### GetVirtualDrivesOk

`func (o *ComputeStorageVirtualDriveOperation) GetVirtualDrivesOk() (*[]ComputeStorageVirtualDrive, bool)`

GetVirtualDrivesOk returns a tuple with the VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrives

`func (o *ComputeStorageVirtualDriveOperation) SetVirtualDrives(v []ComputeStorageVirtualDrive)`

SetVirtualDrives sets VirtualDrives field to given value.

### HasVirtualDrives

`func (o *ComputeStorageVirtualDriveOperation) HasVirtualDrives() bool`

HasVirtualDrives returns a boolean if a field has been set.

### SetVirtualDrivesNil

`func (o *ComputeStorageVirtualDriveOperation) SetVirtualDrivesNil(b bool)`

 SetVirtualDrivesNil sets the value for VirtualDrives to be an explicit nil

### UnsetVirtualDrives
`func (o *ComputeStorageVirtualDriveOperation) UnsetVirtualDrives()`

UnsetVirtualDrives ensures that no value is present for VirtualDrives, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


