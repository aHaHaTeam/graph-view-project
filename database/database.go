package database

import "graph-view-project/models"

type DataBase interface {
	Connect(databaseName string) error
	Close() error
	CreateUser(user models.User) error
	CreateGraph(user models.User, graph models.Graph) error
	CreateNode(graph models.Graph, node models.Node) error
	CreateEdge(graph models.Graph, edge models.Edge) error

	GetUserByLogin(login string) (*models.User, error)
	GetGraphById(id int) (*models.Graph, error)
	GetEdgeById(id int) (*models.Edge, error)
	GetNodeById(id int) (*models.Node, error)
}
