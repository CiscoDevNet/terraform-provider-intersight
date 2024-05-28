# VirtualizationVmwareVirtualMachineSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVirtualMachineSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVirtualMachineSnapshot"]
**CreationTime** | Pointer to **time.Time** | Snapshot creation time. Time at which snapshot gets created. | [optional] 
**CurrentSnapshot** | Pointer to **bool** | If yes, it determines it is the latest snapshot of the virtual machine. | [optional] 
**Description** | Pointer to **string** | User provided description of the virtual machine snapshot. | [optional] 
**Golden** | Pointer to **bool** | If yes, the virtual machine snapshot cannot be deleted. | [optional] 
**Key** | Pointer to **int64** | The internally assigned id/key of virtual machine snapshot. | [optional] 
**PredecessorId** | Pointer to **int64** | Predecessor id is the id of the parent snapshot. | [optional] 
**Quiesced** | Pointer to **bool** | Quiesce pauses all the I/O operations on virtual machine till the snapshot is taken. | [optional] 
**RefValue** | Pointer to **string** | Internally assigned MOR reference value. | [optional] 
**SnapshotSize** | Pointer to **int64** | Size of the snapshot file created of the virtual machine, stored in bytes. | [optional] 
**VirtualMachine** | Pointer to [**NullableVirtualizationVmwareVirtualMachineRelationship**](VirtualizationVmwareVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareVirtualMachineSnapshot

`func NewVirtualizationVmwareVirtualMachineSnapshot(classId string, objectType string, ) *VirtualizationVmwareVirtualMachineSnapshot`

NewVirtualizationVmwareVirtualMachineSnapshot instantiates a new VirtualizationVmwareVirtualMachineSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualMachineSnapshotWithDefaults

`func NewVirtualizationVmwareVirtualMachineSnapshotWithDefaults() *VirtualizationVmwareVirtualMachineSnapshot`

NewVirtualizationVmwareVirtualMachineSnapshotWithDefaults instantiates a new VirtualizationVmwareVirtualMachineSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreationTime

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrentSnapshot

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetCurrentSnapshot() bool`

GetCurrentSnapshot returns the CurrentSnapshot field if non-nil, zero value otherwise.

### GetCurrentSnapshotOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetCurrentSnapshotOk() (*bool, bool)`

GetCurrentSnapshotOk returns a tuple with the CurrentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSnapshot

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetCurrentSnapshot(v bool)`

SetCurrentSnapshot sets CurrentSnapshot field to given value.

### HasCurrentSnapshot

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasCurrentSnapshot() bool`

HasCurrentSnapshot returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGolden

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetGolden() bool`

GetGolden returns the Golden field if non-nil, zero value otherwise.

### GetGoldenOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetGoldenOk() (*bool, bool)`

GetGoldenOk returns a tuple with the Golden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGolden

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetGolden(v bool)`

SetGolden sets Golden field to given value.

### HasGolden

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasGolden() bool`

HasGolden returns a boolean if a field has been set.

### GetKey

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetKey() int64`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetKeyOk() (*int64, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetKey(v int64)`

SetKey sets Key field to given value.

### HasKey

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPredecessorId

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetPredecessorId() int64`

GetPredecessorId returns the PredecessorId field if non-nil, zero value otherwise.

### GetPredecessorIdOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetPredecessorIdOk() (*int64, bool)`

GetPredecessorIdOk returns a tuple with the PredecessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorId

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetPredecessorId(v int64)`

SetPredecessorId sets PredecessorId field to given value.

### HasPredecessorId

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasPredecessorId() bool`

HasPredecessorId returns a boolean if a field has been set.

### GetQuiesced

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetQuiesced() bool`

GetQuiesced returns the Quiesced field if non-nil, zero value otherwise.

### GetQuiescedOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetQuiescedOk() (*bool, bool)`

GetQuiescedOk returns a tuple with the Quiesced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesced

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetQuiesced(v bool)`

SetQuiesced sets Quiesced field to given value.

### HasQuiesced

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasQuiesced() bool`

HasQuiesced returns a boolean if a field has been set.

### GetRefValue

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetRefValue() string`

GetRefValue returns the RefValue field if non-nil, zero value otherwise.

### GetRefValueOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetRefValueOk() (*string, bool)`

GetRefValueOk returns a tuple with the RefValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefValue

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetRefValue(v string)`

SetRefValue sets RefValue field to given value.

### HasRefValue

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasRefValue() bool`

HasRefValue returns a boolean if a field has been set.

### GetSnapshotSize

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetSnapshotSize() int64`

GetSnapshotSize returns the SnapshotSize field if non-nil, zero value otherwise.

### GetSnapshotSizeOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetSnapshotSizeOk() (*int64, bool)`

GetSnapshotSizeOk returns a tuple with the SnapshotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSize

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetSnapshotSize(v int64)`

SetSnapshotSize sets SnapshotSize field to given value.

### HasSnapshotSize

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasSnapshotSize() bool`

HasSnapshotSize returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetVirtualMachine() VirtualizationVmwareVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VirtualizationVmwareVirtualMachineSnapshot) GetVirtualMachineOk() (*VirtualizationVmwareVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetVirtualMachine(v VirtualizationVmwareVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VirtualizationVmwareVirtualMachineSnapshot) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### SetVirtualMachineNil

`func (o *VirtualizationVmwareVirtualMachineSnapshot) SetVirtualMachineNil(b bool)`

 SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil

### UnsetVirtualMachine
`func (o *VirtualizationVmwareVirtualMachineSnapshot) UnsetVirtualMachine()`

UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


