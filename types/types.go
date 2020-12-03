package types

// // Entity interface
// type Entity interface {
// 	GetInsertQuery(schemaName string) (iq *InsertQuery)
// }

// func CreateEntity(typeName string) (e Entity, err error) {
// 	switch typeName {
// 	case "Badges":
// 		e = &Badge{}
// 	case "Comments":
// 		e = &Comment{}
// 	case "Posts":
// 		e = &Post{}
// 	case "PostLinks":
// 		e = &PostLink{}
// 	case "PostHistory":
// 		e = &PostHistory{}
// 	case "Tags":
// 		e = &Tag{}
// 	case "Users":
// 		e = &User{}
// 	case "Votes":
// 		e = &Vote{}
// 	default:
// 		err = fmt.Errorf("Undefined Entity Name: %s", entrp.entityName)
// 	}
// 	return
// }
