package constants

const (
	//DB errors
	ErrFetchingTeams = "failed to fetch teams"
	ErrAddingTeam    = "failed to add team"
	ErrUpdatingTeam  = "failed to update team"
	ErrDeleteTeam    = "failed to delete team"
	ErrFetchingUsers = "failed to fetch users"
	ErrAddingUser    = "failed to add user"
	ErrUpdatingUser  = "failed to update user"
	ErrDeleteUser    = "failed to update user"

	//Validation errors
	ErrInvalidInput = "invalid input"
	ErrInvalidID    = "invalid id"

	//Server errors
	ErrInternalServerError = "internal serveer error"

	//OTHER
	ErrFailedToExtractFromPath = "failed to extract ID from url"

	//TEAM
	ErrTeamNotFound = "team not found"
	ErrUserNotFound = "user not found"
)
