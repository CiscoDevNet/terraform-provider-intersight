# ApplianceDataExportPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.DataExportPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.DataExportPolicy"]
**Enable** | Pointer to **bool** | Status of the data collection mode. If the value is &#39;true&#39;, then data collection is enabled. | [optional] 
**Name** | Pointer to **string** | Name of the Data Export Policy. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**ParentConfig** | Pointer to [**ApplianceDataExportPolicyRelationship**](ApplianceDataExportPolicyRelationship.md) |  | [optional] 
**SubConfigs** | Pointer to [**[]ApplianceDataExportPolicyRelationship**](ApplianceDataExportPolicyRelationship.md) | An array of relationships to applianceDataExportPolicy resources. | [optional] [readonly] 

## Methods

### NewApplianceDataExportPolicy

`func NewApplianceDataExportPolicy(classId string, objectType string, ) *ApplianceDataExportPolicy`

NewApplianceDataExportPolicy instantiates a new ApplianceDataExportPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceDataExportPolicyWithDefaults

`func NewApplianceDataExportPolicyWithDefaults() *ApplianceDataExportPolicy`

NewApplianceDataExportPolicyWithDefaults instantiates a new ApplianceDataExportPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceDataExportPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceDataExportPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceDataExportPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceDataExportPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceDataExportPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceDataExportPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnable

`func (o *ApplianceDataExportPolicy) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ApplianceDataExportPolicy) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ApplianceDataExportPolicy) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ApplianceDataExportPolicy) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetName

`func (o *ApplianceDataExportPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplianceDataExportPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplianceDataExportPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplianceDataExportPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceDataExportPolicy) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceDataExportPolicy) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceDataExportPolicy) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceDataExportPolicy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetParentConfig

`func (o *ApplianceDataExportPolicy) GetParentConfig() ApplianceDataExportPolicyRelationship`

GetParentConfig returns the ParentConfig field if non-nil, zero value otherwise.

### GetParentConfigOk

`func (o *ApplianceDataExportPolicy) GetParentConfigOk() (*ApplianceDataExportPolicyRelationship, bool)`

GetParentConfigOk returns a tuple with the ParentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConfig

`func (o *ApplianceDataExportPolicy) SetParentConfig(v ApplianceDataExportPolicyRelationship)`

SetParentConfig sets ParentConfig field to given value.

### HasParentConfig

`func (o *ApplianceDataExportPolicy) HasParentConfig() bool`

HasParentConfig returns a boolean if a field has been set.

### GetSubConfigs

`func (o *ApplianceDataExportPolicy) GetSubConfigs() []ApplianceDataExportPolicyRelationship`

GetSubConfigs returns the SubConfigs field if non-nil, zero value otherwise.

### GetSubConfigsOk

`func (o *ApplianceDataExportPolicy) GetSubConfigsOk() (*[]ApplianceDataExportPolicyRelationship, bool)`

GetSubConfigsOk returns a tuple with the SubConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubConfigs

`func (o *ApplianceDataExportPolicy) SetSubConfigs(v []ApplianceDataExportPolicyRelationship)`

SetSubConfigs sets SubConfigs field to given value.

### HasSubConfigs

`func (o *ApplianceDataExportPolicy) HasSubConfigs() bool`

HasSubConfigs returns a boolean if a field has been set.

### SetSubConfigsNil

`func (o *ApplianceDataExportPolicy) SetSubConfigsNil(b bool)`

 SetSubConfigsNil sets the value for SubConfigs to be an explicit nil

### UnsetSubConfigs
`func (o *ApplianceDataExportPolicy) UnsetSubConfigs()`

UnsetSubConfigs ensures that no value is present for SubConfigs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


