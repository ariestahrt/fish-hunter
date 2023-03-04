package response

import (
	"fish-hunter/businesses/datasets"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dataset struct {
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

func FromDomain(domain datasets.Domain) Dataset {
	return Dataset{
		Id:         domain.Id,
		Ref_Url:    domain.Ref_Url,
		Ref_Job:    domain.Ref_Job,
		FolderPath: domain.FolderPath,
		HtmlDom: domain.HtmlDom,
		Status: domain.Status,
		HttpCode: domain.HttpCode,
		Assets: domain.Assets,
		Brands: domain.Brands,
		UrlScanUUID: domain.UrlScanUUID,
		ScreenshotPath: domain.ScreenshotPath,
		DomainName: domain.DomainName,
		WhoisRegistrar: domain.WhoisRegistrar,
		WhoisRegistrarUrl: domain.WhoisRegistrarUrl,
		WhoisRegistryCreatedAt: domain.WhoisRegistryCreatedAt,
		WhoisRegistryExpiredAt: domain.WhoisRegistryExpiredAt,
		WhoisRegistryUpdatedAt: domain.WhoisRegistryUpdatedAt,
		WhoisDomainAge: domain.WhoisDomainAge,
		RemoteIpAddress: domain.RemoteIpAddress,
		RemotePort: domain.RemotePort,
		RemoteIpCountryName: domain.RemoteIpCountryName,
		RemoteIpIsp: domain.RemoteIpIsp,
		RemoteIpDomain: domain.RemoteIpDomain,
		RemoteIpAsn: domain.RemoteIpAsn,
		RemoteIpIspOrg: domain.RemoteIpIspOrg,
		Protocol: domain.Protocol,
		SecurityState: domain.SecurityState,
		SecurityProtocol: domain.SecurityProtocol,
		SecurityIssuer: domain.SecurityIssuer,
		SecurityValidFrom: domain.SecurityValidFrom,
		SecurityValidTo: domain.SecurityValidTo,
		Language: domain.Language,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func FromDomainArray(domain []datasets.Domain) []Dataset {
	var response []Dataset
	for _, value := range domain {
		response = append(response, FromDomain(value))
	}
	return response
}