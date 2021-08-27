# KubernetesAciCniApicAllOf

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

### NewKubernetesAciCniApicAllOf

`func NewKubernetesAciCniApicAllOf(classId string, objectType string, ) *KubernetesAciCniApicAllOf`

NewKubernetesAciCniApicAllOf instantiates a new KubernetesAciCniApicAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAciCniApicAllOfWithDefaults

`func NewKubernetesAciCniApicAllOfWithDefaults() *KubernetesAciCniApicAllOf`

NewKubernetesAciCniApicAllOfWithDefaults instantiates a new KubernetesAciCniApicAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAciCniApicAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAciCniApicAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAciCniApicAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAciCniApicAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAciCniApicAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAciCniApicAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssetApicMoid

`func (o *KubernetesAciCniApicAllOf) GetAssetApicMoid() string`

GetAssetApicMoid returns the AssetApicMoid field if non-nil, zero value otherwise.

### GetAssetApicMoidOk

`func (o *KubernetesAciCniApicAllOf) GetAssetApicMoidOk() (*string, bool)`

GetAssetApicMoidOk returns a tuple with the AssetApicMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetApicMoid

`func (o *KubernetesAciCniApicAllOf) SetAssetApicMoid(v string)`

SetAssetApicMoid sets AssetApicMoid field to given value.

### HasAssetApicMoid

`func (o *KubernetesAciCniApicAllOf) HasAssetApicMoid() bool`

HasAssetApicMoid returns a boolean if a field has been set.

### GetNumAciCniProfiles

`func (o *KubernetesAciCniApicAllOf) GetNumAciCniProfiles() int64`

GetNumAciCniProfiles returns the NumAciCniProfiles field if non-nil, zero value otherwise.

### GetNumAciCniProfilesOk

`func (o *KubernetesAciCniApicAllOf) GetNumAciCniProfilesOk() (*int64, bool)`

GetNumAciCniProfilesOk returns a tuple with the NumAciCniProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAciCniProfiles

`func (o *KubernetesAciCniApicAllOf) SetNumAciCniProfiles(v int64)`

SetNumAciCniProfiles sets NumAciCniProfiles field to given value.

### HasNumAciCniProfiles

`func (o *KubernetesAciCniApicAllOf) HasNumAciCniProfiles() bool`

HasNumAciCniProfiles returns a boolean if a field has been set.

### GetAciCniProfiles

`func (o *KubernetesAciCniApicAllOf) GetAciCniProfiles() []KubernetesAciCniProfileRelationship`

GetAciCniProfiles returns the AciCniProfiles field if non-nil, zero value otherwise.

### GetAciCniProfilesOk

`func (o *KubernetesAciCniApicAllOf) GetAciCniProfilesOk() (*[]KubernetesAciCniProfileRelationship, bool)`

GetAciCniProfilesOk returns a tuple with the AciCniProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAciCniProfiles

`func (o *KubernetesAciCniApicAllOf) SetAciCniProfiles(v []KubernetesAciCniProfileRelationship)`

SetAciCniProfiles sets AciCniProfiles field to given value.

### HasAciCniProfiles

`func (o *KubernetesAciCniApicAllOf) HasAciCniProfiles() bool`

HasAciCniProfiles returns a boolean if a field has been set.

### SetAciCniProfilesNil

`func (o *KubernetesAciCniApicAllOf) SetAciCniProfilesNil(b bool)`

 SetAciCniProfilesNil sets the value for AciCniProfiles to be an explicit nil

### UnsetAciCniProfiles
`func (o *KubernetesAciCniApicAllOf) UnsetAciCniProfiles()`

UnsetAciCniProfiles ensures that no value is present for AciCniProfiles, not even an explicit nil
### GetOrganization

`func (o *KubernetesAciCniApicAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAciCniApicAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAciCniApicAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAciCniApicAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *KubernetesAciCniApicAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *KubernetesAciCniApicAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *KubernetesAciCniApicAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *KubernetesAciCniApicAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


