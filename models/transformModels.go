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
	UrlBusqueda      string `json:"url-busqueda" `
	Cantidad         int    `json:"cantidad" `
	Busqueda         string `json:"busqueda" `
	Region           string `json:"region" `
	FechaCreacionABE string `json:"fechaCreacionABE" `
	Jobs             []Jobs `json:"JOBS" `
}

type Jobs struct {
	UrlOferta     string `json:"urlOferta" `
	TituloOferta  string `json:"tituloOferta" `
	UrlEmpresa    string `json:"urlEmpresa" `
	NombreEmpresa string `json:"nombreEmpresa" `
	Region        string `json:"region" `
	Sueldo        string `json:"sueldo" `
}

type SendgridJson struct {
	JobID            []int              `json:"jobIds,omitempty"`
	LoginID          string             `json:"loginId,omitempty"`
	Personalizations []Personalizations `json:"personalizations,omitempty"`
	Platform         string             `json:"platform,omitempty"`
	Template_ID      string             `json:"template_id,omitempty"`
}

type Personalizations struct {
	DynamicTemplateData DynamicTemplateData `json:"dynamicTemplateData,omitempty"`
	To                  []string            `json:"to,omitempty"`
}

type Agents struct {
	Agent  Agent            `json:"agent" `
	Jobs   []SubVacantesJob `json:"jobs" `
	SeoUrl string           `json:"seourl" `
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
	Id                  int      `json:"id" `
	Title               string   `json:"title" `
	Date_expires        string   `json:"date_expires" `
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
	Url                 string   `json:"url" `
	Urlcompany          string   `json:"urlcompany" `
	FriendlyCompanyUrl  string   `json:"friendlycompanyurl" `
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
