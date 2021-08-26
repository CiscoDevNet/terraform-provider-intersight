# VirtualizationHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.Host"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.Host"]
**Action** | Pointer to **string** | Action to be performed on a host (Create, PowerState, Migrate, Clone etc). * &#x60;None&#x60; - A place holder for the default value. * &#x60;PowerOffStorageController&#x60; - Power off HyperFlex storage controller node running on selected hypervisor host. * &#x60;PowerOnStorageController&#x60; - Power on HyperFlex storage controller node running on selected hypervisor host. | [optional] [default to "None"]
**HypervisorType** | Pointer to **string** | Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with. * &#x60;ESXi&#x60; - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * &#x60;HyperFlexAp&#x60; - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * &#x60;Hyper-V&#x60; - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * &#x60;Unknown&#x60; - The hypervisor running on the HyperFlex cluster is not known. | [optional] [readonly] [default to "ESXi"]
**Identity** | Pointer to **string** | Unique identifier assigned to the hypervisor host. | [optional] [readonly] 
**Model** | Pointer to **string** | Commercial model information about this hardware. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the hypervisor host. It must be unique within the target endpoint. | [optional] [readonly] 
**Serial** | Pointer to **string** | Serial number of this host (internally generated). | [optional] [readonly] 
**Vendor** | Pointer to **string** | Commercial vendor details of this hardware. | [optional] [readonly] 
**Inventory** | Pointer to [**VirtualizationBaseHostRelationship**](VirtualizationBaseHostRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationHostAllOf

`func NewVirtualizationHostAllOf(classId string, objectType string, ) *VirtualizationHostAllOf`

NewVirtualizationHostAllOf instantiates a new VirtualizationHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationHostAllOfWithDefaults

`func NewVirtualizationHostAllOfWithDefaults() *VirtualizationHostAllOf`

NewVirtualizationHostAllOfWithDefaults instantiates a new VirtualizationHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationHostAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationHostAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationHostAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationHostAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationHostAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationHostAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *VirtualizationHostAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VirtualizationHostAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VirtualizationHostAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VirtualizationHostAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetHypervisorType

`func (o *VirtualizationHostAllOf) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationHostAllOf) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationHostAllOf) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationHostAllOf) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationHostAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationHostAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationHostAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationHostAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetModel

`func (o *VirtualizationHostAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VirtualizationHostAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VirtualizationHostAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VirtualizationHostAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationHostAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationHostAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationHostAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationHostAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *VirtualizationHostAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *VirtualizationHostAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *VirtualizationHostAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *VirtualizationHostAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *VirtualizationHostAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VirtualizationHostAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VirtualizationHostAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VirtualizationHostAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetInventory

`func (o *VirtualizationHostAllOf) GetInventory() VirtualizationBaseHostRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *VirtualizationHostAllOf) GetInventoryOk() (*VirtualizationBaseHostRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *VirtualizationHostAllOf) SetInventory(v VirtualizationBaseHostRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *VirtualizationHostAllOf) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationHostAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationHostAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationHostAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationHostAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *VirtualizationHostAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *VirtualizationHostAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *VirtualizationHostAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *VirtualizationHostAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


