# AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerMicrosoftAzureApplicationInsightsOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerMicrosoftAzureApplicationInsightsOptions"]
**EnrollmentNumber** | Pointer to **string** | Enrollment number of this Azure account. | [optional] 
**OfferId** | Pointer to **string** | The offer ID of this account. Default value is \&quot;MS-AZR-0003P\&quot;, a pay-as-you-go subscription lets you pay for the services and resources that you use on a monthly basis. | [optional] 
**SubscriptionId** | Pointer to **string** | The Azure Subscription ID. | [optional] 
**TenantId** | Pointer to **string** | Tenant ID associated with Azure Account. | [optional] 

## Methods

### NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions

`func NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions(classId string, objectType string, ) *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions`

NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions instantiates a new AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithDefaults

`func NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions`

NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnrollmentNumber

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetEnrollmentNumber() string`

GetEnrollmentNumber returns the EnrollmentNumber field if non-nil, zero value otherwise.

### GetEnrollmentNumberOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetEnrollmentNumberOk() (*string, bool)`

GetEnrollmentNumberOk returns a tuple with the EnrollmentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentNumber

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetEnrollmentNumber(v string)`

SetEnrollmentNumber sets EnrollmentNumber field to given value.

### HasEnrollmentNumber

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) HasEnrollmentNumber() bool`

HasEnrollmentNumber returns a boolean if a field has been set.

### GetOfferId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


