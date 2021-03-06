package reflect

// Different operations to match parameters with.
type ParameterOperation string

const (
	EqualsOperation              = ParameterOperation("=")
	NotEqualsOperation           = ParameterOperation("!=")
	GreaterThanOperation         = ParameterOperation(">")
	GreaterThanOrEqualsOperation = ParameterOperation(">=")
	LessThanOperation            = ParameterOperation("<")
	LessThanOrEqualsOperation    = ParameterOperation("<=")
	ContainsOperation            = ParameterOperation("=~")
	NotContainsOperation         = ParameterOperation("!~")
)

// A parameter you want to include when generating signed authentication tokens
// for your clients.
type Parameter struct {
	// The name of the field this parameter applies to.
	Field string `json:"field"`

	// The operation to apply to this field and value.
	Op ParameterOperation `json:"op"`

	// The value to compare against.
	Value string `json:"value"`

	// A slice of possible values to compare against, any of which will be
	// considered in the operation specified.
	AnyValues []string `json:"any"`
}
