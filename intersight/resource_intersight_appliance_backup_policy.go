package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplianceBackupPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplianceBackupPolicyCreate,
		ReadContext:   resourceApplianceBackupPolicyRead,
		UpdateContext: resourceApplianceBackupPolicyUpdate,
		DeleteContext: resourceApplianceBackupPolicyDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		CustomizeDiff: CombinedCustomizeDiff,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"ancestors": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"backup_time": {
				Description: "The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "appliance.BackupPolicy",
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"filename": {
				Description:  "Backup filename to backup or restore.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^$|^[a-zA-Z0-9][a-zA-Z0-9_\\.\\-\\+]*$"), ""),
				Optional:     true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"manual_backup": {
				Description: "Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to 'true' and the backup schedule field is ignored.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"mod_time": {
				Description: "The time when this managed object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "appliance.BackupPolicy",
			},
			"owners": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}},
			"parent": {
				Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"password": {
				Description:  "Password to authenticate the file server.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^$|^[^`]+$"), ""),
				Optional:     true,
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"protocol": {
				Description:  "Communication protocol used by the file server (e.g. scp, sftp, or CIFS).\n* `scp` - Secure Copy Protocol (SCP) to access the file server.\n* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.\n* `cifs` - Common Internet File System (CIFS) Protocol to access file server.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"scp", "sftp", "cifs"}, false),
				Optional:     true,
				Default:      "scp",
			},
			"remote_host": {
				Description: "Hostname of the remote file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"remote_path": {
				Description:  "File server directory or share name to copy the file.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^$|^[^`]+$"), ""),
				Optional:     true,
			},
			"remote_port": {
				Description: "Remote TCP port on the file server (e.g. 22 for scp).",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"retention_count": {
				Description: "The number of backups before earliest backup is overwritten. Requires cleanup policy to be enabled.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
			},
			"retention_policy_enabled": {
				Description: "If backup rotate policy is set, older backups will automatically be overwritten. The number of backups before overwriting is defined by the retentionCount property.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"schedule": {
				Description: "Schedule to create a backup of the Intersight Appliance. Manualbackup field must be set to 'false' for this schedule to be active.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "onprem.Schedule",
						},
						"day_of_month": {
							Description:  "Schedule a task on a specific day of the month. Valid values are 1 through 31. If monthOfYear is specified, then dayOfMonth value must be valid for that month. DayOfMonth may not be set when dayOfWeek is specfied.",
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 31),
							Optional:     true,
						},
						"day_of_week": {
							Description:  "Schedule a task on a specific day of the week. Valid values are 1 through 7, with 1 being Sunday. DayOfWeek may not be specfied when dayOfMonth is specified.",
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
						},
						"month_of_year": {
							Description:  "Schedule a task on a specific month of the year. Valid values are 1 through 12, with 1 being January.",
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 12),
							Optional:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "onprem.Schedule",
						},
						"repeat_interval": {
							Description: "Schedule a task to run periodically at an interval. Default unit of the RepeatInterval is in minutes. If the RepeateInterval value is set, then all other properties are ignored by the scheduler. RepeateInterval constraints are enforced by the services that use the schedule. Each service has pre-configured service specific properties for enforcing minimum and maximum values of the RepeatInterval.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"time_of_day": {
							Description:  "Time of the day in seconds. TimeOfDay is required for all schedule configurations, except when the RepeateInterval field is specified.",
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 86399),
							Optional:     true,
						},
						"time_zone": {
							Description:  "Timezone to use for the schedule calculation. If a timezone value is not speficied, then the schedule calculation will be based on UTC.\n* `Pacific/Niue` - \n* `Africa/Abidjan` - \n* `Africa/Accra` - \n* `Africa/Addis_Ababa` - \n* `Africa/Algiers` - \n* `Africa/Asmara` - \n* `Africa/Bamako` - \n* `Africa/Bangui` - \n* `Africa/Banjul` - \n* `Africa/Bissau` - \n* `Africa/Blantyre` - \n* `Africa/Brazzaville` - \n* `Africa/Bujumbura` - \n* `Africa/Cairo` - \n* `Africa/Casablanca` - \n* `Africa/Ceuta` - \n* `Africa/Conakry` - \n* `Africa/Dakar` - \n* `Africa/Dar_es_Salaam` - \n* `Africa/Djibouti` - \n* `Africa/Douala` - \n* `Africa/El_Aaiun` - \n* `Africa/Freetown` - \n* `Africa/Gaborone` - \n* `Africa/Harare` - \n* `Africa/Johannesburg` - \n* `Africa/Juba` - \n* `Africa/Kampala` - \n* `Africa/Khartoum` - \n* `Africa/Kigali` - \n* `Africa/Kinshasa` - \n* `Africa/Lagos` - \n* `Africa/Libreville` - \n* `Africa/Lome` - \n* `Africa/Luanda` - \n* `Africa/Lubumbashi` - \n* `Africa/Lusaka` - \n* `Africa/Malabo` - \n* `Africa/Maputo` - \n* `Africa/Maseru` - \n* `Africa/Mbabane` - \n* `Africa/Mogadishu` - \n* `Africa/Monrovia` - \n* `Africa/Nairobi` - \n* `Africa/Ndjamena` - \n* `Africa/Niamey` - \n* `Africa/Nouakchott` - \n* `Africa/Ouagadougou` - \n* `Africa/Porto-Novo` - \n* `Africa/Sao_Tome` - \n* `Africa/Tripoli` - \n* `Africa/Tunis` - \n* `Africa/Windhoek` - \n* `America/Adak` - \n* `America/Anchorage` - \n* `America/Anguilla` - \n* `America/Antigua` - \n* `America/Araguaina` - \n* `America/Argentina/Buenos_Aires` - \n* `America/Argentina/Catamarca` - \n* `America/Argentina/Cordoba` - \n* `America/Argentina/Jujuy` - \n* `America/Argentina/La_Rioja` - \n* `America/Argentina/Mendoza` - \n* `America/Argentina/Rio_Gallegos` - \n* `America/Argentina/Salta` - \n* `America/Argentina/San_Juan` - \n* `America/Argentina/San_Luis` - \n* `America/Argentina/Tucuman` - \n* `America/Argentina/Ushuaia` - \n* `America/Aruba` - \n* `America/Asuncion` - \n* `America/Atikokan` - \n* `America/Bahia` - \n* `America/Bahia_Banderas` - \n* `America/Barbados` - \n* `America/Belem` - \n* `America/Belize` - \n* `America/Blanc-Sablon` - \n* `America/Boa_Vista` - \n* `America/Bogota` - \n* `America/Boise` - \n* `America/Cambridge_Bay` - \n* `America/Campo_Grande` - \n* `America/Cancun` - \n* `America/Caracas` - \n* `America/Cayenne` - \n* `America/Cayman` - \n* `America/Chicago` - \n* `America/Chihuahua` - \n* `America/Costa_Rica` - \n* `America/Creston` - \n* `America/Cuiaba` - \n* `America/Curacao` - \n* `America/Danmarkshavn` - \n* `America/Dawson` - \n* `America/Dawson_Creek` - \n* `America/Denver` - \n* `America/Detroit` - \n* `America/Dominica` - \n* `America/Edmonton` - \n* `America/Eirunepe` - \n* `America/El_Salvador` - \n* `America/Fortaleza` - \n* `America/Glace_Bay` - \n* `America/Godthab` - \n* `America/Goose_Bay` - \n* `America/Grand_Turk` - \n* `America/Grenada` - \n* `America/Guadeloupe` - \n* `America/Guatemala` - \n* `America/Guayaquil` - \n* `America/Guyana` - \n* `America/Halifax` - \n* `America/Havana` - \n* `America/Hermosillo` - \n* `America/Indiana/Indianapolis` - \n* `America/Indiana/Knox` - \n* `America/Indiana/Marengo` - \n* `America/Indiana/Petersburg` - \n* `America/Indiana/Tell_City` - \n* `America/Indiana/Vevay` - \n* `America/Indiana/Vincennes` - \n* `America/Indiana/Winamac` - \n* `America/Inuvik` - \n* `America/Iqaluit` - \n* `America/Jamaica` - \n* `America/Juneau` - \n* `America/Kentucky/Louisville` - \n* `America/Kentucky/Monticello` - \n* `America/Kralendijk` - \n* `America/La_Paz` - \n* `America/Lima` - \n* `America/Los_Angeles` - \n* `America/Lower_Princes` - \n* `America/Maceio` - \n* `America/Managua` - \n* `America/Manaus` - \n* `America/Marigot` - \n* `America/Martinique` - \n* `America/Matamoros` - \n* `America/Mazatlan` - \n* `America/Menominee` - \n* `America/Merida` - \n* `America/Metlakatla` - \n* `America/Mexico_City` - \n* `America/Miquelon` - \n* `America/Moncton` - \n* `America/Monterrey` - \n* `America/Montevideo` - \n* `America/Montreal` - \n* `America/Montserrat` - \n* `America/Nassau` - \n* `America/New_York` - \n* `America/Nipigon` - \n* `America/Nome` - \n* `America/Noronha` - \n* `America/North_Dakota/Beulah` - \n* `America/North_Dakota/Center` - \n* `America/North_Dakota/New_Salem` - \n* `America/Ojinaga` - \n* `America/Panama` - \n* `America/Pangnirtung` - \n* `America/Paramaribo` - \n* `America/Phoenix` - \n* `America/Port-au-Prince` - \n* `America/Port_of_Spain` - \n* `America/Porto_Velho` - \n* `America/Puerto_Rico` - \n* `America/Rainy_River` - \n* `America/Rankin_Inlet` - \n* `America/Recife` - \n* `America/Regina` - \n* `America/Resolute` - \n* `America/Rio_Branco` - \n* `America/Santa_Isabel` - \n* `America/Santarem` - \n* `America/Santiago` - \n* `America/Santo_Domingo` - \n* `America/Sao_Paulo` - \n* `America/Scoresbysund` - \n* `America/Shiprock` - \n* `America/Sitka` - \n* `America/St_Barthelemy` - \n* `America/St_Johns` - \n* `America/St_Kitts` - \n* `America/St_Lucia` - \n* `America/St_Thomas` - \n* `America/St_Vincent` - \n* `America/Swift_Current` - \n* `America/Tegucigalpa` - \n* `America/Thule` - \n* `America/Thunder_Bay` - \n* `America/Tijuana` - \n* `America/Toronto` - \n* `America/Tortola` - \n* `America/Vancouver` - \n* `America/Whitehorse` - \n* `America/Winnipeg` - \n* `America/Yakutat` - \n* `America/Yellowknife` - \n* `Antarctica/Casey` - \n* `Antarctica/Davis` - \n* `Antarctica/DumontDUrville` - \n* `Antarctica/Macquarie` - \n* `Antarctica/Mawson` - \n* `Antarctica/McMurdo` - \n* `Antarctica/Palmer` - \n* `Antarctica/Rothera` - \n* `Antarctica/South_Pole` - \n* `Antarctica/Syowa` - \n* `Antarctica/Troll` - \n* `Antarctica/Vostok` - \n* `Arctic/Longyearbyen` - \n* `Asia/Aden` - \n* `Asia/Almaty` - \n* `Asia/Amman` - \n* `Asia/Anadyr` - \n* `Asia/Aqtau` - \n* `Asia/Aqtobe` - \n* `Asia/Ashgabat` - \n* `Asia/Baghdad` - \n* `Asia/Bahrain` - \n* `Asia/Baku` - \n* `Asia/Bangkok` - \n* `Asia/Beirut` - \n* `Asia/Bishkek` - \n* `Asia/Brunei` - \n* `Asia/Calcutta` - \n* `Asia/Choibalsan` - \n* `Asia/Chongqing` - \n* `Asia/Colombo` - \n* `Asia/Damascus` - \n* `Asia/Dhaka` - \n* `Asia/Dili` - \n* `Asia/Dubai` - \n* `Asia/Dushanbe` - \n* `Asia/Gaza` - \n* `Asia/Harbin` - \n* `Asia/Hebron` - \n* `Asia/Ho_Chi_Minh` - \n* `Asia/Hong_Kong` - \n* `Asia/Hovd` - \n* `Asia/Irkutsk` - \n* `Asia/Jakarta` - \n* `Asia/Jayapura` - \n* `Asia/Jerusalem` - \n* `Asia/Kabul` - \n* `Asia/Kamchatka` - \n* `Asia/Karachi` - \n* `Asia/Kashgar` - \n* `Asia/Kathmandu` - \n* `Asia/Katmandu` - \n* `Asia/Khandyga` - \n* `Asia/Kolkata` - \n* `Asia/Krasnoyarsk` - \n* `Asia/Kuala_Lumpur` - \n* `Asia/Kuching` - \n* `Asia/Kuwait` - \n* `Asia/Macau` - \n* `Asia/Magadan` - \n* `Asia/Makassar` - \n* `Asia/Manila` - \n* `Asia/Muscat` - \n* `Asia/Nicosia` - \n* `Asia/Novokuznetsk` - \n* `Asia/Novosibirsk` - \n* `Asia/Omsk` - \n* `Asia/Oral` - \n* `Asia/Phnom_Penh` - \n* `Asia/Pontianak` - \n* `Asia/Pyongyang` - \n* `Asia/Qatar` - \n* `Asia/Qyzylorda` - \n* `Asia/Rangoon` - \n* `Asia/Riyadh` - \n* `Asia/Saigon` - \n* `Asia/Sakhalin` - \n* `Asia/Samarkand` - \n* `Asia/Seoul` - \n* `Asia/Shanghai` - \n* `Asia/Singapore` - \n* `Asia/Taipei` - \n* `Asia/Tashkent` - \n* `Asia/Tbilisi` - \n* `Asia/Tehran` - \n* `Asia/Thimphu` - \n* `Asia/Tokyo` - \n* `Asia/Ulaanbaatar` - \n* `Asia/Urumqi` - \n* `Asia/Ust-Nera` - \n* `Asia/Vientiane` - \n* `Asia/Vladivostok` - \n* `Asia/Yakutsk` - \n* `Asia/Yekaterinburg` - \n* `Asia/Yerevan` - \n* `Atlantic/Azores` - \n* `Atlantic/Bermuda` - \n* `Atlantic/Canary` - \n* `Atlantic/Cape_Verde` - \n* `Atlantic/Faroe` - \n* `Atlantic/Madeira` - \n* `Atlantic/Reykjavik` - \n* `Atlantic/South_Georgia` - \n* `Atlantic/St_Helena` - \n* `Atlantic/Stanley` - \n* `Australia/Adelaide` - \n* `Australia/Brisbane` - \n* `Australia/Broken_Hill` - \n* `Australia/Currie` - \n* `Australia/Darwin` - \n* `Australia/Eucla` - \n* `Australia/Hobart` - \n* `Australia/Lindeman` - \n* `Australia/Lord_Howe` - \n* `Australia/Melbourne` - \n* `Australia/Perth` - \n* `Australia/Sydney` - \n* `Etc/GMT` - \n* `Europe/Amsterdam` - \n* `Europe/Andorra` - \n* `Europe/Athens` - \n* `Europe/Belgrade` - \n* `Europe/Berlin` - \n* `Europe/Bratislava` - \n* `Europe/Brussels` - \n* `Europe/Bucharest` - \n* `Europe/Budapest` - \n* `Europe/Busingen` - \n* `Europe/Chisinau` - \n* `Europe/Copenhagen` - \n* `Europe/Dublin` - \n* `Europe/Gibraltar` - \n* `Europe/Guernsey` - \n* `Europe/Helsinki` - \n* `Europe/Isle_of_Man` - \n* `Europe/Istanbul` - \n* `Europe/Jersey` - \n* `Europe/Kaliningrad` - \n* `Europe/Kiev` - \n* `Europe/Lisbon` - \n* `Europe/Ljubljana` - \n* `Europe/London` - \n* `Europe/Luxembourg` - \n* `Europe/Madrid` - \n* `Europe/Malta` - \n* `Europe/Mariehamn` - \n* `Europe/Minsk` - \n* `Europe/Monaco` - \n* `Europe/Moscow` - \n* `Europe/Oslo` - \n* `Europe/Paris` - \n* `Europe/Podgorica` - \n* `Europe/Prague` - \n* `Europe/Riga` - \n* `Europe/Rome` - \n* `Europe/Samara` - \n* `Europe/San_Marino` - \n* `Europe/Sarajevo` - \n* `Europe/Simferopol` - \n* `Europe/Skopje` - \n* `Europe/Sofia` - \n* `Europe/Stockholm` - \n* `Europe/Tallinn` - \n* `Europe/Tirane` - \n* `Europe/Uzhgorod` - \n* `Europe/Vaduz` - \n* `Europe/Vatican` - \n* `Europe/Vienna` - \n* `Europe/Vilnius` - \n* `Europe/Volgograd` - \n* `Europe/Warsaw` - \n* `Europe/Zagreb` - \n* `Europe/Zaporozhye` - \n* `Europe/Zurich` - \n* `Indian/Antananarivo` - \n* `Indian/Chagos` - \n* `Indian/Christmas` - \n* `Indian/Cocos` - \n* `Indian/Comoro` - \n* `Indian/Kerguelen` - \n* `Indian/Mahe` - \n* `Indian/Maldives` - \n* `Indian/Mauritius` - \n* `Indian/Mayotte` - \n* `Indian/Reunion` - \n* `Pacific/Apia` - \n* `Pacific/Auckland` - \n* `Pacific/Chatham` - \n* `Pacific/Chuuk` - \n* `Pacific/Easter` - \n* `Pacific/Efate` - \n* `Pacific/Enderbury` - \n* `Pacific/Fakaofo` - \n* `Pacific/Fiji` - \n* `Pacific/Funafuti` - \n* `Pacific/Galapagos` - \n* `Pacific/Gambier` - \n* `Pacific/Guadalcanal` - \n* `Pacific/Guam` - \n* `Pacific/Honolulu` - \n* `Pacific/Johnston` - \n* `Pacific/Kiritimati` - \n* `Pacific/Kosrae` - \n* `Pacific/Kwajalein` - \n* `Pacific/Majuro` - \n* `Pacific/Marquesas` - \n* `Pacific/Midway` - \n* `Pacific/Nauru` - \n* `Pacific/Norfolk` - \n* `Pacific/Noumea` - \n* `Pacific/Pago_Pago` - \n* `Pacific/Palau` - \n* `Pacific/Pitcairn` - \n* `Pacific/Pohnpei` - \n* `Pacific/Port_Moresby` - \n* `Pacific/Rarotonga` - \n* `Pacific/Saipan` - \n* `Pacific/Tahiti` - \n* `Pacific/Tarawa` - \n* `Pacific/Tongatapu` - \n* `Pacific/Wake` - \n* `Pacific/Wallis` - \n* `UTC` -",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"Pacific/Niue", "Africa/Abidjan", "Africa/Accra", "Africa/Addis_Ababa", "Africa/Algiers", "Africa/Asmara", "Africa/Bamako", "Africa/Bangui", "Africa/Banjul", "Africa/Bissau", "Africa/Blantyre", "Africa/Brazzaville", "Africa/Bujumbura", "Africa/Cairo", "Africa/Casablanca", "Africa/Ceuta", "Africa/Conakry", "Africa/Dakar", "Africa/Dar_es_Salaam", "Africa/Djibouti", "Africa/Douala", "Africa/El_Aaiun", "Africa/Freetown", "Africa/Gaborone", "Africa/Harare", "Africa/Johannesburg", "Africa/Juba", "Africa/Kampala", "Africa/Khartoum", "Africa/Kigali", "Africa/Kinshasa", "Africa/Lagos", "Africa/Libreville", "Africa/Lome", "Africa/Luanda", "Africa/Lubumbashi", "Africa/Lusaka", "Africa/Malabo", "Africa/Maputo", "Africa/Maseru", "Africa/Mbabane", "Africa/Mogadishu", "Africa/Monrovia", "Africa/Nairobi", "Africa/Ndjamena", "Africa/Niamey", "Africa/Nouakchott", "Africa/Ouagadougou", "Africa/Porto-Novo", "Africa/Sao_Tome", "Africa/Tripoli", "Africa/Tunis", "Africa/Windhoek", "America/Adak", "America/Anchorage", "America/Anguilla", "America/Antigua", "America/Araguaina", "America/Argentina/Buenos_Aires", "America/Argentina/Catamarca", "America/Argentina/Cordoba", "America/Argentina/Jujuy", "America/Argentina/La_Rioja", "America/Argentina/Mendoza", "America/Argentina/Rio_Gallegos", "America/Argentina/Salta", "America/Argentina/San_Juan", "America/Argentina/San_Luis", "America/Argentina/Tucuman", "America/Argentina/Ushuaia", "America/Aruba", "America/Asuncion", "America/Atikokan", "America/Bahia", "America/Bahia_Banderas", "America/Barbados", "America/Belem", "America/Belize", "America/Blanc-Sablon", "America/Boa_Vista", "America/Bogota", "America/Boise", "America/Cambridge_Bay", "America/Campo_Grande", "America/Cancun", "America/Caracas", "America/Cayenne", "America/Cayman", "America/Chicago", "America/Chihuahua", "America/Costa_Rica", "America/Creston", "America/Cuiaba", "America/Curacao", "America/Danmarkshavn", "America/Dawson", "America/Dawson_Creek", "America/Denver", "America/Detroit", "America/Dominica", "America/Edmonton", "America/Eirunepe", "America/El_Salvador", "America/Fortaleza", "America/Glace_Bay", "America/Godthab", "America/Goose_Bay", "America/Grand_Turk", "America/Grenada", "America/Guadeloupe", "America/Guatemala", "America/Guayaquil", "America/Guyana", "America/Halifax", "America/Havana", "America/Hermosillo", "America/Indiana/Indianapolis", "America/Indiana/Knox", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Tell_City", "America/Indiana/Vevay", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Inuvik", "America/Iqaluit", "America/Jamaica", "America/Juneau", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Kralendijk", "America/La_Paz", "America/Lima", "America/Los_Angeles", "America/Lower_Princes", "America/Maceio", "America/Managua", "America/Manaus", "America/Marigot", "America/Martinique", "America/Matamoros", "America/Mazatlan", "America/Menominee", "America/Merida", "America/Metlakatla", "America/Mexico_City", "America/Miquelon", "America/Moncton", "America/Monterrey", "America/Montevideo", "America/Montreal", "America/Montserrat", "America/Nassau", "America/New_York", "America/Nipigon", "America/Nome", "America/Noronha", "America/North_Dakota/Beulah", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/Ojinaga", "America/Panama", "America/Pangnirtung", "America/Paramaribo", "America/Phoenix", "America/Port-au-Prince", "America/Port_of_Spain", "America/Porto_Velho", "America/Puerto_Rico", "America/Rainy_River", "America/Rankin_Inlet", "America/Recife", "America/Regina", "America/Resolute", "America/Rio_Branco", "America/Santa_Isabel", "America/Santarem", "America/Santiago", "America/Santo_Domingo", "America/Sao_Paulo", "America/Scoresbysund", "America/Shiprock", "America/Sitka", "America/St_Barthelemy", "America/St_Johns", "America/St_Kitts", "America/St_Lucia", "America/St_Thomas", "America/St_Vincent", "America/Swift_Current", "America/Tegucigalpa", "America/Thule", "America/Thunder_Bay", "America/Tijuana", "America/Toronto", "America/Tortola", "America/Vancouver", "America/Whitehorse", "America/Winnipeg", "America/Yakutat", "America/Yellowknife", "Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Macquarie", "Antarctica/Mawson", "Antarctica/McMurdo", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/South_Pole", "Antarctica/Syowa", "Antarctica/Troll", "Antarctica/Vostok", "Arctic/Longyearbyen", "Asia/Aden", "Asia/Almaty", "Asia/Amman", "Asia/Anadyr", "Asia/Aqtau", "Asia/Aqtobe", "Asia/Ashgabat", "Asia/Baghdad", "Asia/Bahrain", "Asia/Baku", "Asia/Bangkok", "Asia/Beirut", "Asia/Bishkek", "Asia/Brunei", "Asia/Calcutta", "Asia/Choibalsan", "Asia/Chongqing", "Asia/Colombo", "Asia/Damascus", "Asia/Dhaka", "Asia/Dili", "Asia/Dubai", "Asia/Dushanbe", "Asia/Gaza", "Asia/Harbin", "Asia/Hebron", "Asia/Ho_Chi_Minh", "Asia/Hong_Kong", "Asia/Hovd", "Asia/Irkutsk", "Asia/Jakarta", "Asia/Jayapura", "Asia/Jerusalem", "Asia/Kabul", "Asia/Kamchatka", "Asia/Karachi", "Asia/Kashgar", "Asia/Kathmandu", "Asia/Katmandu", "Asia/Khandyga", "Asia/Kolkata", "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Kuching", "Asia/Kuwait", "Asia/Macau", "Asia/Magadan", "Asia/Makassar", "Asia/Manila", "Asia/Muscat", "Asia/Nicosia", "Asia/Novokuznetsk", "Asia/Novosibirsk", "Asia/Omsk", "Asia/Oral", "Asia/Phnom_Penh", "Asia/Pontianak", "Asia/Pyongyang", "Asia/Qatar", "Asia/Qyzylorda", "Asia/Rangoon", "Asia/Riyadh", "Asia/Saigon", "Asia/Sakhalin", "Asia/Samarkand", "Asia/Seoul", "Asia/Shanghai", "Asia/Singapore", "Asia/Taipei", "Asia/Tashkent", "Asia/Tbilisi", "Asia/Tehran", "Asia/Thimphu", "Asia/Tokyo", "Asia/Ulaanbaatar", "Asia/Urumqi", "Asia/Ust-Nera", "Asia/Vientiane", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yekaterinburg", "Asia/Yerevan", "Atlantic/Azores", "Atlantic/Bermuda", "Atlantic/Canary", "Atlantic/Cape_Verde", "Atlantic/Faroe", "Atlantic/Madeira", "Atlantic/Reykjavik", "Atlantic/South_Georgia", "Atlantic/St_Helena", "Atlantic/Stanley", "Australia/Adelaide", "Australia/Brisbane", "Australia/Broken_Hill", "Australia/Currie", "Australia/Darwin", "Australia/Eucla", "Australia/Hobart", "Australia/Lindeman", "Australia/Lord_Howe", "Australia/Melbourne", "Australia/Perth", "Australia/Sydney", "Etc/GMT", "Europe/Amsterdam", "Europe/Andorra", "Europe/Athens", "Europe/Belgrade", "Europe/Berlin", "Europe/Bratislava", "Europe/Brussels", "Europe/Bucharest", "Europe/Budapest", "Europe/Busingen", "Europe/Chisinau", "Europe/Copenhagen", "Europe/Dublin", "Europe/Gibraltar", "Europe/Guernsey", "Europe/Helsinki", "Europe/Isle_of_Man", "Europe/Istanbul", "Europe/Jersey", "Europe/Kaliningrad", "Europe/Kiev", "Europe/Lisbon", "Europe/Ljubljana", "Europe/London", "Europe/Luxembourg", "Europe/Madrid", "Europe/Malta", "Europe/Mariehamn", "Europe/Minsk", "Europe/Monaco", "Europe/Moscow", "Europe/Oslo", "Europe/Paris", "Europe/Podgorica", "Europe/Prague", "Europe/Riga", "Europe/Rome", "Europe/Samara", "Europe/San_Marino", "Europe/Sarajevo", "Europe/Simferopol", "Europe/Skopje", "Europe/Sofia", "Europe/Stockholm", "Europe/Tallinn", "Europe/Tirane", "Europe/Uzhgorod", "Europe/Vaduz", "Europe/Vatican", "Europe/Vienna", "Europe/Vilnius", "Europe/Volgograd", "Europe/Warsaw", "Europe/Zagreb", "Europe/Zaporozhye", "Europe/Zurich", "Indian/Antananarivo", "Indian/Chagos", "Indian/Christmas", "Indian/Cocos", "Indian/Comoro", "Indian/Kerguelen", "Indian/Mahe", "Indian/Maldives", "Indian/Mauritius", "Indian/Mayotte", "Indian/Reunion", "Pacific/Apia", "Pacific/Auckland", "Pacific/Chatham", "Pacific/Chuuk", "Pacific/Easter", "Pacific/Efate", "Pacific/Enderbury", "Pacific/Fakaofo", "Pacific/Fiji", "Pacific/Funafuti", "Pacific/Galapagos", "Pacific/Gambier", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Honolulu", "Pacific/Johnston", "Pacific/Kiritimati", "Pacific/Kosrae", "Pacific/Kwajalein", "Pacific/Majuro", "Pacific/Marquesas", "Pacific/Midway", "Pacific/Nauru", "Pacific/Norfolk", "Pacific/Noumea", "Pacific/Pago_Pago", "Pacific/Palau", "Pacific/Pitcairn", "Pacific/Pohnpei", "Pacific/Port_Moresby", "Pacific/Rarotonga", "Pacific/Saipan", "Pacific/Tahiti", "Pacific/Tarawa", "Pacific/Tongatapu", "Pacific/Wake", "Pacific/Wallis", "UTC"}, false),
							Optional:     true,
							Default:      "Pacific/Niue",
						},
						"week_of_month": {
							Description:  "Schedule a task on a specific week of the month. Valid values are 1 through 5. Value of 5 means last week of the month. WeekOfMonth may not be set when dayOfMonth is specified.",
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 5),
							Optional:     true,
						},
					},
				},
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"tags": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"ancestor_definitions": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "mo.MoRef",
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"selector": {
										Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"definition": {
							Description: "The definition is a reference to the tag definition object.\nThe tag definition object contains the properties of the tag such as name, type, and description.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							ConfigMode:  schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "mo.MoRef",
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"selector": {
										Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"key": {
							Description:  "The string representation of a tag key.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(1, 256),
							Optional:     true,
						},
						"propagated": {
							Description: "Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"type": {
							Description: "An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.\n* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.\n* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \"A/B/C\", then \"A\" is the parent tag, \"B\" is the child tag of \"A\" and \"C\" is the child tag of \"B\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"value": {
							Description:  "The string representation of a tag value.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 256),
							Optional:     true,
						},
					},
				},
			},
			"username": {
				Description:  "Username to authenticate the fileserver.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^$|^[a-zA-Z0-9_][a-zA-Z0-9_\\.\\@\\\\\\-\\+]*$"), ""),
				Optional:     true,
			},
			"version_context": {
				Description: "The versioning info for this managed object.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
						},
						"interested_mos": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "mo.MoRef",
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"selector": {
										Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"marked_for_deletion": {
							Description: "The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
						},
						"ref_mo": {
							Description: "A reference to the original Managed Object.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							ConfigMode:  schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "mo.MoRef",
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"selector": {
										Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"timestamp": {
							Description: "The time this versioned Managed Object was created.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"nr_version": {
							Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"version_type": {
							Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
					},
				},
			},
		},
	}
}

func resourceApplianceBackupPolicyCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewApplianceBackupPolicyWithDefaults()
	if v, ok := d.GetOk("account"); ok {
		p := make([]models.IamAccountRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsIamAccountRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAccount(x)
		}
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("appliance.BackupPolicy")

	if v, ok := d.GetOk("filename"); ok {
		x := (v.(string))
		o.SetFilename(x)
	}

	if v, ok := d.GetOkExists("manual_backup"); ok {
		x := (v.(bool))
		o.SetManualBackup(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("appliance.BackupPolicy")

	if v, ok := d.GetOk("password"); ok {
		x := (v.(string))
		o.SetPassword(x)
	}

	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}

	if v, ok := d.GetOk("remote_host"); ok {
		x := (v.(string))
		o.SetRemoteHost(x)
	}

	if v, ok := d.GetOk("remote_path"); ok {
		x := (v.(string))
		o.SetRemotePath(x)
	}

	if v, ok := d.GetOkExists("remote_port"); ok {
		x := int64(v.(int))
		o.SetRemotePort(x)
	}

	if v, ok := d.GetOkExists("retention_count"); ok {
		x := int64(v.(int))
		o.SetRetentionCount(x)
	}

	if v, ok := d.GetOkExists("retention_policy_enabled"); ok {
		x := (v.(bool))
		o.SetRetentionPolicyEnabled(x)
	}

	if v, ok := d.GetOk("schedule"); ok {
		p := make([]models.OnpremSchedule, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewOnpremScheduleWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("onprem.Schedule")
			if v, ok := l["day_of_month"]; ok {
				{
					x := int64(v.(int))
					o.SetDayOfMonth(x)
				}
			}
			if v, ok := l["day_of_week"]; ok {
				{
					x := int64(v.(int))
					o.SetDayOfWeek(x)
				}
			}
			if v, ok := l["month_of_year"]; ok {
				{
					x := int64(v.(int))
					o.SetMonthOfYear(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["repeat_interval"]; ok {
				{
					x := int64(v.(int))
					o.SetRepeatInterval(x)
				}
			}
			if v, ok := l["time_of_day"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeOfDay(x)
				}
			}
			if v, ok := l["time_zone"]; ok {
				{
					x := (v.(string))
					o.SetTimeZone(x)
				}
			}
			if v, ok := l["week_of_month"]; ok {
				{
					x := int64(v.(int))
					o.SetWeekOfMonth(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSchedule(x)
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["ancestor_definitions"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetAncestorDefinitions(x)
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	r := conn.ApiClient.ApplianceApi.CreateApplianceBackupPolicy(conn.ctx).ApplianceBackupPolicy(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating ApplianceBackupPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating ApplianceBackupPolicy: %s", responseErr.Error())
	}
	if len(resultMo.GetMoid()) != 0 {
		log.Printf("Moid: %s", resultMo.GetMoid())
		d.SetId(resultMo.GetMoid())
	} else {
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		log.Printf("Mo: %v", resultMo)
	}
	if len(resultMo.GetMoid()) == 0 {
		return de
	}
	return append(de, resourceApplianceBackupPolicyRead(c, d, meta)...)
}

func resourceApplianceBackupPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	if len(d.Id()) == 0 {
		return de
	}
	conn := meta.(*Config)
	r := conn.ApiClient.ApplianceApi.GetApplianceBackupPolicyByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "ApplianceBackupPolicy object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ApplianceBackupPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching ApplianceBackupPolicy: %s", responseErr.Error())
	}

	if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Account in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("backup_time", (s.GetBackupTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property BackupTime in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("filename", (s.GetFilename())); err != nil {
		return diag.Errorf("error occurred while setting property Filename in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("is_password_set", (s.GetIsPasswordSet())); err != nil {
		return diag.Errorf("error occurred while setting property IsPasswordSet in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("manual_backup", (s.GetManualBackup())); err != nil {
		return diag.Errorf("error occurred while setting property ManualBackup in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("protocol", (s.GetProtocol())); err != nil {
		return diag.Errorf("error occurred while setting property Protocol in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("remote_host", (s.GetRemoteHost())); err != nil {
		return diag.Errorf("error occurred while setting property RemoteHost in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("remote_path", (s.GetRemotePath())); err != nil {
		return diag.Errorf("error occurred while setting property RemotePath in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("remote_port", (s.GetRemotePort())); err != nil {
		return diag.Errorf("error occurred while setting property RemotePort in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("retention_count", (s.GetRetentionCount())); err != nil {
		return diag.Errorf("error occurred while setting property RetentionCount in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("retention_policy_enabled", (s.GetRetentionPolicyEnabled())); err != nil {
		return diag.Errorf("error occurred while setting property RetentionPolicyEnabled in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("schedule", flattenMapOnpremSchedule(s.GetSchedule(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Schedule in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("username", (s.GetUsername())); err != nil {
		return diag.Errorf("error occurred while setting property Username in ApplianceBackupPolicy object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in ApplianceBackupPolicy object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceApplianceBackupPolicyUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceBackupPolicy{}
	if d.HasChange("account") {
		v := d.Get("account")
		p := make([]models.IamAccountRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsIamAccountRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAccount(x)
		}
	}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("appliance.BackupPolicy")

	if d.HasChange("filename") {
		v := d.Get("filename")
		x := (v.(string))
		o.SetFilename(x)
	}

	if d.HasChange("manual_backup") {
		v := d.Get("manual_backup")
		x := (v.(bool))
		o.SetManualBackup(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("appliance.BackupPolicy")

	if d.HasChange("password") {
		v := d.Get("password")
		x := (v.(string))
		o.SetPassword(x)
	}

	if d.HasChange("protocol") {
		v := d.Get("protocol")
		x := (v.(string))
		o.SetProtocol(x)
	}

	if d.HasChange("remote_host") {
		v := d.Get("remote_host")
		x := (v.(string))
		o.SetRemoteHost(x)
	}

	if d.HasChange("remote_path") {
		v := d.Get("remote_path")
		x := (v.(string))
		o.SetRemotePath(x)
	}

	if d.HasChange("remote_port") {
		v := d.Get("remote_port")
		x := int64(v.(int))
		o.SetRemotePort(x)
	}

	if d.HasChange("retention_count") {
		v := d.Get("retention_count")
		x := int64(v.(int))
		o.SetRetentionCount(x)
	}

	if d.HasChange("retention_policy_enabled") {
		v := d.Get("retention_policy_enabled")
		x := (v.(bool))
		o.SetRetentionPolicyEnabled(x)
	}

	if d.HasChange("schedule") {
		v := d.Get("schedule")
		p := make([]models.OnpremSchedule, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.OnpremSchedule{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("onprem.Schedule")
			if v, ok := l["day_of_month"]; ok {
				{
					x := int64(v.(int))
					o.SetDayOfMonth(x)
				}
			}
			if v, ok := l["day_of_week"]; ok {
				{
					x := int64(v.(int))
					o.SetDayOfWeek(x)
				}
			}
			if v, ok := l["month_of_year"]; ok {
				{
					x := int64(v.(int))
					o.SetMonthOfYear(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["repeat_interval"]; ok {
				{
					x := int64(v.(int))
					o.SetRepeatInterval(x)
				}
			}
			if v, ok := l["time_of_day"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeOfDay(x)
				}
			}
			if v, ok := l["time_zone"]; ok {
				{
					x := (v.(string))
					o.SetTimeZone(x)
				}
			}
			if v, ok := l["week_of_month"]; ok {
				{
					x := int64(v.(int))
					o.SetWeekOfMonth(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSchedule(x)
		}
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["ancestor_definitions"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetAncestorDefinitions(x)
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if d.HasChange("username") {
		v := d.Get("username")
		x := (v.(string))
		o.SetUsername(x)
	}

	r := conn.ApiClient.ApplianceApi.UpdateApplianceBackupPolicy(conn.ctx, d.Id()).ApplianceBackupPolicy(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating ApplianceBackupPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating ApplianceBackupPolicy: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceApplianceBackupPolicyRead(c, d, meta)...)
}

func resourceApplianceBackupPolicyDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	var warning = diag.Diagnostic{Severity: diag.Warning, Summary: "ApplianceBackupPolicy does not allow delete functionality"}
	de = append(de, warning)
	return de
}
