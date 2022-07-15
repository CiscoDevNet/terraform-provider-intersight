# ConvergedinfraServerComplianceDetailsAllOf

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
**PodCompliance** | Pointer to [**ConvergedinfraPodComplianceInfoRelationship**](ConvergedinfraPodComplianceInfoRelationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalSummaryRelationship**](ComputePhysicalSummaryRelationship.md) |  | [optional] 

## Methods

### NewConvergedinfraServerComplianceDetailsAllOf

`func NewConvergedinfraServerComplianceDetailsAllOf(classId string, objectType string, ) *ConvergedinfraServerComplianceDetailsAllOf`

NewConvergedinfraServerComplianceDetailsAllOf instantiates a new ConvergedinfraServerComplianceDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraServerComplianceDetailsAllOfWithDefaults

`func NewConvergedinfraServerComplianceDetailsAllOfWithDefaults() *ConvergedinfraServerComplianceDetailsAllOf`

NewConvergedinfraServerComplianceDetailsAllOfWithDefaults instantiates a new ConvergedinfraServerComplianceDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterCount

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetAdapterCount() int64`

GetAdapterCount returns the AdapterCount field if non-nil, zero value otherwise.

### GetAdapterCountOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetAdapterCountOk() (*int64, bool)`

GetAdapterCountOk returns a tuple with the AdapterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterCount

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetAdapterCount(v int64)`

SetAdapterCount sets AdapterCount field to given value.

### HasAdapterCount

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasAdapterCount() bool`

HasAdapterCount returns a boolean if a field has been set.

### GetFirmware

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetHclStatus

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetHclStatus() string`

GetHclStatus returns the HclStatus field if non-nil, zero value otherwise.

### GetHclStatusOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetHclStatusOk() (*string, bool)`

GetHclStatusOk returns a tuple with the HclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatus

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetHclStatus(v string)`

SetHclStatus sets HclStatus field to given value.

### HasHclStatus

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasHclStatus() bool`

HasHclStatus returns a boolean if a field has been set.

### GetHclStatusReason

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetHclStatusReason() string`

GetHclStatusReason returns the HclStatusReason field if non-nil, zero value otherwise.

### GetHclStatusReasonOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetHclStatusReasonOk() (*string, bool)`

GetHclStatusReasonOk returns a tuple with the HclStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatusReason

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetHclStatusReason(v string)`

SetHclStatusReason sets HclStatusReason field to given value.

### HasHclStatusReason

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasHclStatusReason() bool`

HasHclStatusReason returns a boolean if a field has been set.

### GetModel

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOs

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPlatform

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProcessor

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetProcessor(v string)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetPodCompliance

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship`

GetPodCompliance returns the PodCompliance field if non-nil, zero value otherwise.

### GetPodComplianceOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool)`

GetPodComplianceOk returns a tuple with the PodCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCompliance

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship)`

SetPodCompliance sets PodCompliance field to given value.

### HasPodCompliance

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasPodCompliance() bool`

HasPodCompliance returns a boolean if a field has been set.

### GetServer

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetServer() ComputePhysicalSummaryRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ConvergedinfraServerComplianceDetailsAllOf) GetServerOk() (*ComputePhysicalSummaryRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ConvergedinfraServerComplianceDetailsAllOf) SetServer(v ComputePhysicalSummaryRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *ConvergedinfraServerComplianceDetailsAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


