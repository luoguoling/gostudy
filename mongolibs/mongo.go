package mongolibs

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)
type Monitor struct {
	Msg string	`bson:"msg"`
}
var monitor = Monitor{Msg:"rolin"}
func OpMongo()  {
	fmt.Println("开始操作mongo")
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://119.45.29.189:27017", Database: "monitor", Coll: "monitor_log"})
	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
	one := Monitor{}
	fmt.Println(cli.GetDatabaseName(),cli.Find(ctx,bson.M{"msg":monitor.Msg}).One(&one))


}

