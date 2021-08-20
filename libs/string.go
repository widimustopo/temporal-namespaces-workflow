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
	RegisterWorkflow       = "RegisterWorkflow"
	PaymentWorkflow        = "PaymentWorkflow"
	OrderWorkflow          = "OrderWorkflow"
	ExpiredWorkflow        = "ExpiredWorkflow"
	PaymentFailWorkflow    = "PaymentFailWorkflow"
	CounterProductWorkflow = "CounterProductWorkflow"
	AddProductWorkflow     = "AddProductWorkflow"

	//Register Member
	ActivityRegisterMember = "Register"

	//Payment
	ActivityPayment = "Payment"
	ActivityOrder   = "Order"
	Counter         = "Counter"

	//Expired
	ActivityExpired = "Expired"

	//PaymentFail
	ActivityPaymentFail = "PaymentFail"

	//Product
	ActivityAddProduct = "AddProduct"

	MaxConcurrentSquareActivitySize = 10
)
