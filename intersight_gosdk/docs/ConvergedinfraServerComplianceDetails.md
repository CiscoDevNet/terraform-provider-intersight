# ConvergedinfraServerComplianceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.ServerComplianceDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.ServerComplianceDetails"]
**AdapterCount** | Pointer to **int64** | The number of ethernet NIC adapters in the server. | [optional] [readonly] 
**Firmware** | Pointer to **string** | The Cisco IMC firmware version of the server. | [optional] [readonly] 
**HclStatus** | Pointer to **string** | The HCL compatibility status of the server. * &#x60;NotEvaluated&#x60; - The interoperability compliance for the component has not be checked. * &#x60;Approved&#x60; - The component is valid as per the interoperability compliance check. * &#x60;NotApproved&#x60; - The component is not valid as per the interoperability compliance check. * &#x60;Incomplete&#x60; - The interoperability compliance check could not be completed for the component due to incomplete data. | [optional] [readonly] [default to "NotEvaluated"]
**HclStatusReason** | Pointer to **string** | The reason for server&#39;s HCL status. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information. * &#x60;Incompatible-Server&#x60; - The validation failed for this server because the server&#39;s model was not listed in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not listed for the given server model. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has returned a service error or unrecognized result. * &#x60;Not-Evaluated&#x60; - This means the HclStatus for the sever has not been evaluated because it is exempted. * &#x60;Incompatible-Components&#x60; - The validation has failed for this server because one or more components have \&quot;Not-Listed\&quot; status. * &#x60;Compatible&#x60; - The validation has passed for this server&#39;s model, processor, OS vendor and version. | [optional] [readonly] [default to "Missing-Os-Driver-Info"]
**Model** | Pointer to **string** | The model information of the server. | [optional] [readonly] 
**Os** | Pointer to **string** | Details of name and version of the operating system running on the server. | [optional] [readonly] 
**Platform** | Pointer to **string** | Details of platform of the server, examples are B-Series, C-Series, X-Series etc. | [optional] [readonly] 
**Processor** | Pointer to **string** | The processor information of the server. | [optional] [readonly] 
**PodCompliance** | Pointer to [**NullableConvergedinfraPodComplianceInfoRelationship**](ConvergedinfraPodComplianceInfoRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalSummaryRelationship**](ComputePhysicalSummaryRelationship.md) |  | [optional] 

## Methods

### NewConvergedinfraServerComplianceDetails

`func NewConvergedinfraServerComplianceDetails(classId string, objectType string, ) *ConvergedinfraServerComplianceDetails`

NewConvergedinfraServerComplianceDetails instantiates a new ConvergedinfraServerComplianceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraServerComplianceDetailsWithDefaults

`func NewConvergedinfraServerComplianceDetailsWithDefaults() *ConvergedinfraServerComplianceDetails`

NewConvergedinfraServerComplianceDetailsWithDefaults instantiates a new ConvergedinfraServerComplianceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraServerComplianceDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraServerComplianceDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraServerComplianceDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraServerComplianceDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraServerComplianceDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraServerComplianceDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterCount

`func (o *ConvergedinfraServerComplianceDetails) GetAdapterCount() int64`

GetAdapterCount returns the AdapterCount field if non-nil, zero value otherwise.

### GetAdapterCountOk

`func (o *ConvergedinfraServerComplianceDetails) GetAdapterCountOk() (*int64, bool)`

GetAdapterCountOk returns a tuple with the AdapterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterCount

`func (o *ConvergedinfraServerComplianceDetails) SetAdapterCount(v int64)`

SetAdapterCount sets AdapterCount field to given value.

### HasAdapterCount

`func (o *ConvergedinfraServerComplianceDetails) HasAdapterCount() bool`

HasAdapterCount returns a boolean if a field has been set.

### GetFirmware

`func (o *ConvergedinfraServerComplianceDetails) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *ConvergedinfraServerComplianceDetails) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *ConvergedinfraServerComplianceDetails) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *ConvergedinfraServerComplianceDetails) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetHclStatus

`func (o *ConvergedinfraServerComplianceDetails) GetHclStatus() string`

GetHclStatus returns the HclStatus field if non-nil, zero value otherwise.

### GetHclStatusOk

`func (o *ConvergedinfraServerComplianceDetails) GetHclStatusOk() (*string, bool)`

GetHclStatusOk returns a tuple with the HclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatus

`func (o *ConvergedinfraServerComplianceDetails) SetHclStatus(v string)`

SetHclStatus sets HclStatus field to given value.

### HasHclStatus

`func (o *ConvergedinfraServerComplianceDetails) HasHclStatus() bool`

HasHclStatus returns a boolean if a field has been set.

### GetHclStatusReason

`func (o *ConvergedinfraServerComplianceDetails) GetHclStatusReason() string`

GetHclStatusReason returns the HclStatusReason field if non-nil, zero value otherwise.

### GetHclStatusReasonOk

`func (o *ConvergedinfraServerComplianceDetails) GetHclStatusReasonOk() (*string, bool)`

GetHclStatusReasonOk returns a tuple with the HclStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatusReason

`func (o *ConvergedinfraServerComplianceDetails) SetHclStatusReason(v string)`

SetHclStatusReason sets HclStatusReason field to given value.

### HasHclStatusReason

`func (o *ConvergedinfraServerComplianceDetails) HasHclStatusReason() bool`

HasHclStatusReason returns a boolean if a field has been set.

### GetModel

`func (o *ConvergedinfraServerComplianceDetails) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConvergedinfraServerComplianceDetails) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConvergedinfraServerComplianceDetails) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConvergedinfraServerComplianceDetails) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOs

`func (o *ConvergedinfraServerComplianceDetails) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ConvergedinfraServerComplianceDetails) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ConvergedinfraServerComplianceDetails) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ConvergedinfraServerComplianceDetails) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPlatform

`func (o *ConvergedinfraServerComplianceDetails) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ConvergedinfraServerComplianceDetails) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ConvergedinfraServerComplianceDetails) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ConvergedinfraServerComplianceDetails) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProcessor

`func (o *ConvergedinfraServerComplianceDetails) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *ConvergedinfraServerComplianceDetails) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *ConvergedinfraServerComplianceDetails) SetProcessor(v string)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *ConvergedinfraServerComplianceDetails) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetPodCompliance

`func (o *ConvergedinfraServerComplianceDetails) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship`

GetPodCompliance returns the PodCompliance field if non-nil, zero value otherwise.

### GetPodComplianceOk

`func (o *ConvergedinfraServerComplianceDetails) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool)`

GetPodComplianceOk returns a tuple with the PodCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCompliance

`func (o *ConvergedinfraServerComplianceDetails) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship)`

SetPodCompliance sets PodCompliance field to given value.

### HasPodCompliance

`func (o *ConvergedinfraServerComplianceDetails) HasPodCompliance() bool`

HasPodCompliance returns a boolean if a field has been set.

### SetPodComplianceNil

`func (o *ConvergedinfraServerComplianceDetails) SetPodComplianceNil(b bool)`

 SetPodComplianceNil sets the value for PodCompliance to be an explicit nil

### UnsetPodCompliance
`func (o *ConvergedinfraServerComplianceDetails) UnsetPodCompliance()`

UnsetPodCompliance ensures that no value is present for PodCompliance, not even an explicit nil
### GetServer

`func (o *ConvergedinfraServerComplianceDetails) GetServer() ComputePhysicalSummaryRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ConvergedinfraServerComplianceDetails) GetServerOk() (*ComputePhysicalSummaryRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ConvergedinfraServerComplianceDetails) SetServer(v ComputePhysicalSummaryRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *ConvergedinfraServerComplianceDetails) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *ConvergedinfraServerComplianceDetails) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *ConvergedinfraServerComplianceDetails) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


