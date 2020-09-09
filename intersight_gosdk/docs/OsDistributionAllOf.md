# OsDistributionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the OS distribution such as ESXi, CentOS. | [optional] 
**SupportedEditions** | Pointer to **[]string** |  | [optional] 
**Catalog** | Pointer to [**OsCatalogRelationship**](os.Catalog.Relationship.md) |  | [optional] 
**Version** | Pointer to [**HclOperatingSystemRelationship**](hcl.OperatingSystem.Relationship.md) |  | [optional] 

## Methods

### NewOsDistributionAllOf

`func NewOsDistributionAllOf() *OsDistributionAllOf`

NewOsDistributionAllOf instantiates a new OsDistributionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsDistributionAllOfWithDefaults

`func NewOsDistributionAllOfWithDefaults() *OsDistributionAllOf`

NewOsDistributionAllOfWithDefaults instantiates a new OsDistributionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


