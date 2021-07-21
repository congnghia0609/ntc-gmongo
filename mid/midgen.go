//Author: nghiatc
//Since: Jul 15, 2021

package mid

import (
	"context"
	"fmt"
	"log"

	"github.com/congnghia0609/ntc-gmongo/gmongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const tableName = "midgen"

type MIdGen struct {
	Id  string `bson:"_id" json:"id"`
	Seq int64  `bson:"seq" json:"seq"`
}

func (id MIdGen) String() string {
	return fmt.Sprintf("MIdGen{Id: %v, Seq: %v}", id.Id, id.Seq)
}

// Next generate a auto increment version ID for the given key
func GetNext(name string) (int64, error) {
	var idgen MIdGen
	filter := bson.D{{"_id", name}}
	update := bson.D{{"$inc", bson.D{{"seq", int64(1)}}}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{ReturnDocument: &after, Upsert: &upsert}

	collection := gmongo.TestDB.Collection(tableName)
	err := collection.FindOneAndUpdate(context.Background(), filter, update, &opt).Decode(&idgen)
	// log.Println(err)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 1, nil
		} else {
			log.Println(err)
			return 0, err
		}
	}

	// jidgen, _ := json.Marshal(idgen)
	// fmt.Printf("idgen = %s\n", string(jidgen))
	// fmt.Println(idgen)

	return idgen.Seq, nil
}

// ResetID reset id gen to value
func ResetID(id string, value int64) (int64, error) {
	collection := gmongo.TestDB.Collection(tableName)

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"seq", value}}}}

	result, err := collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return 0, err
	}
	// log.Println("nid:", nid)

	return result.MatchedCount, nil
}
