# TechsupportmanagementNiaParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "techsupportmanagement.NiaParam"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "techsupportmanagement.NiaParam"]
**CollectionClass** | Pointer to **int32** | CollectionClass specifies which app to collect ND techsupport from - infra/NDI/NDO/NDFC. * &#x60;1&#x60; - Collect infra logs for Nexus Dashboard TACASSIST. * &#x60;2&#x60; - Collect logs for Nexus Dashboard Insights app through NDTACASSIST. * &#x60;3&#x60; - Collect logs for Nexus Dashboard Orchestrator app through NDTACASSIST. * &#x60;4&#x60; - Collect logs for Nexus Dashboard Fabric Controller app through NDTACASSIST. * &#x60;5&#x60; - Collect logs for Nexus Data Broker app through NDTACASSIST. * &#x60;6&#x60; - Collect logs for Nexus Data Broker Orchestrator app through NDTACASSIST. | [optional] [default to 1]
**CollectionLevel** | Pointer to **int32** | CollectionLevel controls the size / depth of the tech support files collected. * &#x60;1&#x60; - Use cached data in the returned collection. * &#x60;2&#x60; - Use current data in the returned collection. | [optional] [default to 1]
**Filename** | Pointer to **string** | Filename specifies an individual filename to collect from the endpoint. | [optional] 
**ForceFresh** | Pointer to **bool** | ForceFresh controls whether to return pre-collected files or force the collection of new files. | [optional] 
**Period** | Pointer to **int64** | Number of days for which to collect techsupport. | [optional] [default to 2]
**Pids** | Pointer to **[]string** |  | [optional] 
**SerialNumbers** | Pointer to **[]string** |  | [optional] 
**UpgradeLogs** | Pointer to **bool** | UpgradeLogs controls the inclusion of upgrade logs in tech support bundles. | [optional] 

## Methods

### NewTechsupportmanagementNiaParam

`func NewTechsupportmanagementNiaParam(classId string, objectType string, ) *TechsupportmanagementNiaParam`

NewTechsupportmanagementNiaParam instantiates a new TechsupportmanagementNiaParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementNiaParamWithDefaults

`func NewTechsupportmanagementNiaParamWithDefaults() *TechsupportmanagementNiaParam`

NewTechsupportmanagementNiaParamWithDefaults instantiates a new TechsupportmanagementNiaParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TechsupportmanagementNiaParam) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementNiaParam) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementNiaParam) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TechsupportmanagementNiaParam) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementNiaParam) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementNiaParam) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollectionClass

`func (o *TechsupportmanagementNiaParam) GetCollectionClass() int32`

GetCollectionClass returns the CollectionClass field if non-nil, zero value otherwise.

### GetCollectionClassOk

`func (o *TechsupportmanagementNiaParam) GetCollectionClassOk() (*int32, bool)`

GetCollectionClassOk returns a tuple with the CollectionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionClass

`func (o *TechsupportmanagementNiaParam) SetCollectionClass(v int32)`

SetCollectionClass sets CollectionClass field to given value.

### HasCollectionClass

`func (o *TechsupportmanagementNiaParam) HasCollectionClass() bool`

HasCollectionClass returns a boolean if a field has been set.

### GetCollectionLevel

`func (o *TechsupportmanagementNiaParam) GetCollectionLevel() int32`

GetCollectionLevel returns the CollectionLevel field if non-nil, zero value otherwise.

### GetCollectionLevelOk

`func (o *TechsupportmanagementNiaParam) GetCollectionLevelOk() (*int32, bool)`

GetCollectionLevelOk returns a tuple with the CollectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionLevel

`func (o *TechsupportmanagementNiaParam) SetCollectionLevel(v int32)`

SetCollectionLevel sets CollectionLevel field to given value.

### HasCollectionLevel

`func (o *TechsupportmanagementNiaParam) HasCollectionLevel() bool`

HasCollectionLevel returns a boolean if a field has been set.

### GetFilename

`func (o *TechsupportmanagementNiaParam) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *TechsupportmanagementNiaParam) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *TechsupportmanagementNiaParam) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *TechsupportmanagementNiaParam) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetForceFresh

`func (o *TechsupportmanagementNiaParam) GetForceFresh() bool`

GetForceFresh returns the ForceFresh field if non-nil, zero value otherwise.

### GetForceFreshOk

`func (o *TechsupportmanagementNiaParam) GetForceFreshOk() (*bool, bool)`

GetForceFreshOk returns a tuple with the ForceFresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceFresh

`func (o *TechsupportmanagementNiaParam) SetForceFresh(v bool)`

SetForceFresh sets ForceFresh field to given value.

### HasForceFresh

`func (o *TechsupportmanagementNiaParam) HasForceFresh() bool`

HasForceFresh returns a boolean if a field has been set.

### GetPeriod

`func (o *TechsupportmanagementNiaParam) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TechsupportmanagementNiaParam) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TechsupportmanagementNiaParam) SetPeriod(v int64)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TechsupportmanagementNiaParam) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPids

`func (o *TechsupportmanagementNiaParam) GetPids() []string`

GetPids returns the Pids field if non-nil, zero value otherwise.

### GetPidsOk

`func (o *TechsupportmanagementNiaParam) GetPidsOk() (*[]string, bool)`

GetPidsOk returns a tuple with the Pids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPids

`func (o *TechsupportmanagementNiaParam) SetPids(v []string)`

SetPids sets Pids field to given value.

### HasPids

`func (o *TechsupportmanagementNiaParam) HasPids() bool`

HasPids returns a boolean if a field has been set.

### SetPidsNil

`func (o *TechsupportmanagementNiaParam) SetPidsNil(b bool)`

 SetPidsNil sets the value for Pids to be an explicit nil

### UnsetPids
`func (o *TechsupportmanagementNiaParam) UnsetPids()`

UnsetPids ensures that no value is present for Pids, not even an explicit nil
### GetSerialNumbers

`func (o *TechsupportmanagementNiaParam) GetSerialNumbers() []string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *TechsupportmanagementNiaParam) GetSerialNumbersOk() (*[]string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *TechsupportmanagementNiaParam) SetSerialNumbers(v []string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *TechsupportmanagementNiaParam) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### SetSerialNumbersNil

`func (o *TechsupportmanagementNiaParam) SetSerialNumbersNil(b bool)`

 SetSerialNumbersNil sets the value for SerialNumbers to be an explicit nil

### UnsetSerialNumbers
`func (o *TechsupportmanagementNiaParam) UnsetSerialNumbers()`

UnsetSerialNumbers ensures that no value is present for SerialNumbers, not even an explicit nil
### GetUpgradeLogs

`func (o *TechsupportmanagementNiaParam) GetUpgradeLogs() bool`

GetUpgradeLogs returns the UpgradeLogs field if non-nil, zero value otherwise.

### GetUpgradeLogsOk

`func (o *TechsupportmanagementNiaParam) GetUpgradeLogsOk() (*bool, bool)`

GetUpgradeLogsOk returns a tuple with the UpgradeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeLogs

`func (o *TechsupportmanagementNiaParam) SetUpgradeLogs(v bool)`

SetUpgradeLogs sets UpgradeLogs field to given value.

### HasUpgradeLogs

`func (o *TechsupportmanagementNiaParam) HasUpgradeLogs() bool`

HasUpgradeLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


