# NiatelemetryMsoTenantDetails

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

### NewNiatelemetryMsoTenantDetails

`func NewNiatelemetryMsoTenantDetails(classId string, objectType string, ) *NiatelemetryMsoTenantDetails`

NewNiatelemetryMsoTenantDetails instantiates a new NiatelemetryMsoTenantDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMsoTenantDetailsWithDefaults

`func NewNiatelemetryMsoTenantDetailsWithDefaults() *NiatelemetryMsoTenantDetails`

NewNiatelemetryMsoTenantDetailsWithDefaults instantiates a new NiatelemetryMsoTenantDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMsoTenantDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMsoTenantDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMsoTenantDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMsoTenantDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMsoTenantDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMsoTenantDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeployedSites

`func (o *NiatelemetryMsoTenantDetails) GetDeployedSites() string`

GetDeployedSites returns the DeployedSites field if non-nil, zero value otherwise.

### GetDeployedSitesOk

`func (o *NiatelemetryMsoTenantDetails) GetDeployedSitesOk() (*string, bool)`

GetDeployedSitesOk returns a tuple with the DeployedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSites

`func (o *NiatelemetryMsoTenantDetails) SetDeployedSites(v string)`

SetDeployedSites sets DeployedSites field to given value.

### HasDeployedSites

`func (o *NiatelemetryMsoTenantDetails) HasDeployedSites() bool`

HasDeployedSites returns a boolean if a field has been set.

### GetNumberOfSchemasAssignedPerTenantInMso

`func (o *NiatelemetryMsoTenantDetails) GetNumberOfSchemasAssignedPerTenantInMso() int64`

GetNumberOfSchemasAssignedPerTenantInMso returns the NumberOfSchemasAssignedPerTenantInMso field if non-nil, zero value otherwise.

### GetNumberOfSchemasAssignedPerTenantInMsoOk

`func (o *NiatelemetryMsoTenantDetails) GetNumberOfSchemasAssignedPerTenantInMsoOk() (*int64, bool)`

GetNumberOfSchemasAssignedPerTenantInMsoOk returns a tuple with the NumberOfSchemasAssignedPerTenantInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSchemasAssignedPerTenantInMso

`func (o *NiatelemetryMsoTenantDetails) SetNumberOfSchemasAssignedPerTenantInMso(v int64)`

SetNumberOfSchemasAssignedPerTenantInMso sets NumberOfSchemasAssignedPerTenantInMso field to given value.

### HasNumberOfSchemasAssignedPerTenantInMso

`func (o *NiatelemetryMsoTenantDetails) HasNumberOfSchemasAssignedPerTenantInMso() bool`

HasNumberOfSchemasAssignedPerTenantInMso returns a boolean if a field has been set.

### GetSitesEachTenantIsDeployedToInMso

`func (o *NiatelemetryMsoTenantDetails) GetSitesEachTenantIsDeployedToInMso() int64`

GetSitesEachTenantIsDeployedToInMso returns the SitesEachTenantIsDeployedToInMso field if non-nil, zero value otherwise.

### GetSitesEachTenantIsDeployedToInMsoOk

`func (o *NiatelemetryMsoTenantDetails) GetSitesEachTenantIsDeployedToInMsoOk() (*int64, bool)`

GetSitesEachTenantIsDeployedToInMsoOk returns a tuple with the SitesEachTenantIsDeployedToInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesEachTenantIsDeployedToInMso

`func (o *NiatelemetryMsoTenantDetails) SetSitesEachTenantIsDeployedToInMso(v int64)`

SetSitesEachTenantIsDeployedToInMso sets SitesEachTenantIsDeployedToInMso field to given value.

### HasSitesEachTenantIsDeployedToInMso

`func (o *NiatelemetryMsoTenantDetails) HasSitesEachTenantIsDeployedToInMso() bool`

HasSitesEachTenantIsDeployedToInMso returns a boolean if a field has been set.

### GetTenantId

`func (o *NiatelemetryMsoTenantDetails) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NiatelemetryMsoTenantDetails) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NiatelemetryMsoTenantDetails) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *NiatelemetryMsoTenantDetails) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTenantName

`func (o *NiatelemetryMsoTenantDetails) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *NiatelemetryMsoTenantDetails) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *NiatelemetryMsoTenantDetails) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *NiatelemetryMsoTenantDetails) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMsoTenantDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMsoTenantDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMsoTenantDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMsoTenantDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


