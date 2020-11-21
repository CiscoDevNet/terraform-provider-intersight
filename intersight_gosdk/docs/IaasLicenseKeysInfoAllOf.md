# IaasLicenseKeysInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.LicenseKeysInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.LicenseKeysInfo"]
**Count** | Pointer to **int64** | Number of licenses available for the UCSD PID (Product ID). | [optional] [readonly] 
**ExpirationDate** | Pointer to **string** | Expiration date for the license. | [optional] [readonly] 
**LicenseId** | Pointer to **string** | UCS Director Unique license ID. | [optional] [readonly] 
**Pid** | Pointer to **string** | PID (Product ID) for UCSD License. | [optional] [readonly] 

## Methods

### NewIaasLicenseKeysInfoAllOf

`func NewIaasLicenseKeysInfoAllOf(classId string, objectType string, ) *IaasLicenseKeysInfoAllOf`

NewIaasLicenseKeysInfoAllOf instantiates a new IaasLicenseKeysInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasLicenseKeysInfoAllOfWithDefaults

`func NewIaasLicenseKeysInfoAllOfWithDefaults() *IaasLicenseKeysInfoAllOf`

NewIaasLicenseKeysInfoAllOfWithDefaults instantiates a new IaasLicenseKeysInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasLicenseKeysInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasLicenseKeysInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasLicenseKeysInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasLicenseKeysInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasLicenseKeysInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasLicenseKeysInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *IaasLicenseKeysInfoAllOf) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IaasLicenseKeysInfoAllOf) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IaasLicenseKeysInfoAllOf) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *IaasLicenseKeysInfoAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetExpirationDate

`func (o *IaasLicenseKeysInfoAllOf) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *IaasLicenseKeysInfoAllOf) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *IaasLicenseKeysInfoAllOf) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *IaasLicenseKeysInfoAllOf) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLicenseId

`func (o *IaasLicenseKeysInfoAllOf) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *IaasLicenseKeysInfoAllOf) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *IaasLicenseKeysInfoAllOf) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *IaasLicenseKeysInfoAllOf) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetPid

`func (o *IaasLicenseKeysInfoAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *IaasLicenseKeysInfoAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *IaasLicenseKeysInfoAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *IaasLicenseKeysInfoAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


