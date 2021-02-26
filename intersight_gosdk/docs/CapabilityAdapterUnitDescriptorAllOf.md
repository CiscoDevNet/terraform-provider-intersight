# CapabilityAdapterUnitDescriptorAllOf

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

### NewCapabilityAdapterUnitDescriptorAllOf

`func NewCapabilityAdapterUnitDescriptorAllOf(classId string, objectType string, ) *CapabilityAdapterUnitDescriptorAllOf`

NewCapabilityAdapterUnitDescriptorAllOf instantiates a new CapabilityAdapterUnitDescriptorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityAdapterUnitDescriptorAllOfWithDefaults

`func NewCapabilityAdapterUnitDescriptorAllOfWithDefaults() *CapabilityAdapterUnitDescriptorAllOf`

NewCapabilityAdapterUnitDescriptorAllOfWithDefaults instantiates a new CapabilityAdapterUnitDescriptorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectivityOrder

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetConnectivityOrder() string`

GetConnectivityOrder returns the ConnectivityOrder field if non-nil, zero value otherwise.

### GetConnectivityOrderOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetConnectivityOrderOk() (*string, bool)`

GetConnectivityOrderOk returns a tuple with the ConnectivityOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityOrder

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetConnectivityOrder(v string)`

SetConnectivityOrder sets ConnectivityOrder field to given value.

### HasConnectivityOrder

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasConnectivityOrder() bool`

HasConnectivityOrder returns a boolean if a field has been set.

### GetEthernetPortSpeed

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetEthernetPortSpeed() int64`

GetEthernetPortSpeed returns the EthernetPortSpeed field if non-nil, zero value otherwise.

### GetEthernetPortSpeedOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetEthernetPortSpeedOk() (*int64, bool)`

GetEthernetPortSpeedOk returns a tuple with the EthernetPortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetPortSpeed

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetEthernetPortSpeed(v int64)`

SetEthernetPortSpeed sets EthernetPortSpeed field to given value.

### HasEthernetPortSpeed

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasEthernetPortSpeed() bool`

HasEthernetPortSpeed returns a boolean if a field has been set.

### GetFibreChannelPortSpeed

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelPortSpeed() int64`

GetFibreChannelPortSpeed returns the FibreChannelPortSpeed field if non-nil, zero value otherwise.

### GetFibreChannelPortSpeedOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelPortSpeedOk() (*int64, bool)`

GetFibreChannelPortSpeedOk returns a tuple with the FibreChannelPortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFibreChannelPortSpeed

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetFibreChannelPortSpeed(v int64)`

SetFibreChannelPortSpeed sets FibreChannelPortSpeed field to given value.

### HasFibreChannelPortSpeed

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasFibreChannelPortSpeed() bool`

HasFibreChannelPortSpeed returns a boolean if a field has been set.

### GetFibreChannelScsiIoqLimit

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelScsiIoqLimit() int64`

GetFibreChannelScsiIoqLimit returns the FibreChannelScsiIoqLimit field if non-nil, zero value otherwise.

### GetFibreChannelScsiIoqLimitOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelScsiIoqLimitOk() (*int64, bool)`

GetFibreChannelScsiIoqLimitOk returns a tuple with the FibreChannelScsiIoqLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFibreChannelScsiIoqLimit

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetFibreChannelScsiIoqLimit(v int64)`

SetFibreChannelScsiIoqLimit sets FibreChannelScsiIoqLimit field to given value.

### HasFibreChannelScsiIoqLimit

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasFibreChannelScsiIoqLimit() bool`

HasFibreChannelScsiIoqLimit returns a boolean if a field has been set.

### GetNumDcePorts

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetNumDcePorts() int64`

GetNumDcePorts returns the NumDcePorts field if non-nil, zero value otherwise.

### GetNumDcePortsOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetNumDcePortsOk() (*int64, bool)`

GetNumDcePortsOk returns a tuple with the NumDcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDcePorts

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetNumDcePorts(v int64)`

SetNumDcePorts sets NumDcePorts field to given value.

### HasNumDcePorts

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasNumDcePorts() bool`

HasNumDcePorts returns a boolean if a field has been set.

### GetPromCardType

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetPromCardType() string`

GetPromCardType returns the PromCardType field if non-nil, zero value otherwise.

### GetPromCardTypeOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetPromCardTypeOk() (*string, bool)`

GetPromCardTypeOk returns a tuple with the PromCardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromCardType

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetPromCardType(v string)`

SetPromCardType sets PromCardType field to given value.

### HasPromCardType

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasPromCardType() bool`

HasPromCardType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


