package constants

// Postgres-Setup Errors
const (
	DatabaseRetrievalError      = "error retrieving database %w"
	DatabaseAuthenticationError = "error authenticating with the DataBase %w"
	DatabaseNilError            = "database is nil "
)

// Validations Related Errprs

const (
	FieldRequiredError = "this field is required"
)

// Main

const (
	UnableToConnectServerError = "unable to connect to server"
)
