# KubernetesAciCniApic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AciCniApic"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AciCniApic"]
**AssetApicMoid** | Pointer to **string** | The Moid of the apic device under the asset service. | [optional] 
**NumAciCniProfiles** | Pointer to **int64** | The number of ACI CNI profiles configured for this APIC. | [optional] 
**AciCniProfiles** | Pointer to [**[]KubernetesAciCniProfileRelationship**](KubernetesAciCniProfileRelationship.md) | An array of relationships to kubernetesAciCniProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesAciCniApic

`func NewKubernetesAciCniApic(classId string, objectType string, ) *KubernetesAciCniApic`

NewKubernetesAciCniApic instantiates a new KubernetesAciCniApic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAciCniApicWithDefaults

`func NewKubernetesAciCniApicWithDefaults() *KubernetesAciCniApic`

NewKubernetesAciCniApicWithDefaults instantiates a new KubernetesAciCniApic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAciCniApic) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAciCniApic) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAciCniApic) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAciCniApic) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAciCniApic) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAciCniApic) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssetApicMoid

`func (o *KubernetesAciCniApic) GetAssetApicMoid() string`

GetAssetApicMoid returns the AssetApicMoid field if non-nil, zero value otherwise.

### GetAssetApicMoidOk

`func (o *KubernetesAciCniApic) GetAssetApicMoidOk() (*string, bool)`

GetAssetApicMoidOk returns a tuple with the AssetApicMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetApicMoid

`func (o *KubernetesAciCniApic) SetAssetApicMoid(v string)`

SetAssetApicMoid sets AssetApicMoid field to given value.

### HasAssetApicMoid

`func (o *KubernetesAciCniApic) HasAssetApicMoid() bool`

HasAssetApicMoid returns a boolean if a field has been set.

### GetNumAciCniProfiles

`func (o *KubernetesAciCniApic) GetNumAciCniProfiles() int64`

GetNumAciCniProfiles returns the NumAciCniProfiles field if non-nil, zero value otherwise.

### GetNumAciCniProfilesOk

`func (o *KubernetesAciCniApic) GetNumAciCniProfilesOk() (*int64, bool)`

GetNumAciCniProfilesOk returns a tuple with the NumAciCniProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAciCniProfiles

`func (o *KubernetesAciCniApic) SetNumAciCniProfiles(v int64)`

SetNumAciCniProfiles sets NumAciCniProfiles field to given value.

### HasNumAciCniProfiles

`func (o *KubernetesAciCniApic) HasNumAciCniProfiles() bool`

HasNumAciCniProfiles returns a boolean if a field has been set.

### GetAciCniProfiles

`func (o *KubernetesAciCniApic) GetAciCniProfiles() []KubernetesAciCniProfileRelationship`

GetAciCniProfiles returns the AciCniProfiles field if non-nil, zero value otherwise.

### GetAciCniProfilesOk

`func (o *KubernetesAciCniApic) GetAciCniProfilesOk() (*[]KubernetesAciCniProfileRelationship, bool)`

GetAciCniProfilesOk returns a tuple with the AciCniProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAciCniProfiles

`func (o *KubernetesAciCniApic) SetAciCniProfiles(v []KubernetesAciCniProfileRelationship)`

SetAciCniProfiles sets AciCniProfiles field to given value.

### HasAciCniProfiles

`func (o *KubernetesAciCniApic) HasAciCniProfiles() bool`

HasAciCniProfiles returns a boolean if a field has been set.

### SetAciCniProfilesNil

`func (o *KubernetesAciCniApic) SetAciCniProfilesNil(b bool)`

 SetAciCniProfilesNil sets the value for AciCniProfiles to be an explicit nil

### UnsetAciCniProfiles
`func (o *KubernetesAciCniApic) UnsetAciCniProfiles()`

UnsetAciCniProfiles ensures that no value is present for AciCniProfiles, not even an explicit nil
### GetOrganization

`func (o *KubernetesAciCniApic) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAciCniApic) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAciCniApic) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAciCniApic) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *KubernetesAciCniApic) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *KubernetesAciCniApic) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *KubernetesAciCniApic) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *KubernetesAciCniApic) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


