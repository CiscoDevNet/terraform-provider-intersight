# CapabilityAdapterUnitDescriptor

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


### GetAdapterGeneration

`func (o *CapabilityAdapterUnitDescriptor) GetAdapterGeneration() int32`

GetAdapterGeneration returns the AdapterGeneration field if non-nil, zero value otherwise.

### GetAdapterGenerationOk

`func (o *CapabilityAdapterUnitDescriptor) GetAdapterGenerationOk() (*int32, bool)`

GetAdapterGenerationOk returns a tuple with the AdapterGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterGeneration

`func (o *CapabilityAdapterUnitDescriptor) SetAdapterGeneration(v int32)`

SetAdapterGeneration sets AdapterGeneration field to given value.

### HasAdapterGeneration

`func (o *CapabilityAdapterUnitDescriptor) HasAdapterGeneration() bool`

HasAdapterGeneration returns a boolean if a field has been set.

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

### GetFeatures

`func (o *CapabilityAdapterUnitDescriptor) GetFeatures() []CapabilityFeatureConfig`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CapabilityAdapterUnitDescriptor) GetFeaturesOk() (*[]CapabilityFeatureConfig, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CapabilityAdapterUnitDescriptor) SetFeatures(v []CapabilityFeatureConfig)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CapabilityAdapterUnitDescriptor) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *CapabilityAdapterUnitDescriptor) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *CapabilityAdapterUnitDescriptor) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
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

### GetIsAzureQosSupported

`func (o *CapabilityAdapterUnitDescriptor) GetIsAzureQosSupported() bool`

GetIsAzureQosSupported returns the IsAzureQosSupported field if non-nil, zero value otherwise.

### GetIsAzureQosSupportedOk

`func (o *CapabilityAdapterUnitDescriptor) GetIsAzureQosSupportedOk() (*bool, bool)`

GetIsAzureQosSupportedOk returns a tuple with the IsAzureQosSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzureQosSupported

`func (o *CapabilityAdapterUnitDescriptor) SetIsAzureQosSupported(v bool)`

SetIsAzureQosSupported sets IsAzureQosSupported field to given value.

### HasIsAzureQosSupported

`func (o *CapabilityAdapterUnitDescriptor) HasIsAzureQosSupported() bool`

HasIsAzureQosSupported returns a boolean if a field has been set.

### GetIsGeneveSupported

`func (o *CapabilityAdapterUnitDescriptor) GetIsGeneveSupported() bool`

GetIsGeneveSupported returns the IsGeneveSupported field if non-nil, zero value otherwise.

### GetIsGeneveSupportedOk

`func (o *CapabilityAdapterUnitDescriptor) GetIsGeneveSupportedOk() (*bool, bool)`

GetIsGeneveSupportedOk returns a tuple with the IsGeneveSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGeneveSupported

`func (o *CapabilityAdapterUnitDescriptor) SetIsGeneveSupported(v bool)`

SetIsGeneveSupported sets IsGeneveSupported field to given value.

### HasIsGeneveSupported

`func (o *CapabilityAdapterUnitDescriptor) HasIsGeneveSupported() bool`

HasIsGeneveSupported returns a boolean if a field has been set.

### GetMaxEthRxRingSize

`func (o *CapabilityAdapterUnitDescriptor) GetMaxEthRxRingSize() int64`

GetMaxEthRxRingSize returns the MaxEthRxRingSize field if non-nil, zero value otherwise.

### GetMaxEthRxRingSizeOk

`func (o *CapabilityAdapterUnitDescriptor) GetMaxEthRxRingSizeOk() (*int64, bool)`

GetMaxEthRxRingSizeOk returns a tuple with the MaxEthRxRingSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEthRxRingSize

`func (o *CapabilityAdapterUnitDescriptor) SetMaxEthRxRingSize(v int64)`

SetMaxEthRxRingSize sets MaxEthRxRingSize field to given value.

### HasMaxEthRxRingSize

`func (o *CapabilityAdapterUnitDescriptor) HasMaxEthRxRingSize() bool`

HasMaxEthRxRingSize returns a boolean if a field has been set.

### GetMaxEthTxRingSize

`func (o *CapabilityAdapterUnitDescriptor) GetMaxEthTxRingSize() int64`

GetMaxEthTxRingSize returns the MaxEthTxRingSize field if non-nil, zero value otherwise.

### GetMaxEthTxRingSizeOk

`func (o *CapabilityAdapterUnitDescriptor) GetMaxEthTxRingSizeOk() (*int64, bool)`

GetMaxEthTxRingSizeOk returns a tuple with the MaxEthTxRingSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEthTxRingSize

`func (o *CapabilityAdapterUnitDescriptor) SetMaxEthTxRingSize(v int64)`

SetMaxEthTxRingSize sets MaxEthTxRingSize field to given value.

### HasMaxEthTxRingSize

`func (o *CapabilityAdapterUnitDescriptor) HasMaxEthTxRingSize() bool`

HasMaxEthTxRingSize returns a boolean if a field has been set.

### GetMaxRocev2Interfaces

`func (o *CapabilityAdapterUnitDescriptor) GetMaxRocev2Interfaces() int64`

GetMaxRocev2Interfaces returns the MaxRocev2Interfaces field if non-nil, zero value otherwise.

### GetMaxRocev2InterfacesOk

`func (o *CapabilityAdapterUnitDescriptor) GetMaxRocev2InterfacesOk() (*int64, bool)`

GetMaxRocev2InterfacesOk returns a tuple with the MaxRocev2Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRocev2Interfaces

`func (o *CapabilityAdapterUnitDescriptor) SetMaxRocev2Interfaces(v int64)`

SetMaxRocev2Interfaces sets MaxRocev2Interfaces field to given value.

### HasMaxRocev2Interfaces

`func (o *CapabilityAdapterUnitDescriptor) HasMaxRocev2Interfaces() bool`

HasMaxRocev2Interfaces returns a boolean if a field has been set.

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

### GetPciLink

`func (o *CapabilityAdapterUnitDescriptor) GetPciLink() int64`

GetPciLink returns the PciLink field if non-nil, zero value otherwise.

### GetPciLinkOk

`func (o *CapabilityAdapterUnitDescriptor) GetPciLinkOk() (*int64, bool)`

GetPciLinkOk returns a tuple with the PciLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciLink

`func (o *CapabilityAdapterUnitDescriptor) SetPciLink(v int64)`

SetPciLink sets PciLink field to given value.

### HasPciLink

`func (o *CapabilityAdapterUnitDescriptor) HasPciLink() bool`

HasPciLink returns a boolean if a field has been set.

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


