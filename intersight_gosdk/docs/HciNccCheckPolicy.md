# HciNccCheckPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.NccCheckPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.NccCheckPolicy"]
**Classifications** | Pointer to **[]string** |  | [optional] 
**ClusterConfigs** | Pointer to [**[]HciSdaPolicyClusterConfig**](HciSdaPolicyClusterConfig.md) |  | [optional] 
**Description** | Pointer to **string** | NCC policy description text. | [optional] [readonly] 
**EntityType** | Pointer to **string** | Entity type against which the alert is raised. | [optional] [readonly] 
**ImpactType** | Pointer to **string** | Impact type to which this rule applies. | [optional] [readonly] 
**KbArticles** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | The name of the NCC policy. | [optional] [readonly] 
**NccPolicyExtId** | Pointer to **string** | Unique ID of the NCC policy. | [optional] [readonly] 
**PcExtId** | Pointer to **string** | Unique identifier of the domain manager (Prism Central) instance which owns this policy. | [optional] [readonly] 
**PolicyId** | Pointer to **string** | Unique ID associated with the policy. | [optional] [readonly] 
**Publisher** | Pointer to **string** | Publisher of the NCC policy. | [optional] [readonly] 
**Scope** | Pointer to **string** | Scope for the polcy execution. | [optional] [readonly] 
**SubType** | Pointer to **string** | Sub-type classification of the NCC policy. | [optional] [readonly] 
**TargetClusters** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** | The title of the NCC policy. | [optional] [readonly] 
**Type** | Pointer to **string** | Defines the type of the NCC policy. | [optional] [readonly] 
**DomainManager** | Pointer to [**NullableHciDomainManagerRelationship**](HciDomainManagerRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciNccCheckPolicy

`func NewHciNccCheckPolicy(classId string, objectType string, ) *HciNccCheckPolicy`

NewHciNccCheckPolicy instantiates a new HciNccCheckPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciNccCheckPolicyWithDefaults

`func NewHciNccCheckPolicyWithDefaults() *HciNccCheckPolicy`

NewHciNccCheckPolicyWithDefaults instantiates a new HciNccCheckPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciNccCheckPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciNccCheckPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciNccCheckPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciNccCheckPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciNccCheckPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciNccCheckPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClassifications

`func (o *HciNccCheckPolicy) GetClassifications() []string`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *HciNccCheckPolicy) GetClassificationsOk() (*[]string, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *HciNccCheckPolicy) SetClassifications(v []string)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *HciNccCheckPolicy) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### SetClassificationsNil

`func (o *HciNccCheckPolicy) SetClassificationsNil(b bool)`

 SetClassificationsNil sets the value for Classifications to be an explicit nil

### UnsetClassifications
`func (o *HciNccCheckPolicy) UnsetClassifications()`

UnsetClassifications ensures that no value is present for Classifications, not even an explicit nil
### GetClusterConfigs

`func (o *HciNccCheckPolicy) GetClusterConfigs() []HciSdaPolicyClusterConfig`

GetClusterConfigs returns the ClusterConfigs field if non-nil, zero value otherwise.

### GetClusterConfigsOk

`func (o *HciNccCheckPolicy) GetClusterConfigsOk() (*[]HciSdaPolicyClusterConfig, bool)`

GetClusterConfigsOk returns a tuple with the ClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConfigs

`func (o *HciNccCheckPolicy) SetClusterConfigs(v []HciSdaPolicyClusterConfig)`

SetClusterConfigs sets ClusterConfigs field to given value.

### HasClusterConfigs

`func (o *HciNccCheckPolicy) HasClusterConfigs() bool`

HasClusterConfigs returns a boolean if a field has been set.

### SetClusterConfigsNil

`func (o *HciNccCheckPolicy) SetClusterConfigsNil(b bool)`

 SetClusterConfigsNil sets the value for ClusterConfigs to be an explicit nil

### UnsetClusterConfigs
`func (o *HciNccCheckPolicy) UnsetClusterConfigs()`

UnsetClusterConfigs ensures that no value is present for ClusterConfigs, not even an explicit nil
### GetDescription

`func (o *HciNccCheckPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HciNccCheckPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HciNccCheckPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HciNccCheckPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityType

`func (o *HciNccCheckPolicy) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HciNccCheckPolicy) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HciNccCheckPolicy) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *HciNccCheckPolicy) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetImpactType

`func (o *HciNccCheckPolicy) GetImpactType() string`

GetImpactType returns the ImpactType field if non-nil, zero value otherwise.

### GetImpactTypeOk

`func (o *HciNccCheckPolicy) GetImpactTypeOk() (*string, bool)`

GetImpactTypeOk returns a tuple with the ImpactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactType

`func (o *HciNccCheckPolicy) SetImpactType(v string)`

SetImpactType sets ImpactType field to given value.

### HasImpactType

`func (o *HciNccCheckPolicy) HasImpactType() bool`

HasImpactType returns a boolean if a field has been set.

### GetKbArticles

`func (o *HciNccCheckPolicy) GetKbArticles() []string`

GetKbArticles returns the KbArticles field if non-nil, zero value otherwise.

### GetKbArticlesOk

`func (o *HciNccCheckPolicy) GetKbArticlesOk() (*[]string, bool)`

GetKbArticlesOk returns a tuple with the KbArticles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbArticles

`func (o *HciNccCheckPolicy) SetKbArticles(v []string)`

SetKbArticles sets KbArticles field to given value.

### HasKbArticles

`func (o *HciNccCheckPolicy) HasKbArticles() bool`

HasKbArticles returns a boolean if a field has been set.

### SetKbArticlesNil

`func (o *HciNccCheckPolicy) SetKbArticlesNil(b bool)`

 SetKbArticlesNil sets the value for KbArticles to be an explicit nil

### UnsetKbArticles
`func (o *HciNccCheckPolicy) UnsetKbArticles()`

UnsetKbArticles ensures that no value is present for KbArticles, not even an explicit nil
### GetName

`func (o *HciNccCheckPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciNccCheckPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciNccCheckPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciNccCheckPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNccPolicyExtId

`func (o *HciNccCheckPolicy) GetNccPolicyExtId() string`

GetNccPolicyExtId returns the NccPolicyExtId field if non-nil, zero value otherwise.

### GetNccPolicyExtIdOk

`func (o *HciNccCheckPolicy) GetNccPolicyExtIdOk() (*string, bool)`

GetNccPolicyExtIdOk returns a tuple with the NccPolicyExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNccPolicyExtId

`func (o *HciNccCheckPolicy) SetNccPolicyExtId(v string)`

SetNccPolicyExtId sets NccPolicyExtId field to given value.

### HasNccPolicyExtId

`func (o *HciNccCheckPolicy) HasNccPolicyExtId() bool`

HasNccPolicyExtId returns a boolean if a field has been set.

### GetPcExtId

`func (o *HciNccCheckPolicy) GetPcExtId() string`

GetPcExtId returns the PcExtId field if non-nil, zero value otherwise.

### GetPcExtIdOk

`func (o *HciNccCheckPolicy) GetPcExtIdOk() (*string, bool)`

GetPcExtIdOk returns a tuple with the PcExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcExtId

`func (o *HciNccCheckPolicy) SetPcExtId(v string)`

SetPcExtId sets PcExtId field to given value.

### HasPcExtId

`func (o *HciNccCheckPolicy) HasPcExtId() bool`

HasPcExtId returns a boolean if a field has been set.

### GetPolicyId

`func (o *HciNccCheckPolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *HciNccCheckPolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *HciNccCheckPolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *HciNccCheckPolicy) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPublisher

`func (o *HciNccCheckPolicy) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *HciNccCheckPolicy) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *HciNccCheckPolicy) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *HciNccCheckPolicy) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetScope

`func (o *HciNccCheckPolicy) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HciNccCheckPolicy) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HciNccCheckPolicy) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HciNccCheckPolicy) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSubType

`func (o *HciNccCheckPolicy) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *HciNccCheckPolicy) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *HciNccCheckPolicy) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *HciNccCheckPolicy) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetTargetClusters

`func (o *HciNccCheckPolicy) GetTargetClusters() []string`

GetTargetClusters returns the TargetClusters field if non-nil, zero value otherwise.

### GetTargetClustersOk

`func (o *HciNccCheckPolicy) GetTargetClustersOk() (*[]string, bool)`

GetTargetClustersOk returns a tuple with the TargetClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusters

`func (o *HciNccCheckPolicy) SetTargetClusters(v []string)`

SetTargetClusters sets TargetClusters field to given value.

### HasTargetClusters

`func (o *HciNccCheckPolicy) HasTargetClusters() bool`

HasTargetClusters returns a boolean if a field has been set.

### SetTargetClustersNil

`func (o *HciNccCheckPolicy) SetTargetClustersNil(b bool)`

 SetTargetClustersNil sets the value for TargetClusters to be an explicit nil

### UnsetTargetClusters
`func (o *HciNccCheckPolicy) UnsetTargetClusters()`

UnsetTargetClusters ensures that no value is present for TargetClusters, not even an explicit nil
### GetTitle

`func (o *HciNccCheckPolicy) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HciNccCheckPolicy) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HciNccCheckPolicy) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HciNccCheckPolicy) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *HciNccCheckPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciNccCheckPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciNccCheckPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciNccCheckPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDomainManager

`func (o *HciNccCheckPolicy) GetDomainManager() HciDomainManagerRelationship`

GetDomainManager returns the DomainManager field if non-nil, zero value otherwise.

### GetDomainManagerOk

`func (o *HciNccCheckPolicy) GetDomainManagerOk() (*HciDomainManagerRelationship, bool)`

GetDomainManagerOk returns a tuple with the DomainManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainManager

`func (o *HciNccCheckPolicy) SetDomainManager(v HciDomainManagerRelationship)`

SetDomainManager sets DomainManager field to given value.

### HasDomainManager

`func (o *HciNccCheckPolicy) HasDomainManager() bool`

HasDomainManager returns a boolean if a field has been set.

### SetDomainManagerNil

`func (o *HciNccCheckPolicy) SetDomainManagerNil(b bool)`

 SetDomainManagerNil sets the value for DomainManager to be an explicit nil

### UnsetDomainManager
`func (o *HciNccCheckPolicy) UnsetDomainManager()`

UnsetDomainManager ensures that no value is present for DomainManager, not even an explicit nil
### GetRegisteredDevice

`func (o *HciNccCheckPolicy) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciNccCheckPolicy) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciNccCheckPolicy) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciNccCheckPolicy) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciNccCheckPolicy) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciNccCheckPolicy) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


