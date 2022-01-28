# VirtualizationVmwareDiscoveryProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareDiscoveryProtocol"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareDiscoveryProtocol"]
**Operation** | Pointer to **string** | Operational mode of the ESXI hosts connected to the distributed virtual switch. | [optional] 
**Type** | Pointer to **string** | Discovery protocol type enabled on the distributed virtual switch. | [optional] 

## Methods

### NewVirtualizationVmwareDiscoveryProtocol

`func NewVirtualizationVmwareDiscoveryProtocol(classId string, objectType string, ) *VirtualizationVmwareDiscoveryProtocol`

NewVirtualizationVmwareDiscoveryProtocol instantiates a new VirtualizationVmwareDiscoveryProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDiscoveryProtocolWithDefaults

`func NewVirtualizationVmwareDiscoveryProtocolWithDefaults() *VirtualizationVmwareDiscoveryProtocol`

NewVirtualizationVmwareDiscoveryProtocolWithDefaults instantiates a new VirtualizationVmwareDiscoveryProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareDiscoveryProtocol) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareDiscoveryProtocol) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareDiscoveryProtocol) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareDiscoveryProtocol) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareDiscoveryProtocol) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareDiscoveryProtocol) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperation

`func (o *VirtualizationVmwareDiscoveryProtocol) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VirtualizationVmwareDiscoveryProtocol) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VirtualizationVmwareDiscoveryProtocol) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *VirtualizationVmwareDiscoveryProtocol) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetType

`func (o *VirtualizationVmwareDiscoveryProtocol) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualizationVmwareDiscoveryProtocol) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualizationVmwareDiscoveryProtocol) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualizationVmwareDiscoveryProtocol) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


