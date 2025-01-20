# HciCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Compliance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Compliance"]
**ClusterExtId** | Pointer to **string** | The identifier of a license, usually in LIC-xxxx format. | [optional] [readonly] 
**ComplianceState** | Pointer to **string** | The compliance state of the cluster. The compliance state is determined based on the nonComplianceCount and the license enforcement policy. * &#x60;NotEnforced&#x60; - The license compliance state for a Nutanix cluster is neither in&#x3D;compliance nor not in-compliance.Typically, when a license is not installed on a cluster, or the license enformancement is explicitlyturned off, the cluster is in this state. * &#x60;InCompliance&#x60; - The license compliance state for a Nutanix cluster is in compliance. * &#x60;NotInCompliance&#x60; - The license compliance state for a Nutanix cluster is not in compliance. | [optional] [readonly] [default to "NotEnforced"]
**NonComplianceCount** | Pointer to **int32** | Total number of services that are not in-compliance. A count of 0 does not necessarily mean in compliance. Depending on &#39;complianceState&#39;, the cluster could be either be in compliance or the license check is not enforced. A synthesized property for the services property for ease of querying. | [optional] [readonly] 
**Services** | Pointer to [**[]HciLicensedService**](HciLicensedService.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the cluster which could be NUTANIX, VMWARE, or NON_NUTANIX. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciCompliance

`func NewHciCompliance(classId string, objectType string, ) *HciCompliance`

NewHciCompliance instantiates a new HciCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciComplianceWithDefaults

`func NewHciComplianceWithDefaults() *HciCompliance`

NewHciComplianceWithDefaults instantiates a new HciCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciCompliance) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciCompliance) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciCompliance) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciCompliance) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciCompliance) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciCompliance) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterExtId

`func (o *HciCompliance) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciCompliance) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciCompliance) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciCompliance) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetComplianceState

`func (o *HciCompliance) GetComplianceState() string`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *HciCompliance) GetComplianceStateOk() (*string, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *HciCompliance) SetComplianceState(v string)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *HciCompliance) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### GetNonComplianceCount

`func (o *HciCompliance) GetNonComplianceCount() int32`

GetNonComplianceCount returns the NonComplianceCount field if non-nil, zero value otherwise.

### GetNonComplianceCountOk

`func (o *HciCompliance) GetNonComplianceCountOk() (*int32, bool)`

GetNonComplianceCountOk returns a tuple with the NonComplianceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonComplianceCount

`func (o *HciCompliance) SetNonComplianceCount(v int32)`

SetNonComplianceCount sets NonComplianceCount field to given value.

### HasNonComplianceCount

`func (o *HciCompliance) HasNonComplianceCount() bool`

HasNonComplianceCount returns a boolean if a field has been set.

### GetServices

`func (o *HciCompliance) GetServices() []HciLicensedService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *HciCompliance) GetServicesOk() (*[]HciLicensedService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *HciCompliance) SetServices(v []HciLicensedService)`

SetServices sets Services field to given value.

### HasServices

`func (o *HciCompliance) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *HciCompliance) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *HciCompliance) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetType

`func (o *HciCompliance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciCompliance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciCompliance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciCompliance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCluster

`func (o *HciCompliance) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciCompliance) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciCompliance) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciCompliance) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciCompliance) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciCompliance) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetRegisteredDevice

`func (o *HciCompliance) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciCompliance) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciCompliance) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciCompliance) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciCompliance) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciCompliance) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


