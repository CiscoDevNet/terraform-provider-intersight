# VirtualizationVmwareDiscoveryProtocolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareDiscoveryProtocol"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareDiscoveryProtocol"]
**Operation** | Pointer to **string** | Operational mode of the ESXI hosts connected to the distributed virtual switch. | [optional] 
**Type** | Pointer to **string** | Discovery protocol type enabled on the distributed virtual switch. | [optional] 

## Methods

### NewVirtualizationVmwareDiscoveryProtocolAllOf

`func NewVirtualizationVmwareDiscoveryProtocolAllOf(classId string, objectType string, ) *VirtualizationVmwareDiscoveryProtocolAllOf`

NewVirtualizationVmwareDiscoveryProtocolAllOf instantiates a new VirtualizationVmwareDiscoveryProtocolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDiscoveryProtocolAllOfWithDefaults

`func NewVirtualizationVmwareDiscoveryProtocolAllOfWithDefaults() *VirtualizationVmwareDiscoveryProtocolAllOf`

NewVirtualizationVmwareDiscoveryProtocolAllOfWithDefaults instantiates a new VirtualizationVmwareDiscoveryProtocolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperation

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetType

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualizationVmwareDiscoveryProtocolAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


