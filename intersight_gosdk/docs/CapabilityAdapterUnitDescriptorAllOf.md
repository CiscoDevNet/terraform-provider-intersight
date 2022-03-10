# CapabilityAdapterUnitDescriptorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.AdapterUnitDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.AdapterUnitDescriptor"]
**AdapterGeneration** | Pointer to **int32** | Generation of the adapter. * &#x60;4&#x60; - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04. * &#x60;2&#x60; - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02. * &#x60;3&#x60; - Third generation adapters (13xx). The PIDs of these adapters end with the string 03. * &#x60;5&#x60; - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string. | [optional] [default to 4]
**ConnectivityOrder** | Pointer to **string** | Order in which the ports are connected; sequential or cyclic. | [optional] 
**EthernetPortSpeed** | Pointer to **int64** | The port speed for ethernet ports in Mbps. | [optional] 
**Features** | Pointer to [**[]CapabilityFeatureConfig**](CapabilityFeatureConfig.md) |  | [optional] 
**FibreChannelPortSpeed** | Pointer to **int64** | The port speed for fibre channel ports in Mbps. | [optional] 
**FibreChannelScsiIoqLimit** | Pointer to **int64** | The number of SCSI I/O Queue resources to allocate. | [optional] 
**IsAzureQosSupported** | Pointer to **bool** | Indicates that the Azure Stack Host QoS feature is supported by this adapter. | [optional] [default to true]
**IsGeneveSupported** | Pointer to **bool** | Indicates that the GENEVE offload feature is supported by this adapter. | [optional] [default to true]
**MaxEthRxRingSize** | Pointer to **int64** | Maximum Ring Size value for vNIC Receive Queue. | [optional] [default to 4096]
**MaxEthTxRingSize** | Pointer to **int64** | Maximum Ring Size value for vNIC Transmit Queue. | [optional] [default to 4096]
**MaxRocev2Interfaces** | Pointer to **int64** | Maximum number of vNIC interfaces that can be RoCEv2 enabled. | [optional] [default to 2]
**NumDcePorts** | Pointer to **int64** | Number of Dce Ports for the adapter. | [optional] 
**PciLink** | Pointer to **int64** | Indicates PCI Link status of adapter. | [optional] [default to 0]
**PromCardType** | Pointer to **string** | Prom card type for the adapter. | [optional] 

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


### GetAdapterGeneration

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetAdapterGeneration() int32`

GetAdapterGeneration returns the AdapterGeneration field if non-nil, zero value otherwise.

### GetAdapterGenerationOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetAdapterGenerationOk() (*int32, bool)`

GetAdapterGenerationOk returns a tuple with the AdapterGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterGeneration

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetAdapterGeneration(v int32)`

SetAdapterGeneration sets AdapterGeneration field to given value.

### HasAdapterGeneration

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasAdapterGeneration() bool`

HasAdapterGeneration returns a boolean if a field has been set.

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

### GetFeatures

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetFeatures() []CapabilityFeatureConfig`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetFeaturesOk() (*[]CapabilityFeatureConfig, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetFeatures(v []CapabilityFeatureConfig)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *CapabilityAdapterUnitDescriptorAllOf) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
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

### GetIsAzureQosSupported

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetIsAzureQosSupported() bool`

GetIsAzureQosSupported returns the IsAzureQosSupported field if non-nil, zero value otherwise.

### GetIsAzureQosSupportedOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetIsAzureQosSupportedOk() (*bool, bool)`

GetIsAzureQosSupportedOk returns a tuple with the IsAzureQosSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzureQosSupported

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetIsAzureQosSupported(v bool)`

SetIsAzureQosSupported sets IsAzureQosSupported field to given value.

### HasIsAzureQosSupported

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasIsAzureQosSupported() bool`

HasIsAzureQosSupported returns a boolean if a field has been set.

### GetIsGeneveSupported

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetIsGeneveSupported() bool`

GetIsGeneveSupported returns the IsGeneveSupported field if non-nil, zero value otherwise.

### GetIsGeneveSupportedOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetIsGeneveSupportedOk() (*bool, bool)`

GetIsGeneveSupportedOk returns a tuple with the IsGeneveSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGeneveSupported

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetIsGeneveSupported(v bool)`

SetIsGeneveSupported sets IsGeneveSupported field to given value.

### HasIsGeneveSupported

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasIsGeneveSupported() bool`

HasIsGeneveSupported returns a boolean if a field has been set.

### GetMaxEthRxRingSize

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxEthRxRingSize() int64`

GetMaxEthRxRingSize returns the MaxEthRxRingSize field if non-nil, zero value otherwise.

### GetMaxEthRxRingSizeOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxEthRxRingSizeOk() (*int64, bool)`

GetMaxEthRxRingSizeOk returns a tuple with the MaxEthRxRingSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEthRxRingSize

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetMaxEthRxRingSize(v int64)`

SetMaxEthRxRingSize sets MaxEthRxRingSize field to given value.

### HasMaxEthRxRingSize

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasMaxEthRxRingSize() bool`

HasMaxEthRxRingSize returns a boolean if a field has been set.

### GetMaxEthTxRingSize

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxEthTxRingSize() int64`

GetMaxEthTxRingSize returns the MaxEthTxRingSize field if non-nil, zero value otherwise.

### GetMaxEthTxRingSizeOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxEthTxRingSizeOk() (*int64, bool)`

GetMaxEthTxRingSizeOk returns a tuple with the MaxEthTxRingSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEthTxRingSize

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetMaxEthTxRingSize(v int64)`

SetMaxEthTxRingSize sets MaxEthTxRingSize field to given value.

### HasMaxEthTxRingSize

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasMaxEthTxRingSize() bool`

HasMaxEthTxRingSize returns a boolean if a field has been set.

### GetMaxRocev2Interfaces

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxRocev2Interfaces() int64`

GetMaxRocev2Interfaces returns the MaxRocev2Interfaces field if non-nil, zero value otherwise.

### GetMaxRocev2InterfacesOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxRocev2InterfacesOk() (*int64, bool)`

GetMaxRocev2InterfacesOk returns a tuple with the MaxRocev2Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRocev2Interfaces

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetMaxRocev2Interfaces(v int64)`

SetMaxRocev2Interfaces sets MaxRocev2Interfaces field to given value.

### HasMaxRocev2Interfaces

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasMaxRocev2Interfaces() bool`

HasMaxRocev2Interfaces returns a boolean if a field has been set.

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

### GetPciLink

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetPciLink() int64`

GetPciLink returns the PciLink field if non-nil, zero value otherwise.

### GetPciLinkOk

`func (o *CapabilityAdapterUnitDescriptorAllOf) GetPciLinkOk() (*int64, bool)`

GetPciLinkOk returns a tuple with the PciLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciLink

`func (o *CapabilityAdapterUnitDescriptorAllOf) SetPciLink(v int64)`

SetPciLink sets PciLink field to given value.

### HasPciLink

`func (o *CapabilityAdapterUnitDescriptorAllOf) HasPciLink() bool`

HasPciLink returns a boolean if a field has been set.

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


