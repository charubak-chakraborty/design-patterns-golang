package chainofresponsibility

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckupDone bool
	MedicineDone      bool
	PaymentDone       bool
}
