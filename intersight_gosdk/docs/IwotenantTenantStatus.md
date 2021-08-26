# IwotenantTenantStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iwotenant.TenantStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iwotenant.TenantStatus"]
**DeployStatus** | Pointer to **string** | The deployStatus provides the current status of this deployment. * &#x60;NotStarted&#x60; - The workflow to deploy the tenant cluster has not yet started. * &#x60;InProgress&#x60; - The workflow to deploy the tenant cluster in progress. All the tasks required for thesuccessful tenant creation are running. * &#x60;Completed&#x60; - The workflow to deploy the tenant cluster has completed and health checks have passed. * &#x60;Failed&#x60; - The workflow to deploy the tenant cluster has failed. Detailed reason for the failure isprovided from Tenant.deployStatus. | [optional] [readonly] [default to "NotStarted"]
**IwoId** | Pointer to **string** | The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId. | [optional] [readonly] 
**ReferenceTime** | Pointer to **time.Time** | During IWO tenant upgrade (or reconfiguration), deployStatus is set to InProgress and referenceTime set to current time. When tenant upgrade (or reconfiguration) does not complete within a pre-defined time using this as reference, deployStatus is set as Failed. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIwotenantTenantStatus

`func NewIwotenantTenantStatus(classId string, objectType string, ) *IwotenantTenantStatus`

NewIwotenantTenantStatus instantiates a new IwotenantTenantStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIwotenantTenantStatusWithDefaults

`func NewIwotenantTenantStatusWithDefaults() *IwotenantTenantStatus`

NewIwotenantTenantStatusWithDefaults instantiates a new IwotenantTenantStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IwotenantTenantStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IwotenantTenantStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IwotenantTenantStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IwotenantTenantStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IwotenantTenantStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IwotenantTenantStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeployStatus

`func (o *IwotenantTenantStatus) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *IwotenantTenantStatus) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *IwotenantTenantStatus) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.

### HasDeployStatus

`func (o *IwotenantTenantStatus) HasDeployStatus() bool`

HasDeployStatus returns a boolean if a field has been set.

### GetIwoId

`func (o *IwotenantTenantStatus) GetIwoId() string`

GetIwoId returns the IwoId field if non-nil, zero value otherwise.

### GetIwoIdOk

`func (o *IwotenantTenantStatus) GetIwoIdOk() (*string, bool)`

GetIwoIdOk returns a tuple with the IwoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwoId

`func (o *IwotenantTenantStatus) SetIwoId(v string)`

SetIwoId sets IwoId field to given value.

### HasIwoId

`func (o *IwotenantTenantStatus) HasIwoId() bool`

HasIwoId returns a boolean if a field has been set.

### GetReferenceTime

`func (o *IwotenantTenantStatus) GetReferenceTime() time.Time`

GetReferenceTime returns the ReferenceTime field if non-nil, zero value otherwise.

### GetReferenceTimeOk

`func (o *IwotenantTenantStatus) GetReferenceTimeOk() (*time.Time, bool)`

GetReferenceTimeOk returns a tuple with the ReferenceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTime

`func (o *IwotenantTenantStatus) SetReferenceTime(v time.Time)`

SetReferenceTime sets ReferenceTime field to given value.

### HasReferenceTime

`func (o *IwotenantTenantStatus) HasReferenceTime() bool`

HasReferenceTime returns a boolean if a field has been set.

### GetAccount

`func (o *IwotenantTenantStatus) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IwotenantTenantStatus) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IwotenantTenantStatus) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IwotenantTenantStatus) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


