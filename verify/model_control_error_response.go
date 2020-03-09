/*
 * Nexmo Verify API
 *
 * The Verify API helps you to implement 2FA (two-factor authentication) in your applications. This is useful for:  * Protecting against spam, by preventing spammers from creating multiple accounts * Monitoring suspicious activity, by forcing an account user to verify ownership of a number * Ensuring that you can reach your users at any time because you have their correct phone number More information is available at <https://developer.nexmo.com/verify>
 *
 * API version: 1.0.11
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package verify
// ControlErrorResponse Error
type ControlErrorResponse struct {
	// Code | Text | Description -- | -- | -- 0 | Success | The request was successfully accepted by Nexmo. 1 | Throttled | You are trying to send more than the maximum of 30 requests per second. 2 | Your request is incomplete and missing the mandatory parameter `$parameter` | The stated parameter is missing. 3 | Invalid value for parameter `$parameter` | Invalid value for parameter. If you see Facility not allowed in the error text, check that you are using the correct Base URL in your request. 4 | Invalid credentials were provided | The supplied API key or secret in the request is either invalid or disabled. 5 | Internal Error | An error occurred processing this request in the Cloud Communications Platform. 6 | The Nexmo platform was unable to process this message for the following reason: `$reason` | The request could not be routed. 8 | The api_key you supplied is for an account that has been barred from submitting messages. | 9 | Partner quota exceeded | Your account does not have sufficient credit to process this request. 19 | For `cancel`: Either you have not waited at least 30 secs after sending a Verify request before cancelling or Verify has made too many attempts to deliver the verification code for this request and you must now wait for the process to complete. For `trigger_next_event`: All attempts to deliver the verification code for this request have completed and there are no remaining events to advance to. 101 | No request found | There are no matching verify requests. 
	Status string `json:"status,omitempty"`
	// If the `status` is non-zero, this explains the error encountered.
	ErrorText string `json:"error_text,omitempty"`
}
