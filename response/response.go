package response


type ErrorResponse struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"description"`
}

type CatalogResponse struct {
	Services []Service `json:"services"`
	Error string `json:"error"`
}


type Service struct {
	ID                   string                  `json:"id"`
	Name                 string                  `json:"name"`
	Description          string                  `json:"description"`
	Bindable             bool                    `json:"bindable"`
	InstancesRetrievable bool                    `json:"instances_retrievable,omitempty"`
	BindingsRetrievable  bool                    `json:"bindings_retrievable,omitempty"`
	Tags                 []string                `json:"tags,omitempty"`
	PlanUpdatable        bool                    `json:"plan_updateable"`
	Plans                []ServicePlan           `json:"plans"`
	Requires             []string				 `json:"requires,omitempty"`
	Metadata             *ServiceMetadata        `json:"metadata,omitempty"`
	DashboardClient      *ServiceDashboardClient `json:"dashboard_client,omitempty"`
}

type ServiceDashboardClient struct {
	ID          string `json:"id"`
	Secret      string `json:"secret"`
	RedirectURI string `json:"redirect_uri"`
}

type ServicePlan struct {
	ID              string               `json:"id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	Free            *bool                `json:"free,omitempty"`
	Bindable        *bool                `json:"bindable,omitempty"`
	Metadata        *ServicePlanMetadata `json:"metadata,omitempty"`
	Schemas         *ServiceSchemas      `json:"schemas,omitempty"`
	MaintenanceInfo *MaintenanceInfo     `json:"maintenance_info,omitempty"`
}

type ServiceSchemas struct {
	Instance ServiceInstanceSchema `json:"service_instance,omitempty"`
	Binding  ServiceBindingSchema  `json:"service_binding,omitempty"`
}

type ServiceInstanceSchema struct {
	Create Schema `json:"create,omitempty"`
	Update Schema `json:"update,omitempty"`
}

type ServiceBindingSchema struct {
	Create Schema `json:"create,omitempty"`
}

type Schema struct {
	Parameters map[string]interface{} `json:"parameters"`
}

type ServicePlanMetadata struct {
	DisplayName        string            `json:"displayName,omitempty"`
	Bullets            []string          `json:"bullets,omitempty"`
	Costs              []ServicePlanCost `json:"costs,omitempty"`
	AdditionalMetadata map[string]interface{}
}

type ServicePlanCost struct {
	Amount map[string]float64 `json:"amount"`
	Unit   string             `json:"unit"`
}

type ServiceMetadata struct {
	DisplayName         string `json:"displayName,omitempty"`
	ImageUrl            string `json:"imageUrl,omitempty"`
	LongDescription     string `json:"longDescription,omitempty"`
	ProviderDisplayName string `json:"providerDisplayName,omitempty"`
	DocumentationUrl    string `json:"documentationUrl,omitempty"`
	SupportUrl          string `json:"supportUrl,omitempty"`
	Shareable           *bool  `json:"shareable,omitempty"`
	AdditionalMetadata  map[string]interface{}
}

type MaintenanceInfo struct {
	Public  map[string]string `json:"public,omitempty"`
	Private string            `json:"private,omitempty"`
	Version string            `json:"version,omitempty"`
}