# HciViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Violation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Violation"]
**CapacityViolations** | Pointer to [**[]HciCapacityViolation**](HciCapacityViolation.md) |  | [optional] 
**ClusterExtId** | Pointer to **string** | The unique identifier of the cluster that has the violations. | [optional] [readonly] 
**ExpiredLicenses** | Pointer to [**[]HciExpiredLicense**](HciExpiredLicense.md) |  | [optional] 
**FeatureViolations** | Pointer to [**[]HciFeatureViolation**](HciFeatureViolation.md) |  | [optional] 
**IsMulticluster** | Pointer to **bool** | Indicates if the violation is on multi-cluster which is PrismCentral in Nutanix&#39;s term. Otherwise, it is on a Prism Element cluster. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciViolation

`func NewHciViolation(classId string, objectType string, ) *HciViolation`

NewHciViolation instantiates a new HciViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciViolationWithDefaults

`func NewHciViolationWithDefaults() *HciViolation`

NewHciViolationWithDefaults instantiates a new HciViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciViolation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciViolation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciViolation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciViolation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciViolation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciViolation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacityViolations

`func (o *HciViolation) GetCapacityViolations() []HciCapacityViolation`

GetCapacityViolations returns the CapacityViolations field if non-nil, zero value otherwise.

### GetCapacityViolationsOk

`func (o *HciViolation) GetCapacityViolationsOk() (*[]HciCapacityViolation, bool)`

GetCapacityViolationsOk returns a tuple with the CapacityViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityViolations

`func (o *HciViolation) SetCapacityViolations(v []HciCapacityViolation)`

SetCapacityViolations sets CapacityViolations field to given value.

### HasCapacityViolations

`func (o *HciViolation) HasCapacityViolations() bool`

HasCapacityViolations returns a boolean if a field has been set.

### SetCapacityViolationsNil

`func (o *HciViolation) SetCapacityViolationsNil(b bool)`

 SetCapacityViolationsNil sets the value for CapacityViolations to be an explicit nil

### UnsetCapacityViolations
`func (o *HciViolation) UnsetCapacityViolations()`

UnsetCapacityViolations ensures that no value is present for CapacityViolations, not even an explicit nil
### GetClusterExtId

`func (o *HciViolation) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciViolation) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciViolation) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciViolation) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetExpiredLicenses

`func (o *HciViolation) GetExpiredLicenses() []HciExpiredLicense`

GetExpiredLicenses returns the ExpiredLicenses field if non-nil, zero value otherwise.

### GetExpiredLicensesOk

`func (o *HciViolation) GetExpiredLicensesOk() (*[]HciExpiredLicense, bool)`

GetExpiredLicensesOk returns a tuple with the ExpiredLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredLicenses

`func (o *HciViolation) SetExpiredLicenses(v []HciExpiredLicense)`

SetExpiredLicenses sets ExpiredLicenses field to given value.

### HasExpiredLicenses

`func (o *HciViolation) HasExpiredLicenses() bool`

HasExpiredLicenses returns a boolean if a field has been set.

### SetExpiredLicensesNil

`func (o *HciViolation) SetExpiredLicensesNil(b bool)`

 SetExpiredLicensesNil sets the value for ExpiredLicenses to be an explicit nil

### UnsetExpiredLicenses
`func (o *HciViolation) UnsetExpiredLicenses()`

UnsetExpiredLicenses ensures that no value is present for ExpiredLicenses, not even an explicit nil
### GetFeatureViolations

`func (o *HciViolation) GetFeatureViolations() []HciFeatureViolation`

GetFeatureViolations returns the FeatureViolations field if non-nil, zero value otherwise.

### GetFeatureViolationsOk

`func (o *HciViolation) GetFeatureViolationsOk() (*[]HciFeatureViolation, bool)`

GetFeatureViolationsOk returns a tuple with the FeatureViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureViolations

`func (o *HciViolation) SetFeatureViolations(v []HciFeatureViolation)`

SetFeatureViolations sets FeatureViolations field to given value.

### HasFeatureViolations

`func (o *HciViolation) HasFeatureViolations() bool`

HasFeatureViolations returns a boolean if a field has been set.

### SetFeatureViolationsNil

`func (o *HciViolation) SetFeatureViolationsNil(b bool)`

 SetFeatureViolationsNil sets the value for FeatureViolations to be an explicit nil

### UnsetFeatureViolations
`func (o *HciViolation) UnsetFeatureViolations()`

UnsetFeatureViolations ensures that no value is present for FeatureViolations, not even an explicit nil
### GetIsMulticluster

`func (o *HciViolation) GetIsMulticluster() bool`

GetIsMulticluster returns the IsMulticluster field if non-nil, zero value otherwise.

### GetIsMulticlusterOk

`func (o *HciViolation) GetIsMulticlusterOk() (*bool, bool)`

GetIsMulticlusterOk returns a tuple with the IsMulticluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMulticluster

`func (o *HciViolation) SetIsMulticluster(v bool)`

SetIsMulticluster sets IsMulticluster field to given value.

### HasIsMulticluster

`func (o *HciViolation) HasIsMulticluster() bool`

HasIsMulticluster returns a boolean if a field has been set.

### GetCluster

`func (o *HciViolation) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciViolation) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciViolation) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciViolation) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciViolation) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciViolation) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetRegisteredDevice

`func (o *HciViolation) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciViolation) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciViolation) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciViolation) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciViolation) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciViolation) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


