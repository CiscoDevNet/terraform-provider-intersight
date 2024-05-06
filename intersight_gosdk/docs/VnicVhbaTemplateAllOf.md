# VnicVhbaTemplateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.VhbaTemplate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.VhbaTemplate"]
**Description** | Pointer to **string** | Description of the vHBA template. | [optional] 
**EnableOverride** | Pointer to **bool** | When enabled, the configuration of the derived instances may override the template configuration. | [optional] 
**Name** | Pointer to **string** | Name of the virtual fibre channel interface. | [optional] 
**PeerVhbaName** | Pointer to **string** | Name of the peer vHBA which belongs to the peer FI. | [optional] 
**ScpUsageCount** | Pointer to **int64** | The count of the San Connectivity Policies using vHBA template. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | The fabric port to which the vHBAs will be associated. * &#x60;None&#x60; - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value &#39;None&#39; should be used. * &#x60;A&#x60; - Fabric A of the FI cluster. * &#x60;B&#x60; - Fabric B of the FI cluster. | [optional] [default to "None"]
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**UsageCount** | Pointer to **int64** | The number of objects derived from a Template MO instance. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewVnicVhbaTemplateAllOf

`func NewVnicVhbaTemplateAllOf(classId string, objectType string, ) *VnicVhbaTemplateAllOf`

NewVnicVhbaTemplateAllOf instantiates a new VnicVhbaTemplateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicVhbaTemplateAllOfWithDefaults

`func NewVnicVhbaTemplateAllOfWithDefaults() *VnicVhbaTemplateAllOf`

NewVnicVhbaTemplateAllOfWithDefaults instantiates a new VnicVhbaTemplateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicVhbaTemplateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicVhbaTemplateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicVhbaTemplateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicVhbaTemplateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicVhbaTemplateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicVhbaTemplateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *VnicVhbaTemplateAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VnicVhbaTemplateAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VnicVhbaTemplateAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VnicVhbaTemplateAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableOverride

`func (o *VnicVhbaTemplateAllOf) GetEnableOverride() bool`

GetEnableOverride returns the EnableOverride field if non-nil, zero value otherwise.

### GetEnableOverrideOk

`func (o *VnicVhbaTemplateAllOf) GetEnableOverrideOk() (*bool, bool)`

GetEnableOverrideOk returns a tuple with the EnableOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOverride

`func (o *VnicVhbaTemplateAllOf) SetEnableOverride(v bool)`

SetEnableOverride sets EnableOverride field to given value.

### HasEnableOverride

`func (o *VnicVhbaTemplateAllOf) HasEnableOverride() bool`

HasEnableOverride returns a boolean if a field has been set.

### GetName

`func (o *VnicVhbaTemplateAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicVhbaTemplateAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicVhbaTemplateAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicVhbaTemplateAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeerVhbaName

`func (o *VnicVhbaTemplateAllOf) GetPeerVhbaName() string`

GetPeerVhbaName returns the PeerVhbaName field if non-nil, zero value otherwise.

### GetPeerVhbaNameOk

`func (o *VnicVhbaTemplateAllOf) GetPeerVhbaNameOk() (*string, bool)`

GetPeerVhbaNameOk returns a tuple with the PeerVhbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerVhbaName

`func (o *VnicVhbaTemplateAllOf) SetPeerVhbaName(v string)`

SetPeerVhbaName sets PeerVhbaName field to given value.

### HasPeerVhbaName

`func (o *VnicVhbaTemplateAllOf) HasPeerVhbaName() bool`

HasPeerVhbaName returns a boolean if a field has been set.

### GetScpUsageCount

`func (o *VnicVhbaTemplateAllOf) GetScpUsageCount() int64`

GetScpUsageCount returns the ScpUsageCount field if non-nil, zero value otherwise.

### GetScpUsageCountOk

`func (o *VnicVhbaTemplateAllOf) GetScpUsageCountOk() (*int64, bool)`

GetScpUsageCountOk returns a tuple with the ScpUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpUsageCount

`func (o *VnicVhbaTemplateAllOf) SetScpUsageCount(v int64)`

SetScpUsageCount sets ScpUsageCount field to given value.

### HasScpUsageCount

`func (o *VnicVhbaTemplateAllOf) HasScpUsageCount() bool`

HasScpUsageCount returns a boolean if a field has been set.

### GetSwitchId

`func (o *VnicVhbaTemplateAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *VnicVhbaTemplateAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *VnicVhbaTemplateAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *VnicVhbaTemplateAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetTemplateActions

`func (o *VnicVhbaTemplateAllOf) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *VnicVhbaTemplateAllOf) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *VnicVhbaTemplateAllOf) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *VnicVhbaTemplateAllOf) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *VnicVhbaTemplateAllOf) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *VnicVhbaTemplateAllOf) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetUsageCount

`func (o *VnicVhbaTemplateAllOf) GetUsageCount() int64`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *VnicVhbaTemplateAllOf) GetUsageCountOk() (*int64, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *VnicVhbaTemplateAllOf) SetUsageCount(v int64)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *VnicVhbaTemplateAllOf) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicVhbaTemplateAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicVhbaTemplateAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicVhbaTemplateAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicVhbaTemplateAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


