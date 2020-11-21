# UcsdConnectorPack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ucsd.ConnectorPack"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ucsd.ConnectorPack"]
**ConnectorFeature** | Pointer to **string** | State of the connector pack whether it is enabled or disabled. | [optional] [readonly] 
**DependencyNames** | Pointer to **[]string** |  | [optional] 
**DownloadedVersion** | Pointer to **string** | Version of the connector pack that is last downloaded successfully to UCS Director. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the connector pack running on the UCS Director. | [optional] [readonly] 
**Services** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** | State of the connector pack whether it is enabled or disabled. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the connector pack. | [optional] [readonly] 

## Methods

### NewUcsdConnectorPack

`func NewUcsdConnectorPack(classId string, objectType string, ) *UcsdConnectorPack`

NewUcsdConnectorPack instantiates a new UcsdConnectorPack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUcsdConnectorPackWithDefaults

`func NewUcsdConnectorPackWithDefaults() *UcsdConnectorPack`

NewUcsdConnectorPackWithDefaults instantiates a new UcsdConnectorPack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UcsdConnectorPack) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UcsdConnectorPack) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UcsdConnectorPack) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UcsdConnectorPack) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UcsdConnectorPack) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UcsdConnectorPack) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectorFeature

`func (o *UcsdConnectorPack) GetConnectorFeature() string`

GetConnectorFeature returns the ConnectorFeature field if non-nil, zero value otherwise.

### GetConnectorFeatureOk

`func (o *UcsdConnectorPack) GetConnectorFeatureOk() (*string, bool)`

GetConnectorFeatureOk returns a tuple with the ConnectorFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorFeature

`func (o *UcsdConnectorPack) SetConnectorFeature(v string)`

SetConnectorFeature sets ConnectorFeature field to given value.

### HasConnectorFeature

`func (o *UcsdConnectorPack) HasConnectorFeature() bool`

HasConnectorFeature returns a boolean if a field has been set.

### GetDependencyNames

`func (o *UcsdConnectorPack) GetDependencyNames() []string`

GetDependencyNames returns the DependencyNames field if non-nil, zero value otherwise.

### GetDependencyNamesOk

`func (o *UcsdConnectorPack) GetDependencyNamesOk() (*[]string, bool)`

GetDependencyNamesOk returns a tuple with the DependencyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyNames

`func (o *UcsdConnectorPack) SetDependencyNames(v []string)`

SetDependencyNames sets DependencyNames field to given value.

### HasDependencyNames

`func (o *UcsdConnectorPack) HasDependencyNames() bool`

HasDependencyNames returns a boolean if a field has been set.

### SetDependencyNamesNil

`func (o *UcsdConnectorPack) SetDependencyNamesNil(b bool)`

 SetDependencyNamesNil sets the value for DependencyNames to be an explicit nil

### UnsetDependencyNames
`func (o *UcsdConnectorPack) UnsetDependencyNames()`

UnsetDependencyNames ensures that no value is present for DependencyNames, not even an explicit nil
### GetDownloadedVersion

`func (o *UcsdConnectorPack) GetDownloadedVersion() string`

GetDownloadedVersion returns the DownloadedVersion field if non-nil, zero value otherwise.

### GetDownloadedVersionOk

`func (o *UcsdConnectorPack) GetDownloadedVersionOk() (*string, bool)`

GetDownloadedVersionOk returns a tuple with the DownloadedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedVersion

`func (o *UcsdConnectorPack) SetDownloadedVersion(v string)`

SetDownloadedVersion sets DownloadedVersion field to given value.

### HasDownloadedVersion

`func (o *UcsdConnectorPack) HasDownloadedVersion() bool`

HasDownloadedVersion returns a boolean if a field has been set.

### GetName

`func (o *UcsdConnectorPack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UcsdConnectorPack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UcsdConnectorPack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UcsdConnectorPack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServices

`func (o *UcsdConnectorPack) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *UcsdConnectorPack) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *UcsdConnectorPack) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *UcsdConnectorPack) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *UcsdConnectorPack) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *UcsdConnectorPack) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetState

`func (o *UcsdConnectorPack) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UcsdConnectorPack) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UcsdConnectorPack) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UcsdConnectorPack) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *UcsdConnectorPack) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UcsdConnectorPack) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UcsdConnectorPack) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UcsdConnectorPack) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


