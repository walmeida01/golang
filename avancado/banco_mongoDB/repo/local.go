package repo

import (
	"github.com/walmeida01/golang/avancado/banco_mongoDB/model"
	"gopkg.in/mgo.v2/bson"
)

//ObtemLocal retorna um local do MongoDB
func ObtemLocal(codigoTelefone int) (local model.Local, err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	collections := sessao.DB("auladego").C("local")
	err = collections.Find(bson.M{"telcode": codigoTelefone}).One(&local)
	return
}
