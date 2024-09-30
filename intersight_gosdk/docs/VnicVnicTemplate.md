# VnicVnicTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.VnicTemplate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.VnicTemplate"]
**Description** | Pointer to **string** | Description of the vNIC template. | [optional] 
**EnableOverride** | Pointer to **bool** | When enabled, the configuration of the derived instances may override the template configuration. | [optional] 
**LcpUsageCount** | Pointer to **int64** | The count of the Lan Connectivity Policies using vNIC template. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the vNIC template. | [optional] 
**PeerVnicName** | Pointer to **string** | Name of the peer vNIC which belongs to the peer FI. | [optional] 
**SwitchId** | Pointer to **string** | The fabric port to which the vNICs will be associated. * &#x60;None&#x60; - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value &#39;None&#39; should be used. * &#x60;A&#x60; - Fabric A of the FI cluster. * &#x60;B&#x60; - Fabric B of the FI cluster. | [optional] [default to "None"]
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**UsageCount** | Pointer to **int64** | The number of objects derived from a Template MO instance. | [optional] [readonly] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewVnicVnicTemplate

`func NewVnicVnicTemplate(classId string, objectType string, ) *VnicVnicTemplate`

NewVnicVnicTemplate instantiates a new VnicVnicTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicVnicTemplateWithDefaults

`func NewVnicVnicTemplateWithDefaults() *VnicVnicTemplate`

NewVnicVnicTemplateWithDefaults instantiates a new VnicVnicTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicVnicTemplate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicVnicTemplate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicVnicTemplate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicVnicTemplate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicVnicTemplate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicVnicTemplate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *VnicVnicTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VnicVnicTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VnicVnicTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VnicVnicTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableOverride

`func (o *VnicVnicTemplate) GetEnableOverride() bool`

GetEnableOverride returns the EnableOverride field if non-nil, zero value otherwise.

### GetEnableOverrideOk

`func (o *VnicVnicTemplate) GetEnableOverrideOk() (*bool, bool)`

GetEnableOverrideOk returns a tuple with the EnableOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOverride

`func (o *VnicVnicTemplate) SetEnableOverride(v bool)`

SetEnableOverride sets EnableOverride field to given value.

### HasEnableOverride

`func (o *VnicVnicTemplate) HasEnableOverride() bool`

HasEnableOverride returns a boolean if a field has been set.

### GetLcpUsageCount

`func (o *VnicVnicTemplate) GetLcpUsageCount() int64`

GetLcpUsageCount returns the LcpUsageCount field if non-nil, zero value otherwise.

### GetLcpUsageCountOk

`func (o *VnicVnicTemplate) GetLcpUsageCountOk() (*int64, bool)`

GetLcpUsageCountOk returns a tuple with the LcpUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpUsageCount

`func (o *VnicVnicTemplate) SetLcpUsageCount(v int64)`

SetLcpUsageCount sets LcpUsageCount field to given value.

### HasLcpUsageCount

`func (o *VnicVnicTemplate) HasLcpUsageCount() bool`

HasLcpUsageCount returns a boolean if a field has been set.

### GetName

`func (o *VnicVnicTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicVnicTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicVnicTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicVnicTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeerVnicName

`func (o *VnicVnicTemplate) GetPeerVnicName() string`

GetPeerVnicName returns the PeerVnicName field if non-nil, zero value otherwise.

### GetPeerVnicNameOk

`func (o *VnicVnicTemplate) GetPeerVnicNameOk() (*string, bool)`

GetPeerVnicNameOk returns a tuple with the PeerVnicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerVnicName

`func (o *VnicVnicTemplate) SetPeerVnicName(v string)`

SetPeerVnicName sets PeerVnicName field to given value.

### HasPeerVnicName

`func (o *VnicVnicTemplate) HasPeerVnicName() bool`

HasPeerVnicName returns a boolean if a field has been set.

### GetSwitchId

`func (o *VnicVnicTemplate) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *VnicVnicTemplate) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *VnicVnicTemplate) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *VnicVnicTemplate) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetTemplateActions

`func (o *VnicVnicTemplate) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *VnicVnicTemplate) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *VnicVnicTemplate) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *VnicVnicTemplate) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *VnicVnicTemplate) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *VnicVnicTemplate) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetUsageCount

`func (o *VnicVnicTemplate) GetUsageCount() int64`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *VnicVnicTemplate) GetUsageCountOk() (*int64, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *VnicVnicTemplate) SetUsageCount(v int64)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *VnicVnicTemplate) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicVnicTemplate) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicVnicTemplate) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicVnicTemplate) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicVnicTemplate) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *VnicVnicTemplate) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *VnicVnicTemplate) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


