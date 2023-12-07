# NetworkVethernetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.Vethernet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.Vethernet"]
**Description** | Pointer to **string** | Description for the vNIC. | [optional] [readonly] 
**VethId** | Pointer to **int64** | Vethernet Identifier on a Fabric Interconnect. | [optional] [readonly] 
**AdapterHostEthInterface** | Pointer to [**AdapterHostEthInterfaceRelationship**](AdapterHostEthInterfaceRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVethernetAllOf

`func NewNetworkVethernetAllOf(classId string, objectType string, ) *NetworkVethernetAllOf`

NewNetworkVethernetAllOf instantiates a new NetworkVethernetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVethernetAllOfWithDefaults

`func NewNetworkVethernetAllOfWithDefaults() *NetworkVethernetAllOf`

NewNetworkVethernetAllOfWithDefaults instantiates a new NetworkVethernetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVethernetAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVethernetAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVethernetAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVethernetAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVethernetAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVethernetAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NetworkVethernetAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkVethernetAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkVethernetAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkVethernetAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVethId

`func (o *NetworkVethernetAllOf) GetVethId() int64`

GetVethId returns the VethId field if non-nil, zero value otherwise.

### GetVethIdOk

`func (o *NetworkVethernetAllOf) GetVethIdOk() (*int64, bool)`

GetVethIdOk returns a tuple with the VethId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVethId

`func (o *NetworkVethernetAllOf) SetVethId(v int64)`

SetVethId sets VethId field to given value.

### HasVethId

`func (o *NetworkVethernetAllOf) HasVethId() bool`

HasVethId returns a boolean if a field has been set.

### GetAdapterHostEthInterface

`func (o *NetworkVethernetAllOf) GetAdapterHostEthInterface() AdapterHostEthInterfaceRelationship`

GetAdapterHostEthInterface returns the AdapterHostEthInterface field if non-nil, zero value otherwise.

### GetAdapterHostEthInterfaceOk

`func (o *NetworkVethernetAllOf) GetAdapterHostEthInterfaceOk() (*AdapterHostEthInterfaceRelationship, bool)`

GetAdapterHostEthInterfaceOk returns a tuple with the AdapterHostEthInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterHostEthInterface

`func (o *NetworkVethernetAllOf) SetAdapterHostEthInterface(v AdapterHostEthInterfaceRelationship)`

SetAdapterHostEthInterface sets AdapterHostEthInterface field to given value.

### HasAdapterHostEthInterface

`func (o *NetworkVethernetAllOf) HasAdapterHostEthInterface() bool`

HasAdapterHostEthInterface returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkVethernetAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVethernetAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVethernetAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVethernetAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkVethernetAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVethernetAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVethernetAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVethernetAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


