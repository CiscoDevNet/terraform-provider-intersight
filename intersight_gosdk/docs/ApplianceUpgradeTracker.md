# ApplianceUpgradeTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.UpgradeTracker"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.UpgradeTracker"]
**Description** | Pointer to **string** | A description of the upgrade or upgrade failure. | [optional] [readonly] 
**LastUpgradeFailed** | Pointer to **bool** | Indicates if the last upgrade has failed or not. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the failed upgrade. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceUpgradeTracker

`func NewApplianceUpgradeTracker(classId string, objectType string, ) *ApplianceUpgradeTracker`

NewApplianceUpgradeTracker instantiates a new ApplianceUpgradeTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceUpgradeTrackerWithDefaults

`func NewApplianceUpgradeTrackerWithDefaults() *ApplianceUpgradeTracker`

NewApplianceUpgradeTrackerWithDefaults instantiates a new ApplianceUpgradeTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceUpgradeTracker) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceUpgradeTracker) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceUpgradeTracker) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceUpgradeTracker) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceUpgradeTracker) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceUpgradeTracker) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ApplianceUpgradeTracker) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplianceUpgradeTracker) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplianceUpgradeTracker) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplianceUpgradeTracker) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastUpgradeFailed

`func (o *ApplianceUpgradeTracker) GetLastUpgradeFailed() bool`

GetLastUpgradeFailed returns the LastUpgradeFailed field if non-nil, zero value otherwise.

### GetLastUpgradeFailedOk

`func (o *ApplianceUpgradeTracker) GetLastUpgradeFailedOk() (*bool, bool)`

GetLastUpgradeFailedOk returns a tuple with the LastUpgradeFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgradeFailed

`func (o *ApplianceUpgradeTracker) SetLastUpgradeFailed(v bool)`

SetLastUpgradeFailed sets LastUpgradeFailed field to given value.

### HasLastUpgradeFailed

`func (o *ApplianceUpgradeTracker) HasLastUpgradeFailed() bool`

HasLastUpgradeFailed returns a boolean if a field has been set.

### GetVersion

`func (o *ApplianceUpgradeTracker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceUpgradeTracker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceUpgradeTracker) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceUpgradeTracker) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceUpgradeTracker) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceUpgradeTracker) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceUpgradeTracker) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceUpgradeTracker) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ApplianceUpgradeTracker) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ApplianceUpgradeTracker) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


