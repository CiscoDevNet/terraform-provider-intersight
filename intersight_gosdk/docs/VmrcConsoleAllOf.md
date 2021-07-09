# VmrcConsoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vmrc.Console"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vmrc.Console"]
**Vcenter** | Pointer to [**VirtualizationVmwareVcenterRelationship**](virtualization.VmwareVcenter.Relationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**VirtualizationVmwareVirtualMachineRelationship**](virtualization.VmwareVirtualMachine.Relationship.md) |  | [optional] 

## Methods

### NewVmrcConsoleAllOf

`func NewVmrcConsoleAllOf(classId string, objectType string, ) *VmrcConsoleAllOf`

NewVmrcConsoleAllOf instantiates a new VmrcConsoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmrcConsoleAllOfWithDefaults

`func NewVmrcConsoleAllOfWithDefaults() *VmrcConsoleAllOf`

NewVmrcConsoleAllOfWithDefaults instantiates a new VmrcConsoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VmrcConsoleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VmrcConsoleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VmrcConsoleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VmrcConsoleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VmrcConsoleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VmrcConsoleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVcenter

`func (o *VmrcConsoleAllOf) GetVcenter() VirtualizationVmwareVcenterRelationship`

GetVcenter returns the Vcenter field if non-nil, zero value otherwise.

### GetVcenterOk

`func (o *VmrcConsoleAllOf) GetVcenterOk() (*VirtualizationVmwareVcenterRelationship, bool)`

GetVcenterOk returns a tuple with the Vcenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenter

`func (o *VmrcConsoleAllOf) SetVcenter(v VirtualizationVmwareVcenterRelationship)`

SetVcenter sets Vcenter field to given value.

### HasVcenter

`func (o *VmrcConsoleAllOf) HasVcenter() bool`

HasVcenter returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *VmrcConsoleAllOf) GetVirtualMachine() VirtualizationVmwareVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VmrcConsoleAllOf) GetVirtualMachineOk() (*VirtualizationVmwareVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VmrcConsoleAllOf) SetVirtualMachine(v VirtualizationVmwareVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VmrcConsoleAllOf) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


