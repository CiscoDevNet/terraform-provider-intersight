# VirtualizationVmwareVirtualMachineSnapshotAllOf

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
**VirtualMachine** | Pointer to [**VirtualizationVmwareVirtualMachineRelationship**](VirtualizationVmwareVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareVirtualMachineSnapshotAllOf

`func NewVirtualizationVmwareVirtualMachineSnapshotAllOf(classId string, objectType string, ) *VirtualizationVmwareVirtualMachineSnapshotAllOf`

NewVirtualizationVmwareVirtualMachineSnapshotAllOf instantiates a new VirtualizationVmwareVirtualMachineSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualMachineSnapshotAllOfWithDefaults

`func NewVirtualizationVmwareVirtualMachineSnapshotAllOfWithDefaults() *VirtualizationVmwareVirtualMachineSnapshotAllOf`

NewVirtualizationVmwareVirtualMachineSnapshotAllOfWithDefaults instantiates a new VirtualizationVmwareVirtualMachineSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreationTime

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrentSnapshot

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetCurrentSnapshot() bool`

GetCurrentSnapshot returns the CurrentSnapshot field if non-nil, zero value otherwise.

### GetCurrentSnapshotOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetCurrentSnapshotOk() (*bool, bool)`

GetCurrentSnapshotOk returns a tuple with the CurrentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSnapshot

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetCurrentSnapshot(v bool)`

SetCurrentSnapshot sets CurrentSnapshot field to given value.

### HasCurrentSnapshot

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasCurrentSnapshot() bool`

HasCurrentSnapshot returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGolden

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetGolden() bool`

GetGolden returns the Golden field if non-nil, zero value otherwise.

### GetGoldenOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetGoldenOk() (*bool, bool)`

GetGoldenOk returns a tuple with the Golden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGolden

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetGolden(v bool)`

SetGolden sets Golden field to given value.

### HasGolden

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasGolden() bool`

HasGolden returns a boolean if a field has been set.

### GetKey

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetKey() int64`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetKeyOk() (*int64, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetKey(v int64)`

SetKey sets Key field to given value.

### HasKey

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPredecessorId

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetPredecessorId() int64`

GetPredecessorId returns the PredecessorId field if non-nil, zero value otherwise.

### GetPredecessorIdOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetPredecessorIdOk() (*int64, bool)`

GetPredecessorIdOk returns a tuple with the PredecessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorId

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetPredecessorId(v int64)`

SetPredecessorId sets PredecessorId field to given value.

### HasPredecessorId

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasPredecessorId() bool`

HasPredecessorId returns a boolean if a field has been set.

### GetQuiesced

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetQuiesced() bool`

GetQuiesced returns the Quiesced field if non-nil, zero value otherwise.

### GetQuiescedOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetQuiescedOk() (*bool, bool)`

GetQuiescedOk returns a tuple with the Quiesced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesced

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetQuiesced(v bool)`

SetQuiesced sets Quiesced field to given value.

### HasQuiesced

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasQuiesced() bool`

HasQuiesced returns a boolean if a field has been set.

### GetRefValue

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetRefValue() string`

GetRefValue returns the RefValue field if non-nil, zero value otherwise.

### GetRefValueOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetRefValueOk() (*string, bool)`

GetRefValueOk returns a tuple with the RefValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefValue

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetRefValue(v string)`

SetRefValue sets RefValue field to given value.

### HasRefValue

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasRefValue() bool`

HasRefValue returns a boolean if a field has been set.

### GetSnapshotSize

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetSnapshotSize() int64`

GetSnapshotSize returns the SnapshotSize field if non-nil, zero value otherwise.

### GetSnapshotSizeOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetSnapshotSizeOk() (*int64, bool)`

GetSnapshotSizeOk returns a tuple with the SnapshotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSize

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetSnapshotSize(v int64)`

SetSnapshotSize sets SnapshotSize field to given value.

### HasSnapshotSize

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasSnapshotSize() bool`

HasSnapshotSize returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetVirtualMachine() VirtualizationVmwareVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) GetVirtualMachineOk() (*VirtualizationVmwareVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) SetVirtualMachine(v VirtualizationVmwareVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VirtualizationVmwareVirtualMachineSnapshotAllOf) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


