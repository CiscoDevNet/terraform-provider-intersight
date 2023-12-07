# NetworkVfc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.Vfc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.Vfc"]
**Description** | Pointer to **string** | Description for the vHBA. | [optional] [readonly] 
**VfcId** | Pointer to **int64** | Vfc Identifier on a Fabric Interconnect. | [optional] [readonly] 
**AdapterHostFcInterface** | Pointer to [**AdapterHostFcInterfaceRelationship**](AdapterHostFcInterfaceRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVfc

`func NewNetworkVfc(classId string, objectType string, ) *NetworkVfc`

NewNetworkVfc instantiates a new NetworkVfc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVfcWithDefaults

`func NewNetworkVfcWithDefaults() *NetworkVfc`

NewNetworkVfcWithDefaults instantiates a new NetworkVfc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVfc) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVfc) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVfc) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVfc) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVfc) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVfc) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NetworkVfc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkVfc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkVfc) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkVfc) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVfcId

`func (o *NetworkVfc) GetVfcId() int64`

GetVfcId returns the VfcId field if non-nil, zero value otherwise.

### GetVfcIdOk

`func (o *NetworkVfc) GetVfcIdOk() (*int64, bool)`

GetVfcIdOk returns a tuple with the VfcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVfcId

`func (o *NetworkVfc) SetVfcId(v int64)`

SetVfcId sets VfcId field to given value.

### HasVfcId

`func (o *NetworkVfc) HasVfcId() bool`

HasVfcId returns a boolean if a field has been set.

### GetAdapterHostFcInterface

`func (o *NetworkVfc) GetAdapterHostFcInterface() AdapterHostFcInterfaceRelationship`

GetAdapterHostFcInterface returns the AdapterHostFcInterface field if non-nil, zero value otherwise.

### GetAdapterHostFcInterfaceOk

`func (o *NetworkVfc) GetAdapterHostFcInterfaceOk() (*AdapterHostFcInterfaceRelationship, bool)`

GetAdapterHostFcInterfaceOk returns a tuple with the AdapterHostFcInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterHostFcInterface

`func (o *NetworkVfc) SetAdapterHostFcInterface(v AdapterHostFcInterfaceRelationship)`

SetAdapterHostFcInterface sets AdapterHostFcInterface field to given value.

### HasAdapterHostFcInterface

`func (o *NetworkVfc) HasAdapterHostFcInterface() bool`

HasAdapterHostFcInterface returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkVfc) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVfc) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVfc) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVfc) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkVfc) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVfc) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVfc) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVfc) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


