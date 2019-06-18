package common

import "gopkg.in/mgo.v2"

func ConnectToMongo()(*mgo.Session, error){
	session, err := mgo.Dial("127.0.0.1:27017")
	//session , err are new variables ,,, *mgo.session is type of session and error is type of err
	if err!=nil{
		return nil, err
	}else{
		return session,nil
	}
}