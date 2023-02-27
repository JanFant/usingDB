package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func MongoExamples(client *mongo.Client) {
	coll := client.Database("test").Collection("posts")

	//update := bson.D{{"$inc", bson.D{{"likes", 5}}}}
	//res, err := coll.UpdateOne(context.TODO(), bson.M{"category": "News"}, update)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)

	cur, err := coll.Find(context.TODO(), bson.M{}, options.Find())

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var posts []Post
	if err := cur.All(context.TODO(), &posts); err != nil {
		fmt.Println(err)
	}

	fmt.Println(posts)

	cur.Close(context.TODO())

}

type likes struct {
	coll int `bson:"likes"`
}

type Post struct {
	ID       primitive.ObjectID `bson:"_id"`
	Title    string             `bson:"title"`
	Category string             `bson:"category"`
	Date     string             `bson:"date"`
	Likes    int                `bson:"likes"`
}

type Article struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	CDate    time.Time          `bson:"crateDay"`
	Content  string             `bson:"content"`
	Comments []Comments         `bson:"comments"`
}

type Comments struct {
	Name    string `bson:"name"`
	Content string `bson:"content"`
}
