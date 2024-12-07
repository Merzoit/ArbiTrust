package constants

const (
	ErrEncodingData     = "failed to encode data: %v"
	ErrSendingRequest   = "failed to send request to API: %v"
	ErrFetchingResponse = "failed to fetch response from API: %v"
	ErrDecodingResponse = "failed to decode response from API: %v"
	ErrSendingMessage   = "failed to send message to user: %v"

	//TEAM
	LogErrorAddingTeam   = "error adding team: %v"
	ErrFetchingTeam      = "failed to fetch teams: %v"
	ErrSendingTeamList   = "error sending team list: %v"
	ErrNavigationHandler = "error handler navigation: %v"
	ErrHandlerNoFound    = "no handler found for step: %v"
	ErrHandlerStep       = "error on step %d for user %v: %v"
	ErrSendingStep       = "error sending step: %v"
	//USER
	LogErrorAddingUser = "error adding user: %v"
	//PUBLIC
	LogErrorAddingPublic = "error adding public: %v"
	ErrFetchingPublic    = "failed to fetch public: %v"
	//OTHER
	LogEmptyList = "list is empty"
)
