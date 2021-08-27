# NiatelemetryMsoContractDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.MsoContractDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.MsoContractDetails"]
**ContractName** | Pointer to **string** | Name of contract in Multi-Site Orchestrator. | [optional] 
**DeployedSites** | Pointer to **string** | Site Ids to which this contract is deployed to. | [optional] 
**IsLocal** | Pointer to **string** | Is the contract local to the Multi-Site Orchestrator. | [optional] 
**Reference** | Pointer to **string** | Unique reference for the contract in the Multi-Site Orchestrator. | [optional] 
**SchemaId** | Pointer to **string** | Schema ID in Multi-Site Orchestrator. | [optional] 
**SchemaName** | Pointer to **string** | Schema name this contract belongs to in Multi-Site Orchestrator. | [optional] 
**TemplateName** | Pointer to **string** | Template name this contract belongs to in Multi-Site Orchestrator. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryMsoContractDetailsAllOf

`func NewNiatelemetryMsoContractDetailsAllOf(classId string, objectType string, ) *NiatelemetryMsoContractDetailsAllOf`

NewNiatelemetryMsoContractDetailsAllOf instantiates a new NiatelemetryMsoContractDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMsoContractDetailsAllOfWithDefaults

`func NewNiatelemetryMsoContractDetailsAllOfWithDefaults() *NiatelemetryMsoContractDetailsAllOf`

NewNiatelemetryMsoContractDetailsAllOfWithDefaults instantiates a new NiatelemetryMsoContractDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMsoContractDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMsoContractDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMsoContractDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMsoContractDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContractName

`func (o *NiatelemetryMsoContractDetailsAllOf) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *NiatelemetryMsoContractDetailsAllOf) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *NiatelemetryMsoContractDetailsAllOf) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetDeployedSites

`func (o *NiatelemetryMsoContractDetailsAllOf) GetDeployedSites() string`

GetDeployedSites returns the DeployedSites field if non-nil, zero value otherwise.

### GetDeployedSitesOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetDeployedSitesOk() (*string, bool)`

GetDeployedSitesOk returns a tuple with the DeployedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSites

`func (o *NiatelemetryMsoContractDetailsAllOf) SetDeployedSites(v string)`

SetDeployedSites sets DeployedSites field to given value.

### HasDeployedSites

`func (o *NiatelemetryMsoContractDetailsAllOf) HasDeployedSites() bool`

HasDeployedSites returns a boolean if a field has been set.

### GetIsLocal

`func (o *NiatelemetryMsoContractDetailsAllOf) GetIsLocal() string`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetIsLocalOk() (*string, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *NiatelemetryMsoContractDetailsAllOf) SetIsLocal(v string)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *NiatelemetryMsoContractDetailsAllOf) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetReference

`func (o *NiatelemetryMsoContractDetailsAllOf) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *NiatelemetryMsoContractDetailsAllOf) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *NiatelemetryMsoContractDetailsAllOf) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSchemaId

`func (o *NiatelemetryMsoContractDetailsAllOf) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *NiatelemetryMsoContractDetailsAllOf) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *NiatelemetryMsoContractDetailsAllOf) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetSchemaName

`func (o *NiatelemetryMsoContractDetailsAllOf) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *NiatelemetryMsoContractDetailsAllOf) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *NiatelemetryMsoContractDetailsAllOf) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetTemplateName

`func (o *NiatelemetryMsoContractDetailsAllOf) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *NiatelemetryMsoContractDetailsAllOf) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *NiatelemetryMsoContractDetailsAllOf) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMsoContractDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMsoContractDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMsoContractDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMsoContractDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


