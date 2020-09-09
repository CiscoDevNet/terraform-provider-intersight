# ApplianceDataExportPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Status of the data collection mode. If the value is &#39;true&#39;, then data collection is enabled. | [optional] 
**Name** | Pointer to **string** | Name of the Data Export Policy. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**ParentConfig** | Pointer to [**ApplianceDataExportPolicyRelationship**](appliance.DataExportPolicy.Relationship.md) |  | [optional] 
**SubConfigs** | Pointer to [**[]ApplianceDataExportPolicyRelationship**](appliance.DataExportPolicy.Relationship.md) | An array of relationships to applianceDataExportPolicy resources. | [optional] [readonly] 

## Methods

### NewApplianceDataExportPolicyAllOf

`func NewApplianceDataExportPolicyAllOf() *ApplianceDataExportPolicyAllOf`

NewApplianceDataExportPolicyAllOf instantiates a new ApplianceDataExportPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceDataExportPolicyAllOfWithDefaults

`func NewApplianceDataExportPolicyAllOfWithDefaults() *ApplianceDataExportPolicyAllOf`

NewApplianceDataExportPolicyAllOfWithDefaults instantiates a new ApplianceDataExportPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ApplianceDataExportPolicyAllOf) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ApplianceDataExportPolicyAllOf) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ApplianceDataExportPolicyAllOf) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ApplianceDataExportPolicyAllOf) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetName

`func (o *ApplianceDataExportPolicyAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplianceDataExportPolicyAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplianceDataExportPolicyAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplianceDataExportPolicyAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceDataExportPolicyAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceDataExportPolicyAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceDataExportPolicyAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceDataExportPolicyAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetParentConfig

`func (o *ApplianceDataExportPolicyAllOf) GetParentConfig() ApplianceDataExportPolicyRelationship`

GetParentConfig returns the ParentConfig field if non-nil, zero value otherwise.

### GetParentConfigOk

`func (o *ApplianceDataExportPolicyAllOf) GetParentConfigOk() (*ApplianceDataExportPolicyRelationship, bool)`

GetParentConfigOk returns a tuple with the ParentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConfig

`func (o *ApplianceDataExportPolicyAllOf) SetParentConfig(v ApplianceDataExportPolicyRelationship)`

SetParentConfig sets ParentConfig field to given value.

### HasParentConfig

`func (o *ApplianceDataExportPolicyAllOf) HasParentConfig() bool`

HasParentConfig returns a boolean if a field has been set.

### GetSubConfigs

`func (o *ApplianceDataExportPolicyAllOf) GetSubConfigs() []ApplianceDataExportPolicyRelationship`

GetSubConfigs returns the SubConfigs field if non-nil, zero value otherwise.

### GetSubConfigsOk

`func (o *ApplianceDataExportPolicyAllOf) GetSubConfigsOk() (*[]ApplianceDataExportPolicyRelationship, bool)`

GetSubConfigsOk returns a tuple with the SubConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubConfigs

`func (o *ApplianceDataExportPolicyAllOf) SetSubConfigs(v []ApplianceDataExportPolicyRelationship)`

SetSubConfigs sets SubConfigs field to given value.

### HasSubConfigs

`func (o *ApplianceDataExportPolicyAllOf) HasSubConfigs() bool`

HasSubConfigs returns a boolean if a field has been set.

### SetSubConfigsNil

`func (o *ApplianceDataExportPolicyAllOf) SetSubConfigsNil(b bool)`

 SetSubConfigsNil sets the value for SubConfigs to be an explicit nil

### UnsetSubConfigs
`func (o *ApplianceDataExportPolicyAllOf) UnsetSubConfigs()`

UnsetSubConfigs ensures that no value is present for SubConfigs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


