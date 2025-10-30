package exception

/*
Pacote para padronizar erros da API.
A 'Exception' será enviada como resposta caso
um arequisição falhar.
*/
type Exception struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func (e Exception) Error() string {
	return e.Message
}

func New(message string) error {
	return &Exception{
		Message: message, 
	}
}

func WithStatus(status int, message string) error {
	return &Exception{
		Message: message,
		Status: status,
	}	
}