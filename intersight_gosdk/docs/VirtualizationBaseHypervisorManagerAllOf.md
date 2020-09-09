# VirtualizationBaseHypervisorManagerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to **string** | Identity of the hypervisor (not manipulated by user). It could be a UUID too. For example, c917093f-5443-4748-bc09-eec72ded7608. | [optional] [readonly] 
**Name** | Pointer to **string** | The user provided name for the hypervisor manager. For example, vCenterIreland. Usually, this name is subject to manipulation by the user. It is not the identity of the hypervisor. | [optional] 
**Version** | Pointer to **string** | Release version of the Hypervisor Manger (VMware vCenter Server 6.0.0 build-4541947). | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationBaseHypervisorManagerAllOf

`func NewVirtualizationBaseHypervisorManagerAllOf() *VirtualizationBaseHypervisorManagerAllOf`

NewVirtualizationBaseHypervisorManagerAllOf instantiates a new VirtualizationBaseHypervisorManagerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseHypervisorManagerAllOfWithDefaults

`func NewVirtualizationBaseHypervisorManagerAllOfWithDefaults() *VirtualizationBaseHypervisorManagerAllOf`

NewVirtualizationBaseHypervisorManagerAllOfWithDefaults instantiates a new VirtualizationBaseHypervisorManagerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *VirtualizationBaseHypervisorManagerAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseHypervisorManagerAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseHypervisorManagerAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseHypervisorManagerAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseHypervisorManagerAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseHypervisorManagerAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseHypervisorManagerAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseHypervisorManagerAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualizationBaseHypervisorManagerAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualizationBaseHypervisorManagerAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualizationBaseHypervisorManagerAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualizationBaseHypervisorManagerAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationBaseHypervisorManagerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationBaseHypervisorManagerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationBaseHypervisorManagerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationBaseHypervisorManagerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


