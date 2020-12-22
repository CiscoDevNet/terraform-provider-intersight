# FabricQosClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.QosClass"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.QosClass"]
**AdminState** | Pointer to **string** | Administrative state for this QoS class. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**BandwidthPercent** | Pointer to **int64** | Percentage of bandwidth received by the traffic tagged with this QoS. | [optional] 
**Cos** | Pointer to **int64** | Class of service received by the traffic tagged with this QoS. | [optional] 
**Mtu** | Pointer to **int64** | Maximum transmission unit (MTU) is the largest size packet or frame, that can be sent in a packet- or frame-based network such as the Internet. User can select from the following: 1. Any value between 1500 and 9216 2. &#39;Normal&#39; (default) mapping to a value of 1500. 3. &#39;FC&#39; mapping to a value of 2240. | [optional] [default to 1500]
**MulticastOptimize** | Pointer to **bool** | If enabled, this QoS class will be optimized to send multiple packets. | [optional] [default to false]
**Name** | Pointer to **string** | The &#39;name&#39; of this QoS Class. * &#x60;Best Effort&#x60; - QoS Priority for Best-effort traffic. * &#x60;FC&#x60; - QoS Priority for FC traffic. * &#x60;Platinum&#x60; - QoS Priority for Platinum traffic. * &#x60;Gold&#x60; - QoS Priority for Gold traffic. * &#x60;Silver&#x60; - QoS Priority for Silver traffic. * &#x60;Bronze&#x60; - QoS Priority for Bronze traffic. | [optional] [default to "Best Effort"]
**PacketDrop** | Pointer to **bool** | If enabled, this QoS class will allow packet drops within an acceptable limit. | [optional] [default to true]
**Weight** | Pointer to **int64** | The weight of the QoS Class controls the distribution of bandwidth between QoS Classes, with the same priority after the Guarantees for the QoS Classes are reached. | [optional] 

## Methods

### NewFabricQosClass

`func NewFabricQosClass(classId string, objectType string, ) *FabricQosClass`

NewFabricQosClass instantiates a new FabricQosClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricQosClassWithDefaults

`func NewFabricQosClassWithDefaults() *FabricQosClass`

NewFabricQosClassWithDefaults instantiates a new FabricQosClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricQosClass) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricQosClass) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricQosClass) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricQosClass) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricQosClass) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricQosClass) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *FabricQosClass) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FabricQosClass) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FabricQosClass) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FabricQosClass) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetBandwidthPercent

`func (o *FabricQosClass) GetBandwidthPercent() int64`

GetBandwidthPercent returns the BandwidthPercent field if non-nil, zero value otherwise.

### GetBandwidthPercentOk

`func (o *FabricQosClass) GetBandwidthPercentOk() (*int64, bool)`

GetBandwidthPercentOk returns a tuple with the BandwidthPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthPercent

`func (o *FabricQosClass) SetBandwidthPercent(v int64)`

SetBandwidthPercent sets BandwidthPercent field to given value.

### HasBandwidthPercent

`func (o *FabricQosClass) HasBandwidthPercent() bool`

HasBandwidthPercent returns a boolean if a field has been set.

### GetCos

`func (o *FabricQosClass) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *FabricQosClass) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *FabricQosClass) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *FabricQosClass) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetMtu

`func (o *FabricQosClass) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *FabricQosClass) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *FabricQosClass) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *FabricQosClass) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetMulticastOptimize

`func (o *FabricQosClass) GetMulticastOptimize() bool`

GetMulticastOptimize returns the MulticastOptimize field if non-nil, zero value otherwise.

### GetMulticastOptimizeOk

`func (o *FabricQosClass) GetMulticastOptimizeOk() (*bool, bool)`

GetMulticastOptimizeOk returns a tuple with the MulticastOptimize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastOptimize

`func (o *FabricQosClass) SetMulticastOptimize(v bool)`

SetMulticastOptimize sets MulticastOptimize field to given value.

### HasMulticastOptimize

`func (o *FabricQosClass) HasMulticastOptimize() bool`

HasMulticastOptimize returns a boolean if a field has been set.

### GetName

`func (o *FabricQosClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricQosClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricQosClass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricQosClass) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPacketDrop

`func (o *FabricQosClass) GetPacketDrop() bool`

GetPacketDrop returns the PacketDrop field if non-nil, zero value otherwise.

### GetPacketDropOk

`func (o *FabricQosClass) GetPacketDropOk() (*bool, bool)`

GetPacketDropOk returns a tuple with the PacketDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketDrop

`func (o *FabricQosClass) SetPacketDrop(v bool)`

SetPacketDrop sets PacketDrop field to given value.

### HasPacketDrop

`func (o *FabricQosClass) HasPacketDrop() bool`

HasPacketDrop returns a boolean if a field has been set.

### GetWeight

`func (o *FabricQosClass) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *FabricQosClass) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *FabricQosClass) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *FabricQosClass) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


