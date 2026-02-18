# CondAdapterBaselineStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AdapterBaselineStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AdapterBaselineStatus"]
**Adapter** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Reason** | Pointer to **string** | The baseline status reason of the adapter. * &#x60;NotApplicable&#x60; - Reason is not applicable for the specified baseline status. * &#x60;Compliant&#x60; - Custom Hcl Baseline properties are matched with server&#39;s properties. * &#x60;ServerModelNotMatched&#x60; - Server model of the server does not matched with custom Hcl baseline&#39;s server model. * &#x60;ProcessorNotMatched&#x60; - Processor of the server does not matched with custom Hcl baseline&#39;s processor. * &#x60;FirmwareVersionNotMatched&#x60; - Firmware version of the server does not matched with custom Hcl baseline&#39;s firmware version. * &#x60;ManagementModeNotMatched&#x60; - ManagementMode of the server does not matched with custom Hcl baseline&#39;s management mode. * &#x60;GenerationNotMatched&#x60; - Generation of the server does not matched with custom Hcl baseline&#39;s generation. * &#x60;PersonalityNotMatched&#x60; - Personality of the server does not matched with custom Hcl baseline&#39;s personality. * &#x60;OsVendorNotMatched&#x60; - Operating system vendor of the server does not matched with custom Hcl baseline&#39;s operating system vendor. * &#x60;OsVersionNotMatched&#x60; - Operating system version of the server does not matched with custom Hcl baseline&#39;s operating system version. * &#x60;AdapterModelNotMatched&#x60; - Adapter model of the server does not matched with custom Hcl baseline&#39;s adapter model. * &#x60;AdapterFirmwareNotMatched&#x60; - Adapter firmware version of the server does not matched with custom Hcl baseline&#39;s adapter adapter firmware version. * &#x60;AdapterTypeNotMatched&#x60; - Adapter Type of the server does not matched with custom Hcl baseline&#39;s adapter type. * &#x60;AdapterDriverProtocolNotMatched&#x60; - Adapter driver name of the server does not matched with custom Hcl baseline&#39;s adapter driver name. * &#x60;AdapterDriverVersionNotMatched&#x60; - Adapter driver version of the server does not matched with custom Hcl baseline&#39;s adapter driver version. * &#x60;AdapterNotMatched&#x60; - One or morecustom Hcl  adapter  of the server does not matched with custom Hcl baseline&#39;s adapter details. * &#x60;AdapterVendorNotMatched&#x60; - Adapter vendor of the server does not matched with custom Hcl baseline&#39;s adapter vendor. | [optional] [readonly] [default to "NotApplicable"]
**Status** | Pointer to **string** | The baseline status of the adapter. * &#x60;NotApplicable&#x60; - Indicates that the custom Hcl baseline is not applicable to the server. * &#x60;Compliant&#x60; - Indicates that the server baseline status is validated and compliant with the defined custom Hcl baseline. * &#x60;NonCompliant&#x60; - Indicates that the server baseline status is not validated or does not meet the defined custom Hcl baseline. | [optional] [readonly] [default to "NotApplicable"]

## Methods

### NewCondAdapterBaselineStatus

`func NewCondAdapterBaselineStatus(classId string, objectType string, ) *CondAdapterBaselineStatus`

NewCondAdapterBaselineStatus instantiates a new CondAdapterBaselineStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAdapterBaselineStatusWithDefaults

`func NewCondAdapterBaselineStatusWithDefaults() *CondAdapterBaselineStatus`

NewCondAdapterBaselineStatusWithDefaults instantiates a new CondAdapterBaselineStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAdapterBaselineStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAdapterBaselineStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAdapterBaselineStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAdapterBaselineStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAdapterBaselineStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAdapterBaselineStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapter

`func (o *CondAdapterBaselineStatus) GetAdapter() MoMoRef`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *CondAdapterBaselineStatus) GetAdapterOk() (*MoMoRef, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *CondAdapterBaselineStatus) SetAdapter(v MoMoRef)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *CondAdapterBaselineStatus) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetReason

`func (o *CondAdapterBaselineStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CondAdapterBaselineStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CondAdapterBaselineStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CondAdapterBaselineStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *CondAdapterBaselineStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CondAdapterBaselineStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CondAdapterBaselineStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CondAdapterBaselineStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


