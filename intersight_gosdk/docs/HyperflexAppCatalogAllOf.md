# HyperflexAppCatalogAllOf

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

### NewHyperflexAppCatalogAllOf

`func NewHyperflexAppCatalogAllOf(classId string, objectType string, ) *HyperflexAppCatalogAllOf`

NewHyperflexAppCatalogAllOf instantiates a new HyperflexAppCatalogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAppCatalogAllOfWithDefaults

`func NewHyperflexAppCatalogAllOfWithDefaults() *HyperflexAppCatalogAllOf`

NewHyperflexAppCatalogAllOfWithDefaults instantiates a new HyperflexAppCatalogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexAppCatalogAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexAppCatalogAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexAppCatalogAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexAppCatalogAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexAppCatalogAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexAppCatalogAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVersion

`func (o *HyperflexAppCatalogAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexAppCatalogAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexAppCatalogAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexAppCatalogAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFeatureLimitExternal

`func (o *HyperflexAppCatalogAllOf) GetFeatureLimitExternal() HyperflexFeatureLimitExternalRelationship`

GetFeatureLimitExternal returns the FeatureLimitExternal field if non-nil, zero value otherwise.

### GetFeatureLimitExternalOk

`func (o *HyperflexAppCatalogAllOf) GetFeatureLimitExternalOk() (*HyperflexFeatureLimitExternalRelationship, bool)`

GetFeatureLimitExternalOk returns a tuple with the FeatureLimitExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimitExternal

`func (o *HyperflexAppCatalogAllOf) SetFeatureLimitExternal(v HyperflexFeatureLimitExternalRelationship)`

SetFeatureLimitExternal sets FeatureLimitExternal field to given value.

### HasFeatureLimitExternal

`func (o *HyperflexAppCatalogAllOf) HasFeatureLimitExternal() bool`

HasFeatureLimitExternal returns a boolean if a field has been set.

### GetFeatureLimitInternal

`func (o *HyperflexAppCatalogAllOf) GetFeatureLimitInternal() HyperflexFeatureLimitInternalRelationship`

GetFeatureLimitInternal returns the FeatureLimitInternal field if non-nil, zero value otherwise.

### GetFeatureLimitInternalOk

`func (o *HyperflexAppCatalogAllOf) GetFeatureLimitInternalOk() (*HyperflexFeatureLimitInternalRelationship, bool)`

GetFeatureLimitInternalOk returns a tuple with the FeatureLimitInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimitInternal

`func (o *HyperflexAppCatalogAllOf) SetFeatureLimitInternal(v HyperflexFeatureLimitInternalRelationship)`

SetFeatureLimitInternal sets FeatureLimitInternal field to given value.

### HasFeatureLimitInternal

`func (o *HyperflexAppCatalogAllOf) HasFeatureLimitInternal() bool`

HasFeatureLimitInternal returns a boolean if a field has been set.

### GetHxdpVersions

`func (o *HyperflexAppCatalogAllOf) GetHxdpVersions() []HyperflexHxdpVersionRelationship`

GetHxdpVersions returns the HxdpVersions field if non-nil, zero value otherwise.

### GetHxdpVersionsOk

`func (o *HyperflexAppCatalogAllOf) GetHxdpVersionsOk() (*[]HyperflexHxdpVersionRelationship, bool)`

GetHxdpVersionsOk returns a tuple with the HxdpVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersions

`func (o *HyperflexAppCatalogAllOf) SetHxdpVersions(v []HyperflexHxdpVersionRelationship)`

SetHxdpVersions sets HxdpVersions field to given value.

### HasHxdpVersions

`func (o *HyperflexAppCatalogAllOf) HasHxdpVersions() bool`

HasHxdpVersions returns a boolean if a field has been set.

### SetHxdpVersionsNil

`func (o *HyperflexAppCatalogAllOf) SetHxdpVersionsNil(b bool)`

 SetHxdpVersionsNil sets the value for HxdpVersions to be an explicit nil

### UnsetHxdpVersions
`func (o *HyperflexAppCatalogAllOf) UnsetHxdpVersions()`

UnsetHxdpVersions ensures that no value is present for HxdpVersions, not even an explicit nil
### GetHyperflexCapabilityInfos

`func (o *HyperflexAppCatalogAllOf) GetHyperflexCapabilityInfos() []HyperflexCapabilityInfoRelationship`

GetHyperflexCapabilityInfos returns the HyperflexCapabilityInfos field if non-nil, zero value otherwise.

### GetHyperflexCapabilityInfosOk

`func (o *HyperflexAppCatalogAllOf) GetHyperflexCapabilityInfosOk() (*[]HyperflexCapabilityInfoRelationship, bool)`

GetHyperflexCapabilityInfosOk returns a tuple with the HyperflexCapabilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperflexCapabilityInfos

`func (o *HyperflexAppCatalogAllOf) SetHyperflexCapabilityInfos(v []HyperflexCapabilityInfoRelationship)`

SetHyperflexCapabilityInfos sets HyperflexCapabilityInfos field to given value.

### HasHyperflexCapabilityInfos

`func (o *HyperflexAppCatalogAllOf) HasHyperflexCapabilityInfos() bool`

HasHyperflexCapabilityInfos returns a boolean if a field has been set.

### SetHyperflexCapabilityInfosNil

`func (o *HyperflexAppCatalogAllOf) SetHyperflexCapabilityInfosNil(b bool)`

 SetHyperflexCapabilityInfosNil sets the value for HyperflexCapabilityInfos to be an explicit nil

### UnsetHyperflexCapabilityInfos
`func (o *HyperflexAppCatalogAllOf) UnsetHyperflexCapabilityInfos()`

UnsetHyperflexCapabilityInfos ensures that no value is present for HyperflexCapabilityInfos, not even an explicit nil
### GetHyperflexSoftwareCompatibilityInfos

`func (o *HyperflexAppCatalogAllOf) GetHyperflexSoftwareCompatibilityInfos() []HclHyperflexSoftwareCompatibilityInfoRelationship`

GetHyperflexSoftwareCompatibilityInfos returns the HyperflexSoftwareCompatibilityInfos field if non-nil, zero value otherwise.

### GetHyperflexSoftwareCompatibilityInfosOk

`func (o *HyperflexAppCatalogAllOf) GetHyperflexSoftwareCompatibilityInfosOk() (*[]HclHyperflexSoftwareCompatibilityInfoRelationship, bool)`

GetHyperflexSoftwareCompatibilityInfosOk returns a tuple with the HyperflexSoftwareCompatibilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperflexSoftwareCompatibilityInfos

`func (o *HyperflexAppCatalogAllOf) SetHyperflexSoftwareCompatibilityInfos(v []HclHyperflexSoftwareCompatibilityInfoRelationship)`

SetHyperflexSoftwareCompatibilityInfos sets HyperflexSoftwareCompatibilityInfos field to given value.

### HasHyperflexSoftwareCompatibilityInfos

`func (o *HyperflexAppCatalogAllOf) HasHyperflexSoftwareCompatibilityInfos() bool`

HasHyperflexSoftwareCompatibilityInfos returns a boolean if a field has been set.

### SetHyperflexSoftwareCompatibilityInfosNil

`func (o *HyperflexAppCatalogAllOf) SetHyperflexSoftwareCompatibilityInfosNil(b bool)`

 SetHyperflexSoftwareCompatibilityInfosNil sets the value for HyperflexSoftwareCompatibilityInfos to be an explicit nil

### UnsetHyperflexSoftwareCompatibilityInfos
`func (o *HyperflexAppCatalogAllOf) UnsetHyperflexSoftwareCompatibilityInfos()`

UnsetHyperflexSoftwareCompatibilityInfos ensures that no value is present for HyperflexSoftwareCompatibilityInfos, not even an explicit nil
### GetServerFirmwareVersion

`func (o *HyperflexAppCatalogAllOf) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship`

GetServerFirmwareVersion returns the ServerFirmwareVersion field if non-nil, zero value otherwise.

### GetServerFirmwareVersionOk

`func (o *HyperflexAppCatalogAllOf) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool)`

GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersion

`func (o *HyperflexAppCatalogAllOf) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship)`

SetServerFirmwareVersion sets ServerFirmwareVersion field to given value.

### HasServerFirmwareVersion

`func (o *HyperflexAppCatalogAllOf) HasServerFirmwareVersion() bool`

HasServerFirmwareVersion returns a boolean if a field has been set.

### GetServerModel

`func (o *HyperflexAppCatalogAllOf) GetServerModel() HyperflexServerModelRelationship`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *HyperflexAppCatalogAllOf) GetServerModelOk() (*HyperflexServerModelRelationship, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *HyperflexAppCatalogAllOf) SetServerModel(v HyperflexServerModelRelationship)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *HyperflexAppCatalogAllOf) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### GetSoftwareDistributions

`func (o *HyperflexAppCatalogAllOf) GetSoftwareDistributions() []HyperflexSoftwareDistributionEntryRelationship`

GetSoftwareDistributions returns the SoftwareDistributions field if non-nil, zero value otherwise.

### GetSoftwareDistributionsOk

`func (o *HyperflexAppCatalogAllOf) GetSoftwareDistributionsOk() (*[]HyperflexSoftwareDistributionEntryRelationship, bool)`

GetSoftwareDistributionsOk returns a tuple with the SoftwareDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDistributions

`func (o *HyperflexAppCatalogAllOf) SetSoftwareDistributions(v []HyperflexSoftwareDistributionEntryRelationship)`

SetSoftwareDistributions sets SoftwareDistributions field to given value.

### HasSoftwareDistributions

`func (o *HyperflexAppCatalogAllOf) HasSoftwareDistributions() bool`

HasSoftwareDistributions returns a boolean if a field has been set.

### SetSoftwareDistributionsNil

`func (o *HyperflexAppCatalogAllOf) SetSoftwareDistributionsNil(b bool)`

 SetSoftwareDistributionsNil sets the value for SoftwareDistributions to be an explicit nil

### UnsetSoftwareDistributions
`func (o *HyperflexAppCatalogAllOf) UnsetSoftwareDistributions()`

UnsetSoftwareDistributions ensures that no value is present for SoftwareDistributions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


