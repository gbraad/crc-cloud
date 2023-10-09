// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The S3 object data source allows access to the metadata and
// _optionally_ (see below) content of an object stored inside S3 bucket.
//
// > **Note:** The content of an object (`body` field) is available only for objects which have a human-readable `Content-Type` (`text/*` and `application/json`). This is to prevent printing unsafe characters and potentially downloading large amount of data which would be thrown away in favour of metadata.
//
// ## Example Usage
//
// The following example retrieves a text object (which must have a `Content-Type`
// value starting with `text/`) and uses it as the `userData` for an EC2 instance:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bootstrapScript, err := s3.GetObject(ctx, &s3.GetObjectArgs{
//				Bucket: "ourcorp-deploy-config",
//				Key:    "ec2-bootstrap-script.sh",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewInstance(ctx, "example", &ec2.InstanceArgs{
//				InstanceType: pulumi.String("t2.micro"),
//				Ami:          pulumi.String("ami-2757f631"),
//				UserData:     *pulumi.String(bootstrapScript.Body),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// The following, more-complex example retrieves only the metadata for a zip
// file stored in S3, which is then used to pass the most recent `versionId`
// to AWS Lambda for use as a function implementation. More information about
// Lambda functions is available in the documentation for
// `lambda.Function`.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			lambda, err := s3.GetObject(ctx, &s3.GetObjectArgs{
//				Bucket: "ourcorp-lambda-functions",
//				Key:    "hello-world.zip",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = lambda.NewFunction(ctx, "testLambda", &lambda.FunctionArgs{
//				S3Bucket:        *pulumi.String(lambda.Bucket),
//				S3Key:           *pulumi.String(lambda.Key),
//				S3ObjectVersion: *pulumi.String(lambda.VersionId),
//				Role:            pulumi.Any(aws_iam_role.Iam_for_lambda.Arn),
//				Handler:         pulumi.String("exports.test"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetObject(ctx *pulumi.Context, args *GetObjectArgs, opts ...pulumi.InvokeOption) (*GetObjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetObjectResult
	err := ctx.Invoke("aws:s3/getObject:getObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getObject.
type GetObjectArgs struct {
	// Name of the bucket to read the object from. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
	Bucket string `pulumi:"bucket"`
	// To retrieve the object's checksum, this argument must be `ENABLED`. If you enable `checksumMode` and the object is encrypted with KMS, you must have permission to use the `kms:Decrypt` action. Valid values: `ENABLED`
	ChecksumMode *string `pulumi:"checksumMode"`
	// Full path to the object inside the bucket
	Key   string  `pulumi:"key"`
	Range *string `pulumi:"range"`
	// Map of tags assigned to the object.
	Tags map[string]string `pulumi:"tags"`
	// Specific version ID of the object returned (defaults to latest version)
	VersionId *string `pulumi:"versionId"`
}

// A collection of values returned by getObject.
type GetObjectResult struct {
	// Object data (see **limitations above** to understand cases in which this field is actually available)
	Body   string `pulumi:"body"`
	Bucket string `pulumi:"bucket"`
	// (Optional) Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
	BucketKeyEnabled bool `pulumi:"bucketKeyEnabled"`
	// Caching behavior along the request/reply chain.
	CacheControl string `pulumi:"cacheControl"`
	// The base64-encoded, 32-bit CRC32 checksum of the object.
	ChecksumCrc32 string `pulumi:"checksumCrc32"`
	// The base64-encoded, 32-bit CRC32C checksum of the object.
	ChecksumCrc32c string  `pulumi:"checksumCrc32c"`
	ChecksumMode   *string `pulumi:"checksumMode"`
	// The base64-encoded, 160-bit SHA-1 digest of the object.
	ChecksumSha1 string `pulumi:"checksumSha1"`
	// The base64-encoded, 256-bit SHA-256 digest of the object.
	ChecksumSha256 string `pulumi:"checksumSha256"`
	// Presentational information for the object.
	ContentDisposition string `pulumi:"contentDisposition"`
	// What content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	ContentEncoding string `pulumi:"contentEncoding"`
	// Language the content is in.
	ContentLanguage string `pulumi:"contentLanguage"`
	// Size of the body in bytes.
	ContentLength int `pulumi:"contentLength"`
	// Standard MIME type describing the format of the object data.
	ContentType string `pulumi:"contentType"`
	// [ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)
	Etag string `pulumi:"etag"`
	// If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.
	Expiration string `pulumi:"expiration"`
	// Date and time at which the object is no longer cacheable.
	Expires string `pulumi:"expires"`
	// The provider-assigned unique ID for this managed resource.
	Id  string `pulumi:"id"`
	Key string `pulumi:"key"`
	// Last modified date of the object in RFC1123 format (e.g., `Mon, 02 Jan 2006 15:04:05 MST`)
	LastModified string `pulumi:"lastModified"`
	// Map of metadata stored with the object in S3. Keys are always returned in lowercase.
	Metadata map[string]string `pulumi:"metadata"`
	// Indicates whether this object has an active [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds). This field is only returned if you have permission to view an object's legal hold status.
	ObjectLockLegalHoldStatus string `pulumi:"objectLockLegalHoldStatus"`
	// Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) currently in place for this object.
	ObjectLockMode string `pulumi:"objectLockMode"`
	// The date and time when this object's object lock will expire.
	ObjectLockRetainUntilDate string  `pulumi:"objectLockRetainUntilDate"`
	Range                     *string `pulumi:"range"`
	// If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.
	ServerSideEncryption string `pulumi:"serverSideEncryption"`
	// If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.
	SseKmsKeyId string `pulumi:"sseKmsKeyId"`
	// [Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for `Standard` storage class objects.
	StorageClass string `pulumi:"storageClass"`
	// Map of tags assigned to the object.
	Tags map[string]string `pulumi:"tags"`
	// Latest version ID of the object returned.
	VersionId string `pulumi:"versionId"`
	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.
	WebsiteRedirectLocation string `pulumi:"websiteRedirectLocation"`
}

func GetObjectOutput(ctx *pulumi.Context, args GetObjectOutputArgs, opts ...pulumi.InvokeOption) GetObjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetObjectResult, error) {
			args := v.(GetObjectArgs)
			r, err := GetObject(ctx, &args, opts...)
			var s GetObjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetObjectResultOutput)
}

// A collection of arguments for invoking getObject.
type GetObjectOutputArgs struct {
	// Name of the bucket to read the object from. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// To retrieve the object's checksum, this argument must be `ENABLED`. If you enable `checksumMode` and the object is encrypted with KMS, you must have permission to use the `kms:Decrypt` action. Valid values: `ENABLED`
	ChecksumMode pulumi.StringPtrInput `pulumi:"checksumMode"`
	// Full path to the object inside the bucket
	Key   pulumi.StringInput    `pulumi:"key"`
	Range pulumi.StringPtrInput `pulumi:"range"`
	// Map of tags assigned to the object.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// Specific version ID of the object returned (defaults to latest version)
	VersionId pulumi.StringPtrInput `pulumi:"versionId"`
}

func (GetObjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetObjectArgs)(nil)).Elem()
}

// A collection of values returned by getObject.
type GetObjectResultOutput struct{ *pulumi.OutputState }

func (GetObjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetObjectResult)(nil)).Elem()
}

func (o GetObjectResultOutput) ToGetObjectResultOutput() GetObjectResultOutput {
	return o
}

func (o GetObjectResultOutput) ToGetObjectResultOutputWithContext(ctx context.Context) GetObjectResultOutput {
	return o
}

func (o GetObjectResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetObjectResult] {
	return pulumix.Output[GetObjectResult]{
		OutputState: o.OutputState,
	}
}

// Object data (see **limitations above** to understand cases in which this field is actually available)
func (o GetObjectResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.Body }).(pulumi.StringOutput)
}

func (o GetObjectResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// (Optional) Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
func (o GetObjectResultOutput) BucketKeyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetObjectResult) bool { return v.BucketKeyEnabled }).(pulumi.BoolOutput)
}

// Caching behavior along the request/reply chain.
func (o GetObjectResultOutput) CacheControl() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.CacheControl }).(pulumi.StringOutput)
}

// The base64-encoded, 32-bit CRC32 checksum of the object.
func (o GetObjectResultOutput) ChecksumCrc32() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ChecksumCrc32 }).(pulumi.StringOutput)
}

// The base64-encoded, 32-bit CRC32C checksum of the object.
func (o GetObjectResultOutput) ChecksumCrc32c() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ChecksumCrc32c }).(pulumi.StringOutput)
}

func (o GetObjectResultOutput) ChecksumMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetObjectResult) *string { return v.ChecksumMode }).(pulumi.StringPtrOutput)
}

// The base64-encoded, 160-bit SHA-1 digest of the object.
func (o GetObjectResultOutput) ChecksumSha1() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ChecksumSha1 }).(pulumi.StringOutput)
}

// The base64-encoded, 256-bit SHA-256 digest of the object.
func (o GetObjectResultOutput) ChecksumSha256() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ChecksumSha256 }).(pulumi.StringOutput)
}

// Presentational information for the object.
func (o GetObjectResultOutput) ContentDisposition() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ContentDisposition }).(pulumi.StringOutput)
}

// What content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
func (o GetObjectResultOutput) ContentEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ContentEncoding }).(pulumi.StringOutput)
}

// Language the content is in.
func (o GetObjectResultOutput) ContentLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ContentLanguage }).(pulumi.StringOutput)
}

// Size of the body in bytes.
func (o GetObjectResultOutput) ContentLength() pulumi.IntOutput {
	return o.ApplyT(func(v GetObjectResult) int { return v.ContentLength }).(pulumi.IntOutput)
}

// Standard MIME type describing the format of the object data.
func (o GetObjectResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ContentType }).(pulumi.StringOutput)
}

// [ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)
func (o GetObjectResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.Etag }).(pulumi.StringOutput)
}

// If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.
func (o GetObjectResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.Expiration }).(pulumi.StringOutput)
}

// Date and time at which the object is no longer cacheable.
func (o GetObjectResultOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.Expires }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetObjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetObjectResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.Key }).(pulumi.StringOutput)
}

// Last modified date of the object in RFC1123 format (e.g., `Mon, 02 Jan 2006 15:04:05 MST`)
func (o GetObjectResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// Map of metadata stored with the object in S3. Keys are always returned in lowercase.
func (o GetObjectResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetObjectResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// Indicates whether this object has an active [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds). This field is only returned if you have permission to view an object's legal hold status.
func (o GetObjectResultOutput) ObjectLockLegalHoldStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ObjectLockLegalHoldStatus }).(pulumi.StringOutput)
}

// Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) currently in place for this object.
func (o GetObjectResultOutput) ObjectLockMode() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ObjectLockMode }).(pulumi.StringOutput)
}

// The date and time when this object's object lock will expire.
func (o GetObjectResultOutput) ObjectLockRetainUntilDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ObjectLockRetainUntilDate }).(pulumi.StringOutput)
}

func (o GetObjectResultOutput) Range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetObjectResult) *string { return v.Range }).(pulumi.StringPtrOutput)
}

// If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.
func (o GetObjectResultOutput) ServerSideEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.ServerSideEncryption }).(pulumi.StringOutput)
}

// If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.
func (o GetObjectResultOutput) SseKmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.SseKmsKeyId }).(pulumi.StringOutput)
}

// [Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for `Standard` storage class objects.
func (o GetObjectResultOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.StorageClass }).(pulumi.StringOutput)
}

// Map of tags assigned to the object.
func (o GetObjectResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetObjectResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Latest version ID of the object returned.
func (o GetObjectResultOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.VersionId }).(pulumi.StringOutput)
}

// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.
func (o GetObjectResultOutput) WebsiteRedirectLocation() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectResult) string { return v.WebsiteRedirectLocation }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetObjectResultOutput{})
}
