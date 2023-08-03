# HyperflexNetworkConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.NetworkConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.NetworkConfiguration"]
**ClusterDataIp** | Pointer to **string** | Cluster data IP of the HyperFlex cluster. | [optional] 
**ClusterManagementIp** | Pointer to **string** | Cluster management IP of the HyperFlex cluster. | [optional] 
**DataDefaultGateway** | Pointer to **string** | Default gateway of the data network. | [optional] 
**DataJumboFrame** | Pointer to **bool** | Boolean value to indicate if jumboframes is enabled for storage-data network. | [optional] 
**DataSubNetmask** | Pointer to **string** | Subnet mask of the data network. | [optional] 
**DataVlanId** | Pointer to **int64** | Data VLAN ID. Enter the correct VLAN tags if you are using trunk ports. The VLAN tags must be different when using trunk mode. | [optional] 
**LiveMigrationVlanId** | Pointer to **int64** | VLAN ID for virtual machine live migration. | [optional] 
**ManagementDefaultGateway** | Pointer to **string** | Default gateway of the management network. | [optional] 
**ManagementSubNetmask** | Pointer to **string** | Subnet mask of the management network. | [optional] 
**ManagementVlanId** | Pointer to **int64** | Management VLAN ID. Enter the correct VLAN tags if you are using trunk ports. The VLAN tags must be different when using trunk mode. | [optional] 
**VmNetworkVlanId** | Pointer to **int64** | VM network VLAN ID. Used for VM data traffic. | [optional] 

## Methods

### NewHyperflexNetworkConfigurationAllOf

`func NewHyperflexNetworkConfigurationAllOf(classId string, objectType string, ) *HyperflexNetworkConfigurationAllOf`

NewHyperflexNetworkConfigurationAllOf instantiates a new HyperflexNetworkConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNetworkConfigurationAllOfWithDefaults

`func NewHyperflexNetworkConfigurationAllOfWithDefaults() *HyperflexNetworkConfigurationAllOf`

NewHyperflexNetworkConfigurationAllOfWithDefaults instantiates a new HyperflexNetworkConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexNetworkConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNetworkConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNetworkConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexNetworkConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNetworkConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNetworkConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterDataIp

`func (o *HyperflexNetworkConfigurationAllOf) GetClusterDataIp() string`

GetClusterDataIp returns the ClusterDataIp field if non-nil, zero value otherwise.

### GetClusterDataIpOk

`func (o *HyperflexNetworkConfigurationAllOf) GetClusterDataIpOk() (*string, bool)`

GetClusterDataIpOk returns a tuple with the ClusterDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDataIp

`func (o *HyperflexNetworkConfigurationAllOf) SetClusterDataIp(v string)`

SetClusterDataIp sets ClusterDataIp field to given value.

### HasClusterDataIp

`func (o *HyperflexNetworkConfigurationAllOf) HasClusterDataIp() bool`

HasClusterDataIp returns a boolean if a field has been set.

### GetClusterManagementIp

`func (o *HyperflexNetworkConfigurationAllOf) GetClusterManagementIp() string`

GetClusterManagementIp returns the ClusterManagementIp field if non-nil, zero value otherwise.

### GetClusterManagementIpOk

`func (o *HyperflexNetworkConfigurationAllOf) GetClusterManagementIpOk() (*string, bool)`

GetClusterManagementIpOk returns a tuple with the ClusterManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterManagementIp

`func (o *HyperflexNetworkConfigurationAllOf) SetClusterManagementIp(v string)`

SetClusterManagementIp sets ClusterManagementIp field to given value.

### HasClusterManagementIp

`func (o *HyperflexNetworkConfigurationAllOf) HasClusterManagementIp() bool`

HasClusterManagementIp returns a boolean if a field has been set.

### GetDataDefaultGateway

`func (o *HyperflexNetworkConfigurationAllOf) GetDataDefaultGateway() string`

GetDataDefaultGateway returns the DataDefaultGateway field if non-nil, zero value otherwise.

### GetDataDefaultGatewayOk

`func (o *HyperflexNetworkConfigurationAllOf) GetDataDefaultGatewayOk() (*string, bool)`

GetDataDefaultGatewayOk returns a tuple with the DataDefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDefaultGateway

`func (o *HyperflexNetworkConfigurationAllOf) SetDataDefaultGateway(v string)`

SetDataDefaultGateway sets DataDefaultGateway field to given value.

### HasDataDefaultGateway

`func (o *HyperflexNetworkConfigurationAllOf) HasDataDefaultGateway() bool`

HasDataDefaultGateway returns a boolean if a field has been set.

### GetDataJumboFrame

`func (o *HyperflexNetworkConfigurationAllOf) GetDataJumboFrame() bool`

GetDataJumboFrame returns the DataJumboFrame field if non-nil, zero value otherwise.

### GetDataJumboFrameOk

`func (o *HyperflexNetworkConfigurationAllOf) GetDataJumboFrameOk() (*bool, bool)`

GetDataJumboFrameOk returns a tuple with the DataJumboFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataJumboFrame

`func (o *HyperflexNetworkConfigurationAllOf) SetDataJumboFrame(v bool)`

SetDataJumboFrame sets DataJumboFrame field to given value.

### HasDataJumboFrame

`func (o *HyperflexNetworkConfigurationAllOf) HasDataJumboFrame() bool`

HasDataJumboFrame returns a boolean if a field has been set.

### GetDataSubNetmask

`func (o *HyperflexNetworkConfigurationAllOf) GetDataSubNetmask() string`

GetDataSubNetmask returns the DataSubNetmask field if non-nil, zero value otherwise.

### GetDataSubNetmaskOk

`func (o *HyperflexNetworkConfigurationAllOf) GetDataSubNetmaskOk() (*string, bool)`

GetDataSubNetmaskOk returns a tuple with the DataSubNetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSubNetmask

`func (o *HyperflexNetworkConfigurationAllOf) SetDataSubNetmask(v string)`

SetDataSubNetmask sets DataSubNetmask field to given value.

### HasDataSubNetmask

`func (o *HyperflexNetworkConfigurationAllOf) HasDataSubNetmask() bool`

HasDataSubNetmask returns a boolean if a field has been set.

### GetDataVlanId

`func (o *HyperflexNetworkConfigurationAllOf) GetDataVlanId() int64`

GetDataVlanId returns the DataVlanId field if non-nil, zero value otherwise.

### GetDataVlanIdOk

`func (o *HyperflexNetworkConfigurationAllOf) GetDataVlanIdOk() (*int64, bool)`

GetDataVlanIdOk returns a tuple with the DataVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVlanId

`func (o *HyperflexNetworkConfigurationAllOf) SetDataVlanId(v int64)`

SetDataVlanId sets DataVlanId field to given value.

### HasDataVlanId

`func (o *HyperflexNetworkConfigurationAllOf) HasDataVlanId() bool`

HasDataVlanId returns a boolean if a field has been set.

### GetLiveMigrationVlanId

`func (o *HyperflexNetworkConfigurationAllOf) GetLiveMigrationVlanId() int64`

GetLiveMigrationVlanId returns the LiveMigrationVlanId field if non-nil, zero value otherwise.

### GetLiveMigrationVlanIdOk

`func (o *HyperflexNetworkConfigurationAllOf) GetLiveMigrationVlanIdOk() (*int64, bool)`

GetLiveMigrationVlanIdOk returns a tuple with the LiveMigrationVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveMigrationVlanId

`func (o *HyperflexNetworkConfigurationAllOf) SetLiveMigrationVlanId(v int64)`

SetLiveMigrationVlanId sets LiveMigrationVlanId field to given value.

### HasLiveMigrationVlanId

`func (o *HyperflexNetworkConfigurationAllOf) HasLiveMigrationVlanId() bool`

HasLiveMigrationVlanId returns a boolean if a field has been set.

### GetManagementDefaultGateway

`func (o *HyperflexNetworkConfigurationAllOf) GetManagementDefaultGateway() string`

GetManagementDefaultGateway returns the ManagementDefaultGateway field if non-nil, zero value otherwise.

### GetManagementDefaultGatewayOk

`func (o *HyperflexNetworkConfigurationAllOf) GetManagementDefaultGatewayOk() (*string, bool)`

GetManagementDefaultGatewayOk returns a tuple with the ManagementDefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDefaultGateway

`func (o *HyperflexNetworkConfigurationAllOf) SetManagementDefaultGateway(v string)`

SetManagementDefaultGateway sets ManagementDefaultGateway field to given value.

### HasManagementDefaultGateway

`func (o *HyperflexNetworkConfigurationAllOf) HasManagementDefaultGateway() bool`

HasManagementDefaultGateway returns a boolean if a field has been set.

### GetManagementSubNetmask

`func (o *HyperflexNetworkConfigurationAllOf) GetManagementSubNetmask() string`

GetManagementSubNetmask returns the ManagementSubNetmask field if non-nil, zero value otherwise.

### GetManagementSubNetmaskOk

`func (o *HyperflexNetworkConfigurationAllOf) GetManagementSubNetmaskOk() (*string, bool)`

GetManagementSubNetmaskOk returns a tuple with the ManagementSubNetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementSubNetmask

`func (o *HyperflexNetworkConfigurationAllOf) SetManagementSubNetmask(v string)`

SetManagementSubNetmask sets ManagementSubNetmask field to given value.

### HasManagementSubNetmask

`func (o *HyperflexNetworkConfigurationAllOf) HasManagementSubNetmask() bool`

HasManagementSubNetmask returns a boolean if a field has been set.

### GetManagementVlanId

`func (o *HyperflexNetworkConfigurationAllOf) GetManagementVlanId() int64`

GetManagementVlanId returns the ManagementVlanId field if non-nil, zero value otherwise.

### GetManagementVlanIdOk

`func (o *HyperflexNetworkConfigurationAllOf) GetManagementVlanIdOk() (*int64, bool)`

GetManagementVlanIdOk returns a tuple with the ManagementVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementVlanId

`func (o *HyperflexNetworkConfigurationAllOf) SetManagementVlanId(v int64)`

SetManagementVlanId sets ManagementVlanId field to given value.

### HasManagementVlanId

`func (o *HyperflexNetworkConfigurationAllOf) HasManagementVlanId() bool`

HasManagementVlanId returns a boolean if a field has been set.

### GetVmNetworkVlanId

`func (o *HyperflexNetworkConfigurationAllOf) GetVmNetworkVlanId() int64`

GetVmNetworkVlanId returns the VmNetworkVlanId field if non-nil, zero value otherwise.

### GetVmNetworkVlanIdOk

`func (o *HyperflexNetworkConfigurationAllOf) GetVmNetworkVlanIdOk() (*int64, bool)`

GetVmNetworkVlanIdOk returns a tuple with the VmNetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetworkVlanId

`func (o *HyperflexNetworkConfigurationAllOf) SetVmNetworkVlanId(v int64)`

SetVmNetworkVlanId sets VmNetworkVlanId field to given value.

### HasVmNetworkVlanId

`func (o *HyperflexNetworkConfigurationAllOf) HasVmNetworkVlanId() bool`

HasVmNetworkVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


