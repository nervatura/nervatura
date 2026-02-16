//
//  Generated code. Do not modify.
//  source: store.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'store.pbenum.dart';

export 'store.pbenum.dart';

/// No request parameters
class RequestEmpty extends $pb.GeneratedMessage {
  factory RequestEmpty() => create();
  RequestEmpty._() : super();
  factory RequestEmpty.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestEmpty.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestEmpty', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestEmpty clone() => RequestEmpty()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestEmpty copyWith(void Function(RequestEmpty) updates) => super.copyWith((message) => updates(message as RequestEmpty)) as RequestEmpty;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestEmpty create() => RequestEmpty._();
  RequestEmpty createEmptyInstance() => create();
  static $pb.PbList<RequestEmpty> createRepeated() => $pb.PbList<RequestEmpty>();
  @$core.pragma('dart2js:noInline')
  static RequestEmpty getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestEmpty>(create);
  static RequestEmpty? _defaultInstance;
}

/// Does not return content.
class ResponseStatus extends $pb.GeneratedMessage {
  factory ResponseStatus({
    $core.bool? success,
    $core.String? message,
  }) {
    final $result = create();
    if (success != null) {
      $result.success = success;
    }
    if (message != null) {
      $result.message = message;
    }
    return $result;
  }
  ResponseStatus._() : super();
  factory ResponseStatus.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ResponseStatus.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ResponseStatus', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'success')
    ..aOS(2, _omitFieldNames ? '' : 'message')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ResponseStatus clone() => ResponseStatus()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ResponseStatus copyWith(void Function(ResponseStatus) updates) => super.copyWith((message) => updates(message as ResponseStatus)) as ResponseStatus;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ResponseStatus create() => ResponseStatus._();
  ResponseStatus createEmptyInstance() => create();
  static $pb.PbList<ResponseStatus> createRepeated() => $pb.PbList<ResponseStatus>();
  @$core.pragma('dart2js:noInline')
  static ResponseStatus getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ResponseStatus>(create);
  static ResponseStatus? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get success => $_getBF(0);
  @$pb.TagNumber(1)
  set success($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSuccess() => $_has(0);
  @$pb.TagNumber(1)
  void clearSuccess() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get message => $_getSZ(1);
  @$pb.TagNumber(2)
  set message($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearMessage() => clearField(2);
}

class RequestAuthLogin extends $pb.GeneratedMessage {
  factory RequestAuthLogin({
    $core.String? userName,
    $core.String? password,
    $core.String? database,
  }) {
    final $result = create();
    if (userName != null) {
      $result.userName = userName;
    }
    if (password != null) {
      $result.password = password;
    }
    if (database != null) {
      $result.database = database;
    }
    return $result;
  }
  RequestAuthLogin._() : super();
  factory RequestAuthLogin.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestAuthLogin.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestAuthLogin', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'userName')
    ..aOS(2, _omitFieldNames ? '' : 'password')
    ..aOS(3, _omitFieldNames ? '' : 'database')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestAuthLogin clone() => RequestAuthLogin()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestAuthLogin copyWith(void Function(RequestAuthLogin) updates) => super.copyWith((message) => updates(message as RequestAuthLogin)) as RequestAuthLogin;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestAuthLogin create() => RequestAuthLogin._();
  RequestAuthLogin createEmptyInstance() => create();
  static $pb.PbList<RequestAuthLogin> createRepeated() => $pb.PbList<RequestAuthLogin>();
  @$core.pragma('dart2js:noInline')
  static RequestAuthLogin getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestAuthLogin>(create);
  static RequestAuthLogin? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userName => $_getSZ(0);
  @$pb.TagNumber(1)
  set userName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserName() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get password => $_getSZ(1);
  @$pb.TagNumber(2)
  set password($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPassword() => $_has(1);
  @$pb.TagNumber(2)
  void clearPassword() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get database => $_getSZ(2);
  @$pb.TagNumber(3)
  set database($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDatabase() => $_has(2);
  @$pb.TagNumber(3)
  void clearDatabase() => clearField(3);
}

class ResponseAuthLogin extends $pb.GeneratedMessage {
  factory ResponseAuthLogin({
    $core.String? token,
    $core.String? version,
  }) {
    final $result = create();
    if (token != null) {
      $result.token = token;
    }
    if (version != null) {
      $result.version = version;
    }
    return $result;
  }
  ResponseAuthLogin._() : super();
  factory ResponseAuthLogin.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ResponseAuthLogin.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ResponseAuthLogin', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'token')
    ..aOS(2, _omitFieldNames ? '' : 'version')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ResponseAuthLogin clone() => ResponseAuthLogin()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ResponseAuthLogin copyWith(void Function(ResponseAuthLogin) updates) => super.copyWith((message) => updates(message as ResponseAuthLogin)) as ResponseAuthLogin;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ResponseAuthLogin create() => ResponseAuthLogin._();
  ResponseAuthLogin createEmptyInstance() => create();
  static $pb.PbList<ResponseAuthLogin> createRepeated() => $pb.PbList<ResponseAuthLogin>();
  @$core.pragma('dart2js:noInline')
  static ResponseAuthLogin getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ResponseAuthLogin>(create);
  static ResponseAuthLogin? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get version => $_getSZ(1);
  @$pb.TagNumber(2)
  set version($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasVersion() => $_has(1);
  @$pb.TagNumber(2)
  void clearVersion() => clearField(2);
}

class RequestPasswordChange extends $pb.GeneratedMessage {
  factory RequestPasswordChange({
    $core.String? password,
    $core.String? confirm,
  }) {
    final $result = create();
    if (password != null) {
      $result.password = password;
    }
    if (confirm != null) {
      $result.confirm = confirm;
    }
    return $result;
  }
  RequestPasswordChange._() : super();
  factory RequestPasswordChange.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestPasswordChange.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestPasswordChange', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'password')
    ..aOS(2, _omitFieldNames ? '' : 'confirm')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestPasswordChange clone() => RequestPasswordChange()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestPasswordChange copyWith(void Function(RequestPasswordChange) updates) => super.copyWith((message) => updates(message as RequestPasswordChange)) as RequestPasswordChange;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestPasswordChange create() => RequestPasswordChange._();
  RequestPasswordChange createEmptyInstance() => create();
  static $pb.PbList<RequestPasswordChange> createRepeated() => $pb.PbList<RequestPasswordChange>();
  @$core.pragma('dart2js:noInline')
  static RequestPasswordChange getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestPasswordChange>(create);
  static RequestPasswordChange? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get password => $_getSZ(0);
  @$pb.TagNumber(1)
  set password($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasPassword() => $_has(0);
  @$pb.TagNumber(1)
  void clearPassword() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get confirm => $_getSZ(1);
  @$pb.TagNumber(2)
  set confirm($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasConfirm() => $_has(1);
  @$pb.TagNumber(2)
  void clearConfirm() => clearField(2);
}

class RequestTokenDecode extends $pb.GeneratedMessage {
  factory RequestTokenDecode({
    $core.String? token,
  }) {
    final $result = create();
    if (token != null) {
      $result.token = token;
    }
    return $result;
  }
  RequestTokenDecode._() : super();
  factory RequestTokenDecode.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestTokenDecode.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestTokenDecode', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'token')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestTokenDecode clone() => RequestTokenDecode()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestTokenDecode copyWith(void Function(RequestTokenDecode) updates) => super.copyWith((message) => updates(message as RequestTokenDecode)) as RequestTokenDecode;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestTokenDecode create() => RequestTokenDecode._();
  RequestTokenDecode createEmptyInstance() => create();
  static $pb.PbList<RequestTokenDecode> createRepeated() => $pb.PbList<RequestTokenDecode>();
  @$core.pragma('dart2js:noInline')
  static RequestTokenDecode getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestTokenDecode>(create);
  static RequestTokenDecode? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);
}

class ResponseTokenDecode extends $pb.GeneratedMessage {
  factory ResponseTokenDecode({
    $core.String? code,
    $core.String? userName,
    $core.String? database,
    $core.String? exp,
    $core.String? iss,
  }) {
    final $result = create();
    if (code != null) {
      $result.code = code;
    }
    if (userName != null) {
      $result.userName = userName;
    }
    if (database != null) {
      $result.database = database;
    }
    if (exp != null) {
      $result.exp = exp;
    }
    if (iss != null) {
      $result.iss = iss;
    }
    return $result;
  }
  ResponseTokenDecode._() : super();
  factory ResponseTokenDecode.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ResponseTokenDecode.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ResponseTokenDecode', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'code')
    ..aOS(2, _omitFieldNames ? '' : 'userName')
    ..aOS(3, _omitFieldNames ? '' : 'database')
    ..aOS(4, _omitFieldNames ? '' : 'exp')
    ..aOS(5, _omitFieldNames ? '' : 'iss')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ResponseTokenDecode clone() => ResponseTokenDecode()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ResponseTokenDecode copyWith(void Function(ResponseTokenDecode) updates) => super.copyWith((message) => updates(message as ResponseTokenDecode)) as ResponseTokenDecode;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ResponseTokenDecode create() => ResponseTokenDecode._();
  ResponseTokenDecode createEmptyInstance() => create();
  static $pb.PbList<ResponseTokenDecode> createRepeated() => $pb.PbList<ResponseTokenDecode>();
  @$core.pragma('dart2js:noInline')
  static ResponseTokenDecode getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ResponseTokenDecode>(create);
  static ResponseTokenDecode? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get code => $_getSZ(0);
  @$pb.TagNumber(1)
  set code($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get userName => $_getSZ(1);
  @$pb.TagNumber(2)
  set userName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUserName() => $_has(1);
  @$pb.TagNumber(2)
  void clearUserName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get database => $_getSZ(2);
  @$pb.TagNumber(3)
  set database($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDatabase() => $_has(2);
  @$pb.TagNumber(3)
  void clearDatabase() => clearField(3);

  /// JWT expiration time
  @$pb.TagNumber(4)
  $core.String get exp => $_getSZ(3);
  @$pb.TagNumber(4)
  set exp($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasExp() => $_has(3);
  @$pb.TagNumber(4)
  void clearExp() => clearField(4);

  /// JWT issuer
  @$pb.TagNumber(5)
  $core.String get iss => $_getSZ(4);
  @$pb.TagNumber(5)
  set iss($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasIss() => $_has(4);
  @$pb.TagNumber(5)
  void clearIss() => clearField(5);
}

class RequestGet extends $pb.GeneratedMessage {
  factory RequestGet({
    $fixnum.Int64? id,
    $core.String? code,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    return $result;
  }
  RequestGet._() : super();
  factory RequestGet.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestGet.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestGet', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestGet clone() => RequestGet()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestGet copyWith(void Function(RequestGet) updates) => super.copyWith((message) => updates(message as RequestGet)) as RequestGet;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestGet create() => RequestGet._();
  RequestGet createEmptyInstance() => create();
  static $pb.PbList<RequestGet> createRepeated() => $pb.PbList<RequestGet>();
  @$core.pragma('dart2js:noInline')
  static RequestGet getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestGet>(create);
  static RequestGet? _defaultInstance;

  /// Database primary key. If specified, the code does not need to be set.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If specified, the id does not need to be set.
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);
}

class RequestDelete extends $pb.GeneratedMessage {
  factory RequestDelete({
    $fixnum.Int64? id,
    $core.String? code,
    Model? model,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (model != null) {
      $result.model = model;
    }
    return $result;
  }
  RequestDelete._() : super();
  factory RequestDelete.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestDelete.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestDelete', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<Model>(3, _omitFieldNames ? '' : 'model', $pb.PbFieldType.OE, defaultOrMaker: Model.AUTH, valueOf: Model.valueOf, enumValues: Model.values)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestDelete clone() => RequestDelete()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestDelete copyWith(void Function(RequestDelete) updates) => super.copyWith((message) => updates(message as RequestDelete)) as RequestDelete;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestDelete create() => RequestDelete._();
  RequestDelete createEmptyInstance() => create();
  static $pb.PbList<RequestDelete> createRepeated() => $pb.PbList<RequestDelete>();
  @$core.pragma('dart2js:noInline')
  static RequestDelete getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestDelete>(create);
  static RequestDelete? _defaultInstance;

  /// Database primary key. If specified, the code does not need to be set.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If specified, the id does not need to be set.
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// Model name
  @$pb.TagNumber(3)
  Model get model => $_getN(2);
  @$pb.TagNumber(3)
  set model(Model v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasModel() => $_has(2);
  @$pb.TagNumber(3)
  void clearModel() => clearField(3);
}

class RequestQueryFilter extends $pb.GeneratedMessage {
  factory RequestQueryFilter({
    $core.String? field_1,
    $core.String? value,
  }) {
    final $result = create();
    if (field_1 != null) {
      $result.field_1 = field_1;
    }
    if (value != null) {
      $result.value = value;
    }
    return $result;
  }
  RequestQueryFilter._() : super();
  factory RequestQueryFilter.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestQueryFilter.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestQueryFilter', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'field')
    ..aOS(2, _omitFieldNames ? '' : 'value')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestQueryFilter clone() => RequestQueryFilter()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestQueryFilter copyWith(void Function(RequestQueryFilter) updates) => super.copyWith((message) => updates(message as RequestQueryFilter)) as RequestQueryFilter;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestQueryFilter create() => RequestQueryFilter._();
  RequestQueryFilter createEmptyInstance() => create();
  static $pb.PbList<RequestQueryFilter> createRepeated() => $pb.PbList<RequestQueryFilter>();
  @$core.pragma('dart2js:noInline')
  static RequestQueryFilter getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestQueryFilter>(create);
  static RequestQueryFilter? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get field_1 => $_getSZ(0);
  @$pb.TagNumber(1)
  set field_1($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasField_1() => $_has(0);
  @$pb.TagNumber(1)
  void clearField_1() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get value => $_getSZ(1);
  @$pb.TagNumber(2)
  set value($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => clearField(2);
}

class RequestQuery extends $pb.GeneratedMessage {
  factory RequestQuery({
    $core.Iterable<RequestQueryFilter>? filters,
    $fixnum.Int64? limit,
    $fixnum.Int64? offset,
  }) {
    final $result = create();
    if (filters != null) {
      $result.filters.addAll(filters);
    }
    if (limit != null) {
      $result.limit = limit;
    }
    if (offset != null) {
      $result.offset = offset;
    }
    return $result;
  }
  RequestQuery._() : super();
  factory RequestQuery.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestQuery.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestQuery', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<RequestQueryFilter>(1, _omitFieldNames ? '' : 'filters', $pb.PbFieldType.PM, subBuilder: RequestQueryFilter.create)
    ..aInt64(2, _omitFieldNames ? '' : 'limit')
    ..aInt64(3, _omitFieldNames ? '' : 'offset')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestQuery clone() => RequestQuery()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestQuery copyWith(void Function(RequestQuery) updates) => super.copyWith((message) => updates(message as RequestQuery)) as RequestQuery;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestQuery create() => RequestQuery._();
  RequestQuery createEmptyInstance() => create();
  static $pb.PbList<RequestQuery> createRepeated() => $pb.PbList<RequestQuery>();
  @$core.pragma('dart2js:noInline')
  static RequestQuery getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestQuery>(create);
  static RequestQuery? _defaultInstance;

  /// Fieldname and filter value
  @$pb.TagNumber(1)
  $core.List<RequestQueryFilter> get filters => $_getList(0);

  /// Limit the number of records
  @$pb.TagNumber(2)
  $fixnum.Int64 get limit => $_getI64(1);
  @$pb.TagNumber(2)
  set limit($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLimit() => $_has(1);
  @$pb.TagNumber(2)
  void clearLimit() => clearField(2);

  /// Skip the first N records
  @$pb.TagNumber(3)
  $fixnum.Int64 get offset => $_getI64(2);
  @$pb.TagNumber(3)
  set offset($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasOffset() => $_has(2);
  @$pb.TagNumber(3)
  void clearOffset() => clearField(3);
}

class RequestFunction extends $pb.GeneratedMessage {
  factory RequestFunction({
    $core.String? function,
    $core.Map<$core.String, $core.String>? args,
  }) {
    final $result = create();
    if (function != null) {
      $result.function = function;
    }
    if (args != null) {
      $result.args.addAll(args);
    }
    return $result;
  }
  RequestFunction._() : super();
  factory RequestFunction.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestFunction.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestFunction', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'function')
    ..m<$core.String, $core.String>(2, _omitFieldNames ? '' : 'args', entryClassName: 'RequestFunction.ArgsEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('nervatura'))
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestFunction clone() => RequestFunction()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestFunction copyWith(void Function(RequestFunction) updates) => super.copyWith((message) => updates(message as RequestFunction)) as RequestFunction;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestFunction create() => RequestFunction._();
  RequestFunction createEmptyInstance() => create();
  static $pb.PbList<RequestFunction> createRepeated() => $pb.PbList<RequestFunction>();
  @$core.pragma('dart2js:noInline')
  static RequestFunction getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestFunction>(create);
  static RequestFunction? _defaultInstance;

  /// Server function name.
  /// Valid function names: report_install, report_list, report_get, product_price, email_send
  @$pb.TagNumber(1)
  $core.String get function => $_getSZ(0);
  @$pb.TagNumber(1)
  set function($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFunction() => $_has(0);
  @$pb.TagNumber(1)
  void clearFunction() => clearField(1);

  /// Function arguments
  @$pb.TagNumber(2)
  $core.Map<$core.String, $core.String> get args => $_getMap(1);
}

class RequestView extends $pb.GeneratedMessage {
  factory RequestView({
    ViewName? name,
    $core.Iterable<RequestQueryFilter>? filters,
    $core.Iterable<$core.String>? orderBy,
    $fixnum.Int64? limit,
    $fixnum.Int64? offset,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    if (filters != null) {
      $result.filters.addAll(filters);
    }
    if (orderBy != null) {
      $result.orderBy.addAll(orderBy);
    }
    if (limit != null) {
      $result.limit = limit;
    }
    if (offset != null) {
      $result.offset = offset;
    }
    return $result;
  }
  RequestView._() : super();
  factory RequestView.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestView.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestView', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..e<ViewName>(1, _omitFieldNames ? '' : 'name', $pb.PbFieldType.OE, defaultOrMaker: ViewName.VIEW_CONFIG_DATA, valueOf: ViewName.valueOf, enumValues: ViewName.values)
    ..pc<RequestQueryFilter>(2, _omitFieldNames ? '' : 'filters', $pb.PbFieldType.PM, subBuilder: RequestQueryFilter.create)
    ..pPS(3, _omitFieldNames ? '' : 'orderBy')
    ..aInt64(4, _omitFieldNames ? '' : 'limit')
    ..aInt64(5, _omitFieldNames ? '' : 'offset')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestView clone() => RequestView()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestView copyWith(void Function(RequestView) updates) => super.copyWith((message) => updates(message as RequestView)) as RequestView;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestView create() => RequestView._();
  RequestView createEmptyInstance() => create();
  static $pb.PbList<RequestView> createRepeated() => $pb.PbList<RequestView>();
  @$core.pragma('dart2js:noInline')
  static RequestView getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestView>(create);
  static RequestView? _defaultInstance;

  /// Predefined view name
  @$pb.TagNumber(1)
  ViewName get name => $_getN(0);
  @$pb.TagNumber(1)
  set name(ViewName v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  /// Fieldname and filter value
  @$pb.TagNumber(2)
  $core.List<RequestQueryFilter> get filters => $_getList(1);

  /// Valid SQL ORDER BY clause
  @$pb.TagNumber(3)
  $core.List<$core.String> get orderBy => $_getList(2);

  /// Limit the number of records
  @$pb.TagNumber(4)
  $fixnum.Int64 get limit => $_getI64(3);
  @$pb.TagNumber(4)
  set limit($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasLimit() => $_has(3);
  @$pb.TagNumber(4)
  void clearLimit() => clearField(4);

  /// Skip the first N records
  @$pb.TagNumber(5)
  $fixnum.Int64 get offset => $_getI64(4);
  @$pb.TagNumber(5)
  set offset($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasOffset() => $_has(4);
  @$pb.TagNumber(5)
  void clearOffset() => clearField(5);
}

class RequestDatabase extends $pb.GeneratedMessage {
  factory RequestDatabase({
    $core.String? alias,
    $core.bool? demo,
  }) {
    final $result = create();
    if (alias != null) {
      $result.alias = alias;
    }
    if (demo != null) {
      $result.demo = demo;
    }
    return $result;
  }
  RequestDatabase._() : super();
  factory RequestDatabase.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RequestDatabase.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RequestDatabase', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'alias')
    ..aOB(2, _omitFieldNames ? '' : 'demo')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RequestDatabase clone() => RequestDatabase()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RequestDatabase copyWith(void Function(RequestDatabase) updates) => super.copyWith((message) => updates(message as RequestDatabase)) as RequestDatabase;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RequestDatabase create() => RequestDatabase._();
  RequestDatabase createEmptyInstance() => create();
  static $pb.PbList<RequestDatabase> createRepeated() => $pb.PbList<RequestDatabase>();
  @$core.pragma('dart2js:noInline')
  static RequestDatabase getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RequestDatabase>(create);
  static RequestDatabase? _defaultInstance;

  /// Database alias name
  @$pb.TagNumber(1)
  $core.String get alias => $_getSZ(0);
  @$pb.TagNumber(1)
  set alias($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAlias() => $_has(0);
  @$pb.TagNumber(1)
  void clearAlias() => clearField(1);

  /// Create demo data
  @$pb.TagNumber(2)
  $core.bool get demo => $_getBF(1);
  @$pb.TagNumber(2)
  set demo($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDemo() => $_has(1);
  @$pb.TagNumber(2)
  void clearDemo() => clearField(2);
}

class JsonString extends $pb.GeneratedMessage {
  factory JsonString({
    $core.Map<$core.String, $core.String>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  JsonString._() : super();
  factory JsonString.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory JsonString.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'JsonString', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..m<$core.String, $core.String>(1, _omitFieldNames ? '' : 'data', entryClassName: 'JsonString.DataEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('nervatura'))
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  JsonString clone() => JsonString()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  JsonString copyWith(void Function(JsonString) updates) => super.copyWith((message) => updates(message as JsonString)) as JsonString;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static JsonString create() => JsonString._();
  JsonString createEmptyInstance() => create();
  static $pb.PbList<JsonString> createRepeated() => $pb.PbList<JsonString>();
  @$core.pragma('dart2js:noInline')
  static JsonString getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<JsonString>(create);
  static JsonString? _defaultInstance;

  /// Build in convertion of number, boolean, array and object types to string
  @$pb.TagNumber(1)
  $core.Map<$core.String, $core.String> get data => $_getMap(0);
}

class JsonBytes extends $pb.GeneratedMessage {
  factory JsonBytes({
    $core.List<$core.int>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data = data;
    }
    return $result;
  }
  JsonBytes._() : super();
  factory JsonBytes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory JsonBytes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'JsonBytes', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  JsonBytes clone() => JsonBytes()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  JsonBytes copyWith(void Function(JsonBytes) updates) => super.copyWith((message) => updates(message as JsonBytes)) as JsonBytes;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static JsonBytes create() => JsonBytes._();
  JsonBytes createEmptyInstance() => create();
  static $pb.PbList<JsonBytes> createRepeated() => $pb.PbList<JsonBytes>();
  @$core.pragma('dart2js:noInline')
  static JsonBytes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<JsonBytes>(create);
  static JsonBytes? _defaultInstance;

  /// JSON data as bytes
  @$pb.TagNumber(1)
  $core.List<$core.int> get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
}

/// Address message
class Address extends $pb.GeneratedMessage {
  factory Address({
    $core.String? country,
    $core.String? state,
    $core.String? zipCode,
    $core.String? city,
    $core.String? street,
    $core.String? notes,
    $core.Iterable<$core.String>? tags,
    JsonString? addressMap,
  }) {
    final $result = create();
    if (country != null) {
      $result.country = country;
    }
    if (state != null) {
      $result.state = state;
    }
    if (zipCode != null) {
      $result.zipCode = zipCode;
    }
    if (city != null) {
      $result.city = city;
    }
    if (street != null) {
      $result.street = street;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    if (addressMap != null) {
      $result.addressMap = addressMap;
    }
    return $result;
  }
  Address._() : super();
  factory Address.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Address.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Address', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'country')
    ..aOS(2, _omitFieldNames ? '' : 'state')
    ..aOS(3, _omitFieldNames ? '' : 'zipCode')
    ..aOS(4, _omitFieldNames ? '' : 'city')
    ..aOS(5, _omitFieldNames ? '' : 'street')
    ..aOS(6, _omitFieldNames ? '' : 'notes')
    ..pPS(7, _omitFieldNames ? '' : 'tags')
    ..aOM<JsonString>(8, _omitFieldNames ? '' : 'addressMap', subBuilder: JsonString.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Address clone() => Address()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Address copyWith(void Function(Address) updates) => super.copyWith((message) => updates(message as Address)) as Address;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Address create() => Address._();
  Address createEmptyInstance() => create();
  static $pb.PbList<Address> createRepeated() => $pb.PbList<Address>();
  @$core.pragma('dart2js:noInline')
  static Address getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Address>(create);
  static Address? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get country => $_getSZ(0);
  @$pb.TagNumber(1)
  set country($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCountry() => $_has(0);
  @$pb.TagNumber(1)
  void clearCountry() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get state => $_getSZ(1);
  @$pb.TagNumber(2)
  set state($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasState() => $_has(1);
  @$pb.TagNumber(2)
  void clearState() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get zipCode => $_getSZ(2);
  @$pb.TagNumber(3)
  set zipCode($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasZipCode() => $_has(2);
  @$pb.TagNumber(3)
  void clearZipCode() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get city => $_getSZ(3);
  @$pb.TagNumber(4)
  set city($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCity() => $_has(3);
  @$pb.TagNumber(4)
  void clearCity() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get street => $_getSZ(4);
  @$pb.TagNumber(5)
  set street($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasStreet() => $_has(4);
  @$pb.TagNumber(5)
  void clearStreet() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get notes => $_getSZ(5);
  @$pb.TagNumber(6)
  set notes($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasNotes() => $_has(5);
  @$pb.TagNumber(6)
  void clearNotes() => clearField(6);

  /// Additional tags for the address
  @$pb.TagNumber(7)
  $core.List<$core.String> get tags => $_getList(6);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(8)
  JsonString get addressMap => $_getN(7);
  @$pb.TagNumber(8)
  set addressMap(JsonString v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasAddressMap() => $_has(7);
  @$pb.TagNumber(8)
  void clearAddressMap() => clearField(8);
  @$pb.TagNumber(8)
  JsonString ensureAddressMap() => $_ensure(7);
}

/// Contact message
class Contact extends $pb.GeneratedMessage {
  factory Contact({
    $core.String? firstName,
    $core.String? surname,
    $core.String? status,
    $core.String? phone,
    $core.String? mobile,
    $core.String? email,
    $core.String? notes,
    $core.Iterable<$core.String>? tags,
    JsonString? contactMap,
  }) {
    final $result = create();
    if (firstName != null) {
      $result.firstName = firstName;
    }
    if (surname != null) {
      $result.surname = surname;
    }
    if (status != null) {
      $result.status = status;
    }
    if (phone != null) {
      $result.phone = phone;
    }
    if (mobile != null) {
      $result.mobile = mobile;
    }
    if (email != null) {
      $result.email = email;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    if (contactMap != null) {
      $result.contactMap = contactMap;
    }
    return $result;
  }
  Contact._() : super();
  factory Contact.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Contact.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Contact', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'firstName')
    ..aOS(2, _omitFieldNames ? '' : 'surname')
    ..aOS(3, _omitFieldNames ? '' : 'status')
    ..aOS(4, _omitFieldNames ? '' : 'phone')
    ..aOS(5, _omitFieldNames ? '' : 'mobile')
    ..aOS(6, _omitFieldNames ? '' : 'email')
    ..aOS(7, _omitFieldNames ? '' : 'notes')
    ..pPS(8, _omitFieldNames ? '' : 'tags')
    ..aOM<JsonString>(9, _omitFieldNames ? '' : 'contactMap', subBuilder: JsonString.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Contact clone() => Contact()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Contact copyWith(void Function(Contact) updates) => super.copyWith((message) => updates(message as Contact)) as Contact;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Contact create() => Contact._();
  Contact createEmptyInstance() => create();
  static $pb.PbList<Contact> createRepeated() => $pb.PbList<Contact>();
  @$core.pragma('dart2js:noInline')
  static Contact getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Contact>(create);
  static Contact? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get firstName => $_getSZ(0);
  @$pb.TagNumber(1)
  set firstName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFirstName() => $_has(0);
  @$pb.TagNumber(1)
  void clearFirstName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get surname => $_getSZ(1);
  @$pb.TagNumber(2)
  set surname($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSurname() => $_has(1);
  @$pb.TagNumber(2)
  void clearSurname() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get status => $_getSZ(2);
  @$pb.TagNumber(3)
  set status($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasStatus() => $_has(2);
  @$pb.TagNumber(3)
  void clearStatus() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get phone => $_getSZ(3);
  @$pb.TagNumber(4)
  set phone($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPhone() => $_has(3);
  @$pb.TagNumber(4)
  void clearPhone() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get mobile => $_getSZ(4);
  @$pb.TagNumber(5)
  set mobile($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasMobile() => $_has(4);
  @$pb.TagNumber(5)
  void clearMobile() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get email => $_getSZ(5);
  @$pb.TagNumber(6)
  set email($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasEmail() => $_has(5);
  @$pb.TagNumber(6)
  void clearEmail() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get notes => $_getSZ(6);
  @$pb.TagNumber(7)
  set notes($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasNotes() => $_has(6);
  @$pb.TagNumber(7)
  void clearNotes() => clearField(7);

  /// Additional tags for the contact
  @$pb.TagNumber(8)
  $core.List<$core.String> get tags => $_getList(7);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(9)
  JsonString get contactMap => $_getN(8);
  @$pb.TagNumber(9)
  set contactMap(JsonString v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasContactMap() => $_has(8);
  @$pb.TagNumber(9)
  void clearContactMap() => clearField(9);
  @$pb.TagNumber(9)
  JsonString ensureContactMap() => $_ensure(8);
}

/// Event message
class Event extends $pb.GeneratedMessage {
  factory Event({
    $core.String? uid,
    $core.String? subject,
    $core.String? startTime,
    $core.String? endTime,
    $core.String? place,
    $core.String? description,
    $core.Iterable<$core.String>? tags,
    JsonString? eventMap,
  }) {
    final $result = create();
    if (uid != null) {
      $result.uid = uid;
    }
    if (subject != null) {
      $result.subject = subject;
    }
    if (startTime != null) {
      $result.startTime = startTime;
    }
    if (endTime != null) {
      $result.endTime = endTime;
    }
    if (place != null) {
      $result.place = place;
    }
    if (description != null) {
      $result.description = description;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    if (eventMap != null) {
      $result.eventMap = eventMap;
    }
    return $result;
  }
  Event._() : super();
  factory Event.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Event.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Event', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'uid')
    ..aOS(2, _omitFieldNames ? '' : 'subject')
    ..aOS(3, _omitFieldNames ? '' : 'startTime')
    ..aOS(4, _omitFieldNames ? '' : 'endTime')
    ..aOS(5, _omitFieldNames ? '' : 'place')
    ..aOS(6, _omitFieldNames ? '' : 'description')
    ..pPS(7, _omitFieldNames ? '' : 'tags')
    ..aOM<JsonString>(8, _omitFieldNames ? '' : 'eventMap', subBuilder: JsonString.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Event clone() => Event()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Event copyWith(void Function(Event) updates) => super.copyWith((message) => updates(message as Event)) as Event;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Event create() => Event._();
  Event createEmptyInstance() => create();
  static $pb.PbList<Event> createRepeated() => $pb.PbList<Event>();
  @$core.pragma('dart2js:noInline')
  static Event getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Event>(create);
  static Event? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get uid => $_getSZ(0);
  @$pb.TagNumber(1)
  set uid($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUid() => $_has(0);
  @$pb.TagNumber(1)
  void clearUid() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get subject => $_getSZ(1);
  @$pb.TagNumber(2)
  set subject($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSubject() => $_has(1);
  @$pb.TagNumber(2)
  void clearSubject() => clearField(2);

  /// ISO 8601 date-time
  @$pb.TagNumber(3)
  $core.String get startTime => $_getSZ(2);
  @$pb.TagNumber(3)
  set startTime($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasStartTime() => $_has(2);
  @$pb.TagNumber(3)
  void clearStartTime() => clearField(3);

  /// ISO 8601 date-time
  @$pb.TagNumber(4)
  $core.String get endTime => $_getSZ(3);
  @$pb.TagNumber(4)
  set endTime($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasEndTime() => $_has(3);
  @$pb.TagNumber(4)
  void clearEndTime() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get place => $_getSZ(4);
  @$pb.TagNumber(5)
  set place($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasPlace() => $_has(4);
  @$pb.TagNumber(5)
  void clearPlace() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get description => $_getSZ(5);
  @$pb.TagNumber(6)
  set description($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasDescription() => $_has(5);
  @$pb.TagNumber(6)
  void clearDescription() => clearField(6);

  /// Additional tags for the event
  @$pb.TagNumber(7)
  $core.List<$core.String> get tags => $_getList(6);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(8)
  JsonString get eventMap => $_getN(7);
  @$pb.TagNumber(8)
  set eventMap(JsonString v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasEventMap() => $_has(7);
  @$pb.TagNumber(8)
  void clearEventMap() => clearField(8);
  @$pb.TagNumber(8)
  JsonString ensureEventMap() => $_ensure(7);
}

/// AuthMeta Bookmark message
class Bookmark extends $pb.GeneratedMessage {
  factory Bookmark({
    BookmarkType? bookmarkType,
    $core.String? label,
    $core.String? key,
    $core.String? code,
    JsonString? filters,
    $core.Map<$core.String, $core.bool>? columns,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (bookmarkType != null) {
      $result.bookmarkType = bookmarkType;
    }
    if (label != null) {
      $result.label = label;
    }
    if (key != null) {
      $result.key = key;
    }
    if (code != null) {
      $result.code = code;
    }
    if (filters != null) {
      $result.filters = filters;
    }
    if (columns != null) {
      $result.columns.addAll(columns);
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Bookmark._() : super();
  factory Bookmark.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Bookmark.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Bookmark', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..e<BookmarkType>(1, _omitFieldNames ? '' : 'bookmarkType', $pb.PbFieldType.OE, defaultOrMaker: BookmarkType.BOOKMARK_BROWSER, valueOf: BookmarkType.valueOf, enumValues: BookmarkType.values)
    ..aOS(2, _omitFieldNames ? '' : 'label')
    ..aOS(3, _omitFieldNames ? '' : 'key')
    ..aOS(4, _omitFieldNames ? '' : 'code')
    ..aOM<JsonString>(5, _omitFieldNames ? '' : 'filters', subBuilder: JsonString.create)
    ..m<$core.String, $core.bool>(6, _omitFieldNames ? '' : 'columns', entryClassName: 'Bookmark.ColumnsEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OB, packageName: const $pb.PackageName('nervatura'))
    ..aOS(7, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Bookmark clone() => Bookmark()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Bookmark copyWith(void Function(Bookmark) updates) => super.copyWith((message) => updates(message as Bookmark)) as Bookmark;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Bookmark create() => Bookmark._();
  Bookmark createEmptyInstance() => create();
  static $pb.PbList<Bookmark> createRepeated() => $pb.PbList<Bookmark>();
  @$core.pragma('dart2js:noInline')
  static Bookmark getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Bookmark>(create);
  static Bookmark? _defaultInstance;

  /// ENUM field. Valid values: BROWSER, EDITOR
  @$pb.TagNumber(1)
  BookmarkType get bookmarkType => $_getN(0);
  @$pb.TagNumber(1)
  set bookmarkType(BookmarkType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasBookmarkType() => $_has(0);
  @$pb.TagNumber(1)
  void clearBookmarkType() => clearField(1);

  /// Label of the bookmark
  @$pb.TagNumber(2)
  $core.String get label => $_getSZ(1);
  @$pb.TagNumber(2)
  set label($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLabel() => $_has(1);
  @$pb.TagNumber(2)
  void clearLabel() => clearField(2);

  /// Editor model or browser view name
  @$pb.TagNumber(3)
  $core.String get key => $_getSZ(2);
  @$pb.TagNumber(3)
  set key($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasKey() => $_has(2);
  @$pb.TagNumber(3)
  void clearKey() => clearField(3);

  /// Model code
  @$pb.TagNumber(4)
  $core.String get code => $_getSZ(3);
  @$pb.TagNumber(4)
  set code($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCode() => $_has(3);
  @$pb.TagNumber(4)
  void clearCode() => clearField(4);

  /// Browser filters
  @$pb.TagNumber(5)
  JsonString get filters => $_getN(4);
  @$pb.TagNumber(5)
  set filters(JsonString v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasFilters() => $_has(4);
  @$pb.TagNumber(5)
  void clearFilters() => clearField(5);
  @$pb.TagNumber(5)
  JsonString ensureFilters() => $_ensure(4);

  /// Browser visible columns
  @$pb.TagNumber(6)
  $core.Map<$core.String, $core.bool> get columns => $_getMap(5);

  /// Timestamp of data creation
  @$pb.TagNumber(7)
  $core.String get timeStamp => $_getSZ(6);
  @$pb.TagNumber(7)
  set timeStamp($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasTimeStamp() => $_has(6);
  @$pb.TagNumber(7)
  void clearTimeStamp() => clearField(7);
}

/// AuthMeta message
class AuthMeta extends $pb.GeneratedMessage {
  factory AuthMeta({
    $core.Iterable<$core.String>? tags,
    $core.Iterable<Bookmark>? bookmarks,
  }) {
    final $result = create();
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    if (bookmarks != null) {
      $result.bookmarks.addAll(bookmarks);
    }
    return $result;
  }
  AuthMeta._() : super();
  factory AuthMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AuthMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AuthMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pPS(1, _omitFieldNames ? '' : 'tags')
    ..pc<Bookmark>(2, _omitFieldNames ? '' : 'bookmarks', $pb.PbFieldType.PM, subBuilder: Bookmark.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AuthMeta clone() => AuthMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AuthMeta copyWith(void Function(AuthMeta) updates) => super.copyWith((message) => updates(message as AuthMeta)) as AuthMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AuthMeta create() => AuthMeta._();
  AuthMeta createEmptyInstance() => create();
  static $pb.PbList<AuthMeta> createRepeated() => $pb.PbList<AuthMeta>();
  @$core.pragma('dart2js:noInline')
  static AuthMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AuthMeta>(create);
  static AuthMeta? _defaultInstance;

  /// Additional tags for the user
  @$pb.TagNumber(1)
  $core.List<$core.String> get tags => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<Bookmark> get bookmarks => $_getList(1);
}

/// Auth message
class Auth extends $pb.GeneratedMessage {
  factory Auth({
    $fixnum.Int64? id,
    $core.String? code,
    $core.String? userName,
    UserGroup? userGroup,
    $core.bool? disabled,
    AuthMeta? authMeta,
    JsonString? authMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (userName != null) {
      $result.userName = userName;
    }
    if (userGroup != null) {
      $result.userGroup = userGroup;
    }
    if (disabled != null) {
      $result.disabled = disabled;
    }
    if (authMeta != null) {
      $result.authMeta = authMeta;
    }
    if (authMap != null) {
      $result.authMap = authMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Auth._() : super();
  factory Auth.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Auth.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Auth', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..aOS(3, _omitFieldNames ? '' : 'userName')
    ..e<UserGroup>(4, _omitFieldNames ? '' : 'userGroup', $pb.PbFieldType.OE, defaultOrMaker: UserGroup.GROUP_USER, valueOf: UserGroup.valueOf, enumValues: UserGroup.values)
    ..aOB(5, _omitFieldNames ? '' : 'disabled')
    ..aOM<AuthMeta>(6, _omitFieldNames ? '' : 'authMeta', subBuilder: AuthMeta.create)
    ..aOM<JsonString>(7, _omitFieldNames ? '' : 'authMap', subBuilder: JsonString.create)
    ..aOS(8, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Auth clone() => Auth()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Auth copyWith(void Function(Auth) updates) => super.copyWith((message) => updates(message as Auth)) as Auth;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Auth create() => Auth._();
  Auth createEmptyInstance() => create();
  static $pb.PbList<Auth> createRepeated() => $pb.PbList<Auth>();
  @$core.pragma('dart2js:noInline')
  static Auth getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Auth>(create);
  static Auth? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: USR1731101982N123 ("USR" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Friendly username, nickname or email for login.
  /// It can be changed, but it is a unique identifier at the database level.
  @$pb.TagNumber(3)
  $core.String get userName => $_getSZ(2);
  @$pb.TagNumber(3)
  set userName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasUserName() => $_has(2);
  @$pb.TagNumber(3)
  void clearUserName() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field. Default value: GROUP_USER
  @$pb.TagNumber(4)
  UserGroup get userGroup => $_getN(3);
  @$pb.TagNumber(4)
  set userGroup(UserGroup v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasUserGroup() => $_has(3);
  @$pb.TagNumber(4)
  void clearUserGroup() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get disabled => $_getBF(4);
  @$pb.TagNumber(5)
  set disabled($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasDisabled() => $_has(4);
  @$pb.TagNumber(5)
  void clearDisabled() => clearField(5);

  @$pb.TagNumber(6)
  AuthMeta get authMeta => $_getN(5);
  @$pb.TagNumber(6)
  set authMeta(AuthMeta v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasAuthMeta() => $_has(5);
  @$pb.TagNumber(6)
  void clearAuthMeta() => clearField(6);
  @$pb.TagNumber(6)
  AuthMeta ensureAuthMeta() => $_ensure(5);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(7)
  JsonString get authMap => $_getN(6);
  @$pb.TagNumber(7)
  set authMap(JsonString v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasAuthMap() => $_has(6);
  @$pb.TagNumber(7)
  void clearAuthMap() => clearField(7);
  @$pb.TagNumber(7)
  JsonString ensureAuthMap() => $_ensure(6);

  /// Timestamp of data creation
  @$pb.TagNumber(8)
  $core.String get timeStamp => $_getSZ(7);
  @$pb.TagNumber(8)
  set timeStamp($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasTimeStamp() => $_has(7);
  @$pb.TagNumber(8)
  void clearTimeStamp() => clearField(8);
}

/// CurrencyMeta message
class CurrencyMeta extends $pb.GeneratedMessage {
  factory CurrencyMeta({
    $core.String? description,
    $fixnum.Int64? digit,
    $fixnum.Int64? cashRound,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (description != null) {
      $result.description = description;
    }
    if (digit != null) {
      $result.digit = digit;
    }
    if (cashRound != null) {
      $result.cashRound = cashRound;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  CurrencyMeta._() : super();
  factory CurrencyMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CurrencyMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CurrencyMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'description')
    ..aInt64(2, _omitFieldNames ? '' : 'digit')
    ..aInt64(3, _omitFieldNames ? '' : 'cashRound')
    ..pPS(4, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CurrencyMeta clone() => CurrencyMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CurrencyMeta copyWith(void Function(CurrencyMeta) updates) => super.copyWith((message) => updates(message as CurrencyMeta)) as CurrencyMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CurrencyMeta create() => CurrencyMeta._();
  CurrencyMeta createEmptyInstance() => create();
  static $pb.PbList<CurrencyMeta> createRepeated() => $pb.PbList<CurrencyMeta>();
  @$core.pragma('dart2js:noInline')
  static CurrencyMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CurrencyMeta>(create);
  static CurrencyMeta? _defaultInstance;

  /// The name of the currency.
  @$pb.TagNumber(1)
  $core.String get description => $_getSZ(0);
  @$pb.TagNumber(1)
  set description($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDescription() => $_has(0);
  @$pb.TagNumber(1)
  void clearDescription() => clearField(1);

  /// The number of decimal places used for recording and rounding by the program.
  @$pb.TagNumber(2)
  $fixnum.Int64 get digit => $_getI64(1);
  @$pb.TagNumber(2)
  set digit($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDigit() => $_has(1);
  @$pb.TagNumber(2)
  void clearDigit() => clearField(2);

  /// Rounding value for cash. Could be used in case the smallest banknote in circulation for that certain currency is not 1.
  @$pb.TagNumber(3)
  $fixnum.Int64 get cashRound => $_getI64(2);
  @$pb.TagNumber(3)
  set cashRound($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasCashRound() => $_has(2);
  @$pb.TagNumber(3)
  void clearCashRound() => clearField(3);

  /// Additional tags for the currency
  @$pb.TagNumber(4)
  $core.List<$core.String> get tags => $_getList(3);
}

/// Currency message
class Currency extends $pb.GeneratedMessage {
  factory Currency({
    $fixnum.Int64? id,
    $core.String? code,
    CurrencyMeta? currencyMeta,
    JsonString? currencyMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (currencyMeta != null) {
      $result.currencyMeta = currencyMeta;
    }
    if (currencyMap != null) {
      $result.currencyMap = currencyMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Currency._() : super();
  factory Currency.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Currency.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Currency', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..aOM<CurrencyMeta>(3, _omitFieldNames ? '' : 'currencyMeta', subBuilder: CurrencyMeta.create)
    ..aOM<JsonString>(4, _omitFieldNames ? '' : 'currencyMap', subBuilder: JsonString.create)
    ..aOS(5, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Currency clone() => Currency()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Currency copyWith(void Function(Currency) updates) => super.copyWith((message) => updates(message as Currency)) as Currency;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Currency create() => Currency._();
  Currency createEmptyInstance() => create();
  static $pb.PbList<Currency> createRepeated() => $pb.PbList<Currency>();
  @$core.pragma('dart2js:noInline')
  static Currency getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Currency>(create);
  static Currency? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// The ISO 4217 code of the currency. It cannot be modified after creation.
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  @$pb.TagNumber(3)
  CurrencyMeta get currencyMeta => $_getN(2);
  @$pb.TagNumber(3)
  set currencyMeta(CurrencyMeta v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasCurrencyMeta() => $_has(2);
  @$pb.TagNumber(3)
  void clearCurrencyMeta() => clearField(3);
  @$pb.TagNumber(3)
  CurrencyMeta ensureCurrencyMeta() => $_ensure(2);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(4)
  JsonString get currencyMap => $_getN(3);
  @$pb.TagNumber(4)
  set currencyMap(JsonString v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasCurrencyMap() => $_has(3);
  @$pb.TagNumber(4)
  void clearCurrencyMap() => clearField(4);
  @$pb.TagNumber(4)
  JsonString ensureCurrencyMap() => $_ensure(3);

  /// Timestamp of data creation
  @$pb.TagNumber(5)
  $core.String get timeStamp => $_getSZ(4);
  @$pb.TagNumber(5)
  set timeStamp($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTimeStamp() => $_has(4);
  @$pb.TagNumber(5)
  void clearTimeStamp() => clearField(5);
}

class Currencies extends $pb.GeneratedMessage {
  factory Currencies({
    $core.Iterable<Currency>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Currencies._() : super();
  factory Currencies.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Currencies.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Currencies', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Currency>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Currency.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Currencies clone() => Currencies()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Currencies copyWith(void Function(Currencies) updates) => super.copyWith((message) => updates(message as Currencies)) as Currencies;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Currencies create() => Currencies._();
  Currencies createEmptyInstance() => create();
  static $pb.PbList<Currencies> createRepeated() => $pb.PbList<Currencies>();
  @$core.pragma('dart2js:noInline')
  static Currencies getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Currencies>(create);
  static Currencies? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Currency> get data => $_getList(0);
}

/// CustomerMeta message
class CustomerMeta extends $pb.GeneratedMessage {
  factory CustomerMeta({
    $core.String? taxNumber,
    $core.String? account,
    $core.bool? taxFree,
    $fixnum.Int64? terms,
    $core.double? creditLimit,
    $core.double? discount,
    $core.String? notes,
    $core.bool? inactive,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (taxNumber != null) {
      $result.taxNumber = taxNumber;
    }
    if (account != null) {
      $result.account = account;
    }
    if (taxFree != null) {
      $result.taxFree = taxFree;
    }
    if (terms != null) {
      $result.terms = terms;
    }
    if (creditLimit != null) {
      $result.creditLimit = creditLimit;
    }
    if (discount != null) {
      $result.discount = discount;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (inactive != null) {
      $result.inactive = inactive;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  CustomerMeta._() : super();
  factory CustomerMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CustomerMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CustomerMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'taxNumber')
    ..aOS(2, _omitFieldNames ? '' : 'account')
    ..aOB(3, _omitFieldNames ? '' : 'taxFree')
    ..aInt64(4, _omitFieldNames ? '' : 'terms')
    ..a<$core.double>(5, _omitFieldNames ? '' : 'creditLimit', $pb.PbFieldType.OD)
    ..a<$core.double>(6, _omitFieldNames ? '' : 'discount', $pb.PbFieldType.OD)
    ..aOS(7, _omitFieldNames ? '' : 'notes')
    ..aOB(8, _omitFieldNames ? '' : 'inactive')
    ..pPS(9, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CustomerMeta clone() => CustomerMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CustomerMeta copyWith(void Function(CustomerMeta) updates) => super.copyWith((message) => updates(message as CustomerMeta)) as CustomerMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CustomerMeta create() => CustomerMeta._();
  CustomerMeta createEmptyInstance() => create();
  static $pb.PbList<CustomerMeta> createRepeated() => $pb.PbList<CustomerMeta>();
  @$core.pragma('dart2js:noInline')
  static CustomerMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CustomerMeta>(create);
  static CustomerMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get taxNumber => $_getSZ(0);
  @$pb.TagNumber(1)
  set taxNumber($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTaxNumber() => $_has(0);
  @$pb.TagNumber(1)
  void clearTaxNumber() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get account => $_getSZ(1);
  @$pb.TagNumber(2)
  set account($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasAccount() => $_has(1);
  @$pb.TagNumber(2)
  void clearAccount() => clearField(2);

  /// Tax-free
  @$pb.TagNumber(3)
  $core.bool get taxFree => $_getBF(2);
  @$pb.TagNumber(3)
  set taxFree($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTaxFree() => $_has(2);
  @$pb.TagNumber(3)
  void clearTaxFree() => clearField(3);

  /// Payment per.
  @$pb.TagNumber(4)
  $fixnum.Int64 get terms => $_getI64(3);
  @$pb.TagNumber(4)
  set terms($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTerms() => $_has(3);
  @$pb.TagNumber(4)
  void clearTerms() => clearField(4);

  /// Customer's credit limit. Data is used by financial reports.
  @$pb.TagNumber(5)
  $core.double get creditLimit => $_getN(4);
  @$pb.TagNumber(5)
  set creditLimit($core.double v) { $_setDouble(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasCreditLimit() => $_has(4);
  @$pb.TagNumber(5)
  void clearCreditLimit() => clearField(5);

  /// If new product line is added (offer, order, invoice etc.) all products will receive the discount percentage specified in this field. If the product has a separate customer price, the value specified here will not be considered by the program.
  @$pb.TagNumber(6)
  $core.double get discount => $_getN(5);
  @$pb.TagNumber(6)
  set discount($core.double v) { $_setDouble(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasDiscount() => $_has(5);
  @$pb.TagNumber(6)
  void clearDiscount() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get notes => $_getSZ(6);
  @$pb.TagNumber(7)
  set notes($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasNotes() => $_has(6);
  @$pb.TagNumber(7)
  void clearNotes() => clearField(7);

  @$pb.TagNumber(8)
  $core.bool get inactive => $_getBF(7);
  @$pb.TagNumber(8)
  set inactive($core.bool v) { $_setBool(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasInactive() => $_has(7);
  @$pb.TagNumber(8)
  void clearInactive() => clearField(8);

  /// Additional tags for the customer
  @$pb.TagNumber(9)
  $core.List<$core.String> get tags => $_getList(8);
}

/// Customer message
class Customer extends $pb.GeneratedMessage {
  factory Customer({
    $fixnum.Int64? id,
    $core.String? code,
    CustomerType? customerType,
    $core.String? customerName,
    $core.Iterable<Address>? addresses,
    $core.Iterable<Contact>? contacts,
    $core.Iterable<Event>? events,
    CustomerMeta? customerMeta,
    JsonString? customerMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (customerType != null) {
      $result.customerType = customerType;
    }
    if (customerName != null) {
      $result.customerName = customerName;
    }
    if (addresses != null) {
      $result.addresses.addAll(addresses);
    }
    if (contacts != null) {
      $result.contacts.addAll(contacts);
    }
    if (events != null) {
      $result.events.addAll(events);
    }
    if (customerMeta != null) {
      $result.customerMeta = customerMeta;
    }
    if (customerMap != null) {
      $result.customerMap = customerMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Customer._() : super();
  factory Customer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Customer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Customer', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<CustomerType>(3, _omitFieldNames ? '' : 'customerType', $pb.PbFieldType.OE, defaultOrMaker: CustomerType.CUSTOMER_COMPANY, valueOf: CustomerType.valueOf, enumValues: CustomerType.values)
    ..aOS(4, _omitFieldNames ? '' : 'customerName')
    ..pc<Address>(5, _omitFieldNames ? '' : 'addresses', $pb.PbFieldType.PM, subBuilder: Address.create)
    ..pc<Contact>(6, _omitFieldNames ? '' : 'contacts', $pb.PbFieldType.PM, subBuilder: Contact.create)
    ..pc<Event>(7, _omitFieldNames ? '' : 'events', $pb.PbFieldType.PM, subBuilder: Event.create)
    ..aOM<CustomerMeta>(8, _omitFieldNames ? '' : 'customerMeta', subBuilder: CustomerMeta.create)
    ..aOM<JsonString>(9, _omitFieldNames ? '' : 'customerMap', subBuilder: JsonString.create)
    ..aOS(10, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Customer clone() => Customer()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Customer copyWith(void Function(Customer) updates) => super.copyWith((message) => updates(message as Customer)) as Customer;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Customer create() => Customer._();
  Customer createEmptyInstance() => create();
  static $pb.PbList<Customer> createRepeated() => $pb.PbList<Customer>();
  @$core.pragma('dart2js:noInline')
  static Customer getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Customer>(create);
  static Customer? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: CUS1731101982N123 ("CUS" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field. Default value: CUSTOMER_COMPANY
  @$pb.TagNumber(3)
  CustomerType get customerType => $_getN(2);
  @$pb.TagNumber(3)
  set customerType(CustomerType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasCustomerType() => $_has(2);
  @$pb.TagNumber(3)
  void clearCustomerType() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Full name of the customer
  @$pb.TagNumber(4)
  $core.String get customerName => $_getSZ(3);
  @$pb.TagNumber(4)
  set customerName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCustomerName() => $_has(3);
  @$pb.TagNumber(4)
  void clearCustomerName() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<Address> get addresses => $_getList(4);

  @$pb.TagNumber(6)
  $core.List<Contact> get contacts => $_getList(5);

  @$pb.TagNumber(7)
  $core.List<Event> get events => $_getList(6);

  @$pb.TagNumber(8)
  CustomerMeta get customerMeta => $_getN(7);
  @$pb.TagNumber(8)
  set customerMeta(CustomerMeta v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasCustomerMeta() => $_has(7);
  @$pb.TagNumber(8)
  void clearCustomerMeta() => clearField(8);
  @$pb.TagNumber(8)
  CustomerMeta ensureCustomerMeta() => $_ensure(7);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(9)
  JsonString get customerMap => $_getN(8);
  @$pb.TagNumber(9)
  set customerMap(JsonString v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasCustomerMap() => $_has(8);
  @$pb.TagNumber(9)
  void clearCustomerMap() => clearField(9);
  @$pb.TagNumber(9)
  JsonString ensureCustomerMap() => $_ensure(8);

  /// Timestamp of data creation
  @$pb.TagNumber(10)
  $core.String get timeStamp => $_getSZ(9);
  @$pb.TagNumber(10)
  set timeStamp($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasTimeStamp() => $_has(9);
  @$pb.TagNumber(10)
  void clearTimeStamp() => clearField(10);
}

class Customers extends $pb.GeneratedMessage {
  factory Customers({
    $core.Iterable<Customer>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Customers._() : super();
  factory Customers.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Customers.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Customers', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Customer>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Customer.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Customers clone() => Customers()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Customers copyWith(void Function(Customers) updates) => super.copyWith((message) => updates(message as Customers)) as Customers;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Customers create() => Customers._();
  Customers createEmptyInstance() => create();
  static $pb.PbList<Customers> createRepeated() => $pb.PbList<Customers>();
  @$core.pragma('dart2js:noInline')
  static Customers getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Customers>(create);
  static Customers? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Customer> get data => $_getList(0);
}

/// EmployeeMeta message
class EmployeeMeta extends $pb.GeneratedMessage {
  factory EmployeeMeta({
    $core.String? startDate,
    $core.String? endDate,
    $core.bool? inactive,
    $core.String? notes,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (startDate != null) {
      $result.startDate = startDate;
    }
    if (endDate != null) {
      $result.endDate = endDate;
    }
    if (inactive != null) {
      $result.inactive = inactive;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  EmployeeMeta._() : super();
  factory EmployeeMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory EmployeeMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'EmployeeMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'startDate')
    ..aOS(2, _omitFieldNames ? '' : 'endDate')
    ..aOB(3, _omitFieldNames ? '' : 'inactive')
    ..aOS(4, _omitFieldNames ? '' : 'notes')
    ..pPS(5, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  EmployeeMeta clone() => EmployeeMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  EmployeeMeta copyWith(void Function(EmployeeMeta) updates) => super.copyWith((message) => updates(message as EmployeeMeta)) as EmployeeMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static EmployeeMeta create() => EmployeeMeta._();
  EmployeeMeta createEmptyInstance() => create();
  static $pb.PbList<EmployeeMeta> createRepeated() => $pb.PbList<EmployeeMeta>();
  @$core.pragma('dart2js:noInline')
  static EmployeeMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<EmployeeMeta>(create);
  static EmployeeMeta? _defaultInstance;

  /// ISO 8601 date
  @$pb.TagNumber(1)
  $core.String get startDate => $_getSZ(0);
  @$pb.TagNumber(1)
  set startDate($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasStartDate() => $_has(0);
  @$pb.TagNumber(1)
  void clearStartDate() => clearField(1);

  /// ISO 8601 date
  @$pb.TagNumber(2)
  $core.String get endDate => $_getSZ(1);
  @$pb.TagNumber(2)
  set endDate($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasEndDate() => $_has(1);
  @$pb.TagNumber(2)
  void clearEndDate() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get inactive => $_getBF(2);
  @$pb.TagNumber(3)
  set inactive($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasInactive() => $_has(2);
  @$pb.TagNumber(3)
  void clearInactive() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get notes => $_getSZ(3);
  @$pb.TagNumber(4)
  set notes($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasNotes() => $_has(3);
  @$pb.TagNumber(4)
  void clearNotes() => clearField(4);

  /// Additional tags for the employee
  @$pb.TagNumber(5)
  $core.List<$core.String> get tags => $_getList(4);
}

/// Employee message
class Employee extends $pb.GeneratedMessage {
  factory Employee({
    $fixnum.Int64? id,
    $core.String? code,
    Address? address,
    Contact? contact,
    $core.Iterable<Event>? events,
    EmployeeMeta? employeeMeta,
    JsonString? employeeMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (address != null) {
      $result.address = address;
    }
    if (contact != null) {
      $result.contact = contact;
    }
    if (events != null) {
      $result.events.addAll(events);
    }
    if (employeeMeta != null) {
      $result.employeeMeta = employeeMeta;
    }
    if (employeeMap != null) {
      $result.employeeMap = employeeMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Employee._() : super();
  factory Employee.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Employee.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Employee', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..aOM<Address>(3, _omitFieldNames ? '' : 'address', subBuilder: Address.create)
    ..aOM<Contact>(4, _omitFieldNames ? '' : 'contact', subBuilder: Contact.create)
    ..pc<Event>(5, _omitFieldNames ? '' : 'events', $pb.PbFieldType.PM, subBuilder: Event.create)
    ..aOM<EmployeeMeta>(6, _omitFieldNames ? '' : 'employeeMeta', subBuilder: EmployeeMeta.create)
    ..aOM<JsonString>(7, _omitFieldNames ? '' : 'employeeMap', subBuilder: JsonString.create)
    ..aOS(8, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Employee clone() => Employee()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Employee copyWith(void Function(Employee) updates) => super.copyWith((message) => updates(message as Employee)) as Employee;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Employee create() => Employee._();
  Employee createEmptyInstance() => create();
  static $pb.PbList<Employee> createRepeated() => $pb.PbList<Employee>();
  @$core.pragma('dart2js:noInline')
  static Employee getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Employee>(create);
  static Employee? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: EMP1731101982N123 ("EMP" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  @$pb.TagNumber(3)
  Address get address => $_getN(2);
  @$pb.TagNumber(3)
  set address(Address v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasAddress() => $_has(2);
  @$pb.TagNumber(3)
  void clearAddress() => clearField(3);
  @$pb.TagNumber(3)
  Address ensureAddress() => $_ensure(2);

  @$pb.TagNumber(4)
  Contact get contact => $_getN(3);
  @$pb.TagNumber(4)
  set contact(Contact v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasContact() => $_has(3);
  @$pb.TagNumber(4)
  void clearContact() => clearField(4);
  @$pb.TagNumber(4)
  Contact ensureContact() => $_ensure(3);

  @$pb.TagNumber(5)
  $core.List<Event> get events => $_getList(4);

  @$pb.TagNumber(6)
  EmployeeMeta get employeeMeta => $_getN(5);
  @$pb.TagNumber(6)
  set employeeMeta(EmployeeMeta v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasEmployeeMeta() => $_has(5);
  @$pb.TagNumber(6)
  void clearEmployeeMeta() => clearField(6);
  @$pb.TagNumber(6)
  EmployeeMeta ensureEmployeeMeta() => $_ensure(5);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(7)
  JsonString get employeeMap => $_getN(6);
  @$pb.TagNumber(7)
  set employeeMap(JsonString v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasEmployeeMap() => $_has(6);
  @$pb.TagNumber(7)
  void clearEmployeeMap() => clearField(7);
  @$pb.TagNumber(7)
  JsonString ensureEmployeeMap() => $_ensure(6);

  /// Timestamp of data creation
  @$pb.TagNumber(8)
  $core.String get timeStamp => $_getSZ(7);
  @$pb.TagNumber(8)
  set timeStamp($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasTimeStamp() => $_has(7);
  @$pb.TagNumber(8)
  void clearTimeStamp() => clearField(8);
}

class Employees extends $pb.GeneratedMessage {
  factory Employees({
    $core.Iterable<Employee>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Employees._() : super();
  factory Employees.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Employees.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Employees', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Employee>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Employee.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Employees clone() => Employees()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Employees copyWith(void Function(Employees) updates) => super.copyWith((message) => updates(message as Employees)) as Employees;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Employees create() => Employees._();
  Employees createEmptyInstance() => create();
  static $pb.PbList<Employees> createRepeated() => $pb.PbList<Employees>();
  @$core.pragma('dart2js:noInline')
  static Employees getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Employees>(create);
  static Employees? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Employee> get data => $_getList(0);
}

/// ItemMeta message
class ItemMeta extends $pb.GeneratedMessage {
  factory ItemMeta({
    $core.String? unit,
    $core.double? qty,
    $core.double? fxPrice,
    $core.double? netAmount,
    $core.double? discount,
    $core.double? vatAmount,
    $core.double? amount,
    $core.String? description,
    $core.bool? deposit,
    $core.double? ownStock,
    $core.bool? actionPrice,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (unit != null) {
      $result.unit = unit;
    }
    if (qty != null) {
      $result.qty = qty;
    }
    if (fxPrice != null) {
      $result.fxPrice = fxPrice;
    }
    if (netAmount != null) {
      $result.netAmount = netAmount;
    }
    if (discount != null) {
      $result.discount = discount;
    }
    if (vatAmount != null) {
      $result.vatAmount = vatAmount;
    }
    if (amount != null) {
      $result.amount = amount;
    }
    if (description != null) {
      $result.description = description;
    }
    if (deposit != null) {
      $result.deposit = deposit;
    }
    if (ownStock != null) {
      $result.ownStock = ownStock;
    }
    if (actionPrice != null) {
      $result.actionPrice = actionPrice;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  ItemMeta._() : super();
  factory ItemMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ItemMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ItemMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'unit')
    ..a<$core.double>(2, _omitFieldNames ? '' : 'qty', $pb.PbFieldType.OD)
    ..a<$core.double>(3, _omitFieldNames ? '' : 'fxPrice', $pb.PbFieldType.OD)
    ..a<$core.double>(4, _omitFieldNames ? '' : 'netAmount', $pb.PbFieldType.OD)
    ..a<$core.double>(5, _omitFieldNames ? '' : 'discount', $pb.PbFieldType.OD)
    ..a<$core.double>(6, _omitFieldNames ? '' : 'vatAmount', $pb.PbFieldType.OD)
    ..a<$core.double>(7, _omitFieldNames ? '' : 'amount', $pb.PbFieldType.OD)
    ..aOS(8, _omitFieldNames ? '' : 'description')
    ..aOB(9, _omitFieldNames ? '' : 'deposit')
    ..a<$core.double>(10, _omitFieldNames ? '' : 'ownStock', $pb.PbFieldType.OD)
    ..aOB(11, _omitFieldNames ? '' : 'actionPrice')
    ..pPS(12, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ItemMeta clone() => ItemMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ItemMeta copyWith(void Function(ItemMeta) updates) => super.copyWith((message) => updates(message as ItemMeta)) as ItemMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ItemMeta create() => ItemMeta._();
  ItemMeta createEmptyInstance() => create();
  static $pb.PbList<ItemMeta> createRepeated() => $pb.PbList<ItemMeta>();
  @$core.pragma('dart2js:noInline')
  static ItemMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ItemMeta>(create);
  static ItemMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get unit => $_getSZ(0);
  @$pb.TagNumber(1)
  set unit($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUnit() => $_has(0);
  @$pb.TagNumber(1)
  void clearUnit() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get qty => $_getN(1);
  @$pb.TagNumber(2)
  set qty($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasQty() => $_has(1);
  @$pb.TagNumber(2)
  void clearQty() => clearField(2);

  @$pb.TagNumber(3)
  $core.double get fxPrice => $_getN(2);
  @$pb.TagNumber(3)
  set fxPrice($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasFxPrice() => $_has(2);
  @$pb.TagNumber(3)
  void clearFxPrice() => clearField(3);

  @$pb.TagNumber(4)
  $core.double get netAmount => $_getN(3);
  @$pb.TagNumber(4)
  set netAmount($core.double v) { $_setDouble(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasNetAmount() => $_has(3);
  @$pb.TagNumber(4)
  void clearNetAmount() => clearField(4);

  @$pb.TagNumber(5)
  $core.double get discount => $_getN(4);
  @$pb.TagNumber(5)
  set discount($core.double v) { $_setDouble(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasDiscount() => $_has(4);
  @$pb.TagNumber(5)
  void clearDiscount() => clearField(5);

  @$pb.TagNumber(6)
  $core.double get vatAmount => $_getN(5);
  @$pb.TagNumber(6)
  set vatAmount($core.double v) { $_setDouble(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasVatAmount() => $_has(5);
  @$pb.TagNumber(6)
  void clearVatAmount() => clearField(6);

  @$pb.TagNumber(7)
  $core.double get amount => $_getN(6);
  @$pb.TagNumber(7)
  set amount($core.double v) { $_setDouble(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasAmount() => $_has(6);
  @$pb.TagNumber(7)
  void clearAmount() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get description => $_getSZ(7);
  @$pb.TagNumber(8)
  set description($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasDescription() => $_has(7);
  @$pb.TagNumber(8)
  void clearDescription() => clearField(8);

  @$pb.TagNumber(9)
  $core.bool get deposit => $_getBF(8);
  @$pb.TagNumber(9)
  set deposit($core.bool v) { $_setBool(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasDeposit() => $_has(8);
  @$pb.TagNumber(9)
  void clearDeposit() => clearField(9);

  @$pb.TagNumber(10)
  $core.double get ownStock => $_getN(9);
  @$pb.TagNumber(10)
  set ownStock($core.double v) { $_setDouble(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasOwnStock() => $_has(9);
  @$pb.TagNumber(10)
  void clearOwnStock() => clearField(10);

  @$pb.TagNumber(11)
  $core.bool get actionPrice => $_getBF(10);
  @$pb.TagNumber(11)
  set actionPrice($core.bool v) { $_setBool(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasActionPrice() => $_has(10);
  @$pb.TagNumber(11)
  void clearActionPrice() => clearField(11);

  /// Additional tags for the item
  @$pb.TagNumber(12)
  $core.List<$core.String> get tags => $_getList(11);
}

/// Item message
class Item extends $pb.GeneratedMessage {
  factory Item({
    $fixnum.Int64? id,
    $core.String? code,
    $core.String? transCode,
    $core.String? productCode,
    $core.String? taxCode,
    ItemMeta? itemMeta,
    JsonString? itemMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (transCode != null) {
      $result.transCode = transCode;
    }
    if (productCode != null) {
      $result.productCode = productCode;
    }
    if (taxCode != null) {
      $result.taxCode = taxCode;
    }
    if (itemMeta != null) {
      $result.itemMeta = itemMeta;
    }
    if (itemMap != null) {
      $result.itemMap = itemMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Item._() : super();
  factory Item.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Item.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Item', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..aOS(3, _omitFieldNames ? '' : 'transCode')
    ..aOS(4, _omitFieldNames ? '' : 'productCode')
    ..aOS(5, _omitFieldNames ? '' : 'taxCode')
    ..aOM<ItemMeta>(6, _omitFieldNames ? '' : 'itemMeta', subBuilder: ItemMeta.create)
    ..aOM<JsonString>(7, _omitFieldNames ? '' : 'itemMap', subBuilder: JsonString.create)
    ..aOS(8, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Item clone() => Item()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Item copyWith(void Function(Item) updates) => super.copyWith((message) => updates(message as Item)) as Item;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Item create() => Item._();
  Item createEmptyInstance() => create();
  static $pb.PbList<Item> createRepeated() => $pb.PbList<Item>();
  @$core.pragma('dart2js:noInline')
  static Item getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Item>(create);
  static Item? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: ITM1731101982N123 ("ITM" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Trans">Trans</a>.Code
  @$pb.TagNumber(3)
  $core.String get transCode => $_getSZ(2);
  @$pb.TagNumber(3)
  set transCode($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTransCode() => $_has(2);
  @$pb.TagNumber(3)
  void clearTransCode() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Product">Product</a>.Code
  @$pb.TagNumber(4)
  $core.String get productCode => $_getSZ(3);
  @$pb.TagNumber(4)
  set productCode($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasProductCode() => $_has(3);
  @$pb.TagNumber(4)
  void clearProductCode() => clearField(4);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Tax">Tax</a>.Code
  @$pb.TagNumber(5)
  $core.String get taxCode => $_getSZ(4);
  @$pb.TagNumber(5)
  set taxCode($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTaxCode() => $_has(4);
  @$pb.TagNumber(5)
  void clearTaxCode() => clearField(5);

  @$pb.TagNumber(6)
  ItemMeta get itemMeta => $_getN(5);
  @$pb.TagNumber(6)
  set itemMeta(ItemMeta v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasItemMeta() => $_has(5);
  @$pb.TagNumber(6)
  void clearItemMeta() => clearField(6);
  @$pb.TagNumber(6)
  ItemMeta ensureItemMeta() => $_ensure(5);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(7)
  JsonString get itemMap => $_getN(6);
  @$pb.TagNumber(7)
  set itemMap(JsonString v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasItemMap() => $_has(6);
  @$pb.TagNumber(7)
  void clearItemMap() => clearField(7);
  @$pb.TagNumber(7)
  JsonString ensureItemMap() => $_ensure(6);

  /// Timestamp of data creation
  @$pb.TagNumber(8)
  $core.String get timeStamp => $_getSZ(7);
  @$pb.TagNumber(8)
  set timeStamp($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasTimeStamp() => $_has(7);
  @$pb.TagNumber(8)
  void clearTimeStamp() => clearField(8);
}

class Items extends $pb.GeneratedMessage {
  factory Items({
    $core.Iterable<Item>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Items._() : super();
  factory Items.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Items.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Items', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Item>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Item.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Items clone() => Items()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Items copyWith(void Function(Items) updates) => super.copyWith((message) => updates(message as Items)) as Items;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Items create() => Items._();
  Items createEmptyInstance() => create();
  static $pb.PbList<Items> createRepeated() => $pb.PbList<Items>();
  @$core.pragma('dart2js:noInline')
  static Items getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Items>(create);
  static Items? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Item> get data => $_getList(0);
}

/// LinkMeta message
class LinkMeta extends $pb.GeneratedMessage {
  factory LinkMeta({
    $core.double? qty,
    $core.double? rate,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (qty != null) {
      $result.qty = qty;
    }
    if (rate != null) {
      $result.rate = rate;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  LinkMeta._() : super();
  factory LinkMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LinkMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'LinkMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..a<$core.double>(1, _omitFieldNames ? '' : 'qty', $pb.PbFieldType.OD)
    ..a<$core.double>(2, _omitFieldNames ? '' : 'rate', $pb.PbFieldType.OD)
    ..pPS(3, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  LinkMeta clone() => LinkMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  LinkMeta copyWith(void Function(LinkMeta) updates) => super.copyWith((message) => updates(message as LinkMeta)) as LinkMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static LinkMeta create() => LinkMeta._();
  LinkMeta createEmptyInstance() => create();
  static $pb.PbList<LinkMeta> createRepeated() => $pb.PbList<LinkMeta>();
  @$core.pragma('dart2js:noInline')
  static LinkMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LinkMeta>(create);
  static LinkMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.double get qty => $_getN(0);
  @$pb.TagNumber(1)
  set qty($core.double v) { $_setDouble(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasQty() => $_has(0);
  @$pb.TagNumber(1)
  void clearQty() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get rate => $_getN(1);
  @$pb.TagNumber(2)
  set rate($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRate() => $_has(1);
  @$pb.TagNumber(2)
  void clearRate() => clearField(2);

  /// Additional tags for the link
  @$pb.TagNumber(3)
  $core.List<$core.String> get tags => $_getList(2);
}

/// Link message
class Link extends $pb.GeneratedMessage {
  factory Link({
    $fixnum.Int64? id,
    $core.String? code,
    LinkType? linkType1,
    $core.String? linkCode1,
    LinkType? linkType2,
    $core.String? linkCode2,
    LinkMeta? linkMeta,
    JsonString? linkMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (linkType1 != null) {
      $result.linkType1 = linkType1;
    }
    if (linkCode1 != null) {
      $result.linkCode1 = linkCode1;
    }
    if (linkType2 != null) {
      $result.linkType2 = linkType2;
    }
    if (linkCode2 != null) {
      $result.linkCode2 = linkCode2;
    }
    if (linkMeta != null) {
      $result.linkMeta = linkMeta;
    }
    if (linkMap != null) {
      $result.linkMap = linkMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Link._() : super();
  factory Link.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Link.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Link', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<LinkType>(3, _omitFieldNames ? '' : 'linkType1', $pb.PbFieldType.OE, protoName: 'link_type_1', defaultOrMaker: LinkType.LINK_CUSTOMER, valueOf: LinkType.valueOf, enumValues: LinkType.values)
    ..aOS(4, _omitFieldNames ? '' : 'linkCode1', protoName: 'link_code_1')
    ..e<LinkType>(5, _omitFieldNames ? '' : 'linkType2', $pb.PbFieldType.OE, protoName: 'link_type_2', defaultOrMaker: LinkType.LINK_CUSTOMER, valueOf: LinkType.valueOf, enumValues: LinkType.values)
    ..aOS(6, _omitFieldNames ? '' : 'linkCode2', protoName: 'link_code_2')
    ..aOM<LinkMeta>(7, _omitFieldNames ? '' : 'linkMeta', subBuilder: LinkMeta.create)
    ..aOM<JsonString>(8, _omitFieldNames ? '' : 'linkMap', subBuilder: JsonString.create)
    ..aOS(9, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Link clone() => Link()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Link copyWith(void Function(Link) updates) => super.copyWith((message) => updates(message as Link)) as Link;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Link create() => Link._();
  Link createEmptyInstance() => create();
  static $pb.PbList<Link> createRepeated() => $pb.PbList<Link>();
  @$core.pragma('dart2js:noInline')
  static Link getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Link>(create);
  static Link? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: LNK1731101982N123 ("LNK" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(3)
  LinkType get linkType1 => $_getN(2);
  @$pb.TagNumber(3)
  set linkType1(LinkType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasLinkType1() => $_has(2);
  @$pb.TagNumber(3)
  void clearLinkType1() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to LinkType1.code
  @$pb.TagNumber(4)
  $core.String get linkCode1 => $_getSZ(3);
  @$pb.TagNumber(4)
  set linkCode1($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasLinkCode1() => $_has(3);
  @$pb.TagNumber(4)
  void clearLinkCode1() => clearField(4);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(5)
  LinkType get linkType2 => $_getN(4);
  @$pb.TagNumber(5)
  set linkType2(LinkType v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasLinkType2() => $_has(4);
  @$pb.TagNumber(5)
  void clearLinkType2() => clearField(5);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to LinkType2.code
  @$pb.TagNumber(6)
  $core.String get linkCode2 => $_getSZ(5);
  @$pb.TagNumber(6)
  set linkCode2($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasLinkCode2() => $_has(5);
  @$pb.TagNumber(6)
  void clearLinkCode2() => clearField(6);

  @$pb.TagNumber(7)
  LinkMeta get linkMeta => $_getN(6);
  @$pb.TagNumber(7)
  set linkMeta(LinkMeta v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasLinkMeta() => $_has(6);
  @$pb.TagNumber(7)
  void clearLinkMeta() => clearField(7);
  @$pb.TagNumber(7)
  LinkMeta ensureLinkMeta() => $_ensure(6);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(8)
  JsonString get linkMap => $_getN(7);
  @$pb.TagNumber(8)
  set linkMap(JsonString v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasLinkMap() => $_has(7);
  @$pb.TagNumber(8)
  void clearLinkMap() => clearField(8);
  @$pb.TagNumber(8)
  JsonString ensureLinkMap() => $_ensure(7);

  /// Timestamp of data creation
  @$pb.TagNumber(9)
  $core.String get timeStamp => $_getSZ(8);
  @$pb.TagNumber(9)
  set timeStamp($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasTimeStamp() => $_has(8);
  @$pb.TagNumber(9)
  void clearTimeStamp() => clearField(9);
}

class Links extends $pb.GeneratedMessage {
  factory Links({
    $core.Iterable<Link>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Links._() : super();
  factory Links.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Links.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Links', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Link>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Link.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Links clone() => Links()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Links copyWith(void Function(Links) updates) => super.copyWith((message) => updates(message as Links)) as Links;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Links create() => Links._();
  Links createEmptyInstance() => create();
  static $pb.PbList<Links> createRepeated() => $pb.PbList<Links>();
  @$core.pragma('dart2js:noInline')
  static Links getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Links>(create);
  static Links? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Link> get data => $_getList(0);
}

/// Log message
class Log extends $pb.GeneratedMessage {
  factory Log({
    $fixnum.Int64? id,
    $core.String? code,
    LogType? logType,
    $core.String? refType,
    $core.String? refCode,
    $core.String? authCode,
    JsonString? logMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (logType != null) {
      $result.logType = logType;
    }
    if (refType != null) {
      $result.refType = refType;
    }
    if (refCode != null) {
      $result.refCode = refCode;
    }
    if (authCode != null) {
      $result.authCode = authCode;
    }
    if (logMap != null) {
      $result.logMap = logMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Log._() : super();
  factory Log.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Log.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Log', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<LogType>(3, _omitFieldNames ? '' : 'logType', $pb.PbFieldType.OE, defaultOrMaker: LogType.LOG_INSERT, valueOf: LogType.valueOf, enumValues: LogType.values)
    ..aOS(4, _omitFieldNames ? '' : 'refType')
    ..aOS(5, _omitFieldNames ? '' : 'refCode')
    ..aOS(6, _omitFieldNames ? '' : 'authCode')
    ..aOM<JsonString>(7, _omitFieldNames ? '' : 'logMap', subBuilder: JsonString.create)
    ..aOS(8, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Log clone() => Log()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Log copyWith(void Function(Log) updates) => super.copyWith((message) => updates(message as Log)) as Log;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Log create() => Log._();
  Log createEmptyInstance() => create();
  static $pb.PbList<Log> createRepeated() => $pb.PbList<Log>();
  @$core.pragma('dart2js:noInline')
  static Log getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Log>(create);
  static Log? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: LOG1731101982N123 ("LOG" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// ENUM field.
  @$pb.TagNumber(3)
  LogType get logType => $_getN(2);
  @$pb.TagNumber(3)
  set logType(LogType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasLogType() => $_has(2);
  @$pb.TagNumber(3)
  void clearLogType() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get refType => $_getSZ(3);
  @$pb.TagNumber(4)
  set refType($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRefType() => $_has(3);
  @$pb.TagNumber(4)
  void clearRefType() => clearField(4);

  /// Reference to RefType.code
  @$pb.TagNumber(5)
  $core.String get refCode => $_getSZ(4);
  @$pb.TagNumber(5)
  set refCode($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasRefCode() => $_has(4);
  @$pb.TagNumber(5)
  void clearRefCode() => clearField(5);

  /// Reference to <a href="#Auth">Auth</a>.Code
  @$pb.TagNumber(6)
  $core.String get authCode => $_getSZ(5);
  @$pb.TagNumber(6)
  set authCode($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasAuthCode() => $_has(5);
  @$pb.TagNumber(6)
  void clearAuthCode() => clearField(6);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(7)
  JsonString get logMap => $_getN(6);
  @$pb.TagNumber(7)
  set logMap(JsonString v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasLogMap() => $_has(6);
  @$pb.TagNumber(7)
  void clearLogMap() => clearField(7);
  @$pb.TagNumber(7)
  JsonString ensureLogMap() => $_ensure(6);

  /// Timestamp of data creation
  @$pb.TagNumber(8)
  $core.String get timeStamp => $_getSZ(7);
  @$pb.TagNumber(8)
  set timeStamp($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasTimeStamp() => $_has(7);
  @$pb.TagNumber(8)
  void clearTimeStamp() => clearField(8);
}

class Logs extends $pb.GeneratedMessage {
  factory Logs({
    $core.Iterable<Log>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Logs._() : super();
  factory Logs.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Logs.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Logs', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Log>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Log.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Logs clone() => Logs()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Logs copyWith(void Function(Logs) updates) => super.copyWith((message) => updates(message as Logs)) as Logs;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Logs create() => Logs._();
  Logs createEmptyInstance() => create();
  static $pb.PbList<Logs> createRepeated() => $pb.PbList<Logs>();
  @$core.pragma('dart2js:noInline')
  static Logs getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Logs>(create);
  static Logs? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Log> get data => $_getList(0);
}

/// MovementMeta message
class MovementMeta extends $pb.GeneratedMessage {
  factory MovementMeta({
    $core.double? qty,
    $core.String? notes,
    $core.bool? shared,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (qty != null) {
      $result.qty = qty;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (shared != null) {
      $result.shared = shared;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  MovementMeta._() : super();
  factory MovementMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MovementMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'MovementMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..a<$core.double>(1, _omitFieldNames ? '' : 'qty', $pb.PbFieldType.OD)
    ..aOS(2, _omitFieldNames ? '' : 'notes')
    ..aOB(3, _omitFieldNames ? '' : 'shared')
    ..pPS(4, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MovementMeta clone() => MovementMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MovementMeta copyWith(void Function(MovementMeta) updates) => super.copyWith((message) => updates(message as MovementMeta)) as MovementMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static MovementMeta create() => MovementMeta._();
  MovementMeta createEmptyInstance() => create();
  static $pb.PbList<MovementMeta> createRepeated() => $pb.PbList<MovementMeta>();
  @$core.pragma('dart2js:noInline')
  static MovementMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MovementMeta>(create);
  static MovementMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.double get qty => $_getN(0);
  @$pb.TagNumber(1)
  set qty($core.double v) { $_setDouble(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasQty() => $_has(0);
  @$pb.TagNumber(1)
  void clearQty() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get notes => $_getSZ(1);
  @$pb.TagNumber(2)
  set notes($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNotes() => $_has(1);
  @$pb.TagNumber(2)
  void clearNotes() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get shared => $_getBF(2);
  @$pb.TagNumber(3)
  set shared($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasShared() => $_has(2);
  @$pb.TagNumber(3)
  void clearShared() => clearField(3);

  /// Additional tags for the movement
  @$pb.TagNumber(4)
  $core.List<$core.String> get tags => $_getList(3);
}

/// Movement message
class Movement extends $pb.GeneratedMessage {
  factory Movement({
    $fixnum.Int64? id,
    $core.String? code,
    MovementType? movementType,
    $core.String? shippingTime,
    $core.String? transCode,
    $core.String? productCode,
    $core.String? toolCode,
    $core.String? placeCode,
    $core.String? itemCode,
    $core.String? movementCode,
    MovementMeta? movementMeta,
    JsonString? movementMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (movementType != null) {
      $result.movementType = movementType;
    }
    if (shippingTime != null) {
      $result.shippingTime = shippingTime;
    }
    if (transCode != null) {
      $result.transCode = transCode;
    }
    if (productCode != null) {
      $result.productCode = productCode;
    }
    if (toolCode != null) {
      $result.toolCode = toolCode;
    }
    if (placeCode != null) {
      $result.placeCode = placeCode;
    }
    if (itemCode != null) {
      $result.itemCode = itemCode;
    }
    if (movementCode != null) {
      $result.movementCode = movementCode;
    }
    if (movementMeta != null) {
      $result.movementMeta = movementMeta;
    }
    if (movementMap != null) {
      $result.movementMap = movementMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Movement._() : super();
  factory Movement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Movement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Movement', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<MovementType>(3, _omitFieldNames ? '' : 'movementType', $pb.PbFieldType.OE, defaultOrMaker: MovementType.MOVEMENT_INVENTORY, valueOf: MovementType.valueOf, enumValues: MovementType.values)
    ..aOS(4, _omitFieldNames ? '' : 'shippingTime')
    ..aOS(5, _omitFieldNames ? '' : 'transCode')
    ..aOS(6, _omitFieldNames ? '' : 'productCode')
    ..aOS(7, _omitFieldNames ? '' : 'toolCode')
    ..aOS(8, _omitFieldNames ? '' : 'placeCode')
    ..aOS(9, _omitFieldNames ? '' : 'itemCode')
    ..aOS(10, _omitFieldNames ? '' : 'movementCode')
    ..aOM<MovementMeta>(11, _omitFieldNames ? '' : 'movementMeta', subBuilder: MovementMeta.create)
    ..aOM<JsonString>(12, _omitFieldNames ? '' : 'movementMap', subBuilder: JsonString.create)
    ..aOS(13, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Movement clone() => Movement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Movement copyWith(void Function(Movement) updates) => super.copyWith((message) => updates(message as Movement)) as Movement;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Movement create() => Movement._();
  Movement createEmptyInstance() => create();
  static $pb.PbList<Movement> createRepeated() => $pb.PbList<Movement>();
  @$core.pragma('dart2js:noInline')
  static Movement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Movement>(create);
  static Movement? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: MOV1731101982N123 ("MOV" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(3)
  MovementType get movementType => $_getN(2);
  @$pb.TagNumber(3)
  set movementType(MovementType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasMovementType() => $_has(2);
  @$pb.TagNumber(3)
  void clearMovementType() => clearField(3);

  /// ISO 8601 date-time
  @$pb.TagNumber(4)
  $core.String get shippingTime => $_getSZ(3);
  @$pb.TagNumber(4)
  set shippingTime($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasShippingTime() => $_has(3);
  @$pb.TagNumber(4)
  void clearShippingTime() => clearField(4);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Trans">Trans</a>.Code
  @$pb.TagNumber(5)
  $core.String get transCode => $_getSZ(4);
  @$pb.TagNumber(5)
  set transCode($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTransCode() => $_has(4);
  @$pb.TagNumber(5)
  void clearTransCode() => clearField(5);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Product">Product</a>.Code
  @$pb.TagNumber(6)
  $core.String get productCode => $_getSZ(5);
  @$pb.TagNumber(6)
  set productCode($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasProductCode() => $_has(5);
  @$pb.TagNumber(6)
  void clearProductCode() => clearField(6);

  /// Reference to <a href="#Tool">Tool</a>.Code
  @$pb.TagNumber(7)
  $core.String get toolCode => $_getSZ(6);
  @$pb.TagNumber(7)
  set toolCode($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasToolCode() => $_has(6);
  @$pb.TagNumber(7)
  void clearToolCode() => clearField(7);

  /// Reference to <a href="#Place">Place</a>.Code
  @$pb.TagNumber(8)
  $core.String get placeCode => $_getSZ(7);
  @$pb.TagNumber(8)
  set placeCode($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasPlaceCode() => $_has(7);
  @$pb.TagNumber(8)
  void clearPlaceCode() => clearField(8);

  /// Reference to <a href="#Item">Item</a>.Code
  @$pb.TagNumber(9)
  $core.String get itemCode => $_getSZ(8);
  @$pb.TagNumber(9)
  set itemCode($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasItemCode() => $_has(8);
  @$pb.TagNumber(9)
  void clearItemCode() => clearField(9);

  /// Reference to <a href="#Movement">Movement</a>.Code
  @$pb.TagNumber(10)
  $core.String get movementCode => $_getSZ(9);
  @$pb.TagNumber(10)
  set movementCode($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasMovementCode() => $_has(9);
  @$pb.TagNumber(10)
  void clearMovementCode() => clearField(10);

  @$pb.TagNumber(11)
  MovementMeta get movementMeta => $_getN(10);
  @$pb.TagNumber(11)
  set movementMeta(MovementMeta v) { setField(11, v); }
  @$pb.TagNumber(11)
  $core.bool hasMovementMeta() => $_has(10);
  @$pb.TagNumber(11)
  void clearMovementMeta() => clearField(11);
  @$pb.TagNumber(11)
  MovementMeta ensureMovementMeta() => $_ensure(10);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(12)
  JsonString get movementMap => $_getN(11);
  @$pb.TagNumber(12)
  set movementMap(JsonString v) { setField(12, v); }
  @$pb.TagNumber(12)
  $core.bool hasMovementMap() => $_has(11);
  @$pb.TagNumber(12)
  void clearMovementMap() => clearField(12);
  @$pb.TagNumber(12)
  JsonString ensureMovementMap() => $_ensure(11);

  /// Timestamp of data creation
  @$pb.TagNumber(13)
  $core.String get timeStamp => $_getSZ(12);
  @$pb.TagNumber(13)
  set timeStamp($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasTimeStamp() => $_has(12);
  @$pb.TagNumber(13)
  void clearTimeStamp() => clearField(13);
}

class Movements extends $pb.GeneratedMessage {
  factory Movements({
    $core.Iterable<Movement>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Movements._() : super();
  factory Movements.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Movements.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Movements', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Movement>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Movement.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Movements clone() => Movements()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Movements copyWith(void Function(Movements) updates) => super.copyWith((message) => updates(message as Movements)) as Movements;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Movements create() => Movements._();
  Movements createEmptyInstance() => create();
  static $pb.PbList<Movements> createRepeated() => $pb.PbList<Movements>();
  @$core.pragma('dart2js:noInline')
  static Movements getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Movements>(create);
  static Movements? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Movement> get data => $_getList(0);
}

/// PaymentMeta message
class PaymentMeta extends $pb.GeneratedMessage {
  factory PaymentMeta({
    $core.double? amount,
    $core.String? notes,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (amount != null) {
      $result.amount = amount;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  PaymentMeta._() : super();
  factory PaymentMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PaymentMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PaymentMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..a<$core.double>(1, _omitFieldNames ? '' : 'amount', $pb.PbFieldType.OD)
    ..aOS(2, _omitFieldNames ? '' : 'notes')
    ..pPS(3, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PaymentMeta clone() => PaymentMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PaymentMeta copyWith(void Function(PaymentMeta) updates) => super.copyWith((message) => updates(message as PaymentMeta)) as PaymentMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PaymentMeta create() => PaymentMeta._();
  PaymentMeta createEmptyInstance() => create();
  static $pb.PbList<PaymentMeta> createRepeated() => $pb.PbList<PaymentMeta>();
  @$core.pragma('dart2js:noInline')
  static PaymentMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PaymentMeta>(create);
  static PaymentMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.double get amount => $_getN(0);
  @$pb.TagNumber(1)
  set amount($core.double v) { $_setDouble(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAmount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAmount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get notes => $_getSZ(1);
  @$pb.TagNumber(2)
  set notes($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNotes() => $_has(1);
  @$pb.TagNumber(2)
  void clearNotes() => clearField(2);

  /// Additional tags for the payment
  @$pb.TagNumber(3)
  $core.List<$core.String> get tags => $_getList(2);
}

/// Payment message
class Payment extends $pb.GeneratedMessage {
  factory Payment({
    $fixnum.Int64? id,
    $core.String? code,
    $core.String? paidDate,
    $core.String? transCode,
    PaymentMeta? paymentMeta,
    JsonString? paymentMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (paidDate != null) {
      $result.paidDate = paidDate;
    }
    if (transCode != null) {
      $result.transCode = transCode;
    }
    if (paymentMeta != null) {
      $result.paymentMeta = paymentMeta;
    }
    if (paymentMap != null) {
      $result.paymentMap = paymentMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Payment._() : super();
  factory Payment.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Payment.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Payment', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..aOS(3, _omitFieldNames ? '' : 'paidDate')
    ..aOS(4, _omitFieldNames ? '' : 'transCode')
    ..aOM<PaymentMeta>(5, _omitFieldNames ? '' : 'paymentMeta', subBuilder: PaymentMeta.create)
    ..aOM<JsonString>(6, _omitFieldNames ? '' : 'paymentMap', subBuilder: JsonString.create)
    ..aOS(7, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Payment clone() => Payment()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Payment copyWith(void Function(Payment) updates) => super.copyWith((message) => updates(message as Payment)) as Payment;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Payment create() => Payment._();
  Payment createEmptyInstance() => create();
  static $pb.PbList<Payment> createRepeated() => $pb.PbList<Payment>();
  @$core.pragma('dart2js:noInline')
  static Payment getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Payment>(create);
  static Payment? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: PMT1731101982N123 ("PMT" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ISO 8601 date-time
  @$pb.TagNumber(3)
  $core.String get paidDate => $_getSZ(2);
  @$pb.TagNumber(3)
  set paidDate($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasPaidDate() => $_has(2);
  @$pb.TagNumber(3)
  void clearPaidDate() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Trans">Trans</a>.Code
  @$pb.TagNumber(4)
  $core.String get transCode => $_getSZ(3);
  @$pb.TagNumber(4)
  set transCode($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTransCode() => $_has(3);
  @$pb.TagNumber(4)
  void clearTransCode() => clearField(4);

  @$pb.TagNumber(5)
  PaymentMeta get paymentMeta => $_getN(4);
  @$pb.TagNumber(5)
  set paymentMeta(PaymentMeta v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasPaymentMeta() => $_has(4);
  @$pb.TagNumber(5)
  void clearPaymentMeta() => clearField(5);
  @$pb.TagNumber(5)
  PaymentMeta ensurePaymentMeta() => $_ensure(4);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(6)
  JsonString get paymentMap => $_getN(5);
  @$pb.TagNumber(6)
  set paymentMap(JsonString v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasPaymentMap() => $_has(5);
  @$pb.TagNumber(6)
  void clearPaymentMap() => clearField(6);
  @$pb.TagNumber(6)
  JsonString ensurePaymentMap() => $_ensure(5);

  /// Timestamp of data creation
  @$pb.TagNumber(7)
  $core.String get timeStamp => $_getSZ(6);
  @$pb.TagNumber(7)
  set timeStamp($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasTimeStamp() => $_has(6);
  @$pb.TagNumber(7)
  void clearTimeStamp() => clearField(7);
}

class Payments extends $pb.GeneratedMessage {
  factory Payments({
    $core.Iterable<Payment>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Payments._() : super();
  factory Payments.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Payments.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Payments', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Payment>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Payment.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Payments clone() => Payments()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Payments copyWith(void Function(Payments) updates) => super.copyWith((message) => updates(message as Payments)) as Payments;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Payments create() => Payments._();
  Payments createEmptyInstance() => create();
  static $pb.PbList<Payments> createRepeated() => $pb.PbList<Payments>();
  @$core.pragma('dart2js:noInline')
  static Payments getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Payments>(create);
  static Payments? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Payment> get data => $_getList(0);
}

/// PlaceMeta message
class PlaceMeta extends $pb.GeneratedMessage {
  factory PlaceMeta({
    $core.String? notes,
    $core.bool? inactive,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (notes != null) {
      $result.notes = notes;
    }
    if (inactive != null) {
      $result.inactive = inactive;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  PlaceMeta._() : super();
  factory PlaceMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PlaceMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PlaceMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'notes')
    ..aOB(2, _omitFieldNames ? '' : 'inactive')
    ..pPS(3, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PlaceMeta clone() => PlaceMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PlaceMeta copyWith(void Function(PlaceMeta) updates) => super.copyWith((message) => updates(message as PlaceMeta)) as PlaceMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PlaceMeta create() => PlaceMeta._();
  PlaceMeta createEmptyInstance() => create();
  static $pb.PbList<PlaceMeta> createRepeated() => $pb.PbList<PlaceMeta>();
  @$core.pragma('dart2js:noInline')
  static PlaceMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PlaceMeta>(create);
  static PlaceMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get notes => $_getSZ(0);
  @$pb.TagNumber(1)
  set notes($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasNotes() => $_has(0);
  @$pb.TagNumber(1)
  void clearNotes() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get inactive => $_getBF(1);
  @$pb.TagNumber(2)
  set inactive($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasInactive() => $_has(1);
  @$pb.TagNumber(2)
  void clearInactive() => clearField(2);

  /// Additional tags for the place
  @$pb.TagNumber(3)
  $core.List<$core.String> get tags => $_getList(2);
}

/// Place message
class Place extends $pb.GeneratedMessage {
  factory Place({
    $fixnum.Int64? id,
    $core.String? code,
    PlaceType? placeType,
    $core.String? placeName,
    $core.String? currencyCode,
    Address? address,
    $core.Iterable<Contact>? contacts,
    $core.Iterable<Event>? events,
    PlaceMeta? placeMeta,
    JsonString? placeMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (placeType != null) {
      $result.placeType = placeType;
    }
    if (placeName != null) {
      $result.placeName = placeName;
    }
    if (currencyCode != null) {
      $result.currencyCode = currencyCode;
    }
    if (address != null) {
      $result.address = address;
    }
    if (contacts != null) {
      $result.contacts.addAll(contacts);
    }
    if (events != null) {
      $result.events.addAll(events);
    }
    if (placeMeta != null) {
      $result.placeMeta = placeMeta;
    }
    if (placeMap != null) {
      $result.placeMap = placeMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Place._() : super();
  factory Place.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Place.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Place', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<PlaceType>(3, _omitFieldNames ? '' : 'placeType', $pb.PbFieldType.OE, defaultOrMaker: PlaceType.PLACE_WAREHOUSE, valueOf: PlaceType.valueOf, enumValues: PlaceType.values)
    ..aOS(4, _omitFieldNames ? '' : 'placeName')
    ..aOS(5, _omitFieldNames ? '' : 'currencyCode')
    ..aOM<Address>(6, _omitFieldNames ? '' : 'address', subBuilder: Address.create)
    ..pc<Contact>(7, _omitFieldNames ? '' : 'contacts', $pb.PbFieldType.PM, subBuilder: Contact.create)
    ..pc<Event>(8, _omitFieldNames ? '' : 'events', $pb.PbFieldType.PM, subBuilder: Event.create)
    ..aOM<PlaceMeta>(9, _omitFieldNames ? '' : 'placeMeta', subBuilder: PlaceMeta.create)
    ..aOM<JsonString>(10, _omitFieldNames ? '' : 'placeMap', subBuilder: JsonString.create)
    ..aOS(11, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Place clone() => Place()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Place copyWith(void Function(Place) updates) => super.copyWith((message) => updates(message as Place)) as Place;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Place create() => Place._();
  Place createEmptyInstance() => create();
  static $pb.PbList<Place> createRepeated() => $pb.PbList<Place>();
  @$core.pragma('dart2js:noInline')
  static Place getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Place>(create);
  static Place? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: PLA1731101982N123 ("PLA" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(3)
  PlaceType get placeType => $_getN(2);
  @$pb.TagNumber(3)
  set placeType(PlaceType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasPlaceType() => $_has(2);
  @$pb.TagNumber(3)
  void clearPlaceType() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// The full name of the place.
  @$pb.TagNumber(4)
  $core.String get placeName => $_getSZ(3);
  @$pb.TagNumber(4)
  set placeName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPlaceName() => $_has(3);
  @$pb.TagNumber(4)
  void clearPlaceName() => clearField(4);

  /// Reference to <a href="#Currency">Currency</a>.Code
  @$pb.TagNumber(5)
  $core.String get currencyCode => $_getSZ(4);
  @$pb.TagNumber(5)
  set currencyCode($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasCurrencyCode() => $_has(4);
  @$pb.TagNumber(5)
  void clearCurrencyCode() => clearField(5);

  @$pb.TagNumber(6)
  Address get address => $_getN(5);
  @$pb.TagNumber(6)
  set address(Address v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasAddress() => $_has(5);
  @$pb.TagNumber(6)
  void clearAddress() => clearField(6);
  @$pb.TagNumber(6)
  Address ensureAddress() => $_ensure(5);

  @$pb.TagNumber(7)
  $core.List<Contact> get contacts => $_getList(6);

  @$pb.TagNumber(8)
  $core.List<Event> get events => $_getList(7);

  @$pb.TagNumber(9)
  PlaceMeta get placeMeta => $_getN(8);
  @$pb.TagNumber(9)
  set placeMeta(PlaceMeta v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasPlaceMeta() => $_has(8);
  @$pb.TagNumber(9)
  void clearPlaceMeta() => clearField(9);
  @$pb.TagNumber(9)
  PlaceMeta ensurePlaceMeta() => $_ensure(8);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(10)
  JsonString get placeMap => $_getN(9);
  @$pb.TagNumber(10)
  set placeMap(JsonString v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasPlaceMap() => $_has(9);
  @$pb.TagNumber(10)
  void clearPlaceMap() => clearField(10);
  @$pb.TagNumber(10)
  JsonString ensurePlaceMap() => $_ensure(9);

  /// Timestamp of data creation
  @$pb.TagNumber(11)
  $core.String get timeStamp => $_getSZ(10);
  @$pb.TagNumber(11)
  set timeStamp($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasTimeStamp() => $_has(10);
  @$pb.TagNumber(11)
  void clearTimeStamp() => clearField(11);
}

class Places extends $pb.GeneratedMessage {
  factory Places({
    $core.Iterable<Place>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Places._() : super();
  factory Places.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Places.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Places', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Place>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Place.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Places clone() => Places()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Places copyWith(void Function(Places) updates) => super.copyWith((message) => updates(message as Places)) as Places;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Places create() => Places._();
  Places createEmptyInstance() => create();
  static $pb.PbList<Places> createRepeated() => $pb.PbList<Places>();
  @$core.pragma('dart2js:noInline')
  static Places getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Places>(create);
  static Places? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Place> get data => $_getList(0);
}

/// PriceMeta message
class PriceMeta extends $pb.GeneratedMessage {
  factory PriceMeta({
    $core.double? priceValue,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (priceValue != null) {
      $result.priceValue = priceValue;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  PriceMeta._() : super();
  factory PriceMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PriceMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PriceMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..a<$core.double>(1, _omitFieldNames ? '' : 'priceValue', $pb.PbFieldType.OD)
    ..pPS(2, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PriceMeta clone() => PriceMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PriceMeta copyWith(void Function(PriceMeta) updates) => super.copyWith((message) => updates(message as PriceMeta)) as PriceMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PriceMeta create() => PriceMeta._();
  PriceMeta createEmptyInstance() => create();
  static $pb.PbList<PriceMeta> createRepeated() => $pb.PbList<PriceMeta>();
  @$core.pragma('dart2js:noInline')
  static PriceMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PriceMeta>(create);
  static PriceMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.double get priceValue => $_getN(0);
  @$pb.TagNumber(1)
  set priceValue($core.double v) { $_setDouble(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasPriceValue() => $_has(0);
  @$pb.TagNumber(1)
  void clearPriceValue() => clearField(1);

  /// Additional tags for the price
  @$pb.TagNumber(2)
  $core.List<$core.String> get tags => $_getList(1);
}

/// Price message
class Price extends $pb.GeneratedMessage {
  factory Price({
    $fixnum.Int64? id,
    $core.String? code,
    PriceType? priceType,
    $core.String? validFrom,
    $core.String? validTo,
    $core.String? productCode,
    $core.String? currencyCode,
    $core.String? customerCode,
    $core.double? qty,
    PriceMeta? priceMeta,
    JsonString? priceMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (priceType != null) {
      $result.priceType = priceType;
    }
    if (validFrom != null) {
      $result.validFrom = validFrom;
    }
    if (validTo != null) {
      $result.validTo = validTo;
    }
    if (productCode != null) {
      $result.productCode = productCode;
    }
    if (currencyCode != null) {
      $result.currencyCode = currencyCode;
    }
    if (customerCode != null) {
      $result.customerCode = customerCode;
    }
    if (qty != null) {
      $result.qty = qty;
    }
    if (priceMeta != null) {
      $result.priceMeta = priceMeta;
    }
    if (priceMap != null) {
      $result.priceMap = priceMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Price._() : super();
  factory Price.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Price.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Price', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<PriceType>(3, _omitFieldNames ? '' : 'priceType', $pb.PbFieldType.OE, defaultOrMaker: PriceType.PRICE_CUSTOMER, valueOf: PriceType.valueOf, enumValues: PriceType.values)
    ..aOS(4, _omitFieldNames ? '' : 'validFrom')
    ..aOS(5, _omitFieldNames ? '' : 'validTo')
    ..aOS(6, _omitFieldNames ? '' : 'productCode')
    ..aOS(7, _omitFieldNames ? '' : 'currencyCode')
    ..aOS(8, _omitFieldNames ? '' : 'customerCode')
    ..a<$core.double>(9, _omitFieldNames ? '' : 'qty', $pb.PbFieldType.OD)
    ..aOM<PriceMeta>(10, _omitFieldNames ? '' : 'priceMeta', subBuilder: PriceMeta.create)
    ..aOM<JsonString>(11, _omitFieldNames ? '' : 'priceMap', subBuilder: JsonString.create)
    ..aOS(12, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Price clone() => Price()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Price copyWith(void Function(Price) updates) => super.copyWith((message) => updates(message as Price)) as Price;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Price create() => Price._();
  Price createEmptyInstance() => create();
  static $pb.PbList<Price> createRepeated() => $pb.PbList<Price>();
  @$core.pragma('dart2js:noInline')
  static Price getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Price>(create);
  static Price? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: PRC1731101982N123 ("PRC" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(3)
  PriceType get priceType => $_getN(2);
  @$pb.TagNumber(3)
  set priceType(PriceType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasPriceType() => $_has(2);
  @$pb.TagNumber(3)
  void clearPriceType() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Start of validity, mandatory data. ISO 8601 date
  @$pb.TagNumber(4)
  $core.String get validFrom => $_getSZ(3);
  @$pb.TagNumber(4)
  set validFrom($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasValidFrom() => $_has(3);
  @$pb.TagNumber(4)
  void clearValidFrom() => clearField(4);

  /// End of validity, can be left empty. ISO 8601 date
  @$pb.TagNumber(5)
  $core.String get validTo => $_getSZ(4);
  @$pb.TagNumber(5)
  set validTo($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasValidTo() => $_has(4);
  @$pb.TagNumber(5)
  void clearValidTo() => clearField(5);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Product">Product</a>.Code
  @$pb.TagNumber(6)
  $core.String get productCode => $_getSZ(5);
  @$pb.TagNumber(6)
  set productCode($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasProductCode() => $_has(5);
  @$pb.TagNumber(6)
  void clearProductCode() => clearField(6);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Currency">Currency</a>.Code
  @$pb.TagNumber(7)
  $core.String get currencyCode => $_getSZ(6);
  @$pb.TagNumber(7)
  set currencyCode($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasCurrencyCode() => $_has(6);
  @$pb.TagNumber(7)
  void clearCurrencyCode() => clearField(7);

  /// Optional field. Reference to <a href="#Customer">Customer</a>.Code
  @$pb.TagNumber(8)
  $core.String get customerCode => $_getSZ(7);
  @$pb.TagNumber(8)
  set customerCode($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasCustomerCode() => $_has(7);
  @$pb.TagNumber(8)
  void clearCustomerCode() => clearField(8);

  /// Price ranges can also be specified, thus different price can be set for a smaller and bigger quantity of the same product.
  /// The quantity should be used as the lower threshold, ie. this should be the minimum quantity for the price set.
  @$pb.TagNumber(9)
  $core.double get qty => $_getN(8);
  @$pb.TagNumber(9)
  set qty($core.double v) { $_setDouble(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasQty() => $_has(8);
  @$pb.TagNumber(9)
  void clearQty() => clearField(9);

  @$pb.TagNumber(10)
  PriceMeta get priceMeta => $_getN(9);
  @$pb.TagNumber(10)
  set priceMeta(PriceMeta v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasPriceMeta() => $_has(9);
  @$pb.TagNumber(10)
  void clearPriceMeta() => clearField(10);
  @$pb.TagNumber(10)
  PriceMeta ensurePriceMeta() => $_ensure(9);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(11)
  JsonString get priceMap => $_getN(10);
  @$pb.TagNumber(11)
  set priceMap(JsonString v) { setField(11, v); }
  @$pb.TagNumber(11)
  $core.bool hasPriceMap() => $_has(10);
  @$pb.TagNumber(11)
  void clearPriceMap() => clearField(11);
  @$pb.TagNumber(11)
  JsonString ensurePriceMap() => $_ensure(10);

  /// Timestamp of data creation
  @$pb.TagNumber(12)
  $core.String get timeStamp => $_getSZ(11);
  @$pb.TagNumber(12)
  set timeStamp($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasTimeStamp() => $_has(11);
  @$pb.TagNumber(12)
  void clearTimeStamp() => clearField(12);
}

class Prices extends $pb.GeneratedMessage {
  factory Prices({
    $core.Iterable<Price>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Prices._() : super();
  factory Prices.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Prices.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Prices', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Price>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Price.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Prices clone() => Prices()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Prices copyWith(void Function(Prices) updates) => super.copyWith((message) => updates(message as Prices)) as Prices;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Prices create() => Prices._();
  Prices createEmptyInstance() => create();
  static $pb.PbList<Prices> createRepeated() => $pb.PbList<Prices>();
  @$core.pragma('dart2js:noInline')
  static Prices getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Prices>(create);
  static Prices? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Price> get data => $_getList(0);
}

/// ProductMeta message
class ProductMeta extends $pb.GeneratedMessage {
  factory ProductMeta({
    $core.String? unit,
    BarcodeType? barcodeType,
    $core.String? barcodeData,
    $core.double? barcodeQty,
    $core.String? notes,
    $core.bool? inactive,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (unit != null) {
      $result.unit = unit;
    }
    if (barcodeType != null) {
      $result.barcodeType = barcodeType;
    }
    if (barcodeData != null) {
      $result.barcodeData = barcodeData;
    }
    if (barcodeQty != null) {
      $result.barcodeQty = barcodeQty;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (inactive != null) {
      $result.inactive = inactive;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  ProductMeta._() : super();
  factory ProductMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProductMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ProductMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'unit')
    ..e<BarcodeType>(2, _omitFieldNames ? '' : 'barcodeType', $pb.PbFieldType.OE, defaultOrMaker: BarcodeType.BARCODE_CODE_39, valueOf: BarcodeType.valueOf, enumValues: BarcodeType.values)
    ..aOS(3, _omitFieldNames ? '' : 'barcodeData')
    ..a<$core.double>(4, _omitFieldNames ? '' : 'barcodeQty', $pb.PbFieldType.OD)
    ..aOS(5, _omitFieldNames ? '' : 'notes')
    ..aOB(6, _omitFieldNames ? '' : 'inactive')
    ..pPS(7, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProductMeta clone() => ProductMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProductMeta copyWith(void Function(ProductMeta) updates) => super.copyWith((message) => updates(message as ProductMeta)) as ProductMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ProductMeta create() => ProductMeta._();
  ProductMeta createEmptyInstance() => create();
  static $pb.PbList<ProductMeta> createRepeated() => $pb.PbList<ProductMeta>();
  @$core.pragma('dart2js:noInline')
  static ProductMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProductMeta>(create);
  static ProductMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get unit => $_getSZ(0);
  @$pb.TagNumber(1)
  set unit($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUnit() => $_has(0);
  @$pb.TagNumber(1)
  void clearUnit() => clearField(1);

  /// ENUM field.
  @$pb.TagNumber(2)
  BarcodeType get barcodeType => $_getN(1);
  @$pb.TagNumber(2)
  set barcodeType(BarcodeType v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasBarcodeType() => $_has(1);
  @$pb.TagNumber(2)
  void clearBarcodeType() => clearField(2);

  /// Any barcode or QR code data
  @$pb.TagNumber(3)
  $core.String get barcodeData => $_getSZ(2);
  @$pb.TagNumber(3)
  set barcodeData($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasBarcodeData() => $_has(2);
  @$pb.TagNumber(3)
  void clearBarcodeData() => clearField(3);

  /// The actual amount of the products identified by the barcode. For example can be used for packaged goods, tray packaging.
  @$pb.TagNumber(4)
  $core.double get barcodeQty => $_getN(3);
  @$pb.TagNumber(4)
  set barcodeQty($core.double v) { $_setDouble(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasBarcodeQty() => $_has(3);
  @$pb.TagNumber(4)
  void clearBarcodeQty() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get notes => $_getSZ(4);
  @$pb.TagNumber(5)
  set notes($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasNotes() => $_has(4);
  @$pb.TagNumber(5)
  void clearNotes() => clearField(5);

  @$pb.TagNumber(6)
  $core.bool get inactive => $_getBF(5);
  @$pb.TagNumber(6)
  set inactive($core.bool v) { $_setBool(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasInactive() => $_has(5);
  @$pb.TagNumber(6)
  void clearInactive() => clearField(6);

  /// Additional tags for the product
  @$pb.TagNumber(7)
  $core.List<$core.String> get tags => $_getList(6);
}

/// Product message
class Product extends $pb.GeneratedMessage {
  factory Product({
    $fixnum.Int64? id,
    $core.String? code,
    ProductType? productType,
    $core.String? productName,
    $core.String? taxCode,
    $core.Iterable<Event>? events,
    ProductMeta? productMeta,
    JsonString? productMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (productType != null) {
      $result.productType = productType;
    }
    if (productName != null) {
      $result.productName = productName;
    }
    if (taxCode != null) {
      $result.taxCode = taxCode;
    }
    if (events != null) {
      $result.events.addAll(events);
    }
    if (productMeta != null) {
      $result.productMeta = productMeta;
    }
    if (productMap != null) {
      $result.productMap = productMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Product._() : super();
  factory Product.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Product.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Product', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<ProductType>(3, _omitFieldNames ? '' : 'productType', $pb.PbFieldType.OE, defaultOrMaker: ProductType.PRODUCT_ITEM, valueOf: ProductType.valueOf, enumValues: ProductType.values)
    ..aOS(4, _omitFieldNames ? '' : 'productName')
    ..aOS(5, _omitFieldNames ? '' : 'taxCode')
    ..pc<Event>(6, _omitFieldNames ? '' : 'events', $pb.PbFieldType.PM, subBuilder: Event.create)
    ..aOM<ProductMeta>(7, _omitFieldNames ? '' : 'productMeta', subBuilder: ProductMeta.create)
    ..aOM<JsonString>(8, _omitFieldNames ? '' : 'productMap', subBuilder: JsonString.create)
    ..aOS(9, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Product clone() => Product()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Product copyWith(void Function(Product) updates) => super.copyWith((message) => updates(message as Product)) as Product;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Product create() => Product._();
  Product createEmptyInstance() => create();
  static $pb.PbList<Product> createRepeated() => $pb.PbList<Product>();
  @$core.pragma('dart2js:noInline')
  static Product getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Product>(create);
  static Product? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: PRD1731101982N123 ("PRD" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(3)
  ProductType get productType => $_getN(2);
  @$pb.TagNumber(3)
  set productType(ProductType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasProductType() => $_has(2);
  @$pb.TagNumber(3)
  void clearProductType() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// The full name of the product or short description.
  @$pb.TagNumber(4)
  $core.String get productName => $_getSZ(3);
  @$pb.TagNumber(4)
  set productName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasProductName() => $_has(3);
  @$pb.TagNumber(4)
  void clearProductName() => clearField(4);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Tax">Tax</a>.Code
  @$pb.TagNumber(5)
  $core.String get taxCode => $_getSZ(4);
  @$pb.TagNumber(5)
  set taxCode($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTaxCode() => $_has(4);
  @$pb.TagNumber(5)
  void clearTaxCode() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<Event> get events => $_getList(5);

  @$pb.TagNumber(7)
  ProductMeta get productMeta => $_getN(6);
  @$pb.TagNumber(7)
  set productMeta(ProductMeta v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasProductMeta() => $_has(6);
  @$pb.TagNumber(7)
  void clearProductMeta() => clearField(7);
  @$pb.TagNumber(7)
  ProductMeta ensureProductMeta() => $_ensure(6);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(8)
  JsonString get productMap => $_getN(7);
  @$pb.TagNumber(8)
  set productMap(JsonString v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasProductMap() => $_has(7);
  @$pb.TagNumber(8)
  void clearProductMap() => clearField(8);
  @$pb.TagNumber(8)
  JsonString ensureProductMap() => $_ensure(7);

  /// Timestamp of data creation
  @$pb.TagNumber(9)
  $core.String get timeStamp => $_getSZ(8);
  @$pb.TagNumber(9)
  set timeStamp($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasTimeStamp() => $_has(8);
  @$pb.TagNumber(9)
  void clearTimeStamp() => clearField(9);
}

class Products extends $pb.GeneratedMessage {
  factory Products({
    $core.Iterable<Product>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Products._() : super();
  factory Products.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Products.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Products', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Product>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Product.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Products clone() => Products()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Products copyWith(void Function(Products) updates) => super.copyWith((message) => updates(message as Products)) as Products;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Products create() => Products._();
  Products createEmptyInstance() => create();
  static $pb.PbList<Products> createRepeated() => $pb.PbList<Products>();
  @$core.pragma('dart2js:noInline')
  static Products getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Products>(create);
  static Products? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Product> get data => $_getList(0);
}

/// ProjectMeta message
class ProjectMeta extends $pb.GeneratedMessage {
  factory ProjectMeta({
    $core.String? startDate,
    $core.String? endDate,
    $core.String? notes,
    $core.bool? inactive,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (startDate != null) {
      $result.startDate = startDate;
    }
    if (endDate != null) {
      $result.endDate = endDate;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (inactive != null) {
      $result.inactive = inactive;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  ProjectMeta._() : super();
  factory ProjectMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ProjectMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ProjectMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'startDate')
    ..aOS(2, _omitFieldNames ? '' : 'endDate')
    ..aOS(3, _omitFieldNames ? '' : 'notes')
    ..aOB(4, _omitFieldNames ? '' : 'inactive')
    ..pPS(5, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ProjectMeta clone() => ProjectMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ProjectMeta copyWith(void Function(ProjectMeta) updates) => super.copyWith((message) => updates(message as ProjectMeta)) as ProjectMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ProjectMeta create() => ProjectMeta._();
  ProjectMeta createEmptyInstance() => create();
  static $pb.PbList<ProjectMeta> createRepeated() => $pb.PbList<ProjectMeta>();
  @$core.pragma('dart2js:noInline')
  static ProjectMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ProjectMeta>(create);
  static ProjectMeta? _defaultInstance;

  /// ISO 8601 date
  @$pb.TagNumber(1)
  $core.String get startDate => $_getSZ(0);
  @$pb.TagNumber(1)
  set startDate($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasStartDate() => $_has(0);
  @$pb.TagNumber(1)
  void clearStartDate() => clearField(1);

  /// ISO 8601 date
  @$pb.TagNumber(2)
  $core.String get endDate => $_getSZ(1);
  @$pb.TagNumber(2)
  set endDate($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasEndDate() => $_has(1);
  @$pb.TagNumber(2)
  void clearEndDate() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get notes => $_getSZ(2);
  @$pb.TagNumber(3)
  set notes($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasNotes() => $_has(2);
  @$pb.TagNumber(3)
  void clearNotes() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get inactive => $_getBF(3);
  @$pb.TagNumber(4)
  set inactive($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasInactive() => $_has(3);
  @$pb.TagNumber(4)
  void clearInactive() => clearField(4);

  /// Additional tags for the project
  @$pb.TagNumber(5)
  $core.List<$core.String> get tags => $_getList(4);
}

/// Project message
class Project extends $pb.GeneratedMessage {
  factory Project({
    $fixnum.Int64? id,
    $core.String? code,
    $core.String? projectName,
    $core.String? customerCode,
    $core.Iterable<Address>? addresses,
    $core.Iterable<Contact>? contacts,
    $core.Iterable<Event>? events,
    ProjectMeta? projectMeta,
    JsonString? projectMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (projectName != null) {
      $result.projectName = projectName;
    }
    if (customerCode != null) {
      $result.customerCode = customerCode;
    }
    if (addresses != null) {
      $result.addresses.addAll(addresses);
    }
    if (contacts != null) {
      $result.contacts.addAll(contacts);
    }
    if (events != null) {
      $result.events.addAll(events);
    }
    if (projectMeta != null) {
      $result.projectMeta = projectMeta;
    }
    if (projectMap != null) {
      $result.projectMap = projectMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Project._() : super();
  factory Project.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Project.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Project', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..aOS(3, _omitFieldNames ? '' : 'projectName')
    ..aOS(4, _omitFieldNames ? '' : 'customerCode')
    ..pc<Address>(5, _omitFieldNames ? '' : 'addresses', $pb.PbFieldType.PM, subBuilder: Address.create)
    ..pc<Contact>(6, _omitFieldNames ? '' : 'contacts', $pb.PbFieldType.PM, subBuilder: Contact.create)
    ..pc<Event>(7, _omitFieldNames ? '' : 'events', $pb.PbFieldType.PM, subBuilder: Event.create)
    ..aOM<ProjectMeta>(8, _omitFieldNames ? '' : 'projectMeta', subBuilder: ProjectMeta.create)
    ..aOM<JsonString>(9, _omitFieldNames ? '' : 'projectMap', subBuilder: JsonString.create)
    ..aOS(10, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Project clone() => Project()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Project copyWith(void Function(Project) updates) => super.copyWith((message) => updates(message as Project)) as Project;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Project create() => Project._();
  Project createEmptyInstance() => create();
  static $pb.PbList<Project> createRepeated() => $pb.PbList<Project>();
  @$core.pragma('dart2js:noInline')
  static Project getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Project>(create);
  static Project? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: PRJ1731101982N123 ("PRJ" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// The name of the project.
  @$pb.TagNumber(3)
  $core.String get projectName => $_getSZ(2);
  @$pb.TagNumber(3)
  set projectName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasProjectName() => $_has(2);
  @$pb.TagNumber(3)
  void clearProjectName() => clearField(3);

  /// Reference to <a href="#Customer">Customer</a>.Code
  @$pb.TagNumber(4)
  $core.String get customerCode => $_getSZ(3);
  @$pb.TagNumber(4)
  set customerCode($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCustomerCode() => $_has(3);
  @$pb.TagNumber(4)
  void clearCustomerCode() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<Address> get addresses => $_getList(4);

  @$pb.TagNumber(6)
  $core.List<Contact> get contacts => $_getList(5);

  @$pb.TagNumber(7)
  $core.List<Event> get events => $_getList(6);

  @$pb.TagNumber(8)
  ProjectMeta get projectMeta => $_getN(7);
  @$pb.TagNumber(8)
  set projectMeta(ProjectMeta v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasProjectMeta() => $_has(7);
  @$pb.TagNumber(8)
  void clearProjectMeta() => clearField(8);
  @$pb.TagNumber(8)
  ProjectMeta ensureProjectMeta() => $_ensure(7);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(9)
  JsonString get projectMap => $_getN(8);
  @$pb.TagNumber(9)
  set projectMap(JsonString v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasProjectMap() => $_has(8);
  @$pb.TagNumber(9)
  void clearProjectMap() => clearField(9);
  @$pb.TagNumber(9)
  JsonString ensureProjectMap() => $_ensure(8);

  /// Timestamp of data creation
  @$pb.TagNumber(10)
  $core.String get timeStamp => $_getSZ(9);
  @$pb.TagNumber(10)
  set timeStamp($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasTimeStamp() => $_has(9);
  @$pb.TagNumber(10)
  void clearTimeStamp() => clearField(10);
}

class Projects extends $pb.GeneratedMessage {
  factory Projects({
    $core.Iterable<Project>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Projects._() : super();
  factory Projects.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Projects.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Projects', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Project>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Project.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Projects clone() => Projects()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Projects copyWith(void Function(Projects) updates) => super.copyWith((message) => updates(message as Projects)) as Projects;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Projects create() => Projects._();
  Projects createEmptyInstance() => create();
  static $pb.PbList<Projects> createRepeated() => $pb.PbList<Projects>();
  @$core.pragma('dart2js:noInline')
  static Projects getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Projects>(create);
  static Projects? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Project> get data => $_getList(0);
}

/// RateMeta message
class RateMeta extends $pb.GeneratedMessage {
  factory RateMeta({
    $core.double? rateValue,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (rateValue != null) {
      $result.rateValue = rateValue;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  RateMeta._() : super();
  factory RateMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RateMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RateMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..a<$core.double>(1, _omitFieldNames ? '' : 'rateValue', $pb.PbFieldType.OD)
    ..pPS(2, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RateMeta clone() => RateMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RateMeta copyWith(void Function(RateMeta) updates) => super.copyWith((message) => updates(message as RateMeta)) as RateMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RateMeta create() => RateMeta._();
  RateMeta createEmptyInstance() => create();
  static $pb.PbList<RateMeta> createRepeated() => $pb.PbList<RateMeta>();
  @$core.pragma('dart2js:noInline')
  static RateMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RateMeta>(create);
  static RateMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.double get rateValue => $_getN(0);
  @$pb.TagNumber(1)
  set rateValue($core.double v) { $_setDouble(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRateValue() => $_has(0);
  @$pb.TagNumber(1)
  void clearRateValue() => clearField(1);

  /// Additional tags for the rate
  @$pb.TagNumber(2)
  $core.List<$core.String> get tags => $_getList(1);
}

/// Rate message
class Rate extends $pb.GeneratedMessage {
  factory Rate({
    $fixnum.Int64? id,
    $core.String? code,
    RateType? rateType,
    $core.String? rateDate,
    $core.String? placeCode,
    $core.String? currencyCode,
    RateMeta? rateMeta,
    JsonString? rateMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (rateType != null) {
      $result.rateType = rateType;
    }
    if (rateDate != null) {
      $result.rateDate = rateDate;
    }
    if (placeCode != null) {
      $result.placeCode = placeCode;
    }
    if (currencyCode != null) {
      $result.currencyCode = currencyCode;
    }
    if (rateMeta != null) {
      $result.rateMeta = rateMeta;
    }
    if (rateMap != null) {
      $result.rateMap = rateMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Rate._() : super();
  factory Rate.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Rate.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Rate', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<RateType>(3, _omitFieldNames ? '' : 'rateType', $pb.PbFieldType.OE, defaultOrMaker: RateType.RATE_RATE, valueOf: RateType.valueOf, enumValues: RateType.values)
    ..aOS(4, _omitFieldNames ? '' : 'rateDate')
    ..aOS(5, _omitFieldNames ? '' : 'placeCode')
    ..aOS(6, _omitFieldNames ? '' : 'currencyCode')
    ..aOM<RateMeta>(7, _omitFieldNames ? '' : 'rateMeta', subBuilder: RateMeta.create)
    ..aOM<JsonString>(8, _omitFieldNames ? '' : 'rateMap', subBuilder: JsonString.create)
    ..aOS(9, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Rate clone() => Rate()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Rate copyWith(void Function(Rate) updates) => super.copyWith((message) => updates(message as Rate)) as Rate;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Rate create() => Rate._();
  Rate createEmptyInstance() => create();
  static $pb.PbList<Rate> createRepeated() => $pb.PbList<Rate>();
  @$core.pragma('dart2js:noInline')
  static Rate getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Rate>(create);
  static Rate? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: RAT1731101982N123 ("RAT" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(3)
  RateType get rateType => $_getN(2);
  @$pb.TagNumber(3)
  set rateType(RateType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasRateType() => $_has(2);
  @$pb.TagNumber(3)
  void clearRateType() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ISO 8601 date
  @$pb.TagNumber(4)
  $core.String get rateDate => $_getSZ(3);
  @$pb.TagNumber(4)
  set rateDate($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRateDate() => $_has(3);
  @$pb.TagNumber(4)
  void clearRateDate() => clearField(4);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Place">Place</a>.Code
  @$pb.TagNumber(5)
  $core.String get placeCode => $_getSZ(4);
  @$pb.TagNumber(5)
  set placeCode($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasPlaceCode() => $_has(4);
  @$pb.TagNumber(5)
  void clearPlaceCode() => clearField(5);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Currency">Currency</a>.Code
  @$pb.TagNumber(6)
  $core.String get currencyCode => $_getSZ(5);
  @$pb.TagNumber(6)
  set currencyCode($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasCurrencyCode() => $_has(5);
  @$pb.TagNumber(6)
  void clearCurrencyCode() => clearField(6);

  @$pb.TagNumber(7)
  RateMeta get rateMeta => $_getN(6);
  @$pb.TagNumber(7)
  set rateMeta(RateMeta v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasRateMeta() => $_has(6);
  @$pb.TagNumber(7)
  void clearRateMeta() => clearField(7);
  @$pb.TagNumber(7)
  RateMeta ensureRateMeta() => $_ensure(6);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(8)
  JsonString get rateMap => $_getN(7);
  @$pb.TagNumber(8)
  set rateMap(JsonString v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasRateMap() => $_has(7);
  @$pb.TagNumber(8)
  void clearRateMap() => clearField(8);
  @$pb.TagNumber(8)
  JsonString ensureRateMap() => $_ensure(7);

  /// Timestamp of data creation
  @$pb.TagNumber(9)
  $core.String get timeStamp => $_getSZ(8);
  @$pb.TagNumber(9)
  set timeStamp($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasTimeStamp() => $_has(8);
  @$pb.TagNumber(9)
  void clearTimeStamp() => clearField(9);
}

class Rates extends $pb.GeneratedMessage {
  factory Rates({
    $core.Iterable<Rate>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Rates._() : super();
  factory Rates.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Rates.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Rates', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Rate>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Rate.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Rates clone() => Rates()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Rates copyWith(void Function(Rates) updates) => super.copyWith((message) => updates(message as Rates)) as Rates;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Rates create() => Rates._();
  Rates createEmptyInstance() => create();
  static $pb.PbList<Rates> createRepeated() => $pb.PbList<Rates>();
  @$core.pragma('dart2js:noInline')
  static Rates getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Rates>(create);
  static Rates? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Rate> get data => $_getList(0);
}

/// TaxMeta message
class TaxMeta extends $pb.GeneratedMessage {
  factory TaxMeta({
    $core.String? description,
    $core.double? rateValue,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (description != null) {
      $result.description = description;
    }
    if (rateValue != null) {
      $result.rateValue = rateValue;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  TaxMeta._() : super();
  factory TaxMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TaxMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'TaxMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'description')
    ..a<$core.double>(2, _omitFieldNames ? '' : 'rateValue', $pb.PbFieldType.OD)
    ..pPS(3, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TaxMeta clone() => TaxMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TaxMeta copyWith(void Function(TaxMeta) updates) => super.copyWith((message) => updates(message as TaxMeta)) as TaxMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static TaxMeta create() => TaxMeta._();
  TaxMeta createEmptyInstance() => create();
  static $pb.PbList<TaxMeta> createRepeated() => $pb.PbList<TaxMeta>();
  @$core.pragma('dart2js:noInline')
  static TaxMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TaxMeta>(create);
  static TaxMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get description => $_getSZ(0);
  @$pb.TagNumber(1)
  set description($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDescription() => $_has(0);
  @$pb.TagNumber(1)
  void clearDescription() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get rateValue => $_getN(1);
  @$pb.TagNumber(2)
  set rateValue($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRateValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearRateValue() => clearField(2);

  /// Additional tags for the tax
  @$pb.TagNumber(3)
  $core.List<$core.String> get tags => $_getList(2);
}

/// Tax message
class Tax extends $pb.GeneratedMessage {
  factory Tax({
    $fixnum.Int64? id,
    $core.String? code,
    TaxMeta? taxMeta,
    JsonString? taxMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (taxMeta != null) {
      $result.taxMeta = taxMeta;
    }
    if (taxMap != null) {
      $result.taxMap = taxMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Tax._() : super();
  factory Tax.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Tax.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Tax', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..aOM<TaxMeta>(3, _omitFieldNames ? '' : 'taxMeta', subBuilder: TaxMeta.create)
    ..aOM<JsonString>(4, _omitFieldNames ? '' : 'taxMap', subBuilder: JsonString.create)
    ..aOS(5, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Tax clone() => Tax()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Tax copyWith(void Function(Tax) updates) => super.copyWith((message) => updates(message as Tax)) as Tax;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Tax create() => Tax._();
  Tax createEmptyInstance() => create();
  static $pb.PbList<Tax> createRepeated() => $pb.PbList<Tax>();
  @$core.pragma('dart2js:noInline')
  static Tax getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Tax>(create);
  static Tax? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Database independent unique external key. It cannot be modified after creation.
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  @$pb.TagNumber(3)
  TaxMeta get taxMeta => $_getN(2);
  @$pb.TagNumber(3)
  set taxMeta(TaxMeta v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasTaxMeta() => $_has(2);
  @$pb.TagNumber(3)
  void clearTaxMeta() => clearField(3);
  @$pb.TagNumber(3)
  TaxMeta ensureTaxMeta() => $_ensure(2);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(4)
  JsonString get taxMap => $_getN(3);
  @$pb.TagNumber(4)
  set taxMap(JsonString v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasTaxMap() => $_has(3);
  @$pb.TagNumber(4)
  void clearTaxMap() => clearField(4);
  @$pb.TagNumber(4)
  JsonString ensureTaxMap() => $_ensure(3);

  /// Timestamp of data creation
  @$pb.TagNumber(5)
  $core.String get timeStamp => $_getSZ(4);
  @$pb.TagNumber(5)
  set timeStamp($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTimeStamp() => $_has(4);
  @$pb.TagNumber(5)
  void clearTimeStamp() => clearField(5);
}

class Taxes extends $pb.GeneratedMessage {
  factory Taxes({
    $core.Iterable<Tax>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Taxes._() : super();
  factory Taxes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Taxes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Taxes', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Tax>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Tax.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Taxes clone() => Taxes()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Taxes copyWith(void Function(Taxes) updates) => super.copyWith((message) => updates(message as Taxes)) as Taxes;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Taxes create() => Taxes._();
  Taxes createEmptyInstance() => create();
  static $pb.PbList<Taxes> createRepeated() => $pb.PbList<Taxes>();
  @$core.pragma('dart2js:noInline')
  static Taxes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Taxes>(create);
  static Taxes? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Tax> get data => $_getList(0);
}

/// ToolMeta message
class ToolMeta extends $pb.GeneratedMessage {
  factory ToolMeta({
    $core.String? serialNumber,
    $core.String? notes,
    $core.bool? inactive,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (serialNumber != null) {
      $result.serialNumber = serialNumber;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (inactive != null) {
      $result.inactive = inactive;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  ToolMeta._() : super();
  factory ToolMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ToolMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ToolMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'serialNumber')
    ..aOS(2, _omitFieldNames ? '' : 'notes')
    ..aOB(3, _omitFieldNames ? '' : 'inactive')
    ..pPS(4, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ToolMeta clone() => ToolMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ToolMeta copyWith(void Function(ToolMeta) updates) => super.copyWith((message) => updates(message as ToolMeta)) as ToolMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ToolMeta create() => ToolMeta._();
  ToolMeta createEmptyInstance() => create();
  static $pb.PbList<ToolMeta> createRepeated() => $pb.PbList<ToolMeta>();
  @$core.pragma('dart2js:noInline')
  static ToolMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ToolMeta>(create);
  static ToolMeta? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get serialNumber => $_getSZ(0);
  @$pb.TagNumber(1)
  set serialNumber($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSerialNumber() => $_has(0);
  @$pb.TagNumber(1)
  void clearSerialNumber() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get notes => $_getSZ(1);
  @$pb.TagNumber(2)
  set notes($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNotes() => $_has(1);
  @$pb.TagNumber(2)
  void clearNotes() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get inactive => $_getBF(2);
  @$pb.TagNumber(3)
  set inactive($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasInactive() => $_has(2);
  @$pb.TagNumber(3)
  void clearInactive() => clearField(3);

  /// Additional tags for the tool
  @$pb.TagNumber(4)
  $core.List<$core.String> get tags => $_getList(3);
}

/// Tool message
class Tool extends $pb.GeneratedMessage {
  factory Tool({
    $fixnum.Int64? id,
    $core.String? code,
    $core.String? description,
    $core.String? productCode,
    $core.Iterable<Event>? events,
    ToolMeta? toolMeta,
    JsonString? toolMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (description != null) {
      $result.description = description;
    }
    if (productCode != null) {
      $result.productCode = productCode;
    }
    if (events != null) {
      $result.events.addAll(events);
    }
    if (toolMeta != null) {
      $result.toolMeta = toolMeta;
    }
    if (toolMap != null) {
      $result.toolMap = toolMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Tool._() : super();
  factory Tool.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Tool.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Tool', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..aOS(3, _omitFieldNames ? '' : 'description')
    ..aOS(4, _omitFieldNames ? '' : 'productCode')
    ..pc<Event>(5, _omitFieldNames ? '' : 'events', $pb.PbFieldType.PM, subBuilder: Event.create)
    ..aOM<ToolMeta>(6, _omitFieldNames ? '' : 'toolMeta', subBuilder: ToolMeta.create)
    ..aOM<JsonString>(7, _omitFieldNames ? '' : 'toolMap', subBuilder: JsonString.create)
    ..aOS(8, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Tool clone() => Tool()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Tool copyWith(void Function(Tool) updates) => super.copyWith((message) => updates(message as Tool)) as Tool;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Tool create() => Tool._();
  Tool createEmptyInstance() => create();
  static $pb.PbList<Tool> createRepeated() => $pb.PbList<Tool>();
  @$core.pragma('dart2js:noInline')
  static Tool getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Tool>(create);
  static Tool? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: SER1731101982N123 ("SER" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// The description of the tool.
  @$pb.TagNumber(3)
  $core.String get description => $_getSZ(2);
  @$pb.TagNumber(3)
  set description($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDescription() => $_has(2);
  @$pb.TagNumber(3)
  void clearDescription() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// Reference to <a href="#Product">Product</a>.Code
  @$pb.TagNumber(4)
  $core.String get productCode => $_getSZ(3);
  @$pb.TagNumber(4)
  set productCode($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasProductCode() => $_has(3);
  @$pb.TagNumber(4)
  void clearProductCode() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<Event> get events => $_getList(4);

  @$pb.TagNumber(6)
  ToolMeta get toolMeta => $_getN(5);
  @$pb.TagNumber(6)
  set toolMeta(ToolMeta v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasToolMeta() => $_has(5);
  @$pb.TagNumber(6)
  void clearToolMeta() => clearField(6);
  @$pb.TagNumber(6)
  ToolMeta ensureToolMeta() => $_ensure(5);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(7)
  JsonString get toolMap => $_getN(6);
  @$pb.TagNumber(7)
  set toolMap(JsonString v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasToolMap() => $_has(6);
  @$pb.TagNumber(7)
  void clearToolMap() => clearField(7);
  @$pb.TagNumber(7)
  JsonString ensureToolMap() => $_ensure(6);

  /// Timestamp of data creation
  @$pb.TagNumber(8)
  $core.String get timeStamp => $_getSZ(7);
  @$pb.TagNumber(8)
  set timeStamp($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasTimeStamp() => $_has(7);
  @$pb.TagNumber(8)
  void clearTimeStamp() => clearField(8);
}

class Tools extends $pb.GeneratedMessage {
  factory Tools({
    $core.Iterable<Tool>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Tools._() : super();
  factory Tools.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Tools.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Tools', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Tool>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Tool.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Tools clone() => Tools()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Tools copyWith(void Function(Tools) updates) => super.copyWith((message) => updates(message as Tools)) as Tools;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Tools create() => Tools._();
  Tools createEmptyInstance() => create();
  static $pb.PbList<Tools> createRepeated() => $pb.PbList<Tools>();
  @$core.pragma('dart2js:noInline')
  static Tools getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Tools>(create);
  static Tools? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Tool> get data => $_getList(0);
}

/// Trans Worksheet message
class Worksheet extends $pb.GeneratedMessage {
  factory Worksheet({
    $core.double? distance,
    $core.double? repair,
    $core.double? total,
    $core.String? justification,
  }) {
    final $result = create();
    if (distance != null) {
      $result.distance = distance;
    }
    if (repair != null) {
      $result.repair = repair;
    }
    if (total != null) {
      $result.total = total;
    }
    if (justification != null) {
      $result.justification = justification;
    }
    return $result;
  }
  Worksheet._() : super();
  factory Worksheet.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Worksheet.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Worksheet', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..a<$core.double>(1, _omitFieldNames ? '' : 'distance', $pb.PbFieldType.OD)
    ..a<$core.double>(2, _omitFieldNames ? '' : 'repair', $pb.PbFieldType.OD)
    ..a<$core.double>(3, _omitFieldNames ? '' : 'total', $pb.PbFieldType.OD)
    ..aOS(4, _omitFieldNames ? '' : 'justification')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Worksheet clone() => Worksheet()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Worksheet copyWith(void Function(Worksheet) updates) => super.copyWith((message) => updates(message as Worksheet)) as Worksheet;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Worksheet create() => Worksheet._();
  Worksheet createEmptyInstance() => create();
  static $pb.PbList<Worksheet> createRepeated() => $pb.PbList<Worksheet>();
  @$core.pragma('dart2js:noInline')
  static Worksheet getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Worksheet>(create);
  static Worksheet? _defaultInstance;

  /// Distance (km)
  @$pb.TagNumber(1)
  $core.double get distance => $_getN(0);
  @$pb.TagNumber(1)
  set distance($core.double v) { $_setDouble(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDistance() => $_has(0);
  @$pb.TagNumber(1)
  void clearDistance() => clearField(1);

  /// Repair time (hour)
  @$pb.TagNumber(2)
  $core.double get repair => $_getN(1);
  @$pb.TagNumber(2)
  set repair($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRepair() => $_has(1);
  @$pb.TagNumber(2)
  void clearRepair() => clearField(2);

  /// Total time (hour)
  @$pb.TagNumber(3)
  $core.double get total => $_getN(2);
  @$pb.TagNumber(3)
  set total($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTotal() => $_has(2);
  @$pb.TagNumber(3)
  void clearTotal() => clearField(3);

  /// Worksheet justification
  @$pb.TagNumber(4)
  $core.String get justification => $_getSZ(3);
  @$pb.TagNumber(4)
  set justification($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasJustification() => $_has(3);
  @$pb.TagNumber(4)
  void clearJustification() => clearField(4);
}

/// Trans Rent message
class Rent extends $pb.GeneratedMessage {
  factory Rent({
    $core.double? holiday,
    $core.double? badTool,
    $core.double? other,
    $core.String? justification,
  }) {
    final $result = create();
    if (holiday != null) {
      $result.holiday = holiday;
    }
    if (badTool != null) {
      $result.badTool = badTool;
    }
    if (other != null) {
      $result.other = other;
    }
    if (justification != null) {
      $result.justification = justification;
    }
    return $result;
  }
  Rent._() : super();
  factory Rent.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Rent.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Rent', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..a<$core.double>(1, _omitFieldNames ? '' : 'holiday', $pb.PbFieldType.OD)
    ..a<$core.double>(2, _omitFieldNames ? '' : 'badTool', $pb.PbFieldType.OD)
    ..a<$core.double>(3, _omitFieldNames ? '' : 'other', $pb.PbFieldType.OD)
    ..aOS(4, _omitFieldNames ? '' : 'justification')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Rent clone() => Rent()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Rent copyWith(void Function(Rent) updates) => super.copyWith((message) => updates(message as Rent)) as Rent;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Rent create() => Rent._();
  Rent createEmptyInstance() => create();
  static $pb.PbList<Rent> createRepeated() => $pb.PbList<Rent>();
  @$core.pragma('dart2js:noInline')
  static Rent getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Rent>(create);
  static Rent? _defaultInstance;

  /// Holidays (day)
  @$pb.TagNumber(1)
  $core.double get holiday => $_getN(0);
  @$pb.TagNumber(1)
  set holiday($core.double v) { $_setDouble(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasHoliday() => $_has(0);
  @$pb.TagNumber(1)
  void clearHoliday() => clearField(1);

  /// Bad tool / machine (hour)
  @$pb.TagNumber(2)
  $core.double get badTool => $_getN(1);
  @$pb.TagNumber(2)
  set badTool($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasBadTool() => $_has(1);
  @$pb.TagNumber(2)
  void clearBadTool() => clearField(2);

  /// Other non-eligible
  @$pb.TagNumber(3)
  $core.double get other => $_getN(2);
  @$pb.TagNumber(3)
  set other($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasOther() => $_has(2);
  @$pb.TagNumber(3)
  void clearOther() => clearField(3);

  /// Rent justification
  @$pb.TagNumber(4)
  $core.String get justification => $_getSZ(3);
  @$pb.TagNumber(4)
  set justification($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasJustification() => $_has(3);
  @$pb.TagNumber(4)
  void clearJustification() => clearField(4);
}

/// TransInvoice message
class Invoice extends $pb.GeneratedMessage {
  factory Invoice({
    $core.String? companyName,
    $core.String? companyAddress,
    $core.String? companyTaxNumber,
    $core.String? customerName,
    $core.String? customerAddress,
    $core.String? customerTaxNumber,
  }) {
    final $result = create();
    if (companyName != null) {
      $result.companyName = companyName;
    }
    if (companyAddress != null) {
      $result.companyAddress = companyAddress;
    }
    if (companyTaxNumber != null) {
      $result.companyTaxNumber = companyTaxNumber;
    }
    if (customerName != null) {
      $result.customerName = customerName;
    }
    if (customerAddress != null) {
      $result.customerAddress = customerAddress;
    }
    if (customerTaxNumber != null) {
      $result.customerTaxNumber = customerTaxNumber;
    }
    return $result;
  }
  Invoice._() : super();
  factory Invoice.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Invoice.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Invoice', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'companyName')
    ..aOS(2, _omitFieldNames ? '' : 'companyAddress')
    ..aOS(3, _omitFieldNames ? '' : 'companyTaxNumber')
    ..aOS(4, _omitFieldNames ? '' : 'customerName')
    ..aOS(5, _omitFieldNames ? '' : 'customerAddress')
    ..aOS(6, _omitFieldNames ? '' : 'customerTaxNumber')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Invoice clone() => Invoice()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Invoice copyWith(void Function(Invoice) updates) => super.copyWith((message) => updates(message as Invoice)) as Invoice;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Invoice create() => Invoice._();
  Invoice createEmptyInstance() => create();
  static $pb.PbList<Invoice> createRepeated() => $pb.PbList<Invoice>();
  @$core.pragma('dart2js:noInline')
  static Invoice getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Invoice>(create);
  static Invoice? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get companyName => $_getSZ(0);
  @$pb.TagNumber(1)
  set companyName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCompanyName() => $_has(0);
  @$pb.TagNumber(1)
  void clearCompanyName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get companyAddress => $_getSZ(1);
  @$pb.TagNumber(2)
  set companyAddress($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCompanyAddress() => $_has(1);
  @$pb.TagNumber(2)
  void clearCompanyAddress() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get companyTaxNumber => $_getSZ(2);
  @$pb.TagNumber(3)
  set companyTaxNumber($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasCompanyTaxNumber() => $_has(2);
  @$pb.TagNumber(3)
  void clearCompanyTaxNumber() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get customerName => $_getSZ(3);
  @$pb.TagNumber(4)
  set customerName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCustomerName() => $_has(3);
  @$pb.TagNumber(4)
  void clearCustomerName() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get customerAddress => $_getSZ(4);
  @$pb.TagNumber(5)
  set customerAddress($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasCustomerAddress() => $_has(4);
  @$pb.TagNumber(5)
  void clearCustomerAddress() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get customerTaxNumber => $_getSZ(5);
  @$pb.TagNumber(6)
  set customerTaxNumber($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasCustomerTaxNumber() => $_has(5);
  @$pb.TagNumber(6)
  void clearCustomerTaxNumber() => clearField(6);
}

/// TransMeta message
class TransMeta extends $pb.GeneratedMessage {
  factory TransMeta({
    $core.String? dueTime,
    $core.String? refNumber,
    PaidType? paidType,
    $core.bool? taxFree,
    $core.bool? paid,
    $core.double? rate,
    $core.bool? closed,
    TransStatus? status,
    TransState? transState,
    $core.String? notes,
    $core.String? internalNotes,
    $core.String? reportNotes,
    Worksheet? worksheet,
    Rent? rent,
    Invoice? invoice,
    $core.Iterable<$core.String>? tags,
  }) {
    final $result = create();
    if (dueTime != null) {
      $result.dueTime = dueTime;
    }
    if (refNumber != null) {
      $result.refNumber = refNumber;
    }
    if (paidType != null) {
      $result.paidType = paidType;
    }
    if (taxFree != null) {
      $result.taxFree = taxFree;
    }
    if (paid != null) {
      $result.paid = paid;
    }
    if (rate != null) {
      $result.rate = rate;
    }
    if (closed != null) {
      $result.closed = closed;
    }
    if (status != null) {
      $result.status = status;
    }
    if (transState != null) {
      $result.transState = transState;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (internalNotes != null) {
      $result.internalNotes = internalNotes;
    }
    if (reportNotes != null) {
      $result.reportNotes = reportNotes;
    }
    if (worksheet != null) {
      $result.worksheet = worksheet;
    }
    if (rent != null) {
      $result.rent = rent;
    }
    if (invoice != null) {
      $result.invoice = invoice;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    return $result;
  }
  TransMeta._() : super();
  factory TransMeta.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TransMeta.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'TransMeta', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'dueTime')
    ..aOS(2, _omitFieldNames ? '' : 'refNumber')
    ..e<PaidType>(3, _omitFieldNames ? '' : 'paidType', $pb.PbFieldType.OE, defaultOrMaker: PaidType.PAID_ONLINE, valueOf: PaidType.valueOf, enumValues: PaidType.values)
    ..aOB(4, _omitFieldNames ? '' : 'taxFree')
    ..aOB(5, _omitFieldNames ? '' : 'paid')
    ..a<$core.double>(6, _omitFieldNames ? '' : 'rate', $pb.PbFieldType.OD)
    ..aOB(7, _omitFieldNames ? '' : 'closed')
    ..e<TransStatus>(8, _omitFieldNames ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: TransStatus.TRANS_NORMAL, valueOf: TransStatus.valueOf, enumValues: TransStatus.values)
    ..e<TransState>(9, _omitFieldNames ? '' : 'transState', $pb.PbFieldType.OE, defaultOrMaker: TransState.TRANS_OK, valueOf: TransState.valueOf, enumValues: TransState.values)
    ..aOS(10, _omitFieldNames ? '' : 'notes')
    ..aOS(11, _omitFieldNames ? '' : 'internalNotes')
    ..aOS(12, _omitFieldNames ? '' : 'reportNotes')
    ..aOM<Worksheet>(13, _omitFieldNames ? '' : 'worksheet', subBuilder: Worksheet.create)
    ..aOM<Rent>(14, _omitFieldNames ? '' : 'rent', subBuilder: Rent.create)
    ..aOM<Invoice>(15, _omitFieldNames ? '' : 'invoice', subBuilder: Invoice.create)
    ..pPS(16, _omitFieldNames ? '' : 'tags')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TransMeta clone() => TransMeta()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TransMeta copyWith(void Function(TransMeta) updates) => super.copyWith((message) => updates(message as TransMeta)) as TransMeta;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static TransMeta create() => TransMeta._();
  TransMeta createEmptyInstance() => create();
  static $pb.PbList<TransMeta> createRepeated() => $pb.PbList<TransMeta>();
  @$core.pragma('dart2js:noInline')
  static TransMeta getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TransMeta>(create);
  static TransMeta? _defaultInstance;

  /// ISO 8601 date-time
  @$pb.TagNumber(1)
  $core.String get dueTime => $_getSZ(0);
  @$pb.TagNumber(1)
  set dueTime($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDueTime() => $_has(0);
  @$pb.TagNumber(1)
  void clearDueTime() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get refNumber => $_getSZ(1);
  @$pb.TagNumber(2)
  set refNumber($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRefNumber() => $_has(1);
  @$pb.TagNumber(2)
  void clearRefNumber() => clearField(2);

  /// ENUM field.
  @$pb.TagNumber(3)
  PaidType get paidType => $_getN(2);
  @$pb.TagNumber(3)
  set paidType(PaidType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasPaidType() => $_has(2);
  @$pb.TagNumber(3)
  void clearPaidType() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get taxFree => $_getBF(3);
  @$pb.TagNumber(4)
  set taxFree($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTaxFree() => $_has(3);
  @$pb.TagNumber(4)
  void clearTaxFree() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get paid => $_getBF(4);
  @$pb.TagNumber(5)
  set paid($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasPaid() => $_has(4);
  @$pb.TagNumber(5)
  void clearPaid() => clearField(5);

  @$pb.TagNumber(6)
  $core.double get rate => $_getN(5);
  @$pb.TagNumber(6)
  set rate($core.double v) { $_setDouble(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasRate() => $_has(5);
  @$pb.TagNumber(6)
  void clearRate() => clearField(6);

  @$pb.TagNumber(7)
  $core.bool get closed => $_getBF(6);
  @$pb.TagNumber(7)
  set closed($core.bool v) { $_setBool(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasClosed() => $_has(6);
  @$pb.TagNumber(7)
  void clearClosed() => clearField(7);

  /// ENUM field.
  @$pb.TagNumber(8)
  TransStatus get status => $_getN(7);
  @$pb.TagNumber(8)
  set status(TransStatus v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasStatus() => $_has(7);
  @$pb.TagNumber(8)
  void clearStatus() => clearField(8);

  /// ENUM field.
  @$pb.TagNumber(9)
  TransState get transState => $_getN(8);
  @$pb.TagNumber(9)
  set transState(TransState v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasTransState() => $_has(8);
  @$pb.TagNumber(9)
  void clearTransState() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get notes => $_getSZ(9);
  @$pb.TagNumber(10)
  set notes($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasNotes() => $_has(9);
  @$pb.TagNumber(10)
  void clearNotes() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get internalNotes => $_getSZ(10);
  @$pb.TagNumber(11)
  set internalNotes($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasInternalNotes() => $_has(10);
  @$pb.TagNumber(11)
  void clearInternalNotes() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get reportNotes => $_getSZ(11);
  @$pb.TagNumber(12)
  set reportNotes($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasReportNotes() => $_has(11);
  @$pb.TagNumber(12)
  void clearReportNotes() => clearField(12);

  @$pb.TagNumber(13)
  Worksheet get worksheet => $_getN(12);
  @$pb.TagNumber(13)
  set worksheet(Worksheet v) { setField(13, v); }
  @$pb.TagNumber(13)
  $core.bool hasWorksheet() => $_has(12);
  @$pb.TagNumber(13)
  void clearWorksheet() => clearField(13);
  @$pb.TagNumber(13)
  Worksheet ensureWorksheet() => $_ensure(12);

  @$pb.TagNumber(14)
  Rent get rent => $_getN(13);
  @$pb.TagNumber(14)
  set rent(Rent v) { setField(14, v); }
  @$pb.TagNumber(14)
  $core.bool hasRent() => $_has(13);
  @$pb.TagNumber(14)
  void clearRent() => clearField(14);
  @$pb.TagNumber(14)
  Rent ensureRent() => $_ensure(13);

  @$pb.TagNumber(15)
  Invoice get invoice => $_getN(14);
  @$pb.TagNumber(15)
  set invoice(Invoice v) { setField(15, v); }
  @$pb.TagNumber(15)
  $core.bool hasInvoice() => $_has(14);
  @$pb.TagNumber(15)
  void clearInvoice() => clearField(15);
  @$pb.TagNumber(15)
  Invoice ensureInvoice() => $_ensure(14);

  /// Additional tags for the trans
  @$pb.TagNumber(16)
  $core.List<$core.String> get tags => $_getList(15);
}

/// Trans message
class Trans extends $pb.GeneratedMessage {
  factory Trans({
    $fixnum.Int64? id,
    $core.String? code,
    TransType? transType,
    $core.String? transDate,
    Direction? direction,
    $core.String? transCode,
    $core.String? customerCode,
    $core.String? employeeCode,
    $core.String? projectCode,
    $core.String? placeCode,
    $core.String? currencyCode,
    $core.String? authCode,
    TransMeta? transMeta,
    JsonString? transMap,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (transType != null) {
      $result.transType = transType;
    }
    if (transDate != null) {
      $result.transDate = transDate;
    }
    if (direction != null) {
      $result.direction = direction;
    }
    if (transCode != null) {
      $result.transCode = transCode;
    }
    if (customerCode != null) {
      $result.customerCode = customerCode;
    }
    if (employeeCode != null) {
      $result.employeeCode = employeeCode;
    }
    if (projectCode != null) {
      $result.projectCode = projectCode;
    }
    if (placeCode != null) {
      $result.placeCode = placeCode;
    }
    if (currencyCode != null) {
      $result.currencyCode = currencyCode;
    }
    if (authCode != null) {
      $result.authCode = authCode;
    }
    if (transMeta != null) {
      $result.transMeta = transMeta;
    }
    if (transMap != null) {
      $result.transMap = transMap;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Trans._() : super();
  factory Trans.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Trans.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Trans', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<TransType>(3, _omitFieldNames ? '' : 'transType', $pb.PbFieldType.OE, defaultOrMaker: TransType.TRANS_INVOICE, valueOf: TransType.valueOf, enumValues: TransType.values)
    ..aOS(4, _omitFieldNames ? '' : 'transDate')
    ..e<Direction>(5, _omitFieldNames ? '' : 'direction', $pb.PbFieldType.OE, defaultOrMaker: Direction.DIRECTION_OUT, valueOf: Direction.valueOf, enumValues: Direction.values)
    ..aOS(6, _omitFieldNames ? '' : 'transCode')
    ..aOS(7, _omitFieldNames ? '' : 'customerCode')
    ..aOS(8, _omitFieldNames ? '' : 'employeeCode')
    ..aOS(9, _omitFieldNames ? '' : 'projectCode')
    ..aOS(10, _omitFieldNames ? '' : 'placeCode')
    ..aOS(11, _omitFieldNames ? '' : 'currencyCode')
    ..aOS(12, _omitFieldNames ? '' : 'authCode')
    ..aOM<TransMeta>(13, _omitFieldNames ? '' : 'transMeta', subBuilder: TransMeta.create)
    ..aOM<JsonString>(14, _omitFieldNames ? '' : 'transMap', subBuilder: JsonString.create)
    ..aOS(15, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Trans clone() => Trans()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Trans copyWith(void Function(Trans) updates) => super.copyWith((message) => updates(message as Trans)) as Trans;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Trans create() => Trans._();
  Trans createEmptyInstance() => create();
  static $pb.PbList<Trans> createRepeated() => $pb.PbList<Trans>();
  @$core.pragma('dart2js:noInline')
  static Trans getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Trans>(create);
  static Trans? _defaultInstance;

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: TRN1731101982N123 ("TRN" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(3)
  TransType get transType => $_getN(2);
  @$pb.TagNumber(3)
  set transType(TransType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasTransType() => $_has(2);
  @$pb.TagNumber(3)
  void clearTransType() => clearField(3);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ISO 8601 date
  @$pb.TagNumber(4)
  $core.String get transDate => $_getSZ(3);
  @$pb.TagNumber(4)
  set transDate($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTransDate() => $_has(3);
  @$pb.TagNumber(4)
  void clearTransDate() => clearField(4);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(5)
  Direction get direction => $_getN(4);
  @$pb.TagNumber(5)
  set direction(Direction v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasDirection() => $_has(4);
  @$pb.TagNumber(5)
  void clearDirection() => clearField(5);

  /// Reference to <a href="#Trans">Trans</a>.Code
  @$pb.TagNumber(6)
  $core.String get transCode => $_getSZ(5);
  @$pb.TagNumber(6)
  set transCode($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasTransCode() => $_has(5);
  @$pb.TagNumber(6)
  void clearTransCode() => clearField(6);

  /// Reference to <a href="#Customer">Customer</a>.Code
  @$pb.TagNumber(7)
  $core.String get customerCode => $_getSZ(6);
  @$pb.TagNumber(7)
  set customerCode($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasCustomerCode() => $_has(6);
  @$pb.TagNumber(7)
  void clearCustomerCode() => clearField(7);

  /// Reference to <a href="#Employee">Employee</a>.Code
  @$pb.TagNumber(8)
  $core.String get employeeCode => $_getSZ(7);
  @$pb.TagNumber(8)
  set employeeCode($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasEmployeeCode() => $_has(7);
  @$pb.TagNumber(8)
  void clearEmployeeCode() => clearField(8);

  /// Reference to <a href="#Project">Project</a>.Code
  @$pb.TagNumber(9)
  $core.String get projectCode => $_getSZ(8);
  @$pb.TagNumber(9)
  set projectCode($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasProjectCode() => $_has(8);
  @$pb.TagNumber(9)
  void clearProjectCode() => clearField(9);

  /// Reference to <a href="#Place">Place</a>.Code
  @$pb.TagNumber(10)
  $core.String get placeCode => $_getSZ(9);
  @$pb.TagNumber(10)
  set placeCode($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasPlaceCode() => $_has(9);
  @$pb.TagNumber(10)
  void clearPlaceCode() => clearField(10);

  /// Reference to <a href="#Currency">Currency</a>.Code
  @$pb.TagNumber(11)
  $core.String get currencyCode => $_getSZ(10);
  @$pb.TagNumber(11)
  set currencyCode($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasCurrencyCode() => $_has(10);
  @$pb.TagNumber(11)
  void clearCurrencyCode() => clearField(11);

  /// Reference to <a href="#Auth">Auth</a>.Code
  @$pb.TagNumber(12)
  $core.String get authCode => $_getSZ(11);
  @$pb.TagNumber(12)
  set authCode($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasAuthCode() => $_has(11);
  @$pb.TagNumber(12)
  void clearAuthCode() => clearField(12);

  @$pb.TagNumber(13)
  TransMeta get transMeta => $_getN(12);
  @$pb.TagNumber(13)
  set transMeta(TransMeta v) { setField(13, v); }
  @$pb.TagNumber(13)
  $core.bool hasTransMeta() => $_has(12);
  @$pb.TagNumber(13)
  void clearTransMeta() => clearField(13);
  @$pb.TagNumber(13)
  TransMeta ensureTransMeta() => $_ensure(12);

  /// Flexible key-value map for additional metadata. The value is any json type.
  @$pb.TagNumber(14)
  JsonString get transMap => $_getN(13);
  @$pb.TagNumber(14)
  set transMap(JsonString v) { setField(14, v); }
  @$pb.TagNumber(14)
  $core.bool hasTransMap() => $_has(13);
  @$pb.TagNumber(14)
  void clearTransMap() => clearField(14);
  @$pb.TagNumber(14)
  JsonString ensureTransMap() => $_ensure(13);

  /// Timestamp of data creation
  @$pb.TagNumber(15)
  $core.String get timeStamp => $_getSZ(14);
  @$pb.TagNumber(15)
  set timeStamp($core.String v) { $_setString(14, v); }
  @$pb.TagNumber(15)
  $core.bool hasTimeStamp() => $_has(14);
  @$pb.TagNumber(15)
  void clearTimeStamp() => clearField(15);
}

class Transactions extends $pb.GeneratedMessage {
  factory Transactions({
    $core.Iterable<Trans>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Transactions._() : super();
  factory Transactions.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Transactions.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Transactions', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Trans>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Trans.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Transactions clone() => Transactions()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Transactions copyWith(void Function(Transactions) updates) => super.copyWith((message) => updates(message as Transactions)) as Transactions;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Transactions create() => Transactions._();
  Transactions createEmptyInstance() => create();
  static $pb.PbList<Transactions> createRepeated() => $pb.PbList<Transactions>();
  @$core.pragma('dart2js:noInline')
  static Transactions getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Transactions>(create);
  static Transactions? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Trans> get data => $_getList(0);
}

enum Config_Data {
  map, 
  shortcut, 
  message, 
  pattern, 
  report, 
  printQueue, 
  configData, 
  notSet
}

/// Config message
class Config extends $pb.GeneratedMessage {
  factory Config({
    $fixnum.Int64? id,
    $core.String? code,
    ConfigType? configType,
    ConfigMap? map,
    ConfigShortcut? shortcut,
    ConfigMessage? message,
    ConfigPattern? pattern,
    ConfigReport? report,
    ConfigPrintQueue? printQueue,
    JsonString? configData,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (code != null) {
      $result.code = code;
    }
    if (configType != null) {
      $result.configType = configType;
    }
    if (map != null) {
      $result.map = map;
    }
    if (shortcut != null) {
      $result.shortcut = shortcut;
    }
    if (message != null) {
      $result.message = message;
    }
    if (pattern != null) {
      $result.pattern = pattern;
    }
    if (report != null) {
      $result.report = report;
    }
    if (printQueue != null) {
      $result.printQueue = printQueue;
    }
    if (configData != null) {
      $result.configData = configData;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  Config._() : super();
  factory Config.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Config.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static const $core.Map<$core.int, Config_Data> _Config_DataByTag = {
    4 : Config_Data.map,
    5 : Config_Data.shortcut,
    6 : Config_Data.message,
    7 : Config_Data.pattern,
    8 : Config_Data.report,
    9 : Config_Data.printQueue,
    10 : Config_Data.configData,
    0 : Config_Data.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Config', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..oo(0, [4, 5, 6, 7, 8, 9, 10])
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'code')
    ..e<ConfigType>(3, _omitFieldNames ? '' : 'configType', $pb.PbFieldType.OE, defaultOrMaker: ConfigType.CONFIG_MAP, valueOf: ConfigType.valueOf, enumValues: ConfigType.values)
    ..aOM<ConfigMap>(4, _omitFieldNames ? '' : 'map', subBuilder: ConfigMap.create)
    ..aOM<ConfigShortcut>(5, _omitFieldNames ? '' : 'shortcut', subBuilder: ConfigShortcut.create)
    ..aOM<ConfigMessage>(6, _omitFieldNames ? '' : 'message', subBuilder: ConfigMessage.create)
    ..aOM<ConfigPattern>(7, _omitFieldNames ? '' : 'pattern', subBuilder: ConfigPattern.create)
    ..aOM<ConfigReport>(8, _omitFieldNames ? '' : 'report', subBuilder: ConfigReport.create)
    ..aOM<ConfigPrintQueue>(9, _omitFieldNames ? '' : 'printQueue', subBuilder: ConfigPrintQueue.create)
    ..aOM<JsonString>(10, _omitFieldNames ? '' : 'configData', subBuilder: JsonString.create)
    ..aOS(11, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Config clone() => Config()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Config copyWith(void Function(Config) updates) => super.copyWith((message) => updates(message as Config)) as Config;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Config create() => Config._();
  Config createEmptyInstance() => create();
  static $pb.PbList<Config> createRepeated() => $pb.PbList<Config>();
  @$core.pragma('dart2js:noInline')
  static Config getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Config>(create);
  static Config? _defaultInstance;

  Config_Data whichData() => _Config_DataByTag[$_whichOneof(0)]!;
  void clearData() => clearField($_whichOneof(0));

  /// Database primary key
  /// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  /// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
  /// Example: CNF1731101982N123 ("CNF" + UNIX Time stamp + "N" + current ID)
  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  /// <span class="text-red-500 font-bold">*</span> Required field.
  /// ENUM field.
  @$pb.TagNumber(3)
  ConfigType get configType => $_getN(2);
  @$pb.TagNumber(3)
  set configType(ConfigType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasConfigType() => $_has(2);
  @$pb.TagNumber(3)
  void clearConfigType() => clearField(3);

  @$pb.TagNumber(4)
  ConfigMap get map => $_getN(3);
  @$pb.TagNumber(4)
  set map(ConfigMap v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasMap() => $_has(3);
  @$pb.TagNumber(4)
  void clearMap() => clearField(4);
  @$pb.TagNumber(4)
  ConfigMap ensureMap() => $_ensure(3);

  @$pb.TagNumber(5)
  ConfigShortcut get shortcut => $_getN(4);
  @$pb.TagNumber(5)
  set shortcut(ConfigShortcut v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasShortcut() => $_has(4);
  @$pb.TagNumber(5)
  void clearShortcut() => clearField(5);
  @$pb.TagNumber(5)
  ConfigShortcut ensureShortcut() => $_ensure(4);

  @$pb.TagNumber(6)
  ConfigMessage get message => $_getN(5);
  @$pb.TagNumber(6)
  set message(ConfigMessage v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasMessage() => $_has(5);
  @$pb.TagNumber(6)
  void clearMessage() => clearField(6);
  @$pb.TagNumber(6)
  ConfigMessage ensureMessage() => $_ensure(5);

  @$pb.TagNumber(7)
  ConfigPattern get pattern => $_getN(6);
  @$pb.TagNumber(7)
  set pattern(ConfigPattern v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasPattern() => $_has(6);
  @$pb.TagNumber(7)
  void clearPattern() => clearField(7);
  @$pb.TagNumber(7)
  ConfigPattern ensurePattern() => $_ensure(6);

  @$pb.TagNumber(8)
  ConfigReport get report => $_getN(7);
  @$pb.TagNumber(8)
  set report(ConfigReport v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasReport() => $_has(7);
  @$pb.TagNumber(8)
  void clearReport() => clearField(8);
  @$pb.TagNumber(8)
  ConfigReport ensureReport() => $_ensure(7);

  @$pb.TagNumber(9)
  ConfigPrintQueue get printQueue => $_getN(8);
  @$pb.TagNumber(9)
  set printQueue(ConfigPrintQueue v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasPrintQueue() => $_has(8);
  @$pb.TagNumber(9)
  void clearPrintQueue() => clearField(9);
  @$pb.TagNumber(9)
  ConfigPrintQueue ensurePrintQueue() => $_ensure(8);

  @$pb.TagNumber(10)
  JsonString get configData => $_getN(9);
  @$pb.TagNumber(10)
  set configData(JsonString v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasConfigData() => $_has(9);
  @$pb.TagNumber(10)
  void clearConfigData() => clearField(10);
  @$pb.TagNumber(10)
  JsonString ensureConfigData() => $_ensure(9);

  /// Timestamp of data creation
  @$pb.TagNumber(11)
  $core.String get timeStamp => $_getSZ(10);
  @$pb.TagNumber(11)
  set timeStamp($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasTimeStamp() => $_has(10);
  @$pb.TagNumber(11)
  void clearTimeStamp() => clearField(11);
}

class Configs extends $pb.GeneratedMessage {
  factory Configs({
    $core.Iterable<Config>? data,
  }) {
    final $result = create();
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  Configs._() : super();
  factory Configs.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Configs.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Configs', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..pc<Config>(1, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Config.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Configs clone() => Configs()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Configs copyWith(void Function(Configs) updates) => super.copyWith((message) => updates(message as Configs)) as Configs;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Configs create() => Configs._();
  Configs createEmptyInstance() => create();
  static $pb.PbList<Configs> createRepeated() => $pb.PbList<Configs>();
  @$core.pragma('dart2js:noInline')
  static Configs getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Configs>(create);
  static Configs? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Config> get data => $_getList(0);
}

/// ConfigMap message
class ConfigMap extends $pb.GeneratedMessage {
  factory ConfigMap({
    $core.String? fieldName,
    FieldType? fieldType,
    $core.String? description,
    $core.Iterable<$core.String>? tags,
    $core.Iterable<MapFilter>? filter,
  }) {
    final $result = create();
    if (fieldName != null) {
      $result.fieldName = fieldName;
    }
    if (fieldType != null) {
      $result.fieldType = fieldType;
    }
    if (description != null) {
      $result.description = description;
    }
    if (tags != null) {
      $result.tags.addAll(tags);
    }
    if (filter != null) {
      $result.filter.addAll(filter);
    }
    return $result;
  }
  ConfigMap._() : super();
  factory ConfigMap.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfigMap.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ConfigMap', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'fieldName')
    ..e<FieldType>(2, _omitFieldNames ? '' : 'fieldType', $pb.PbFieldType.OE, defaultOrMaker: FieldType.FIELD_STRING, valueOf: FieldType.valueOf, enumValues: FieldType.values)
    ..aOS(3, _omitFieldNames ? '' : 'description')
    ..pPS(4, _omitFieldNames ? '' : 'tags')
    ..pc<MapFilter>(5, _omitFieldNames ? '' : 'filter', $pb.PbFieldType.KE, valueOf: MapFilter.valueOf, enumValues: MapFilter.values, defaultEnumValue: MapFilter.FILTER_ADDRESS)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ConfigMap clone() => ConfigMap()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ConfigMap copyWith(void Function(ConfigMap) updates) => super.copyWith((message) => updates(message as ConfigMap)) as ConfigMap;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ConfigMap create() => ConfigMap._();
  ConfigMap createEmptyInstance() => create();
  static $pb.PbList<ConfigMap> createRepeated() => $pb.PbList<ConfigMap>();
  @$core.pragma('dart2js:noInline')
  static ConfigMap getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfigMap>(create);
  static ConfigMap? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get fieldName => $_getSZ(0);
  @$pb.TagNumber(1)
  set fieldName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFieldName() => $_has(0);
  @$pb.TagNumber(1)
  void clearFieldName() => clearField(1);

  /// ENUM field.
  @$pb.TagNumber(2)
  FieldType get fieldType => $_getN(1);
  @$pb.TagNumber(2)
  set fieldType(FieldType v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasFieldType() => $_has(1);
  @$pb.TagNumber(2)
  void clearFieldType() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get description => $_getSZ(2);
  @$pb.TagNumber(3)
  set description($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDescription() => $_has(2);
  @$pb.TagNumber(3)
  void clearDescription() => clearField(3);

  /// ENUM list if FieldType is ENUM
  /// ENUM field. Valid values: ADDRESS, BARCODE, CONTACT, CURRENCY, CUSTOMER, EMPLOYEE, EVENT, ITEM,
  /// MOVEMENT, PAYMENT, PLACE, PRICE, PRODUCT, PROJECT, RATE, TAX, TOOL, USER, TRANS,
  /// INVOICE, RECEIPT, ORDER, OFFER, WORKSHEET, RENT, DELIVERY,
  /// INVENTORY, WAYBILL, PRODUCTION, FORMULA, BANK, CASH
  @$pb.TagNumber(4)
  $core.List<$core.String> get tags => $_getList(3);

  /// ENUM list if FieldType is FILTER
  @$pb.TagNumber(5)
  $core.List<MapFilter> get filter => $_getList(4);
}

/// ConfigShortcutField message
class ConfigShortcutField extends $pb.GeneratedMessage {
  factory ConfigShortcutField({
    $core.String? fieldName,
    $core.String? description,
    ShortcutField? fieldType,
    $fixnum.Int64? order,
  }) {
    final $result = create();
    if (fieldName != null) {
      $result.fieldName = fieldName;
    }
    if (description != null) {
      $result.description = description;
    }
    if (fieldType != null) {
      $result.fieldType = fieldType;
    }
    if (order != null) {
      $result.order = order;
    }
    return $result;
  }
  ConfigShortcutField._() : super();
  factory ConfigShortcutField.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfigShortcutField.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ConfigShortcutField', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'fieldName')
    ..aOS(2, _omitFieldNames ? '' : 'description')
    ..e<ShortcutField>(3, _omitFieldNames ? '' : 'fieldType', $pb.PbFieldType.OE, defaultOrMaker: ShortcutField.SHORTCUT_STRING, valueOf: ShortcutField.valueOf, enumValues: ShortcutField.values)
    ..aInt64(4, _omitFieldNames ? '' : 'order')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ConfigShortcutField clone() => ConfigShortcutField()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ConfigShortcutField copyWith(void Function(ConfigShortcutField) updates) => super.copyWith((message) => updates(message as ConfigShortcutField)) as ConfigShortcutField;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ConfigShortcutField create() => ConfigShortcutField._();
  ConfigShortcutField createEmptyInstance() => create();
  static $pb.PbList<ConfigShortcutField> createRepeated() => $pb.PbList<ConfigShortcutField>();
  @$core.pragma('dart2js:noInline')
  static ConfigShortcutField getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfigShortcutField>(create);
  static ConfigShortcutField? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get fieldName => $_getSZ(0);
  @$pb.TagNumber(1)
  set fieldName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFieldName() => $_has(0);
  @$pb.TagNumber(1)
  void clearFieldName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get description => $_getSZ(1);
  @$pb.TagNumber(2)
  set description($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDescription() => $_has(1);
  @$pb.TagNumber(2)
  void clearDescription() => clearField(2);

  /// ENUM field.
  @$pb.TagNumber(3)
  ShortcutField get fieldType => $_getN(2);
  @$pb.TagNumber(3)
  set fieldType(ShortcutField v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasFieldType() => $_has(2);
  @$pb.TagNumber(3)
  void clearFieldType() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get order => $_getI64(3);
  @$pb.TagNumber(4)
  set order($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasOrder() => $_has(3);
  @$pb.TagNumber(4)
  void clearOrder() => clearField(4);
}

/// ConfigShortcut message
class ConfigShortcut extends $pb.GeneratedMessage {
  factory ConfigShortcut({
    $core.String? shortcutKey,
    $core.String? description,
    $core.String? modul,
    $core.String? icon,
    ShortcutMethod? method,
    $core.String? funcname,
    $core.String? address,
    $core.Iterable<ConfigShortcutField>? fields,
  }) {
    final $result = create();
    if (shortcutKey != null) {
      $result.shortcutKey = shortcutKey;
    }
    if (description != null) {
      $result.description = description;
    }
    if (modul != null) {
      $result.modul = modul;
    }
    if (icon != null) {
      $result.icon = icon;
    }
    if (method != null) {
      $result.method = method;
    }
    if (funcname != null) {
      $result.funcname = funcname;
    }
    if (address != null) {
      $result.address = address;
    }
    if (fields != null) {
      $result.fields.addAll(fields);
    }
    return $result;
  }
  ConfigShortcut._() : super();
  factory ConfigShortcut.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfigShortcut.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ConfigShortcut', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'shortcutKey')
    ..aOS(2, _omitFieldNames ? '' : 'description')
    ..aOS(3, _omitFieldNames ? '' : 'modul')
    ..aOS(4, _omitFieldNames ? '' : 'icon')
    ..e<ShortcutMethod>(5, _omitFieldNames ? '' : 'method', $pb.PbFieldType.OE, defaultOrMaker: ShortcutMethod.METHOD_GET, valueOf: ShortcutMethod.valueOf, enumValues: ShortcutMethod.values)
    ..aOS(6, _omitFieldNames ? '' : 'funcname')
    ..aOS(7, _omitFieldNames ? '' : 'address')
    ..pc<ConfigShortcutField>(8, _omitFieldNames ? '' : 'fields', $pb.PbFieldType.PM, subBuilder: ConfigShortcutField.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ConfigShortcut clone() => ConfigShortcut()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ConfigShortcut copyWith(void Function(ConfigShortcut) updates) => super.copyWith((message) => updates(message as ConfigShortcut)) as ConfigShortcut;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ConfigShortcut create() => ConfigShortcut._();
  ConfigShortcut createEmptyInstance() => create();
  static $pb.PbList<ConfigShortcut> createRepeated() => $pb.PbList<ConfigShortcut>();
  @$core.pragma('dart2js:noInline')
  static ConfigShortcut getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfigShortcut>(create);
  static ConfigShortcut? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get shortcutKey => $_getSZ(0);
  @$pb.TagNumber(1)
  set shortcutKey($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasShortcutKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearShortcutKey() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get description => $_getSZ(1);
  @$pb.TagNumber(2)
  set description($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDescription() => $_has(1);
  @$pb.TagNumber(2)
  void clearDescription() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get modul => $_getSZ(2);
  @$pb.TagNumber(3)
  set modul($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasModul() => $_has(2);
  @$pb.TagNumber(3)
  void clearModul() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get icon => $_getSZ(3);
  @$pb.TagNumber(4)
  set icon($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasIcon() => $_has(3);
  @$pb.TagNumber(4)
  void clearIcon() => clearField(4);

  /// ENUM field.
  @$pb.TagNumber(5)
  ShortcutMethod get method => $_getN(4);
  @$pb.TagNumber(5)
  set method(ShortcutMethod v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasMethod() => $_has(4);
  @$pb.TagNumber(5)
  void clearMethod() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get funcname => $_getSZ(5);
  @$pb.TagNumber(6)
  set funcname($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasFuncname() => $_has(5);
  @$pb.TagNumber(6)
  void clearFuncname() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get address => $_getSZ(6);
  @$pb.TagNumber(7)
  set address($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasAddress() => $_has(6);
  @$pb.TagNumber(7)
  void clearAddress() => clearField(7);

  @$pb.TagNumber(8)
  $core.List<ConfigShortcutField> get fields => $_getList(7);
}

/// ConfigMessage message
class ConfigMessage extends $pb.GeneratedMessage {
  factory ConfigMessage({
    $core.String? section,
    $core.String? key,
    $core.String? lang,
    $core.String? value,
  }) {
    final $result = create();
    if (section != null) {
      $result.section = section;
    }
    if (key != null) {
      $result.key = key;
    }
    if (lang != null) {
      $result.lang = lang;
    }
    if (value != null) {
      $result.value = value;
    }
    return $result;
  }
  ConfigMessage._() : super();
  factory ConfigMessage.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfigMessage.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ConfigMessage', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'section')
    ..aOS(2, _omitFieldNames ? '' : 'key')
    ..aOS(3, _omitFieldNames ? '' : 'lang')
    ..aOS(4, _omitFieldNames ? '' : 'value')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ConfigMessage clone() => ConfigMessage()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ConfigMessage copyWith(void Function(ConfigMessage) updates) => super.copyWith((message) => updates(message as ConfigMessage)) as ConfigMessage;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ConfigMessage create() => ConfigMessage._();
  ConfigMessage createEmptyInstance() => create();
  static $pb.PbList<ConfigMessage> createRepeated() => $pb.PbList<ConfigMessage>();
  @$core.pragma('dart2js:noInline')
  static ConfigMessage getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfigMessage>(create);
  static ConfigMessage? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get section => $_getSZ(0);
  @$pb.TagNumber(1)
  set section($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSection() => $_has(0);
  @$pb.TagNumber(1)
  void clearSection() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get key => $_getSZ(1);
  @$pb.TagNumber(2)
  set key($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearKey() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get lang => $_getSZ(2);
  @$pb.TagNumber(3)
  set lang($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLang() => $_has(2);
  @$pb.TagNumber(3)
  void clearLang() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get value => $_getSZ(3);
  @$pb.TagNumber(4)
  set value($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasValue() => $_has(3);
  @$pb.TagNumber(4)
  void clearValue() => clearField(4);
}

/// ConfigPattern message
class ConfigPattern extends $pb.GeneratedMessage {
  factory ConfigPattern({
    TransType? transType,
    $core.String? description,
    $core.String? notes,
    $core.bool? defaultPattern,
  }) {
    final $result = create();
    if (transType != null) {
      $result.transType = transType;
    }
    if (description != null) {
      $result.description = description;
    }
    if (notes != null) {
      $result.notes = notes;
    }
    if (defaultPattern != null) {
      $result.defaultPattern = defaultPattern;
    }
    return $result;
  }
  ConfigPattern._() : super();
  factory ConfigPattern.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfigPattern.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ConfigPattern', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..e<TransType>(1, _omitFieldNames ? '' : 'transType', $pb.PbFieldType.OE, defaultOrMaker: TransType.TRANS_INVOICE, valueOf: TransType.valueOf, enumValues: TransType.values)
    ..aOS(2, _omitFieldNames ? '' : 'description')
    ..aOS(3, _omitFieldNames ? '' : 'notes')
    ..aOB(4, _omitFieldNames ? '' : 'defaultPattern')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ConfigPattern clone() => ConfigPattern()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ConfigPattern copyWith(void Function(ConfigPattern) updates) => super.copyWith((message) => updates(message as ConfigPattern)) as ConfigPattern;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ConfigPattern create() => ConfigPattern._();
  ConfigPattern createEmptyInstance() => create();
  static $pb.PbList<ConfigPattern> createRepeated() => $pb.PbList<ConfigPattern>();
  @$core.pragma('dart2js:noInline')
  static ConfigPattern getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfigPattern>(create);
  static ConfigPattern? _defaultInstance;

  /// ENUM field.
  @$pb.TagNumber(1)
  TransType get transType => $_getN(0);
  @$pb.TagNumber(1)
  set transType(TransType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTransType() => $_has(0);
  @$pb.TagNumber(1)
  void clearTransType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get description => $_getSZ(1);
  @$pb.TagNumber(2)
  set description($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDescription() => $_has(1);
  @$pb.TagNumber(2)
  void clearDescription() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get notes => $_getSZ(2);
  @$pb.TagNumber(3)
  set notes($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasNotes() => $_has(2);
  @$pb.TagNumber(3)
  void clearNotes() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get defaultPattern => $_getBF(3);
  @$pb.TagNumber(4)
  set defaultPattern($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasDefaultPattern() => $_has(3);
  @$pb.TagNumber(4)
  void clearDefaultPattern() => clearField(4);
}

/// ConfigPrintQueue message
class ConfigPrintQueue extends $pb.GeneratedMessage {
  factory ConfigPrintQueue({
    $core.String? refType,
    $core.String? refCode,
    $fixnum.Int64? qty,
    $core.String? reportCode,
    $core.String? authCode,
    $core.String? timeStamp,
  }) {
    final $result = create();
    if (refType != null) {
      $result.refType = refType;
    }
    if (refCode != null) {
      $result.refCode = refCode;
    }
    if (qty != null) {
      $result.qty = qty;
    }
    if (reportCode != null) {
      $result.reportCode = reportCode;
    }
    if (authCode != null) {
      $result.authCode = authCode;
    }
    if (timeStamp != null) {
      $result.timeStamp = timeStamp;
    }
    return $result;
  }
  ConfigPrintQueue._() : super();
  factory ConfigPrintQueue.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfigPrintQueue.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ConfigPrintQueue', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'refType')
    ..aOS(2, _omitFieldNames ? '' : 'refCode')
    ..aInt64(3, _omitFieldNames ? '' : 'qty')
    ..aOS(4, _omitFieldNames ? '' : 'reportCode')
    ..aOS(5, _omitFieldNames ? '' : 'authCode')
    ..aOS(6, _omitFieldNames ? '' : 'timeStamp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ConfigPrintQueue clone() => ConfigPrintQueue()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ConfigPrintQueue copyWith(void Function(ConfigPrintQueue) updates) => super.copyWith((message) => updates(message as ConfigPrintQueue)) as ConfigPrintQueue;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ConfigPrintQueue create() => ConfigPrintQueue._();
  ConfigPrintQueue createEmptyInstance() => create();
  static $pb.PbList<ConfigPrintQueue> createRepeated() => $pb.PbList<ConfigPrintQueue>();
  @$core.pragma('dart2js:noInline')
  static ConfigPrintQueue getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfigPrintQueue>(create);
  static ConfigPrintQueue? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get refType => $_getSZ(0);
  @$pb.TagNumber(1)
  set refType($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRefType() => $_has(0);
  @$pb.TagNumber(1)
  void clearRefType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get refCode => $_getSZ(1);
  @$pb.TagNumber(2)
  set refCode($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRefCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearRefCode() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get qty => $_getI64(2);
  @$pb.TagNumber(3)
  set qty($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasQty() => $_has(2);
  @$pb.TagNumber(3)
  void clearQty() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get reportCode => $_getSZ(3);
  @$pb.TagNumber(4)
  set reportCode($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasReportCode() => $_has(3);
  @$pb.TagNumber(4)
  void clearReportCode() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get authCode => $_getSZ(4);
  @$pb.TagNumber(5)
  set authCode($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasAuthCode() => $_has(4);
  @$pb.TagNumber(5)
  void clearAuthCode() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get timeStamp => $_getSZ(5);
  @$pb.TagNumber(6)
  set timeStamp($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasTimeStamp() => $_has(5);
  @$pb.TagNumber(6)
  void clearTimeStamp() => clearField(6);
}

/// ConfigReport message
class ConfigReport extends $pb.GeneratedMessage {
  factory ConfigReport({
    $core.String? reportKey,
    $core.String? reportType,
    TransType? transType,
    Direction? direction,
    FileType? fileType,
    $core.String? reportName,
    $core.String? description,
    $core.String? label,
    $core.String? template,
  }) {
    final $result = create();
    if (reportKey != null) {
      $result.reportKey = reportKey;
    }
    if (reportType != null) {
      $result.reportType = reportType;
    }
    if (transType != null) {
      $result.transType = transType;
    }
    if (direction != null) {
      $result.direction = direction;
    }
    if (fileType != null) {
      $result.fileType = fileType;
    }
    if (reportName != null) {
      $result.reportName = reportName;
    }
    if (description != null) {
      $result.description = description;
    }
    if (label != null) {
      $result.label = label;
    }
    if (template != null) {
      $result.template = template;
    }
    return $result;
  }
  ConfigReport._() : super();
  factory ConfigReport.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfigReport.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ConfigReport', package: const $pb.PackageName(_omitMessageNames ? '' : 'nervatura'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'reportKey')
    ..aOS(2, _omitFieldNames ? '' : 'reportType')
    ..e<TransType>(3, _omitFieldNames ? '' : 'transType', $pb.PbFieldType.OE, defaultOrMaker: TransType.TRANS_INVOICE, valueOf: TransType.valueOf, enumValues: TransType.values)
    ..e<Direction>(4, _omitFieldNames ? '' : 'direction', $pb.PbFieldType.OE, defaultOrMaker: Direction.DIRECTION_OUT, valueOf: Direction.valueOf, enumValues: Direction.values)
    ..e<FileType>(5, _omitFieldNames ? '' : 'fileType', $pb.PbFieldType.OE, defaultOrMaker: FileType.FILE_PDF, valueOf: FileType.valueOf, enumValues: FileType.values)
    ..aOS(6, _omitFieldNames ? '' : 'reportName')
    ..aOS(7, _omitFieldNames ? '' : 'description')
    ..aOS(8, _omitFieldNames ? '' : 'label')
    ..aOS(9, _omitFieldNames ? '' : 'template')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ConfigReport clone() => ConfigReport()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ConfigReport copyWith(void Function(ConfigReport) updates) => super.copyWith((message) => updates(message as ConfigReport)) as ConfigReport;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ConfigReport create() => ConfigReport._();
  ConfigReport createEmptyInstance() => create();
  static $pb.PbList<ConfigReport> createRepeated() => $pb.PbList<ConfigReport>();
  @$core.pragma('dart2js:noInline')
  static ConfigReport getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfigReport>(create);
  static ConfigReport? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get reportKey => $_getSZ(0);
  @$pb.TagNumber(1)
  set reportKey($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasReportKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearReportKey() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get reportType => $_getSZ(1);
  @$pb.TagNumber(2)
  set reportType($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasReportType() => $_has(1);
  @$pb.TagNumber(2)
  void clearReportType() => clearField(2);

  /// ENUM field.
  @$pb.TagNumber(3)
  TransType get transType => $_getN(2);
  @$pb.TagNumber(3)
  set transType(TransType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasTransType() => $_has(2);
  @$pb.TagNumber(3)
  void clearTransType() => clearField(3);

  /// ENUM field.
  @$pb.TagNumber(4)
  Direction get direction => $_getN(3);
  @$pb.TagNumber(4)
  set direction(Direction v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasDirection() => $_has(3);
  @$pb.TagNumber(4)
  void clearDirection() => clearField(4);

  /// ENUM field.
  @$pb.TagNumber(5)
  FileType get fileType => $_getN(4);
  @$pb.TagNumber(5)
  set fileType(FileType v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasFileType() => $_has(4);
  @$pb.TagNumber(5)
  void clearFileType() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get reportName => $_getSZ(5);
  @$pb.TagNumber(6)
  set reportName($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasReportName() => $_has(5);
  @$pb.TagNumber(6)
  void clearReportName() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get description => $_getSZ(6);
  @$pb.TagNumber(7)
  set description($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasDescription() => $_has(6);
  @$pb.TagNumber(7)
  void clearDescription() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get label => $_getSZ(7);
  @$pb.TagNumber(8)
  set label($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasLabel() => $_has(7);
  @$pb.TagNumber(8)
  void clearLabel() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get template => $_getSZ(8);
  @$pb.TagNumber(9)
  set template($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasTemplate() => $_has(8);
  @$pb.TagNumber(9)
  void clearTemplate() => clearField(9);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
