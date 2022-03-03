package constant

//ContextKey see this link for possible alternative: https://blog.golang.org/context#TOC_3.2.
type ContextKey string

const (
	ContextAction                     ContextKey = "request-action"
	ContextUUID                       ContextKey = "request-uuid"
	ContextUserAgent                  ContextKey = "client-user-agent"
	ContextUserIP                     ContextKey = "client-net-ip"
	ContextDeviceID                   ContextKey = "client-device-id"
	ContextSessionID                  ContextKey = "session-id"
	ContextChannelID                  ContextKey = "channel-id"
	ContextClientID                   ContextKey = "client-id"
	ContextRouterParam                ContextKey = "router-params"
	ContextReferenceUUID              ContextKey = "reference-request-uuid"
	ContextMessageID                  ContextKey = "nsq-message-id"
	ContextBirthTime                  ContextKey = "birth-time"
	ContextUserID                     ContextKey = "user-id"
	ContextCIFID                      ContextKey = "cif-id"
	ContextApplicationStatus          ContextKey = "application-status"
	ContextExistingApplicationInvoice ContextKey = "existing-application-invoice"
	ContextExternalUserID             ContextKey = "external-user-id"
	ContextExternalApplicationID      ContextKey = "external-application-id"
	ContextTime                       ContextKey = "current-time"
	ContextUserData                   ContextKey = "user-data"
	ContextParent                     ContextKey = "parent-context"
)
