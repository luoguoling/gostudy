package mongolibs

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"reflect"

	//"golang_demo/mongo/db"
	"log"
	"time"
)
func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error)  {
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// 通过传进来的uri连接相关的配置
	o := options.Client().ApplyURI(uri)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	o.SetMaxPoolSize(num)
	// 发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 返回 client
	return client.Database(name), nil
}
// 数据结构体
type Student struct {
	Id int32	`json:"id"`
	Name string `json:"name"`
	Age int	`json:"age"`
}
type Resp struct {

	Code int `json:"code"`
	Data []Student `json:"data"`
}
var (
	opt = "mongodb://119.45.29.189:27017" //  带账号名的链接
	name = "user1" // 数据库名
	maxTime = time.Duration(2) // 链接超时时间
	num uint64 = 50 // 链接数
	table = "student" // 表名
	toDB *mongo.Database // database 话柄
	collection *mongo.Collection // collection 话柄
)

func init()  {
	var err error
	toDB, err = ConnectToDB(opt, name,maxTime,num)
	if err!= nil {
		panic("链接数据库有误!")
	}
	collection = toDB.Collection(table)
}

// GetList 获取全量的数据
func GetList()  {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var all []Student

	//all := make([]Student,0)
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())

	log.Println("collection.Find curl.All: ", all)
	fmt.Println(reflect.TypeOf(all))
	result := Resp{1,all}
	fmt.Println(result.Code,result.Data)
	res ,_ := json.Marshal(result)
	fmt.Println("res is ...." )
	fmt.Println("结果是:",res)
	//for _, one := range all {
	//	log.Println("Id:",one.Id," - name:",one.Name," - age:",one.Age)
	//}
	//return res
}

// AddOne 新增一条数据
func AddOne(s1 *Student)  {
	objId, err := collection.InsertOne(context.TODO(), &s1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("录入数据成功，objId:",objId)
}

// EditOne 编辑一条数据
func EditOne(student *Student,m bson.M)  {
	update := bson.M{"$set": student}
	updateResult, err := collection.UpdateOne(context.Background(),  m, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
}

// 更新数据 - 存在更新，不存在就新增
func Update(student *Student,m bson.M)  {
	update := bson.M{"$set": student}
	updateOpts := options.Update().SetUpsert(true)
	updateResult, err := collection.UpdateOne(context.Background(), m, update, updateOpts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
}

// 删除一条数据
func Del(m bson.M)  {
	deleteResult, err := collection.DeleteOne(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.DeleteOne:", deleteResult)
}

// Sectle 模糊查询
// bson.M{"name": primitive.Regex{Pattern: "深入"}}
func Sectle(m bson.M)  {
	cur, err := collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var s Student
		if err = cur.Decode(&s); err != nil {
			log.Fatal(err)
		}
		log.Println("collection.Find name=primitive.Regex{xx}: ", s)
	}
	cur.Close(context.Background())
}

// 统计collection的数据总数
func Count()  {
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(count)
	}
	log.Println("collection.CountDocuments:", count)
}

// 准确搜索一条数据
func GetOne(m bson.M)  {
	var one Student
	err := collection.FindOne(context.Background(), m).Decode(&one)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.FindOne: ", one)
}