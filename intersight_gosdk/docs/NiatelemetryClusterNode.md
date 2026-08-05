# NiatelemetryClusterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ClusterNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ClusterNode"]
**BootTime** | Pointer to **string** | The boot time of the node in the cluster. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | The name of the cluster this node belongs to. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | The firmware version of the node in the cluster. | [optional] [readonly] 
**ManagementIp** | Pointer to **string** | The management IP address of the node in the cluster. | [optional] [readonly] 
**Model** | Pointer to **string** | The model of the node in the cluster. | [optional] [readonly] 
**Name** | Pointer to **string** | Returns the name of the node. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type of the node in the cluster. | [optional] [readonly] 
**Role** | Pointer to **string** | Role of the node in the cluster. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | Serial number of the node in the cluster. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the node in the cluster. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableNiatelemetryClusterRelationship**](NiatelemetryClusterRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryClusterNode

`func NewNiatelemetryClusterNode(classId string, objectType string, ) *NiatelemetryClusterNode`

NewNiatelemetryClusterNode instantiates a new NiatelemetryClusterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryClusterNodeWithDefaults

`func NewNiatelemetryClusterNodeWithDefaults() *NiatelemetryClusterNode`

NewNiatelemetryClusterNodeWithDefaults instantiates a new NiatelemetryClusterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryClusterNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryClusterNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryClusterNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryClusterNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryClusterNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryClusterNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootTime

`func (o *NiatelemetryClusterNode) GetBootTime() string`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *NiatelemetryClusterNode) GetBootTimeOk() (*string, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *NiatelemetryClusterNode) SetBootTime(v string)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *NiatelemetryClusterNode) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetClusterName

`func (o *NiatelemetryClusterNode) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *NiatelemetryClusterNode) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *NiatelemetryClusterNode) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *NiatelemetryClusterNode) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *NiatelemetryClusterNode) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *NiatelemetryClusterNode) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *NiatelemetryClusterNode) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *NiatelemetryClusterNode) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetManagementIp

`func (o *NiatelemetryClusterNode) GetManagementIp() string`

GetManagementIp returns the ManagementIp field if non-nil, zero value otherwise.

### GetManagementIpOk

`func (o *NiatelemetryClusterNode) GetManagementIpOk() (*string, bool)`

GetManagementIpOk returns a tuple with the ManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIp

`func (o *NiatelemetryClusterNode) SetManagementIp(v string)`

SetManagementIp sets ManagementIp field to given value.

### HasManagementIp

`func (o *NiatelemetryClusterNode) HasManagementIp() bool`

HasManagementIp returns a boolean if a field has been set.

### GetModel

`func (o *NiatelemetryClusterNode) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NiatelemetryClusterNode) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NiatelemetryClusterNode) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NiatelemetryClusterNode) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryClusterNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryClusterNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryClusterNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryClusterNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformType

`func (o *NiatelemetryClusterNode) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *NiatelemetryClusterNode) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *NiatelemetryClusterNode) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *NiatelemetryClusterNode) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRole

`func (o *NiatelemetryClusterNode) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NiatelemetryClusterNode) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NiatelemetryClusterNode) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NiatelemetryClusterNode) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryClusterNode) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryClusterNode) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryClusterNode) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryClusterNode) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *NiatelemetryClusterNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NiatelemetryClusterNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NiatelemetryClusterNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NiatelemetryClusterNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCluster

`func (o *NiatelemetryClusterNode) GetCluster() NiatelemetryClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NiatelemetryClusterNode) GetClusterOk() (*NiatelemetryClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NiatelemetryClusterNode) SetCluster(v NiatelemetryClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NiatelemetryClusterNode) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *NiatelemetryClusterNode) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *NiatelemetryClusterNode) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


