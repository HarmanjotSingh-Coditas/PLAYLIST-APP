package constants

// Router
const (
	DatabaseConnectionError = "error connecting to database %w"
	DatabasePingingError    = "error pinging database %w"
)

// Create Playlist Service
const (
	SongIdsDoesNotExistsError  = "empty playlist created as one or more song IDs do not exist"
	PlaylistAlreadyExistsError = "playlist already exists"
)

// Create Playlist Handler
const (
	JsonBindingFieldError       = "JsonBindingFieldError"
	PlaylistCreationFailedError = "playlist could not be created"
	FailedToCreatePlaylist      = "failed to create playlist"
	UnexpectedError             = "An unexpected error occurred"
)

// Ad playlist Repo

const (
	PlaylistExistenceCheckError   = "failed to check playlist existence"
	SongExistenceCheckingError    = "failed to check song existence"
	NoSongsProvidedError          = "no songs provided to add"
	SongAddingToPlaylistError     = "failed to add songs to playlist"
	SongDeletionfromPlaylistError = "failed to delete songs from playlist"
	NoSongsDeletedFromPlaylist    = "no songs were deleted from playlist"
)

// Ad playlist Service

const (
	PlaylistDoesNotExistsError    = "playlist does not exist"
	SongIdsAlreadyInPlaylistError = "songs with IDs [%v] already exist in playlist"
	NoValidSongsToDeleteError     = "no valid songs to delete. Non-existent songs: %v"
	InvalidAction                 = "invalid Action"
	PlaylistNotFoundError         = "playlist not found"
	NoValidSongsToAddError        = "no valid songs to add"
	NoValidSongsToBeDeletedError  = "no valid songs to delete"
	InvalidActionsError           = "Invalid action. Must be either 'ADD' or 'DELETE'"
)
