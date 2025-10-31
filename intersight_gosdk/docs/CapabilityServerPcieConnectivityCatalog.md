# CapabilityServerPcieConnectivityCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerPcieConnectivityCatalog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerPcieConnectivityCatalog"]
**SupportedLayouts** | Pointer to [**[]CapabilityServerPcieConnectivityLayout**](CapabilityServerPcieConnectivityLayout.md) |  | [optional] 
**SupportedNumberOfGpus** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCapabilityServerPcieConnectivityCatalog

`func NewCapabilityServerPcieConnectivityCatalog(classId string, objectType string, ) *CapabilityServerPcieConnectivityCatalog`

NewCapabilityServerPcieConnectivityCatalog instantiates a new CapabilityServerPcieConnectivityCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerPcieConnectivityCatalogWithDefaults

`func NewCapabilityServerPcieConnectivityCatalogWithDefaults() *CapabilityServerPcieConnectivityCatalog`

NewCapabilityServerPcieConnectivityCatalogWithDefaults instantiates a new CapabilityServerPcieConnectivityCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerPcieConnectivityCatalog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerPcieConnectivityCatalog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerPcieConnectivityCatalog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerPcieConnectivityCatalog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerPcieConnectivityCatalog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerPcieConnectivityCatalog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSupportedLayouts

`func (o *CapabilityServerPcieConnectivityCatalog) GetSupportedLayouts() []CapabilityServerPcieConnectivityLayout`

GetSupportedLayouts returns the SupportedLayouts field if non-nil, zero value otherwise.

### GetSupportedLayoutsOk

`func (o *CapabilityServerPcieConnectivityCatalog) GetSupportedLayoutsOk() (*[]CapabilityServerPcieConnectivityLayout, bool)`

GetSupportedLayoutsOk returns a tuple with the SupportedLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedLayouts

`func (o *CapabilityServerPcieConnectivityCatalog) SetSupportedLayouts(v []CapabilityServerPcieConnectivityLayout)`

SetSupportedLayouts sets SupportedLayouts field to given value.

### HasSupportedLayouts

`func (o *CapabilityServerPcieConnectivityCatalog) HasSupportedLayouts() bool`

HasSupportedLayouts returns a boolean if a field has been set.

### SetSupportedLayoutsNil

`func (o *CapabilityServerPcieConnectivityCatalog) SetSupportedLayoutsNil(b bool)`

 SetSupportedLayoutsNil sets the value for SupportedLayouts to be an explicit nil

### UnsetSupportedLayouts
`func (o *CapabilityServerPcieConnectivityCatalog) UnsetSupportedLayouts()`

UnsetSupportedLayouts ensures that no value is present for SupportedLayouts, not even an explicit nil
### GetSupportedNumberOfGpus

`func (o *CapabilityServerPcieConnectivityCatalog) GetSupportedNumberOfGpus() []int64`

GetSupportedNumberOfGpus returns the SupportedNumberOfGpus field if non-nil, zero value otherwise.

### GetSupportedNumberOfGpusOk

`func (o *CapabilityServerPcieConnectivityCatalog) GetSupportedNumberOfGpusOk() (*[]int64, bool)`

GetSupportedNumberOfGpusOk returns a tuple with the SupportedNumberOfGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedNumberOfGpus

`func (o *CapabilityServerPcieConnectivityCatalog) SetSupportedNumberOfGpus(v []int64)`

SetSupportedNumberOfGpus sets SupportedNumberOfGpus field to given value.

### HasSupportedNumberOfGpus

`func (o *CapabilityServerPcieConnectivityCatalog) HasSupportedNumberOfGpus() bool`

HasSupportedNumberOfGpus returns a boolean if a field has been set.

### SetSupportedNumberOfGpusNil

`func (o *CapabilityServerPcieConnectivityCatalog) SetSupportedNumberOfGpusNil(b bool)`

 SetSupportedNumberOfGpusNil sets the value for SupportedNumberOfGpus to be an explicit nil

### UnsetSupportedNumberOfGpus
`func (o *CapabilityServerPcieConnectivityCatalog) UnsetSupportedNumberOfGpus()`

UnsetSupportedNumberOfGpus ensures that no value is present for SupportedNumberOfGpus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


