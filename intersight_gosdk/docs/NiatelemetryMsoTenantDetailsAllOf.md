# NiatelemetryMsoTenantDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.MsoTenantDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.MsoTenantDetails"]
**DeployedSites** | Pointer to **string** | Site IDs to which this tenant is deployed to. | [optional] 
**NumberOfSchemasAssignedPerTenantInMso** | Pointer to **int64** | Number of schemas assigned to each tenant in Multi-Site Orchestrator. | [optional] 
**SitesEachTenantIsDeployedToInMso** | Pointer to **int64** | Number of sites each tenant is deployed to. | [optional] 
**TenantId** | Pointer to **string** | ID of tenant in Multi-Site Orchestrator. | [optional] 
**TenantName** | Pointer to **string** | Name of the tenant in Multi-Site Orchestrator. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryMsoTenantDetailsAllOf

`func NewNiatelemetryMsoTenantDetailsAllOf(classId string, objectType string, ) *NiatelemetryMsoTenantDetailsAllOf`

NewNiatelemetryMsoTenantDetailsAllOf instantiates a new NiatelemetryMsoTenantDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMsoTenantDetailsAllOfWithDefaults

`func NewNiatelemetryMsoTenantDetailsAllOfWithDefaults() *NiatelemetryMsoTenantDetailsAllOf`

NewNiatelemetryMsoTenantDetailsAllOfWithDefaults instantiates a new NiatelemetryMsoTenantDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMsoTenantDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMsoTenantDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeployedSites

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetDeployedSites() string`

GetDeployedSites returns the DeployedSites field if non-nil, zero value otherwise.

### GetDeployedSitesOk

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetDeployedSitesOk() (*string, bool)`

GetDeployedSitesOk returns a tuple with the DeployedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSites

`func (o *NiatelemetryMsoTenantDetailsAllOf) SetDeployedSites(v string)`

SetDeployedSites sets DeployedSites field to given value.

### HasDeployedSites

`func (o *NiatelemetryMsoTenantDetailsAllOf) HasDeployedSites() bool`

HasDeployedSites returns a boolean if a field has been set.

### GetNumberOfSchemasAssignedPerTenantInMso

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetNumberOfSchemasAssignedPerTenantInMso() int64`

GetNumberOfSchemasAssignedPerTenantInMso returns the NumberOfSchemasAssignedPerTenantInMso field if non-nil, zero value otherwise.

### GetNumberOfSchemasAssignedPerTenantInMsoOk

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetNumberOfSchemasAssignedPerTenantInMsoOk() (*int64, bool)`

GetNumberOfSchemasAssignedPerTenantInMsoOk returns a tuple with the NumberOfSchemasAssignedPerTenantInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSchemasAssignedPerTenantInMso

`func (o *NiatelemetryMsoTenantDetailsAllOf) SetNumberOfSchemasAssignedPerTenantInMso(v int64)`

SetNumberOfSchemasAssignedPerTenantInMso sets NumberOfSchemasAssignedPerTenantInMso field to given value.

### HasNumberOfSchemasAssignedPerTenantInMso

`func (o *NiatelemetryMsoTenantDetailsAllOf) HasNumberOfSchemasAssignedPerTenantInMso() bool`

HasNumberOfSchemasAssignedPerTenantInMso returns a boolean if a field has been set.

### GetSitesEachTenantIsDeployedToInMso

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetSitesEachTenantIsDeployedToInMso() int64`

GetSitesEachTenantIsDeployedToInMso returns the SitesEachTenantIsDeployedToInMso field if non-nil, zero value otherwise.

### GetSitesEachTenantIsDeployedToInMsoOk

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetSitesEachTenantIsDeployedToInMsoOk() (*int64, bool)`

GetSitesEachTenantIsDeployedToInMsoOk returns a tuple with the SitesEachTenantIsDeployedToInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesEachTenantIsDeployedToInMso

`func (o *NiatelemetryMsoTenantDetailsAllOf) SetSitesEachTenantIsDeployedToInMso(v int64)`

SetSitesEachTenantIsDeployedToInMso sets SitesEachTenantIsDeployedToInMso field to given value.

### HasSitesEachTenantIsDeployedToInMso

`func (o *NiatelemetryMsoTenantDetailsAllOf) HasSitesEachTenantIsDeployedToInMso() bool`

HasSitesEachTenantIsDeployedToInMso returns a boolean if a field has been set.

### GetTenantId

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NiatelemetryMsoTenantDetailsAllOf) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *NiatelemetryMsoTenantDetailsAllOf) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTenantName

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *NiatelemetryMsoTenantDetailsAllOf) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *NiatelemetryMsoTenantDetailsAllOf) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMsoTenantDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMsoTenantDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMsoTenantDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


