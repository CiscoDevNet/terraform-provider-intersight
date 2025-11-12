package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getSchedulerSchedulePolicySchema() map[string]*schema.Schema {
	var schemaMap = make(map[string]*schema.Schema)
	schemaMap = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
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
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"associated_objects": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
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
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"block_dates": {
			Type:     schema.TypeList,
			Optional: true,
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
					},
					"end_time": {
						Description: "The end date time of the block date.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_time": {
						Description: "The start date time of the block date.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
		},
		"description": {
			Description: "Description of the policy.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"enable_block_dates": {
			Description: "Enable or disable block dates. If set to true, the schedule will not run during the block date interval.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"name": {
			Description: "Name of the concrete policy.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"organization": {
			Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
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
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
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
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
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
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"schedule_params": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"duration": {
						Description: "The duration of the schedule. Its syntax is specified at https://www.w3.org/TR/xmlschema11-2/#nt-durationRep For example, P20DT10H5M2.3S is for 20 days, 10 hours, 5 minutes and 2.3 seconds. It is a mandatory input property for Policy based schedules.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "The name of the schedule. It is a mandatory input property for Policy based schedules.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_time": {
						Description: "The schedule start time. A future time is required.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"time_zone": {
						Description: "The timezone for the startTime specified. It is a mandatory input property when start time is provided.\n* `Pacific/Niue` - \n* `Africa/Abidjan` - \n* `Africa/Accra` - \n* `Africa/Addis_Ababa` - \n* `Africa/Algiers` - \n* `Africa/Asmara` - \n* `Africa/Bamako` - \n* `Africa/Bangui` - \n* `Africa/Banjul` - \n* `Africa/Bissau` - \n* `Africa/Blantyre` - \n* `Africa/Brazzaville` - \n* `Africa/Bujumbura` - \n* `Africa/Cairo` - \n* `Africa/Casablanca` - \n* `Africa/Ceuta` - \n* `Africa/Conakry` - \n* `Africa/Dakar` - \n* `Africa/Dar_es_Salaam` - \n* `Africa/Djibouti` - \n* `Africa/Douala` - \n* `Africa/El_Aaiun` - \n* `Africa/Freetown` - \n* `Africa/Gaborone` - \n* `Africa/Harare` - \n* `Africa/Johannesburg` - \n* `Africa/Juba` - \n* `Africa/Kampala` - \n* `Africa/Khartoum` - \n* `Africa/Kigali` - \n* `Africa/Kinshasa` - \n* `Africa/Lagos` - \n* `Africa/Libreville` - \n* `Africa/Lome` - \n* `Africa/Luanda` - \n* `Africa/Lubumbashi` - \n* `Africa/Lusaka` - \n* `Africa/Malabo` - \n* `Africa/Maputo` - \n* `Africa/Maseru` - \n* `Africa/Mbabane` - \n* `Africa/Mogadishu` - \n* `Africa/Monrovia` - \n* `Africa/Nairobi` - \n* `Africa/Ndjamena` - \n* `Africa/Niamey` - \n* `Africa/Nouakchott` - \n* `Africa/Ouagadougou` - \n* `Africa/Porto-Novo` - \n* `Africa/Sao_Tome` - \n* `Africa/Tripoli` - \n* `Africa/Tunis` - \n* `Africa/Windhoek` - \n* `America/Adak` - \n* `America/Anchorage` - \n* `America/Anguilla` - \n* `America/Antigua` - \n* `America/Araguaina` - \n* `America/Argentina/Buenos_Aires` - \n* `America/Argentina/Catamarca` - \n* `America/Argentina/Cordoba` - \n* `America/Argentina/Jujuy` - \n* `America/Argentina/La_Rioja` - \n* `America/Argentina/Mendoza` - \n* `America/Argentina/Rio_Gallegos` - \n* `America/Argentina/Salta` - \n* `America/Argentina/San_Juan` - \n* `America/Argentina/San_Luis` - \n* `America/Argentina/Tucuman` - \n* `America/Argentina/Ushuaia` - \n* `America/Aruba` - \n* `America/Asuncion` - \n* `America/Atikokan` - \n* `America/Bahia` - \n* `America/Bahia_Banderas` - \n* `America/Barbados` - \n* `America/Belem` - \n* `America/Belize` - \n* `America/Blanc-Sablon` - \n* `America/Boa_Vista` - \n* `America/Bogota` - \n* `America/Boise` - \n* `America/Cambridge_Bay` - \n* `America/Campo_Grande` - \n* `America/Cancun` - \n* `America/Caracas` - \n* `America/Cayenne` - \n* `America/Cayman` - \n* `America/Chicago` - \n* `America/Chihuahua` - \n* `America/Costa_Rica` - \n* `America/Creston` - \n* `America/Cuiaba` - \n* `America/Curacao` - \n* `America/Danmarkshavn` - \n* `America/Dawson` - \n* `America/Dawson_Creek` - \n* `America/Denver` - \n* `America/Detroit` - \n* `America/Dominica` - \n* `America/Edmonton` - \n* `America/Eirunepe` - \n* `America/El_Salvador` - \n* `America/Fortaleza` - \n* `America/Glace_Bay` - \n* `America/Godthab` - \n* `America/Goose_Bay` - \n* `America/Grand_Turk` - \n* `America/Grenada` - \n* `America/Guadeloupe` - \n* `America/Guatemala` - \n* `America/Guayaquil` - \n* `America/Guyana` - \n* `America/Halifax` - \n* `America/Havana` - \n* `America/Hermosillo` - \n* `America/Indiana/Indianapolis` - \n* `America/Indiana/Knox` - \n* `America/Indiana/Marengo` - \n* `America/Indiana/Petersburg` - \n* `America/Indiana/Tell_City` - \n* `America/Indiana/Vevay` - \n* `America/Indiana/Vincennes` - \n* `America/Indiana/Winamac` - \n* `America/Inuvik` - \n* `America/Iqaluit` - \n* `America/Jamaica` - \n* `America/Juneau` - \n* `America/Kentucky/Louisville` - \n* `America/Kentucky/Monticello` - \n* `America/Kralendijk` - \n* `America/La_Paz` - \n* `America/Lima` - \n* `America/Los_Angeles` - \n* `America/Lower_Princes` - \n* `America/Maceio` - \n* `America/Managua` - \n* `America/Manaus` - \n* `America/Marigot` - \n* `America/Martinique` - \n* `America/Matamoros` - \n* `America/Mazatlan` - \n* `America/Menominee` - \n* `America/Merida` - \n* `America/Metlakatla` - \n* `America/Mexico_City` - \n* `America/Miquelon` - \n* `America/Moncton` - \n* `America/Monterrey` - \n* `America/Montevideo` - \n* `America/Montreal` - \n* `America/Montserrat` - \n* `America/Nassau` - \n* `America/New_York` - \n* `America/Nipigon` - \n* `America/Nome` - \n* `America/Noronha` - \n* `America/North_Dakota/Beulah` - \n* `America/North_Dakota/Center` - \n* `America/North_Dakota/New_Salem` - \n* `America/Ojinaga` - \n* `America/Panama` - \n* `America/Pangnirtung` - \n* `America/Paramaribo` - \n* `America/Phoenix` - \n* `America/Port-au-Prince` - \n* `America/Port_of_Spain` - \n* `America/Porto_Velho` - \n* `America/Puerto_Rico` - \n* `America/Rainy_River` - \n* `America/Rankin_Inlet` - \n* `America/Recife` - \n* `America/Regina` - \n* `America/Resolute` - \n* `America/Rio_Branco` - \n* `America/Santa_Isabel` - \n* `America/Santarem` - \n* `America/Santiago` - \n* `America/Santo_Domingo` - \n* `America/Sao_Paulo` - \n* `America/Scoresbysund` - \n* `America/Shiprock` - \n* `America/Sitka` - \n* `America/St_Barthelemy` - \n* `America/St_Johns` - \n* `America/St_Kitts` - \n* `America/St_Lucia` - \n* `America/St_Thomas` - \n* `America/St_Vincent` - \n* `America/Swift_Current` - \n* `America/Tegucigalpa` - \n* `America/Thule` - \n* `America/Thunder_Bay` - \n* `America/Tijuana` - \n* `America/Toronto` - \n* `America/Tortola` - \n* `America/Vancouver` - \n* `America/Whitehorse` - \n* `America/Winnipeg` - \n* `America/Yakutat` - \n* `America/Yellowknife` - \n* `Antarctica/Casey` - \n* `Antarctica/Davis` - \n* `Antarctica/DumontDUrville` - \n* `Antarctica/Macquarie` - \n* `Antarctica/Mawson` - \n* `Antarctica/McMurdo` - \n* `Antarctica/Palmer` - \n* `Antarctica/Rothera` - \n* `Antarctica/South_Pole` - \n* `Antarctica/Syowa` - \n* `Antarctica/Troll` - \n* `Antarctica/Vostok` - \n* `Arctic/Longyearbyen` - \n* `Asia/Aden` - \n* `Asia/Almaty` - \n* `Asia/Amman` - \n* `Asia/Anadyr` - \n* `Asia/Aqtau` - \n* `Asia/Aqtobe` - \n* `Asia/Ashgabat` - \n* `Asia/Baghdad` - \n* `Asia/Bahrain` - \n* `Asia/Baku` - \n* `Asia/Bangkok` - \n* `Asia/Beirut` - \n* `Asia/Bishkek` - \n* `Asia/Brunei` - \n* `Asia/Calcutta` - \n* `Asia/Choibalsan` - \n* `Asia/Chongqing` - \n* `Asia/Colombo` - \n* `Asia/Damascus` - \n* `Asia/Dhaka` - \n* `Asia/Dili` - \n* `Asia/Dubai` - \n* `Asia/Dushanbe` - \n* `Asia/Gaza` - \n* `Asia/Harbin` - \n* `Asia/Hebron` - \n* `Asia/Ho_Chi_Minh` - \n* `Asia/Hong_Kong` - \n* `Asia/Hovd` - \n* `Asia/Irkutsk` - \n* `Asia/Jakarta` - \n* `Asia/Jayapura` - \n* `Asia/Jerusalem` - \n* `Asia/Kabul` - \n* `Asia/Kamchatka` - \n* `Asia/Karachi` - \n* `Asia/Kashgar` - \n* `Asia/Kathmandu` - \n* `Asia/Katmandu` - \n* `Asia/Khandyga` - \n* `Asia/Kolkata` - \n* `Asia/Krasnoyarsk` - \n* `Asia/Kuala_Lumpur` - \n* `Asia/Kuching` - \n* `Asia/Kuwait` - \n* `Asia/Macau` - \n* `Asia/Magadan` - \n* `Asia/Makassar` - \n* `Asia/Manila` - \n* `Asia/Muscat` - \n* `Asia/Nicosia` - \n* `Asia/Novokuznetsk` - \n* `Asia/Novosibirsk` - \n* `Asia/Omsk` - \n* `Asia/Oral` - \n* `Asia/Phnom_Penh` - \n* `Asia/Pontianak` - \n* `Asia/Pyongyang` - \n* `Asia/Qatar` - \n* `Asia/Qyzylorda` - \n* `Asia/Rangoon` - \n* `Asia/Riyadh` - \n* `Asia/Saigon` - \n* `Asia/Sakhalin` - \n* `Asia/Samarkand` - \n* `Asia/Seoul` - \n* `Asia/Shanghai` - \n* `Asia/Singapore` - \n* `Asia/Taipei` - \n* `Asia/Tashkent` - \n* `Asia/Tbilisi` - \n* `Asia/Tehran` - \n* `Asia/Thimphu` - \n* `Asia/Tokyo` - \n* `Asia/Ulaanbaatar` - \n* `Asia/Urumqi` - \n* `Asia/Ust-Nera` - \n* `Asia/Vientiane` - \n* `Asia/Vladivostok` - \n* `Asia/Yakutsk` - \n* `Asia/Yekaterinburg` - \n* `Asia/Yerevan` - \n* `Atlantic/Azores` - \n* `Atlantic/Bermuda` - \n* `Atlantic/Canary` - \n* `Atlantic/Cape_Verde` - \n* `Atlantic/Faroe` - \n* `Atlantic/Madeira` - \n* `Atlantic/Reykjavik` - \n* `Atlantic/South_Georgia` - \n* `Atlantic/St_Helena` - \n* `Atlantic/Stanley` - \n* `Australia/Adelaide` - \n* `Australia/Brisbane` - \n* `Australia/Broken_Hill` - \n* `Australia/Currie` - \n* `Australia/Darwin` - \n* `Australia/Eucla` - \n* `Australia/Hobart` - \n* `Australia/Lindeman` - \n* `Australia/Lord_Howe` - \n* `Australia/Melbourne` - \n* `Australia/Perth` - \n* `Australia/Sydney` - \n* `Etc/GMT` - \n* `Europe/Amsterdam` - \n* `Europe/Andorra` - \n* `Europe/Athens` - \n* `Europe/Belgrade` - \n* `Europe/Berlin` - \n* `Europe/Bratislava` - \n* `Europe/Brussels` - \n* `Europe/Bucharest` - \n* `Europe/Budapest` - \n* `Europe/Busingen` - \n* `Europe/Chisinau` - \n* `Europe/Copenhagen` - \n* `Europe/Dublin` - \n* `Europe/Gibraltar` - \n* `Europe/Guernsey` - \n* `Europe/Helsinki` - \n* `Europe/Isle_of_Man` - \n* `Europe/Istanbul` - \n* `Europe/Jersey` - \n* `Europe/Kaliningrad` - \n* `Europe/Kiev` - \n* `Europe/Lisbon` - \n* `Europe/Ljubljana` - \n* `Europe/London` - \n* `Europe/Luxembourg` - \n* `Europe/Madrid` - \n* `Europe/Malta` - \n* `Europe/Mariehamn` - \n* `Europe/Minsk` - \n* `Europe/Monaco` - \n* `Europe/Moscow` - \n* `Europe/Oslo` - \n* `Europe/Paris` - \n* `Europe/Podgorica` - \n* `Europe/Prague` - \n* `Europe/Riga` - \n* `Europe/Rome` - \n* `Europe/Samara` - \n* `Europe/San_Marino` - \n* `Europe/Sarajevo` - \n* `Europe/Simferopol` - \n* `Europe/Skopje` - \n* `Europe/Sofia` - \n* `Europe/Stockholm` - \n* `Europe/Tallinn` - \n* `Europe/Tirane` - \n* `Europe/Uzhgorod` - \n* `Europe/Vaduz` - \n* `Europe/Vatican` - \n* `Europe/Vienna` - \n* `Europe/Vilnius` - \n* `Europe/Volgograd` - \n* `Europe/Warsaw` - \n* `Europe/Zagreb` - \n* `Europe/Zaporozhye` - \n* `Europe/Zurich` - \n* `Indian/Antananarivo` - \n* `Indian/Chagos` - \n* `Indian/Christmas` - \n* `Indian/Cocos` - \n* `Indian/Comoro` - \n* `Indian/Kerguelen` - \n* `Indian/Mahe` - \n* `Indian/Maldives` - \n* `Indian/Mauritius` - \n* `Indian/Mayotte` - \n* `Indian/Reunion` - \n* `Pacific/Apia` - \n* `Pacific/Auckland` - \n* `Pacific/Chatham` - \n* `Pacific/Chuuk` - \n* `Pacific/Easter` - \n* `Pacific/Efate` - \n* `Pacific/Enderbury` - \n* `Pacific/Fakaofo` - \n* `Pacific/Fiji` - \n* `Pacific/Funafuti` - \n* `Pacific/Galapagos` - \n* `Pacific/Gambier` - \n* `Pacific/Guadalcanal` - \n* `Pacific/Guam` - \n* `Pacific/Honolulu` - \n* `Pacific/Johnston` - \n* `Pacific/Kiritimati` - \n* `Pacific/Kosrae` - \n* `Pacific/Kwajalein` - \n* `Pacific/Majuro` - \n* `Pacific/Marquesas` - \n* `Pacific/Midway` - \n* `Pacific/Nauru` - \n* `Pacific/Norfolk` - \n* `Pacific/Noumea` - \n* `Pacific/Pago_Pago` - \n* `Pacific/Palau` - \n* `Pacific/Pitcairn` - \n* `Pacific/Pohnpei` - \n* `Pacific/Port_Moresby` - \n* `Pacific/Rarotonga` - \n* `Pacific/Saipan` - \n* `Pacific/Tahiti` - \n* `Pacific/Tarawa` - \n* `Pacific/Tongatapu` - \n* `Pacific/Wake` - \n* `Pacific/Wallis` - \n* `UTC` -",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"ancestor_definitions": {
						Type:     schema.TypeList,
						Optional: true,
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
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
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
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
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
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"propagated": {
						Description: "Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"type": {
						Description: "An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.\n* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.\n* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \"A/B/C\", then \"A\" is the parent tag, \"B\" is the child tag of \"A\" and \"C\" is the child tag of \"B\".",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"usage_count": {
			Description: "The number of profiles, templates and deployments that are using this policy. \nThis is used to determine if the policy can be deleted.\nIf the usageCount is greater than 0, the policy cannot be deleted.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
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
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
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
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
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
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
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
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
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
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
	return schemaMap
}

func dataSourceSchedulerSchedulePolicy() *schema.Resource {
	var subSchema = getSchedulerSchedulePolicySchema()
	var model = getSchedulerSchedulePolicySchema()
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceSchedulerSchedulePolicyRead,
		Schema:      model}
}

func dataSourceSchedulerSchedulePolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SchedulerSchedulePolicy{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("ancestors"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
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
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAncestors(x)
	}

	if v, ok := d.GetOk("associated_objects"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
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
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAssociatedObjects(x)
	}

	if v, ok := d.GetOk("block_dates"); ok {
		x := make([]models.SchedulerBlockDate, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.SchedulerBlockDate{}
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
			o.SetClassId("scheduler.BlockDate")
			if v, ok := l["end_time"]; ok {
				{
					x, _ := time.Parse(time.RFC1123, v.(string))
					o.SetEndTime(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_time"]; ok {
				{
					x, _ := time.Parse(time.RFC1123, v.(string))
					o.SetStartTime(x)
				}
			}
			x = append(x, *o)
		}
		o.SetBlockDates(x)
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
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

	if v, ok := d.GetOkExists("enable_block_dates"); ok {
		x := (v.(bool))
		o.SetEnableBlockDates(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
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

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOwners(x)
	}

	if v, ok := d.GetOk("parent"); ok {
		p := make([]models.MoBaseMoRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
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
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("schedule_params"); ok {
		x := make([]models.SchedulerBaseScheduleParams, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.SchedulerBaseScheduleParams{}
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
			o.SetClassId("scheduler.BaseScheduleParams")
			if v, ok := l["duration"]; ok {
				{
					x := (v.(string))
					o.SetDuration(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_time"]; ok {
				{
					x, _ := time.Parse(time.RFC1123, v.(string))
					o.SetStartTime(x)
				}
			}
			if v, ok := l["time_zone"]; ok {
				{
					x := (v.(string))
					o.SetTimeZone(x)
				}
			}
			x = append(x, *o)
		}
		o.SetScheduleParams(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOk("tags"); ok {
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

	if v, ok := d.GetOkExists("usage_count"); ok {
		x := int64(v.(int))
		o.SetUsageCount(x)
	}

	if v, ok := d.GetOk("version_context"); ok {
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
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
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
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
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SchedulerSchedulePolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.SchedulerApi.GetSchedulerSchedulePolicyList(conn.ctx).Filter(getRequestParams(data)).Count(true).Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of SchedulerSchedulePolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of SchedulerSchedulePolicy: %s", responseErr.Error())
	}
	count := countResponse.MoDocumentCount.GetCount()
	if count == 0 {
		return diag.Errorf("your query for SchedulerSchedulePolicy data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var schedulerSchedulePolicyResults = make([]map[string]interface{}, 0, 0)
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.SchedulerApi.GetSchedulerSchedulePolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching SchedulerSchedulePolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching SchedulerSchedulePolicy: %s", responseErr.Error())
		}
		results := resMo.SchedulerSchedulePolicyList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for k := 0; k < len(results); k++ {
				var s = results[k]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["associated_objects"] = flattenListMoBaseMoRelationship(s.GetAssociatedObjects(), d)

				temp["block_dates"] = flattenListSchedulerBlockDate(s.GetBlockDates(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["description"] = (s.GetDescription())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["enable_block_dates"] = (s.GetEnableBlockDates())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["schedule_params"] = flattenListSchedulerBaseScheduleParams(s.GetScheduleParams(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["usage_count"] = (s.GetUsageCount())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				schedulerSchedulePolicyResults = append(schedulerSchedulePolicyResults, temp)
			}
		}
	}
	log.Println("length of results: ", len(schedulerSchedulePolicyResults))
	if err := d.Set("results", schedulerSchedulePolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(schedulerSchedulePolicyResults[0]["moid"].(string))
	return de
}
