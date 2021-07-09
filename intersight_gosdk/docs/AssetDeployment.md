# AssetDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.Deployment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.Deployment"]
**DeploymentRefId** | Pointer to **string** | Identifies the consumption-based subscription&#39;s deployment. | [optional] [readonly] 
**EndCustomer** | Pointer to [**NullableAssetCustomerInformation**](asset.CustomerInformation.md) |  | [optional] 
**EndDate** | Pointer to **time.Time** | End Date for the consumption-based subscription&#39;s deployment. | [optional] [readonly] 
**LicenseType** | Pointer to **string** | Active license tier for the purchased Cisco device&#39;s deployment. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. | [optional] [readonly] [default to "Base"]
**MlbOfferType** | Pointer to **string** | Identifier for consumption based pricing. MLB refers to MultiLine Bundle. | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | Start Date for the consumption-based subscription&#39;s deployment. | [optional] [readonly] 
**SubscriptionRefId** | Pointer to **string** | Identifies the consumption-based subscription. | [optional] [readonly] 
**UnitOfMeasure** | Pointer to [**[]AssetMeteringType**](AssetMeteringType.md) |  | [optional] 
**Workloads** | Pointer to **[]string** |  | [optional] 
**Devices** | Pointer to [**[]AssetDeploymentDeviceRelationship**](AssetDeploymentDeviceRelationship.md) | An array of relationships to assetDeploymentDevice resources. | [optional] [readonly] 
**Subscription** | Pointer to [**AssetSubscriptionRelationship**](asset.Subscription.Relationship.md) |  | [optional] 

## Methods

### NewAssetDeployment

`func NewAssetDeployment(classId string, objectType string, ) *AssetDeployment`

NewAssetDeployment instantiates a new AssetDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeploymentWithDefaults

`func NewAssetDeploymentWithDefaults() *AssetDeployment`

NewAssetDeploymentWithDefaults instantiates a new AssetDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeployment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeployment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeployment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeployment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeployment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeployment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeploymentRefId

`func (o *AssetDeployment) GetDeploymentRefId() string`

GetDeploymentRefId returns the DeploymentRefId field if non-nil, zero value otherwise.

### GetDeploymentRefIdOk

`func (o *AssetDeployment) GetDeploymentRefIdOk() (*string, bool)`

GetDeploymentRefIdOk returns a tuple with the DeploymentRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRefId

`func (o *AssetDeployment) SetDeploymentRefId(v string)`

SetDeploymentRefId sets DeploymentRefId field to given value.

### HasDeploymentRefId

`func (o *AssetDeployment) HasDeploymentRefId() bool`

HasDeploymentRefId returns a boolean if a field has been set.

### GetEndCustomer

`func (o *AssetDeployment) GetEndCustomer() AssetCustomerInformation`

GetEndCustomer returns the EndCustomer field if non-nil, zero value otherwise.

### GetEndCustomerOk

`func (o *AssetDeployment) GetEndCustomerOk() (*AssetCustomerInformation, bool)`

GetEndCustomerOk returns a tuple with the EndCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomer

`func (o *AssetDeployment) SetEndCustomer(v AssetCustomerInformation)`

SetEndCustomer sets EndCustomer field to given value.

### HasEndCustomer

`func (o *AssetDeployment) HasEndCustomer() bool`

HasEndCustomer returns a boolean if a field has been set.

### SetEndCustomerNil

`func (o *AssetDeployment) SetEndCustomerNil(b bool)`

 SetEndCustomerNil sets the value for EndCustomer to be an explicit nil

### UnsetEndCustomer
`func (o *AssetDeployment) UnsetEndCustomer()`

UnsetEndCustomer ensures that no value is present for EndCustomer, not even an explicit nil
### GetEndDate

`func (o *AssetDeployment) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AssetDeployment) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AssetDeployment) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AssetDeployment) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLicenseType

`func (o *AssetDeployment) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *AssetDeployment) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *AssetDeployment) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *AssetDeployment) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetMlbOfferType

`func (o *AssetDeployment) GetMlbOfferType() string`

GetMlbOfferType returns the MlbOfferType field if non-nil, zero value otherwise.

### GetMlbOfferTypeOk

`func (o *AssetDeployment) GetMlbOfferTypeOk() (*string, bool)`

GetMlbOfferTypeOk returns a tuple with the MlbOfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlbOfferType

`func (o *AssetDeployment) SetMlbOfferType(v string)`

SetMlbOfferType sets MlbOfferType field to given value.

### HasMlbOfferType

`func (o *AssetDeployment) HasMlbOfferType() bool`

HasMlbOfferType returns a boolean if a field has been set.

### GetStartDate

`func (o *AssetDeployment) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AssetDeployment) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AssetDeployment) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AssetDeployment) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionRefId

`func (o *AssetDeployment) GetSubscriptionRefId() string`

GetSubscriptionRefId returns the SubscriptionRefId field if non-nil, zero value otherwise.

### GetSubscriptionRefIdOk

`func (o *AssetDeployment) GetSubscriptionRefIdOk() (*string, bool)`

GetSubscriptionRefIdOk returns a tuple with the SubscriptionRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRefId

`func (o *AssetDeployment) SetSubscriptionRefId(v string)`

SetSubscriptionRefId sets SubscriptionRefId field to given value.

### HasSubscriptionRefId

`func (o *AssetDeployment) HasSubscriptionRefId() bool`

HasSubscriptionRefId returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *AssetDeployment) GetUnitOfMeasure() []AssetMeteringType`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *AssetDeployment) GetUnitOfMeasureOk() (*[]AssetMeteringType, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *AssetDeployment) SetUnitOfMeasure(v []AssetMeteringType)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *AssetDeployment) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### SetUnitOfMeasureNil

`func (o *AssetDeployment) SetUnitOfMeasureNil(b bool)`

 SetUnitOfMeasureNil sets the value for UnitOfMeasure to be an explicit nil

### UnsetUnitOfMeasure
`func (o *AssetDeployment) UnsetUnitOfMeasure()`

UnsetUnitOfMeasure ensures that no value is present for UnitOfMeasure, not even an explicit nil
### GetWorkloads

`func (o *AssetDeployment) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *AssetDeployment) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *AssetDeployment) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *AssetDeployment) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### SetWorkloadsNil

`func (o *AssetDeployment) SetWorkloadsNil(b bool)`

 SetWorkloadsNil sets the value for Workloads to be an explicit nil

### UnsetWorkloads
`func (o *AssetDeployment) UnsetWorkloads()`

UnsetWorkloads ensures that no value is present for Workloads, not even an explicit nil
### GetDevices

`func (o *AssetDeployment) GetDevices() []AssetDeploymentDeviceRelationship`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *AssetDeployment) GetDevicesOk() (*[]AssetDeploymentDeviceRelationship, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *AssetDeployment) SetDevices(v []AssetDeploymentDeviceRelationship)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *AssetDeployment) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### SetDevicesNil

`func (o *AssetDeployment) SetDevicesNil(b bool)`

 SetDevicesNil sets the value for Devices to be an explicit nil

### UnsetDevices
`func (o *AssetDeployment) UnsetDevices()`

UnsetDevices ensures that no value is present for Devices, not even an explicit nil
### GetSubscription

`func (o *AssetDeployment) GetSubscription() AssetSubscriptionRelationship`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AssetDeployment) GetSubscriptionOk() (*AssetSubscriptionRelationship, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AssetDeployment) SetSubscription(v AssetSubscriptionRelationship)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *AssetDeployment) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


