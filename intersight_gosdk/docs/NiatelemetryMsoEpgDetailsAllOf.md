# NiatelemetryMsoEpgDetailsAllOf

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

### NewNiatelemetryMsoEpgDetailsAllOf

`func NewNiatelemetryMsoEpgDetailsAllOf(classId string, objectType string, ) *NiatelemetryMsoEpgDetailsAllOf`

NewNiatelemetryMsoEpgDetailsAllOf instantiates a new NiatelemetryMsoEpgDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMsoEpgDetailsAllOfWithDefaults

`func NewNiatelemetryMsoEpgDetailsAllOfWithDefaults() *NiatelemetryMsoEpgDetailsAllOf`

NewNiatelemetryMsoEpgDetailsAllOfWithDefaults instantiates a new NiatelemetryMsoEpgDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeployedSites

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetDeployedSites() string`

GetDeployedSites returns the DeployedSites field if non-nil, zero value otherwise.

### GetDeployedSitesOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetDeployedSitesOk() (*string, bool)`

GetDeployedSitesOk returns a tuple with the DeployedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSites

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetDeployedSites(v string)`

SetDeployedSites sets DeployedSites field to given value.

### HasDeployedSites

`func (o *NiatelemetryMsoEpgDetailsAllOf) HasDeployedSites() bool`

HasDeployedSites returns a boolean if a field has been set.

### GetEpgName

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetEpgName() string`

GetEpgName returns the EpgName field if non-nil, zero value otherwise.

### GetEpgNameOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetEpgNameOk() (*string, bool)`

GetEpgNameOk returns a tuple with the EpgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpgName

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetEpgName(v string)`

SetEpgName sets EpgName field to given value.

### HasEpgName

`func (o *NiatelemetryMsoEpgDetailsAllOf) HasEpgName() bool`

HasEpgName returns a boolean if a field has been set.

### GetIsLocal

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetIsLocal() string`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetIsLocalOk() (*string, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetIsLocal(v string)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *NiatelemetryMsoEpgDetailsAllOf) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetReference

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *NiatelemetryMsoEpgDetailsAllOf) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSchemaId

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *NiatelemetryMsoEpgDetailsAllOf) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetSchemaName

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *NiatelemetryMsoEpgDetailsAllOf) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetTemplateName

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *NiatelemetryMsoEpgDetailsAllOf) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMsoEpgDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMsoEpgDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMsoEpgDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


