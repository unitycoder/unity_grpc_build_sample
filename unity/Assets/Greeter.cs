﻿using System;
using System.Collections;
using System.Collections.Generic;
using System.Threading;
using UnityEngine;
using grpc = global::Grpc.Core;

/// <summary>
/// The greeting service definition.
/// </summary>
public static partial class Greeter
{
    static readonly string __ServiceName = "helloworld.Greeter";

    static readonly grpc::Marshaller<global::Helloworld.HelloRequest> __Marshaller_HelloRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Helloworld.HelloRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Helloworld.HelloReply> __Marshaller_HelloReply = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Helloworld.HelloReply.Parser.ParseFrom);

    static readonly grpc::Method<global::Helloworld.HelloRequest, global::Helloworld.HelloReply> __Method_SayHello = new grpc::Method<global::Helloworld.HelloRequest, global::Helloworld.HelloReply>(
        grpc::MethodType.Unary,
        __ServiceName,
        "SayHello",
        __Marshaller_HelloRequest,
        __Marshaller_HelloReply);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
        get { return global::Helloworld.HelloReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of Greeter</summary>
    public abstract partial class GreeterBase
    {
        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request received from the client.</param>
        /// <param name="context">The context of the server-side call handler being invoked.</param>
        /// <returns>The response to send back to the client (wrapped by a task).</returns>
        public virtual global::System.Threading.Tasks.Task<global::Helloworld.HelloReply> SayHello(global::Helloworld.HelloRequest request, grpc::ServerCallContext context)
        {
            throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
        }

    }

    /// <summary>Client for Greeter</summary>
    public partial class GreeterClient : grpc::ClientBase<GreeterClient>
    {
        /// <summary>Creates a new client for Greeter</summary>
        /// <param name="channel">The channel to use to make remote calls.</param>
        public GreeterClient(grpc::Channel channel) : base(channel)
        {
        }
        /// <summary>Creates a new client for Greeter that uses a custom <c>CallInvoker</c>.</summary>
        /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
        public GreeterClient(grpc::CallInvoker callInvoker) : base(callInvoker)
        {
        }
        /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
        protected GreeterClient() : base()
        {
        }
        /// <summary>Protected constructor to allow creation of configured clients.</summary>
        /// <param name="configuration">The client configuration.</param>
        protected GreeterClient(ClientBaseConfiguration configuration) : base(configuration)
        {
        }

        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request to send to the server.</param>
        /// <param name="headers">The initial metadata to send with the call. This parameter is optional.</param>
        /// <param name="deadline">An optional deadline for the call. The call will be cancelled if deadline is hit.</param>
        /// <param name="cancellationToken">An optional token for canceling the call.</param>
        /// <returns>The response received from the server.</returns>
        public virtual global::Helloworld.HelloReply SayHello(global::Helloworld.HelloRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
        {
            return SayHello(request, new grpc::CallOptions(headers, deadline, cancellationToken));
        }
        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request to send to the server.</param>
        /// <param name="options">The options for the call.</param>
        /// <returns>The response received from the server.</returns>
        public virtual global::Helloworld.HelloReply SayHello(global::Helloworld.HelloRequest request, grpc::CallOptions options)
        {
            return CallInvoker.BlockingUnaryCall(__Method_SayHello, null, options, request);
        }
        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request to send to the server.</param>
        /// <param name="headers">The initial metadata to send with the call. This parameter is optional.</param>
        /// <param name="deadline">An optional deadline for the call. The call will be cancelled if deadline is hit.</param>
        /// <param name="cancellationToken">An optional token for canceling the call.</param>
        /// <returns>The call object.</returns>
        public virtual grpc::AsyncUnaryCall<global::Helloworld.HelloReply> SayHelloAsync(global::Helloworld.HelloRequest request, grpc::Metadata headers = null, DateTime? deadline = null, CancellationToken cancellationToken = default(CancellationToken))
        {
            return SayHelloAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
        }
        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request to send to the server.</param>
        /// <param name="options">The options for the call.</param>
        /// <returns>The call object.</returns>
        public virtual grpc::AsyncUnaryCall<global::Helloworld.HelloReply> SayHelloAsync(global::Helloworld.HelloRequest request, grpc::CallOptions options)
        {
            return CallInvoker.AsyncUnaryCall(__Method_SayHello, null, options, request);
        }
        /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
        protected override GreeterClient NewInstance(ClientBaseConfiguration configuration)
        {
            return new GreeterClient(configuration);
        }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static grpc::ServerServiceDefinition BindService(GreeterBase serviceImpl)
    {
        return grpc::ServerServiceDefinition.CreateBuilder()
            .AddMethod(__Method_SayHello, serviceImpl.SayHello).Build();
    }

}