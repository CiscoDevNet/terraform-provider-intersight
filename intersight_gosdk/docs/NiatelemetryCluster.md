# NiatelemetryCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Cluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Cluster"]
**ConnectionState** | Pointer to **string** | The current Intersight connection state of the cluster. | [optional] [readonly] 
**Name** | Pointer to **string** | Returns the name of the fabric. | [optional] [readonly] 
**Fabrics** | Pointer to [**[]NiatelemetryFabricRelationship**](NiatelemetryFabricRelationship.md) | An array of relationships to niatelemetryFabric resources. | [optional] [readonly] 
**Nodes** | Pointer to [**[]NiatelemetryClusterNodeRelationship**](NiatelemetryClusterNodeRelationship.md) | An array of relationships to niatelemetryClusterNode resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryCluster

`func NewNiatelemetryCluster(classId string, objectType string, ) *NiatelemetryCluster`

NewNiatelemetryCluster instantiates a new NiatelemetryCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryClusterWithDefaults

`func NewNiatelemetryClusterWithDefaults() *NiatelemetryCluster`

NewNiatelemetryClusterWithDefaults instantiates a new NiatelemetryCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionState

`func (o *NiatelemetryCluster) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *NiatelemetryCluster) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *NiatelemetryCluster) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *NiatelemetryCluster) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFabrics

`func (o *NiatelemetryCluster) GetFabrics() []NiatelemetryFabricRelationship`

GetFabrics returns the Fabrics field if non-nil, zero value otherwise.

### GetFabricsOk

`func (o *NiatelemetryCluster) GetFabricsOk() (*[]NiatelemetryFabricRelationship, bool)`

GetFabricsOk returns a tuple with the Fabrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabrics

`func (o *NiatelemetryCluster) SetFabrics(v []NiatelemetryFabricRelationship)`

SetFabrics sets Fabrics field to given value.

### HasFabrics

`func (o *NiatelemetryCluster) HasFabrics() bool`

HasFabrics returns a boolean if a field has been set.

### SetFabricsNil

`func (o *NiatelemetryCluster) SetFabricsNil(b bool)`

 SetFabricsNil sets the value for Fabrics to be an explicit nil

### UnsetFabrics
`func (o *NiatelemetryCluster) UnsetFabrics()`

UnsetFabrics ensures that no value is present for Fabrics, not even an explicit nil
### GetNodes

`func (o *NiatelemetryCluster) GetNodes() []NiatelemetryClusterNodeRelationship`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NiatelemetryCluster) GetNodesOk() (*[]NiatelemetryClusterNodeRelationship, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NiatelemetryCluster) SetNodes(v []NiatelemetryClusterNodeRelationship)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *NiatelemetryCluster) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *NiatelemetryCluster) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *NiatelemetryCluster) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetRegisteredDevice

`func (o *NiatelemetryCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryCluster) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryCluster) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryCluster) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


