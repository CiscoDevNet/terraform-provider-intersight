# AssetDeploymentDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeploymentDevice"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeploymentDevice"]
**DeviceId** | Pointer to **string** | Unique identifier of the Cisco device. | [optional] [readonly] 
**DeviceInformation** | Pointer to [**NullableAssetDeploymentDeviceInformation**](asset.DeploymentDeviceInformation.md) |  | [optional] 
**DevicePid** | Pointer to **string** | Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices. | [optional] [readonly] 
**DeviceStatistics** | Pointer to [**NullableAssetDeviceStatistics**](asset.DeviceStatistics.md) |  | [optional] 
**ProductSubgroup** | Pointer to **string** | Product Subgroup type helps to determine if device subgroup within Product type has to be billed using consumption metering. example \&quot;N9300 Series\&quot; in Product type \&quot;SWITCH\&quot;. | [optional] [readonly] 
**ProductType** | Pointer to **string** | Product type helps to determine if device has to be billed using consumption metering. example \&quot;SERVER\&quot;. | [optional] [readonly] 
**UnitOfMeasure** | Pointer to [**[]AssetMeteringType**](AssetMeteringType.md) |  | [optional] 
**VirtualizationPlatform** | Pointer to **string** | Virtualization platform is used to identify the hypervisor type. example \&quot;ESXi\&quot;. | [optional] [readonly] 
**Workload** | Pointer to **string** | Workload/Usecase running on the device. | [optional] [readonly] 
**Deployment** | Pointer to [**AssetDeploymentRelationship**](asset.Deployment.Relationship.md) |  | [optional] 
**DeviceContractInformation** | Pointer to [**AssetDeviceContractInformationRelationship**](asset.DeviceContractInformation.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Subscription** | Pointer to [**AssetSubscriptionRelationship**](asset.Subscription.Relationship.md) |  | [optional] 
**SubscriptionAccount** | Pointer to [**AssetSubscriptionAccountRelationship**](asset.SubscriptionAccount.Relationship.md) |  | [optional] 

## Methods

### NewAssetDeploymentDevice

`func NewAssetDeploymentDevice(classId string, objectType string, ) *AssetDeploymentDevice`

NewAssetDeploymentDevice instantiates a new AssetDeploymentDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeploymentDeviceWithDefaults

`func NewAssetDeploymentDeviceWithDefaults() *AssetDeploymentDevice`

NewAssetDeploymentDeviceWithDefaults instantiates a new AssetDeploymentDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeploymentDevice) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeploymentDevice) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeploymentDevice) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeploymentDevice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeploymentDevice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeploymentDevice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceId

`func (o *AssetDeploymentDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AssetDeploymentDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AssetDeploymentDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AssetDeploymentDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceInformation

`func (o *AssetDeploymentDevice) GetDeviceInformation() AssetDeploymentDeviceInformation`

GetDeviceInformation returns the DeviceInformation field if non-nil, zero value otherwise.

### GetDeviceInformationOk

`func (o *AssetDeploymentDevice) GetDeviceInformationOk() (*AssetDeploymentDeviceInformation, bool)`

GetDeviceInformationOk returns a tuple with the DeviceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInformation

`func (o *AssetDeploymentDevice) SetDeviceInformation(v AssetDeploymentDeviceInformation)`

SetDeviceInformation sets DeviceInformation field to given value.

### HasDeviceInformation

`func (o *AssetDeploymentDevice) HasDeviceInformation() bool`

HasDeviceInformation returns a boolean if a field has been set.

### SetDeviceInformationNil

`func (o *AssetDeploymentDevice) SetDeviceInformationNil(b bool)`

 SetDeviceInformationNil sets the value for DeviceInformation to be an explicit nil

### UnsetDeviceInformation
`func (o *AssetDeploymentDevice) UnsetDeviceInformation()`

UnsetDeviceInformation ensures that no value is present for DeviceInformation, not even an explicit nil
### GetDevicePid

`func (o *AssetDeploymentDevice) GetDevicePid() string`

GetDevicePid returns the DevicePid field if non-nil, zero value otherwise.

### GetDevicePidOk

`func (o *AssetDeploymentDevice) GetDevicePidOk() (*string, bool)`

GetDevicePidOk returns a tuple with the DevicePid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePid

`func (o *AssetDeploymentDevice) SetDevicePid(v string)`

SetDevicePid sets DevicePid field to given value.

### HasDevicePid

`func (o *AssetDeploymentDevice) HasDevicePid() bool`

HasDevicePid returns a boolean if a field has been set.

### GetDeviceStatistics

`func (o *AssetDeploymentDevice) GetDeviceStatistics() AssetDeviceStatistics`

GetDeviceStatistics returns the DeviceStatistics field if non-nil, zero value otherwise.

### GetDeviceStatisticsOk

`func (o *AssetDeploymentDevice) GetDeviceStatisticsOk() (*AssetDeviceStatistics, bool)`

GetDeviceStatisticsOk returns a tuple with the DeviceStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatistics

`func (o *AssetDeploymentDevice) SetDeviceStatistics(v AssetDeviceStatistics)`

SetDeviceStatistics sets DeviceStatistics field to given value.

### HasDeviceStatistics

`func (o *AssetDeploymentDevice) HasDeviceStatistics() bool`

HasDeviceStatistics returns a boolean if a field has been set.

### SetDeviceStatisticsNil

`func (o *AssetDeploymentDevice) SetDeviceStatisticsNil(b bool)`

 SetDeviceStatisticsNil sets the value for DeviceStatistics to be an explicit nil

### UnsetDeviceStatistics
`func (o *AssetDeploymentDevice) UnsetDeviceStatistics()`

UnsetDeviceStatistics ensures that no value is present for DeviceStatistics, not even an explicit nil
### GetProductSubgroup

`func (o *AssetDeploymentDevice) GetProductSubgroup() string`

GetProductSubgroup returns the ProductSubgroup field if non-nil, zero value otherwise.

### GetProductSubgroupOk

`func (o *AssetDeploymentDevice) GetProductSubgroupOk() (*string, bool)`

GetProductSubgroupOk returns a tuple with the ProductSubgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubgroup

`func (o *AssetDeploymentDevice) SetProductSubgroup(v string)`

SetProductSubgroup sets ProductSubgroup field to given value.

### HasProductSubgroup

`func (o *AssetDeploymentDevice) HasProductSubgroup() bool`

HasProductSubgroup returns a boolean if a field has been set.

### GetProductType

`func (o *AssetDeploymentDevice) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AssetDeploymentDevice) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AssetDeploymentDevice) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AssetDeploymentDevice) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *AssetDeploymentDevice) GetUnitOfMeasure() []AssetMeteringType`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *AssetDeploymentDevice) GetUnitOfMeasureOk() (*[]AssetMeteringType, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *AssetDeploymentDevice) SetUnitOfMeasure(v []AssetMeteringType)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *AssetDeploymentDevice) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### SetUnitOfMeasureNil

`func (o *AssetDeploymentDevice) SetUnitOfMeasureNil(b bool)`

 SetUnitOfMeasureNil sets the value for UnitOfMeasure to be an explicit nil

### UnsetUnitOfMeasure
`func (o *AssetDeploymentDevice) UnsetUnitOfMeasure()`

UnsetUnitOfMeasure ensures that no value is present for UnitOfMeasure, not even an explicit nil
### GetVirtualizationPlatform

`func (o *AssetDeploymentDevice) GetVirtualizationPlatform() string`

GetVirtualizationPlatform returns the VirtualizationPlatform field if non-nil, zero value otherwise.

### GetVirtualizationPlatformOk

`func (o *AssetDeploymentDevice) GetVirtualizationPlatformOk() (*string, bool)`

GetVirtualizationPlatformOk returns a tuple with the VirtualizationPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationPlatform

`func (o *AssetDeploymentDevice) SetVirtualizationPlatform(v string)`

SetVirtualizationPlatform sets VirtualizationPlatform field to given value.

### HasVirtualizationPlatform

`func (o *AssetDeploymentDevice) HasVirtualizationPlatform() bool`

HasVirtualizationPlatform returns a boolean if a field has been set.

### GetWorkload

`func (o *AssetDeploymentDevice) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *AssetDeploymentDevice) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *AssetDeploymentDevice) SetWorkload(v string)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *AssetDeploymentDevice) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### GetDeployment

`func (o *AssetDeploymentDevice) GetDeployment() AssetDeploymentRelationship`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *AssetDeploymentDevice) GetDeploymentOk() (*AssetDeploymentRelationship, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *AssetDeploymentDevice) SetDeployment(v AssetDeploymentRelationship)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *AssetDeploymentDevice) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetDeviceContractInformation

`func (o *AssetDeploymentDevice) GetDeviceContractInformation() AssetDeviceContractInformationRelationship`

GetDeviceContractInformation returns the DeviceContractInformation field if non-nil, zero value otherwise.

### GetDeviceContractInformationOk

`func (o *AssetDeploymentDevice) GetDeviceContractInformationOk() (*AssetDeviceContractInformationRelationship, bool)`

GetDeviceContractInformationOk returns a tuple with the DeviceContractInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceContractInformation

`func (o *AssetDeploymentDevice) SetDeviceContractInformation(v AssetDeviceContractInformationRelationship)`

SetDeviceContractInformation sets DeviceContractInformation field to given value.

### HasDeviceContractInformation

`func (o *AssetDeploymentDevice) HasDeviceContractInformation() bool`

HasDeviceContractInformation returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AssetDeploymentDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AssetDeploymentDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AssetDeploymentDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AssetDeploymentDevice) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSubscription

`func (o *AssetDeploymentDevice) GetSubscription() AssetSubscriptionRelationship`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AssetDeploymentDevice) GetSubscriptionOk() (*AssetSubscriptionRelationship, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AssetDeploymentDevice) SetSubscription(v AssetSubscriptionRelationship)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *AssetDeploymentDevice) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSubscriptionAccount

`func (o *AssetDeploymentDevice) GetSubscriptionAccount() AssetSubscriptionAccountRelationship`

GetSubscriptionAccount returns the SubscriptionAccount field if non-nil, zero value otherwise.

### GetSubscriptionAccountOk

`func (o *AssetDeploymentDevice) GetSubscriptionAccountOk() (*AssetSubscriptionAccountRelationship, bool)`

GetSubscriptionAccountOk returns a tuple with the SubscriptionAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAccount

`func (o *AssetDeploymentDevice) SetSubscriptionAccount(v AssetSubscriptionAccountRelationship)`

SetSubscriptionAccount sets SubscriptionAccount field to given value.

### HasSubscriptionAccount

`func (o *AssetDeploymentDevice) HasSubscriptionAccount() bool`

HasSubscriptionAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


