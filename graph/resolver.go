package graph

import "github.com/ruancaetano/go-graphql/graph/model"

type Resolver struct {
	Categories []*model.Category
	Courses    []*model.Course
	Chapters   []*model.Chapter
}
