# IwotenantTenantCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iwotenant.TenantCustomization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iwotenant.TenantCustomization"]
**EnableDataExtractor** | Pointer to **bool** | Enables IWO tenant data to be available for reporting.  This will start &#39;extractor&#39; pod. | [optional] 
**IsWriteUserAccessKeyIdSet** | Pointer to **bool** | Indicates whether the value of the &#39;writeUserAccessKeyId&#39; property has been set. | [optional] [readonly] [default to false]
**IsWriteUserSecretAccessKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;writeUserSecretAccessKey&#39; property has been set. | [optional] [readonly] [default to false]
**IwoId** | Pointer to **string** | The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId. | [optional] 
**MskServerForDataExtractor** | Pointer to **string** | MSK cluster endpoint that data extractor can send reporting data to. This  MS cluster in turn populates data into tables in Redshift cluster. | [optional] 
**WriteUserAccessKeyId** | Pointer to **string** | AWS access key Id to write data to redshift.  Refer to AWS cloud formation stack &#39;Output&#39; of the tenant. | [optional] 
**WriteUserSecretAccessKey** | Pointer to **string** | AWS secret access key to write data to redshift.  Refer to AWS cloud formation stack &#39;Output&#39; of the tenant. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIwotenantTenantCustomization

`func NewIwotenantTenantCustomization(classId string, objectType string, ) *IwotenantTenantCustomization`

NewIwotenantTenantCustomization instantiates a new IwotenantTenantCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIwotenantTenantCustomizationWithDefaults

`func NewIwotenantTenantCustomizationWithDefaults() *IwotenantTenantCustomization`

NewIwotenantTenantCustomizationWithDefaults instantiates a new IwotenantTenantCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IwotenantTenantCustomization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IwotenantTenantCustomization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IwotenantTenantCustomization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IwotenantTenantCustomization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IwotenantTenantCustomization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IwotenantTenantCustomization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableDataExtractor

`func (o *IwotenantTenantCustomization) GetEnableDataExtractor() bool`

GetEnableDataExtractor returns the EnableDataExtractor field if non-nil, zero value otherwise.

### GetEnableDataExtractorOk

`func (o *IwotenantTenantCustomization) GetEnableDataExtractorOk() (*bool, bool)`

GetEnableDataExtractorOk returns a tuple with the EnableDataExtractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDataExtractor

`func (o *IwotenantTenantCustomization) SetEnableDataExtractor(v bool)`

SetEnableDataExtractor sets EnableDataExtractor field to given value.

### HasEnableDataExtractor

`func (o *IwotenantTenantCustomization) HasEnableDataExtractor() bool`

HasEnableDataExtractor returns a boolean if a field has been set.

### GetIsWriteUserAccessKeyIdSet

`func (o *IwotenantTenantCustomization) GetIsWriteUserAccessKeyIdSet() bool`

GetIsWriteUserAccessKeyIdSet returns the IsWriteUserAccessKeyIdSet field if non-nil, zero value otherwise.

### GetIsWriteUserAccessKeyIdSetOk

`func (o *IwotenantTenantCustomization) GetIsWriteUserAccessKeyIdSetOk() (*bool, bool)`

GetIsWriteUserAccessKeyIdSetOk returns a tuple with the IsWriteUserAccessKeyIdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWriteUserAccessKeyIdSet

`func (o *IwotenantTenantCustomization) SetIsWriteUserAccessKeyIdSet(v bool)`

SetIsWriteUserAccessKeyIdSet sets IsWriteUserAccessKeyIdSet field to given value.

### HasIsWriteUserAccessKeyIdSet

`func (o *IwotenantTenantCustomization) HasIsWriteUserAccessKeyIdSet() bool`

HasIsWriteUserAccessKeyIdSet returns a boolean if a field has been set.

### GetIsWriteUserSecretAccessKeySet

`func (o *IwotenantTenantCustomization) GetIsWriteUserSecretAccessKeySet() bool`

GetIsWriteUserSecretAccessKeySet returns the IsWriteUserSecretAccessKeySet field if non-nil, zero value otherwise.

### GetIsWriteUserSecretAccessKeySetOk

`func (o *IwotenantTenantCustomization) GetIsWriteUserSecretAccessKeySetOk() (*bool, bool)`

GetIsWriteUserSecretAccessKeySetOk returns a tuple with the IsWriteUserSecretAccessKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWriteUserSecretAccessKeySet

`func (o *IwotenantTenantCustomization) SetIsWriteUserSecretAccessKeySet(v bool)`

SetIsWriteUserSecretAccessKeySet sets IsWriteUserSecretAccessKeySet field to given value.

### HasIsWriteUserSecretAccessKeySet

`func (o *IwotenantTenantCustomization) HasIsWriteUserSecretAccessKeySet() bool`

HasIsWriteUserSecretAccessKeySet returns a boolean if a field has been set.

### GetIwoId

`func (o *IwotenantTenantCustomization) GetIwoId() string`

GetIwoId returns the IwoId field if non-nil, zero value otherwise.

### GetIwoIdOk

`func (o *IwotenantTenantCustomization) GetIwoIdOk() (*string, bool)`

GetIwoIdOk returns a tuple with the IwoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwoId

`func (o *IwotenantTenantCustomization) SetIwoId(v string)`

SetIwoId sets IwoId field to given value.

### HasIwoId

`func (o *IwotenantTenantCustomization) HasIwoId() bool`

HasIwoId returns a boolean if a field has been set.

### GetMskServerForDataExtractor

`func (o *IwotenantTenantCustomization) GetMskServerForDataExtractor() string`

GetMskServerForDataExtractor returns the MskServerForDataExtractor field if non-nil, zero value otherwise.

### GetMskServerForDataExtractorOk

`func (o *IwotenantTenantCustomization) GetMskServerForDataExtractorOk() (*string, bool)`

GetMskServerForDataExtractorOk returns a tuple with the MskServerForDataExtractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMskServerForDataExtractor

`func (o *IwotenantTenantCustomization) SetMskServerForDataExtractor(v string)`

SetMskServerForDataExtractor sets MskServerForDataExtractor field to given value.

### HasMskServerForDataExtractor

`func (o *IwotenantTenantCustomization) HasMskServerForDataExtractor() bool`

HasMskServerForDataExtractor returns a boolean if a field has been set.

### GetWriteUserAccessKeyId

`func (o *IwotenantTenantCustomization) GetWriteUserAccessKeyId() string`

GetWriteUserAccessKeyId returns the WriteUserAccessKeyId field if non-nil, zero value otherwise.

### GetWriteUserAccessKeyIdOk

`func (o *IwotenantTenantCustomization) GetWriteUserAccessKeyIdOk() (*string, bool)`

GetWriteUserAccessKeyIdOk returns a tuple with the WriteUserAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteUserAccessKeyId

`func (o *IwotenantTenantCustomization) SetWriteUserAccessKeyId(v string)`

SetWriteUserAccessKeyId sets WriteUserAccessKeyId field to given value.

### HasWriteUserAccessKeyId

`func (o *IwotenantTenantCustomization) HasWriteUserAccessKeyId() bool`

HasWriteUserAccessKeyId returns a boolean if a field has been set.

### GetWriteUserSecretAccessKey

`func (o *IwotenantTenantCustomization) GetWriteUserSecretAccessKey() string`

GetWriteUserSecretAccessKey returns the WriteUserSecretAccessKey field if non-nil, zero value otherwise.

### GetWriteUserSecretAccessKeyOk

`func (o *IwotenantTenantCustomization) GetWriteUserSecretAccessKeyOk() (*string, bool)`

GetWriteUserSecretAccessKeyOk returns a tuple with the WriteUserSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteUserSecretAccessKey

`func (o *IwotenantTenantCustomization) SetWriteUserSecretAccessKey(v string)`

SetWriteUserSecretAccessKey sets WriteUserSecretAccessKey field to given value.

### HasWriteUserSecretAccessKey

`func (o *IwotenantTenantCustomization) HasWriteUserSecretAccessKey() bool`

HasWriteUserSecretAccessKey returns a boolean if a field has been set.

### GetAccount

`func (o *IwotenantTenantCustomization) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IwotenantTenantCustomization) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IwotenantTenantCustomization) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IwotenantTenantCustomization) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


