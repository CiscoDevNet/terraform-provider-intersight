package intersight

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSoftwareSolutionDistributable() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSoftwareSolutionDistributableRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"bundle_type": {
				Description: "The bundle type of the image, as published on cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "User provided description about the file. Cisco provided description for image inventoried from a Cisco repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"download_count": {
				Description: "The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"file_path": {
				Description: "The path of the file in S3/minio bucket.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"guid": {
				Description: "The unique identifier for an image in a Cisco repository.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"import_action": {
				Description: "The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.\n* `None` - No action should be taken on the imported file.\n* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.\n* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.\n* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.\n* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.\n* `PreCache` - Cache the file into the Intersight Appliance.\n* `Cancel` - The cancel import process for the file into the repository.\n* `Extract` - The action to extract the file in the external repository.\n* `Evict` - Evict the cached file from the Intersight Appliance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"import_state": {
				Description: "The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.\n* `ReadyForImport` - The image is ready to be imported into the repository.\n* `Importing` - The image is being imported into the repository.\n* `Imported` - The image has been extracted and imported into the repository.\n* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.\n* `Extracting` - Indicates that the image is being extracted into the repository.\n* `Extracted` - Indicates that the image has been extracted into the repository.\n* `Failed` - The image import from an external source to the repository has failed.\n* `MetaOnly` - The image is present in an external repository.\n* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.\n* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.\n* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.\n* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.\n* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.\n* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.\n* `Invalid` - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"imported_time": {
				Description: "The time at which this image or file was imported/cached into the repositry. if the 'ImportState' is 'Imported', the time at which this image or file was imported. if the 'ImportState' is 'Cached', the time at which this image or file was cached.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_access_time": {
				Description: "The time at which this file was last downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"md5e_tag": {
				Description: "The MD5 ETag for a file that is stored in Intersight repository or in the appliance cache. Warning - MD5 is currently broken and this will be migrated to SHA shortly.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"md5sum": {
				Description: "The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mdfid": {
				Description: "The mdfid of the image provided by cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mod_time": {
				Description: "The time when this managed object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"model": {
				Description: "The endpoint model for which this firmware image is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The name of the file. It is populated as part of the image import operation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"platform_type": {
				Description: "The platform type of the image.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"recommended_build": {
				Description: "The build which is recommended by Cisco.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"release_date": {
				Description: "The date on which the file was released or distributed by its vendor.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"release_notes_url": {
				Description: "The url for the release notes of this image.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sha512sum": {
				Description: "The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"size": {
				Description: "The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"software_advisory_url": {
				Description: "The software advisory, if any, provided by the vendor for this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_type_id": {
				Description: "The software type id provided by cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"solution_name": {
				Description: "The name of the solution in which the image belongs.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sub_type": {
				Description: "The type of the file like OS image, Script etc.\n* `osimage` - The solution OS image for deployment.\n* `script` - The Python script for the solution VM configuration and deployment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vendor": {
				Description: "The vendor or publisher of this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "Vendor provided version for the file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceSoftwareSolutionDistributable().Schema},
				Computed: true,
			}},
	}
}

func dataSourceSoftwareSolutionDistributableRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SoftwareSolutionDistributable{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("bundle_type"); ok {
		x := (v.(string))
		o.SetBundleType(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("download_count"); ok {
		x := int64(v.(int))
		o.SetDownloadCount(x)
	}
	if v, ok := d.GetOk("file_path"); ok {
		x := (v.(string))
		o.SetFilePath(x)
	}
	if v, ok := d.GetOk("guid"); ok {
		x := (v.(string))
		o.SetGuid(x)
	}
	if v, ok := d.GetOk("import_action"); ok {
		x := (v.(string))
		o.SetImportAction(x)
	}
	if v, ok := d.GetOk("import_state"); ok {
		x := (v.(string))
		o.SetImportState(x)
	}
	if v, ok := d.GetOk("imported_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetImportedTime(x)
	}
	if v, ok := d.GetOk("last_access_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastAccessTime(x)
	}
	if v, ok := d.GetOk("md5e_tag"); ok {
		x := (v.(string))
		o.SetMd5eTag(x)
	}
	if v, ok := d.GetOk("md5sum"); ok {
		x := (v.(string))
		o.SetMd5sum(x)
	}
	if v, ok := d.GetOk("mdfid"); ok {
		x := (v.(string))
		o.SetMdfid(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}
	if v, ok := d.GetOk("recommended_build"); ok {
		x := (v.(string))
		o.SetRecommendedBuild(x)
	}
	if v, ok := d.GetOk("release_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetReleaseDate(x)
	}
	if v, ok := d.GetOk("release_notes_url"); ok {
		x := (v.(string))
		o.SetReleaseNotesUrl(x)
	}
	if v, ok := d.GetOk("sha512sum"); ok {
		x := (v.(string))
		o.SetSha512sum(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("software_advisory_url"); ok {
		x := (v.(string))
		o.SetSoftwareAdvisoryUrl(x)
	}
	if v, ok := d.GetOk("software_type_id"); ok {
		x := (v.(string))
		o.SetSoftwareTypeId(x)
	}
	if v, ok := d.GetOk("solution_name"); ok {
		x := (v.(string))
		o.SetSolutionName(x)
	}
	if v, ok := d.GetOk("sub_type"); ok {
		x := (v.(string))
		o.SetSubType(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SoftwareSolutionDistributable object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.SoftwareApi.GetSoftwareSolutionDistributableList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of SoftwareSolutionDistributable: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of SoftwareSolutionDistributable: %s", responseErr.Error())
	}
	count := countResponse.SoftwareSolutionDistributableList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for SoftwareSolutionDistributable data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var softwareSolutionDistributableResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.SoftwareApi.GetSoftwareSolutionDistributableList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching SoftwareSolutionDistributable: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching SoftwareSolutionDistributable: %s", responseErr.Error())
		}
		results := resMo.SoftwareSolutionDistributableList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["bundle_type"] = (s.GetBundleType())

				temp["catalog"] = flattenMapSoftwarerepositoryCatalogRelationship(s.GetCatalog(), d)
				temp["class_id"] = (s.GetClassId())

				temp["component_meta"] = flattenListFirmwareComponentMeta(s.GetComponentMeta(), d)

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["description"] = (s.GetDescription())

				temp["distributable_metas"] = flattenListFirmwareDistributableMetaRelationship(s.GetDistributableMetas(), d)
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["download_count"] = (s.GetDownloadCount())
				temp["file_path"] = (s.GetFilePath())
				temp["guid"] = (s.GetGuid())
				temp["import_action"] = (s.GetImportAction())
				temp["import_state"] = (s.GetImportState())

				temp["imported_time"] = (s.GetImportedTime()).String()

				temp["last_access_time"] = (s.GetLastAccessTime()).String()
				temp["md5e_tag"] = (s.GetMd5eTag())
				temp["md5sum"] = (s.GetMd5sum())
				temp["mdfid"] = (s.GetMdfid())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["platform_type"] = (s.GetPlatformType())
				temp["recommended_build"] = (s.GetRecommendedBuild())

				temp["release"] = flattenMapSoftwarerepositoryReleaseRelationship(s.GetRelease(), d)

				temp["release_date"] = (s.GetReleaseDate()).String()
				temp["release_notes_url"] = (s.GetReleaseNotesUrl())
				temp["sha512sum"] = (s.GetSha512sum())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["size"] = (s.GetSize())
				temp["software_advisory_url"] = (s.GetSoftwareAdvisoryUrl())
				temp["software_type_id"] = (s.GetSoftwareTypeId())
				temp["solution_name"] = (s.GetSolutionName())

				temp["nr_source"] = flattenMapSoftwarerepositoryFileServer(s.GetSource(), d)
				temp["sub_type"] = (s.GetSubType())
				temp["supported_models"] = (s.GetSupportedModels())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vendor"] = (s.GetVendor())
				temp["nr_version"] = (s.GetVersion())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				softwareSolutionDistributableResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(softwareSolutionDistributableResults))
	if err := d.Set("results", softwareSolutionDistributableResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(softwareSolutionDistributableResults[0]["moid"].(string))
	return de
}
