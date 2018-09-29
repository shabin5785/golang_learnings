package impl

type MemoryDataStore struct {
	dns   string
	Users map[int64]string
}

func (mds *MemoryDataStore) Name() string{
	return "MemoryDataStore"
}

func(mds *MemoryDataStore) FindUserById(id int64) (string,error){
	//get user name
	return "",nil
}
