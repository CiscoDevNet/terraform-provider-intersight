# HyperflexNetworkConfiguration

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

### NewHyperflexNetworkConfiguration

`func NewHyperflexNetworkConfiguration(classId string, objectType string, ) *HyperflexNetworkConfiguration`

NewHyperflexNetworkConfiguration instantiates a new HyperflexNetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNetworkConfigurationWithDefaults

`func NewHyperflexNetworkConfigurationWithDefaults() *HyperflexNetworkConfiguration`

NewHyperflexNetworkConfigurationWithDefaults instantiates a new HyperflexNetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexNetworkConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNetworkConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNetworkConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexNetworkConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNetworkConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNetworkConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterDataIp

`func (o *HyperflexNetworkConfiguration) GetClusterDataIp() string`

GetClusterDataIp returns the ClusterDataIp field if non-nil, zero value otherwise.

### GetClusterDataIpOk

`func (o *HyperflexNetworkConfiguration) GetClusterDataIpOk() (*string, bool)`

GetClusterDataIpOk returns a tuple with the ClusterDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDataIp

`func (o *HyperflexNetworkConfiguration) SetClusterDataIp(v string)`

SetClusterDataIp sets ClusterDataIp field to given value.

### HasClusterDataIp

`func (o *HyperflexNetworkConfiguration) HasClusterDataIp() bool`

HasClusterDataIp returns a boolean if a field has been set.

### GetClusterManagementIp

`func (o *HyperflexNetworkConfiguration) GetClusterManagementIp() string`

GetClusterManagementIp returns the ClusterManagementIp field if non-nil, zero value otherwise.

### GetClusterManagementIpOk

`func (o *HyperflexNetworkConfiguration) GetClusterManagementIpOk() (*string, bool)`

GetClusterManagementIpOk returns a tuple with the ClusterManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterManagementIp

`func (o *HyperflexNetworkConfiguration) SetClusterManagementIp(v string)`

SetClusterManagementIp sets ClusterManagementIp field to given value.

### HasClusterManagementIp

`func (o *HyperflexNetworkConfiguration) HasClusterManagementIp() bool`

HasClusterManagementIp returns a boolean if a field has been set.

### GetDataDefaultGateway

`func (o *HyperflexNetworkConfiguration) GetDataDefaultGateway() string`

GetDataDefaultGateway returns the DataDefaultGateway field if non-nil, zero value otherwise.

### GetDataDefaultGatewayOk

`func (o *HyperflexNetworkConfiguration) GetDataDefaultGatewayOk() (*string, bool)`

GetDataDefaultGatewayOk returns a tuple with the DataDefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDefaultGateway

`func (o *HyperflexNetworkConfiguration) SetDataDefaultGateway(v string)`

SetDataDefaultGateway sets DataDefaultGateway field to given value.

### HasDataDefaultGateway

`func (o *HyperflexNetworkConfiguration) HasDataDefaultGateway() bool`

HasDataDefaultGateway returns a boolean if a field has been set.

### GetDataJumboFrame

`func (o *HyperflexNetworkConfiguration) GetDataJumboFrame() bool`

GetDataJumboFrame returns the DataJumboFrame field if non-nil, zero value otherwise.

### GetDataJumboFrameOk

`func (o *HyperflexNetworkConfiguration) GetDataJumboFrameOk() (*bool, bool)`

GetDataJumboFrameOk returns a tuple with the DataJumboFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataJumboFrame

`func (o *HyperflexNetworkConfiguration) SetDataJumboFrame(v bool)`

SetDataJumboFrame sets DataJumboFrame field to given value.

### HasDataJumboFrame

`func (o *HyperflexNetworkConfiguration) HasDataJumboFrame() bool`

HasDataJumboFrame returns a boolean if a field has been set.

### GetDataSubNetmask

`func (o *HyperflexNetworkConfiguration) GetDataSubNetmask() string`

GetDataSubNetmask returns the DataSubNetmask field if non-nil, zero value otherwise.

### GetDataSubNetmaskOk

`func (o *HyperflexNetworkConfiguration) GetDataSubNetmaskOk() (*string, bool)`

GetDataSubNetmaskOk returns a tuple with the DataSubNetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSubNetmask

`func (o *HyperflexNetworkConfiguration) SetDataSubNetmask(v string)`

SetDataSubNetmask sets DataSubNetmask field to given value.

### HasDataSubNetmask

`func (o *HyperflexNetworkConfiguration) HasDataSubNetmask() bool`

HasDataSubNetmask returns a boolean if a field has been set.

### GetDataVlanId

`func (o *HyperflexNetworkConfiguration) GetDataVlanId() int64`

GetDataVlanId returns the DataVlanId field if non-nil, zero value otherwise.

### GetDataVlanIdOk

`func (o *HyperflexNetworkConfiguration) GetDataVlanIdOk() (*int64, bool)`

GetDataVlanIdOk returns a tuple with the DataVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVlanId

`func (o *HyperflexNetworkConfiguration) SetDataVlanId(v int64)`

SetDataVlanId sets DataVlanId field to given value.

### HasDataVlanId

`func (o *HyperflexNetworkConfiguration) HasDataVlanId() bool`

HasDataVlanId returns a boolean if a field has been set.

### GetLiveMigrationVlanId

`func (o *HyperflexNetworkConfiguration) GetLiveMigrationVlanId() int64`

GetLiveMigrationVlanId returns the LiveMigrationVlanId field if non-nil, zero value otherwise.

### GetLiveMigrationVlanIdOk

`func (o *HyperflexNetworkConfiguration) GetLiveMigrationVlanIdOk() (*int64, bool)`

GetLiveMigrationVlanIdOk returns a tuple with the LiveMigrationVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveMigrationVlanId

`func (o *HyperflexNetworkConfiguration) SetLiveMigrationVlanId(v int64)`

SetLiveMigrationVlanId sets LiveMigrationVlanId field to given value.

### HasLiveMigrationVlanId

`func (o *HyperflexNetworkConfiguration) HasLiveMigrationVlanId() bool`

HasLiveMigrationVlanId returns a boolean if a field has been set.

### GetManagementDefaultGateway

`func (o *HyperflexNetworkConfiguration) GetManagementDefaultGateway() string`

GetManagementDefaultGateway returns the ManagementDefaultGateway field if non-nil, zero value otherwise.

### GetManagementDefaultGatewayOk

`func (o *HyperflexNetworkConfiguration) GetManagementDefaultGatewayOk() (*string, bool)`

GetManagementDefaultGatewayOk returns a tuple with the ManagementDefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDefaultGateway

`func (o *HyperflexNetworkConfiguration) SetManagementDefaultGateway(v string)`

SetManagementDefaultGateway sets ManagementDefaultGateway field to given value.

### HasManagementDefaultGateway

`func (o *HyperflexNetworkConfiguration) HasManagementDefaultGateway() bool`

HasManagementDefaultGateway returns a boolean if a field has been set.

### GetManagementSubNetmask

`func (o *HyperflexNetworkConfiguration) GetManagementSubNetmask() string`

GetManagementSubNetmask returns the ManagementSubNetmask field if non-nil, zero value otherwise.

### GetManagementSubNetmaskOk

`func (o *HyperflexNetworkConfiguration) GetManagementSubNetmaskOk() (*string, bool)`

GetManagementSubNetmaskOk returns a tuple with the ManagementSubNetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementSubNetmask

`func (o *HyperflexNetworkConfiguration) SetManagementSubNetmask(v string)`

SetManagementSubNetmask sets ManagementSubNetmask field to given value.

### HasManagementSubNetmask

`func (o *HyperflexNetworkConfiguration) HasManagementSubNetmask() bool`

HasManagementSubNetmask returns a boolean if a field has been set.

### GetManagementVlanId

`func (o *HyperflexNetworkConfiguration) GetManagementVlanId() int64`

GetManagementVlanId returns the ManagementVlanId field if non-nil, zero value otherwise.

### GetManagementVlanIdOk

`func (o *HyperflexNetworkConfiguration) GetManagementVlanIdOk() (*int64, bool)`

GetManagementVlanIdOk returns a tuple with the ManagementVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementVlanId

`func (o *HyperflexNetworkConfiguration) SetManagementVlanId(v int64)`

SetManagementVlanId sets ManagementVlanId field to given value.

### HasManagementVlanId

`func (o *HyperflexNetworkConfiguration) HasManagementVlanId() bool`

HasManagementVlanId returns a boolean if a field has been set.

### GetVmNetworkVlanId

`func (o *HyperflexNetworkConfiguration) GetVmNetworkVlanId() int64`

GetVmNetworkVlanId returns the VmNetworkVlanId field if non-nil, zero value otherwise.

### GetVmNetworkVlanIdOk

`func (o *HyperflexNetworkConfiguration) GetVmNetworkVlanIdOk() (*int64, bool)`

GetVmNetworkVlanIdOk returns a tuple with the VmNetworkVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetworkVlanId

`func (o *HyperflexNetworkConfiguration) SetVmNetworkVlanId(v int64)`

SetVmNetworkVlanId sets VmNetworkVlanId field to given value.

### HasVmNetworkVlanId

`func (o *HyperflexNetworkConfiguration) HasVmNetworkVlanId() bool`

HasVmNetworkVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


