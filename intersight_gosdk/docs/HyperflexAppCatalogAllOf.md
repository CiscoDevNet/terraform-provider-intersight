# HyperflexAppCatalogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The catalog version used in HyperFlex cluster configuration service. | [optional] 
**FeatureLimitExternal** | Pointer to [**HyperflexFeatureLimitExternalRelationship**](hyperflex.FeatureLimitExternal.Relationship.md) |  | [optional] 
**FeatureLimitInternal** | Pointer to [**HyperflexFeatureLimitInternalRelationship**](hyperflex.FeatureLimitInternal.Relationship.md) |  | [optional] 
**HxdpVersions** | Pointer to [**[]HyperflexHxdpVersionRelationship**](hyperflex.HxdpVersion.Relationship.md) | An array of relationships to hyperflexHxdpVersion resources. | [optional] 
**HyperflexCapabilityInfos** | Pointer to [**[]HyperflexCapabilityInfoRelationship**](hyperflex.CapabilityInfo.Relationship.md) | An array of relationships to hyperflexCapabilityInfo resources. | [optional] 
**HyperflexSoftwareCompatibilityInfos** | Pointer to [**[]HclHyperflexSoftwareCompatibilityInfoRelationship**](hcl.HyperflexSoftwareCompatibilityInfo.Relationship.md) | An array of relationships to hclHyperflexSoftwareCompatibilityInfo resources. | [optional] 
**ServerFirmwareVersion** | Pointer to [**HyperflexServerFirmwareVersionRelationship**](hyperflex.ServerFirmwareVersion.Relationship.md) |  | [optional] 
**ServerModel** | Pointer to [**HyperflexServerModelRelationship**](hyperflex.ServerModel.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexAppCatalogAllOf

`func NewHyperflexAppCatalogAllOf() *HyperflexAppCatalogAllOf`

NewHyperflexAppCatalogAllOf instantiates a new HyperflexAppCatalogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAppCatalogAllOfWithDefaults

`func NewHyperflexAppCatalogAllOfWithDefaults() *HyperflexAppCatalogAllOf`

NewHyperflexAppCatalogAllOfWithDefaults instantiates a new HyperflexAppCatalogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


