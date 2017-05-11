// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cognitosync provides the client and types for making API
// requests to Amazon Cognito Sync.
//
// Amazon Cognito Sync provides an AWS service and client library that enable
// cross-device syncing of application-related user data. High-level client
// libraries are available for both iOS and Android. You can use these libraries
// to persist data locally so that it's available even if the device is offline.
// Developer credentials don't need to be stored on the mobile device to access
// the service. You can use Amazon Cognito to obtain a normalized user ID and
// credentials. User data is persisted in a dataset that can store up to 1 MB
// of key-value pairs, and you can have up to 20 datasets per user identity.
//
// With Amazon Cognito Sync, the data stored for each identity is accessible
// only to credentials assigned to that identity. In order to use the Cognito
// Sync service, you need to make API calls using credentials retrieved with
// Amazon Cognito Identity service (http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/Welcome.html).
//
// If you want to use Cognito Sync in an Android or iOS application, you will
// probably want to make API calls via the AWS Mobile SDK. To learn more, see
// the Developer Guide for Android (http://docs.aws.amazon.com/mobile/sdkforandroid/developerguide/cognito-sync.html)
// and the Developer Guide for iOS (http://docs.aws.amazon.com/mobile/sdkforios/developerguide/cognito-sync.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30 for more information on this service.
//
// See cognitosync package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cognitosync/
//
// Using the Client
//
// To use the client for Amazon Cognito Sync you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := cognitosync.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Cognito Sync client CognitoSync for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cognitosync/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.BulkPublish(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("BulkPublish result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.BulkPublishWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package cognitosync
