package services

import(
	"fmt"
	"github.com/techswarn/flowdocapi/models"

    "time"

    "github.com/techswarn/flowdocapi/database"
	"github.com/google/uuid"
)

func CreateNode(nodeRequest models.NodeRequest) (models.Node, error) {

	var node models.Node = models.Node{
		ID : uuid.New().String(),
		NodeType: nodeRequest.NodeType,
		Label: nodeRequest.Label,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if nodeResult := database.DB.Create(&node); nodeResult.Error != nil {
		fmt.Printf("DB write error: %s", &nodeResult.Error)
		return node, nodeResult.Error
	}

	//Check if Apex parent node is present if not present set source as target for parent apex nodes
	var Nodes []models.Node = []models.Node{}
	database.DB.Order("created_at asc").Find(&Nodes)
	fmt.Printf("%d", len(Nodes))
	var source string 
	if len(Nodes) == 1 {
		source = node.ID
	} else {
		fmt.Println("NODE LENGTH GREATER THAN 1")
		source = nodeRequest.Source
		fmt.Printf("%s \n", nodeRequest.Source)
	}

    var article models.Article = models.Article{
		NodeID: node.ID,
		Heading: nodeRequest.Heading,
		Description: nodeRequest.Description,
		Error: nodeRequest.Error,
	}

    if articleResult := database.DB.Create(&article); articleResult.Error != nil {
		fmt.Printf("DB write error: %s", &articleResult.Error)
	}
	fmt.Printf("add article: %#v", article)
	fmt.Printf("Links from node request: %#v",nodeRequest.Links)

	for _, v := range nodeRequest.Links{
		var link models.Url = models.Url{
			ArticleID: article.ID,
			Label: v.LinkLabel,
			Link: v.Url,
		}
		if urlResult := database.DB.Create(&link); urlResult.Error != nil {
			fmt.Printf("DB write error: %s", &urlResult.Error)
		}
	}
    
	var edge models.Edge = models.Edge{
		ID : uuid.New().String(),
		Source : source,
		Target : node.ID,			
		Type :	nodeRequest.EdgeType,	
		Animated :  nodeRequest.Animated,   
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),    
	}
	

	if edgeResult := database.DB.Create(&edge); edgeResult.Error != nil {
		fmt.Printf("DB write error: %s", &edgeResult.Error)
		return node, edgeResult.Error
	}

	return node, nil
}

func GetAllNodes() []models.Node {
	// create a variable to store items data
	var Nodes []models.Node = []models.Node{}

	// get all data from the database order by created_at
	database.DB.Order("created_at asc").Find(&Nodes)

	// return all items from the database
	return Nodes
}

func GetAllEdges() []models.Edge {
	// create a variable to store items data
	var Edges []models.Edge = []models.Edge{}

	// get all data from the database order by created_at
	database.DB.Order("created_at asc").Find(&Edges)

	// return all items from the database
	return Edges
}

func GetArticle() []models.Article {
	// create a variable to store items data

	var Article []models.Article = []models.Article{}

	// get all data from the database order by created_at
	// if err := database.DB.Model(&models.Article{}).Select("articles.id, articles.node_id, articles.heading, articles.description, articles.error").Joins("join urls on urls.article_id = articles.id").Where("articles.node_id = ?", "fe5b1fc7-d54a-479b-b65c-9e82717f16d2").Scan(&Article); err != nil {
	// 	fmt.Println(err) 
	// }
	// var urlID int32= Article.ID
	// if err := database.DB.Model(&models.Url{}).Select("urls.id").Joins("join articles on articles.id = urls.article_id").Where("urls.article_id = ?", urlID).Scan(&urls); err != nil {
	// 	fmt.Println(err) 
	// }


    if err := database.DB.Model(&models.Article{}).Preload("Urls").Find(&Article).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Articles: %#v \n", Article)
	// if err := database.DB.Table("articles").Select("articles.id, articles.node_id, articles.heading, articles.description, articles.error, urls.id, urls.label, urls.link").Joins("JOIN urls on urls.article_id = articles.id").Find(&Article).Error; err != nil {
	// 	fmt.Println("Error") 
	// }
	

	// if err := database.DB.Model(&models.Article{}).Select("articles.id, articles.node_id, articles.heading, articles.description, articles.error, urls.id, urls.label, urls.link").Joins("join urls on urls.article_id = articles.id").Scan(&Article); err != nil {
	// 	fmt.Println(err) 
	// }

	// return all items from the database
	return Article
}