# HyperflexVcenterConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VcenterConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VcenterConfiguration"]
**ClusterId** | Pointer to **string** | The vCenter compute cluster identifier for the HyperFlex Cluster. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | The vCenter compute cluster name for the HyperFlex cluster. | [optional] [readonly] 
**DatacenterId** | Pointer to **string** | The identifier of the datacenter in vCenter that the HyperFlex cluster belongs to. | [optional] [readonly] 
**DatacenterName** | Pointer to **string** | The name of the datacenter in vCenter that the HyperFlex cluster belongs to. | [optional] [readonly] 
**Url** | Pointer to **string** | The URL of the vCenter for this HyperFlex Cluster. | [optional] [readonly] 

## Methods

### NewHyperflexVcenterConfiguration

`func NewHyperflexVcenterConfiguration(classId string, objectType string, ) *HyperflexVcenterConfiguration`

NewHyperflexVcenterConfiguration instantiates a new HyperflexVcenterConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVcenterConfigurationWithDefaults

`func NewHyperflexVcenterConfigurationWithDefaults() *HyperflexVcenterConfiguration`

NewHyperflexVcenterConfigurationWithDefaults instantiates a new HyperflexVcenterConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVcenterConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVcenterConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVcenterConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVcenterConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVcenterConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVcenterConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterId

`func (o *HyperflexVcenterConfiguration) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HyperflexVcenterConfiguration) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HyperflexVcenterConfiguration) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HyperflexVcenterConfiguration) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterName

`func (o *HyperflexVcenterConfiguration) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HyperflexVcenterConfiguration) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HyperflexVcenterConfiguration) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HyperflexVcenterConfiguration) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetDatacenterId

`func (o *HyperflexVcenterConfiguration) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *HyperflexVcenterConfiguration) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *HyperflexVcenterConfiguration) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *HyperflexVcenterConfiguration) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetDatacenterName

`func (o *HyperflexVcenterConfiguration) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *HyperflexVcenterConfiguration) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *HyperflexVcenterConfiguration) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *HyperflexVcenterConfiguration) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetUrl

`func (o *HyperflexVcenterConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HyperflexVcenterConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HyperflexVcenterConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HyperflexVcenterConfiguration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


