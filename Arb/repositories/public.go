package repositories

import (
	"arb/constants"
	"arb/db"
	"arb/structures"
	"context"
	"log"
)

type PublicRepository interface {
	CreatePublic(public *structures.Public) error
	GetPublicByID(id int) (*structures.Public, error)
	GetAllPublics() ([]*structures.Public, error)
	UpdatePublic(public *structures.Public) error
	DeletePublic(id int) error
}

type PgPublicRepository struct{}

func NewPublicRepository() PublicRepository {
	return &PgPublicRepository{}
}

func (repo *PgPublicRepository) CreatePublic(public *structures.Public) error {
	query := `
		INSERT INTO publics (
			name, tag, contacts, topic, subscriber_price, ad_price, 
			wants_op, description, is_selling, monthly_users, sale_price, 
			is_scammer, is_verified
			)
		VALEUS ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	log.Println("DB: Adding public to datebase..")

	err := db.DatabasePool.QueryRow(
		context.Background(), query, public.Name, public.Tag, public.Owner,
		public.Contacts, public.Topic, public.SubcriberPrice, public.AdPrice,
		public.WantsOP, public.Description, public.IsSelling, public.MonthlyUsers,
		public.SalePrice, public.IsScammer, public.IsVerified,
	)

	if err != nil {
		log.Printf("DB: "+constants.LogErrorAddingPublic, err)
	}

	return nil
}

func (repo *PgPublicRepository) GetPublicByID(id int) (*structures.Public, error) {
	public := &structures.Public{}
	query := `
		SELECT * FROM publics WHERE id = $1
	`

	log.Println("DB: Get public from database..")
	err := db.DatabasePool.QueryRow(context.Background(), query, id).Scan(
		&public.ID, &public.Name, &public.Tag, &public.Owner, &public.Contacts,
		&public.Topic, &public.SubcriberPrice, &public.AdPrice,
		&public.WantsOP, &public.Description, &public.IsSelling,
		&public.MonthlyUsers, &public.SalePrice, &public.IsScammer,
	)

	if err != nil {
		log.Printf("DB: "+constants.LogErrorFetchingPublic, err)
		return nil, err
	}

	return public, nil
}

func (repo *PgPublicRepository) GetAllPublics() ([]*structures.Public, error) {
	query := `
		SELECT * FROM publics
	`
	log.Printf("DB: Get all publics from database...")

	rows, err := db.DatabasePool.Query(context.Background(), query)
	if err != nil {
		log.Printf("DB: "+constants.LogErrorFetchingAllPublics, err)
		return nil, err
	}
	defer rows.Close()

	var publics []*structures.Public

	for rows.Next() {
		public := &structures.Public{}

		err := rows.Scan(
			&public.ID, &public.Name, &public.Tag, &public.Owner, &public.Contacts,
			&public.Topic, &public.SubcriberPrice, &public.AdPrice,
			&public.WantsOP, &public.Description, &public.IsSelling,
			&public.MonthlyUsers, &public.SalePrice, &public.IsScammer,
		)
		if err != nil {
			log.Printf(constants.LogErrorScaningPublic, err)
			continue
		}

		publics = append(publics, public)
	}
	return publics, nil
}

func (repo *PgPublicRepository) UpdatePublic(public *structures.Public) error {
	query := `
		UPDATE publics SET
			name = $1, tag = &2, owner = &3, contacts = &4, topic = &5,
			subscriber_price = &6, ad_price = &7, wants_op = &8,
			description = &9, is_selling = &10, monthly_users = &11,
			sale_price = &12, is_scammer = &13
		WHERE id = &14
	`
	log.Printf("DB: Updating public from database...")

	_, err := db.DatabasePool.Exec(context.Background(), query,
		&public.Name, &public.Tag, &public.Owner, &public.Contacts,
		&public.Topic, &public.SubcriberPrice, &public.AdPrice,
		&public.WantsOP, &public.Description, &public.IsSelling,
		&public.MonthlyUsers, &public.SalePrice, &public.IsScammer, &public.ID,
	)
	if err != nil {
		log.Printf(constants.LogErrorUpdatingPublic, err)
		return err
	}
	return nil
}

func (repo *PgPublicRepository) DeletePublic(id int) error {
	query := `
		DELETE FROM publics WHERE id = $1
	`
	log.Println("DB: Deleting public from database...")

	_, err := db.DatabasePool.Exec(context.Background(), query, id)
	if err != nil {
		log.Printf(constants.LogErrorDeletingPublic, err)
		return err
	}

	return nil
}
