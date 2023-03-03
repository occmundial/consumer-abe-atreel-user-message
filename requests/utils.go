package requests

import (
	"strconv"
	"strings"
	"time"

	"github.com/leekchan/accounting"
	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

const (
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

type ConvertJSONToHTMLAbeData struct {
	OAgentsResults []models.Agents
	Name           string
	StateDic       map[string]string
}

func ConvertJSONToHTMLABE(htmlAbeData *ConvertJSONToHTMLAbeData, config *models.Configuration) ([]int, models.DynamicTemplateData) {
	jobIds := []int{}
	dynamicTemplateData := models.DynamicTemplateData{}
	locationAgent := ""
	locationAgena := ""
	nameAgenta := ""
	countAgent := 0
	countAgena := 0
	ABEs := []models.Abes{}
	for i := range htmlAbeData.OAgentsResults {
		urlCompany := ""
		if htmlAbeData.OAgentsResults[i].Agent.JobsCount > 0 {
			countAgent = htmlAbeData.OAgentsResults[i].Agent.JobsCount
			SearchURL := htmlAbeData.OAgentsResults[i].SeoURL
			nameAgent := htmlAbeData.OAgentsResults[i].Agent.NameAgent
			locationAgent = getLocationAgent(&htmlAbeData.OAgentsResults[i].Agent)
			if countAgent >= countAgena {
				countAgena = countAgent
				locationAgena = locationAgent
				nameAgenta = nameAgent
			}
			dataAbe := models.Abes{}
			dataAbe.URLBusqueda = SearchURL
			dataAbe.Cantidad = htmlAbeData.OAgentsResults[i].Agent.JobsCount
			dataAbe.Busqueda = htmlAbeData.OAgentsResults[i].Agent.NameAgent
			dataAbe.Region = htmlAbeData.OAgentsResults[i].Agent.Location
			dataAbe.FechaCreacionABE = htmlAbeData.OAgentsResults[i].Agent.FechaCreacionABE
			dataJobs := []models.Jobs{}
			for j := range htmlAbeData.OAgentsResults[i].Jobs {
				jobIds = append(jobIds, htmlAbeData.OAgentsResults[i].Jobs[j].ID)
				params := strings.ReplaceAll(config.AbeUtmEmpresa, `|idHtml|`, constants.HTMLDefault)
				url := htmlAbeData.OAgentsResults[i].Jobs[j].FriendlyCompanyURL
				companyNamePretty := htmlAbeData.OAgentsResults[i].Jobs[j].Companynamepretty
				urlCompany = config.Occ + "/" + url + params + companyNamePretty
				dataJobs = append(dataJobs, getJobData(&htmlAbeData.OAgentsResults[i].Jobs[j], config, urlCompany))
			}
			dataAbe.Jobs = dataJobs
			ABEs = append(ABEs, dataAbe)
		}
	}
	dynamicTemplateData.Abes = ABEs
	dynamicTemplateData.Nombre = getFirstName(htmlAbeData.Name)
	if countAgent > 0 {
		subjectPush := getSubjectByDay(SubjectData{nameAgenta, strconv.Itoa(countAgena), locationAgena, int(time.Now().Weekday())})
		dynamicTemplateData.Subject = subjectPush
	}
	return jobIds, dynamicTemplateData
}

func getJobData(jobs *models.SubVacantesJob, config *models.Configuration, urlCompany string) models.Jobs {
	jobData := models.Jobs{}
	jobData.URLEmpresa = urlCompany
	jobData.TituloOferta = jobs.Title
	jobData.URLOferta = jobs.URL
	jobData.NombreEmpresa = jobs.Companyname
	jobData.Region = jobs.Locationdescription
	if jobs.Showsalary {
		ac := accounting.Accounting{Symbol: "$", Precision: 2}
		salaryFrom := ac.FormatMoney(jobs.Salaryfrom)
		salaryTo := ac.FormatMoney(jobs.Salaryto)
		jobData.Sueldo = salaryFrom + " - " + salaryTo + " " + getLookup(jobs.Salarytime)
	} else {
		jobData.Sueldo = config.TxtSalary
	}
	return jobData
}

func getLocationAgent(agent *models.Agent) string {
	if agent.LocationCity != "" {
		if agent.LocationCity == constants.YZonaMetro {
			return agent.LocationState + " " + agent.LocationCity
		} else {
			return agent.LocationCity
		}
	} else if agent.LocationState != "" {
		var locationAgent = stateDic[agent.LocationState]
		if locationAgent == "" {
			return constants.MEXICO
		} else {
			return locationAgent
		}
	} else {
		return constants.MEXICO
	}
}

type SubjectData struct {
	nombre     string
	cantidad   string
	localidad  string
	intWeekday int
}

func getSubjectByDay(subjectData SubjectData) string {
	switch subjectData.intWeekday {
	case Monday:
		return getMondySubject(subjectData)
	case Tuesday:
		return getTuesdyaSubject(subjectData)
	case Wednesday:
		return getWednesdaySubject(subjectData)
	case Thursday:
		return getThurdaySubject(subjectData)
	case Friday:
		return getFridaySubject(subjectData)
	case Saturday:
		return getSaturdaySubject(subjectData)
	default:
		return getDefaultSubject(subjectData)
	}
}

func getMondySubject(subjectData SubjectData) string {
	if subjectData.nombre != "" {
		subjectData.nombre = " de " + subjectData.nombre
	}
	complemento := "Nuevos "
	emp := " empleos"
	if subjectData.cantidad == "1" {
		complemento = "Nuevo "
		emp = "empleo"
		subjectData.cantidad = ""
	}
	if subjectData.localidad != "" {
		subjectData.localidad = " en " + subjectData.localidad
	}
	return complemento + subjectData.cantidad + emp + subjectData.localidad + subjectData.nombre
}

func getTuesdyaSubject(subjectData SubjectData) string {
	if subjectData.nombre != "" {
		subjectData.nombre = "de " + subjectData.nombre
	}
	complemento := " empleos nuevos "
	if subjectData.cantidad == "1" {
		complemento = " empleo nuevo "
	}
	if subjectData.localidad != "" {
		subjectData.localidad = "En " + subjectData.localidad + " "
	}
	return subjectData.localidad + "hay " + subjectData.cantidad + complemento + subjectData.nombre
}

func getWednesdaySubject(subjectData SubjectData) string {
	complemento := " empleos nuevos"
	if subjectData.cantidad == "1" {
		complemento = " empleo nuevo"
	}
	if subjectData.nombre != "" && subjectData.localidad != "" {
		subjectData.localidad = " en " + subjectData.localidad
	}
	return subjectData.nombre + subjectData.localidad + " ¡" + subjectData.cantidad + complemento + "!"
}

func getThurdaySubject(subjectData SubjectData) string {
	if subjectData.nombre != "" {
		subjectData.nombre = "de " + subjectData.nombre
	}
	complemento := " vacantes nuevas "
	if subjectData.cantidad == "1" {
		complemento = " vacante nueva "
	}
	if subjectData.localidad != "" {
		subjectData.localidad += " tiene "
	}
	return subjectData.localidad + subjectData.cantidad + complemento + subjectData.nombre
}

func getFridaySubject(subjectData SubjectData) string {
	complemento := " vacantes perfectas "
	if subjectData.cantidad == "1" {
		complemento = " vacante perfecta "
	}
	if subjectData.localidad != "" {
		subjectData.localidad = " en " + subjectData.localidad
	}
	return "Hay " + subjectData.cantidad + complemento + "para ti" + subjectData.localidad
}

func getSaturdaySubject(subjectData SubjectData) string {
	if subjectData.nombre != "" {
		subjectData.nombre = " de " + subjectData.nombre
	}
	complemento := " nuevos empleos"
	if subjectData.cantidad == "1" {
		complemento = " nuevo empleo"
	}
	if subjectData.localidad != "" {
		subjectData.localidad = " en " + subjectData.localidad
	}
	return subjectData.cantidad + complemento + subjectData.localidad + subjectData.nombre
}

func getDefaultSubject(subjectData SubjectData) string {
	complemento := " vacantes nuevas"
	if subjectData.cantidad == "1" {
		complemento = " vacante nueva"
	}
	if subjectData.localidad != "" {
		subjectData.localidad = " en " + subjectData.localidad
	}
	return "¡" + subjectData.cantidad + complemento + "! " + subjectData.nombre + subjectData.localidad
}

func getLookup(salarytime int) string {
	switch salarytime {
	case 1:
		return "Hora"
	case 2:
		return "Semanal"
	case 3:
		return "Anual"
	default:
		return "Mensual"
	}
}

func getFirstName(fullName string) string {
	arrayFullName := strings.Split(fullName, " ")
	if len(arrayFullName) == 0 {
		return ""
	} else {
		return arrayFullName[0]
	}
}
