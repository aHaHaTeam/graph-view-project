package database

import "graph-view-project/models"

type DataBase interface {
	Connect(databaseName string) error
	Disconnect() error

	CreateUser(user models.User) (*models.User, error)
	CreateGraph(user models.User, graph models.Graph) (*models.Graph, error)
	CreateNode(graph models.Graph, node models.Node) (*models.Node, error)
	CreateEdge(graph models.Graph, edge models.Edge) (*models.Edge, error)

	GetUser(id int) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
	GetGraph(id int) (*models.Graph, error)
	GetEdge(id int) (*models.Edge, error)
	GetNode(id int) (*models.Node, error)

	UpdateUser(newUser models.User) error
	UpdateGraph(newGraph models.Graph) error
	UpdateEdge(newEdge models.Edge) error
	UpdateNode(newNode models.Node) error
}
