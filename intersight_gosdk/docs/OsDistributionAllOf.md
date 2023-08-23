# OsDistributionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.Distribution"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.Distribution"]
**IsSupported** | Pointer to **bool** | An internal property that is used to denote if the OS Distribution is supported by Intersight for Automated Installation. | [optional] [readonly] 
**Label** | Pointer to **string** | The label of the OS distribution such as ESXi, CentOS to be displayed. | [optional] 
**Name** | Pointer to **string** | The name of the OS distribution such as ESXi, CentOS. | [optional] 
**ScuSupported** | Pointer to **bool** | An internal property that is used to denote if the OS Distribution is supported by the Server Configuration Utility. | [optional] [readonly] 
**SupportedEditions** | Pointer to **[]string** |  | [optional] 
**Catalog** | Pointer to [**OsCatalogRelationship**](OsCatalogRelationship.md) |  | [optional] 
**Vendor** | Pointer to [**HclOperatingSystemVendorRelationship**](HclOperatingSystemVendorRelationship.md) |  | [optional] 
**Version** | Pointer to [**HclOperatingSystemRelationship**](HclOperatingSystemRelationship.md) |  | [optional] 

## Methods

### NewOsDistributionAllOf

`func NewOsDistributionAllOf(classId string, objectType string, ) *OsDistributionAllOf`

NewOsDistributionAllOf instantiates a new OsDistributionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsDistributionAllOfWithDefaults

`func NewOsDistributionAllOfWithDefaults() *OsDistributionAllOf`

NewOsDistributionAllOfWithDefaults instantiates a new OsDistributionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsDistributionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsDistributionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsDistributionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsDistributionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsDistributionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsDistributionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsSupported

`func (o *OsDistributionAllOf) GetIsSupported() bool`

GetIsSupported returns the IsSupported field if non-nil, zero value otherwise.

### GetIsSupportedOk

`func (o *OsDistributionAllOf) GetIsSupportedOk() (*bool, bool)`

GetIsSupportedOk returns a tuple with the IsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupported

`func (o *OsDistributionAllOf) SetIsSupported(v bool)`

SetIsSupported sets IsSupported field to given value.

### HasIsSupported

`func (o *OsDistributionAllOf) HasIsSupported() bool`

HasIsSupported returns a boolean if a field has been set.

### GetLabel

`func (o *OsDistributionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OsDistributionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OsDistributionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *OsDistributionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *OsDistributionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsDistributionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsDistributionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsDistributionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScuSupported

`func (o *OsDistributionAllOf) GetScuSupported() bool`

GetScuSupported returns the ScuSupported field if non-nil, zero value otherwise.

### GetScuSupportedOk

`func (o *OsDistributionAllOf) GetScuSupportedOk() (*bool, bool)`

GetScuSupportedOk returns a tuple with the ScuSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScuSupported

`func (o *OsDistributionAllOf) SetScuSupported(v bool)`

SetScuSupported sets ScuSupported field to given value.

### HasScuSupported

`func (o *OsDistributionAllOf) HasScuSupported() bool`

HasScuSupported returns a boolean if a field has been set.

### GetSupportedEditions

`func (o *OsDistributionAllOf) GetSupportedEditions() []string`

GetSupportedEditions returns the SupportedEditions field if non-nil, zero value otherwise.

### GetSupportedEditionsOk

`func (o *OsDistributionAllOf) GetSupportedEditionsOk() (*[]string, bool)`

GetSupportedEditionsOk returns a tuple with the SupportedEditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEditions

`func (o *OsDistributionAllOf) SetSupportedEditions(v []string)`

SetSupportedEditions sets SupportedEditions field to given value.

### HasSupportedEditions

`func (o *OsDistributionAllOf) HasSupportedEditions() bool`

HasSupportedEditions returns a boolean if a field has been set.

### SetSupportedEditionsNil

`func (o *OsDistributionAllOf) SetSupportedEditionsNil(b bool)`

 SetSupportedEditionsNil sets the value for SupportedEditions to be an explicit nil

### UnsetSupportedEditions
`func (o *OsDistributionAllOf) UnsetSupportedEditions()`

UnsetSupportedEditions ensures that no value is present for SupportedEditions, not even an explicit nil
### GetCatalog

`func (o *OsDistributionAllOf) GetCatalog() OsCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *OsDistributionAllOf) GetCatalogOk() (*OsCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *OsDistributionAllOf) SetCatalog(v OsCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *OsDistributionAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetVendor

`func (o *OsDistributionAllOf) GetVendor() HclOperatingSystemVendorRelationship`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *OsDistributionAllOf) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *OsDistributionAllOf) SetVendor(v HclOperatingSystemVendorRelationship)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *OsDistributionAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *OsDistributionAllOf) GetVersion() HclOperatingSystemRelationship`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OsDistributionAllOf) GetVersionOk() (*HclOperatingSystemRelationship, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OsDistributionAllOf) SetVersion(v HclOperatingSystemRelationship)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OsDistributionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


