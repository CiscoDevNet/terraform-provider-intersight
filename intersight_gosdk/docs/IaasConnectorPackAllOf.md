# IaasConnectorPackAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.ConnectorPack"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.ConnectorPack"]
**CompleteVersion** | Pointer to **string** | Complete version of the connector pack including build number. | [optional] [readonly] 
**DependencyNames** | Pointer to **[]string** |  | [optional] 
**DownloadedVersion** | Pointer to **string** | Version of the connector pack that is last downloaded successfully to UCSD. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the connector pack running on the UCSD. | [optional] [readonly] 
**State** | Pointer to **string** | State of the connector pack whether it is enabled or disabled. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the connector pack. | [optional] [readonly] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](IaasUcsdInfoRelationship.md) |  | [optional] 

## Methods

### NewIaasConnectorPackAllOf

`func NewIaasConnectorPackAllOf(classId string, objectType string, ) *IaasConnectorPackAllOf`

NewIaasConnectorPackAllOf instantiates a new IaasConnectorPackAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasConnectorPackAllOfWithDefaults

`func NewIaasConnectorPackAllOfWithDefaults() *IaasConnectorPackAllOf`

NewIaasConnectorPackAllOfWithDefaults instantiates a new IaasConnectorPackAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasConnectorPackAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasConnectorPackAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasConnectorPackAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasConnectorPackAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasConnectorPackAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasConnectorPackAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompleteVersion

`func (o *IaasConnectorPackAllOf) GetCompleteVersion() string`

GetCompleteVersion returns the CompleteVersion field if non-nil, zero value otherwise.

### GetCompleteVersionOk

`func (o *IaasConnectorPackAllOf) GetCompleteVersionOk() (*string, bool)`

GetCompleteVersionOk returns a tuple with the CompleteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteVersion

`func (o *IaasConnectorPackAllOf) SetCompleteVersion(v string)`

SetCompleteVersion sets CompleteVersion field to given value.

### HasCompleteVersion

`func (o *IaasConnectorPackAllOf) HasCompleteVersion() bool`

HasCompleteVersion returns a boolean if a field has been set.

### GetDependencyNames

`func (o *IaasConnectorPackAllOf) GetDependencyNames() []string`

GetDependencyNames returns the DependencyNames field if non-nil, zero value otherwise.

### GetDependencyNamesOk

`func (o *IaasConnectorPackAllOf) GetDependencyNamesOk() (*[]string, bool)`

GetDependencyNamesOk returns a tuple with the DependencyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyNames

`func (o *IaasConnectorPackAllOf) SetDependencyNames(v []string)`

SetDependencyNames sets DependencyNames field to given value.

### HasDependencyNames

`func (o *IaasConnectorPackAllOf) HasDependencyNames() bool`

HasDependencyNames returns a boolean if a field has been set.

### SetDependencyNamesNil

`func (o *IaasConnectorPackAllOf) SetDependencyNamesNil(b bool)`

 SetDependencyNamesNil sets the value for DependencyNames to be an explicit nil

### UnsetDependencyNames
`func (o *IaasConnectorPackAllOf) UnsetDependencyNames()`

UnsetDependencyNames ensures that no value is present for DependencyNames, not even an explicit nil
### GetDownloadedVersion

`func (o *IaasConnectorPackAllOf) GetDownloadedVersion() string`

GetDownloadedVersion returns the DownloadedVersion field if non-nil, zero value otherwise.

### GetDownloadedVersionOk

`func (o *IaasConnectorPackAllOf) GetDownloadedVersionOk() (*string, bool)`

GetDownloadedVersionOk returns a tuple with the DownloadedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedVersion

`func (o *IaasConnectorPackAllOf) SetDownloadedVersion(v string)`

SetDownloadedVersion sets DownloadedVersion field to given value.

### HasDownloadedVersion

`func (o *IaasConnectorPackAllOf) HasDownloadedVersion() bool`

HasDownloadedVersion returns a boolean if a field has been set.

### GetName

`func (o *IaasConnectorPackAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IaasConnectorPackAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IaasConnectorPackAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IaasConnectorPackAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *IaasConnectorPackAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IaasConnectorPackAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IaasConnectorPackAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IaasConnectorPackAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *IaasConnectorPackAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IaasConnectorPackAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IaasConnectorPackAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IaasConnectorPackAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetGuid

`func (o *IaasConnectorPackAllOf) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasConnectorPackAllOf) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasConnectorPackAllOf) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasConnectorPackAllOf) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


