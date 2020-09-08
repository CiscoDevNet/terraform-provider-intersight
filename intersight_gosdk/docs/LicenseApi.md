# \LicenseApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLicenseLicenseInfo**](LicenseApi.md#CreateLicenseLicenseInfo) | **Post** /api/v1/license/LicenseInfos | Create a &#39;license.LicenseInfo&#39; resource.
[**CreateLicenseLicenseReservationOp**](LicenseApi.md#CreateLicenseLicenseReservationOp) | **Post** /api/v1/license/LicenseReservationOps | Create a &#39;license.LicenseReservationOp&#39; resource.
[**GetLicenseAccountLicenseDataByMoid**](LicenseApi.md#GetLicenseAccountLicenseDataByMoid) | **Get** /api/v1/license/AccountLicenseData/{Moid} | Read a &#39;license.AccountLicenseData&#39; resource.
[**GetLicenseAccountLicenseDataList**](LicenseApi.md#GetLicenseAccountLicenseDataList) | **Get** /api/v1/license/AccountLicenseData | Read a &#39;license.AccountLicenseData&#39; resource.
[**GetLicenseCustomerOpByMoid**](LicenseApi.md#GetLicenseCustomerOpByMoid) | **Get** /api/v1/license/CustomerOps/{Moid} | Read a &#39;license.CustomerOp&#39; resource.
[**GetLicenseCustomerOpList**](LicenseApi.md#GetLicenseCustomerOpList) | **Get** /api/v1/license/CustomerOps | Read a &#39;license.CustomerOp&#39; resource.
[**GetLicenseLicenseInfoByMoid**](LicenseApi.md#GetLicenseLicenseInfoByMoid) | **Get** /api/v1/license/LicenseInfos/{Moid} | Read a &#39;license.LicenseInfo&#39; resource.
[**GetLicenseLicenseInfoList**](LicenseApi.md#GetLicenseLicenseInfoList) | **Get** /api/v1/license/LicenseInfos | Read a &#39;license.LicenseInfo&#39; resource.
[**GetLicenseLicenseReservationOpByMoid**](LicenseApi.md#GetLicenseLicenseReservationOpByMoid) | **Get** /api/v1/license/LicenseReservationOps/{Moid} | Read a &#39;license.LicenseReservationOp&#39; resource.
[**GetLicenseLicenseReservationOpList**](LicenseApi.md#GetLicenseLicenseReservationOpList) | **Get** /api/v1/license/LicenseReservationOps | Read a &#39;license.LicenseReservationOp&#39; resource.
[**GetLicenseSmartlicenseTokenByMoid**](LicenseApi.md#GetLicenseSmartlicenseTokenByMoid) | **Get** /api/v1/license/SmartlicenseTokens/{Moid} | Read a &#39;license.SmartlicenseToken&#39; resource.
[**GetLicenseSmartlicenseTokenList**](LicenseApi.md#GetLicenseSmartlicenseTokenList) | **Get** /api/v1/license/SmartlicenseTokens | Read a &#39;license.SmartlicenseToken&#39; resource.
[**PatchLicenseAccountLicenseData**](LicenseApi.md#PatchLicenseAccountLicenseData) | **Patch** /api/v1/license/AccountLicenseData/{Moid} | Update a &#39;license.AccountLicenseData&#39; resource.
[**PatchLicenseCustomerOp**](LicenseApi.md#PatchLicenseCustomerOp) | **Patch** /api/v1/license/CustomerOps/{Moid} | Update a &#39;license.CustomerOp&#39; resource.
[**PatchLicenseLicenseInfo**](LicenseApi.md#PatchLicenseLicenseInfo) | **Patch** /api/v1/license/LicenseInfos/{Moid} | Update a &#39;license.LicenseInfo&#39; resource.
[**PatchLicenseLicenseReservationOp**](LicenseApi.md#PatchLicenseLicenseReservationOp) | **Patch** /api/v1/license/LicenseReservationOps/{Moid} | Update a &#39;license.LicenseReservationOp&#39; resource.
[**PatchLicenseSmartlicenseToken**](LicenseApi.md#PatchLicenseSmartlicenseToken) | **Patch** /api/v1/license/SmartlicenseTokens/{Moid} | Update a &#39;license.SmartlicenseToken&#39; resource.
[**UpdateLicenseAccountLicenseData**](LicenseApi.md#UpdateLicenseAccountLicenseData) | **Post** /api/v1/license/AccountLicenseData/{Moid} | Update a &#39;license.AccountLicenseData&#39; resource.
[**UpdateLicenseCustomerOp**](LicenseApi.md#UpdateLicenseCustomerOp) | **Post** /api/v1/license/CustomerOps/{Moid} | Update a &#39;license.CustomerOp&#39; resource.
[**UpdateLicenseLicenseInfo**](LicenseApi.md#UpdateLicenseLicenseInfo) | **Post** /api/v1/license/LicenseInfos/{Moid} | Update a &#39;license.LicenseInfo&#39; resource.
[**UpdateLicenseLicenseReservationOp**](LicenseApi.md#UpdateLicenseLicenseReservationOp) | **Post** /api/v1/license/LicenseReservationOps/{Moid} | Update a &#39;license.LicenseReservationOp&#39; resource.
[**UpdateLicenseSmartlicenseToken**](LicenseApi.md#UpdateLicenseSmartlicenseToken) | **Post** /api/v1/license/SmartlicenseTokens/{Moid} | Update a &#39;license.SmartlicenseToken&#39; resource.



## CreateLicenseLicenseInfo

> LicenseLicenseInfo CreateLicenseLicenseInfo(ctx).LicenseLicenseInfo(licenseLicenseInfo).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'license.LicenseInfo' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    licenseLicenseInfo := openapiclient.license.LicenseInfo{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}), RefMo: openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}, Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{), RefMo: , Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ActiveAdmin: false, DaysLeft: int64(123), EndTime: "TODO", EnforceMode: "EnforceMode_example", ErrorDesc: "ErrorDesc_example", EvaluationPeriod: int64(123), ExtraEvaluation: int64(123), LicenseCount: int64(123), LicenseState: "LicenseState_example", LicenseType: "LicenseType_example", StartTime: "TODO", TrialAdmin: false, AccountLicenseData: openapiclient.license.AccountLicenseData.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountId: "AccountId_example", AgentData: "AgentData_example", AuthExpireTime: "AuthExpireTime_example", AuthInitialTime: "AuthInitialTime_example", AuthNextTime: "AuthNextTime_example", Category: "Category_example", DefaultLicenseType: "DefaultLicenseType_example", ErrorDesc: "ErrorDesc_example", Group: "Group_example", HighestCompliantLicenseTier: "HighestCompliantLicenseTier_example", LastSync: "TODO", LastUpdatedTime: "TODO", LicenseState: "LicenseState_example", LicenseTechSupportInfo: "LicenseTechSupportInfo_example", RegisterExpireTime: "RegisterExpireTime_example", RegisterInitialTime: "RegisterInitialTime_example", RegisterNextTime: "RegisterNextTime_example", RegistrationStatus: "RegistrationStatus_example", RenewFailureString: "RenewFailureString_example", SmartAccount: "SmartAccount_example", SyncStatus: "SyncStatus_example", VirtualAccount: "VirtualAccount_example", Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", Var2LicenseReservationOp: openapiclient.license.LicenseReservationOp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AuthCode: "AuthCode_example", AuthCodeInstalled: false, ConfirmCode: "ConfirmCode_example", GenerateRequestCode: false, GenerateReturnCode: false, RequestCode: "RequestCode_example", ReturnCode: "ReturnCode_example", Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", Var2LicenseReservationOp: openapiclient.license.LicenseReservationOp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AuthCode: "AuthCode_example", AuthCodeInstalled: false, ConfirmCode: "ConfirmCode_example", GenerateRequestCode: false, GenerateReturnCode: false, RequestCode: "RequestCode_example", ReturnCode: "ReturnCode_example", Account: }, AppRegistrations: []IamAppRegistrationRelationship{openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: , OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: , OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: , Permission: openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{), EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{), System: }), Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: , ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{)})})})}, Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: }}, Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, UserUniqueIdentifier: "UserUniqueIdentifier_example", Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, UserUniqueIdentifier: "UserUniqueIdentifier_example", Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , Resource: , Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), System: })}), Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{), System: }), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", UserUniqueIdentifier: "UserUniqueIdentifier_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", UserUniqueIdentifier: "UserUniqueIdentifier_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: }), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })}}), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })})}), Users: []IamUserRelationship{)}), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{)}), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: , Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: , Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), EndPointRoles: []IamEndPointRoleRelationship{), Idp: , PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), System: }), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , Resource: , Roles: []IamRoleRelationship{)}), Roles: []IamRoleRelationship{), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, User: }), Permission: , Roles: []IamRoleRelationship{), User: }, Permission: , User: }), Permission: , Roles: []IamRoleRelationship{), User: }), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , IpRulesConfiguration: openapiclient.iam.IpAccessManagement.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Enable: false, LastRecoveryTime: "TODO", Holder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , IpRulesConfiguration: openapiclient.iam.IpAccessManagement.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Enable: false, LastRecoveryTime: "TODO", Holder: , IpAddresses: []IamIpAddressRelationship{openapiclient.iam.IpAddress.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Address: "Address_example", Description: "Description_example", IpAccessManagement: })}, ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"}, Roles: []CmrfCmRf{openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"})}), TargetApp: "TargetApp_example", Holder: , Resource: })}, IpAddresses: []IamIpAddressRelationship{openapiclient.iam.IpAddress.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Address: "Address_example", Description: "Description_example", IpAccessManagement: })}, ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: , Roles: []CmrfCmRf{)}), TargetApp: "TargetApp_example", Holder: , Resource: })}, SessionLimits: }}, AppRegistrations: []IamAppRegistrationRelationship{), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: , SessionLimits: }, CustomerOp: openapiclient.license.CustomerOp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ActiveAdmin: false, DeregisterDevice: false, EnableTrial: false, EvaluationPeriod: int64(123), ExtraEvaluation: int64(123), RenewAuthorization: false, RenewIdCertificate: false, ShowAgentTechSupport: false, AccountLicenseData: openapiclient.license.AccountLicenseData.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountId: "AccountId_example", AgentData: "AgentData_example", AuthExpireTime: "AuthExpireTime_example", AuthInitialTime: "AuthInitialTime_example", AuthNextTime: "AuthNextTime_example", Category: "Category_example", DefaultLicenseType: "DefaultLicenseType_example", ErrorDesc: "ErrorDesc_example", Group: "Group_example", HighestCompliantLicenseTier: "HighestCompliantLicenseTier_example", LastSync: "TODO", LastUpdatedTime: "TODO", LicenseState: "LicenseState_example", LicenseTechSupportInfo: "LicenseTechSupportInfo_example", RegisterExpireTime: "RegisterExpireTime_example", RegisterInitialTime: "RegisterInitialTime_example", RegisterNextTime: "RegisterNextTime_example", RegistrationStatus: "RegistrationStatus_example", RenewFailureString: "RenewFailureString_example", SmartAccount: "SmartAccount_example", SyncStatus: "SyncStatus_example", VirtualAccount: "VirtualAccount_example", Account: , CustomerOp: openapiclient.license.CustomerOp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ActiveAdmin: false, DeregisterDevice: false, EnableTrial: false, EvaluationPeriod: int64(123), ExtraEvaluation: int64(123), RenewAuthorization: false, RenewIdCertificate: false, ShowAgentTechSupport: false, AccountLicenseData: }, Licenseinfos: []LicenseLicenseInfoRelationship{openapiclient.license.LicenseInfo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ActiveAdmin: false, DaysLeft: int64(123), EndTime: "TODO", EnforceMode: "EnforceMode_example", ErrorDesc: "ErrorDesc_example", EvaluationPeriod: int64(123), ExtraEvaluation: int64(123), LicenseCount: int64(123), LicenseState: "LicenseState_example", LicenseType: "LicenseType_example", StartTime: "TODO", TrialAdmin: false, AccountLicenseData: }), SmartlicenseToken: openapiclient.license.SmartlicenseToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Token: "Token_example", AccountLicenseData: }}}, Licenseinfos: []LicenseLicenseInfoRelationship{openapiclient.license.LicenseInfo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ActiveAdmin: false, DaysLeft: int64(123), EndTime: "TODO", EnforceMode: "EnforceMode_example", ErrorDesc: "ErrorDesc_example", EvaluationPeriod: int64(123), ExtraEvaluation: int64(123), LicenseCount: int64(123), LicenseState: "LicenseState_example", LicenseType: "LicenseType_example", StartTime: "TODO", TrialAdmin: false, AccountLicenseData: }), SmartlicenseToken: openapiclient.license.SmartlicenseToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Token: "Token_example", AccountLicenseData: }}} // LicenseLicenseInfo | The 'license.LicenseInfo' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.CreateLicenseLicenseInfo(context.Background(), licenseLicenseInfo).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.CreateLicenseLicenseInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLicenseLicenseInfo`: LicenseLicenseInfo
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.CreateLicenseLicenseInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLicenseLicenseInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licenseLicenseInfo** | [**LicenseLicenseInfo**](LicenseLicenseInfo.md) | The &#39;license.LicenseInfo&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**LicenseLicenseInfo**](license.LicenseInfo.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLicenseLicenseReservationOp

> LicenseLicenseReservationOp CreateLicenseLicenseReservationOp(ctx).LicenseLicenseReservationOp(licenseLicenseReservationOp).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'license.LicenseReservationOp' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    licenseLicenseReservationOp := openapiclient.license.LicenseReservationOp{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AuthCode: "AuthCode_example", AuthCodeInstalled: false, ConfirmCode: "ConfirmCode_example", GenerateRequestCode: false, GenerateReturnCode: false, RequestCode: "RequestCode_example", ReturnCode: "ReturnCode_example", Account: } // LicenseLicenseReservationOp | The 'license.LicenseReservationOp' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.CreateLicenseLicenseReservationOp(context.Background(), licenseLicenseReservationOp).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.CreateLicenseLicenseReservationOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLicenseLicenseReservationOp`: LicenseLicenseReservationOp
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.CreateLicenseLicenseReservationOp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLicenseLicenseReservationOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licenseLicenseReservationOp** | [**LicenseLicenseReservationOp**](LicenseLicenseReservationOp.md) | The &#39;license.LicenseReservationOp&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**LicenseLicenseReservationOp**](license.LicenseReservationOp.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseAccountLicenseDataByMoid

> LicenseAccountLicenseData GetLicenseAccountLicenseDataByMoid(ctx, moid).Execute()

Read a 'license.AccountLicenseData' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseAccountLicenseDataByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseAccountLicenseDataByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseAccountLicenseDataByMoid`: LicenseAccountLicenseData
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseAccountLicenseDataByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseAccountLicenseDataByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseAccountLicenseData**](license.AccountLicenseData.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseAccountLicenseDataList

> LicenseAccountLicenseDataResponse GetLicenseAccountLicenseDataList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'license.AccountLicenseData' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseAccountLicenseDataList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseAccountLicenseDataList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseAccountLicenseDataList`: LicenseAccountLicenseDataResponse
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseAccountLicenseDataList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseAccountLicenseDataListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**LicenseAccountLicenseDataResponse**](license.AccountLicenseData.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseCustomerOpByMoid

> LicenseCustomerOp GetLicenseCustomerOpByMoid(ctx, moid).Execute()

Read a 'license.CustomerOp' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseCustomerOpByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseCustomerOpByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseCustomerOpByMoid`: LicenseCustomerOp
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseCustomerOpByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseCustomerOpByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseCustomerOp**](license.CustomerOp.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseCustomerOpList

> LicenseCustomerOpResponse GetLicenseCustomerOpList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'license.CustomerOp' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseCustomerOpList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseCustomerOpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseCustomerOpList`: LicenseCustomerOpResponse
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseCustomerOpList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseCustomerOpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**LicenseCustomerOpResponse**](license.CustomerOp.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLicenseInfoByMoid

> LicenseLicenseInfo GetLicenseLicenseInfoByMoid(ctx, moid).Execute()

Read a 'license.LicenseInfo' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseLicenseInfoByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseLicenseInfoByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLicenseInfoByMoid`: LicenseLicenseInfo
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseLicenseInfoByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLicenseInfoByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseLicenseInfo**](license.LicenseInfo.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLicenseInfoList

> LicenseLicenseInfoResponse GetLicenseLicenseInfoList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'license.LicenseInfo' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseLicenseInfoList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseLicenseInfoList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLicenseInfoList`: LicenseLicenseInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseLicenseInfoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLicenseInfoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**LicenseLicenseInfoResponse**](license.LicenseInfo.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLicenseReservationOpByMoid

> LicenseLicenseReservationOp GetLicenseLicenseReservationOpByMoid(ctx, moid).Execute()

Read a 'license.LicenseReservationOp' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseLicenseReservationOpByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseLicenseReservationOpByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLicenseReservationOpByMoid`: LicenseLicenseReservationOp
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseLicenseReservationOpByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLicenseReservationOpByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseLicenseReservationOp**](license.LicenseReservationOp.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLicenseReservationOpList

> LicenseLicenseReservationOpResponse GetLicenseLicenseReservationOpList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'license.LicenseReservationOp' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseLicenseReservationOpList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseLicenseReservationOpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLicenseReservationOpList`: LicenseLicenseReservationOpResponse
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseLicenseReservationOpList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLicenseReservationOpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**LicenseLicenseReservationOpResponse**](license.LicenseReservationOp.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseSmartlicenseTokenByMoid

> LicenseSmartlicenseToken GetLicenseSmartlicenseTokenByMoid(ctx, moid).Execute()

Read a 'license.SmartlicenseToken' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseSmartlicenseTokenByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseSmartlicenseTokenByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseSmartlicenseTokenByMoid`: LicenseSmartlicenseToken
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseSmartlicenseTokenByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseSmartlicenseTokenByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseSmartlicenseToken**](license.SmartlicenseToken.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseSmartlicenseTokenList

> LicenseSmartlicenseTokenResponse GetLicenseSmartlicenseTokenList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'license.SmartlicenseToken' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetLicenseSmartlicenseTokenList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicenseSmartlicenseTokenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseSmartlicenseTokenList`: LicenseSmartlicenseTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicenseSmartlicenseTokenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseSmartlicenseTokenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**LicenseSmartlicenseTokenResponse**](license.SmartlicenseToken.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLicenseAccountLicenseData

> LicenseAccountLicenseData PatchLicenseAccountLicenseData(ctx, moid).LicenseAccountLicenseData(licenseAccountLicenseData).IfMatch(ifMatch).Execute()

Update a 'license.AccountLicenseData' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseAccountLicenseData := openapiclient.license.AccountLicenseData{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountId: "AccountId_example", AgentData: "AgentData_example", AuthExpireTime: "AuthExpireTime_example", AuthInitialTime: "AuthInitialTime_example", AuthNextTime: "AuthNextTime_example", Category: "Category_example", DefaultLicenseType: "DefaultLicenseType_example", ErrorDesc: "ErrorDesc_example", Group: "Group_example", HighestCompliantLicenseTier: "HighestCompliantLicenseTier_example", LastSync: "TODO", LastUpdatedTime: "TODO", LicenseState: "LicenseState_example", LicenseTechSupportInfo: "LicenseTechSupportInfo_example", RegisterExpireTime: "RegisterExpireTime_example", RegisterInitialTime: "RegisterInitialTime_example", RegisterNextTime: "RegisterNextTime_example", RegistrationStatus: "RegistrationStatus_example", RenewFailureString: "RenewFailureString_example", SmartAccount: "SmartAccount_example", SyncStatus: "SyncStatus_example", VirtualAccount: "VirtualAccount_example", Account: , CustomerOp: , Licenseinfos: []LicenseLicenseInfoRelationship{), SmartlicenseToken: } // LicenseAccountLicenseData | The 'license.AccountLicenseData' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.PatchLicenseAccountLicenseData(context.Background(), moid, licenseAccountLicenseData).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.PatchLicenseAccountLicenseData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchLicenseAccountLicenseData`: LicenseAccountLicenseData
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.PatchLicenseAccountLicenseData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLicenseAccountLicenseDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseAccountLicenseData** | [**LicenseAccountLicenseData**](LicenseAccountLicenseData.md) | The &#39;license.AccountLicenseData&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseAccountLicenseData**](license.AccountLicenseData.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLicenseCustomerOp

> LicenseCustomerOp PatchLicenseCustomerOp(ctx, moid).LicenseCustomerOp(licenseCustomerOp).IfMatch(ifMatch).Execute()

Update a 'license.CustomerOp' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseCustomerOp := openapiclient.license.CustomerOp{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ActiveAdmin: false, DeregisterDevice: false, EnableTrial: false, EvaluationPeriod: int64(123), ExtraEvaluation: int64(123), RenewAuthorization: false, RenewIdCertificate: false, ShowAgentTechSupport: false, AccountLicenseData: } // LicenseCustomerOp | The 'license.CustomerOp' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.PatchLicenseCustomerOp(context.Background(), moid, licenseCustomerOp).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.PatchLicenseCustomerOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchLicenseCustomerOp`: LicenseCustomerOp
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.PatchLicenseCustomerOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLicenseCustomerOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseCustomerOp** | [**LicenseCustomerOp**](LicenseCustomerOp.md) | The &#39;license.CustomerOp&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseCustomerOp**](license.CustomerOp.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLicenseLicenseInfo

> LicenseLicenseInfo PatchLicenseLicenseInfo(ctx, moid).LicenseLicenseInfo(licenseLicenseInfo).IfMatch(ifMatch).Execute()

Update a 'license.LicenseInfo' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseLicenseInfo := openapiclient.license.LicenseInfo{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ActiveAdmin: false, DaysLeft: int64(123), EndTime: "TODO", EnforceMode: "EnforceMode_example", ErrorDesc: "ErrorDesc_example", EvaluationPeriod: int64(123), ExtraEvaluation: int64(123), LicenseCount: int64(123), LicenseState: "LicenseState_example", LicenseType: "LicenseType_example", StartTime: "TODO", TrialAdmin: false, AccountLicenseData: } // LicenseLicenseInfo | The 'license.LicenseInfo' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.PatchLicenseLicenseInfo(context.Background(), moid, licenseLicenseInfo).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.PatchLicenseLicenseInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchLicenseLicenseInfo`: LicenseLicenseInfo
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.PatchLicenseLicenseInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLicenseLicenseInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseLicenseInfo** | [**LicenseLicenseInfo**](LicenseLicenseInfo.md) | The &#39;license.LicenseInfo&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseLicenseInfo**](license.LicenseInfo.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLicenseLicenseReservationOp

> LicenseLicenseReservationOp PatchLicenseLicenseReservationOp(ctx, moid).LicenseLicenseReservationOp(licenseLicenseReservationOp).IfMatch(ifMatch).Execute()

Update a 'license.LicenseReservationOp' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseLicenseReservationOp := openapiclient.license.LicenseReservationOp{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AuthCode: "AuthCode_example", AuthCodeInstalled: false, ConfirmCode: "ConfirmCode_example", GenerateRequestCode: false, GenerateReturnCode: false, RequestCode: "RequestCode_example", ReturnCode: "ReturnCode_example", Account: } // LicenseLicenseReservationOp | The 'license.LicenseReservationOp' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.PatchLicenseLicenseReservationOp(context.Background(), moid, licenseLicenseReservationOp).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.PatchLicenseLicenseReservationOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchLicenseLicenseReservationOp`: LicenseLicenseReservationOp
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.PatchLicenseLicenseReservationOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLicenseLicenseReservationOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseLicenseReservationOp** | [**LicenseLicenseReservationOp**](LicenseLicenseReservationOp.md) | The &#39;license.LicenseReservationOp&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseLicenseReservationOp**](license.LicenseReservationOp.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLicenseSmartlicenseToken

> LicenseSmartlicenseToken PatchLicenseSmartlicenseToken(ctx, moid).LicenseSmartlicenseToken(licenseSmartlicenseToken).IfMatch(ifMatch).Execute()

Update a 'license.SmartlicenseToken' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseSmartlicenseToken := openapiclient.license.SmartlicenseToken{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Token: "Token_example", AccountLicenseData: } // LicenseSmartlicenseToken | The 'license.SmartlicenseToken' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.PatchLicenseSmartlicenseToken(context.Background(), moid, licenseSmartlicenseToken).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.PatchLicenseSmartlicenseToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchLicenseSmartlicenseToken`: LicenseSmartlicenseToken
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.PatchLicenseSmartlicenseToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLicenseSmartlicenseTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseSmartlicenseToken** | [**LicenseSmartlicenseToken**](LicenseSmartlicenseToken.md) | The &#39;license.SmartlicenseToken&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseSmartlicenseToken**](license.SmartlicenseToken.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicenseAccountLicenseData

> LicenseAccountLicenseData UpdateLicenseAccountLicenseData(ctx, moid).LicenseAccountLicenseData(licenseAccountLicenseData).IfMatch(ifMatch).Execute()

Update a 'license.AccountLicenseData' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseAccountLicenseData := openapiclient.license.AccountLicenseData{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountId: "AccountId_example", AgentData: "AgentData_example", AuthExpireTime: "AuthExpireTime_example", AuthInitialTime: "AuthInitialTime_example", AuthNextTime: "AuthNextTime_example", Category: "Category_example", DefaultLicenseType: "DefaultLicenseType_example", ErrorDesc: "ErrorDesc_example", Group: "Group_example", HighestCompliantLicenseTier: "HighestCompliantLicenseTier_example", LastSync: "TODO", LastUpdatedTime: "TODO", LicenseState: "LicenseState_example", LicenseTechSupportInfo: "LicenseTechSupportInfo_example", RegisterExpireTime: "RegisterExpireTime_example", RegisterInitialTime: "RegisterInitialTime_example", RegisterNextTime: "RegisterNextTime_example", RegistrationStatus: "RegistrationStatus_example", RenewFailureString: "RenewFailureString_example", SmartAccount: "SmartAccount_example", SyncStatus: "SyncStatus_example", VirtualAccount: "VirtualAccount_example", Account: , CustomerOp: , Licenseinfos: []LicenseLicenseInfoRelationship{), SmartlicenseToken: } // LicenseAccountLicenseData | The 'license.AccountLicenseData' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.UpdateLicenseAccountLicenseData(context.Background(), moid, licenseAccountLicenseData).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.UpdateLicenseAccountLicenseData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicenseAccountLicenseData`: LicenseAccountLicenseData
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.UpdateLicenseAccountLicenseData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseAccountLicenseDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseAccountLicenseData** | [**LicenseAccountLicenseData**](LicenseAccountLicenseData.md) | The &#39;license.AccountLicenseData&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseAccountLicenseData**](license.AccountLicenseData.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicenseCustomerOp

> LicenseCustomerOp UpdateLicenseCustomerOp(ctx, moid).LicenseCustomerOp(licenseCustomerOp).IfMatch(ifMatch).Execute()

Update a 'license.CustomerOp' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseCustomerOp := openapiclient.license.CustomerOp{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ActiveAdmin: false, DeregisterDevice: false, EnableTrial: false, EvaluationPeriod: int64(123), ExtraEvaluation: int64(123), RenewAuthorization: false, RenewIdCertificate: false, ShowAgentTechSupport: false, AccountLicenseData: } // LicenseCustomerOp | The 'license.CustomerOp' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.UpdateLicenseCustomerOp(context.Background(), moid, licenseCustomerOp).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.UpdateLicenseCustomerOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicenseCustomerOp`: LicenseCustomerOp
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.UpdateLicenseCustomerOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseCustomerOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseCustomerOp** | [**LicenseCustomerOp**](LicenseCustomerOp.md) | The &#39;license.CustomerOp&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseCustomerOp**](license.CustomerOp.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicenseLicenseInfo

> LicenseLicenseInfo UpdateLicenseLicenseInfo(ctx, moid).LicenseLicenseInfo(licenseLicenseInfo).IfMatch(ifMatch).Execute()

Update a 'license.LicenseInfo' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseLicenseInfo :=  // LicenseLicenseInfo | The 'license.LicenseInfo' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.UpdateLicenseLicenseInfo(context.Background(), moid, licenseLicenseInfo).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.UpdateLicenseLicenseInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicenseLicenseInfo`: LicenseLicenseInfo
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.UpdateLicenseLicenseInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseLicenseInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseLicenseInfo** | [**LicenseLicenseInfo**](LicenseLicenseInfo.md) | The &#39;license.LicenseInfo&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseLicenseInfo**](license.LicenseInfo.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicenseLicenseReservationOp

> LicenseLicenseReservationOp UpdateLicenseLicenseReservationOp(ctx, moid).LicenseLicenseReservationOp(licenseLicenseReservationOp).IfMatch(ifMatch).Execute()

Update a 'license.LicenseReservationOp' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseLicenseReservationOp :=  // LicenseLicenseReservationOp | The 'license.LicenseReservationOp' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.UpdateLicenseLicenseReservationOp(context.Background(), moid, licenseLicenseReservationOp).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.UpdateLicenseLicenseReservationOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicenseLicenseReservationOp`: LicenseLicenseReservationOp
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.UpdateLicenseLicenseReservationOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseLicenseReservationOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseLicenseReservationOp** | [**LicenseLicenseReservationOp**](LicenseLicenseReservationOp.md) | The &#39;license.LicenseReservationOp&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseLicenseReservationOp**](license.LicenseReservationOp.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicenseSmartlicenseToken

> LicenseSmartlicenseToken UpdateLicenseSmartlicenseToken(ctx, moid).LicenseSmartlicenseToken(licenseSmartlicenseToken).IfMatch(ifMatch).Execute()

Update a 'license.SmartlicenseToken' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    licenseSmartlicenseToken := openapiclient.license.SmartlicenseToken{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Token: "Token_example", AccountLicenseData: } // LicenseSmartlicenseToken | The 'license.SmartlicenseToken' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.UpdateLicenseSmartlicenseToken(context.Background(), moid, licenseSmartlicenseToken).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.UpdateLicenseSmartlicenseToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicenseSmartlicenseToken`: LicenseSmartlicenseToken
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.UpdateLicenseSmartlicenseToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseSmartlicenseTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseSmartlicenseToken** | [**LicenseSmartlicenseToken**](LicenseSmartlicenseToken.md) | The &#39;license.SmartlicenseToken&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**LicenseSmartlicenseToken**](license.SmartlicenseToken.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

