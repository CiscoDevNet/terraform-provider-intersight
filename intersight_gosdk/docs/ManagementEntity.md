# ManagementEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterLinkState** | Pointer to **string** | Cluster link state between the Fabric Interconnects. | [optional] [readonly] 
**ClusterReadiness** | Pointer to **string** | Cluster readiness of the Fabric Interconnect. | [optional] [readonly] 
**ClusterState** | Pointer to **string** | Cluster state of the Fabric Interconnect. | [optional] [readonly] 
**EntityId** | Pointer to **string** | Identity of the Fabric Interconnect - A/B. | [optional] [readonly] 
**Leadership** | Pointer to **string** | Role (Primary / Subordinate) of the Fabric Interconnect. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewManagementEntity

`func NewManagementEntity() *ManagementEntity`

NewManagementEntity instantiates a new ManagementEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementEntityWithDefaults

`func NewManagementEntityWithDefaults() *ManagementEntity`

NewManagementEntityWithDefaults instantiates a new ManagementEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterLinkState

`func (o *ManagementEntity) GetClusterLinkState() string`

GetClusterLinkState returns the ClusterLinkState field if non-nil, zero value otherwise.

### GetClusterLinkStateOk

`func (o *ManagementEntity) GetClusterLinkStateOk() (*string, bool)`

GetClusterLinkStateOk returns a tuple with the ClusterLinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterLinkState

`func (o *ManagementEntity) SetClusterLinkState(v string)`

SetClusterLinkState sets ClusterLinkState field to given value.

### HasClusterLinkState

`func (o *ManagementEntity) HasClusterLinkState() bool`

HasClusterLinkState returns a boolean if a field has been set.

### GetClusterReadiness

`func (o *ManagementEntity) GetClusterReadiness() string`

GetClusterReadiness returns the ClusterReadiness field if non-nil, zero value otherwise.

### GetClusterReadinessOk

`func (o *ManagementEntity) GetClusterReadinessOk() (*string, bool)`

GetClusterReadinessOk returns a tuple with the ClusterReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReadiness

`func (o *ManagementEntity) SetClusterReadiness(v string)`

SetClusterReadiness sets ClusterReadiness field to given value.

### HasClusterReadiness

`func (o *ManagementEntity) HasClusterReadiness() bool`

HasClusterReadiness returns a boolean if a field has been set.

### GetClusterState

`func (o *ManagementEntity) GetClusterState() string`

GetClusterState returns the ClusterState field if non-nil, zero value otherwise.

### GetClusterStateOk

`func (o *ManagementEntity) GetClusterStateOk() (*string, bool)`

GetClusterStateOk returns a tuple with the ClusterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterState

`func (o *ManagementEntity) SetClusterState(v string)`

SetClusterState sets ClusterState field to given value.

### HasClusterState

`func (o *ManagementEntity) HasClusterState() bool`

HasClusterState returns a boolean if a field has been set.

### GetEntityId

`func (o *ManagementEntity) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ManagementEntity) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ManagementEntity) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ManagementEntity) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetLeadership

`func (o *ManagementEntity) GetLeadership() string`

GetLeadership returns the Leadership field if non-nil, zero value otherwise.

### GetLeadershipOk

`func (o *ManagementEntity) GetLeadershipOk() (*string, bool)`

GetLeadershipOk returns a tuple with the Leadership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadership

`func (o *ManagementEntity) SetLeadership(v string)`

SetLeadership sets Leadership field to given value.

### HasLeadership

`func (o *ManagementEntity) HasLeadership() bool`

HasLeadership returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ManagementEntity) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ManagementEntity) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ManagementEntity) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ManagementEntity) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *ManagementEntity) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *ManagementEntity) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *ManagementEntity) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *ManagementEntity) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ManagementEntity) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ManagementEntity) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ManagementEntity) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ManagementEntity) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


