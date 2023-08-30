package logrus_rollbar

import (
	"context"
	"net/http"
)

// SetContext
func (hook *RollbarHook) SetContext(ctx context.Context) {
	hook.client.SetContext(ctx)
}

// SetToken
func (hook *RollbarHook) SetToken(token string) {
	hook.client.SetToken(token)
}

// SetEnvironment
func (hook *RollbarHook) SetEnvironment(environment string) {
	hook.client.SetEnvironment(environment)
}

// SetItemsPerMinute
func (hook *RollbarHook) SetItemsPerMinute(itemsPerMinute int) {
	hook.client.SetItemsPerMinute(itemsPerMinute)
}

// SetPlatform
func (hook *RollbarHook) SetPlatform(platform string) {
	hook.client.SetPlatform(platform)
}

// SetCodeVersion
func (hook *RollbarHook) SetCodeVersion(codeVersion string) {
	hook.client.SetCodeVersion(codeVersion)
}

// SetServerHost
func (hook *RollbarHook) SetServerHost(serverHost string) {
	hook.client.SetServerHost(serverHost)
}

// SetServerRoot
func (hook *RollbarHook) SetServerRoot(serverRoot string) {
	hook.client.SetServerRoot(serverRoot)
}

// SetCustom
func (hook *RollbarHook) SetCustom(custom map[string]interface{}) {
	hook.client.SetCustom(custom)
}

// SetPerson
func (hook *RollbarHook) SetPerson(id, username, email string) {
	hook.client.SetPerson(id, username, email)
}

// ClearPerson
func (hook *RollbarHook) ClearPerson() {
	hook.client.ClearPerson()
}

// SetFingerprint
func (hook *RollbarHook) SetFingerprint(fingerprint bool) {
	hook.client.SetFingerprint(fingerprint)
}

// SetRetryAttempts
func (hook *RollbarHook) SetRetryAttempts(retryAttempts int) {
	hook.client.SetRetryAttempts(retryAttempts)
}

// SetPrintPayloadOnError
func (hook *RollbarHook) SetPrintPayloadOnError(printPayloadOnError bool) {
	hook.client.SetPrintPayloadOnError(printPayloadOnError)
}

// SetHTTPClient
func (hook *RollbarHook) SetHTTPClient(httpClient *http.Client) {
	hook.client.SetHTTPClient(httpClient)
}
