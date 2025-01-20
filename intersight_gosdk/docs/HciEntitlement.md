# HciEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Entitlement"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Entitlement"]
**ClusterExtId** | Pointer to **string** | The identifier of a license, usually in LIC-xxxx format. | [optional] [readonly] 
**Details** | Pointer to [**[]HciEntitlementDetail**](HciEntitlementDetail.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the cluster which could be NUTANIX, VMWARE, or NON_NUTANIX. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciEntitlement

`func NewHciEntitlement(classId string, objectType string, ) *HciEntitlement`

NewHciEntitlement instantiates a new HciEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciEntitlementWithDefaults

`func NewHciEntitlementWithDefaults() *HciEntitlement`

NewHciEntitlementWithDefaults instantiates a new HciEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciEntitlement) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciEntitlement) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciEntitlement) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciEntitlement) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciEntitlement) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciEntitlement) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterExtId

`func (o *HciEntitlement) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciEntitlement) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciEntitlement) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciEntitlement) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetDetails

`func (o *HciEntitlement) GetDetails() []HciEntitlementDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HciEntitlement) GetDetailsOk() (*[]HciEntitlementDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HciEntitlement) SetDetails(v []HciEntitlementDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *HciEntitlement) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *HciEntitlement) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *HciEntitlement) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetType

`func (o *HciEntitlement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciEntitlement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciEntitlement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciEntitlement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCluster

`func (o *HciEntitlement) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciEntitlement) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciEntitlement) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciEntitlement) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciEntitlement) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciEntitlement) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetRegisteredDevice

`func (o *HciEntitlement) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciEntitlement) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciEntitlement) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciEntitlement) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciEntitlement) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciEntitlement) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


