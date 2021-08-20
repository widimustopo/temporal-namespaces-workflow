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

	//Temporal Workflow
	RegisterWorkflow    = "RegisterWorkflow"
	PaymentWorkflow     = "PaymentWorkflow"
	OrderWorkflow       = "OrderWorkflow"
	ExpiredWorkflow     = "ExpiredWorkflow"
	PaymentFailWorkflow = "PaymentFailWorkflow"

	//Register Member
	ActivityRegisterMember = "Register"
	RegisterMemberQueue    = "REG-MEMBER-QUEUE"

	//Payment
	ActivityPayment = "Payment"
	ActivityOrder   = "Order"
	PaymentQueue    = "ORDER-QUEUE"

	//Expired
	ActivityExpired = "Expired"
	ExpiredQueue    = "EXPIRED-QUEUE"

	//PaymentFail
	ActivityPaymentFail = "PaymentFail"
	PaymentFailQueue    = "PAYMENT-FAIL-QUEUE"

	MaxConcurrentSquareActivitySize = 10
)
