package api

import "arbbot/structures"

func AddTeamToAPI(team structures.Team) error {
	return PostToAPI("http://localhost:8080/api/team", team)
}
