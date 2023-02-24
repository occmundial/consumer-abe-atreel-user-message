package requests

import (
	"strings"
	"testing"

	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/stretchr/testify/assert"
)

var (
	configuration = models.Configuration{Occ: "www.occ.com.mx", TxtSalary: "Sin sueldo"}
	nombre        = "nombreJob"
	testStateDic  = make(map[string]string)
	agents        = models.Agents{
		Agent: models.Agent{
			Agenteid:         1,
			Name:             "name",
			NameAgent:        "nameAgent",
			Location:         "location",
			JobsCount:        1,
			SearchArgs:       make(map[string]string),
			FechaCreacionABE: "fecha",
			LocationCity:     "",
			LocationState:    "001",
		},
		Jobs:   []models.SubVacantesJob{},
		SeoUrl: "urls",
	}
	agents2 = models.Agents{
		Agent: models.Agent{
			Agenteid:         2,
			Name:             "name",
			NameAgent:        "nameAgent",
			Location:         "location",
			JobsCount:        1,
			SearchArgs:       make(map[string]string),
			FechaCreacionABE: "fecha",
			LocationCity:     "Corregidora",
			LocationState:    "",
		},
		Jobs:   []models.SubVacantesJob{},
		SeoUrl: "urls",
	}

	agents3 = models.Agents{
		Agent: models.Agent{
			Agenteid:         2,
			Name:             "name",
			NameAgent:        "nameAgent",
			Location:         "location",
			JobsCount:        1,
			SearchArgs:       make(map[string]string),
			FechaCreacionABE: "fecha",
			LocationCity:     "",
			LocationState:    "",
		},
		Jobs:   []models.SubVacantesJob{},
		SeoUrl: "urls",
	}
	subVacantesJobs  = models.SubVacantesJob{Companynamepretty: "Companynamepretty", Urlcompany: "urlcompany", Salaryfrom: 15000, Salarytime: 1}
	subVacantesJobs2 = models.SubVacantesJob{Companynamepretty: "Companynamepretty2", Urlcompany: "urlcompany2", Salaryfrom: 1000, Salarytime: 2, Showsalary: true}
	agentsWithJobs   = models.Agents{
		Agent: models.Agent{
			Agenteid:         2,
			Name:             "name",
			NameAgent:        "golang",
			Location:         "location",
			JobsCount:        2,
			SearchArgs:       make(map[string]string),
			FechaCreacionABE: "fecha",
			LocationCity:     constants.YZonaMetro,
			LocationState:    "Puebla",
		},
		Jobs:   []models.SubVacantesJob{subVacantesJobs, subVacantesJobs2},
		SeoUrl: "urls",
	}
)

func Test_GetSubjectByDay_OK_Monday_1(t *testing.T) {
	subject := getSubjectByDay("", "1", "Qro", 1)
	assert.Equal(t, "Nuevo empleo en Qro", subject)
}

func Test_GetSubjectByDay_OK_Monday_2(t *testing.T) {
	subject := getSubjectByDay("golang", "2", "", 1)
	assert.Equal(t, "Nuevos 2 empleos de golang", subject)
}

func Test_GetSubjectByDay_OK_Tuesday(t *testing.T) {
	subject := getSubjectByDay("", "1", "Qro", 2)
	assert.Equal(t, "En Qro hay 1 empleo nuevo ", subject)
}

func Test_GetSubjectByDay_OK_Tuesday_2(t *testing.T) {
	subject := getSubjectByDay("golang", "2", "", 2)
	assert.Equal(t, "hay 2 empleos nuevos de golang", subject)
}

func Test_GetSubjectByDay_OK_Wenesday(t *testing.T) {
	subject := getSubjectByDay("", "1", "Qro", 3)
	assert.Equal(t, "Qro ¡1 empleo nuevo!", subject)
}

func Test_GetSubjectByDay_OK_Wenesday_2(t *testing.T) {
	subject := getSubjectByDay("golang", "2", "", 3)
	assert.Equal(t, "golang ¡2 empleos nuevos!", subject)
}

func Test_GetSubjectByDay_OK_Wenesday_3(t *testing.T) {
	subject := getSubjectByDay("golang", "3", "Qro", 3)
	assert.Equal(t, "golang en Qro ¡3 empleos nuevos!", subject)
}

func Test_GetSubjectByDay_OK_Thuesday(t *testing.T) {
	subject := getSubjectByDay("", "1", "Qro", 4)
	assert.Equal(t, "Qro tiene 1 vacante nueva ", subject)
}

func Test_GetSubjectByDay_OK_Thuesday_2(t *testing.T) {
	subject := getSubjectByDay("golang", "2", "", 4)
	assert.Equal(t, "2 vacantes nuevas de golang", subject)
}

func Test_GetSubjectByDay_OK_Friday(t *testing.T) {
	subject := getSubjectByDay("", "1", "Qro", 5)
	assert.Equal(t, "Hay 1 vacante perfecta para ti en Qro", subject)
}

func Test_GetSubjectByDay_OK_Friday_2(t *testing.T) {
	subject := getSubjectByDay("golang", "2", "", 5)
	assert.Equal(t, "Hay 2 vacantes perfectas para ti", subject)
}

func Test_GetSubjectByDay_OK_Sartuday(t *testing.T) {
	subject := getSubjectByDay("", "1", "Qro", 6)
	assert.Equal(t, "1 nuevo empleo en Qro", subject)
}

func Test_GetSubjectByDay_OK_Sartuday_2(t *testing.T) {
	subject := getSubjectByDay("golang", "2", "", 6)
	assert.Equal(t, "2 nuevos empleos de golang", subject)
}

func Test_GetSubjectByDay_OK_Sunday(t *testing.T) {
	subject := getSubjectByDay("", "1", "Qro", 0)
	assert.Equal(t, "¡1 vacante nueva!  en Qro", subject)
}

func Test_GetSubjectByDay_OK_Sunday_2(t *testing.T) {
	subject := getSubjectByDay("golang", "2", "", 0)
	assert.Equal(t, "¡2 vacantes nuevas! golang", subject)
}

func Test_ConvertJsonToHtml_ABE_Empty_ALL_OK(t *testing.T) {
	jobIds, dynamicTemplateData := ConvertJsonToHtml_ABE([]models.Agents{}, nombre, testStateDic, &configuration)
	assert.Equal(t, 0, len(jobIds))
	assert.Equal(t, nombre, dynamicTemplateData.Nombre)
	assert.Equal(t, "", dynamicTemplateData.Subject)
	assert.Equal(t, 0, len(dynamicTemplateData.Abes))
}

func Test_ConvertJsonToHtml_ABE_Without_SubVacantesJob__Mexico_OK(t *testing.T) {
	agents := []models.Agents{agents}
	jobIds, dynamicTemplateData := ConvertJsonToHtml_ABE(agents, nombre, testStateDic, &configuration)
	assert.Equal(t, 0, len(jobIds))
	assert.Equal(t, nombre, dynamicTemplateData.Nombre)
	assert.True(t, strings.Contains(dynamicTemplateData.Subject, constants.MEXICO))
	assert.Equal(t, 1, len(dynamicTemplateData.Abes))
}

func Test_ConvertJsonToHtml_ABE_Without_SubVacantesJob__Mexico2_OK(t *testing.T) {
	agents := []models.Agents{agents2}
	jobIds, dynamicTemplateData := ConvertJsonToHtml_ABE(agents, nombre, testStateDic, &configuration)
	assert.Equal(t, 0, len(jobIds))
	assert.Equal(t, nombre, dynamicTemplateData.Nombre)
	assert.True(t, strings.Contains(dynamicTemplateData.Subject, "Corregidora"))
	assert.Equal(t, 1, len(dynamicTemplateData.Abes))
}

func Test_ConvertJsonToHtml_ABE_Without_SubVacantesJob__Mexico3_OK(t *testing.T) {
	agents := []models.Agents{agents3}
	jobIds, dynamicTemplateData := ConvertJsonToHtml_ABE(agents, nombre, testStateDic, &configuration)
	assert.Equal(t, 0, len(jobIds))
	assert.Equal(t, nombre, dynamicTemplateData.Nombre)
	assert.True(t, strings.Contains(dynamicTemplateData.Subject, constants.MEXICO))
	assert.Equal(t, 1, len(dynamicTemplateData.Abes))
}

func Test_ConvertJsonToHtml_ABE_Without_SubVacantesJob_Queretaro_OK(t *testing.T) {
	testStateDic["001"] = "Querétaro"
	agents := []models.Agents{agents}
	jobIds, dynamicTemplateData := ConvertJsonToHtml_ABE(agents, nombre, testStateDic, &configuration)
	assert.Equal(t, 0, len(jobIds))
	assert.Equal(t, nombre, dynamicTemplateData.Nombre)
	assert.True(t, strings.Contains(dynamicTemplateData.Subject, "Querétaro"))
	assert.Equal(t, 1, len(dynamicTemplateData.Abes))
}

func Test_ConvertJsonToHtml_ABE_With_SubVacantesJob_OK(t *testing.T) {
	agents := []models.Agents{agentsWithJobs, agentsWithJobs}
	jobIds, dynamicTemplateData := ConvertJsonToHtml_ABE(agents, nombre, testStateDic, &configuration)
	assert.Equal(t, 4, len(jobIds))
	assert.Equal(t, nombre, dynamicTemplateData.Nombre)
	assert.True(t, len(dynamicTemplateData.Subject) > 0)
	assert.Equal(t, 2, len(dynamicTemplateData.Abes))
}
