/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.reportcomm.Report');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.activitycomm.Activity');
goog.require('proto.reportcomm.Interval');


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
proto.reportcomm.Report = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.reportcomm.Report.repeatedFields_, null);
};
goog.inherits(proto.reportcomm.Report, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.reportcomm.Report.displayName = 'proto.reportcomm.Report';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.reportcomm.Report.repeatedFields_ = [5];



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
proto.reportcomm.Report.prototype.toObject = function(opt_includeInstance) {
  return proto.reportcomm.Report.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.reportcomm.Report} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.reportcomm.Report.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    orderid: jspb.Message.getFieldWithDefault(msg, 3, ""),
    period: (f = msg.getPeriod()) && proto.reportcomm.Interval.toObject(includeInstance, f),
    activitiesList: jspb.Message.toObjectList(msg.getActivitiesList(),
    proto.activitycomm.Activity.toObject, includeInstance)
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
 * @return {!proto.reportcomm.Report}
 */
proto.reportcomm.Report.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.reportcomm.Report;
  return proto.reportcomm.Report.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.reportcomm.Report} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.reportcomm.Report}
 */
proto.reportcomm.Report.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setOrderid(value);
      break;
    case 4:
      var value = new proto.reportcomm.Interval;
      reader.readMessage(value,proto.reportcomm.Interval.deserializeBinaryFromReader);
      msg.setPeriod(value);
      break;
    case 5:
      var value = new proto.activitycomm.Activity;
      reader.readMessage(value,proto.activitycomm.Activity.deserializeBinaryFromReader);
      msg.addActivities(value);
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
proto.reportcomm.Report.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.reportcomm.Report.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.reportcomm.Report} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.reportcomm.Report.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getName();
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
      proto.reportcomm.Interval.serializeBinaryToWriter
    );
  }
  f = message.getActivitiesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      5,
      f,
      proto.activitycomm.Activity.serializeBinaryToWriter
    );
  }
};


/**
 * optional string Id = 1;
 * @return {string}
 */
proto.reportcomm.Report.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.reportcomm.Report.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string Name = 2;
 * @return {string}
 */
proto.reportcomm.Report.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.reportcomm.Report.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string OrderId = 3;
 * @return {string}
 */
proto.reportcomm.Report.prototype.getOrderid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.reportcomm.Report.prototype.setOrderid = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional Interval Period = 4;
 * @return {?proto.reportcomm.Interval}
 */
proto.reportcomm.Report.prototype.getPeriod = function() {
  return /** @type{?proto.reportcomm.Interval} */ (
    jspb.Message.getWrapperField(this, proto.reportcomm.Interval, 4));
};


/** @param {?proto.reportcomm.Interval|undefined} value */
proto.reportcomm.Report.prototype.setPeriod = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.reportcomm.Report.prototype.clearPeriod = function() {
  this.setPeriod(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.reportcomm.Report.prototype.hasPeriod = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * repeated activitycomm.Activity Activities = 5;
 * @return {!Array<!proto.activitycomm.Activity>}
 */
proto.reportcomm.Report.prototype.getActivitiesList = function() {
  return /** @type{!Array<!proto.activitycomm.Activity>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.activitycomm.Activity, 5));
};


/** @param {!Array<!proto.activitycomm.Activity>} value */
proto.reportcomm.Report.prototype.setActivitiesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 5, value);
};


/**
 * @param {!proto.activitycomm.Activity=} opt_value
 * @param {number=} opt_index
 * @return {!proto.activitycomm.Activity}
 */
proto.reportcomm.Report.prototype.addActivities = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 5, opt_value, proto.activitycomm.Activity, opt_index);
};


proto.reportcomm.Report.prototype.clearActivitiesList = function() {
  this.setActivitiesList([]);
};


