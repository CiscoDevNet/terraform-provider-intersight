# VnicBaseEthIfAllOf

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

### NewVnicBaseEthIfAllOf

`func NewVnicBaseEthIfAllOf(classId string, objectType string, ) *VnicBaseEthIfAllOf`

NewVnicBaseEthIfAllOf instantiates a new VnicBaseEthIfAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicBaseEthIfAllOfWithDefaults

`func NewVnicBaseEthIfAllOfWithDefaults() *VnicBaseEthIfAllOf`

NewVnicBaseEthIfAllOfWithDefaults instantiates a new VnicBaseEthIfAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicBaseEthIfAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicBaseEthIfAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicBaseEthIfAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicBaseEthIfAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicBaseEthIfAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicBaseEthIfAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCdn

`func (o *VnicBaseEthIfAllOf) GetCdn() VnicCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *VnicBaseEthIfAllOf) GetCdnOk() (*VnicCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *VnicBaseEthIfAllOf) SetCdn(v VnicCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *VnicBaseEthIfAllOf) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### SetCdnNil

`func (o *VnicBaseEthIfAllOf) SetCdnNil(b bool)`

 SetCdnNil sets the value for Cdn to be an explicit nil

### UnsetCdn
`func (o *VnicBaseEthIfAllOf) UnsetCdn()`

UnsetCdn ensures that no value is present for Cdn, not even an explicit nil
### GetFailoverEnabled

`func (o *VnicBaseEthIfAllOf) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *VnicBaseEthIfAllOf) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *VnicBaseEthIfAllOf) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *VnicBaseEthIfAllOf) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### GetPinGroupName

`func (o *VnicBaseEthIfAllOf) GetPinGroupName() string`

GetPinGroupName returns the PinGroupName field if non-nil, zero value otherwise.

### GetPinGroupNameOk

`func (o *VnicBaseEthIfAllOf) GetPinGroupNameOk() (*string, bool)`

GetPinGroupNameOk returns a tuple with the PinGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinGroupName

`func (o *VnicBaseEthIfAllOf) SetPinGroupName(v string)`

SetPinGroupName sets PinGroupName field to given value.

### HasPinGroupName

`func (o *VnicBaseEthIfAllOf) HasPinGroupName() bool`

HasPinGroupName returns a boolean if a field has been set.

### GetSriovSettings

`func (o *VnicBaseEthIfAllOf) GetSriovSettings() VnicSriovSettings`

GetSriovSettings returns the SriovSettings field if non-nil, zero value otherwise.

### GetSriovSettingsOk

`func (o *VnicBaseEthIfAllOf) GetSriovSettingsOk() (*VnicSriovSettings, bool)`

GetSriovSettingsOk returns a tuple with the SriovSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSriovSettings

`func (o *VnicBaseEthIfAllOf) SetSriovSettings(v VnicSriovSettings)`

SetSriovSettings sets SriovSettings field to given value.

### HasSriovSettings

`func (o *VnicBaseEthIfAllOf) HasSriovSettings() bool`

HasSriovSettings returns a boolean if a field has been set.

### SetSriovSettingsNil

`func (o *VnicBaseEthIfAllOf) SetSriovSettingsNil(b bool)`

 SetSriovSettingsNil sets the value for SriovSettings to be an explicit nil

### UnsetSriovSettings
`func (o *VnicBaseEthIfAllOf) UnsetSriovSettings()`

UnsetSriovSettings ensures that no value is present for SriovSettings, not even an explicit nil
### GetUsnicSettings

`func (o *VnicBaseEthIfAllOf) GetUsnicSettings() VnicUsnicSettings`

GetUsnicSettings returns the UsnicSettings field if non-nil, zero value otherwise.

### GetUsnicSettingsOk

`func (o *VnicBaseEthIfAllOf) GetUsnicSettingsOk() (*VnicUsnicSettings, bool)`

GetUsnicSettingsOk returns a tuple with the UsnicSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicSettings

`func (o *VnicBaseEthIfAllOf) SetUsnicSettings(v VnicUsnicSettings)`

SetUsnicSettings sets UsnicSettings field to given value.

### HasUsnicSettings

`func (o *VnicBaseEthIfAllOf) HasUsnicSettings() bool`

HasUsnicSettings returns a boolean if a field has been set.

### SetUsnicSettingsNil

`func (o *VnicBaseEthIfAllOf) SetUsnicSettingsNil(b bool)`

 SetUsnicSettingsNil sets the value for UsnicSettings to be an explicit nil

### UnsetUsnicSettings
`func (o *VnicBaseEthIfAllOf) UnsetUsnicSettings()`

UnsetUsnicSettings ensures that no value is present for UsnicSettings, not even an explicit nil
### GetVmqSettings

`func (o *VnicBaseEthIfAllOf) GetVmqSettings() VnicVmqSettings`

GetVmqSettings returns the VmqSettings field if non-nil, zero value otherwise.

### GetVmqSettingsOk

`func (o *VnicBaseEthIfAllOf) GetVmqSettingsOk() (*VnicVmqSettings, bool)`

GetVmqSettingsOk returns a tuple with the VmqSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmqSettings

`func (o *VnicBaseEthIfAllOf) SetVmqSettings(v VnicVmqSettings)`

SetVmqSettings sets VmqSettings field to given value.

### HasVmqSettings

`func (o *VnicBaseEthIfAllOf) HasVmqSettings() bool`

HasVmqSettings returns a boolean if a field has been set.

### SetVmqSettingsNil

`func (o *VnicBaseEthIfAllOf) SetVmqSettingsNil(b bool)`

 SetVmqSettingsNil sets the value for VmqSettings to be an explicit nil

### UnsetVmqSettings
`func (o *VnicBaseEthIfAllOf) UnsetVmqSettings()`

UnsetVmqSettings ensures that no value is present for VmqSettings, not even an explicit nil
### GetEthAdapterPolicy

`func (o *VnicBaseEthIfAllOf) GetEthAdapterPolicy() VnicEthAdapterPolicyRelationship`

GetEthAdapterPolicy returns the EthAdapterPolicy field if non-nil, zero value otherwise.

### GetEthAdapterPolicyOk

`func (o *VnicBaseEthIfAllOf) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyRelationship, bool)`

GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthAdapterPolicy

`func (o *VnicBaseEthIfAllOf) SetEthAdapterPolicy(v VnicEthAdapterPolicyRelationship)`

SetEthAdapterPolicy sets EthAdapterPolicy field to given value.

### HasEthAdapterPolicy

`func (o *VnicBaseEthIfAllOf) HasEthAdapterPolicy() bool`

HasEthAdapterPolicy returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *VnicBaseEthIfAllOf) GetEthNetworkPolicy() VnicEthNetworkPolicyRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *VnicBaseEthIfAllOf) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *VnicBaseEthIfAllOf) SetEthNetworkPolicy(v VnicEthNetworkPolicyRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *VnicBaseEthIfAllOf) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetEthQosPolicy

`func (o *VnicBaseEthIfAllOf) GetEthQosPolicy() VnicEthQosPolicyRelationship`

GetEthQosPolicy returns the EthQosPolicy field if non-nil, zero value otherwise.

### GetEthQosPolicyOk

`func (o *VnicBaseEthIfAllOf) GetEthQosPolicyOk() (*VnicEthQosPolicyRelationship, bool)`

GetEthQosPolicyOk returns a tuple with the EthQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthQosPolicy

`func (o *VnicBaseEthIfAllOf) SetEthQosPolicy(v VnicEthQosPolicyRelationship)`

SetEthQosPolicy sets EthQosPolicy field to given value.

### HasEthQosPolicy

`func (o *VnicBaseEthIfAllOf) HasEthQosPolicy() bool`

HasEthQosPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkControlPolicy

`func (o *VnicBaseEthIfAllOf) GetFabricEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship`

GetFabricEthNetworkControlPolicy returns the FabricEthNetworkControlPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkControlPolicyOk

`func (o *VnicBaseEthIfAllOf) GetFabricEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool)`

GetFabricEthNetworkControlPolicyOk returns a tuple with the FabricEthNetworkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkControlPolicy

`func (o *VnicBaseEthIfAllOf) SetFabricEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship)`

SetFabricEthNetworkControlPolicy sets FabricEthNetworkControlPolicy field to given value.

### HasFabricEthNetworkControlPolicy

`func (o *VnicBaseEthIfAllOf) HasFabricEthNetworkControlPolicy() bool`

HasFabricEthNetworkControlPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkGroupPolicy

`func (o *VnicBaseEthIfAllOf) GetFabricEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship`

GetFabricEthNetworkGroupPolicy returns the FabricEthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkGroupPolicyOk

`func (o *VnicBaseEthIfAllOf) GetFabricEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyRelationship, bool)`

GetFabricEthNetworkGroupPolicyOk returns a tuple with the FabricEthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkGroupPolicy

`func (o *VnicBaseEthIfAllOf) SetFabricEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship)`

SetFabricEthNetworkGroupPolicy sets FabricEthNetworkGroupPolicy field to given value.

### HasFabricEthNetworkGroupPolicy

`func (o *VnicBaseEthIfAllOf) HasFabricEthNetworkGroupPolicy() bool`

HasFabricEthNetworkGroupPolicy returns a boolean if a field has been set.

### SetFabricEthNetworkGroupPolicyNil

`func (o *VnicBaseEthIfAllOf) SetFabricEthNetworkGroupPolicyNil(b bool)`

 SetFabricEthNetworkGroupPolicyNil sets the value for FabricEthNetworkGroupPolicy to be an explicit nil

### UnsetFabricEthNetworkGroupPolicy
`func (o *VnicBaseEthIfAllOf) UnsetFabricEthNetworkGroupPolicy()`

UnsetFabricEthNetworkGroupPolicy ensures that no value is present for FabricEthNetworkGroupPolicy, not even an explicit nil
### GetIscsiBootPolicy

`func (o *VnicBaseEthIfAllOf) GetIscsiBootPolicy() VnicIscsiBootPolicyRelationship`

GetIscsiBootPolicy returns the IscsiBootPolicy field if non-nil, zero value otherwise.

### GetIscsiBootPolicyOk

`func (o *VnicBaseEthIfAllOf) GetIscsiBootPolicyOk() (*VnicIscsiBootPolicyRelationship, bool)`

GetIscsiBootPolicyOk returns a tuple with the IscsiBootPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiBootPolicy

`func (o *VnicBaseEthIfAllOf) SetIscsiBootPolicy(v VnicIscsiBootPolicyRelationship)`

SetIscsiBootPolicy sets IscsiBootPolicy field to given value.

### HasIscsiBootPolicy

`func (o *VnicBaseEthIfAllOf) HasIscsiBootPolicy() bool`

HasIscsiBootPolicy returns a boolean if a field has been set.

### GetMacPool

`func (o *VnicBaseEthIfAllOf) GetMacPool() MacpoolPoolRelationship`

GetMacPool returns the MacPool field if non-nil, zero value otherwise.

### GetMacPoolOk

`func (o *VnicBaseEthIfAllOf) GetMacPoolOk() (*MacpoolPoolRelationship, bool)`

GetMacPoolOk returns a tuple with the MacPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPool

`func (o *VnicBaseEthIfAllOf) SetMacPool(v MacpoolPoolRelationship)`

SetMacPool sets MacPool field to given value.

### HasMacPool

`func (o *VnicBaseEthIfAllOf) HasMacPool() bool`

HasMacPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


