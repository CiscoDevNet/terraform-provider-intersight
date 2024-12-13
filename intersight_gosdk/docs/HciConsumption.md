# HciConsumption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Consumption"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Consumption"]
**ClusterExtId** | Pointer to **string** | The id of the cluster that consume the license. | [optional] [readonly] 
**LicenseId** | Pointer to **string** | The id of the license instance being consumed. | [optional] [readonly] 
**QuantityUsed** | Pointer to **float64** | The quantity of a license consumed by the cluster. | [optional] [readonly] 

## Methods

### NewHciConsumption

`func NewHciConsumption(classId string, objectType string, ) *HciConsumption`

NewHciConsumption instantiates a new HciConsumption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciConsumptionWithDefaults

`func NewHciConsumptionWithDefaults() *HciConsumption`

NewHciConsumptionWithDefaults instantiates a new HciConsumption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciConsumption) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciConsumption) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciConsumption) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciConsumption) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciConsumption) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciConsumption) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterExtId

`func (o *HciConsumption) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciConsumption) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciConsumption) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciConsumption) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetLicenseId

`func (o *HciConsumption) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *HciConsumption) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *HciConsumption) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *HciConsumption) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetQuantityUsed

`func (o *HciConsumption) GetQuantityUsed() float64`

GetQuantityUsed returns the QuantityUsed field if non-nil, zero value otherwise.

### GetQuantityUsedOk

`func (o *HciConsumption) GetQuantityUsedOk() (*float64, bool)`

GetQuantityUsedOk returns a tuple with the QuantityUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUsed

`func (o *HciConsumption) SetQuantityUsed(v float64)`

SetQuantityUsed sets QuantityUsed field to given value.

### HasQuantityUsed

`func (o *HciConsumption) HasQuantityUsed() bool`

HasQuantityUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


