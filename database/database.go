package database

import "graph-view-project/models"

type DataBase interface {
	Connect(databaseName string) error
	Disconnect() error
	CompleteReset() error

	CreateUser(user models.User) (*models.User, error)
	CreateGraph(user models.User, graph models.Graph) (*models.Graph, error)
	CreateNode(graph models.Graph, node models.Node) (*models.Node, error)
	CreateEdge(graph models.Graph, edge models.Edge) (*models.Edge, error)

	GetUser(id models.UserId) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
	GetGraph(id models.GraphId) (*models.Graph, error)
	GetEdge(id models.EdgeId) (*models.Edge, error)
	GetNode(id models.NodeId) (*models.Node, error)

	UpdateUser(newUser models.User) error
	UpdateGraph(newGraph models.Graph) error
	UpdateEdge(newEdge models.Edge) error
	UpdateNode(newNode models.Node) error
}
