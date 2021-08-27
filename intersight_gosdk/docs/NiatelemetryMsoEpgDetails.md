# NiatelemetryMsoEpgDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.MsoEpgDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.MsoEpgDetails"]
**DeployedSites** | Pointer to **string** | Site Ids to which this EPG is deployed to. | [optional] 
**EpgName** | Pointer to **string** | Name of EPG in Multi-Site Orchestrator. | [optional] 
**IsLocal** | Pointer to **string** | Is the EPG local to the Multi-Site Orchestrator. | [optional] 
**Reference** | Pointer to **string** | Unique reference for the EPG in the Multi-Site Orchestrator. | [optional] 
**SchemaId** | Pointer to **string** | Schema ID in Multi-Site Orchestrator. | [optional] 
**SchemaName** | Pointer to **string** | Schema name in Multi-Site Orchestrator. | [optional] 
**TemplateName** | Pointer to **string** | Template name in Multi-Site Orchestrator. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryMsoEpgDetails

`func NewNiatelemetryMsoEpgDetails(classId string, objectType string, ) *NiatelemetryMsoEpgDetails`

NewNiatelemetryMsoEpgDetails instantiates a new NiatelemetryMsoEpgDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMsoEpgDetailsWithDefaults

`func NewNiatelemetryMsoEpgDetailsWithDefaults() *NiatelemetryMsoEpgDetails`

NewNiatelemetryMsoEpgDetailsWithDefaults instantiates a new NiatelemetryMsoEpgDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMsoEpgDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMsoEpgDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMsoEpgDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMsoEpgDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMsoEpgDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMsoEpgDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeployedSites

`func (o *NiatelemetryMsoEpgDetails) GetDeployedSites() string`

GetDeployedSites returns the DeployedSites field if non-nil, zero value otherwise.

### GetDeployedSitesOk

`func (o *NiatelemetryMsoEpgDetails) GetDeployedSitesOk() (*string, bool)`

GetDeployedSitesOk returns a tuple with the DeployedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSites

`func (o *NiatelemetryMsoEpgDetails) SetDeployedSites(v string)`

SetDeployedSites sets DeployedSites field to given value.

### HasDeployedSites

`func (o *NiatelemetryMsoEpgDetails) HasDeployedSites() bool`

HasDeployedSites returns a boolean if a field has been set.

### GetEpgName

`func (o *NiatelemetryMsoEpgDetails) GetEpgName() string`

GetEpgName returns the EpgName field if non-nil, zero value otherwise.

### GetEpgNameOk

`func (o *NiatelemetryMsoEpgDetails) GetEpgNameOk() (*string, bool)`

GetEpgNameOk returns a tuple with the EpgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpgName

`func (o *NiatelemetryMsoEpgDetails) SetEpgName(v string)`

SetEpgName sets EpgName field to given value.

### HasEpgName

`func (o *NiatelemetryMsoEpgDetails) HasEpgName() bool`

HasEpgName returns a boolean if a field has been set.

### GetIsLocal

`func (o *NiatelemetryMsoEpgDetails) GetIsLocal() string`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *NiatelemetryMsoEpgDetails) GetIsLocalOk() (*string, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *NiatelemetryMsoEpgDetails) SetIsLocal(v string)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *NiatelemetryMsoEpgDetails) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetReference

`func (o *NiatelemetryMsoEpgDetails) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *NiatelemetryMsoEpgDetails) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *NiatelemetryMsoEpgDetails) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *NiatelemetryMsoEpgDetails) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSchemaId

`func (o *NiatelemetryMsoEpgDetails) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *NiatelemetryMsoEpgDetails) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *NiatelemetryMsoEpgDetails) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *NiatelemetryMsoEpgDetails) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetSchemaName

`func (o *NiatelemetryMsoEpgDetails) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *NiatelemetryMsoEpgDetails) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *NiatelemetryMsoEpgDetails) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *NiatelemetryMsoEpgDetails) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetTemplateName

`func (o *NiatelemetryMsoEpgDetails) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *NiatelemetryMsoEpgDetails) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *NiatelemetryMsoEpgDetails) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *NiatelemetryMsoEpgDetails) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMsoEpgDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMsoEpgDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMsoEpgDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMsoEpgDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


