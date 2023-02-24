package requests

import (
	"strconv"
	"strings"
	"time"

	"github.com/leekchan/accounting"
	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

func ConvertJsonToHtml_ABE(oAgentsResults []models.Agents, nombre string, stateDic map[string]string, config *models.Configuration) ([]int, models.DynamicTemplateData) {
	jobIds := []int{}
	dynamicTemplateData := models.DynamicTemplateData{}
	locationAgent := ""
	locationAgena := ""
	nameAgenta := ""
	countAgent := 0
	countAgena := 0
	ABEs := []models.Abes{}
	for _, agentResult := range oAgentsResults {
		agent := agentResult.Agent
		jobs := agentResult.Jobs
		encontrados := agent.JobsCount
		urlCompany := ""
		if encontrados > 0 {
			countAgent = encontrados
			urlResultados := agentResult.SeoUrl
			nameAgent := agent.NameAgent
			if agent.LocationCity != "" {
				if agent.LocationCity == constants.YZonaMetro {
					locationAgent = agent.LocationState + " " + agent.LocationCity
				} else {
					locationAgent = agent.LocationCity
				}
			} else if agent.LocationState != "" {
				locationAgent = stateDic[agent.LocationState]
				if locationAgent == "" {
					locationAgent = constants.MEXICO
				}
			} else {
				locationAgent = constants.MEXICO
			}
			if countAgent >= countAgena {
				countAgena = countAgent
				locationAgena = locationAgent
				nameAgenta = nameAgent
			}
			dataAbe := models.Abes{}
			dataAbe.UrlBusqueda = urlResultados
			dataAbe.Cantidad = encontrados
			dataAbe.Busqueda = agent.NameAgent
			dataAbe.Region = agent.Location
			dataAbe.FechaCreacionABE = agent.FechaCreacionABE
			dataJobs := []models.Jobs{}
			for _, job := range jobs {
				jobIds = append(jobIds, job.Id)
				urlCompany = config.Occ + "/" + job.FriendlyCompanyUrl + strings.Replace(config.ABE_UTM_EMPRESA, `|idHtml|`, constants.HtmlDefault, -1) + job.Companynamepretty
				jobData := models.Jobs{}
				jobData.UrlEmpresa = urlCompany
				jobData.TituloOferta = job.Title
				jobData.UrlOferta = job.Url
				jobData.NombreEmpresa = job.Companyname
				jobData.Region = job.Locationdescription
				if job.Showsalary {
					ac := accounting.Accounting{Symbol: "$", Precision: 2}
					jobData.Sueldo = ac.FormatMoney(job.Salaryfrom) + " - " + ac.FormatMoney(job.Salaryto) + " " + getLookup(job.Salarytime)
				} else {
					jobData.Sueldo = config.TxtSalary
				}
				dataJobs = append(dataJobs, jobData)
			}
			dataAbe.Jobs = dataJobs
			ABEs = append(ABEs, dataAbe)
		}
	}
	dynamicTemplateData.Abes = ABEs
	dynamicTemplateData.Nombre = getFirstName(nombre)
	if countAgent > 0 {
		subjectPush := getSubjectByDay(nameAgenta, strconv.Itoa(countAgena), locationAgena, int(time.Now().Weekday()))
		dynamicTemplateData.Subject = subjectPush
	}
	return jobIds, dynamicTemplateData
}

func getSubjectByDay(nombre string, cantidad string, localidad string, intWeekday int) string {
	switch intWeekday {
	case 1:
		if nombre != "" {
			nombre = " de " + nombre
		}
		complemento := "Nuevos "
		emp := " empleos"
		if cantidad == "1" {
			complemento = "Nuevo "
			emp = "empleo"
			cantidad = ""
		}
		if localidad != "" {
			localidad = " en " + localidad
		}
		return complemento + cantidad + emp + localidad + nombre
	case 2:
		if nombre != "" {
			nombre = "de " + nombre
		}
		complemento := " empleos nuevos "
		if cantidad == "1" {
			complemento = " empleo nuevo "
		}
		if localidad != "" {
			localidad = "En " + localidad + " "
		}
		return localidad + "hay " + cantidad + complemento + nombre
	case 3:
		complemento := " empleos nuevos"
		if cantidad == "1" {
			complemento = " empleo nuevo"
		}
		if nombre != "" && localidad != "" {
			localidad = " en " + localidad
		}
		return nombre + localidad + " ¡" + cantidad + complemento + "!"
	case 4:
		if nombre != "" {
			nombre = "de " + nombre
		}
		complemento := " vacantes nuevas "
		if cantidad == "1" {
			complemento = " vacante nueva "
		}
		if localidad != "" {
			localidad = localidad + " tiene "
		}
		return localidad + cantidad + complemento + nombre
	case 5:
		complemento := " vacantes perfectas "
		if cantidad == "1" {
			complemento = " vacante perfecta "
		}
		if localidad != "" {
			localidad = " en " + localidad
		}
		return "Hay " + cantidad + complemento + "para ti" + localidad
	case 6:
		if nombre != "" {
			nombre = " de " + nombre
		}
		complemento := " nuevos empleos"
		if cantidad == "1" {
			complemento = " nuevo empleo"
		}
		if localidad != "" {
			localidad = " en " + localidad
		}
		return cantidad + complemento + localidad + nombre
	default:
		complemento := " vacantes nuevas"
		if cantidad == "1" {
			complemento = " vacante nueva"
		}
		if localidad != "" {
			localidad = " en " + localidad
		}
		return "¡" + cantidad + complemento + "! " + nombre + localidad
	}
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
