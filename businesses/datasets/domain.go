package datasets

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Ref

{
  "_id": {
    "$oid": "63d22b61df6255c6bb785e6f"
  },
  "ref_url": {
    "$oid": "63d22a4f1ed2dd0ec7f6042b"
  },
  "ref_job": {
    "$oid": "63d22b0adf6255c6bb785e6e"
  },
  "entry_date": {
    "$date": {
      "$numberLong": "1674746849000"
    }
  },
  "url": "https://www.metamaskpendingverification.duckdns.org/",
  "folder_path": "datasets/63d22a4f1ed2dd0ec7f6042b",
  "htmldom_path": "datasets/63d22a4f1ed2dd0ec7f6042b/index.html",
  "status": "new",
  "http_status_code": 200,
  "assets_downloaded": 0.64,
  "brands": [
    "metamask"
  ],
  "urlscan_uuid": "6601af63-a050-4059-a696-50433547019c",
  "screenshot_path": null,
  "domain_name": null,
  "whois_registrar": null,
  "whois_registrar_url": null,
  "whois_registry_created_at": null,
  "whois_registry_expired_at": null,
  "whois_registry_updated_at": null,
  "whois_domain_age": null,
  "remote_ip_address": "163.123.141.198",
  "remote_port": 443,
  "remote_ip_country_name": "United States",
  "remote_ip_isp": "Des Capital B.V.",
  "remote_ip_domain": "serverion.com",
  "remote_ip_asn": 213035,
  "remote_ip_isp_org": "Serverion LLC",
  "protocol": "http/1.1",
  "security_state": "secure",
  "security_protocol": "TLS 1.2",
  "security_issuer": "R3",
  "security_valid_from": 1674588635,
  "security_valid_to": 1682364634,
  "created_at": {
    "$date": {
      "$numberLong": "1674746849000"
    }
  },
  "updated_at": {
    "$date": {
      "$numberLong": "1674746849000"
    }
  },
  "deleted_at": null
}

*/

type Domain struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Ref_Url    primitive.ObjectID `json:"-" bson:"ref_url,omitempty"`
	Ref_Job    primitive.ObjectID `json:"-" bson:"ref_job,omitempty"`
	EntryDate  primitive.DateTime `json:"entry_date,omitempty" bson:"entry_date,omitempty"`
	Url        string             `json:"url,omitempty" bson:"url,omitempty"`
	FolderPath string             `json:"folder_path,omitempty" bson:"folder_path,omitempty"`
	HtmlDom    string             `json:"htmldom_path,omitempty" bson:"htmldom_path,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	HttpCode   int                `json:"http_status_code,omitempty" bson:"http_status_code,omitempty"`
	Assets     float64            `json:"assets_downloaded,omitempty" bson:"assets_downloaded,omitempty"`
	Brands     []string           `json:"brands,omitempty" bson:"brands,omitempty"`
	UrlScanUUID	string             `json:"urlscan_uuid,omitempty" bson:"urlscan_uuid,omitempty"`
	ScreenshotPath string           `json:"screenshot_path,omitempty" bson:"screenshot_path,omitempty"`
	DomainName string             `json:"domain_name,omitempty" bson:"domain_name,omitempty"`
	WhoisRegistrar string           `json:"whois_registrar,omitempty" bson:"whois_registrar,omitempty"`
	WhoisRegistrarUrl string        `json:"whois_registrar_url,omitempty" bson:"whois_registrar_url,omitempty"`
	WhoisRegistryCreatedAt primitive.DateTime `json:"whois_registry_created_at,omitempty" bson:"whois_registry_created_at,omitempty"`
	WhoisRegistryExpiredAt primitive.DateTime `json:"whois_registry_expired_at,omitempty" bson:"whois_registry_expired_at,omitempty"`
	WhoisRegistryUpdatedAt primitive.DateTime `json:"whois_registry_updated_at,omitempty" bson:"whois_registry_updated_at,omitempty"`
	WhoisDomainAge int64           `json:"whois_domain_age,omitempty" bson:"whois_domain_age,omitempty"`
	RemoteIpAddress string          `json:"remote_ip_address,omitempty" bson:"remote_ip_address,omitempty"`
	RemotePort int                 `json:"remote_port,omitempty" bson:"remote_port,omitempty"`
	RemoteIpCountryName string      `json:"remote_ip_country_name,omitempty" bson:"remote_ip_country_name,omitempty"`
	RemoteIpIsp string             `json:"remote_ip_isp,omitempty" bson:"remote_ip_isp,omitempty"`
	RemoteIpDomain string           `json:"remote_ip_domain,omitempty" bson:"remote_ip_domain,omitempty"`
	RemoteIpAsn int                 `json:"remote_ip_asn,omitempty" bson:"remote_ip_asn,omitempty"`
	RemoteIpIspOrg string           `json:"remote_ip_isp_org,omitempty" bson:"remote_ip_isp_org,omitempty"`
	Protocol string                `json:"protocol,omitempty" bson:"protocol,omitempty"`
	SecurityState string            `json:"security_state,omitempty" bson:"security_state,omitempty"`
	SecurityProtocol string         `json:"security_protocol,omitempty" bson:"security_protocol,omitempty"`
	SecurityIssuer string           `json:"security_issuer,omitempty" bson:"security_issuer,omitempty"`
	SecurityValidFrom primitive.DateTime `json:"security_valid_from,omitempty" bson:"security_valid_from,omitempty"`
	SecurityValidTo primitive.DateTime `json:"security_valid_to,omitempty" bson:"security_valid_to,omitempty"`
  Language  string             `json:"language,omitempty" bson:"language,omitempty"`
	CreatedAt  primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt   primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type UseCase interface {
	Status(status string) ([]Domain, error)
	GetByID(id string) (Domain, error)
	Activate(id string) (string, error)
	Validate(domain Domain) (Domain, error)
	Download(id string) (string, error)
	TopBrands() (map[string]interface{}, error)
}

type Repository interface {
	Status(status string) ([]Domain, error)
	GetByID(id string) (Domain, error)
	Validate(domain Domain) (Domain, error)
	TopBrands() (map[string]interface{}, error)
	CountTotal() (int64, error)
	CountTotalValid() (int64, error)
	GetTotalBetweenDates(start time.Time, end time.Time) (int64, error)
}