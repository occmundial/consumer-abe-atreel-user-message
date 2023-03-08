package models

type State struct {
	StateID   string `json:"StateID,omitempty"`
	StateName string `json:"StateName,omitempty"`
}

type StartStatus struct {
	Success bool  `json:"success" example:"true"`
	Error   error `json:"error" example:"connection error"`
}

type DynamicTemplateData struct {
	Abes    []Abes `json:"ABES" `
	Nombre  string `json:"nombre" `
	Subject string `json:"subject" `
}

type Abes struct {
	URLBusqueda      string `json:"url-busqueda" `
	Cantidad         int    `json:"cantidad" `
	Busqueda         string `json:"busqueda" `
	Region           string `json:"region" `
	FechaCreacionABE string `json:"fechaCreacionABE" `
	Jobs             []Jobs `json:"JOBS" `
}

type Jobs struct {
	URLOferta     string `json:"urlOferta" `
	TituloOferta  string `json:"tituloOferta" `
	URLEmpresa    string `json:"urlEmpresa" `
	NombreEmpresa string `json:"nombreEmpresa" `
	Region        string `json:"region" `
	Sueldo        string `json:"sueldo" `
}

type SendgridJSON struct {
	JobID            []int              `json:"jobIds,omitempty"`
	LoginID          string             `json:"loginId,omitempty"`
	Personalizations []Personalizations `json:"personalizations,omitempty"`
	Platform         string             `json:"platform,omitempty"`
	TemplateID       string             `json:"template_id,omitempty"`
}

type Personalizations struct {
	DynamicTemplateData DynamicTemplateData `json:"dynamicTemplateData,omitempty"`
	To                  []string            `json:"to,omitempty"`
}

type TlalocLocation struct {
	ID            string        `json:"id"`
	StateName     string        `json:"stateName"`
	Compatibility Compatibility `json:"compatibility"`
}

type Compatibility struct {
	Locations string `json:"locations"`
}

type Agents struct {
	Agent  Agent            `json:"agent" `
	Jobs   []SubVacantesJob `json:"jobs" `
	SeoURL string           `json:"seourl" `
}

type Agent struct {
	Agenteid         int               `json:"agenteid" `
	Name             string            `json:"name" `
	NameAgent        string            `json:"name_agent" `
	Location         string            `json:"location" `
	JobsCount        int               `json:"jobs_count" `
	SearchArgs       map[string]string `json:"search_args" `
	FechaCreacionABE string            `json:"fechaCreacionABE" `
	LocationCity     string            `json:"location_city" `
	LocationState    string            `json:"location_state" `
}
type SubVacantesJob struct {
	ID                  int      `json:"id" `
	Title               string   `json:"title" `
	DateExpires         string   `json:"date_expires" `
	Occejecutivo        bool     `json:"occejecutivo" `
	Companyname         string   `json:"companyname" `
	Companynamepretty   string   `json:"companynamepretty" `
	Countryname         string   `json:"countryname" `
	Statename           string   `json:"statename" `
	Cityname            string   `json:"cityname" `
	Showsalary          bool     `json:"showsalary" `
	Salaryfrom          float32  `json:"salaryfrom" `
	Salaryto            float32  `json:"salaryto" `
	Salarytime          int      `json:"salarytime" `
	Locationdescription string   `json:"locationdescription" `
	URL                 string   `json:"url" `
	Urlcompany          string   `json:"urlcompany" `
	FriendlyCompanyURL  string   `json:"friendlycompanyurl" `
	Logo                string   `json:"logo" `
	Jobtype             int      `json:"jobtype" `
	Bullets             []Bullet `json:"bullets" `
	Score               float64  `json:"score" `
}

type Bullet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Tooltip     string `json:"tooltip"`
}

type TlalocResponse struct {
	Locations []TlalocLocation `json:"Locations"`
}
