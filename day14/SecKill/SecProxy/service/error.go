package service

const (
	ErrInvalidRequest = iota + 1001
	ErrNotFoundProductId
	ErrUserCheckAuthFailed
	ErrUserServiceBusy
	ErrActiveNotStart
	ErrActiveAlreadyEnd
	ErrActiveSaleOut
	ErrProcessTimeout
	ErrClientClosed
)
