package repositories

import (
	"arb/db"
	"arb/structures"
	"context"
	"errors"
	"log"
)

type TeamRepository interface {
	AddTeam(team *structures.Team) error
	GetTeamById(id uint) (*structures.Team, error)
	UpdateTeam(team *structures.Team) error
	DeleteTeam(id uint) error
	GetAllTeams() ([]*structures.Team, error)
}

type PgTeamRepository struct{}

func NewPgTeamRepository() TeamRepository {
	return &PgTeamRepository{}
}

func (repo *PgTeamRepository) AddTeam(team *structures.Team) error {
	query := `
		INSERT INTO teams (name, owner, contacts, topic, min_subscriber_price, max_subscriber_price, description, bot_link, is_scammer, team_size, sponsor_count, min_withdrawal_amount, is_verified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id
	`
	log.Printf("Executing SQL: %s with data: %+v", query, team)

	err := db.DatabasePool.QueryRow(
		context.Background(), query,
		team.Name, team.Owner, team.Contacts, team.Topic,
		team.MinSubPrice, team.MaxSubPrice, team.Description, team.BotLink,
		team.IsScummer, team.TeamSize, team.SponsorCount, team.MinWithdrawalAmount, team.IsVerified,
	).Scan(&team.ID)

	if err != nil {
		log.Printf("Error inserting team: %v", err)
	}
	return err
}

func (repo *PgTeamRepository) GetTeamById(id uint) (*structures.Team, error) {
	team := &structures.Team{}
	query := "SELECT id, name, owner, contacts, topic, min_subscriber_price, max_subscriber_price, description, bot_link, is_scammer, team_size, sponsor_count, min_withdrawal_amount FROM teams WHERE id=$1"

	err := db.DatabasePool.QueryRow(context.Background(), query, id).Scan(
		&team.ID, &team.Name, &team.Owner, &team.Contacts, &team.Topic,
		&team.MinSubPrice, &team.MaxSubPrice, &team.Description, &team.BotLink,
		&team.IsScummer, &team.TeamSize, &team.SponsorCount, &team.MinWithdrawalAmount,
	)

	if err != nil {
		log.Printf("Error fetching team with id %d: %v", id, err)
		return nil, errors.New("team not found")
	}
	return team, nil
}

func (repo *PgTeamRepository) UpdateTeam(team *structures.Team) error {
	query := `
		UPDATE teams SET name=$1, owner=$2, contacts=$3, topic=$4, min_subscriber_price=$5, max_subscriber_price=$6, description=$7, bot_link=$8, is_scammer=$9, team_size=$10, sponsor_count=$11, min_withdrawal_amount=$12
		WHERE id=$13
		`
	_, err := db.DatabasePool.Exec(context.Background(), query, team.Name, team.Owner, team.Contacts, team.Topic, team.MinSubPrice, team.MaxSubPrice, team.Description, team.BotLink, team.IsScummer, team.TeamSize, team.SponsorCount, team.MinWithdrawalAmount, team.ID)
	return err
}

func (repo *PgTeamRepository) DeleteTeam(id uint) error {
	query := `
		DELETE FROM teams WHERE id=$1
	`
	_, err := db.DatabasePool.Exec(context.Background(), query, id)
	return err
}

func (repo *PgTeamRepository) GetAllTeams() ([]*structures.Team, error) {
	query := "SELECT id, name, owner, contacts, topic, min_subscriber_price, max_subscriber_price, description, bot_link, is_scammer, team_size, sponsor_count, min_withdrawal_amount FROM teams"

	rows, err := db.DatabasePool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []*structures.Team
	for rows.Next() {
		team := &structures.Team{}

		err := rows.Scan(
			&team.ID, &team.Name, &team.Owner, &team.Contacts, &team.Topic,
			&team.MinSubPrice, &team.MaxSubPrice, &team.Description, &team.BotLink,
			&team.IsScummer, &team.TeamSize, &team.SponsorCount, &team.MinWithdrawalAmount,
		)
		if err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}
	return teams, rows.Err()
}
