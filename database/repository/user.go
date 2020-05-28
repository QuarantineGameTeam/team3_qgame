package repository

import (
	"gihub.com/team3_qgame/database/transaction"
	"github.com/google/uuid"
)

type UserRepository struct {
	manager *transaction.DbContextManager
}

const (
	getAllItems = "SELECT id, train_id, plane_id, user_id, place, " +
		"ticket_type," +
		" discount, price, total_price, name, surname FROM tickets;"
	getOneItem = "SELECT id, train_id, plane_id, user_id, place, " +
		"ticket_type," +
		" discount, price, total_price, name, " +
		"surname FROM tickets WHERE id = $1;"
	addOneItem = "INSERT INTO tickets (id, place, ticket_type, discount, " +
		"price, total_price, name, surname) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	updateItem = "UPDATE tickets SET place=$2, ticket_type=$3, " +
		"discount=$4, " +
		"price=$5, total_price=$6, name=$7, surname=$8 WHERE id=$1;"
	deleteItem = "DELETE FROM tickets WHERE id=$1;"
)

func NewUserRepository(manager *transaction.DbContextManager) *UserRepositoryRepository {
	return &ProjectRepository{
		manager: manager,
	}
}

func (p *ProjectRepository) NewUser(ctx context.Context, project *model.Project) (*model.Project, error) {
	tr, _ := p.manager.GetTransactionContext(ctx)

	project.ProjectID = uuid.New().String()

	err := tr.Provider().Create(project).Error

	if err != nil {
		return nil, err
	}

	return project, nil
}

//CreateTicket sends a query for creating new one ticket
func (*ticketRepository) CreateTicket(tk data.Ticket) error {
	_, err := Db.Exec(addOneItem, tk.ID, tk.Place, tk.TicketType, tk.Discount,
		tk.Price, tk.TotalPrice, tk.Name, tk.Surname)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectRepository) GetByID(ctx context.Context, id uint) (*model.Project, error) {
	tr, _ := p.manager.GetTransactionContext(ctx)

	var projects []*model.Project
	err := tr.Provider().First(&projects).Where("id", id).Error

	if err != nil {
		return nil, err
	}
	if len(projects) == 0 {
		return nil, errors.NotFoundError
	}
	return projects[0], nil
}
