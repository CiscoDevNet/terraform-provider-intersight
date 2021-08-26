# FabricAppliancePcRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.AppliancePcRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.AppliancePcRole"]
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port channel. * &#x60;Auto&#x60; - Admin configurable speed AUTO ( default ). * &#x60;1Gbps&#x60; - Admin configurable speed 1Gbps. * &#x60;10Gbps&#x60; - Admin configurable speed 10Gbps. * &#x60;25Gbps&#x60; - Admin configurable speed 25Gbps. * &#x60;40Gbps&#x60; - Admin configurable speed 40Gbps. * &#x60;100Gbps&#x60; - Admin configurable speed 100Gbps. | [optional] [default to "Auto"]
**Mode** | Pointer to **string** | Port mode to be set on the appliance port-channel. * &#x60;trunk&#x60; - Trunk Mode Switch Port Type. * &#x60;access&#x60; - Access Mode Switch Port Type. | [optional] [default to "trunk"]
**Priority** | Pointer to **string** | The &#39;name&#39; of the System QoS Class. * &#x60;Best Effort&#x60; - QoS Priority for Best-effort traffic. * &#x60;FC&#x60; - QoS Priority for FC traffic. * &#x60;Platinum&#x60; - QoS Priority for Platinum traffic. * &#x60;Gold&#x60; - QoS Priority for Gold traffic. * &#x60;Silver&#x60; - QoS Priority for Silver traffic. * &#x60;Bronze&#x60; - QoS Priority for Bronze traffic. | [optional] [default to "Best Effort"]
**EthNetworkControlPolicy** | Pointer to [**FabricEthNetworkControlPolicyRelationship**](FabricEthNetworkControlPolicyRelationship.md) |  | [optional] 
**EthNetworkGroupPolicy** | Pointer to [**FabricEthNetworkGroupPolicyRelationship**](FabricEthNetworkGroupPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricAppliancePcRole

`func NewFabricAppliancePcRole(classId string, objectType string, ) *FabricAppliancePcRole`

NewFabricAppliancePcRole instantiates a new FabricAppliancePcRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricAppliancePcRoleWithDefaults

`func NewFabricAppliancePcRoleWithDefaults() *FabricAppliancePcRole`

NewFabricAppliancePcRoleWithDefaults instantiates a new FabricAppliancePcRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricAppliancePcRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricAppliancePcRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricAppliancePcRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricAppliancePcRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricAppliancePcRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricAppliancePcRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FabricAppliancePcRole) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricAppliancePcRole) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricAppliancePcRole) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricAppliancePcRole) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetMode

`func (o *FabricAppliancePcRole) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FabricAppliancePcRole) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FabricAppliancePcRole) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FabricAppliancePcRole) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPriority

`func (o *FabricAppliancePcRole) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *FabricAppliancePcRole) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *FabricAppliancePcRole) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *FabricAppliancePcRole) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEthNetworkControlPolicy

`func (o *FabricAppliancePcRole) GetEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship`

GetEthNetworkControlPolicy returns the EthNetworkControlPolicy field if non-nil, zero value otherwise.

### GetEthNetworkControlPolicyOk

`func (o *FabricAppliancePcRole) GetEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool)`

GetEthNetworkControlPolicyOk returns a tuple with the EthNetworkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkControlPolicy

`func (o *FabricAppliancePcRole) SetEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship)`

SetEthNetworkControlPolicy sets EthNetworkControlPolicy field to given value.

### HasEthNetworkControlPolicy

`func (o *FabricAppliancePcRole) HasEthNetworkControlPolicy() bool`

HasEthNetworkControlPolicy returns a boolean if a field has been set.

### GetEthNetworkGroupPolicy

`func (o *FabricAppliancePcRole) GetEthNetworkGroupPolicy() FabricEthNetworkGroupPolicyRelationship`

GetEthNetworkGroupPolicy returns the EthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetEthNetworkGroupPolicyOk

`func (o *FabricAppliancePcRole) GetEthNetworkGroupPolicyOk() (*FabricEthNetworkGroupPolicyRelationship, bool)`

GetEthNetworkGroupPolicyOk returns a tuple with the EthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkGroupPolicy

`func (o *FabricAppliancePcRole) SetEthNetworkGroupPolicy(v FabricEthNetworkGroupPolicyRelationship)`

SetEthNetworkGroupPolicy sets EthNetworkGroupPolicy field to given value.

### HasEthNetworkGroupPolicy

`func (o *FabricAppliancePcRole) HasEthNetworkGroupPolicy() bool`

HasEthNetworkGroupPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


