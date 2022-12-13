# NiatelemetryNexusCloudSiteAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusCloudSite"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusCloudSite"]
**Advisories** | Pointer to **bool** | Advisories setting status, based on license type. | [optional] 
**Anomalies** | Pointer to **bool** | Anomalies setting status, based on license type. | [optional] 
**CapacityUtilization** | Pointer to **bool** | Capacity utilization setting status, based on license type. | [optional] 
**LicenseType** | Pointer to **string** | Returns the license type of the device. | [optional] 
**Name** | Pointer to **string** | Returns the name of the Nexus Cloud site. | [optional] 
**SiteType** | Pointer to **string** | Returns the type of the Nexus Cloud site. | [optional] 
**SoftwareManagement** | Pointer to **bool** | Software management setting status, based on license type. | [optional] 
**Uuid** | Pointer to **string** | Returns the uuid of the Nexus Cloud site. | [optional] 
**NexusCloudAccount** | Pointer to [**NiatelemetryNexusCloudAccountRelationship**](NiatelemetryNexusCloudAccountRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNexusCloudSiteAllOf

`func NewNiatelemetryNexusCloudSiteAllOf(classId string, objectType string, ) *NiatelemetryNexusCloudSiteAllOf`

NewNiatelemetryNexusCloudSiteAllOf instantiates a new NiatelemetryNexusCloudSiteAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNexusCloudSiteAllOfWithDefaults

`func NewNiatelemetryNexusCloudSiteAllOfWithDefaults() *NiatelemetryNexusCloudSiteAllOf`

NewNiatelemetryNexusCloudSiteAllOfWithDefaults instantiates a new NiatelemetryNexusCloudSiteAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNexusCloudSiteAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNexusCloudSiteAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNexusCloudSiteAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNexusCloudSiteAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdvisories

`func (o *NiatelemetryNexusCloudSiteAllOf) GetAdvisories() bool`

GetAdvisories returns the Advisories field if non-nil, zero value otherwise.

### GetAdvisoriesOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetAdvisoriesOk() (*bool, bool)`

GetAdvisoriesOk returns a tuple with the Advisories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisories

`func (o *NiatelemetryNexusCloudSiteAllOf) SetAdvisories(v bool)`

SetAdvisories sets Advisories field to given value.

### HasAdvisories

`func (o *NiatelemetryNexusCloudSiteAllOf) HasAdvisories() bool`

HasAdvisories returns a boolean if a field has been set.

### GetAnomalies

`func (o *NiatelemetryNexusCloudSiteAllOf) GetAnomalies() bool`

GetAnomalies returns the Anomalies field if non-nil, zero value otherwise.

### GetAnomaliesOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetAnomaliesOk() (*bool, bool)`

GetAnomaliesOk returns a tuple with the Anomalies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalies

`func (o *NiatelemetryNexusCloudSiteAllOf) SetAnomalies(v bool)`

SetAnomalies sets Anomalies field to given value.

### HasAnomalies

`func (o *NiatelemetryNexusCloudSiteAllOf) HasAnomalies() bool`

HasAnomalies returns a boolean if a field has been set.

### GetCapacityUtilization

`func (o *NiatelemetryNexusCloudSiteAllOf) GetCapacityUtilization() bool`

GetCapacityUtilization returns the CapacityUtilization field if non-nil, zero value otherwise.

### GetCapacityUtilizationOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetCapacityUtilizationOk() (*bool, bool)`

GetCapacityUtilizationOk returns a tuple with the CapacityUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUtilization

`func (o *NiatelemetryNexusCloudSiteAllOf) SetCapacityUtilization(v bool)`

SetCapacityUtilization sets CapacityUtilization field to given value.

### HasCapacityUtilization

`func (o *NiatelemetryNexusCloudSiteAllOf) HasCapacityUtilization() bool`

HasCapacityUtilization returns a boolean if a field has been set.

### GetLicenseType

`func (o *NiatelemetryNexusCloudSiteAllOf) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *NiatelemetryNexusCloudSiteAllOf) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *NiatelemetryNexusCloudSiteAllOf) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryNexusCloudSiteAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryNexusCloudSiteAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryNexusCloudSiteAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteType

`func (o *NiatelemetryNexusCloudSiteAllOf) GetSiteType() string`

GetSiteType returns the SiteType field if non-nil, zero value otherwise.

### GetSiteTypeOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetSiteTypeOk() (*string, bool)`

GetSiteTypeOk returns a tuple with the SiteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteType

`func (o *NiatelemetryNexusCloudSiteAllOf) SetSiteType(v string)`

SetSiteType sets SiteType field to given value.

### HasSiteType

`func (o *NiatelemetryNexusCloudSiteAllOf) HasSiteType() bool`

HasSiteType returns a boolean if a field has been set.

### GetSoftwareManagement

`func (o *NiatelemetryNexusCloudSiteAllOf) GetSoftwareManagement() bool`

GetSoftwareManagement returns the SoftwareManagement field if non-nil, zero value otherwise.

### GetSoftwareManagementOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetSoftwareManagementOk() (*bool, bool)`

GetSoftwareManagementOk returns a tuple with the SoftwareManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareManagement

`func (o *NiatelemetryNexusCloudSiteAllOf) SetSoftwareManagement(v bool)`

SetSoftwareManagement sets SoftwareManagement field to given value.

### HasSoftwareManagement

`func (o *NiatelemetryNexusCloudSiteAllOf) HasSoftwareManagement() bool`

HasSoftwareManagement returns a boolean if a field has been set.

### GetUuid

`func (o *NiatelemetryNexusCloudSiteAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NiatelemetryNexusCloudSiteAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NiatelemetryNexusCloudSiteAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetNexusCloudAccount

`func (o *NiatelemetryNexusCloudSiteAllOf) GetNexusCloudAccount() NiatelemetryNexusCloudAccountRelationship`

GetNexusCloudAccount returns the NexusCloudAccount field if non-nil, zero value otherwise.

### GetNexusCloudAccountOk

`func (o *NiatelemetryNexusCloudSiteAllOf) GetNexusCloudAccountOk() (*NiatelemetryNexusCloudAccountRelationship, bool)`

GetNexusCloudAccountOk returns a tuple with the NexusCloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusCloudAccount

`func (o *NiatelemetryNexusCloudSiteAllOf) SetNexusCloudAccount(v NiatelemetryNexusCloudAccountRelationship)`

SetNexusCloudAccount sets NexusCloudAccount field to given value.

### HasNexusCloudAccount

`func (o *NiatelemetryNexusCloudSiteAllOf) HasNexusCloudAccount() bool`

HasNexusCloudAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


