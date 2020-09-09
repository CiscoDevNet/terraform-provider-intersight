# FabricQosClassAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **string** | Administrative state for this QoS class. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**BandwidthPercent** | Pointer to **int64** | Percentage of bandwidth received by the traffic tagged with this QoS. | [optional] 
**Cos** | Pointer to **int64** | Class of service received by the traffic tagged with this QoS. | [optional] 
**Mtu** | Pointer to **int64** | Maximum transmission unit (MTU) is the largest size packet or frame, that can be sent in a packet- or frame-based network such as the Internet. User can select from the following: 1. Any value between 1500 and 9216 2. &#39;Normal&#39; (default) mapping to a value of 1500. 3. &#39;FC&#39; mapping to a value of 2240. | [optional] 
**MulticastOptimize** | Pointer to **bool** | If enabled, this QoS class will be optimized to send multiple packets. | [optional] 
**Name** | Pointer to **string** | The &#39;name&#39; of this QoS Class. * &#x60;Best Effort&#x60; - QoS Priority for Best-effort traffic. * &#x60;FC&#x60; - QoS Priority for FC traffic. * &#x60;Platinum&#x60; - QoS Priority for Platinum traffic. * &#x60;Gold&#x60; - QoS Priority for Gold traffic. * &#x60;Silver&#x60; - QoS Priority for Silver traffic. * &#x60;Bronze&#x60; - QoS Priority for Bronze traffic. | [optional] [default to "Best Effort"]
**PacketDrop** | Pointer to **bool** | If enabled, this QoS class will allow packet drops within an acceptable limit. | [optional] 
**Weight** | Pointer to **int64** | The weight of the QoS Class controls the distribution of bandwidth between QoS Classes, with the same priority after the Guarantees for the QoS Classes are reached. | [optional] 

## Methods

### NewFabricQosClassAllOf

`func NewFabricQosClassAllOf() *FabricQosClassAllOf`

NewFabricQosClassAllOf instantiates a new FabricQosClassAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricQosClassAllOfWithDefaults

`func NewFabricQosClassAllOfWithDefaults() *FabricQosClassAllOf`

NewFabricQosClassAllOfWithDefaults instantiates a new FabricQosClassAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *FabricQosClassAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FabricQosClassAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FabricQosClassAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FabricQosClassAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetBandwidthPercent

`func (o *FabricQosClassAllOf) GetBandwidthPercent() int64`

GetBandwidthPercent returns the BandwidthPercent field if non-nil, zero value otherwise.

### GetBandwidthPercentOk

`func (o *FabricQosClassAllOf) GetBandwidthPercentOk() (*int64, bool)`

GetBandwidthPercentOk returns a tuple with the BandwidthPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthPercent

`func (o *FabricQosClassAllOf) SetBandwidthPercent(v int64)`

SetBandwidthPercent sets BandwidthPercent field to given value.

### HasBandwidthPercent

`func (o *FabricQosClassAllOf) HasBandwidthPercent() bool`

HasBandwidthPercent returns a boolean if a field has been set.

### GetCos

`func (o *FabricQosClassAllOf) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *FabricQosClassAllOf) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *FabricQosClassAllOf) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *FabricQosClassAllOf) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetMtu

`func (o *FabricQosClassAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *FabricQosClassAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *FabricQosClassAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *FabricQosClassAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetMulticastOptimize

`func (o *FabricQosClassAllOf) GetMulticastOptimize() bool`

GetMulticastOptimize returns the MulticastOptimize field if non-nil, zero value otherwise.

### GetMulticastOptimizeOk

`func (o *FabricQosClassAllOf) GetMulticastOptimizeOk() (*bool, bool)`

GetMulticastOptimizeOk returns a tuple with the MulticastOptimize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastOptimize

`func (o *FabricQosClassAllOf) SetMulticastOptimize(v bool)`

SetMulticastOptimize sets MulticastOptimize field to given value.

### HasMulticastOptimize

`func (o *FabricQosClassAllOf) HasMulticastOptimize() bool`

HasMulticastOptimize returns a boolean if a field has been set.

### GetName

`func (o *FabricQosClassAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricQosClassAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricQosClassAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricQosClassAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPacketDrop

`func (o *FabricQosClassAllOf) GetPacketDrop() bool`

GetPacketDrop returns the PacketDrop field if non-nil, zero value otherwise.

### GetPacketDropOk

`func (o *FabricQosClassAllOf) GetPacketDropOk() (*bool, bool)`

GetPacketDropOk returns a tuple with the PacketDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketDrop

`func (o *FabricQosClassAllOf) SetPacketDrop(v bool)`

SetPacketDrop sets PacketDrop field to given value.

### HasPacketDrop

`func (o *FabricQosClassAllOf) HasPacketDrop() bool`

HasPacketDrop returns a boolean if a field has been set.

### GetWeight

`func (o *FabricQosClassAllOf) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *FabricQosClassAllOf) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *FabricQosClassAllOf) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *FabricQosClassAllOf) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


