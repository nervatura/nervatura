<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Nervatura;

/**
 *
 * Nervatura gRPC API
 */
class APIClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * Login by username and password and get access token
     * @param \Nervatura\RequestAuthLogin $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AuthLogin(\Nervatura\RequestAuthLogin $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/AuthLogin',
        $argument,
        ['\Nervatura\ResponseAuthLogin', 'decode'],
        $metadata, $options);
    }

    /**
     * Update or create user account. 
     * If id or existing code is set, the user is updated, otherwise a new user is created. 
     * If user is not admin, the user id and code are set to the current token user id or code
     * @param \Nervatura\Auth $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AuthUpdate(\Nervatura\Auth $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/AuthUpdate',
        $argument,
        ['\Nervatura\Auth', 'decode'],
        $metadata, $options);
    }

    /**
     * Get user account by database id or code (unique key) 
     * If user is not admin, the user id and code are set to the current token user id or code
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AuthGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/AuthGet',
        $argument,
        ['\Nervatura\Auth', 'decode'],
        $metadata, $options);
    }

    /**
     * Change token user password
     * @param \Nervatura\RequestPasswordChange $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AuthPassword(\Nervatura\RequestPasswordChange $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/AuthPassword',
        $argument,
        ['\Nervatura\ResponseStatus', 'decode'],
        $metadata, $options);
    }

    /**
     * Reset a user password and result a new password 
     * If user is not admin, the user id and code are set to the current token user id or code
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AuthPasswordReset(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/AuthPasswordReset',
        $argument,
        ['\Nervatura\ResponseStatus', 'decode'],
        $metadata, $options);
    }

    /**
     * Login by access token. Returns token user
     * @param \Nervatura\RequestEmpty $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TokenLogin(\Nervatura\RequestEmpty $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/TokenLogin',
        $argument,
        ['\Nervatura\Auth', 'decode'],
        $metadata, $options);
    }

    /**
     * Refresh access token
     * @param \Nervatura\RequestEmpty $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TokenRefresh(\Nervatura\RequestEmpty $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/TokenRefresh',
        $argument,
        ['\Nervatura\ResponseAuthLogin', 'decode'],
        $metadata, $options);
    }

    /**
     * Decode JWT access token but doesn't validate the signature
     * @param \Nervatura\RequestTokenDecode $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TokenDecode(\Nervatura\RequestTokenDecode $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/TokenDecode',
        $argument,
        ['\Nervatura\ResponseTokenDecode', 'decode'],
        $metadata, $options);
    }

    /**
     * Create new nervatura database schema
     * @param \Nervatura\RequestDatabase $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Database(\Nervatura\RequestDatabase $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/Database',
        $argument,
        ['\Nervatura\JsonBytes', 'decode'],
        $metadata, $options);
    }

    /**
     * Delete data by database id or code (unique key)
     * Only admin users can delete configurations except for CONFIG_PRINT_QUEUE and CONFIG_PATTERN
     * Auth user cannot be deleted, only disabled by admin users
     * @param \Nervatura\RequestDelete $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Delete(\Nervatura\RequestDelete $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/Delete',
        $argument,
        ['\Nervatura\ResponseStatus', 'decode'],
        $metadata, $options);
    }

    /**
     * Call a server side function and get the result 
     * Example: create new PDF invoice, send email or get a product price
     * @param \Nervatura\RequestFunction $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Function(\Nervatura\RequestFunction $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/Function',
        $argument,
        ['\Nervatura\JsonBytes', 'decode'],
        $metadata, $options);
    }

    /**
     * Get a predefined view by name
     * @param \Nervatura\RequestView $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function View(\Nervatura\RequestView $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/View',
        $argument,
        ['\Nervatura\JsonBytes', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create configuration 
     * If id or existing code is set, the configuration is updated, otherwise a new configuration is created.
     * Only admin users can create or update configurations except for CONFIG_PRINT_QUEUE and CONFIG_PATTERN
     * @param \Nervatura\Config $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ConfigUpdate(\Nervatura\Config $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ConfigUpdate',
        $argument,
        ['\Nervatura\Config', 'decode'],
        $metadata, $options);
    }

    /**
     * Get configuration by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ConfigGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ConfigGet',
        $argument,
        ['\Nervatura\Config', 'decode'],
        $metadata, $options);
    }

    /**
     * Get configurations by query parameters
     * Valid filters: config_type
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ConfigQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ConfigQuery',
        $argument,
        ['\Nervatura\Configs', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create currency 
     * If id or existing code is set, the currency is updated, otherwise a new currency is created.
     * @param \Nervatura\Currency $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CurrencyUpdate(\Nervatura\Currency $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/CurrencyUpdate',
        $argument,
        ['\Nervatura\Currency', 'decode'],
        $metadata, $options);
    }

    /**
     * Get currency by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CurrencyGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/CurrencyGet',
        $argument,
        ['\Nervatura\Currency', 'decode'],
        $metadata, $options);
    }

    /**
     * Get currencies by query parameters
     * Valid filters: tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CurrencyQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/CurrencyQuery',
        $argument,
        ['\Nervatura\Currencies', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create customer 
     * If id or existing code is set, the customer is updated, otherwise a new customer is created.
     * @param \Nervatura\Customer $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CustomerUpdate(\Nervatura\Customer $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/CustomerUpdate',
        $argument,
        ['\Nervatura\Customer', 'decode'],
        $metadata, $options);
    }

    /**
     * Get customer by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CustomerGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/CustomerGet',
        $argument,
        ['\Nervatura\Customer', 'decode'],
        $metadata, $options);
    }

    /**
     * Get customers by query parameters
     * Valid filters: customer_type, customer_name, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CustomerQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/CustomerQuery',
        $argument,
        ['\Nervatura\Customers', 'decode'],
        $metadata, $options);
    }

    /**
     *
     * Update (replace) or create employee
     * If id or existing code is set, the employee is updated, otherwise a new employee is created.
     * @param \Nervatura\Employee $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function EmployeeUpdate(\Nervatura\Employee $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/EmployeeUpdate',
        $argument,
        ['\Nervatura\Employee', 'decode'],
        $metadata, $options);
    }

    /**
     * Get employee by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function EmployeeGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/EmployeeGet',
        $argument,
        ['\Nervatura\Employee', 'decode'],
        $metadata, $options);
    }

    /**
     * Get employees by query parameters
     * Valid filters: tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function EmployeeQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/EmployeeQuery',
        $argument,
        ['\Nervatura\Employees', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create item 
     * If id or existing code is set, the item is updated, otherwise a new item is created.
     * @param \Nervatura\Item $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ItemUpdate(\Nervatura\Item $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ItemUpdate',
        $argument,
        ['\Nervatura\Item', 'decode'],
        $metadata, $options);
    }

    /**
     * Get item by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ItemGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ItemGet',
        $argument,
        ['\Nervatura\Item', 'decode'],
        $metadata, $options);
    }

    /**
     * Get items by query parameters 
     * Valid filters: trans_code, product_code, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ItemQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ItemQuery',
        $argument,
        ['\Nervatura\Items', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create link 
     * If id or existing code is set, the link is updated, otherwise a new link is created.
     * @param \Nervatura\Link $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function LinkUpdate(\Nervatura\Link $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/LinkUpdate',
        $argument,
        ['\Nervatura\Link', 'decode'],
        $metadata, $options);
    }

    /**
     * Get link by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function LinkGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/LinkGet',
        $argument,
        ['\Nervatura\Link', 'decode'],
        $metadata, $options);
    }

    /**
     * Get links by query parameters 
     * Valid filters: link_type_1, link_code_1, link_type_2, link_code_2, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function LinkQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/LinkQuery',
        $argument,
        ['\Nervatura\Links', 'decode'],
        $metadata, $options);
    }

    /**
     * Get log by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function LogGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/LogGet',
        $argument,
        ['\Nervatura\Log', 'decode'],
        $metadata, $options);
    }

    /**
     * Get logs by query parameters 
     * Valid filters: log_type, ref_type, ref_code, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function LogQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/LogQuery',
        $argument,
        ['\Nervatura\Logs', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create movement 
     * If id or existing code is set, the movement is updated, otherwise a new movement is created.
     * @param \Nervatura\Movement $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function MovementUpdate(\Nervatura\Movement $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/MovementUpdate',
        $argument,
        ['\Nervatura\Movement', 'decode'],
        $metadata, $options);
    }

    /**
     * Get movement by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function MovementGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/MovementGet',
        $argument,
        ['\Nervatura\Movement', 'decode'],
        $metadata, $options);
    }

    /**
     * Get movements by query parameters 
     * Valid filters: trans_code, movement_type, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function MovementQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/MovementQuery',
        $argument,
        ['\Nervatura\Movements', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create payment 
     * If id or existing code is set, the payment is updated, otherwise a new payment is created.
     * @param \Nervatura\Payment $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function PaymentUpdate(\Nervatura\Payment $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/PaymentUpdate',
        $argument,
        ['\Nervatura\Payment', 'decode'],
        $metadata, $options);
    }

    /**
     * Get payment by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function PaymentGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/PaymentGet',
        $argument,
        ['\Nervatura\Payment', 'decode'],
        $metadata, $options);
    }

    /**
     * Get payments by query parameters   
     * Valid filters: trans_code, paid_date, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function PaymentQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/PaymentQuery',
        $argument,
        ['\Nervatura\Payments', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create place 
     * If id or existing code is set, the place is updated, otherwise a new place is created.
     * @param \Nervatura\Place $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function PlaceUpdate(\Nervatura\Place $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/PlaceUpdate',
        $argument,
        ['\Nervatura\Place', 'decode'],
        $metadata, $options);
    }

    /**
     * Get place by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function PlaceGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/PlaceGet',
        $argument,
        ['\Nervatura\Place', 'decode'],
        $metadata, $options);
    }

    /**
     * Get places by query parameters   
     * Valid filters: place_type, place_name, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function PlaceQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/PlaceQuery',
        $argument,
        ['\Nervatura\Places', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create price 
     * If id or existing code is set, the price is updated, otherwise a new price is created.
     * @param \Nervatura\Price $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function PriceUpdate(\Nervatura\Price $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/PriceUpdate',
        $argument,
        ['\Nervatura\Price', 'decode'],
        $metadata, $options);
    }

    /**
     * Get price by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function PriceGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/PriceGet',
        $argument,
        ['\Nervatura\Price', 'decode'],
        $metadata, $options);
    }

    /**
     * Get prices by query parameters   
     * Valid filters: price_type, product_code, currency_code, customer_code, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function PriceQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/PriceQuery',
        $argument,
        ['\Nervatura\Prices', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create product 
     * If id or existing code is set, the product is updated, otherwise a new product is created.
     * @param \Nervatura\Product $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ProductUpdate(\Nervatura\Product $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ProductUpdate',
        $argument,
        ['\Nervatura\Product', 'decode'],
        $metadata, $options);
    }

    /**
     * Get product by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ProductGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ProductGet',
        $argument,
        ['\Nervatura\Product', 'decode'],
        $metadata, $options);
    }

    /**
     * Get products by query parameters   
     * Valid filters: product_type, product_name, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ProductQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ProductQuery',
        $argument,
        ['\Nervatura\Products', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create project 
     * If id or existing code is set, the project is updated, otherwise a new project is created.
     * @param \Nervatura\Project $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ProjectUpdate(\Nervatura\Project $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ProjectUpdate',
        $argument,
        ['\Nervatura\Project', 'decode'],
        $metadata, $options);
    }

    /**
     * Get project by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ProjectGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ProjectGet',
        $argument,
        ['\Nervatura\Project', 'decode'],
        $metadata, $options);
    }

    /**
     * Get projects by query parameters     
     * Valid filters: project_name, customer_code, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ProjectQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ProjectQuery',
        $argument,
        ['\Nervatura\Projects', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create rate 
     * If id or existing code is set, the rate is updated, otherwise a new rate is created.
     * @param \Nervatura\Rate $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RateUpdate(\Nervatura\Rate $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/RateUpdate',
        $argument,
        ['\Nervatura\Rate', 'decode'],
        $metadata, $options);
    }

    /**
     * Get rate by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RateGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/RateGet',
        $argument,
        ['\Nervatura\Rate', 'decode'],
        $metadata, $options);
    }

    /**
     * Get rates by query parameters     
     * Valid filters: rate_type, currency_code, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RateQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/RateQuery',
        $argument,
        ['\Nervatura\Rates', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create tax 
     * If id or existing code is set, the tax is updated, otherwise a new tax is created.
     * @param \Nervatura\Tax $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TaxUpdate(\Nervatura\Tax $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/TaxUpdate',
        $argument,
        ['\Nervatura\Tax', 'decode'],
        $metadata, $options);
    }

    /**
     * Get tax by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TaxGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/TaxGet',
        $argument,
        ['\Nervatura\Tax', 'decode'],
        $metadata, $options);
    }

    /**
     * Get taxes by query parameters     
     * Valid filters: tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TaxQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/TaxQuery',
        $argument,
        ['\Nervatura\Taxes', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create tool
     * If id or existing code is set, the tool is updated, otherwise a new tool is created.
     * @param \Nervatura\Tool $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ToolUpdate(\Nervatura\Tool $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ToolUpdate',
        $argument,
        ['\Nervatura\Tool', 'decode'],
        $metadata, $options);
    }

    /**
     * Get tool by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ToolGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ToolGet',
        $argument,
        ['\Nervatura\Tool', 'decode'],
        $metadata, $options);
    }

    /**
     * Get tools by query parameters     
     * Valid filters: product_code, description, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ToolQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/ToolQuery',
        $argument,
        ['\Nervatura\Tools', 'decode'],
        $metadata, $options);
    }

    /**
     * Update (replace) or create transaction 
     * If id or existing code is set, the transaction is updated, otherwise a new transaction is created.
     * @param \Nervatura\Trans $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TransUpdate(\Nervatura\Trans $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/TransUpdate',
        $argument,
        ['\Nervatura\Trans', 'decode'],
        $metadata, $options);
    }

    /**
     * Get transaction by database id or code (unique key)
     * @param \Nervatura\RequestGet $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TransGet(\Nervatura\RequestGet $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/TransGet',
        $argument,
        ['\Nervatura\Trans', 'decode'],
        $metadata, $options);
    }

    /**
     * Get transactions by query parameters     
     * Valid filters: trans_type, direction, trans_date, tag
     * @param \Nervatura\RequestQuery $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TransQuery(\Nervatura\RequestQuery $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/nervatura.API/TransQuery',
        $argument,
        ['\Nervatura\Transactions', 'decode'],
        $metadata, $options);
    }

}
