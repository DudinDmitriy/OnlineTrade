
package clients

type clientID int


type clientData struct {
	id clientID
	name string
	email string
}

type listClients struc{
	list map[clientID]clientData
	lastID clientID
}

func (cd *listClients) Init(){
	// Load list clients from storage

}

func (cd *listClients) Save(){

}

func (cd *listClients) Add(){
	
}

