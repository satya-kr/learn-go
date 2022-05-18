package controller

import "go.mongodb.org/mongo-driver/mongo"

//_ =
//public key : rydtului
//public key : 848ec56e-894d-4839-885c-4a7d18ebe959

const connectionString = "mongodb+srv://dbtest7872:dbtest7872@clustergo.c7bnu.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

//most important
var collection *mongo.Collection
