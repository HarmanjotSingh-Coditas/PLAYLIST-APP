package constants

const (
	PlaylistId  = "playlist_id"
	UserId      = "user_id"
	Name        = "name"
	Description = "description"
	SongId      = "song_id"
)

// Constants Related to postgres-Setup
const (
	Host       = "host"
	Port       = "port"
	Dbname     = "dbname"
	Password   = "password"
	User       = "user"
	Sslmode    = "sslmode"
	ConfigType = "yml"
	ConfigName = "postgres"
	ConfigPath = "../../../src/config"
)

// Validations
const (
	Required = "required"
)

// Main function

const (
	ConnectingToServer    = "Connecting to the server on port %s"
	ConnectedToServer     = "Connected to the server on port %s"
	ServerShutdownSignal  = "Termination Signal Recieved , shutting down the server"
)
