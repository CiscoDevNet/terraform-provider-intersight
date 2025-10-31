# ComputeAdapterUnitPropFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.AdapterUnitPropFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.AdapterUnitPropFilter"]
**Count** | Pointer to **int64** | Number of Adapters to be selected. | [optional] [default to 0]
**Model** | Pointer to **string** | Specific Adapter model to select. * &#x60;Any&#x60; - To select any Adapter model available in XFM. * &#x60;UCSC-P-N7S400GF&#x60; - Cisco UCS X-Series Adapter with 400G capability and dual port configuration. * &#x60;UCSC-P-N7D200GF&#x60; - Cisco UCS X-Series Adapter with 200G capability and dual port configuration. * &#x60;UCSC-P-N3140H&#x60; - Cisco UCS X-Series Adapter with 200G capability and single port configuration. * &#x60;UCSC-P-N7S400GFO&#x60; - Cisco UCS X-Series SmartNIC with 400G QSFP112 PCIe Gen5 NIC. * &#x60;UCSC-P-N7D200GFO&#x60; - Cisco UCS X-Series SmartNIC with 200G QSFP112 PCIe Gen5 NIC. | [optional] [default to "Any"]

## Methods

### NewComputeAdapterUnitPropFilter

`func NewComputeAdapterUnitPropFilter(classId string, objectType string, ) *ComputeAdapterUnitPropFilter`

NewComputeAdapterUnitPropFilter instantiates a new ComputeAdapterUnitPropFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAdapterUnitPropFilterWithDefaults

`func NewComputeAdapterUnitPropFilterWithDefaults() *ComputeAdapterUnitPropFilter`

NewComputeAdapterUnitPropFilterWithDefaults instantiates a new ComputeAdapterUnitPropFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeAdapterUnitPropFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeAdapterUnitPropFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeAdapterUnitPropFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeAdapterUnitPropFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeAdapterUnitPropFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeAdapterUnitPropFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *ComputeAdapterUnitPropFilter) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ComputeAdapterUnitPropFilter) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ComputeAdapterUnitPropFilter) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ComputeAdapterUnitPropFilter) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetModel

`func (o *ComputeAdapterUnitPropFilter) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputeAdapterUnitPropFilter) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputeAdapterUnitPropFilter) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputeAdapterUnitPropFilter) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


