# PartnerintegrationMetricsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.MetricsModel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.MetricsModel"]
**Attributes** | Pointer to **interface{}** | Transformation attributes model in yaml format. | [optional] 
**DruidInstrumentExporters** | Pointer to **interface{}** | Transformation druid instrument exporters model in yaml format. | [optional] 
**Instruments** | Pointer to **interface{}** | Transformation instruments model in yaml format. | [optional] 
**MeterProviders** | Pointer to **interface{}** | Transformation meter providers model in yaml format. | [optional] 
**Metrics** | Pointer to **interface{}** | Transformation metrics model in yaml format. | [optional] 

## Methods

### NewPartnerintegrationMetricsModel

`func NewPartnerintegrationMetricsModel(classId string, objectType string, ) *PartnerintegrationMetricsModel`

NewPartnerintegrationMetricsModel instantiates a new PartnerintegrationMetricsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationMetricsModelWithDefaults

`func NewPartnerintegrationMetricsModelWithDefaults() *PartnerintegrationMetricsModel`

NewPartnerintegrationMetricsModelWithDefaults instantiates a new PartnerintegrationMetricsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationMetricsModel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationMetricsModel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationMetricsModel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationMetricsModel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationMetricsModel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationMetricsModel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttributes

`func (o *PartnerintegrationMetricsModel) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PartnerintegrationMetricsModel) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PartnerintegrationMetricsModel) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PartnerintegrationMetricsModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *PartnerintegrationMetricsModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *PartnerintegrationMetricsModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetDruidInstrumentExporters

`func (o *PartnerintegrationMetricsModel) GetDruidInstrumentExporters() interface{}`

GetDruidInstrumentExporters returns the DruidInstrumentExporters field if non-nil, zero value otherwise.

### GetDruidInstrumentExportersOk

`func (o *PartnerintegrationMetricsModel) GetDruidInstrumentExportersOk() (*interface{}, bool)`

GetDruidInstrumentExportersOk returns a tuple with the DruidInstrumentExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDruidInstrumentExporters

`func (o *PartnerintegrationMetricsModel) SetDruidInstrumentExporters(v interface{})`

SetDruidInstrumentExporters sets DruidInstrumentExporters field to given value.

### HasDruidInstrumentExporters

`func (o *PartnerintegrationMetricsModel) HasDruidInstrumentExporters() bool`

HasDruidInstrumentExporters returns a boolean if a field has been set.

### SetDruidInstrumentExportersNil

`func (o *PartnerintegrationMetricsModel) SetDruidInstrumentExportersNil(b bool)`

 SetDruidInstrumentExportersNil sets the value for DruidInstrumentExporters to be an explicit nil

### UnsetDruidInstrumentExporters
`func (o *PartnerintegrationMetricsModel) UnsetDruidInstrumentExporters()`

UnsetDruidInstrumentExporters ensures that no value is present for DruidInstrumentExporters, not even an explicit nil
### GetInstruments

`func (o *PartnerintegrationMetricsModel) GetInstruments() interface{}`

GetInstruments returns the Instruments field if non-nil, zero value otherwise.

### GetInstrumentsOk

`func (o *PartnerintegrationMetricsModel) GetInstrumentsOk() (*interface{}, bool)`

GetInstrumentsOk returns a tuple with the Instruments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruments

`func (o *PartnerintegrationMetricsModel) SetInstruments(v interface{})`

SetInstruments sets Instruments field to given value.

### HasInstruments

`func (o *PartnerintegrationMetricsModel) HasInstruments() bool`

HasInstruments returns a boolean if a field has been set.

### SetInstrumentsNil

`func (o *PartnerintegrationMetricsModel) SetInstrumentsNil(b bool)`

 SetInstrumentsNil sets the value for Instruments to be an explicit nil

### UnsetInstruments
`func (o *PartnerintegrationMetricsModel) UnsetInstruments()`

UnsetInstruments ensures that no value is present for Instruments, not even an explicit nil
### GetMeterProviders

`func (o *PartnerintegrationMetricsModel) GetMeterProviders() interface{}`

GetMeterProviders returns the MeterProviders field if non-nil, zero value otherwise.

### GetMeterProvidersOk

`func (o *PartnerintegrationMetricsModel) GetMeterProvidersOk() (*interface{}, bool)`

GetMeterProvidersOk returns a tuple with the MeterProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterProviders

`func (o *PartnerintegrationMetricsModel) SetMeterProviders(v interface{})`

SetMeterProviders sets MeterProviders field to given value.

### HasMeterProviders

`func (o *PartnerintegrationMetricsModel) HasMeterProviders() bool`

HasMeterProviders returns a boolean if a field has been set.

### SetMeterProvidersNil

`func (o *PartnerintegrationMetricsModel) SetMeterProvidersNil(b bool)`

 SetMeterProvidersNil sets the value for MeterProviders to be an explicit nil

### UnsetMeterProviders
`func (o *PartnerintegrationMetricsModel) UnsetMeterProviders()`

UnsetMeterProviders ensures that no value is present for MeterProviders, not even an explicit nil
### GetMetrics

`func (o *PartnerintegrationMetricsModel) GetMetrics() interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *PartnerintegrationMetricsModel) GetMetricsOk() (*interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *PartnerintegrationMetricsModel) SetMetrics(v interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *PartnerintegrationMetricsModel) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetricsNil

`func (o *PartnerintegrationMetricsModel) SetMetricsNil(b bool)`

 SetMetricsNil sets the value for Metrics to be an explicit nil

### UnsetMetrics
`func (o *PartnerintegrationMetricsModel) UnsetMetrics()`

UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


