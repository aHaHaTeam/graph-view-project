package database

import "graph-view-project/wasm/models"

type Database interface {
	Connect() error
	Disconnect() error

	CreateUser(user models.User) error
	GetUserByID(id int) (models.User, error)
	SaveUser(user models.User) error

	CreateGraph(graph models.Graph) error
	GetGraphByID(id int) (models.Graph, error)
	SaveGraph(graph models.Graph) error

	CreateNode(node models.Node) error
	GetNodeByID(id int) (models.Node, error)
	SaveNode(node models.Node) error

	CreateEdge(user models.User) error
	GetEdgeByID(id int) (models.User, error)
	SaveEdge(user models.User) error
}
