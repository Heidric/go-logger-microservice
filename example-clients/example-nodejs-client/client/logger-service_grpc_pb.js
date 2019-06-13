// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var logger$service_pb = require('./logger-service_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_pb_src_EntriesRequest(arg) {
  if (!(arg instanceof logger$service_pb.EntriesRequest)) {
    throw new Error('Expected argument of type pb_src.EntriesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_src_EntriesRequest(buffer_arg) {
  return logger$service_pb.EntriesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_src_EntriesResponse(arg) {
  if (!(arg instanceof logger$service_pb.EntriesResponse)) {
    throw new Error('Expected argument of type pb_src.EntriesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_src_EntriesResponse(buffer_arg) {
  return logger$service_pb.EntriesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_src_LogCreationResponse(arg) {
  if (!(arg instanceof logger$service_pb.LogCreationResponse)) {
    throw new Error('Expected argument of type pb_src.LogCreationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_src_LogCreationResponse(buffer_arg) {
  return logger$service_pb.LogCreationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_src_LogEntry(arg) {
  if (!(arg instanceof logger$service_pb.LogEntry)) {
    throw new Error('Expected argument of type pb_src.LogEntry');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_src_LogEntry(buffer_arg) {
  return logger$service_pb.LogEntry.deserializeBinary(new Uint8Array(buffer_arg));
}


var LoggerServiceService = exports.LoggerServiceService = {
  createEntry: {
    path: '/pb_src.LoggerService/CreateEntry',
    requestStream: false,
    responseStream: false,
    requestType: logger$service_pb.LogEntry,
    responseType: logger$service_pb.LogCreationResponse,
    requestSerialize: serialize_pb_src_LogEntry,
    requestDeserialize: deserialize_pb_src_LogEntry,
    responseSerialize: serialize_pb_src_LogCreationResponse,
    responseDeserialize: deserialize_pb_src_LogCreationResponse,
  },
  getEntries: {
    path: '/pb_src.LoggerService/GetEntries',
    requestStream: false,
    responseStream: false,
    requestType: logger$service_pb.EntriesRequest,
    responseType: logger$service_pb.EntriesResponse,
    requestSerialize: serialize_pb_src_EntriesRequest,
    requestDeserialize: deserialize_pb_src_EntriesRequest,
    responseSerialize: serialize_pb_src_EntriesResponse,
    responseDeserialize: deserialize_pb_src_EntriesResponse,
  },
};

exports.LoggerServiceClient = grpc.makeGenericClientConstructor(LoggerServiceService);
