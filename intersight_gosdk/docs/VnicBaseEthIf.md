# VnicBaseEthIf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Cdn** | Pointer to [**NullableVnicCdn**](VnicCdn.md) |  | [optional] 
**FailoverEnabled** | Pointer to **bool** | Enabling failover ensures that traffic from the vNIC automatically fails over to the secondary Fabric Interconnect, in case the specified Fabric Interconnect path goes down. Failover applies only to Cisco VICs that are connected to a Fabric Interconnect cluster. | [optional] [default to false]
**PinGroupName** | Pointer to **string** | Pingroup name associated to vNIC for static pinning. LCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vNIC traffic. | [optional] 
**SriovSettings** | Pointer to [**NullableVnicSriovSettings**](VnicSriovSettings.md) |  | [optional] 
**UsnicSettings** | Pointer to [**NullableVnicUsnicSettings**](VnicUsnicSettings.md) |  | [optional] 
**VmqSettings** | Pointer to [**NullableVnicVmqSettings**](VnicVmqSettings.md) |  | [optional] 
**EthAdapterPolicy** | Pointer to [**VnicEthAdapterPolicyRelationship**](VnicEthAdapterPolicyRelationship.md) |  | [optional] 
**EthNetworkPolicy** | Pointer to [**VnicEthNetworkPolicyRelationship**](VnicEthNetworkPolicyRelationship.md) |  | [optional] 
**EthQosPolicy** | Pointer to [**VnicEthQosPolicyRelationship**](VnicEthQosPolicyRelationship.md) |  | [optional] 
**FabricEthNetworkControlPolicy** | Pointer to [**FabricEthNetworkControlPolicyRelationship**](FabricEthNetworkControlPolicyRelationship.md) |  | [optional] 
**FabricEthNetworkGroupPolicy** | Pointer to [**[]FabricEthNetworkGroupPolicyRelationship**](FabricEthNetworkGroupPolicyRelationship.md) | An array of relationships to fabricEthNetworkGroupPolicy resources. | [optional] 
**IscsiBootPolicy** | Pointer to [**VnicIscsiBootPolicyRelationship**](VnicIscsiBootPolicyRelationship.md) |  | [optional] 
**MacPool** | Pointer to [**MacpoolPoolRelationship**](MacpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewVnicBaseEthIf

`func NewVnicBaseEthIf(classId string, objectType string, ) *VnicBaseEthIf`

NewVnicBaseEthIf instantiates a new VnicBaseEthIf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicBaseEthIfWithDefaults

`func NewVnicBaseEthIfWithDefaults() *VnicBaseEthIf`

NewVnicBaseEthIfWithDefaults instantiates a new VnicBaseEthIf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicBaseEthIf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicBaseEthIf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicBaseEthIf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicBaseEthIf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicBaseEthIf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicBaseEthIf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCdn

`func (o *VnicBaseEthIf) GetCdn() VnicCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *VnicBaseEthIf) GetCdnOk() (*VnicCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *VnicBaseEthIf) SetCdn(v VnicCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *VnicBaseEthIf) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### SetCdnNil

`func (o *VnicBaseEthIf) SetCdnNil(b bool)`

 SetCdnNil sets the value for Cdn to be an explicit nil

### UnsetCdn
`func (o *VnicBaseEthIf) UnsetCdn()`

UnsetCdn ensures that no value is present for Cdn, not even an explicit nil
### GetFailoverEnabled

`func (o *VnicBaseEthIf) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *VnicBaseEthIf) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *VnicBaseEthIf) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *VnicBaseEthIf) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### GetPinGroupName

`func (o *VnicBaseEthIf) GetPinGroupName() string`

GetPinGroupName returns the PinGroupName field if non-nil, zero value otherwise.

### GetPinGroupNameOk

`func (o *VnicBaseEthIf) GetPinGroupNameOk() (*string, bool)`

GetPinGroupNameOk returns a tuple with the PinGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinGroupName

`func (o *VnicBaseEthIf) SetPinGroupName(v string)`

SetPinGroupName sets PinGroupName field to given value.

### HasPinGroupName

`func (o *VnicBaseEthIf) HasPinGroupName() bool`

HasPinGroupName returns a boolean if a field has been set.

### GetSriovSettings

`func (o *VnicBaseEthIf) GetSriovSettings() VnicSriovSettings`

GetSriovSettings returns the SriovSettings field if non-nil, zero value otherwise.

### GetSriovSettingsOk

`func (o *VnicBaseEthIf) GetSriovSettingsOk() (*VnicSriovSettings, bool)`

GetSriovSettingsOk returns a tuple with the SriovSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSriovSettings

`func (o *VnicBaseEthIf) SetSriovSettings(v VnicSriovSettings)`

SetSriovSettings sets SriovSettings field to given value.

### HasSriovSettings

`func (o *VnicBaseEthIf) HasSriovSettings() bool`

HasSriovSettings returns a boolean if a field has been set.

### SetSriovSettingsNil

`func (o *VnicBaseEthIf) SetSriovSettingsNil(b bool)`

 SetSriovSettingsNil sets the value for SriovSettings to be an explicit nil

### UnsetSriovSettings
`func (o *VnicBaseEthIf) UnsetSriovSettings()`

UnsetSriovSettings ensures that no value is present for SriovSettings, not even an explicit nil
### GetUsnicSettings

`func (o *VnicBaseEthIf) GetUsnicSettings() VnicUsnicSettings`

GetUsnicSettings returns the UsnicSettings field if non-nil, zero value otherwise.

### GetUsnicSettingsOk

`func (o *VnicBaseEthIf) GetUsnicSettingsOk() (*VnicUsnicSettings, bool)`

GetUsnicSettingsOk returns a tuple with the UsnicSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicSettings

`func (o *VnicBaseEthIf) SetUsnicSettings(v VnicUsnicSettings)`

SetUsnicSettings sets UsnicSettings field to given value.

### HasUsnicSettings

`func (o *VnicBaseEthIf) HasUsnicSettings() bool`

HasUsnicSettings returns a boolean if a field has been set.

### SetUsnicSettingsNil

`func (o *VnicBaseEthIf) SetUsnicSettingsNil(b bool)`

 SetUsnicSettingsNil sets the value for UsnicSettings to be an explicit nil

### UnsetUsnicSettings
`func (o *VnicBaseEthIf) UnsetUsnicSettings()`

UnsetUsnicSettings ensures that no value is present for UsnicSettings, not even an explicit nil
### GetVmqSettings

`func (o *VnicBaseEthIf) GetVmqSettings() VnicVmqSettings`

GetVmqSettings returns the VmqSettings field if non-nil, zero value otherwise.

### GetVmqSettingsOk

`func (o *VnicBaseEthIf) GetVmqSettingsOk() (*VnicVmqSettings, bool)`

GetVmqSettingsOk returns a tuple with the VmqSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmqSettings

`func (o *VnicBaseEthIf) SetVmqSettings(v VnicVmqSettings)`

SetVmqSettings sets VmqSettings field to given value.

### HasVmqSettings

`func (o *VnicBaseEthIf) HasVmqSettings() bool`

HasVmqSettings returns a boolean if a field has been set.

### SetVmqSettingsNil

`func (o *VnicBaseEthIf) SetVmqSettingsNil(b bool)`

 SetVmqSettingsNil sets the value for VmqSettings to be an explicit nil

### UnsetVmqSettings
`func (o *VnicBaseEthIf) UnsetVmqSettings()`

UnsetVmqSettings ensures that no value is present for VmqSettings, not even an explicit nil
### GetEthAdapterPolicy

`func (o *VnicBaseEthIf) GetEthAdapterPolicy() VnicEthAdapterPolicyRelationship`

GetEthAdapterPolicy returns the EthAdapterPolicy field if non-nil, zero value otherwise.

### GetEthAdapterPolicyOk

`func (o *VnicBaseEthIf) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyRelationship, bool)`

GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthAdapterPolicy

`func (o *VnicBaseEthIf) SetEthAdapterPolicy(v VnicEthAdapterPolicyRelationship)`

SetEthAdapterPolicy sets EthAdapterPolicy field to given value.

### HasEthAdapterPolicy

`func (o *VnicBaseEthIf) HasEthAdapterPolicy() bool`

HasEthAdapterPolicy returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *VnicBaseEthIf) GetEthNetworkPolicy() VnicEthNetworkPolicyRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *VnicBaseEthIf) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *VnicBaseEthIf) SetEthNetworkPolicy(v VnicEthNetworkPolicyRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *VnicBaseEthIf) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetEthQosPolicy

`func (o *VnicBaseEthIf) GetEthQosPolicy() VnicEthQosPolicyRelationship`

GetEthQosPolicy returns the EthQosPolicy field if non-nil, zero value otherwise.

### GetEthQosPolicyOk

`func (o *VnicBaseEthIf) GetEthQosPolicyOk() (*VnicEthQosPolicyRelationship, bool)`

GetEthQosPolicyOk returns a tuple with the EthQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthQosPolicy

`func (o *VnicBaseEthIf) SetEthQosPolicy(v VnicEthQosPolicyRelationship)`

SetEthQosPolicy sets EthQosPolicy field to given value.

### HasEthQosPolicy

`func (o *VnicBaseEthIf) HasEthQosPolicy() bool`

HasEthQosPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkControlPolicy

`func (o *VnicBaseEthIf) GetFabricEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship`

GetFabricEthNetworkControlPolicy returns the FabricEthNetworkControlPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkControlPolicyOk

`func (o *VnicBaseEthIf) GetFabricEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool)`

GetFabricEthNetworkControlPolicyOk returns a tuple with the FabricEthNetworkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkControlPolicy

`func (o *VnicBaseEthIf) SetFabricEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship)`

SetFabricEthNetworkControlPolicy sets FabricEthNetworkControlPolicy field to given value.

### HasFabricEthNetworkControlPolicy

`func (o *VnicBaseEthIf) HasFabricEthNetworkControlPolicy() bool`

HasFabricEthNetworkControlPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkGroupPolicy

`func (o *VnicBaseEthIf) GetFabricEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship`

GetFabricEthNetworkGroupPolicy returns the FabricEthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkGroupPolicyOk

`func (o *VnicBaseEthIf) GetFabricEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyRelationship, bool)`

GetFabricEthNetworkGroupPolicyOk returns a tuple with the FabricEthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkGroupPolicy

`func (o *VnicBaseEthIf) SetFabricEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship)`

SetFabricEthNetworkGroupPolicy sets FabricEthNetworkGroupPolicy field to given value.

### HasFabricEthNetworkGroupPolicy

`func (o *VnicBaseEthIf) HasFabricEthNetworkGroupPolicy() bool`

HasFabricEthNetworkGroupPolicy returns a boolean if a field has been set.

### SetFabricEthNetworkGroupPolicyNil

`func (o *VnicBaseEthIf) SetFabricEthNetworkGroupPolicyNil(b bool)`

 SetFabricEthNetworkGroupPolicyNil sets the value for FabricEthNetworkGroupPolicy to be an explicit nil

### UnsetFabricEthNetworkGroupPolicy
`func (o *VnicBaseEthIf) UnsetFabricEthNetworkGroupPolicy()`

UnsetFabricEthNetworkGroupPolicy ensures that no value is present for FabricEthNetworkGroupPolicy, not even an explicit nil
### GetIscsiBootPolicy

`func (o *VnicBaseEthIf) GetIscsiBootPolicy() VnicIscsiBootPolicyRelationship`

GetIscsiBootPolicy returns the IscsiBootPolicy field if non-nil, zero value otherwise.

### GetIscsiBootPolicyOk

`func (o *VnicBaseEthIf) GetIscsiBootPolicyOk() (*VnicIscsiBootPolicyRelationship, bool)`

GetIscsiBootPolicyOk returns a tuple with the IscsiBootPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiBootPolicy

`func (o *VnicBaseEthIf) SetIscsiBootPolicy(v VnicIscsiBootPolicyRelationship)`

SetIscsiBootPolicy sets IscsiBootPolicy field to given value.

### HasIscsiBootPolicy

`func (o *VnicBaseEthIf) HasIscsiBootPolicy() bool`

HasIscsiBootPolicy returns a boolean if a field has been set.

### GetMacPool

`func (o *VnicBaseEthIf) GetMacPool() MacpoolPoolRelationship`

GetMacPool returns the MacPool field if non-nil, zero value otherwise.

### GetMacPoolOk

`func (o *VnicBaseEthIf) GetMacPoolOk() (*MacpoolPoolRelationship, bool)`

GetMacPoolOk returns a tuple with the MacPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPool

`func (o *VnicBaseEthIf) SetMacPool(v MacpoolPoolRelationship)`

SetMacPool sets MacPool field to given value.

### HasMacPool

`func (o *VnicBaseEthIf) HasMacPool() bool`

HasMacPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


