package impl

func newPostGreSQLDatastore(conf map[string]string) Datastore {
	return &PostgreSQLDataStore{
		DSN: conf["dsn"],
		//create db as well
	}
}

func newMemoryDatastore(conf map[string]string) Datastore {
	return &MemoryDataStore{
		dns: "",
	}
}

type DataStoreFactory func(conf map[string]string) Datastore

var datastoreFactories = make(map[string]DataStoreFactory)

func Register(engineName string, dataStoreFactory DataStoreFactory) {
	datastoreFactories[engineName] = dataStoreFactory
}

func init() {
	Register("memory", newMemoryDatastore)
	Register("postgre", newPostGreSQLDatastore)
}

func CreateDatastore(conf map[string]string) Datastore {
	engineName := conf["engine"]
	engineFactory := datastoreFactories[engineName]
	return engineFactory(conf)
}
