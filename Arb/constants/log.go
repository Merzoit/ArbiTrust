package constants

const (
	//Errors logs
	//USER
	LogErrorAddingUser   = "Error adding user: %v"
	LogErrorFetchingUser = "Error fetching user: %v"
	LogErrorUpdatingUser = "Error updating user: %v"
	LogErrorDeletingUser = "Error deleting user: %v"
	LogErrorUserScanning = "error scanning row: %v"
	//TEAM
	LogErrorAddingTeam   = "Error adding team: %v"
	LogErrorUpdatingTeam = "Error updating team: %v"
	LogErrorFetchingTeam = "Error fetching team: %v"
	LogErrorDeletingTeam = "Error deleting team: %v"
	LogErrorTeamScanning = "error scanning row: %v"
	//REQUESTS
	LogErrorDecodingRequestBody = "Error decoding request body: %v"
	LogErrorEncodingResponse    = "Error encoding response: %v"
	//OTHER
	LogErrorExtractingID = "Error extracting ID from: %v"

	//Success logs
	//USER
	LogUserCreateSuccessfully = "User created successfully: %v"
	LogUserFetchSuccessfully  = "User fetched successfully: %v"
	LogUserUpdateSuccessfully = "User updated successfully: %v"
	LogUserDeleteSuccessfully = "User deleted successfully: %v"
	//TEAM
	LogTeamCreateSuccessfully = "Team created successfully: %v"
	LogTeamFetchSuccessfully  = "Team fetched successfully: %v"
	LogTeamUpdateSuccessfully = "Team updated successfully: %v"
	LogTeamDeleteSuccessfully = "Team deleted successfully: %v"
	LogFetchTeamCount         = "Successfully fetched %d teams"
	//OTHER
	LogFetchedUserID         = "Successfully fetched user with id: %d"
	LogExtractIDSuccessfully = "Exctract id successfully from: %v"
	LogErrorIteration        = "Error iterating over rows: %v"
)
