package reclameaqui

// Pagination limits
const (
	maxSize        = 50
	maxPerPage     = 10
	defaultSize    = 10
	defaultPerPage = 10
)

// HTTP client configuration
const (
	baseURL        = "https://morpheus-bff.reclameaqui.com.br"
	userAgent      = "PostmanRuntime/7.51.0"
	defaultTimeout = 30
	ja3Fingerprint = "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-17513,29-23-24,0"
)

// API endpoints
const (
	endpointSearchCompany     = "/v1/companies"
	endpointDescribeCompany   = "/v1/companies/%s"
	endpointListComplaints    = "/v1/complaints"
	endpointDescribeComplaint = "/v1/complaints/%s"
)
