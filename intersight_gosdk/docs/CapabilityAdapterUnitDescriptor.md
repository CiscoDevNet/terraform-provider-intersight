# CapabilityAdapterUnitDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.AdapterUnitDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.AdapterUnitDescriptor"]
**ConnectivityOrder** | Pointer to **string** | Order in which the ports are connected; sequential or cyclic. | [optional] 
**EthernetPortSpeed** | Pointer to **int64** | The port speed for ethernet ports in Mbps. | [optional] 
**FibreChannelPortSpeed** | Pointer to **int64** | The port speed for fibre channel ports in Mbps. | [optional] 
**FibreChannelScsiIoqLimit** | Pointer to **int64** | The number of SCSI I/O Queue resources to allocate. | [optional] 
**NumDcePorts** | Pointer to **int64** | Number of Dce Ports for the adaptor. | [optional] 
**PromCardType** | Pointer to **string** | Prom card type for the adaptor. | [optional] 

## Methods

### NewCapabilityAdapterUnitDescriptor

`func NewCapabilityAdapterUnitDescriptor(classId string, objectType string, ) *CapabilityAdapterUnitDescriptor`

NewCapabilityAdapterUnitDescriptor instantiates a new CapabilityAdapterUnitDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityAdapterUnitDescriptorWithDefaults

`func NewCapabilityAdapterUnitDescriptorWithDefaults() *CapabilityAdapterUnitDescriptor`

NewCapabilityAdapterUnitDescriptorWithDefaults instantiates a new CapabilityAdapterUnitDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityAdapterUnitDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityAdapterUnitDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityAdapterUnitDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityAdapterUnitDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityAdapterUnitDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityAdapterUnitDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectivityOrder

`func (o *CapabilityAdapterUnitDescriptor) GetConnectivityOrder() string`

GetConnectivityOrder returns the ConnectivityOrder field if non-nil, zero value otherwise.

### GetConnectivityOrderOk

`func (o *CapabilityAdapterUnitDescriptor) GetConnectivityOrderOk() (*string, bool)`

GetConnectivityOrderOk returns a tuple with the ConnectivityOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityOrder

`func (o *CapabilityAdapterUnitDescriptor) SetConnectivityOrder(v string)`

SetConnectivityOrder sets ConnectivityOrder field to given value.

### HasConnectivityOrder

`func (o *CapabilityAdapterUnitDescriptor) HasConnectivityOrder() bool`

HasConnectivityOrder returns a boolean if a field has been set.

### GetEthernetPortSpeed

`func (o *CapabilityAdapterUnitDescriptor) GetEthernetPortSpeed() int64`

GetEthernetPortSpeed returns the EthernetPortSpeed field if non-nil, zero value otherwise.

### GetEthernetPortSpeedOk

`func (o *CapabilityAdapterUnitDescriptor) GetEthernetPortSpeedOk() (*int64, bool)`

GetEthernetPortSpeedOk returns a tuple with the EthernetPortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetPortSpeed

`func (o *CapabilityAdapterUnitDescriptor) SetEthernetPortSpeed(v int64)`

SetEthernetPortSpeed sets EthernetPortSpeed field to given value.

### HasEthernetPortSpeed

`func (o *CapabilityAdapterUnitDescriptor) HasEthernetPortSpeed() bool`

HasEthernetPortSpeed returns a boolean if a field has been set.

### GetFibreChannelPortSpeed

`func (o *CapabilityAdapterUnitDescriptor) GetFibreChannelPortSpeed() int64`

GetFibreChannelPortSpeed returns the FibreChannelPortSpeed field if non-nil, zero value otherwise.

### GetFibreChannelPortSpeedOk

`func (o *CapabilityAdapterUnitDescriptor) GetFibreChannelPortSpeedOk() (*int64, bool)`

GetFibreChannelPortSpeedOk returns a tuple with the FibreChannelPortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFibreChannelPortSpeed

`func (o *CapabilityAdapterUnitDescriptor) SetFibreChannelPortSpeed(v int64)`

SetFibreChannelPortSpeed sets FibreChannelPortSpeed field to given value.

### HasFibreChannelPortSpeed

`func (o *CapabilityAdapterUnitDescriptor) HasFibreChannelPortSpeed() bool`

HasFibreChannelPortSpeed returns a boolean if a field has been set.

### GetFibreChannelScsiIoqLimit

`func (o *CapabilityAdapterUnitDescriptor) GetFibreChannelScsiIoqLimit() int64`

GetFibreChannelScsiIoqLimit returns the FibreChannelScsiIoqLimit field if non-nil, zero value otherwise.

### GetFibreChannelScsiIoqLimitOk

`func (o *CapabilityAdapterUnitDescriptor) GetFibreChannelScsiIoqLimitOk() (*int64, bool)`

GetFibreChannelScsiIoqLimitOk returns a tuple with the FibreChannelScsiIoqLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFibreChannelScsiIoqLimit

`func (o *CapabilityAdapterUnitDescriptor) SetFibreChannelScsiIoqLimit(v int64)`

SetFibreChannelScsiIoqLimit sets FibreChannelScsiIoqLimit field to given value.

### HasFibreChannelScsiIoqLimit

`func (o *CapabilityAdapterUnitDescriptor) HasFibreChannelScsiIoqLimit() bool`

HasFibreChannelScsiIoqLimit returns a boolean if a field has been set.

### GetNumDcePorts

`func (o *CapabilityAdapterUnitDescriptor) GetNumDcePorts() int64`

GetNumDcePorts returns the NumDcePorts field if non-nil, zero value otherwise.

### GetNumDcePortsOk

`func (o *CapabilityAdapterUnitDescriptor) GetNumDcePortsOk() (*int64, bool)`

GetNumDcePortsOk returns a tuple with the NumDcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDcePorts

`func (o *CapabilityAdapterUnitDescriptor) SetNumDcePorts(v int64)`

SetNumDcePorts sets NumDcePorts field to given value.

### HasNumDcePorts

`func (o *CapabilityAdapterUnitDescriptor) HasNumDcePorts() bool`

HasNumDcePorts returns a boolean if a field has been set.

### GetPromCardType

`func (o *CapabilityAdapterUnitDescriptor) GetPromCardType() string`

GetPromCardType returns the PromCardType field if non-nil, zero value otherwise.

### GetPromCardTypeOk

`func (o *CapabilityAdapterUnitDescriptor) GetPromCardTypeOk() (*string, bool)`

GetPromCardTypeOk returns a tuple with the PromCardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromCardType

`func (o *CapabilityAdapterUnitDescriptor) SetPromCardType(v string)`

SetPromCardType sets PromCardType field to given value.

### HasPromCardType

`func (o *CapabilityAdapterUnitDescriptor) HasPromCardType() bool`

HasPromCardType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


