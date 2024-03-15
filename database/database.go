package database

import "graph-view-project/models"

type DataBase interface {
	Connect(databaseName string) error
	Disconnect() error

	CreateUser(user models.User) error
	CreateGraph(graph models.Graph) error
	CreateNode(node models.Node) error
	CreateEdge(edge models.Edge) error

	GetUserByLogin(login string) (*models.User, error)
	GetGraphById(id int) (*models.Graph, error)
	GetEdgeById(id int) (*models.Edge, error)
	GetNodeById(id int) (*models.Node, error)

	UpdateUserByLogin(login string, newUser models.User) error
	UpdateGraphById(id int, newGraph models.Graph) error
	UpdateEdgeById(id int, newEdge models.Edge) error
	UpdateNodeById(id int, newNode models.Node) error
}
