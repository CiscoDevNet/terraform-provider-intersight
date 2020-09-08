# FirmwareBaseImpactAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputationError** | Pointer to **string** | Details for the error that occurred during the reboot validation analysis. | [optional] 
**ComputationStatusDetail** | Pointer to **string** | The computation status of the estimate operation for a component. * &#x60;Inprogress&#x60; - Upgrade impact calculation is in progress. * &#x60;Completed&#x60; - Upgrade impact calculation is completed. * &#x60;Unavailable&#x60; - Upgrade impact is not available since the image is not present in the Fabric Interconnect. * &#x60;Failed&#x60; - Upgrade impact is not available due to an unknown error. | [optional] [default to "Inprogress"]
**DomainName** | Pointer to **string** | The endpoint type or name. | [optional] 
**EndPoint** | Pointer to **string** | A reference to a REST resource, uniquely identified by object type and MOID. | [optional] 
**FirmwareVersion** | Pointer to **string** | The current firmware version of the component. | [optional] 
**ImpactType** | Pointer to **string** | The impact type of the endpoint, whether a reboot is required or not. * &#x60;NoReboot&#x60; - A reboot is not required for the endpoint after upgrade. * &#x60;Reboot&#x60; - A reboot is required to the endpoint after upgrade. | [optional] [default to "NoReboot"]
**Model** | Pointer to **string** | The model details of the component. | [optional] 
**TargetFirmwareVersion** | Pointer to **string** | The target firmware version of the component. | [optional] 
**VersionImpact** | Pointer to **string** | The change of version impact for the endpoint. * &#x60;None&#x60; - No change in version for the component. * &#x60;Upgrade&#x60; - The component will be upgraded. * &#x60;Downgrade&#x60; - The component will be downgraded. | [optional] [default to "None"]

## Methods

### NewFirmwareBaseImpactAllOf

`func NewFirmwareBaseImpactAllOf() *FirmwareBaseImpactAllOf`

NewFirmwareBaseImpactAllOf instantiates a new FirmwareBaseImpactAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBaseImpactAllOfWithDefaults

`func NewFirmwareBaseImpactAllOfWithDefaults() *FirmwareBaseImpactAllOf`

NewFirmwareBaseImpactAllOfWithDefaults instantiates a new FirmwareBaseImpactAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputationError

`func (o *FirmwareBaseImpactAllOf) GetComputationError() string`

GetComputationError returns the ComputationError field if non-nil, zero value otherwise.

### GetComputationErrorOk

`func (o *FirmwareBaseImpactAllOf) GetComputationErrorOk() (*string, bool)`

GetComputationErrorOk returns a tuple with the ComputationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationError

`func (o *FirmwareBaseImpactAllOf) SetComputationError(v string)`

SetComputationError sets ComputationError field to given value.

### HasComputationError

`func (o *FirmwareBaseImpactAllOf) HasComputationError() bool`

HasComputationError returns a boolean if a field has been set.

### GetComputationStatusDetail

`func (o *FirmwareBaseImpactAllOf) GetComputationStatusDetail() string`

GetComputationStatusDetail returns the ComputationStatusDetail field if non-nil, zero value otherwise.

### GetComputationStatusDetailOk

`func (o *FirmwareBaseImpactAllOf) GetComputationStatusDetailOk() (*string, bool)`

GetComputationStatusDetailOk returns a tuple with the ComputationStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationStatusDetail

`func (o *FirmwareBaseImpactAllOf) SetComputationStatusDetail(v string)`

SetComputationStatusDetail sets ComputationStatusDetail field to given value.

### HasComputationStatusDetail

`func (o *FirmwareBaseImpactAllOf) HasComputationStatusDetail() bool`

HasComputationStatusDetail returns a boolean if a field has been set.

### GetDomainName

`func (o *FirmwareBaseImpactAllOf) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *FirmwareBaseImpactAllOf) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *FirmwareBaseImpactAllOf) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *FirmwareBaseImpactAllOf) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEndPoint

`func (o *FirmwareBaseImpactAllOf) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *FirmwareBaseImpactAllOf) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *FirmwareBaseImpactAllOf) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.

### HasEndPoint

`func (o *FirmwareBaseImpactAllOf) HasEndPoint() bool`

HasEndPoint returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *FirmwareBaseImpactAllOf) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *FirmwareBaseImpactAllOf) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *FirmwareBaseImpactAllOf) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *FirmwareBaseImpactAllOf) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetImpactType

`func (o *FirmwareBaseImpactAllOf) GetImpactType() string`

GetImpactType returns the ImpactType field if non-nil, zero value otherwise.

### GetImpactTypeOk

`func (o *FirmwareBaseImpactAllOf) GetImpactTypeOk() (*string, bool)`

GetImpactTypeOk returns a tuple with the ImpactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactType

`func (o *FirmwareBaseImpactAllOf) SetImpactType(v string)`

SetImpactType sets ImpactType field to given value.

### HasImpactType

`func (o *FirmwareBaseImpactAllOf) HasImpactType() bool`

HasImpactType returns a boolean if a field has been set.

### GetModel

`func (o *FirmwareBaseImpactAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FirmwareBaseImpactAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FirmwareBaseImpactAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FirmwareBaseImpactAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTargetFirmwareVersion

`func (o *FirmwareBaseImpactAllOf) GetTargetFirmwareVersion() string`

GetTargetFirmwareVersion returns the TargetFirmwareVersion field if non-nil, zero value otherwise.

### GetTargetFirmwareVersionOk

`func (o *FirmwareBaseImpactAllOf) GetTargetFirmwareVersionOk() (*string, bool)`

GetTargetFirmwareVersionOk returns a tuple with the TargetFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFirmwareVersion

`func (o *FirmwareBaseImpactAllOf) SetTargetFirmwareVersion(v string)`

SetTargetFirmwareVersion sets TargetFirmwareVersion field to given value.

### HasTargetFirmwareVersion

`func (o *FirmwareBaseImpactAllOf) HasTargetFirmwareVersion() bool`

HasTargetFirmwareVersion returns a boolean if a field has been set.

### GetVersionImpact

`func (o *FirmwareBaseImpactAllOf) GetVersionImpact() string`

GetVersionImpact returns the VersionImpact field if non-nil, zero value otherwise.

### GetVersionImpactOk

`func (o *FirmwareBaseImpactAllOf) GetVersionImpactOk() (*string, bool)`

GetVersionImpactOk returns a tuple with the VersionImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionImpact

`func (o *FirmwareBaseImpactAllOf) SetVersionImpact(v string)`

SetVersionImpact sets VersionImpact field to given value.

### HasVersionImpact

`func (o *FirmwareBaseImpactAllOf) HasVersionImpact() bool`

HasVersionImpact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


