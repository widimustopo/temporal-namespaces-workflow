package libs

const (
	OK                            = "OK"
	OperationSuccessfullyExecuted = "Operation Successfully Executed"
	SomethingWentWrong            = "Oops, Something Went Wrong"
	ValidationError               = "Validation Error"
	Unauthorized                  = "Unauthorized"
	UnprocessableEntity           = "Unprocessable Entity"
	BadRequest                    = "Bad Request"
	Forbidden                     = "Forbidden"
	InternalServerError           = "Internal Server Error"
	ServiceIsUnavailable          = "The Service Is Unavailable"
	ServiceIsNotAccessible        = "We are Sorry, The Service Is Not Available Right Now"
	Success                       = "Success"
	NotFound                      = "Not Found"
	DeleteSuccess                 = "Delete Success"

	//Temporal
	RegisterWorkflow = "RegisterWorkflow"
	//Register Member
	ActivityRegisterMember = "Register"
	RegisterMemberQueue    = "REG-MEMBER-QUEUE"
	//Payment
	ActivityPayment = "Payment"
	PaymentQueue    = "PAYMENT-QUEUE"

	MaxConcurrentSquareActivitySize = 10
)
