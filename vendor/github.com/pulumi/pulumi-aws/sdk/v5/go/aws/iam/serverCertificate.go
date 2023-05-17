// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM Server Certificate resource to upload Server Certificates.
// Certs uploaded to IAM can easily work with other AWS services such as:
//
// - AWS Elastic Beanstalk
// - Elastic Load Balancing
// - CloudFront
// - AWS OpsWorks
//
// For information about server certificates in IAM, see [Managing Server
// Certificates][2] in AWS Documentation.
//
// ## Example Usage
//
// **Using certs on file:**
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewServerCertificate(ctx, "testCert", &iam.ServerCertificateArgs{
//				CertificateBody: readFileOrPanic("self-ca-cert.pem"),
//				PrivateKey:      readFileOrPanic("test-key.pem"),
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
// **Example with cert in-line:**
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewServerCertificate(ctx, "testCertAlt", &iam.ServerCertificateArgs{
//				CertificateBody: pulumi.String("-----BEGIN CERTIFICATE-----\n[......] # cert contents\n-----END CERTIFICATE-----\n\n"),
//				PrivateKey:      pulumi.String("-----BEGIN RSA PRIVATE KEY-----\n[......] # cert contents\n-----END RSA PRIVATE KEY-----\n\n"),
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
// **Use in combination with an AWS ELB resource:**
//
// Some properties of an IAM Server Certificates cannot be updated while they are
// in use. In order for the provider to effectively manage a Certificate in this situation, it is
// recommended you utilize the `namePrefix` attribute and enable the
// `createBeforeDestroy`. This will allow this provider
// to create a new, updated `iam.ServerCertificate` resource and replace it in
// dependant resources before attempting to destroy the old version.
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elb"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testCert, err := iam.NewServerCertificate(ctx, "testCert", &iam.ServerCertificateArgs{
//				NamePrefix:      pulumi.String("example-cert"),
//				CertificateBody: readFileOrPanic("self-ca-cert.pem"),
//				PrivateKey:      readFileOrPanic("test-key.pem"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elb.NewLoadBalancer(ctx, "ourapp", &elb.LoadBalancerArgs{
//				AvailabilityZones: pulumi.StringArray{
//					pulumi.String("us-west-2a"),
//				},
//				CrossZoneLoadBalancing: pulumi.Bool(true),
//				Listeners: elb.LoadBalancerListenerArray{
//					&elb.LoadBalancerListenerArgs{
//						InstancePort:     pulumi.Int(8000),
//						InstanceProtocol: pulumi.String("http"),
//						LbPort:           pulumi.Int(443),
//						LbProtocol:       pulumi.String("https"),
//						SslCertificateId: testCert.Arn,
//					},
//				},
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
// ## Import
//
// IAM Server Certificates can be imported using the `name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:iam/serverCertificate:ServerCertificate certificate example.com-certificate-until-2018
//
// ```
//
//	[1]https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html [2]https://docs.aws.amazon.com/IAM/latest/UserGuide/ManagingServerCerts.html
type ServerCertificate struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) specifying the server certificate.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The contents of the public key certificate in
	// PEM-encoded format.
	CertificateBody pulumi.StringOutput `pulumi:"certificateBody"`
	// The contents of the certificate chain.
	// This is typically a concatenation of the PEM-encoded public key certificates
	// of the chain.
	CertificateChain pulumi.StringPtrOutput `pulumi:"certificateChain"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
	Expiration pulumi.StringOutput `pulumi:"expiration"`
	// The name of the Server Certificate. Do not include the
	// path in this value. If omitted, the provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The contents of the private key in PEM-encoded format.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// Map of resource tags for the server certificate. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
	UploadDate pulumi.StringOutput `pulumi:"uploadDate"`
}

// NewServerCertificate registers a new resource with the given unique name, arguments, and options.
func NewServerCertificate(ctx *pulumi.Context,
	name string, args *ServerCertificateArgs, opts ...pulumi.ResourceOption) (*ServerCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateBody == nil {
		return nil, errors.New("invalid value for required argument 'CertificateBody'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	var resource ServerCertificate
	err := ctx.RegisterResource("aws:iam/serverCertificate:ServerCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerCertificate gets an existing ServerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerCertificateState, opts ...pulumi.ResourceOption) (*ServerCertificate, error) {
	var resource ServerCertificate
	err := ctx.ReadResource("aws:iam/serverCertificate:ServerCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerCertificate resources.
type serverCertificateState struct {
	// The Amazon Resource Name (ARN) specifying the server certificate.
	Arn *string `pulumi:"arn"`
	// The contents of the public key certificate in
	// PEM-encoded format.
	CertificateBody *string `pulumi:"certificateBody"`
	// The contents of the certificate chain.
	// This is typically a concatenation of the PEM-encoded public key certificates
	// of the chain.
	CertificateChain *string `pulumi:"certificateChain"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
	Expiration *string `pulumi:"expiration"`
	// The name of the Server Certificate. Do not include the
	// path in this value. If omitted, the provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
	Path *string `pulumi:"path"`
	// The contents of the private key in PEM-encoded format.
	PrivateKey *string `pulumi:"privateKey"`
	// Map of resource tags for the server certificate. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
	UploadDate *string `pulumi:"uploadDate"`
}

type ServerCertificateState struct {
	// The Amazon Resource Name (ARN) specifying the server certificate.
	Arn pulumi.StringPtrInput
	// The contents of the public key certificate in
	// PEM-encoded format.
	CertificateBody pulumi.StringPtrInput
	// The contents of the certificate chain.
	// This is typically a concatenation of the PEM-encoded public key certificates
	// of the chain.
	CertificateChain pulumi.StringPtrInput
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
	Expiration pulumi.StringPtrInput
	// The name of the Server Certificate. Do not include the
	// path in this value. If omitted, the provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
	Path pulumi.StringPtrInput
	// The contents of the private key in PEM-encoded format.
	PrivateKey pulumi.StringPtrInput
	// Map of resource tags for the server certificate. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
	UploadDate pulumi.StringPtrInput
}

func (ServerCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCertificateState)(nil)).Elem()
}

type serverCertificateArgs struct {
	// The contents of the public key certificate in
	// PEM-encoded format.
	CertificateBody string `pulumi:"certificateBody"`
	// The contents of the certificate chain.
	// This is typically a concatenation of the PEM-encoded public key certificates
	// of the chain.
	CertificateChain *string `pulumi:"certificateChain"`
	// The name of the Server Certificate. Do not include the
	// path in this value. If omitted, the provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
	Path *string `pulumi:"path"`
	// The contents of the private key in PEM-encoded format.
	PrivateKey string `pulumi:"privateKey"`
	// Map of resource tags for the server certificate. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServerCertificate resource.
type ServerCertificateArgs struct {
	// The contents of the public key certificate in
	// PEM-encoded format.
	CertificateBody pulumi.StringInput
	// The contents of the certificate chain.
	// This is typically a concatenation of the PEM-encoded public key certificates
	// of the chain.
	CertificateChain pulumi.StringPtrInput
	// The name of the Server Certificate. Do not include the
	// path in this value. If omitted, the provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
	Path pulumi.StringPtrInput
	// The contents of the private key in PEM-encoded format.
	PrivateKey pulumi.StringInput
	// Map of resource tags for the server certificate. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ServerCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCertificateArgs)(nil)).Elem()
}

type ServerCertificateInput interface {
	pulumi.Input

	ToServerCertificateOutput() ServerCertificateOutput
	ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput
}

func (*ServerCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificate)(nil)).Elem()
}

func (i *ServerCertificate) ToServerCertificateOutput() ServerCertificateOutput {
	return i.ToServerCertificateOutputWithContext(context.Background())
}

func (i *ServerCertificate) ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateOutput)
}

// ServerCertificateArrayInput is an input type that accepts ServerCertificateArray and ServerCertificateArrayOutput values.
// You can construct a concrete instance of `ServerCertificateArrayInput` via:
//
//	ServerCertificateArray{ ServerCertificateArgs{...} }
type ServerCertificateArrayInput interface {
	pulumi.Input

	ToServerCertificateArrayOutput() ServerCertificateArrayOutput
	ToServerCertificateArrayOutputWithContext(context.Context) ServerCertificateArrayOutput
}

type ServerCertificateArray []ServerCertificateInput

func (ServerCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerCertificate)(nil)).Elem()
}

func (i ServerCertificateArray) ToServerCertificateArrayOutput() ServerCertificateArrayOutput {
	return i.ToServerCertificateArrayOutputWithContext(context.Background())
}

func (i ServerCertificateArray) ToServerCertificateArrayOutputWithContext(ctx context.Context) ServerCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateArrayOutput)
}

// ServerCertificateMapInput is an input type that accepts ServerCertificateMap and ServerCertificateMapOutput values.
// You can construct a concrete instance of `ServerCertificateMapInput` via:
//
//	ServerCertificateMap{ "key": ServerCertificateArgs{...} }
type ServerCertificateMapInput interface {
	pulumi.Input

	ToServerCertificateMapOutput() ServerCertificateMapOutput
	ToServerCertificateMapOutputWithContext(context.Context) ServerCertificateMapOutput
}

type ServerCertificateMap map[string]ServerCertificateInput

func (ServerCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerCertificate)(nil)).Elem()
}

func (i ServerCertificateMap) ToServerCertificateMapOutput() ServerCertificateMapOutput {
	return i.ToServerCertificateMapOutputWithContext(context.Background())
}

func (i ServerCertificateMap) ToServerCertificateMapOutputWithContext(ctx context.Context) ServerCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateMapOutput)
}

type ServerCertificateOutput struct{ *pulumi.OutputState }

func (ServerCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificate)(nil)).Elem()
}

func (o ServerCertificateOutput) ToServerCertificateOutput() ServerCertificateOutput {
	return o
}

func (o ServerCertificateOutput) ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput {
	return o
}

// The Amazon Resource Name (ARN) specifying the server certificate.
func (o ServerCertificateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The contents of the public key certificate in
// PEM-encoded format.
func (o ServerCertificateOutput) CertificateBody() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.CertificateBody }).(pulumi.StringOutput)
}

// The contents of the certificate chain.
// This is typically a concatenation of the PEM-encoded public key certificates
// of the chain.
func (o ServerCertificateOutput) CertificateChain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.CertificateChain }).(pulumi.StringPtrOutput)
}

// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
func (o ServerCertificateOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.Expiration }).(pulumi.StringOutput)
}

// The name of the Server Certificate. Do not include the
// path in this value. If omitted, the provider will assign a random, unique name.
func (o ServerCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified
// prefix. Conflicts with `name`.
func (o ServerCertificateOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// The IAM path for the server certificate.  If it is not
// included, it defaults to a slash (/). If this certificate is for use with
// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
func (o ServerCertificateOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// The contents of the private key in PEM-encoded format.
func (o ServerCertificateOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// Map of resource tags for the server certificate. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ServerCertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ServerCertificateOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
func (o ServerCertificateOutput) UploadDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.UploadDate }).(pulumi.StringOutput)
}

type ServerCertificateArrayOutput struct{ *pulumi.OutputState }

func (ServerCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerCertificate)(nil)).Elem()
}

func (o ServerCertificateArrayOutput) ToServerCertificateArrayOutput() ServerCertificateArrayOutput {
	return o
}

func (o ServerCertificateArrayOutput) ToServerCertificateArrayOutputWithContext(ctx context.Context) ServerCertificateArrayOutput {
	return o
}

func (o ServerCertificateArrayOutput) Index(i pulumi.IntInput) ServerCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerCertificate {
		return vs[0].([]*ServerCertificate)[vs[1].(int)]
	}).(ServerCertificateOutput)
}

type ServerCertificateMapOutput struct{ *pulumi.OutputState }

func (ServerCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerCertificate)(nil)).Elem()
}

func (o ServerCertificateMapOutput) ToServerCertificateMapOutput() ServerCertificateMapOutput {
	return o
}

func (o ServerCertificateMapOutput) ToServerCertificateMapOutputWithContext(ctx context.Context) ServerCertificateMapOutput {
	return o
}

func (o ServerCertificateMapOutput) MapIndex(k pulumi.StringInput) ServerCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerCertificate {
		return vs[0].(map[string]*ServerCertificate)[vs[1].(string)]
	}).(ServerCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerCertificateInput)(nil)).Elem(), &ServerCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerCertificateArrayInput)(nil)).Elem(), ServerCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerCertificateMapInput)(nil)).Elem(), ServerCertificateMap{})
	pulumi.RegisterOutputType(ServerCertificateOutput{})
	pulumi.RegisterOutputType(ServerCertificateArrayOutput{})
	pulumi.RegisterOutputType(ServerCertificateMapOutput{})
}
