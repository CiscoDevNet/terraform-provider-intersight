# NiatelemetryMsoSchemaDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.MsoSchemaDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.MsoSchemaDetails"]
**DeployedSites** | Pointer to **string** | Site IDs to which this schema is deployed to. | [optional] 
**NumberOfPolicyObjectsPerSchema** | Pointer to **int64** | Number of policy objects per schema. | [optional] 
**NumberOfTemplatesPerSchema** | Pointer to **int64** | Number of tenants assigned per schema in Multi-Site Orchestrator. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**SchemaId** | Pointer to **string** | Schema ID in Multi-Site Orchestrator. | [optional] 
**SchemaName** | Pointer to **string** | Schema name in Multi-Site Orchestrator. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryMsoSchemaDetails

`func NewNiatelemetryMsoSchemaDetails(classId string, objectType string, ) *NiatelemetryMsoSchemaDetails`

NewNiatelemetryMsoSchemaDetails instantiates a new NiatelemetryMsoSchemaDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMsoSchemaDetailsWithDefaults

`func NewNiatelemetryMsoSchemaDetailsWithDefaults() *NiatelemetryMsoSchemaDetails`

NewNiatelemetryMsoSchemaDetailsWithDefaults instantiates a new NiatelemetryMsoSchemaDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMsoSchemaDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMsoSchemaDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMsoSchemaDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMsoSchemaDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMsoSchemaDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMsoSchemaDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeployedSites

`func (o *NiatelemetryMsoSchemaDetails) GetDeployedSites() string`

GetDeployedSites returns the DeployedSites field if non-nil, zero value otherwise.

### GetDeployedSitesOk

`func (o *NiatelemetryMsoSchemaDetails) GetDeployedSitesOk() (*string, bool)`

GetDeployedSitesOk returns a tuple with the DeployedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSites

`func (o *NiatelemetryMsoSchemaDetails) SetDeployedSites(v string)`

SetDeployedSites sets DeployedSites field to given value.

### HasDeployedSites

`func (o *NiatelemetryMsoSchemaDetails) HasDeployedSites() bool`

HasDeployedSites returns a boolean if a field has been set.

### GetNumberOfPolicyObjectsPerSchema

`func (o *NiatelemetryMsoSchemaDetails) GetNumberOfPolicyObjectsPerSchema() int64`

GetNumberOfPolicyObjectsPerSchema returns the NumberOfPolicyObjectsPerSchema field if non-nil, zero value otherwise.

### GetNumberOfPolicyObjectsPerSchemaOk

`func (o *NiatelemetryMsoSchemaDetails) GetNumberOfPolicyObjectsPerSchemaOk() (*int64, bool)`

GetNumberOfPolicyObjectsPerSchemaOk returns a tuple with the NumberOfPolicyObjectsPerSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPolicyObjectsPerSchema

`func (o *NiatelemetryMsoSchemaDetails) SetNumberOfPolicyObjectsPerSchema(v int64)`

SetNumberOfPolicyObjectsPerSchema sets NumberOfPolicyObjectsPerSchema field to given value.

### HasNumberOfPolicyObjectsPerSchema

`func (o *NiatelemetryMsoSchemaDetails) HasNumberOfPolicyObjectsPerSchema() bool`

HasNumberOfPolicyObjectsPerSchema returns a boolean if a field has been set.

### GetNumberOfTemplatesPerSchema

`func (o *NiatelemetryMsoSchemaDetails) GetNumberOfTemplatesPerSchema() int64`

GetNumberOfTemplatesPerSchema returns the NumberOfTemplatesPerSchema field if non-nil, zero value otherwise.

### GetNumberOfTemplatesPerSchemaOk

`func (o *NiatelemetryMsoSchemaDetails) GetNumberOfTemplatesPerSchemaOk() (*int64, bool)`

GetNumberOfTemplatesPerSchemaOk returns a tuple with the NumberOfTemplatesPerSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTemplatesPerSchema

`func (o *NiatelemetryMsoSchemaDetails) SetNumberOfTemplatesPerSchema(v int64)`

SetNumberOfTemplatesPerSchema sets NumberOfTemplatesPerSchema field to given value.

### HasNumberOfTemplatesPerSchema

`func (o *NiatelemetryMsoSchemaDetails) HasNumberOfTemplatesPerSchema() bool`

HasNumberOfTemplatesPerSchema returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryMsoSchemaDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryMsoSchemaDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryMsoSchemaDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryMsoSchemaDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetSchemaId

`func (o *NiatelemetryMsoSchemaDetails) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *NiatelemetryMsoSchemaDetails) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *NiatelemetryMsoSchemaDetails) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *NiatelemetryMsoSchemaDetails) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetSchemaName

`func (o *NiatelemetryMsoSchemaDetails) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *NiatelemetryMsoSchemaDetails) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *NiatelemetryMsoSchemaDetails) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *NiatelemetryMsoSchemaDetails) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMsoSchemaDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMsoSchemaDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMsoSchemaDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMsoSchemaDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


