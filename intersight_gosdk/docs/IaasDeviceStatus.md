# IaasDeviceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.DeviceStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.DeviceStatus"]
**AccountName** | Pointer to **string** | The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD. | [optional] [readonly] 
**AccountType** | Pointer to **string** | The UCSD Infra Account type. | [optional] [readonly] 
**Category** | Pointer to **string** | The UCSD Infra Account category. | [optional] [readonly] 
**ClaimStatus** | Pointer to **string** | Describes if the device is claimed in Intersight or not. * &#x60;Unknown&#x60; - If UCS Director managed account claim status information is unknown. * &#x60;Yes&#x60; - If UCS Director managed account is claimed in Intersight. * &#x60;No&#x60; - If UCS Director managed account is not claimed in Intersight. * &#x60;Not Applicable&#x60; - If UCS Director managed account is not capable of providing claim status information. | [optional] [readonly] [default to "Unknown"]
**ConnectionStatus** | Pointer to **string** | Describes about the connection status between the UCSD and the actual end device. | [optional] [readonly] 
**DeviceModel** | Pointer to **string** | Describes about the device model. | [optional] [readonly] 
**DeviceVendor** | Pointer to **string** | Describes about the device vendor/manufacturer of the device. | [optional] [readonly] 
**DeviceVersion** | Pointer to **string** | Describes about the current firmware version running on the device. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | The IPAddress of the device. | [optional] [readonly] 
**Pod** | Pointer to **string** | Describes about the pod to which this device belongs to in UCSD. | [optional] [readonly] 
**PodType** | Pointer to **string** | Describes about the podType of Pod to which this device belongs to in UCSD. | [optional] [readonly] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](IaasUcsdInfoRelationship.md) |  | [optional] 

## Methods

### NewIaasDeviceStatus

`func NewIaasDeviceStatus(classId string, objectType string, ) *IaasDeviceStatus`

NewIaasDeviceStatus instantiates a new IaasDeviceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasDeviceStatusWithDefaults

`func NewIaasDeviceStatusWithDefaults() *IaasDeviceStatus`

NewIaasDeviceStatusWithDefaults instantiates a new IaasDeviceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasDeviceStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasDeviceStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasDeviceStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasDeviceStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasDeviceStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasDeviceStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountName

`func (o *IaasDeviceStatus) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *IaasDeviceStatus) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *IaasDeviceStatus) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *IaasDeviceStatus) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountType

`func (o *IaasDeviceStatus) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *IaasDeviceStatus) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *IaasDeviceStatus) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *IaasDeviceStatus) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetCategory

`func (o *IaasDeviceStatus) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IaasDeviceStatus) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IaasDeviceStatus) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IaasDeviceStatus) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClaimStatus

`func (o *IaasDeviceStatus) GetClaimStatus() string`

GetClaimStatus returns the ClaimStatus field if non-nil, zero value otherwise.

### GetClaimStatusOk

`func (o *IaasDeviceStatus) GetClaimStatusOk() (*string, bool)`

GetClaimStatusOk returns a tuple with the ClaimStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimStatus

`func (o *IaasDeviceStatus) SetClaimStatus(v string)`

SetClaimStatus sets ClaimStatus field to given value.

### HasClaimStatus

`func (o *IaasDeviceStatus) HasClaimStatus() bool`

HasClaimStatus returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *IaasDeviceStatus) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *IaasDeviceStatus) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *IaasDeviceStatus) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *IaasDeviceStatus) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDeviceModel

`func (o *IaasDeviceStatus) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *IaasDeviceStatus) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *IaasDeviceStatus) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *IaasDeviceStatus) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *IaasDeviceStatus) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *IaasDeviceStatus) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *IaasDeviceStatus) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *IaasDeviceStatus) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDeviceVersion

`func (o *IaasDeviceStatus) GetDeviceVersion() string`

GetDeviceVersion returns the DeviceVersion field if non-nil, zero value otherwise.

### GetDeviceVersionOk

`func (o *IaasDeviceStatus) GetDeviceVersionOk() (*string, bool)`

GetDeviceVersionOk returns a tuple with the DeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVersion

`func (o *IaasDeviceStatus) SetDeviceVersion(v string)`

SetDeviceVersion sets DeviceVersion field to given value.

### HasDeviceVersion

`func (o *IaasDeviceStatus) HasDeviceVersion() bool`

HasDeviceVersion returns a boolean if a field has been set.

### GetIpAddress

`func (o *IaasDeviceStatus) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IaasDeviceStatus) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IaasDeviceStatus) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *IaasDeviceStatus) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPod

`func (o *IaasDeviceStatus) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *IaasDeviceStatus) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *IaasDeviceStatus) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *IaasDeviceStatus) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetPodType

`func (o *IaasDeviceStatus) GetPodType() string`

GetPodType returns the PodType field if non-nil, zero value otherwise.

### GetPodTypeOk

`func (o *IaasDeviceStatus) GetPodTypeOk() (*string, bool)`

GetPodTypeOk returns a tuple with the PodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodType

`func (o *IaasDeviceStatus) SetPodType(v string)`

SetPodType sets PodType field to given value.

### HasPodType

`func (o *IaasDeviceStatus) HasPodType() bool`

HasPodType returns a boolean if a field has been set.

### GetGuid

`func (o *IaasDeviceStatus) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasDeviceStatus) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasDeviceStatus) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasDeviceStatus) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


