# ComputeStoragePhysicalDriveOperationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.StoragePhysicalDriveOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.StoragePhysicalDriveOperation"]
**AdminAction** | Pointer to **string** | Administrative actions that can be performed on the Storage Physical Drives. * &#x60;None&#x60; - No action on the selected Storage Physical Drives. * &#x60;SetJbod&#x60; - Set Jbod action state on the selected Storage Physical Drives. * &#x60;SetUnconfiguredGood&#x60; - Set Unconfigured Good action state on the selected Storage Physical Drives. | [optional] [default to "None"]
**ControllerId** | Pointer to **string** | Storage Controller Id of the storage Physical Drives of the server. | [optional] 
**PhysicalDrives** | Pointer to [**[]ComputeStoragePhysicalDrive**](ComputeStoragePhysicalDrive.md) |  | [optional] 

## Methods

### NewComputeStoragePhysicalDriveOperationAllOf

`func NewComputeStoragePhysicalDriveOperationAllOf(classId string, objectType string, ) *ComputeStoragePhysicalDriveOperationAllOf`

NewComputeStoragePhysicalDriveOperationAllOf instantiates a new ComputeStoragePhysicalDriveOperationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeStoragePhysicalDriveOperationAllOfWithDefaults

`func NewComputeStoragePhysicalDriveOperationAllOfWithDefaults() *ComputeStoragePhysicalDriveOperationAllOf`

NewComputeStoragePhysicalDriveOperationAllOfWithDefaults instantiates a new ComputeStoragePhysicalDriveOperationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeStoragePhysicalDriveOperationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeStoragePhysicalDriveOperationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *ComputeStoragePhysicalDriveOperationAllOf) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *ComputeStoragePhysicalDriveOperationAllOf) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetControllerId

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *ComputeStoragePhysicalDriveOperationAllOf) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *ComputeStoragePhysicalDriveOperationAllOf) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetPhysicalDrives

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetPhysicalDrives() []ComputeStoragePhysicalDrive`

GetPhysicalDrives returns the PhysicalDrives field if non-nil, zero value otherwise.

### GetPhysicalDrivesOk

`func (o *ComputeStoragePhysicalDriveOperationAllOf) GetPhysicalDrivesOk() (*[]ComputeStoragePhysicalDrive, bool)`

GetPhysicalDrivesOk returns a tuple with the PhysicalDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDrives

`func (o *ComputeStoragePhysicalDriveOperationAllOf) SetPhysicalDrives(v []ComputeStoragePhysicalDrive)`

SetPhysicalDrives sets PhysicalDrives field to given value.

### HasPhysicalDrives

`func (o *ComputeStoragePhysicalDriveOperationAllOf) HasPhysicalDrives() bool`

HasPhysicalDrives returns a boolean if a field has been set.

### SetPhysicalDrivesNil

`func (o *ComputeStoragePhysicalDriveOperationAllOf) SetPhysicalDrivesNil(b bool)`

 SetPhysicalDrivesNil sets the value for PhysicalDrives to be an explicit nil

### UnsetPhysicalDrives
`func (o *ComputeStoragePhysicalDriveOperationAllOf) UnsetPhysicalDrives()`

UnsetPhysicalDrives ensures that no value is present for PhysicalDrives, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


