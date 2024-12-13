# HciDomainManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.DomainManager"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.DomainManager"]
**Name** | Pointer to **string** | The name of the domain manager. | [optional] [readonly] 
**PcExtId** | Pointer to **string** | The unique identifier of the domain manager (Prism Central) instance. | [optional] [readonly] 
**Size** | Pointer to **string** | The size of the domain manager such as STARTER, SMALL, LARGE, EXTRALARGE. It determines the resources used by the domain manager. | [optional] [readonly] 
**Clusters** | Pointer to [**[]HciClusterRelationship**](HciClusterRelationship.md) | An array of relationships to hciCluster resources. | [optional] [readonly] 
**Licenses** | Pointer to [**[]HciLicenseRelationship**](HciLicenseRelationship.md) | An array of relationships to hciLicense resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciDomainManager

`func NewHciDomainManager(classId string, objectType string, ) *HciDomainManager`

NewHciDomainManager instantiates a new HciDomainManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciDomainManagerWithDefaults

`func NewHciDomainManagerWithDefaults() *HciDomainManager`

NewHciDomainManagerWithDefaults instantiates a new HciDomainManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciDomainManager) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciDomainManager) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciDomainManager) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciDomainManager) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciDomainManager) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciDomainManager) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *HciDomainManager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciDomainManager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciDomainManager) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciDomainManager) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcExtId

`func (o *HciDomainManager) GetPcExtId() string`

GetPcExtId returns the PcExtId field if non-nil, zero value otherwise.

### GetPcExtIdOk

`func (o *HciDomainManager) GetPcExtIdOk() (*string, bool)`

GetPcExtIdOk returns a tuple with the PcExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcExtId

`func (o *HciDomainManager) SetPcExtId(v string)`

SetPcExtId sets PcExtId field to given value.

### HasPcExtId

`func (o *HciDomainManager) HasPcExtId() bool`

HasPcExtId returns a boolean if a field has been set.

### GetSize

`func (o *HciDomainManager) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HciDomainManager) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HciDomainManager) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *HciDomainManager) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetClusters

`func (o *HciDomainManager) GetClusters() []HciClusterRelationship`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *HciDomainManager) GetClustersOk() (*[]HciClusterRelationship, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *HciDomainManager) SetClusters(v []HciClusterRelationship)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *HciDomainManager) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### SetClustersNil

`func (o *HciDomainManager) SetClustersNil(b bool)`

 SetClustersNil sets the value for Clusters to be an explicit nil

### UnsetClusters
`func (o *HciDomainManager) UnsetClusters()`

UnsetClusters ensures that no value is present for Clusters, not even an explicit nil
### GetLicenses

`func (o *HciDomainManager) GetLicenses() []HciLicenseRelationship`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *HciDomainManager) GetLicensesOk() (*[]HciLicenseRelationship, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *HciDomainManager) SetLicenses(v []HciLicenseRelationship)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *HciDomainManager) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### SetLicensesNil

`func (o *HciDomainManager) SetLicensesNil(b bool)`

 SetLicensesNil sets the value for Licenses to be an explicit nil

### UnsetLicenses
`func (o *HciDomainManager) UnsetLicenses()`

UnsetLicenses ensures that no value is present for Licenses, not even an explicit nil
### GetRegisteredDevice

`func (o *HciDomainManager) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciDomainManager) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciDomainManager) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciDomainManager) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciDomainManager) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciDomainManager) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


