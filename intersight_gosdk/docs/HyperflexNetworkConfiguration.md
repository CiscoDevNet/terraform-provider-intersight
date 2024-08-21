# HyperflexNetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.NetworkConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.NetworkConfiguration"]
**DataGatewayIpAddress** | Pointer to **string** | The data gateway IP of the HyperFlex cluster. | [optional] [readonly] 
**DataIpAddress** | Pointer to **string** | The data IP of the HyperFlex cluster. | [optional] [readonly] 
**DataNetmask** | Pointer to **string** | The data subnet mask of the HyperFlex cluster. | [optional] [readonly] 
**DataVlan** | Pointer to **int64** | The data VLAN of the HyperFlex cluster. | [optional] [readonly] 
**DnsSuffix** | Pointer to **string** | The DNS domain suffix configured for the HyperFlex Cluster. | [optional] [readonly] 
**JumboFrameEnabled** | Pointer to **bool** | The jumbo frame enablement of the HyperFlex cluster. | [optional] [readonly] 
**LiveMigrationVlan** | Pointer to **int64** | The live migration VLAN ID of the HyperFlex cluster. | [optional] [readonly] 
**MgmtGatewayIpAddress** | Pointer to **string** | The management gateway IP of the HyperFlex cluster. | [optional] [readonly] 
**MgmtIpAddress** | Pointer to **string** | The management IP or the hostname of the HyperFlex cluster. | [optional] [readonly] 
**MgmtNetmask** | Pointer to **string** | The management subnet mask of the HyperFlex cluster. | [optional] [readonly] 
**MgmtVlan** | Pointer to **int64** | The management VLAN ID of the HyperFlex cluster. | [optional] [readonly] 
**Timezone** | Pointer to **string** | The timezone configured on the HyperFlex Cluster. | [optional] [readonly] 
**VmNetworkVlans** | Pointer to **[]int64** |  | [optional] 

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


### GetDataGatewayIpAddress

`func (o *HyperflexNetworkConfiguration) GetDataGatewayIpAddress() string`

GetDataGatewayIpAddress returns the DataGatewayIpAddress field if non-nil, zero value otherwise.

### GetDataGatewayIpAddressOk

`func (o *HyperflexNetworkConfiguration) GetDataGatewayIpAddressOk() (*string, bool)`

GetDataGatewayIpAddressOk returns a tuple with the DataGatewayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataGatewayIpAddress

`func (o *HyperflexNetworkConfiguration) SetDataGatewayIpAddress(v string)`

SetDataGatewayIpAddress sets DataGatewayIpAddress field to given value.

### HasDataGatewayIpAddress

`func (o *HyperflexNetworkConfiguration) HasDataGatewayIpAddress() bool`

HasDataGatewayIpAddress returns a boolean if a field has been set.

### GetDataIpAddress

`func (o *HyperflexNetworkConfiguration) GetDataIpAddress() string`

GetDataIpAddress returns the DataIpAddress field if non-nil, zero value otherwise.

### GetDataIpAddressOk

`func (o *HyperflexNetworkConfiguration) GetDataIpAddressOk() (*string, bool)`

GetDataIpAddressOk returns a tuple with the DataIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIpAddress

`func (o *HyperflexNetworkConfiguration) SetDataIpAddress(v string)`

SetDataIpAddress sets DataIpAddress field to given value.

### HasDataIpAddress

`func (o *HyperflexNetworkConfiguration) HasDataIpAddress() bool`

HasDataIpAddress returns a boolean if a field has been set.

### GetDataNetmask

`func (o *HyperflexNetworkConfiguration) GetDataNetmask() string`

GetDataNetmask returns the DataNetmask field if non-nil, zero value otherwise.

### GetDataNetmaskOk

`func (o *HyperflexNetworkConfiguration) GetDataNetmaskOk() (*string, bool)`

GetDataNetmaskOk returns a tuple with the DataNetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNetmask

`func (o *HyperflexNetworkConfiguration) SetDataNetmask(v string)`

SetDataNetmask sets DataNetmask field to given value.

### HasDataNetmask

`func (o *HyperflexNetworkConfiguration) HasDataNetmask() bool`

HasDataNetmask returns a boolean if a field has been set.

### GetDataVlan

`func (o *HyperflexNetworkConfiguration) GetDataVlan() int64`

GetDataVlan returns the DataVlan field if non-nil, zero value otherwise.

### GetDataVlanOk

`func (o *HyperflexNetworkConfiguration) GetDataVlanOk() (*int64, bool)`

GetDataVlanOk returns a tuple with the DataVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVlan

`func (o *HyperflexNetworkConfiguration) SetDataVlan(v int64)`

SetDataVlan sets DataVlan field to given value.

### HasDataVlan

`func (o *HyperflexNetworkConfiguration) HasDataVlan() bool`

HasDataVlan returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *HyperflexNetworkConfiguration) GetDnsSuffix() string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *HyperflexNetworkConfiguration) GetDnsSuffixOk() (*string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *HyperflexNetworkConfiguration) SetDnsSuffix(v string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *HyperflexNetworkConfiguration) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetJumboFrameEnabled

`func (o *HyperflexNetworkConfiguration) GetJumboFrameEnabled() bool`

GetJumboFrameEnabled returns the JumboFrameEnabled field if non-nil, zero value otherwise.

### GetJumboFrameEnabledOk

`func (o *HyperflexNetworkConfiguration) GetJumboFrameEnabledOk() (*bool, bool)`

GetJumboFrameEnabledOk returns a tuple with the JumboFrameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboFrameEnabled

`func (o *HyperflexNetworkConfiguration) SetJumboFrameEnabled(v bool)`

SetJumboFrameEnabled sets JumboFrameEnabled field to given value.

### HasJumboFrameEnabled

`func (o *HyperflexNetworkConfiguration) HasJumboFrameEnabled() bool`

HasJumboFrameEnabled returns a boolean if a field has been set.

### GetLiveMigrationVlan

`func (o *HyperflexNetworkConfiguration) GetLiveMigrationVlan() int64`

GetLiveMigrationVlan returns the LiveMigrationVlan field if non-nil, zero value otherwise.

### GetLiveMigrationVlanOk

`func (o *HyperflexNetworkConfiguration) GetLiveMigrationVlanOk() (*int64, bool)`

GetLiveMigrationVlanOk returns a tuple with the LiveMigrationVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveMigrationVlan

`func (o *HyperflexNetworkConfiguration) SetLiveMigrationVlan(v int64)`

SetLiveMigrationVlan sets LiveMigrationVlan field to given value.

### HasLiveMigrationVlan

`func (o *HyperflexNetworkConfiguration) HasLiveMigrationVlan() bool`

HasLiveMigrationVlan returns a boolean if a field has been set.

### GetMgmtGatewayIpAddress

`func (o *HyperflexNetworkConfiguration) GetMgmtGatewayIpAddress() string`

GetMgmtGatewayIpAddress returns the MgmtGatewayIpAddress field if non-nil, zero value otherwise.

### GetMgmtGatewayIpAddressOk

`func (o *HyperflexNetworkConfiguration) GetMgmtGatewayIpAddressOk() (*string, bool)`

GetMgmtGatewayIpAddressOk returns a tuple with the MgmtGatewayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtGatewayIpAddress

`func (o *HyperflexNetworkConfiguration) SetMgmtGatewayIpAddress(v string)`

SetMgmtGatewayIpAddress sets MgmtGatewayIpAddress field to given value.

### HasMgmtGatewayIpAddress

`func (o *HyperflexNetworkConfiguration) HasMgmtGatewayIpAddress() bool`

HasMgmtGatewayIpAddress returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *HyperflexNetworkConfiguration) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *HyperflexNetworkConfiguration) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *HyperflexNetworkConfiguration) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *HyperflexNetworkConfiguration) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetMgmtNetmask

`func (o *HyperflexNetworkConfiguration) GetMgmtNetmask() string`

GetMgmtNetmask returns the MgmtNetmask field if non-nil, zero value otherwise.

### GetMgmtNetmaskOk

`func (o *HyperflexNetworkConfiguration) GetMgmtNetmaskOk() (*string, bool)`

GetMgmtNetmaskOk returns a tuple with the MgmtNetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtNetmask

`func (o *HyperflexNetworkConfiguration) SetMgmtNetmask(v string)`

SetMgmtNetmask sets MgmtNetmask field to given value.

### HasMgmtNetmask

`func (o *HyperflexNetworkConfiguration) HasMgmtNetmask() bool`

HasMgmtNetmask returns a boolean if a field has been set.

### GetMgmtVlan

`func (o *HyperflexNetworkConfiguration) GetMgmtVlan() int64`

GetMgmtVlan returns the MgmtVlan field if non-nil, zero value otherwise.

### GetMgmtVlanOk

`func (o *HyperflexNetworkConfiguration) GetMgmtVlanOk() (*int64, bool)`

GetMgmtVlanOk returns a tuple with the MgmtVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtVlan

`func (o *HyperflexNetworkConfiguration) SetMgmtVlan(v int64)`

SetMgmtVlan sets MgmtVlan field to given value.

### HasMgmtVlan

`func (o *HyperflexNetworkConfiguration) HasMgmtVlan() bool`

HasMgmtVlan returns a boolean if a field has been set.

### GetTimezone

`func (o *HyperflexNetworkConfiguration) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *HyperflexNetworkConfiguration) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *HyperflexNetworkConfiguration) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *HyperflexNetworkConfiguration) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetVmNetworkVlans

`func (o *HyperflexNetworkConfiguration) GetVmNetworkVlans() []int64`

GetVmNetworkVlans returns the VmNetworkVlans field if non-nil, zero value otherwise.

### GetVmNetworkVlansOk

`func (o *HyperflexNetworkConfiguration) GetVmNetworkVlansOk() (*[]int64, bool)`

GetVmNetworkVlansOk returns a tuple with the VmNetworkVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetworkVlans

`func (o *HyperflexNetworkConfiguration) SetVmNetworkVlans(v []int64)`

SetVmNetworkVlans sets VmNetworkVlans field to given value.

### HasVmNetworkVlans

`func (o *HyperflexNetworkConfiguration) HasVmNetworkVlans() bool`

HasVmNetworkVlans returns a boolean if a field has been set.

### SetVmNetworkVlansNil

`func (o *HyperflexNetworkConfiguration) SetVmNetworkVlansNil(b bool)`

 SetVmNetworkVlansNil sets the value for VmNetworkVlans to be an explicit nil

### UnsetVmNetworkVlans
`func (o *HyperflexNetworkConfiguration) UnsetVmNetworkVlans()`

UnsetVmNetworkVlans ensures that no value is present for VmNetworkVlans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


