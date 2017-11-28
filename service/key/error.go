package key

import "github.com/giantswarm/microerror"

var wrongTypeError = microerror.New("wrong type")

// IsWrongTypeError asserts wrongTypeError.
func IsWrongTypeError(err error) bool {
	return microerror.Cause(err) == wrongTypeError
}

var malformedCloudConfigKeyError = microerror.New("malformed key in the cloudconfig")

// IsMalformedCloudConfigKey asserts malformedCloudConfigKeyError.
func IsMalformedCloudConfigKey(err error) bool {
	return microerror.Cause(err) == malformedCloudConfigKeyError
}

var missingCloudConfigKeyError = microerror.New("missing key in the cloudconfig")

// IsMissingCloudConfigKey asserts missingCloudConfigKeyError.
func IsMissingCloudConfigKey(err error) bool {
	return microerror.Cause(err) == missingCloudConfigKeyError
}
