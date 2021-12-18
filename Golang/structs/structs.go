package structs

type Prueba struct {
	Nombre string `bson:"nombre,omitempty"`
	Edad   int    `bson:"edad,omitempty"`
}

type Persona struct {
	Name         string `bson:"name,omitempty"`
	Location     string `bson:"location,omitempty"`
	Age          int    `bson:"age,omitempty"`
	Vaccine_type string `bson:"vaccine_type,omitempty"`
	N_dose       int    `bson:"n_dose,omitempty"`
}
