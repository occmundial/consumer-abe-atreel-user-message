package constants

const (
	StatusProcessStartError = "process-start-error"
	// StatusReadMessageError : error durante la lectura del mensaje
	StatusReadMessageError = "read-message-error"
	// StatusInvalidMessage : el mensaje no tiene el mínimo de valores para considerarlo como válido
	StatusInvalidMessage = "invalid-message-error"
	// StatusCheckProcessError : error durante verificación de mensaje procesado
	StatusCheckProcessError = "checking-error"
	// StatusAlreadyProcessed : el mensaje fue procesado con anterioridad o ya no se necesita procesarlo
	StatusAlreadyProcessed = "no-need-to-process"
	// StatusFullProcessOK : Procesamiento completo del mensaje (mensaje procesado y borrado de la cola)
	StatusFullProcessOK = "full-processed-message"
	// StatusProcessOK : Mensaje procesado exitosamente pero no borrado de la cola
	StatusProcessOK = "processed-not-committed-message"
	// StatusProcessError : Error en procesamiento del mensaje: falló el procesamiento del mensaje
	StatusProcessError = "processing-error"
	// LimitReader : cantidad máxima de bytes a leer del body del response
	LimitReader = 1048576

	HTMLDefault     = "1"
	YZonaMetro      = "Y Zona Metro."
	MEXICO          = "México"
	TemplateID      = "d-1488308ecee24468b2f83c52bbeadf11"
	Platform        = "ABE"
	JSONContentType = "application/json"
)
