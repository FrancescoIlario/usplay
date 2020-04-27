/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.activitycomm.ListActivitiesReply');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.activitycomm.Activity');


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
proto.activitycomm.ListActivitiesReply = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.activitycomm.ListActivitiesReply.repeatedFields_, null);
};
goog.inherits(proto.activitycomm.ListActivitiesReply, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.activitycomm.ListActivitiesReply.displayName = 'proto.activitycomm.ListActivitiesReply';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.activitycomm.ListActivitiesReply.repeatedFields_ = [1];



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
proto.activitycomm.ListActivitiesReply.prototype.toObject = function(opt_includeInstance) {
  return proto.activitycomm.ListActivitiesReply.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.activitycomm.ListActivitiesReply} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.activitycomm.ListActivitiesReply.toObject = function(includeInstance, msg) {
  var f, obj = {
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
 * @return {!proto.activitycomm.ListActivitiesReply}
 */
proto.activitycomm.ListActivitiesReply.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.activitycomm.ListActivitiesReply;
  return proto.activitycomm.ListActivitiesReply.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.activitycomm.ListActivitiesReply} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.activitycomm.ListActivitiesReply}
 */
proto.activitycomm.ListActivitiesReply.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
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
proto.activitycomm.ListActivitiesReply.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.activitycomm.ListActivitiesReply.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.activitycomm.ListActivitiesReply} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.activitycomm.ListActivitiesReply.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getActivitiesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.activitycomm.Activity.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Activity activities = 1;
 * @return {!Array<!proto.activitycomm.Activity>}
 */
proto.activitycomm.ListActivitiesReply.prototype.getActivitiesList = function() {
  return /** @type{!Array<!proto.activitycomm.Activity>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.activitycomm.Activity, 1));
};


/** @param {!Array<!proto.activitycomm.Activity>} value */
proto.activitycomm.ListActivitiesReply.prototype.setActivitiesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.activitycomm.Activity=} opt_value
 * @param {number=} opt_index
 * @return {!proto.activitycomm.Activity}
 */
proto.activitycomm.ListActivitiesReply.prototype.addActivities = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.activitycomm.Activity, opt_index);
};


proto.activitycomm.ListActivitiesReply.prototype.clearActivitiesList = function() {
  this.setActivitiesList([]);
};


