# ManagementEntityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "management.Entity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "management.Entity"]
**ClusterLinkState** | Pointer to **string** | Cluster link state between the Fabric Interconnects. | [optional] [readonly] 
**ClusterReadiness** | Pointer to **string** | Cluster readiness of the Fabric Interconnect. | [optional] [readonly] 
**ClusterState** | Pointer to **string** | Cluster state of the Fabric Interconnect. | [optional] [readonly] 
**EntityId** | Pointer to **string** | Identity of the Fabric Interconnect - A/B. | [optional] [readonly] 
**Leadership** | Pointer to **string** | Role (Primary / Subordinate) of the Fabric Interconnect. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewManagementEntityAllOf

`func NewManagementEntityAllOf(classId string, objectType string, ) *ManagementEntityAllOf`

NewManagementEntityAllOf instantiates a new ManagementEntityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementEntityAllOfWithDefaults

`func NewManagementEntityAllOfWithDefaults() *ManagementEntityAllOf`

NewManagementEntityAllOfWithDefaults instantiates a new ManagementEntityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ManagementEntityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ManagementEntityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ManagementEntityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ManagementEntityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ManagementEntityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ManagementEntityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterLinkState

`func (o *ManagementEntityAllOf) GetClusterLinkState() string`

GetClusterLinkState returns the ClusterLinkState field if non-nil, zero value otherwise.

### GetClusterLinkStateOk

`func (o *ManagementEntityAllOf) GetClusterLinkStateOk() (*string, bool)`

GetClusterLinkStateOk returns a tuple with the ClusterLinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterLinkState

`func (o *ManagementEntityAllOf) SetClusterLinkState(v string)`

SetClusterLinkState sets ClusterLinkState field to given value.

### HasClusterLinkState

`func (o *ManagementEntityAllOf) HasClusterLinkState() bool`

HasClusterLinkState returns a boolean if a field has been set.

### GetClusterReadiness

`func (o *ManagementEntityAllOf) GetClusterReadiness() string`

GetClusterReadiness returns the ClusterReadiness field if non-nil, zero value otherwise.

### GetClusterReadinessOk

`func (o *ManagementEntityAllOf) GetClusterReadinessOk() (*string, bool)`

GetClusterReadinessOk returns a tuple with the ClusterReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReadiness

`func (o *ManagementEntityAllOf) SetClusterReadiness(v string)`

SetClusterReadiness sets ClusterReadiness field to given value.

### HasClusterReadiness

`func (o *ManagementEntityAllOf) HasClusterReadiness() bool`

HasClusterReadiness returns a boolean if a field has been set.

### GetClusterState

`func (o *ManagementEntityAllOf) GetClusterState() string`

GetClusterState returns the ClusterState field if non-nil, zero value otherwise.

### GetClusterStateOk

`func (o *ManagementEntityAllOf) GetClusterStateOk() (*string, bool)`

GetClusterStateOk returns a tuple with the ClusterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterState

`func (o *ManagementEntityAllOf) SetClusterState(v string)`

SetClusterState sets ClusterState field to given value.

### HasClusterState

`func (o *ManagementEntityAllOf) HasClusterState() bool`

HasClusterState returns a boolean if a field has been set.

### GetEntityId

`func (o *ManagementEntityAllOf) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ManagementEntityAllOf) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ManagementEntityAllOf) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ManagementEntityAllOf) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetLeadership

`func (o *ManagementEntityAllOf) GetLeadership() string`

GetLeadership returns the Leadership field if non-nil, zero value otherwise.

### GetLeadershipOk

`func (o *ManagementEntityAllOf) GetLeadershipOk() (*string, bool)`

GetLeadershipOk returns a tuple with the Leadership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadership

`func (o *ManagementEntityAllOf) SetLeadership(v string)`

SetLeadership sets Leadership field to given value.

### HasLeadership

`func (o *ManagementEntityAllOf) HasLeadership() bool`

HasLeadership returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ManagementEntityAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ManagementEntityAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ManagementEntityAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ManagementEntityAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *ManagementEntityAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *ManagementEntityAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *ManagementEntityAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *ManagementEntityAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ManagementEntityAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ManagementEntityAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ManagementEntityAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ManagementEntityAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


