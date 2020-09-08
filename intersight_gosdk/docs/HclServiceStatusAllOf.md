# HclServiceStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExemptionFileVersion** | Pointer to **string** | Version of the last modified exemption file. | [optional] 
**Identity** | Pointer to **string** | A field to uniquely identify the document with the status. | [optional] 
**LastHclDataModifiedTime** | Pointer to [**time.Time**](time.Time.md) | The timestamp of the last modified record in the HCL tool database. Used to query and get updated records. | [optional] 
**LastImportedDataChecksum** | Pointer to **string** | Checksum of the data dump used as the base for delta updates. | [optional] 
**Status** | Pointer to **string** | Status of the service indicatating if the service is up or under maintenance due to data update. * &#x60;Unknown&#x60; - Default service status. Indicates that the full import of data has failed during service startup. * &#x60;Initializing&#x60; - The service starts or restarts. * &#x60;DataRefreshing&#x60; - Full import of data is in progress. * &#x60;Active&#x60; - The data import is successful and service is ready to serve API requests. | [optional] [default to "Unknown"]

## Methods

### NewHclServiceStatusAllOf

`func NewHclServiceStatusAllOf() *HclServiceStatusAllOf`

NewHclServiceStatusAllOf instantiates a new HclServiceStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclServiceStatusAllOfWithDefaults

`func NewHclServiceStatusAllOfWithDefaults() *HclServiceStatusAllOf`

NewHclServiceStatusAllOfWithDefaults instantiates a new HclServiceStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExemptionFileVersion

`func (o *HclServiceStatusAllOf) GetExemptionFileVersion() string`

GetExemptionFileVersion returns the ExemptionFileVersion field if non-nil, zero value otherwise.

### GetExemptionFileVersionOk

`func (o *HclServiceStatusAllOf) GetExemptionFileVersionOk() (*string, bool)`

GetExemptionFileVersionOk returns a tuple with the ExemptionFileVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptionFileVersion

`func (o *HclServiceStatusAllOf) SetExemptionFileVersion(v string)`

SetExemptionFileVersion sets ExemptionFileVersion field to given value.

### HasExemptionFileVersion

`func (o *HclServiceStatusAllOf) HasExemptionFileVersion() bool`

HasExemptionFileVersion returns a boolean if a field has been set.

### GetIdentity

`func (o *HclServiceStatusAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *HclServiceStatusAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *HclServiceStatusAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *HclServiceStatusAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLastHclDataModifiedTime

`func (o *HclServiceStatusAllOf) GetLastHclDataModifiedTime() time.Time`

GetLastHclDataModifiedTime returns the LastHclDataModifiedTime field if non-nil, zero value otherwise.

### GetLastHclDataModifiedTimeOk

`func (o *HclServiceStatusAllOf) GetLastHclDataModifiedTimeOk() (*time.Time, bool)`

GetLastHclDataModifiedTimeOk returns a tuple with the LastHclDataModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHclDataModifiedTime

`func (o *HclServiceStatusAllOf) SetLastHclDataModifiedTime(v time.Time)`

SetLastHclDataModifiedTime sets LastHclDataModifiedTime field to given value.

### HasLastHclDataModifiedTime

`func (o *HclServiceStatusAllOf) HasLastHclDataModifiedTime() bool`

HasLastHclDataModifiedTime returns a boolean if a field has been set.

### GetLastImportedDataChecksum

`func (o *HclServiceStatusAllOf) GetLastImportedDataChecksum() string`

GetLastImportedDataChecksum returns the LastImportedDataChecksum field if non-nil, zero value otherwise.

### GetLastImportedDataChecksumOk

`func (o *HclServiceStatusAllOf) GetLastImportedDataChecksumOk() (*string, bool)`

GetLastImportedDataChecksumOk returns a tuple with the LastImportedDataChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastImportedDataChecksum

`func (o *HclServiceStatusAllOf) SetLastImportedDataChecksum(v string)`

SetLastImportedDataChecksum sets LastImportedDataChecksum field to given value.

### HasLastImportedDataChecksum

`func (o *HclServiceStatusAllOf) HasLastImportedDataChecksum() bool`

HasLastImportedDataChecksum returns a boolean if a field has been set.

### GetStatus

`func (o *HclServiceStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HclServiceStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HclServiceStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HclServiceStatusAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


