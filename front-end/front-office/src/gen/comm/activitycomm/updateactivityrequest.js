/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.activitycomm.UpdateActivityRequest');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.activitycomm.Interval');


/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.activitycomm.UpdateActivityRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.activitycomm.UpdateActivityRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.activitycomm.UpdateActivityRequest.displayName = 'proto.activitycomm.UpdateActivityRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.activitycomm.UpdateActivityRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.activitycomm.UpdateActivityRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.activitycomm.UpdateActivityRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.activitycomm.UpdateActivityRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    acttypeid: jspb.Message.getFieldWithDefault(msg, 2, ""),
    orderid: jspb.Message.getFieldWithDefault(msg, 3, ""),
    period: (f = msg.getPeriod()) && proto.activitycomm.Interval.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.activitycomm.UpdateActivityRequest}
 */
proto.activitycomm.UpdateActivityRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.activitycomm.UpdateActivityRequest;
  return proto.activitycomm.UpdateActivityRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.activitycomm.UpdateActivityRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.activitycomm.UpdateActivityRequest}
 */
proto.activitycomm.UpdateActivityRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setActtypeid(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setOrderid(value);
      break;
    case 4:
      var value = new proto.activitycomm.Interval;
      reader.readMessage(value,proto.activitycomm.Interval.deserializeBinaryFromReader);
      msg.setPeriod(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.activitycomm.UpdateActivityRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.activitycomm.UpdateActivityRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.activitycomm.UpdateActivityRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.activitycomm.UpdateActivityRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getActtypeid();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getOrderid();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getPeriod();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.activitycomm.Interval.serializeBinaryToWriter
    );
  }
};


/**
 * optional string Id = 1;
 * @return {string}
 */
proto.activitycomm.UpdateActivityRequest.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.activitycomm.UpdateActivityRequest.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string ActTypeID = 2;
 * @return {string}
 */
proto.activitycomm.UpdateActivityRequest.prototype.getActtypeid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.activitycomm.UpdateActivityRequest.prototype.setActtypeid = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string OrderID = 3;
 * @return {string}
 */
proto.activitycomm.UpdateActivityRequest.prototype.getOrderid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.activitycomm.UpdateActivityRequest.prototype.setOrderid = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional Interval Period = 4;
 * @return {?proto.activitycomm.Interval}
 */
proto.activitycomm.UpdateActivityRequest.prototype.getPeriod = function() {
  return /** @type{?proto.activitycomm.Interval} */ (
    jspb.Message.getWrapperField(this, proto.activitycomm.Interval, 4));
};


/** @param {?proto.activitycomm.Interval|undefined} value */
proto.activitycomm.UpdateActivityRequest.prototype.setPeriod = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.activitycomm.UpdateActivityRequest.prototype.clearPeriod = function() {
  this.setPeriod(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.activitycomm.UpdateActivityRequest.prototype.hasPeriod = function() {
  return jspb.Message.getField(this, 4) != null;
};

