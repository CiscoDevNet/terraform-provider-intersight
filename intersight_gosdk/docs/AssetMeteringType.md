# AssetMeteringType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.MeteringType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.MeteringType"]
**Name** | Pointer to **string** | Metric type used to calculate metering for the device sent from the IB Contract. example  Node, vMemory, vCPU. * &#x60;None&#x60; - A default value to catch cases where metric type is not correctly detected. * &#x60;Node&#x60; - The metering of the device is on the basis of Power state. * &#x60;Storage&#x60; - The metering of the device is on the basis of used Storage. * &#x60;vMemory&#x60; - The metering of the device is on the basis of VM Memory. * &#x60;vCPU&#x60; - The metering of the device is on the basis of vCPU. * &#x60;vStorage&#x60; - The metering of the device is on the basis of used virtual Storage. * &#x60;Switch&#x60; - The metering of the device is on the basis of Switch. | [optional] [readonly] [default to "None"]
**Unit** | Pointer to **string** | Metric unit used to calculate metering for the device sent from the IB Contract. example  Node, GiB, Cores. * &#x60;None&#x60; - A default value to catch cases where metric unit is not correctly detected. * &#x60;Node&#x60; - It is applicable for Node Metric type. * &#x60;GiB&#x60; - It is applicable for VMemory, vStorage and Storage Metric types. * &#x60;TiB&#x60; - It is applicable for VMemory, vStorage and Storage Metric types. * &#x60;Cores&#x60; - It is applicable for vCPU Metric type. * &#x60;Switch&#x60; - It is applicable for Switch Metric type. * &#x60;Port&#x60; - It is applicable for Switch Metric type. | [optional] [readonly] [default to "None"]

## Methods

### NewAssetMeteringType

`func NewAssetMeteringType(classId string, objectType string, ) *AssetMeteringType`

NewAssetMeteringType instantiates a new AssetMeteringType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetMeteringTypeWithDefaults

`func NewAssetMeteringTypeWithDefaults() *AssetMeteringType`

NewAssetMeteringTypeWithDefaults instantiates a new AssetMeteringType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetMeteringType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetMeteringType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetMeteringType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetMeteringType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetMeteringType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetMeteringType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *AssetMeteringType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetMeteringType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetMeteringType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetMeteringType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnit

`func (o *AssetMeteringType) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AssetMeteringType) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AssetMeteringType) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AssetMeteringType) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


