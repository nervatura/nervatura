<?php

require(__DIR__.'/vendor/autoload.php');
require(__DIR__.'/lib/cli.php');
require(__DIR__.'/lib/rest.php');

$dotenv = Dotenv\Dotenv::createImmutable(__DIR__ . '/','.env.example');
$dotenv->load();

$api_map = array(
  'cli' => new Cli(),
  'http' => new Rest(),
);

$func_map = array(
  'cli' => array(
    'Database' => function($api){
      return $api->Database(array(
        'demo' => 'true',
        'alias' => 'demo'
      ));
    },
    'Function' => function($api){
      return $api->Function(array(
        'name' => 'test',
        'values' => array()
      ));
    },
    'ResetPassword' => function($api){
      return $api->ResetPassword(array(
        'alias' => 'demo',
        'code' => 'USR0000000000N1'
      ));
    },
    'Create' => function($api){
      return $api->Create('customer', array(
        'alias' => 'demo'
      ), array(
        'code' => 'CLI0000000000N1',
        'customer_name' => 'PHP Test'
      ));
    },
    'Update' => function($api){
      return $api->Update('customer', array(
        'alias' => 'demo',
        'code' => 'CLI0000000000N1',
      ), array(
        'customer_meta' => array(
          'account' => '1234567890'
        )
      ));
    },
    'Get' => function($api){
      return $api->Get('customer', array(
        'alias' => 'demo',
        'code' => 'CLI0000000000N1',
      ));
    },
    'Query' => function($api){
      return $api->Query('customer', array(
        'alias' => 'demo',
        'customer_type' => 'CUSTOMER_COMPANY',
      ));
    },
    'View' => function($api){
      return $api->View(array(
        'alias' => 'demo',
        'name' => 'VIEW_CUSTOMER_EVENTS',
        'filter' => "subject like '%visit%' and place='City1'",
        'limit' => 10,
      ));
    },
    'Delete' => function($api){
      return $api->Delete('customer', array(
        'alias' => 'demo',
        'code' => 'CLI0000000000N1',
      ));
    },
  ),
  'http' => array(
    'Post' => function($api){
      return $api->Post('', 'customer', array(
        'code' => 'REST0000000000N1',
        'customer_name' => 'PHP Test'
      ));
    },
    'Put' => function($api){
      return $api->Put('', 'customer/REST0000000000N1', array(
        'customer_name' => 'Test PHP'
      ));
    },
    'Get' => function($api){
      return $api->Get('', 'customer', array(
        'customer_type' => 'CUSTOMER_COMPANY'
      ));
    },
    'Delete' => function($api){
      return $api->Delete('', 'customer/REST0000000000N1');
    },
  )
);

foreach ($api_map as $api_key => $api) {
  $start_time = microtime(true);
  foreach ($func_map[$api_key] as $func_name => $func_test) {
    $result = $func_test($api);
    if($result['error'] != NULL) {
      echo $api_key.' '.$func_name.' error: '.$result['error'].PHP_EOL;
      continue;
    }
    echo $api_key.' '.$func_name.' OK'.PHP_EOL;
  }
  $end_time = microtime(true);
  echo '--------------------'.PHP_EOL;
  echo $api_key.' time: '.(int)(($end_time-$start_time)*1000).PHP_EOL;
  echo '--------------------'.PHP_EOL;
}

?>