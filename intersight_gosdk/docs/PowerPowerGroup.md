# PowerPowerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "power.PowerGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "power.PowerGroup"]
**AdminState** | Pointer to **string** | Administrative state of the power group. Determines whether the power group is enabled or disabled. * &#x60;Enabled&#x60; - The power group is administratively enabled. * &#x60;Disabled&#x60; - The power group is administratively disabled. | [optional] [default to "Enabled"]
**BufferPowerRequired** | Pointer to **int64** | Reserved power buffer calculated as 10% of total chassis minimum power requirements. FI and FEX do not contribute to buffer power. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | Power Group Configuration State. * &#x60;Ok&#x60; - The configuration state of the power group is ok. All members are configured successfully. * &#x60;Unknown&#x60; - The configuration state of the power group is not known. * &#x60;Configuring&#x60; - The power group is actively being configured. Transitional state during CREATE or UPDATE operations. * &#x60;PartiallyConfigured&#x60; - The power group configuration failed mid-way. Some members may be configured while others are not. Totals may not match actual membership. Reconciliation or manual intervention is required. * &#x60;PartiallyDeleted&#x60; - The configuration state of the power group is Partially Deleted. This occurs when a DELETE operation was initiated but the workflow to remove power limits did not succeed for all members. * &#x60;Disabling&#x60; - The power group is being disabled. A workflow is in progress to push 0 power limits to all member devices. * &#x60;PartiallyDisabled&#x60; - The power group disable operation failed mid-way. Some members may have their power limits removed while others retain their previous values. The power group remains in a disabled state but may require manual intervention. | [optional] [readonly] [default to "Ok"]
**Description** | Pointer to **string** | Description of the power group. | [optional] 
**Name** | Pointer to **string** | User defined name for the power group. | [optional] 
**OperState** | Pointer to **string** | Power Group Operational State. * &#x60;Ok&#x60; - The operational state of the power group is ok. All members are configured successfully. * &#x60;Unknown&#x60; - The operational state of the power group is not known. * &#x60;InsufficientPower&#x60; - The power group does not have sufficient power budget to meet the minimum power requirements of all its chassis members. This typically occurs when new hardware is added to a member chassis, increasing its minimum power requirement beyond the available budget. | [optional] [readonly] [default to "Ok"]
**PowerBudget** | Pointer to **int64** | Maximum power budget allocated to this power group in watts. | [optional] 
**TotalMaximumPowerRequired** | Pointer to **int64** | Sum of maximum power required for each member in this group in watts. | [optional] [readonly] 
**TotalMinimumPowerRequired** | Pointer to **int64** | Sum of minimum power required for each member in this group in watts. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**InventoryDeviceMember** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] 
**PowerGroupMember** | Pointer to [**[]PowerPowerGroupMemberRelationship**](PowerPowerGroupMemberRelationship.md) | An array of relationships to powerPowerGroupMember resources. | [optional] [readonly] 

## Methods

### NewPowerPowerGroup

`func NewPowerPowerGroup(classId string, objectType string, ) *PowerPowerGroup`

NewPowerPowerGroup instantiates a new PowerPowerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPowerGroupWithDefaults

`func NewPowerPowerGroupWithDefaults() *PowerPowerGroup`

NewPowerPowerGroupWithDefaults instantiates a new PowerPowerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PowerPowerGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PowerPowerGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PowerPowerGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PowerPowerGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerPowerGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerPowerGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *PowerPowerGroup) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *PowerPowerGroup) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *PowerPowerGroup) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *PowerPowerGroup) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetBufferPowerRequired

`func (o *PowerPowerGroup) GetBufferPowerRequired() int64`

GetBufferPowerRequired returns the BufferPowerRequired field if non-nil, zero value otherwise.

### GetBufferPowerRequiredOk

`func (o *PowerPowerGroup) GetBufferPowerRequiredOk() (*int64, bool)`

GetBufferPowerRequiredOk returns a tuple with the BufferPowerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferPowerRequired

`func (o *PowerPowerGroup) SetBufferPowerRequired(v int64)`

SetBufferPowerRequired sets BufferPowerRequired field to given value.

### HasBufferPowerRequired

`func (o *PowerPowerGroup) HasBufferPowerRequired() bool`

HasBufferPowerRequired returns a boolean if a field has been set.

### GetConfigState

`func (o *PowerPowerGroup) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PowerPowerGroup) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PowerPowerGroup) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PowerPowerGroup) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetDescription

`func (o *PowerPowerGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerPowerGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerPowerGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerPowerGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PowerPowerGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerPowerGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerPowerGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PowerPowerGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *PowerPowerGroup) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PowerPowerGroup) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PowerPowerGroup) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PowerPowerGroup) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPowerBudget

`func (o *PowerPowerGroup) GetPowerBudget() int64`

GetPowerBudget returns the PowerBudget field if non-nil, zero value otherwise.

### GetPowerBudgetOk

`func (o *PowerPowerGroup) GetPowerBudgetOk() (*int64, bool)`

GetPowerBudgetOk returns a tuple with the PowerBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerBudget

`func (o *PowerPowerGroup) SetPowerBudget(v int64)`

SetPowerBudget sets PowerBudget field to given value.

### HasPowerBudget

`func (o *PowerPowerGroup) HasPowerBudget() bool`

HasPowerBudget returns a boolean if a field has been set.

### GetTotalMaximumPowerRequired

`func (o *PowerPowerGroup) GetTotalMaximumPowerRequired() int64`

GetTotalMaximumPowerRequired returns the TotalMaximumPowerRequired field if non-nil, zero value otherwise.

### GetTotalMaximumPowerRequiredOk

`func (o *PowerPowerGroup) GetTotalMaximumPowerRequiredOk() (*int64, bool)`

GetTotalMaximumPowerRequiredOk returns a tuple with the TotalMaximumPowerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaximumPowerRequired

`func (o *PowerPowerGroup) SetTotalMaximumPowerRequired(v int64)`

SetTotalMaximumPowerRequired sets TotalMaximumPowerRequired field to given value.

### HasTotalMaximumPowerRequired

`func (o *PowerPowerGroup) HasTotalMaximumPowerRequired() bool`

HasTotalMaximumPowerRequired returns a boolean if a field has been set.

### GetTotalMinimumPowerRequired

`func (o *PowerPowerGroup) GetTotalMinimumPowerRequired() int64`

GetTotalMinimumPowerRequired returns the TotalMinimumPowerRequired field if non-nil, zero value otherwise.

### GetTotalMinimumPowerRequiredOk

`func (o *PowerPowerGroup) GetTotalMinimumPowerRequiredOk() (*int64, bool)`

GetTotalMinimumPowerRequiredOk returns a tuple with the TotalMinimumPowerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMinimumPowerRequired

`func (o *PowerPowerGroup) SetTotalMinimumPowerRequired(v int64)`

SetTotalMinimumPowerRequired sets TotalMinimumPowerRequired field to given value.

### HasTotalMinimumPowerRequired

`func (o *PowerPowerGroup) HasTotalMinimumPowerRequired() bool`

HasTotalMinimumPowerRequired returns a boolean if a field has been set.

### GetAccount

`func (o *PowerPowerGroup) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PowerPowerGroup) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PowerPowerGroup) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PowerPowerGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *PowerPowerGroup) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *PowerPowerGroup) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetInventoryDeviceMember

`func (o *PowerPowerGroup) GetInventoryDeviceMember() []MoBaseMoRelationship`

GetInventoryDeviceMember returns the InventoryDeviceMember field if non-nil, zero value otherwise.

### GetInventoryDeviceMemberOk

`func (o *PowerPowerGroup) GetInventoryDeviceMemberOk() (*[]MoBaseMoRelationship, bool)`

GetInventoryDeviceMemberOk returns a tuple with the InventoryDeviceMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceMember

`func (o *PowerPowerGroup) SetInventoryDeviceMember(v []MoBaseMoRelationship)`

SetInventoryDeviceMember sets InventoryDeviceMember field to given value.

### HasInventoryDeviceMember

`func (o *PowerPowerGroup) HasInventoryDeviceMember() bool`

HasInventoryDeviceMember returns a boolean if a field has been set.

### SetInventoryDeviceMemberNil

`func (o *PowerPowerGroup) SetInventoryDeviceMemberNil(b bool)`

 SetInventoryDeviceMemberNil sets the value for InventoryDeviceMember to be an explicit nil

### UnsetInventoryDeviceMember
`func (o *PowerPowerGroup) UnsetInventoryDeviceMember()`

UnsetInventoryDeviceMember ensures that no value is present for InventoryDeviceMember, not even an explicit nil
### GetPowerGroupMember

`func (o *PowerPowerGroup) GetPowerGroupMember() []PowerPowerGroupMemberRelationship`

GetPowerGroupMember returns the PowerGroupMember field if non-nil, zero value otherwise.

### GetPowerGroupMemberOk

`func (o *PowerPowerGroup) GetPowerGroupMemberOk() (*[]PowerPowerGroupMemberRelationship, bool)`

GetPowerGroupMemberOk returns a tuple with the PowerGroupMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerGroupMember

`func (o *PowerPowerGroup) SetPowerGroupMember(v []PowerPowerGroupMemberRelationship)`

SetPowerGroupMember sets PowerGroupMember field to given value.

### HasPowerGroupMember

`func (o *PowerPowerGroup) HasPowerGroupMember() bool`

HasPowerGroupMember returns a boolean if a field has been set.

### SetPowerGroupMemberNil

`func (o *PowerPowerGroup) SetPowerGroupMemberNil(b bool)`

 SetPowerGroupMemberNil sets the value for PowerGroupMember to be an explicit nil

### UnsetPowerGroupMember
`func (o *PowerPowerGroup) UnsetPowerGroupMember()`

UnsetPowerGroupMember ensures that no value is present for PowerGroupMember, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


