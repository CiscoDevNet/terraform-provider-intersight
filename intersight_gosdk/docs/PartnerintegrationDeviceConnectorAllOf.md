# PartnerintegrationDeviceConnectorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.DeviceConnector"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.DeviceConnector"]
**Action** | Pointer to **string** | Action to be performed on the device connector recipe. * &#x60;None&#x60; - Default Value of the action, i.e. do nothing. * &#x60;Build&#x60; - Build the device connector service image. * &#x60;Deploy&#x60; - Deploy the device connector service on the appliance. * &#x60;Upload&#x60; - Upload a file to the Partner Integration Appliance bucket. | [optional] [default to "None"]
**BuildStartTime** | Pointer to **string** | Time when build was triggered. | [optional] [readonly] 
**BuildStatus** | Pointer to **string** | Status of build for device connector recipe. * &#x60;None&#x60; - Default value of the status. i.e. done nothing. * &#x60;BackendInProgress&#x60; - The backend build is in progress. * &#x60;BackendFailed&#x60; - The backend build has failed. * &#x60;DockerInProgress&#x60; - The docker build is in progress. * &#x60;DockerFailed&#x60; - The docker build has failed. * &#x60;Completed&#x60; - The operation completed successfully. | [optional] [readonly] [default to "None"]
**ImageName** | Pointer to **string** | Name of the docker image that is built. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the device connector recipe. | [optional] 
**SrcBucket** | Pointer to **string** | Name of the bucket to pick up the file from. | [optional] [readonly] 
**SrcFileName** | Pointer to **string** | Name of source file to upload. | [optional] [readonly] 
**Logs** | Pointer to [**[]PartnerintegrationDcLogsRelationship**](PartnerintegrationDcLogsRelationship.md) | An array of relationships to partnerintegrationDcLogs resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationDeviceConnectorAllOf

`func NewPartnerintegrationDeviceConnectorAllOf(classId string, objectType string, ) *PartnerintegrationDeviceConnectorAllOf`

NewPartnerintegrationDeviceConnectorAllOf instantiates a new PartnerintegrationDeviceConnectorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationDeviceConnectorAllOfWithDefaults

`func NewPartnerintegrationDeviceConnectorAllOfWithDefaults() *PartnerintegrationDeviceConnectorAllOf`

NewPartnerintegrationDeviceConnectorAllOfWithDefaults instantiates a new PartnerintegrationDeviceConnectorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationDeviceConnectorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationDeviceConnectorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationDeviceConnectorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationDeviceConnectorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *PartnerintegrationDeviceConnectorAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PartnerintegrationDeviceConnectorAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PartnerintegrationDeviceConnectorAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBuildStartTime

`func (o *PartnerintegrationDeviceConnectorAllOf) GetBuildStartTime() string`

GetBuildStartTime returns the BuildStartTime field if non-nil, zero value otherwise.

### GetBuildStartTimeOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetBuildStartTimeOk() (*string, bool)`

GetBuildStartTimeOk returns a tuple with the BuildStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStartTime

`func (o *PartnerintegrationDeviceConnectorAllOf) SetBuildStartTime(v string)`

SetBuildStartTime sets BuildStartTime field to given value.

### HasBuildStartTime

`func (o *PartnerintegrationDeviceConnectorAllOf) HasBuildStartTime() bool`

HasBuildStartTime returns a boolean if a field has been set.

### GetBuildStatus

`func (o *PartnerintegrationDeviceConnectorAllOf) GetBuildStatus() string`

GetBuildStatus returns the BuildStatus field if non-nil, zero value otherwise.

### GetBuildStatusOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetBuildStatusOk() (*string, bool)`

GetBuildStatusOk returns a tuple with the BuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStatus

`func (o *PartnerintegrationDeviceConnectorAllOf) SetBuildStatus(v string)`

SetBuildStatus sets BuildStatus field to given value.

### HasBuildStatus

`func (o *PartnerintegrationDeviceConnectorAllOf) HasBuildStatus() bool`

HasBuildStatus returns a boolean if a field has been set.

### GetImageName

`func (o *PartnerintegrationDeviceConnectorAllOf) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *PartnerintegrationDeviceConnectorAllOf) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *PartnerintegrationDeviceConnectorAllOf) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetName

`func (o *PartnerintegrationDeviceConnectorAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerintegrationDeviceConnectorAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartnerintegrationDeviceConnectorAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSrcBucket

`func (o *PartnerintegrationDeviceConnectorAllOf) GetSrcBucket() string`

GetSrcBucket returns the SrcBucket field if non-nil, zero value otherwise.

### GetSrcBucketOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetSrcBucketOk() (*string, bool)`

GetSrcBucketOk returns a tuple with the SrcBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcBucket

`func (o *PartnerintegrationDeviceConnectorAllOf) SetSrcBucket(v string)`

SetSrcBucket sets SrcBucket field to given value.

### HasSrcBucket

`func (o *PartnerintegrationDeviceConnectorAllOf) HasSrcBucket() bool`

HasSrcBucket returns a boolean if a field has been set.

### GetSrcFileName

`func (o *PartnerintegrationDeviceConnectorAllOf) GetSrcFileName() string`

GetSrcFileName returns the SrcFileName field if non-nil, zero value otherwise.

### GetSrcFileNameOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetSrcFileNameOk() (*string, bool)`

GetSrcFileNameOk returns a tuple with the SrcFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcFileName

`func (o *PartnerintegrationDeviceConnectorAllOf) SetSrcFileName(v string)`

SetSrcFileName sets SrcFileName field to given value.

### HasSrcFileName

`func (o *PartnerintegrationDeviceConnectorAllOf) HasSrcFileName() bool`

HasSrcFileName returns a boolean if a field has been set.

### GetLogs

`func (o *PartnerintegrationDeviceConnectorAllOf) GetLogs() []PartnerintegrationDcLogsRelationship`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetLogsOk() (*[]PartnerintegrationDcLogsRelationship, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PartnerintegrationDeviceConnectorAllOf) SetLogs(v []PartnerintegrationDcLogsRelationship)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PartnerintegrationDeviceConnectorAllOf) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogsNil

`func (o *PartnerintegrationDeviceConnectorAllOf) SetLogsNil(b bool)`

 SetLogsNil sets the value for Logs to be an explicit nil

### UnsetLogs
`func (o *PartnerintegrationDeviceConnectorAllOf) UnsetLogs()`

UnsetLogs ensures that no value is present for Logs, not even an explicit nil
### GetOrganization

`func (o *PartnerintegrationDeviceConnectorAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerintegrationDeviceConnectorAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerintegrationDeviceConnectorAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PartnerintegrationDeviceConnectorAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


