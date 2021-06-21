# HyperflexWitnessConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.WitnessConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.WitnessConfiguration"]
**ConnectionError** | Pointer to **string** | The detailed connection error to the external witness. Empty if status is connected. | [optional] 
**CustomWitnessEnabled** | Pointer to **bool** | Custom witness has been configured by user. | [optional] 
**Fingerprint** | Pointer to **string** | The fingerprint of the witness server, identifies the revision of the witness servers database. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty. | [optional] 
**Status** | Pointer to **string** | Status of the devices connection to the witness. Device will report status as either &#39;Connected&#39; or &#39;NotConnected&#39;. | [optional] 
**Version** | Pointer to **string** | The version of the custom witness server. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty. | [optional] 
**WitnessUrl** | Pointer to **string** | URL of the witness endpoint, including IP/host and path. Only applicable if custom witness has been enabled in the cluster, otherwise value is always empty. | [optional] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexWitnessConfiguration

`func NewHyperflexWitnessConfiguration(classId string, objectType string, ) *HyperflexWitnessConfiguration`

NewHyperflexWitnessConfiguration instantiates a new HyperflexWitnessConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexWitnessConfigurationWithDefaults

`func NewHyperflexWitnessConfigurationWithDefaults() *HyperflexWitnessConfiguration`

NewHyperflexWitnessConfigurationWithDefaults instantiates a new HyperflexWitnessConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexWitnessConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexWitnessConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexWitnessConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexWitnessConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexWitnessConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexWitnessConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionError

`func (o *HyperflexWitnessConfiguration) GetConnectionError() string`

GetConnectionError returns the ConnectionError field if non-nil, zero value otherwise.

### GetConnectionErrorOk

`func (o *HyperflexWitnessConfiguration) GetConnectionErrorOk() (*string, bool)`

GetConnectionErrorOk returns a tuple with the ConnectionError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionError

`func (o *HyperflexWitnessConfiguration) SetConnectionError(v string)`

SetConnectionError sets ConnectionError field to given value.

### HasConnectionError

`func (o *HyperflexWitnessConfiguration) HasConnectionError() bool`

HasConnectionError returns a boolean if a field has been set.

### GetCustomWitnessEnabled

`func (o *HyperflexWitnessConfiguration) GetCustomWitnessEnabled() bool`

GetCustomWitnessEnabled returns the CustomWitnessEnabled field if non-nil, zero value otherwise.

### GetCustomWitnessEnabledOk

`func (o *HyperflexWitnessConfiguration) GetCustomWitnessEnabledOk() (*bool, bool)`

GetCustomWitnessEnabledOk returns a tuple with the CustomWitnessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomWitnessEnabled

`func (o *HyperflexWitnessConfiguration) SetCustomWitnessEnabled(v bool)`

SetCustomWitnessEnabled sets CustomWitnessEnabled field to given value.

### HasCustomWitnessEnabled

`func (o *HyperflexWitnessConfiguration) HasCustomWitnessEnabled() bool`

HasCustomWitnessEnabled returns a boolean if a field has been set.

### GetFingerprint

`func (o *HyperflexWitnessConfiguration) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *HyperflexWitnessConfiguration) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *HyperflexWitnessConfiguration) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *HyperflexWitnessConfiguration) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexWitnessConfiguration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexWitnessConfiguration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexWitnessConfiguration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexWitnessConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexWitnessConfiguration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexWitnessConfiguration) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexWitnessConfiguration) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexWitnessConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWitnessUrl

`func (o *HyperflexWitnessConfiguration) GetWitnessUrl() string`

GetWitnessUrl returns the WitnessUrl field if non-nil, zero value otherwise.

### GetWitnessUrlOk

`func (o *HyperflexWitnessConfiguration) GetWitnessUrlOk() (*string, bool)`

GetWitnessUrlOk returns a tuple with the WitnessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWitnessUrl

`func (o *HyperflexWitnessConfiguration) SetWitnessUrl(v string)`

SetWitnessUrl sets WitnessUrl field to given value.

### HasWitnessUrl

`func (o *HyperflexWitnessConfiguration) HasWitnessUrl() bool`

HasWitnessUrl returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexWitnessConfiguration) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexWitnessConfiguration) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexWitnessConfiguration) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexWitnessConfiguration) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


