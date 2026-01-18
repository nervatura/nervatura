//
//  Generated code. Do not modify.
//  source: store.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'store.pb.dart' as $0;

export 'store.pb.dart';

@$pb.GrpcServiceName('nervatura.API')
class APIClient extends $grpc.Client {
  static final _$authLogin = $grpc.ClientMethod<$0.RequestAuthLogin, $0.ResponseAuthLogin>(
      '/nervatura.API/AuthLogin',
      ($0.RequestAuthLogin value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ResponseAuthLogin.fromBuffer(value));
  static final _$authUpdate = $grpc.ClientMethod<$0.Auth, $0.Auth>(
      '/nervatura.API/AuthUpdate',
      ($0.Auth value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Auth.fromBuffer(value));
  static final _$authGet = $grpc.ClientMethod<$0.RequestGet, $0.Auth>(
      '/nervatura.API/AuthGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Auth.fromBuffer(value));
  static final _$authPassword = $grpc.ClientMethod<$0.RequestPasswordChange, $0.ResponseStatus>(
      '/nervatura.API/AuthPassword',
      ($0.RequestPasswordChange value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ResponseStatus.fromBuffer(value));
  static final _$authPasswordReset = $grpc.ClientMethod<$0.RequestGet, $0.ResponseStatus>(
      '/nervatura.API/AuthPasswordReset',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ResponseStatus.fromBuffer(value));
  static final _$tokenLogin = $grpc.ClientMethod<$0.RequestEmpty, $0.Auth>(
      '/nervatura.API/TokenLogin',
      ($0.RequestEmpty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Auth.fromBuffer(value));
  static final _$tokenRefresh = $grpc.ClientMethod<$0.RequestEmpty, $0.ResponseAuthLogin>(
      '/nervatura.API/TokenRefresh',
      ($0.RequestEmpty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ResponseAuthLogin.fromBuffer(value));
  static final _$tokenDecode = $grpc.ClientMethod<$0.RequestTokenDecode, $0.ResponseTokenDecode>(
      '/nervatura.API/TokenDecode',
      ($0.RequestTokenDecode value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ResponseTokenDecode.fromBuffer(value));
  static final _$database = $grpc.ClientMethod<$0.RequestDatabase, $0.JsonBytes>(
      '/nervatura.API/Database',
      ($0.RequestDatabase value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.JsonBytes.fromBuffer(value));
  static final _$delete = $grpc.ClientMethod<$0.RequestDelete, $0.ResponseStatus>(
      '/nervatura.API/Delete',
      ($0.RequestDelete value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ResponseStatus.fromBuffer(value));
  static final _$function = $grpc.ClientMethod<$0.RequestFunction, $0.JsonBytes>(
      '/nervatura.API/Function',
      ($0.RequestFunction value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.JsonBytes.fromBuffer(value));
  static final _$view = $grpc.ClientMethod<$0.RequestView, $0.JsonBytes>(
      '/nervatura.API/View',
      ($0.RequestView value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.JsonBytes.fromBuffer(value));
  static final _$configUpdate = $grpc.ClientMethod<$0.Config, $0.Config>(
      '/nervatura.API/ConfigUpdate',
      ($0.Config value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Config.fromBuffer(value));
  static final _$configGet = $grpc.ClientMethod<$0.RequestGet, $0.Config>(
      '/nervatura.API/ConfigGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Config.fromBuffer(value));
  static final _$configQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Configs>(
      '/nervatura.API/ConfigQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Configs.fromBuffer(value));
  static final _$currencyUpdate = $grpc.ClientMethod<$0.Currency, $0.Currency>(
      '/nervatura.API/CurrencyUpdate',
      ($0.Currency value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Currency.fromBuffer(value));
  static final _$currencyGet = $grpc.ClientMethod<$0.RequestGet, $0.Currency>(
      '/nervatura.API/CurrencyGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Currency.fromBuffer(value));
  static final _$currencyQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Currencies>(
      '/nervatura.API/CurrencyQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Currencies.fromBuffer(value));
  static final _$customerUpdate = $grpc.ClientMethod<$0.Customer, $0.Customer>(
      '/nervatura.API/CustomerUpdate',
      ($0.Customer value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Customer.fromBuffer(value));
  static final _$customerGet = $grpc.ClientMethod<$0.RequestGet, $0.Customer>(
      '/nervatura.API/CustomerGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Customer.fromBuffer(value));
  static final _$customerQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Customers>(
      '/nervatura.API/CustomerQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Customers.fromBuffer(value));
  static final _$employeeUpdate = $grpc.ClientMethod<$0.Employee, $0.Employee>(
      '/nervatura.API/EmployeeUpdate',
      ($0.Employee value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Employee.fromBuffer(value));
  static final _$employeeGet = $grpc.ClientMethod<$0.RequestGet, $0.Employee>(
      '/nervatura.API/EmployeeGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Employee.fromBuffer(value));
  static final _$employeeQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Employees>(
      '/nervatura.API/EmployeeQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Employees.fromBuffer(value));
  static final _$itemUpdate = $grpc.ClientMethod<$0.Item, $0.Item>(
      '/nervatura.API/ItemUpdate',
      ($0.Item value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Item.fromBuffer(value));
  static final _$itemGet = $grpc.ClientMethod<$0.RequestGet, $0.Item>(
      '/nervatura.API/ItemGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Item.fromBuffer(value));
  static final _$itemQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Items>(
      '/nervatura.API/ItemQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Items.fromBuffer(value));
  static final _$linkUpdate = $grpc.ClientMethod<$0.Link, $0.Link>(
      '/nervatura.API/LinkUpdate',
      ($0.Link value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Link.fromBuffer(value));
  static final _$linkGet = $grpc.ClientMethod<$0.RequestGet, $0.Link>(
      '/nervatura.API/LinkGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Link.fromBuffer(value));
  static final _$linkQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Links>(
      '/nervatura.API/LinkQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Links.fromBuffer(value));
  static final _$logGet = $grpc.ClientMethod<$0.RequestGet, $0.Log>(
      '/nervatura.API/LogGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Log.fromBuffer(value));
  static final _$logQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Logs>(
      '/nervatura.API/LogQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Logs.fromBuffer(value));
  static final _$movementUpdate = $grpc.ClientMethod<$0.Movement, $0.Movement>(
      '/nervatura.API/MovementUpdate',
      ($0.Movement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Movement.fromBuffer(value));
  static final _$movementGet = $grpc.ClientMethod<$0.RequestGet, $0.Movement>(
      '/nervatura.API/MovementGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Movement.fromBuffer(value));
  static final _$movementQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Movements>(
      '/nervatura.API/MovementQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Movements.fromBuffer(value));
  static final _$paymentUpdate = $grpc.ClientMethod<$0.Payment, $0.Payment>(
      '/nervatura.API/PaymentUpdate',
      ($0.Payment value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Payment.fromBuffer(value));
  static final _$paymentGet = $grpc.ClientMethod<$0.RequestGet, $0.Payment>(
      '/nervatura.API/PaymentGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Payment.fromBuffer(value));
  static final _$paymentQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Payments>(
      '/nervatura.API/PaymentQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Payments.fromBuffer(value));
  static final _$placeUpdate = $grpc.ClientMethod<$0.Place, $0.Place>(
      '/nervatura.API/PlaceUpdate',
      ($0.Place value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Place.fromBuffer(value));
  static final _$placeGet = $grpc.ClientMethod<$0.RequestGet, $0.Place>(
      '/nervatura.API/PlaceGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Place.fromBuffer(value));
  static final _$placeQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Places>(
      '/nervatura.API/PlaceQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Places.fromBuffer(value));
  static final _$priceUpdate = $grpc.ClientMethod<$0.Price, $0.Price>(
      '/nervatura.API/PriceUpdate',
      ($0.Price value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Price.fromBuffer(value));
  static final _$priceGet = $grpc.ClientMethod<$0.RequestGet, $0.Price>(
      '/nervatura.API/PriceGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Price.fromBuffer(value));
  static final _$priceQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Prices>(
      '/nervatura.API/PriceQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Prices.fromBuffer(value));
  static final _$productUpdate = $grpc.ClientMethod<$0.Product, $0.Product>(
      '/nervatura.API/ProductUpdate',
      ($0.Product value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Product.fromBuffer(value));
  static final _$productGet = $grpc.ClientMethod<$0.RequestGet, $0.Product>(
      '/nervatura.API/ProductGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Product.fromBuffer(value));
  static final _$productQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Products>(
      '/nervatura.API/ProductQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Products.fromBuffer(value));
  static final _$projectUpdate = $grpc.ClientMethod<$0.Project, $0.Project>(
      '/nervatura.API/ProjectUpdate',
      ($0.Project value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Project.fromBuffer(value));
  static final _$projectGet = $grpc.ClientMethod<$0.RequestGet, $0.Project>(
      '/nervatura.API/ProjectGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Project.fromBuffer(value));
  static final _$projectQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Projects>(
      '/nervatura.API/ProjectQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Projects.fromBuffer(value));
  static final _$rateUpdate = $grpc.ClientMethod<$0.Rate, $0.Rate>(
      '/nervatura.API/RateUpdate',
      ($0.Rate value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Rate.fromBuffer(value));
  static final _$rateGet = $grpc.ClientMethod<$0.RequestGet, $0.Rate>(
      '/nervatura.API/RateGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Rate.fromBuffer(value));
  static final _$rateQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Rates>(
      '/nervatura.API/RateQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Rates.fromBuffer(value));
  static final _$taxUpdate = $grpc.ClientMethod<$0.Tax, $0.Tax>(
      '/nervatura.API/TaxUpdate',
      ($0.Tax value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Tax.fromBuffer(value));
  static final _$taxGet = $grpc.ClientMethod<$0.RequestGet, $0.Tax>(
      '/nervatura.API/TaxGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Tax.fromBuffer(value));
  static final _$taxQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Taxes>(
      '/nervatura.API/TaxQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Taxes.fromBuffer(value));
  static final _$toolUpdate = $grpc.ClientMethod<$0.Tool, $0.Tool>(
      '/nervatura.API/ToolUpdate',
      ($0.Tool value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Tool.fromBuffer(value));
  static final _$toolGet = $grpc.ClientMethod<$0.RequestGet, $0.Tool>(
      '/nervatura.API/ToolGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Tool.fromBuffer(value));
  static final _$toolQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Tools>(
      '/nervatura.API/ToolQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Tools.fromBuffer(value));
  static final _$transUpdate = $grpc.ClientMethod<$0.Trans, $0.Trans>(
      '/nervatura.API/TransUpdate',
      ($0.Trans value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Trans.fromBuffer(value));
  static final _$transGet = $grpc.ClientMethod<$0.RequestGet, $0.Trans>(
      '/nervatura.API/TransGet',
      ($0.RequestGet value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Trans.fromBuffer(value));
  static final _$transQuery = $grpc.ClientMethod<$0.RequestQuery, $0.Transactions>(
      '/nervatura.API/TransQuery',
      ($0.RequestQuery value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Transactions.fromBuffer(value));

  APIClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.ResponseAuthLogin> authLogin($0.RequestAuthLogin request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$authLogin, request, options: options);
  }

  $grpc.ResponseFuture<$0.Auth> authUpdate($0.Auth request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$authUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Auth> authGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$authGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.ResponseStatus> authPassword($0.RequestPasswordChange request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$authPassword, request, options: options);
  }

  $grpc.ResponseFuture<$0.ResponseStatus> authPasswordReset($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$authPasswordReset, request, options: options);
  }

  $grpc.ResponseFuture<$0.Auth> tokenLogin($0.RequestEmpty request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$tokenLogin, request, options: options);
  }

  $grpc.ResponseFuture<$0.ResponseAuthLogin> tokenRefresh($0.RequestEmpty request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$tokenRefresh, request, options: options);
  }

  $grpc.ResponseFuture<$0.ResponseTokenDecode> tokenDecode($0.RequestTokenDecode request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$tokenDecode, request, options: options);
  }

  $grpc.ResponseFuture<$0.JsonBytes> database($0.RequestDatabase request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$database, request, options: options);
  }

  $grpc.ResponseFuture<$0.ResponseStatus> delete($0.RequestDelete request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$delete, request, options: options);
  }

  $grpc.ResponseFuture<$0.JsonBytes> function($0.RequestFunction request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$function, request, options: options);
  }

  $grpc.ResponseFuture<$0.JsonBytes> view($0.RequestView request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$view, request, options: options);
  }

  $grpc.ResponseFuture<$0.Config> configUpdate($0.Config request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$configUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Config> configGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$configGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Configs> configQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$configQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Currency> currencyUpdate($0.Currency request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$currencyUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Currency> currencyGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$currencyGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Currencies> currencyQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$currencyQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Customer> customerUpdate($0.Customer request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$customerUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Customer> customerGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$customerGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Customers> customerQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$customerQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Employee> employeeUpdate($0.Employee request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$employeeUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Employee> employeeGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$employeeGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Employees> employeeQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$employeeQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Item> itemUpdate($0.Item request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$itemUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Item> itemGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$itemGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Items> itemQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$itemQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Link> linkUpdate($0.Link request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$linkUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Link> linkGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$linkGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Links> linkQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$linkQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Log> logGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$logGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Logs> logQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$logQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Movement> movementUpdate($0.Movement request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$movementUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Movement> movementGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$movementGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Movements> movementQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$movementQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Payment> paymentUpdate($0.Payment request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$paymentUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Payment> paymentGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$paymentGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Payments> paymentQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$paymentQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Place> placeUpdate($0.Place request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$placeUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Place> placeGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$placeGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Places> placeQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$placeQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Price> priceUpdate($0.Price request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$priceUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Price> priceGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$priceGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Prices> priceQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$priceQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Product> productUpdate($0.Product request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$productUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Product> productGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$productGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Products> productQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$productQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Project> projectUpdate($0.Project request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$projectUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Project> projectGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$projectGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Projects> projectQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$projectQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Rate> rateUpdate($0.Rate request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$rateUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Rate> rateGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$rateGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Rates> rateQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$rateQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Tax> taxUpdate($0.Tax request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$taxUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Tax> taxGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$taxGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Taxes> taxQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$taxQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Tool> toolUpdate($0.Tool request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toolUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Tool> toolGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toolGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Tools> toolQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$toolQuery, request, options: options);
  }

  $grpc.ResponseFuture<$0.Trans> transUpdate($0.Trans request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$transUpdate, request, options: options);
  }

  $grpc.ResponseFuture<$0.Trans> transGet($0.RequestGet request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$transGet, request, options: options);
  }

  $grpc.ResponseFuture<$0.Transactions> transQuery($0.RequestQuery request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$transQuery, request, options: options);
  }
}

@$pb.GrpcServiceName('nervatura.API')
abstract class APIServiceBase extends $grpc.Service {
  $core.String get $name => 'nervatura.API';

  APIServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.RequestAuthLogin, $0.ResponseAuthLogin>(
        'AuthLogin',
        authLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestAuthLogin.fromBuffer(value),
        ($0.ResponseAuthLogin value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Auth, $0.Auth>(
        'AuthUpdate',
        authUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Auth.fromBuffer(value),
        ($0.Auth value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Auth>(
        'AuthGet',
        authGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Auth value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestPasswordChange, $0.ResponseStatus>(
        'AuthPassword',
        authPassword_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestPasswordChange.fromBuffer(value),
        ($0.ResponseStatus value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.ResponseStatus>(
        'AuthPasswordReset',
        authPasswordReset_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.ResponseStatus value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestEmpty, $0.Auth>(
        'TokenLogin',
        tokenLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestEmpty.fromBuffer(value),
        ($0.Auth value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestEmpty, $0.ResponseAuthLogin>(
        'TokenRefresh',
        tokenRefresh_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestEmpty.fromBuffer(value),
        ($0.ResponseAuthLogin value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestTokenDecode, $0.ResponseTokenDecode>(
        'TokenDecode',
        tokenDecode_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestTokenDecode.fromBuffer(value),
        ($0.ResponseTokenDecode value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestDatabase, $0.JsonBytes>(
        'Database',
        database_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestDatabase.fromBuffer(value),
        ($0.JsonBytes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestDelete, $0.ResponseStatus>(
        'Delete',
        delete_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestDelete.fromBuffer(value),
        ($0.ResponseStatus value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestFunction, $0.JsonBytes>(
        'Function',
        function_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestFunction.fromBuffer(value),
        ($0.JsonBytes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestView, $0.JsonBytes>(
        'View',
        view_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestView.fromBuffer(value),
        ($0.JsonBytes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Config, $0.Config>(
        'ConfigUpdate',
        configUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Config.fromBuffer(value),
        ($0.Config value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Config>(
        'ConfigGet',
        configGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Config value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Configs>(
        'ConfigQuery',
        configQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Configs value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Currency, $0.Currency>(
        'CurrencyUpdate',
        currencyUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Currency.fromBuffer(value),
        ($0.Currency value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Currency>(
        'CurrencyGet',
        currencyGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Currency value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Currencies>(
        'CurrencyQuery',
        currencyQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Currencies value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Customer, $0.Customer>(
        'CustomerUpdate',
        customerUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Customer.fromBuffer(value),
        ($0.Customer value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Customer>(
        'CustomerGet',
        customerGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Customer value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Customers>(
        'CustomerQuery',
        customerQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Customers value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Employee, $0.Employee>(
        'EmployeeUpdate',
        employeeUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Employee.fromBuffer(value),
        ($0.Employee value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Employee>(
        'EmployeeGet',
        employeeGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Employee value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Employees>(
        'EmployeeQuery',
        employeeQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Employees value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Item, $0.Item>(
        'ItemUpdate',
        itemUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Item.fromBuffer(value),
        ($0.Item value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Item>(
        'ItemGet',
        itemGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Item value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Items>(
        'ItemQuery',
        itemQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Items value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Link, $0.Link>(
        'LinkUpdate',
        linkUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Link.fromBuffer(value),
        ($0.Link value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Link>(
        'LinkGet',
        linkGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Link value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Links>(
        'LinkQuery',
        linkQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Links value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Log>(
        'LogGet',
        logGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Log value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Logs>(
        'LogQuery',
        logQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Logs value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Movement, $0.Movement>(
        'MovementUpdate',
        movementUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Movement.fromBuffer(value),
        ($0.Movement value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Movement>(
        'MovementGet',
        movementGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Movement value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Movements>(
        'MovementQuery',
        movementQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Movements value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Payment, $0.Payment>(
        'PaymentUpdate',
        paymentUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Payment.fromBuffer(value),
        ($0.Payment value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Payment>(
        'PaymentGet',
        paymentGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Payment value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Payments>(
        'PaymentQuery',
        paymentQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Payments value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Place, $0.Place>(
        'PlaceUpdate',
        placeUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Place.fromBuffer(value),
        ($0.Place value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Place>(
        'PlaceGet',
        placeGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Place value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Places>(
        'PlaceQuery',
        placeQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Places value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Price, $0.Price>(
        'PriceUpdate',
        priceUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Price.fromBuffer(value),
        ($0.Price value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Price>(
        'PriceGet',
        priceGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Price value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Prices>(
        'PriceQuery',
        priceQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Prices value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Product, $0.Product>(
        'ProductUpdate',
        productUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Product.fromBuffer(value),
        ($0.Product value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Product>(
        'ProductGet',
        productGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Product value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Products>(
        'ProductQuery',
        productQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Products value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Project, $0.Project>(
        'ProjectUpdate',
        projectUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Project.fromBuffer(value),
        ($0.Project value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Project>(
        'ProjectGet',
        projectGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Project value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Projects>(
        'ProjectQuery',
        projectQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Projects value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Rate, $0.Rate>(
        'RateUpdate',
        rateUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Rate.fromBuffer(value),
        ($0.Rate value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Rate>(
        'RateGet',
        rateGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Rate value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Rates>(
        'RateQuery',
        rateQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Rates value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Tax, $0.Tax>(
        'TaxUpdate',
        taxUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Tax.fromBuffer(value),
        ($0.Tax value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Tax>(
        'TaxGet',
        taxGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Tax value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Taxes>(
        'TaxQuery',
        taxQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Taxes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Tool, $0.Tool>(
        'ToolUpdate',
        toolUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Tool.fromBuffer(value),
        ($0.Tool value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Tool>(
        'ToolGet',
        toolGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Tool value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Tools>(
        'ToolQuery',
        toolQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Tools value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Trans, $0.Trans>(
        'TransUpdate',
        transUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Trans.fromBuffer(value),
        ($0.Trans value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestGet, $0.Trans>(
        'TransGet',
        transGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestGet.fromBuffer(value),
        ($0.Trans value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RequestQuery, $0.Transactions>(
        'TransQuery',
        transQuery_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RequestQuery.fromBuffer(value),
        ($0.Transactions value) => value.writeToBuffer()));
  }

  $async.Future<$0.ResponseAuthLogin> authLogin_Pre($grpc.ServiceCall call, $async.Future<$0.RequestAuthLogin> request) async {
    return authLogin(call, await request);
  }

  $async.Future<$0.Auth> authUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Auth> request) async {
    return authUpdate(call, await request);
  }

  $async.Future<$0.Auth> authGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return authGet(call, await request);
  }

  $async.Future<$0.ResponseStatus> authPassword_Pre($grpc.ServiceCall call, $async.Future<$0.RequestPasswordChange> request) async {
    return authPassword(call, await request);
  }

  $async.Future<$0.ResponseStatus> authPasswordReset_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return authPasswordReset(call, await request);
  }

  $async.Future<$0.Auth> tokenLogin_Pre($grpc.ServiceCall call, $async.Future<$0.RequestEmpty> request) async {
    return tokenLogin(call, await request);
  }

  $async.Future<$0.ResponseAuthLogin> tokenRefresh_Pre($grpc.ServiceCall call, $async.Future<$0.RequestEmpty> request) async {
    return tokenRefresh(call, await request);
  }

  $async.Future<$0.ResponseTokenDecode> tokenDecode_Pre($grpc.ServiceCall call, $async.Future<$0.RequestTokenDecode> request) async {
    return tokenDecode(call, await request);
  }

  $async.Future<$0.JsonBytes> database_Pre($grpc.ServiceCall call, $async.Future<$0.RequestDatabase> request) async {
    return database(call, await request);
  }

  $async.Future<$0.ResponseStatus> delete_Pre($grpc.ServiceCall call, $async.Future<$0.RequestDelete> request) async {
    return delete(call, await request);
  }

  $async.Future<$0.JsonBytes> function_Pre($grpc.ServiceCall call, $async.Future<$0.RequestFunction> request) async {
    return function(call, await request);
  }

  $async.Future<$0.JsonBytes> view_Pre($grpc.ServiceCall call, $async.Future<$0.RequestView> request) async {
    return view(call, await request);
  }

  $async.Future<$0.Config> configUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Config> request) async {
    return configUpdate(call, await request);
  }

  $async.Future<$0.Config> configGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return configGet(call, await request);
  }

  $async.Future<$0.Configs> configQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return configQuery(call, await request);
  }

  $async.Future<$0.Currency> currencyUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Currency> request) async {
    return currencyUpdate(call, await request);
  }

  $async.Future<$0.Currency> currencyGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return currencyGet(call, await request);
  }

  $async.Future<$0.Currencies> currencyQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return currencyQuery(call, await request);
  }

  $async.Future<$0.Customer> customerUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Customer> request) async {
    return customerUpdate(call, await request);
  }

  $async.Future<$0.Customer> customerGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return customerGet(call, await request);
  }

  $async.Future<$0.Customers> customerQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return customerQuery(call, await request);
  }

  $async.Future<$0.Employee> employeeUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Employee> request) async {
    return employeeUpdate(call, await request);
  }

  $async.Future<$0.Employee> employeeGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return employeeGet(call, await request);
  }

  $async.Future<$0.Employees> employeeQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return employeeQuery(call, await request);
  }

  $async.Future<$0.Item> itemUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Item> request) async {
    return itemUpdate(call, await request);
  }

  $async.Future<$0.Item> itemGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return itemGet(call, await request);
  }

  $async.Future<$0.Items> itemQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return itemQuery(call, await request);
  }

  $async.Future<$0.Link> linkUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Link> request) async {
    return linkUpdate(call, await request);
  }

  $async.Future<$0.Link> linkGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return linkGet(call, await request);
  }

  $async.Future<$0.Links> linkQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return linkQuery(call, await request);
  }

  $async.Future<$0.Log> logGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return logGet(call, await request);
  }

  $async.Future<$0.Logs> logQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return logQuery(call, await request);
  }

  $async.Future<$0.Movement> movementUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Movement> request) async {
    return movementUpdate(call, await request);
  }

  $async.Future<$0.Movement> movementGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return movementGet(call, await request);
  }

  $async.Future<$0.Movements> movementQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return movementQuery(call, await request);
  }

  $async.Future<$0.Payment> paymentUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Payment> request) async {
    return paymentUpdate(call, await request);
  }

  $async.Future<$0.Payment> paymentGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return paymentGet(call, await request);
  }

  $async.Future<$0.Payments> paymentQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return paymentQuery(call, await request);
  }

  $async.Future<$0.Place> placeUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Place> request) async {
    return placeUpdate(call, await request);
  }

  $async.Future<$0.Place> placeGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return placeGet(call, await request);
  }

  $async.Future<$0.Places> placeQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return placeQuery(call, await request);
  }

  $async.Future<$0.Price> priceUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Price> request) async {
    return priceUpdate(call, await request);
  }

  $async.Future<$0.Price> priceGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return priceGet(call, await request);
  }

  $async.Future<$0.Prices> priceQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return priceQuery(call, await request);
  }

  $async.Future<$0.Product> productUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Product> request) async {
    return productUpdate(call, await request);
  }

  $async.Future<$0.Product> productGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return productGet(call, await request);
  }

  $async.Future<$0.Products> productQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return productQuery(call, await request);
  }

  $async.Future<$0.Project> projectUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Project> request) async {
    return projectUpdate(call, await request);
  }

  $async.Future<$0.Project> projectGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return projectGet(call, await request);
  }

  $async.Future<$0.Projects> projectQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return projectQuery(call, await request);
  }

  $async.Future<$0.Rate> rateUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Rate> request) async {
    return rateUpdate(call, await request);
  }

  $async.Future<$0.Rate> rateGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return rateGet(call, await request);
  }

  $async.Future<$0.Rates> rateQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return rateQuery(call, await request);
  }

  $async.Future<$0.Tax> taxUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Tax> request) async {
    return taxUpdate(call, await request);
  }

  $async.Future<$0.Tax> taxGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return taxGet(call, await request);
  }

  $async.Future<$0.Taxes> taxQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return taxQuery(call, await request);
  }

  $async.Future<$0.Tool> toolUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Tool> request) async {
    return toolUpdate(call, await request);
  }

  $async.Future<$0.Tool> toolGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return toolGet(call, await request);
  }

  $async.Future<$0.Tools> toolQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return toolQuery(call, await request);
  }

  $async.Future<$0.Trans> transUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Trans> request) async {
    return transUpdate(call, await request);
  }

  $async.Future<$0.Trans> transGet_Pre($grpc.ServiceCall call, $async.Future<$0.RequestGet> request) async {
    return transGet(call, await request);
  }

  $async.Future<$0.Transactions> transQuery_Pre($grpc.ServiceCall call, $async.Future<$0.RequestQuery> request) async {
    return transQuery(call, await request);
  }

  $async.Future<$0.ResponseAuthLogin> authLogin($grpc.ServiceCall call, $0.RequestAuthLogin request);
  $async.Future<$0.Auth> authUpdate($grpc.ServiceCall call, $0.Auth request);
  $async.Future<$0.Auth> authGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.ResponseStatus> authPassword($grpc.ServiceCall call, $0.RequestPasswordChange request);
  $async.Future<$0.ResponseStatus> authPasswordReset($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Auth> tokenLogin($grpc.ServiceCall call, $0.RequestEmpty request);
  $async.Future<$0.ResponseAuthLogin> tokenRefresh($grpc.ServiceCall call, $0.RequestEmpty request);
  $async.Future<$0.ResponseTokenDecode> tokenDecode($grpc.ServiceCall call, $0.RequestTokenDecode request);
  $async.Future<$0.JsonBytes> database($grpc.ServiceCall call, $0.RequestDatabase request);
  $async.Future<$0.ResponseStatus> delete($grpc.ServiceCall call, $0.RequestDelete request);
  $async.Future<$0.JsonBytes> function($grpc.ServiceCall call, $0.RequestFunction request);
  $async.Future<$0.JsonBytes> view($grpc.ServiceCall call, $0.RequestView request);
  $async.Future<$0.Config> configUpdate($grpc.ServiceCall call, $0.Config request);
  $async.Future<$0.Config> configGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Configs> configQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Currency> currencyUpdate($grpc.ServiceCall call, $0.Currency request);
  $async.Future<$0.Currency> currencyGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Currencies> currencyQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Customer> customerUpdate($grpc.ServiceCall call, $0.Customer request);
  $async.Future<$0.Customer> customerGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Customers> customerQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Employee> employeeUpdate($grpc.ServiceCall call, $0.Employee request);
  $async.Future<$0.Employee> employeeGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Employees> employeeQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Item> itemUpdate($grpc.ServiceCall call, $0.Item request);
  $async.Future<$0.Item> itemGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Items> itemQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Link> linkUpdate($grpc.ServiceCall call, $0.Link request);
  $async.Future<$0.Link> linkGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Links> linkQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Log> logGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Logs> logQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Movement> movementUpdate($grpc.ServiceCall call, $0.Movement request);
  $async.Future<$0.Movement> movementGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Movements> movementQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Payment> paymentUpdate($grpc.ServiceCall call, $0.Payment request);
  $async.Future<$0.Payment> paymentGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Payments> paymentQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Place> placeUpdate($grpc.ServiceCall call, $0.Place request);
  $async.Future<$0.Place> placeGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Places> placeQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Price> priceUpdate($grpc.ServiceCall call, $0.Price request);
  $async.Future<$0.Price> priceGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Prices> priceQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Product> productUpdate($grpc.ServiceCall call, $0.Product request);
  $async.Future<$0.Product> productGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Products> productQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Project> projectUpdate($grpc.ServiceCall call, $0.Project request);
  $async.Future<$0.Project> projectGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Projects> projectQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Rate> rateUpdate($grpc.ServiceCall call, $0.Rate request);
  $async.Future<$0.Rate> rateGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Rates> rateQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Tax> taxUpdate($grpc.ServiceCall call, $0.Tax request);
  $async.Future<$0.Tax> taxGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Taxes> taxQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Tool> toolUpdate($grpc.ServiceCall call, $0.Tool request);
  $async.Future<$0.Tool> toolGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Tools> toolQuery($grpc.ServiceCall call, $0.RequestQuery request);
  $async.Future<$0.Trans> transUpdate($grpc.ServiceCall call, $0.Trans request);
  $async.Future<$0.Trans> transGet($grpc.ServiceCall call, $0.RequestGet request);
  $async.Future<$0.Transactions> transQuery($grpc.ServiceCall call, $0.RequestQuery request);
}
