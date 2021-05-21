# VirtualizationBaseHypervisorManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Build** | Pointer to **string** | The build number of the Hypervisor Manger (e.g., 4541947, 6.3.9600.18692). The build number may indicate some feature support that applications might rely on. The build number may not always be an integer. | [optional] [readonly] 
**Identity** | Pointer to **string** | Identity of the hypervisor (not manipulated by user). It could be a UUID too. For example, c917093f-5443-4748-bc09-eec72ded7608. | [optional] [readonly] 
**Name** | Pointer to **string** | The user provided name for the hypervisor manager. For example, vCenterIreland. Usually, this name is subject to manipulation by the user. It is not the identity of the hypervisor. | [optional] 
**Version** | Pointer to **string** | Release version of the Hypervisor Manger (VMware vCenter Server 6.0.0 build-4541947). | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationBaseHypervisorManager

`func NewVirtualizationBaseHypervisorManager(classId string, objectType string, ) *VirtualizationBaseHypervisorManager`

NewVirtualizationBaseHypervisorManager instantiates a new VirtualizationBaseHypervisorManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseHypervisorManagerWithDefaults

`func NewVirtualizationBaseHypervisorManagerWithDefaults() *VirtualizationBaseHypervisorManager`

NewVirtualizationBaseHypervisorManagerWithDefaults instantiates a new VirtualizationBaseHypervisorManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseHypervisorManager) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseHypervisorManager) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseHypervisorManager) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseHypervisorManager) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseHypervisorManager) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseHypervisorManager) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBuild

`func (o *VirtualizationBaseHypervisorManager) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *VirtualizationBaseHypervisorManager) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *VirtualizationBaseHypervisorManager) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *VirtualizationBaseHypervisorManager) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseHypervisorManager) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseHypervisorManager) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseHypervisorManager) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseHypervisorManager) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseHypervisorManager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseHypervisorManager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseHypervisorManager) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseHypervisorManager) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualizationBaseHypervisorManager) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualizationBaseHypervisorManager) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualizationBaseHypervisorManager) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualizationBaseHypervisorManager) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationBaseHypervisorManager) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationBaseHypervisorManager) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationBaseHypervisorManager) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationBaseHypervisorManager) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


