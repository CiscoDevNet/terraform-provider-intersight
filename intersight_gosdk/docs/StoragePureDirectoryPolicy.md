# StoragePureDirectoryPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureDirectoryPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureDirectoryPolicy"]
**Destroyed** | Pointer to **bool** | Returns a value of true if the managed directory of the export has been destroyed and is pending eradication. The export can be recovered by recovering the destroyed managed directory. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Returns a value of true if the export policy that manages this export is enabled. | [optional] [readonly] 
**Name** | Pointer to **string** | The export name for accessing this export. | [optional] [readonly] 
**PolicyType** | Pointer to **string** | The export policy that manages this export. An export can be managed by at most one export policy. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ExportMembers** | Pointer to [**[]StoragePureDirectoryExportRelationship**](StoragePureDirectoryExportRelationship.md) | An array of relationships to storagePureDirectoryExport resources. | [optional] [readonly] 
**NfsRules** | Pointer to [**[]StoragePureNfsPolicyRuleRelationship**](StoragePureNfsPolicyRuleRelationship.md) | An array of relationships to storagePureNfsPolicyRule resources. | [optional] [readonly] 
**QuotaMembers** | Pointer to [**[]StoragePureDirectoryQuotaRelationship**](StoragePureDirectoryQuotaRelationship.md) | An array of relationships to storagePureDirectoryQuota resources. | [optional] [readonly] 
**QuotaRules** | Pointer to [**[]StoragePureQuotaPolicyRuleRelationship**](StoragePureQuotaPolicyRuleRelationship.md) | An array of relationships to storagePureQuotaPolicyRule resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SmbRules** | Pointer to [**[]StoragePureSmbPolicyRuleRelationship**](StoragePureSmbPolicyRuleRelationship.md) | An array of relationships to storagePureSmbPolicyRule resources. | [optional] [readonly] 

## Methods

### NewStoragePureDirectoryPolicy

`func NewStoragePureDirectoryPolicy(classId string, objectType string, ) *StoragePureDirectoryPolicy`

NewStoragePureDirectoryPolicy instantiates a new StoragePureDirectoryPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureDirectoryPolicyWithDefaults

`func NewStoragePureDirectoryPolicyWithDefaults() *StoragePureDirectoryPolicy`

NewStoragePureDirectoryPolicyWithDefaults instantiates a new StoragePureDirectoryPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureDirectoryPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureDirectoryPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureDirectoryPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureDirectoryPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureDirectoryPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureDirectoryPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestroyed

`func (o *StoragePureDirectoryPolicy) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePureDirectoryPolicy) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePureDirectoryPolicy) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePureDirectoryPolicy) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetEnabled

`func (o *StoragePureDirectoryPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StoragePureDirectoryPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StoragePureDirectoryPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StoragePureDirectoryPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *StoragePureDirectoryPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureDirectoryPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureDirectoryPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureDirectoryPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyType

`func (o *StoragePureDirectoryPolicy) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *StoragePureDirectoryPolicy) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *StoragePureDirectoryPolicy) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *StoragePureDirectoryPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureDirectoryPolicy) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureDirectoryPolicy) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureDirectoryPolicy) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureDirectoryPolicy) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureDirectoryPolicy) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureDirectoryPolicy) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetExportMembers

`func (o *StoragePureDirectoryPolicy) GetExportMembers() []StoragePureDirectoryExportRelationship`

GetExportMembers returns the ExportMembers field if non-nil, zero value otherwise.

### GetExportMembersOk

`func (o *StoragePureDirectoryPolicy) GetExportMembersOk() (*[]StoragePureDirectoryExportRelationship, bool)`

GetExportMembersOk returns a tuple with the ExportMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMembers

`func (o *StoragePureDirectoryPolicy) SetExportMembers(v []StoragePureDirectoryExportRelationship)`

SetExportMembers sets ExportMembers field to given value.

### HasExportMembers

`func (o *StoragePureDirectoryPolicy) HasExportMembers() bool`

HasExportMembers returns a boolean if a field has been set.

### SetExportMembersNil

`func (o *StoragePureDirectoryPolicy) SetExportMembersNil(b bool)`

 SetExportMembersNil sets the value for ExportMembers to be an explicit nil

### UnsetExportMembers
`func (o *StoragePureDirectoryPolicy) UnsetExportMembers()`

UnsetExportMembers ensures that no value is present for ExportMembers, not even an explicit nil
### GetNfsRules

`func (o *StoragePureDirectoryPolicy) GetNfsRules() []StoragePureNfsPolicyRuleRelationship`

GetNfsRules returns the NfsRules field if non-nil, zero value otherwise.

### GetNfsRulesOk

`func (o *StoragePureDirectoryPolicy) GetNfsRulesOk() (*[]StoragePureNfsPolicyRuleRelationship, bool)`

GetNfsRulesOk returns a tuple with the NfsRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRules

`func (o *StoragePureDirectoryPolicy) SetNfsRules(v []StoragePureNfsPolicyRuleRelationship)`

SetNfsRules sets NfsRules field to given value.

### HasNfsRules

`func (o *StoragePureDirectoryPolicy) HasNfsRules() bool`

HasNfsRules returns a boolean if a field has been set.

### SetNfsRulesNil

`func (o *StoragePureDirectoryPolicy) SetNfsRulesNil(b bool)`

 SetNfsRulesNil sets the value for NfsRules to be an explicit nil

### UnsetNfsRules
`func (o *StoragePureDirectoryPolicy) UnsetNfsRules()`

UnsetNfsRules ensures that no value is present for NfsRules, not even an explicit nil
### GetQuotaMembers

`func (o *StoragePureDirectoryPolicy) GetQuotaMembers() []StoragePureDirectoryQuotaRelationship`

GetQuotaMembers returns the QuotaMembers field if non-nil, zero value otherwise.

### GetQuotaMembersOk

`func (o *StoragePureDirectoryPolicy) GetQuotaMembersOk() (*[]StoragePureDirectoryQuotaRelationship, bool)`

GetQuotaMembersOk returns a tuple with the QuotaMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMembers

`func (o *StoragePureDirectoryPolicy) SetQuotaMembers(v []StoragePureDirectoryQuotaRelationship)`

SetQuotaMembers sets QuotaMembers field to given value.

### HasQuotaMembers

`func (o *StoragePureDirectoryPolicy) HasQuotaMembers() bool`

HasQuotaMembers returns a boolean if a field has been set.

### SetQuotaMembersNil

`func (o *StoragePureDirectoryPolicy) SetQuotaMembersNil(b bool)`

 SetQuotaMembersNil sets the value for QuotaMembers to be an explicit nil

### UnsetQuotaMembers
`func (o *StoragePureDirectoryPolicy) UnsetQuotaMembers()`

UnsetQuotaMembers ensures that no value is present for QuotaMembers, not even an explicit nil
### GetQuotaRules

`func (o *StoragePureDirectoryPolicy) GetQuotaRules() []StoragePureQuotaPolicyRuleRelationship`

GetQuotaRules returns the QuotaRules field if non-nil, zero value otherwise.

### GetQuotaRulesOk

`func (o *StoragePureDirectoryPolicy) GetQuotaRulesOk() (*[]StoragePureQuotaPolicyRuleRelationship, bool)`

GetQuotaRulesOk returns a tuple with the QuotaRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRules

`func (o *StoragePureDirectoryPolicy) SetQuotaRules(v []StoragePureQuotaPolicyRuleRelationship)`

SetQuotaRules sets QuotaRules field to given value.

### HasQuotaRules

`func (o *StoragePureDirectoryPolicy) HasQuotaRules() bool`

HasQuotaRules returns a boolean if a field has been set.

### SetQuotaRulesNil

`func (o *StoragePureDirectoryPolicy) SetQuotaRulesNil(b bool)`

 SetQuotaRulesNil sets the value for QuotaRules to be an explicit nil

### UnsetQuotaRules
`func (o *StoragePureDirectoryPolicy) UnsetQuotaRules()`

UnsetQuotaRules ensures that no value is present for QuotaRules, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureDirectoryPolicy) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureDirectoryPolicy) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureDirectoryPolicy) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureDirectoryPolicy) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureDirectoryPolicy) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureDirectoryPolicy) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSmbRules

`func (o *StoragePureDirectoryPolicy) GetSmbRules() []StoragePureSmbPolicyRuleRelationship`

GetSmbRules returns the SmbRules field if non-nil, zero value otherwise.

### GetSmbRulesOk

`func (o *StoragePureDirectoryPolicy) GetSmbRulesOk() (*[]StoragePureSmbPolicyRuleRelationship, bool)`

GetSmbRulesOk returns a tuple with the SmbRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbRules

`func (o *StoragePureDirectoryPolicy) SetSmbRules(v []StoragePureSmbPolicyRuleRelationship)`

SetSmbRules sets SmbRules field to given value.

### HasSmbRules

`func (o *StoragePureDirectoryPolicy) HasSmbRules() bool`

HasSmbRules returns a boolean if a field has been set.

### SetSmbRulesNil

`func (o *StoragePureDirectoryPolicy) SetSmbRulesNil(b bool)`

 SetSmbRulesNil sets the value for SmbRules to be an explicit nil

### UnsetSmbRules
`func (o *StoragePureDirectoryPolicy) UnsetSmbRules()`

UnsetSmbRules ensures that no value is present for SmbRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


