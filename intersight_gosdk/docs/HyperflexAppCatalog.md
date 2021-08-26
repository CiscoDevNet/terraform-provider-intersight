# HyperflexAppCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.AppCatalog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.AppCatalog"]
**Version** | Pointer to **string** | The catalog version used in HyperFlex cluster configuration service. | [optional] 
**FeatureLimitExternal** | Pointer to [**HyperflexFeatureLimitExternalRelationship**](HyperflexFeatureLimitExternalRelationship.md) |  | [optional] 
**FeatureLimitInternal** | Pointer to [**HyperflexFeatureLimitInternalRelationship**](HyperflexFeatureLimitInternalRelationship.md) |  | [optional] 
**HxdpVersions** | Pointer to [**[]HyperflexHxdpVersionRelationship**](HyperflexHxdpVersionRelationship.md) | An array of relationships to hyperflexHxdpVersion resources. | [optional] 
**HyperflexCapabilityInfos** | Pointer to [**[]HyperflexCapabilityInfoRelationship**](HyperflexCapabilityInfoRelationship.md) | An array of relationships to hyperflexCapabilityInfo resources. | [optional] 
**HyperflexSoftwareCompatibilityInfos** | Pointer to [**[]HclHyperflexSoftwareCompatibilityInfoRelationship**](HclHyperflexSoftwareCompatibilityInfoRelationship.md) | An array of relationships to hclHyperflexSoftwareCompatibilityInfo resources. | [optional] 
**ServerFirmwareVersion** | Pointer to [**HyperflexServerFirmwareVersionRelationship**](HyperflexServerFirmwareVersionRelationship.md) |  | [optional] 
**ServerModel** | Pointer to [**HyperflexServerModelRelationship**](HyperflexServerModelRelationship.md) |  | [optional] 
**SoftwareDistributions** | Pointer to [**[]HyperflexSoftwareDistributionEntryRelationship**](HyperflexSoftwareDistributionEntryRelationship.md) | An array of relationships to hyperflexSoftwareDistributionEntry resources. | [optional] 

## Methods

### NewHyperflexAppCatalog

`func NewHyperflexAppCatalog(classId string, objectType string, ) *HyperflexAppCatalog`

NewHyperflexAppCatalog instantiates a new HyperflexAppCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAppCatalogWithDefaults

`func NewHyperflexAppCatalogWithDefaults() *HyperflexAppCatalog`

NewHyperflexAppCatalogWithDefaults instantiates a new HyperflexAppCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexAppCatalog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexAppCatalog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexAppCatalog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexAppCatalog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexAppCatalog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexAppCatalog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVersion

`func (o *HyperflexAppCatalog) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexAppCatalog) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexAppCatalog) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexAppCatalog) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFeatureLimitExternal

`func (o *HyperflexAppCatalog) GetFeatureLimitExternal() HyperflexFeatureLimitExternalRelationship`

GetFeatureLimitExternal returns the FeatureLimitExternal field if non-nil, zero value otherwise.

### GetFeatureLimitExternalOk

`func (o *HyperflexAppCatalog) GetFeatureLimitExternalOk() (*HyperflexFeatureLimitExternalRelationship, bool)`

GetFeatureLimitExternalOk returns a tuple with the FeatureLimitExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimitExternal

`func (o *HyperflexAppCatalog) SetFeatureLimitExternal(v HyperflexFeatureLimitExternalRelationship)`

SetFeatureLimitExternal sets FeatureLimitExternal field to given value.

### HasFeatureLimitExternal

`func (o *HyperflexAppCatalog) HasFeatureLimitExternal() bool`

HasFeatureLimitExternal returns a boolean if a field has been set.

### GetFeatureLimitInternal

`func (o *HyperflexAppCatalog) GetFeatureLimitInternal() HyperflexFeatureLimitInternalRelationship`

GetFeatureLimitInternal returns the FeatureLimitInternal field if non-nil, zero value otherwise.

### GetFeatureLimitInternalOk

`func (o *HyperflexAppCatalog) GetFeatureLimitInternalOk() (*HyperflexFeatureLimitInternalRelationship, bool)`

GetFeatureLimitInternalOk returns a tuple with the FeatureLimitInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimitInternal

`func (o *HyperflexAppCatalog) SetFeatureLimitInternal(v HyperflexFeatureLimitInternalRelationship)`

SetFeatureLimitInternal sets FeatureLimitInternal field to given value.

### HasFeatureLimitInternal

`func (o *HyperflexAppCatalog) HasFeatureLimitInternal() bool`

HasFeatureLimitInternal returns a boolean if a field has been set.

### GetHxdpVersions

`func (o *HyperflexAppCatalog) GetHxdpVersions() []HyperflexHxdpVersionRelationship`

GetHxdpVersions returns the HxdpVersions field if non-nil, zero value otherwise.

### GetHxdpVersionsOk

`func (o *HyperflexAppCatalog) GetHxdpVersionsOk() (*[]HyperflexHxdpVersionRelationship, bool)`

GetHxdpVersionsOk returns a tuple with the HxdpVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersions

`func (o *HyperflexAppCatalog) SetHxdpVersions(v []HyperflexHxdpVersionRelationship)`

SetHxdpVersions sets HxdpVersions field to given value.

### HasHxdpVersions

`func (o *HyperflexAppCatalog) HasHxdpVersions() bool`

HasHxdpVersions returns a boolean if a field has been set.

### SetHxdpVersionsNil

`func (o *HyperflexAppCatalog) SetHxdpVersionsNil(b bool)`

 SetHxdpVersionsNil sets the value for HxdpVersions to be an explicit nil

### UnsetHxdpVersions
`func (o *HyperflexAppCatalog) UnsetHxdpVersions()`

UnsetHxdpVersions ensures that no value is present for HxdpVersions, not even an explicit nil
### GetHyperflexCapabilityInfos

`func (o *HyperflexAppCatalog) GetHyperflexCapabilityInfos() []HyperflexCapabilityInfoRelationship`

GetHyperflexCapabilityInfos returns the HyperflexCapabilityInfos field if non-nil, zero value otherwise.

### GetHyperflexCapabilityInfosOk

`func (o *HyperflexAppCatalog) GetHyperflexCapabilityInfosOk() (*[]HyperflexCapabilityInfoRelationship, bool)`

GetHyperflexCapabilityInfosOk returns a tuple with the HyperflexCapabilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperflexCapabilityInfos

`func (o *HyperflexAppCatalog) SetHyperflexCapabilityInfos(v []HyperflexCapabilityInfoRelationship)`

SetHyperflexCapabilityInfos sets HyperflexCapabilityInfos field to given value.

### HasHyperflexCapabilityInfos

`func (o *HyperflexAppCatalog) HasHyperflexCapabilityInfos() bool`

HasHyperflexCapabilityInfos returns a boolean if a field has been set.

### SetHyperflexCapabilityInfosNil

`func (o *HyperflexAppCatalog) SetHyperflexCapabilityInfosNil(b bool)`

 SetHyperflexCapabilityInfosNil sets the value for HyperflexCapabilityInfos to be an explicit nil

### UnsetHyperflexCapabilityInfos
`func (o *HyperflexAppCatalog) UnsetHyperflexCapabilityInfos()`

UnsetHyperflexCapabilityInfos ensures that no value is present for HyperflexCapabilityInfos, not even an explicit nil
### GetHyperflexSoftwareCompatibilityInfos

`func (o *HyperflexAppCatalog) GetHyperflexSoftwareCompatibilityInfos() []HclHyperflexSoftwareCompatibilityInfoRelationship`

GetHyperflexSoftwareCompatibilityInfos returns the HyperflexSoftwareCompatibilityInfos field if non-nil, zero value otherwise.

### GetHyperflexSoftwareCompatibilityInfosOk

`func (o *HyperflexAppCatalog) GetHyperflexSoftwareCompatibilityInfosOk() (*[]HclHyperflexSoftwareCompatibilityInfoRelationship, bool)`

GetHyperflexSoftwareCompatibilityInfosOk returns a tuple with the HyperflexSoftwareCompatibilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperflexSoftwareCompatibilityInfos

`func (o *HyperflexAppCatalog) SetHyperflexSoftwareCompatibilityInfos(v []HclHyperflexSoftwareCompatibilityInfoRelationship)`

SetHyperflexSoftwareCompatibilityInfos sets HyperflexSoftwareCompatibilityInfos field to given value.

### HasHyperflexSoftwareCompatibilityInfos

`func (o *HyperflexAppCatalog) HasHyperflexSoftwareCompatibilityInfos() bool`

HasHyperflexSoftwareCompatibilityInfos returns a boolean if a field has been set.

### SetHyperflexSoftwareCompatibilityInfosNil

`func (o *HyperflexAppCatalog) SetHyperflexSoftwareCompatibilityInfosNil(b bool)`

 SetHyperflexSoftwareCompatibilityInfosNil sets the value for HyperflexSoftwareCompatibilityInfos to be an explicit nil

### UnsetHyperflexSoftwareCompatibilityInfos
`func (o *HyperflexAppCatalog) UnsetHyperflexSoftwareCompatibilityInfos()`

UnsetHyperflexSoftwareCompatibilityInfos ensures that no value is present for HyperflexSoftwareCompatibilityInfos, not even an explicit nil
### GetServerFirmwareVersion

`func (o *HyperflexAppCatalog) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship`

GetServerFirmwareVersion returns the ServerFirmwareVersion field if non-nil, zero value otherwise.

### GetServerFirmwareVersionOk

`func (o *HyperflexAppCatalog) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool)`

GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersion

`func (o *HyperflexAppCatalog) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship)`

SetServerFirmwareVersion sets ServerFirmwareVersion field to given value.

### HasServerFirmwareVersion

`func (o *HyperflexAppCatalog) HasServerFirmwareVersion() bool`

HasServerFirmwareVersion returns a boolean if a field has been set.

### GetServerModel

`func (o *HyperflexAppCatalog) GetServerModel() HyperflexServerModelRelationship`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *HyperflexAppCatalog) GetServerModelOk() (*HyperflexServerModelRelationship, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *HyperflexAppCatalog) SetServerModel(v HyperflexServerModelRelationship)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *HyperflexAppCatalog) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### GetSoftwareDistributions

`func (o *HyperflexAppCatalog) GetSoftwareDistributions() []HyperflexSoftwareDistributionEntryRelationship`

GetSoftwareDistributions returns the SoftwareDistributions field if non-nil, zero value otherwise.

### GetSoftwareDistributionsOk

`func (o *HyperflexAppCatalog) GetSoftwareDistributionsOk() (*[]HyperflexSoftwareDistributionEntryRelationship, bool)`

GetSoftwareDistributionsOk returns a tuple with the SoftwareDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDistributions

`func (o *HyperflexAppCatalog) SetSoftwareDistributions(v []HyperflexSoftwareDistributionEntryRelationship)`

SetSoftwareDistributions sets SoftwareDistributions field to given value.

### HasSoftwareDistributions

`func (o *HyperflexAppCatalog) HasSoftwareDistributions() bool`

HasSoftwareDistributions returns a boolean if a field has been set.

### SetSoftwareDistributionsNil

`func (o *HyperflexAppCatalog) SetSoftwareDistributionsNil(b bool)`

 SetSoftwareDistributionsNil sets the value for SoftwareDistributions to be an explicit nil

### UnsetSoftwareDistributions
`func (o *HyperflexAppCatalog) UnsetSoftwareDistributions()`

UnsetSoftwareDistributions ensures that no value is present for SoftwareDistributions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


