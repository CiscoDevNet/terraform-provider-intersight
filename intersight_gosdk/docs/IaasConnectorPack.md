# IaasConnectorPack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompleteVersion** | Pointer to **string** | Complete version of the connector pack including build number. | [optional] [readonly] 
**DependencyNames** | Pointer to **[]string** |  | [optional] 
**DownloadedVersion** | Pointer to **string** | Version of the connector pack that is last downloaded successfully to UCSD. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the connector pack running on the UCSD. | [optional] [readonly] 
**State** | Pointer to **string** | State of the connector pack whether it is enabled or disabled. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the connector pack. | [optional] [readonly] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](iaas.UcsdInfo.Relationship.md) |  | [optional] 

## Methods

### NewIaasConnectorPack

`func NewIaasConnectorPack() *IaasConnectorPack`

NewIaasConnectorPack instantiates a new IaasConnectorPack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasConnectorPackWithDefaults

`func NewIaasConnectorPackWithDefaults() *IaasConnectorPack`

NewIaasConnectorPackWithDefaults instantiates a new IaasConnectorPack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleteVersion

`func (o *IaasConnectorPack) GetCompleteVersion() string`

GetCompleteVersion returns the CompleteVersion field if non-nil, zero value otherwise.

### GetCompleteVersionOk

`func (o *IaasConnectorPack) GetCompleteVersionOk() (*string, bool)`

GetCompleteVersionOk returns a tuple with the CompleteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteVersion

`func (o *IaasConnectorPack) SetCompleteVersion(v string)`

SetCompleteVersion sets CompleteVersion field to given value.

### HasCompleteVersion

`func (o *IaasConnectorPack) HasCompleteVersion() bool`

HasCompleteVersion returns a boolean if a field has been set.

### GetDependencyNames

`func (o *IaasConnectorPack) GetDependencyNames() []string`

GetDependencyNames returns the DependencyNames field if non-nil, zero value otherwise.

### GetDependencyNamesOk

`func (o *IaasConnectorPack) GetDependencyNamesOk() (*[]string, bool)`

GetDependencyNamesOk returns a tuple with the DependencyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyNames

`func (o *IaasConnectorPack) SetDependencyNames(v []string)`

SetDependencyNames sets DependencyNames field to given value.

### HasDependencyNames

`func (o *IaasConnectorPack) HasDependencyNames() bool`

HasDependencyNames returns a boolean if a field has been set.

### GetDownloadedVersion

`func (o *IaasConnectorPack) GetDownloadedVersion() string`

GetDownloadedVersion returns the DownloadedVersion field if non-nil, zero value otherwise.

### GetDownloadedVersionOk

`func (o *IaasConnectorPack) GetDownloadedVersionOk() (*string, bool)`

GetDownloadedVersionOk returns a tuple with the DownloadedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedVersion

`func (o *IaasConnectorPack) SetDownloadedVersion(v string)`

SetDownloadedVersion sets DownloadedVersion field to given value.

### HasDownloadedVersion

`func (o *IaasConnectorPack) HasDownloadedVersion() bool`

HasDownloadedVersion returns a boolean if a field has been set.

### GetName

`func (o *IaasConnectorPack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IaasConnectorPack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IaasConnectorPack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IaasConnectorPack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *IaasConnectorPack) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IaasConnectorPack) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IaasConnectorPack) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IaasConnectorPack) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *IaasConnectorPack) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IaasConnectorPack) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IaasConnectorPack) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IaasConnectorPack) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetGuid

`func (o *IaasConnectorPack) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasConnectorPack) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasConnectorPack) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasConnectorPack) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


