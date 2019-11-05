///
//  Generated code. Do not modify.
//  source: proto/types/error.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/any.pb.dart' as $0;

class Error extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Error', package: const $pb.PackageName('types'), createEmptyInstance: create)
    ..a<$core.int>(1, 'code', $pb.PbFieldType.O3)
    ..aOS(2, 'message')
    ..aOM<$0.Any>(4, 'details', subBuilder: $0.Any.create)
    ..m<$core.String, ValidationError>(5, 'validation', entryClassName: 'Error.ValidationEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OM, valueCreator: ValidationError.create, packageName: const $pb.PackageName('types'))
    ..hasRequiredFields = false
  ;

  Error._() : super();
  factory Error() => create();
  factory Error.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Error.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Error clone() => Error()..mergeFromMessage(this);
  Error copyWith(void Function(Error) updates) => super.copyWith((message) => updates(message as Error));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Error create() => Error._();
  Error createEmptyInstance() => create();
  static $pb.PbList<Error> createRepeated() => $pb.PbList<Error>();
  @$core.pragma('dart2js:noInline')
  static Error getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Error>(create);
  static Error _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get code => $_getIZ(0);
  @$pb.TagNumber(1)
  set code($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get message => $_getSZ(1);
  @$pb.TagNumber(2)
  set message($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearMessage() => clearField(2);

  @$pb.TagNumber(4)
  $0.Any get details => $_getN(2);
  @$pb.TagNumber(4)
  set details($0.Any v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasDetails() => $_has(2);
  @$pb.TagNumber(4)
  void clearDetails() => clearField(4);
  @$pb.TagNumber(4)
  $0.Any ensureDetails() => $_ensure(2);

  @$pb.TagNumber(5)
  $core.Map<$core.String, ValidationError> get validation => $_getMap(3);
}

class ValidationError extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ValidationError', package: const $pb.PackageName('types'), createEmptyInstance: create)
    ..pPS(1, 'errors')
    ..hasRequiredFields = false
  ;

  ValidationError._() : super();
  factory ValidationError() => create();
  factory ValidationError.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ValidationError.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ValidationError clone() => ValidationError()..mergeFromMessage(this);
  ValidationError copyWith(void Function(ValidationError) updates) => super.copyWith((message) => updates(message as ValidationError));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ValidationError create() => ValidationError._();
  ValidationError createEmptyInstance() => create();
  static $pb.PbList<ValidationError> createRepeated() => $pb.PbList<ValidationError>();
  @$core.pragma('dart2js:noInline')
  static ValidationError getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ValidationError>(create);
  static ValidationError _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get errors => $_getList(0);
}

