package version

import (
	"context"
	"errors"
	"goodfs/apiserver/model/meta"
	"goodfs/apiserver/repository"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ErrVersion int32 = -1
)

//Find 根据hash查找版本，返回版本以及版本号
func Find(hash string) (*meta.MetaVersion, int32) {
	collection := repository.GetMongo().Collection("metadata")
	res := struct {
		Index    int32 `bson:"index"`
		versions []*meta.MetaVersion
	}{}
	if e := collection.FindOne(
		nil,
		bson.M{"versions.hash": hash},
		options.FindOne().SetProjection(bson.M{
			"index":      bson.M{"$indexOfArray": bson.A{"$versions.hash", hash}},
			"versions.$": 1,
		}),
	).Decode(&res); e != nil {
		log.Println(e)
		return nil, ErrVersion
	}

	if res.Index > ErrVersion {
		return res.versions[0], res.Index
	} else {
		return nil, res.Index
	}
}

//Add 为metadata添加一个版本，添加到版本数组的末尾，版本号为数组序号
//返回对应版本号,如果失败返回ErrVersion -1
func Add(ctx context.Context, id string, ver *meta.MetaVersion) int32 {
	if ctx == nil {
		ctx = context.Background()
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("id error %v", id)
		return ErrVersion
	}

	ver.Ts = time.Now()
	collection := repository.GetMongo().Collection("metadata")
	data := struct {
		LenOfVersion int32 `bson:"lenOfVersion"`
	}{}

	//returns the pre-modified version of the document
	err = collection.FindOneAndUpdate(ctx, bson.M{
		"_id": oid,
	}, bson.M{
		"$set": bson.M{
			"update_time": time.Now(),
		},
		"$push": bson.M{
			"versions": ver,
		},
	}, options.FindOneAndUpdate().SetProjection(bson.M{
		"lenOfVersion": bson.M{"$size": "$versions"},
		"_id":          0,
	})).Decode(&data)

	if err != nil {
		log.Println(err)
		return ErrVersion
	}

	return data.LenOfVersion
}

func Delete(ctx context.Context, id string, ver *meta.MetaVersion) error {
	if ctx == nil {
		ctx = context.Background()
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("id error %v", id)
		return nil
	}

	collection := repository.GetMongo().Collection("metadata")

	res, err := collection.UpdateOne(ctx, bson.M{
		"_id": oid,
	}, bson.M{
		"$set": bson.M{
			"versions.$[elem].hash": "",
			"versions.$[elem].size": 0,
		},
	}, options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.M{
			"elem.hash": ver.Hash,
		}},
	}).SetHint("metadata_versions_hash"))

	if err != nil {
		return err
	} else if res.ModifiedCount == 0 {
		return errors.New("Delete fail")
	}

	return nil
}