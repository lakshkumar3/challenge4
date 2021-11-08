package redis

//
//import (
//    "challenge1/entity"
//    "encoding/json"
//    "fmt"
//    "github.com/cakemarketing/snowbank/stores"
//    "github.com/go-redis/redis"
//)
//var client *redis.Client
//
//type Redis struct {
//     *redis.Client
//}
//
//func GetRedisDBInstance() *Redis {
//    if client == nil {
//       client= redis.NewClient(&redis.Options{
//            Addr: "localhost:6379",
//            Password: "",
//            DB: 0,
//        })
//        return &Redis{ client}
//    }
//    return &Redis{ client}
//}
//func (obj Redis) SetValues(name string,equation entity.Equation) error {
//    if obj.IsUserExists(name) {
//        oldEquationsJson,_:= obj.Get(name).Result()
//    var oldEquations entity.Equation
//      err:=  json.Unmarshal([]byte(oldEquationsJson),&oldEquations)
//        if err != nil {
//        return err
//        }
//        oldEquations.Equations = append(oldEquations.Equations, equation.Equations...)
//        allEquationJson,err:=json.Marshal(oldEquations)
//        if err != nil {
//        }
//        obj.Set(name,allEquationJson,0)
//    } else {
//        allEquationJson,err:=json.Marshal(equation)
//        if err != nil {
//            return err
//        }
//        obj.Set(name,allEquationJson,0)
//    }
//    return nil
//}
//func (obj Redis) IsUserExists(name string) bool  {
//    _,err:= obj.Get(name).Result()
//    if err != nil {
//    return false
//    }
//    return true
//}
//
//
//// Healthy returns an error if it fails to make a request to AWS
//func (r Redis) Healthy() error {
//    status, err := r.Ping().Result()
//    if err != nil {
//        return err
//    }
//
//    if status != "PONG" {
//        return fmt.Errorf("%s != PONG", status)
//    }
//
//    return nil
//}
//
////
//func (db *Redis) GetModuleType() stores.ModuleType {
//    return stores.ModuleType_DB
//}
//
////
//func (db *Redis) GetDatabaseType() stores.DatabaseType {
//    return stores.DatabaseType_Redis
//}
//
