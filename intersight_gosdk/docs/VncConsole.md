# VncConsole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnc.Console"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnc.Console"]
**VirtualMachine** | Pointer to [**VirtualizationIweVirtualMachineRelationship**](VirtualizationIweVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewVncConsole

`func NewVncConsole(classId string, objectType string, ) *VncConsole`

NewVncConsole instantiates a new VncConsole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVncConsoleWithDefaults

`func NewVncConsoleWithDefaults() *VncConsole`

NewVncConsoleWithDefaults instantiates a new VncConsole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VncConsole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VncConsole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VncConsole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VncConsole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VncConsole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VncConsole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVirtualMachine

`func (o *VncConsole) GetVirtualMachine() VirtualizationIweVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VncConsole) GetVirtualMachineOk() (*VirtualizationIweVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VncConsole) SetVirtualMachine(v VirtualizationIweVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VncConsole) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


