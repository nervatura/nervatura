<?php

class Cli {
  function connect($arg){
    $command = array(getenv('NT_SERVICE_PATH'));
    if (getenv('NT_SERVICE_PATH') == 'docker') {
      $command = array('docker', 'exec', '-i', 'nervatura', '/nervatura');
    }
    $process = proc_open(
      array_merge($command, $arg), 
      array(0 => array('pipe', 'r'), 1 => array('pipe', 'w'), 2 => array('pipe', 'w')), 
      $pipes, __DIR__.'/..', getenv()
    );

    $result_str = ''; $error = '';
    if (is_resource($process)) {      
      fclose($pipes[0]);

      $result_str = stream_get_contents($pipes[1]);
      fclose($pipes[1]);

      $error = stream_get_contents($pipes[2]);
      fclose($pipes[2]);
      proc_close($process);
    }
    if ($error != '') {
      return array(
        'result'=>NULL,
        'error'=>$error
      );
    }
    $result_arr = explode(PHP_EOL, $result_str);
    $result = $result_arr[count($result_arr)-2];
    $result_map = json_decode($result, true);
    if ($result_map == NULL) {
      return array(
        'result'=>NULL,
        'error'=>$result
      );
    }
    if (is_array($result_map) && array_key_exists('code', $result_map) && 
      ($result_map['code'] != 200) && ($result_map['code'] != 201) && ($result_map['code'] != 204)) {
      return array(
        'result'=>NULL,
        'error'=>$result_map['data'] 
      );
    }
    return array(
      'result'=>$result_map,
      'error'=>NULL
    );
  }

  function withOptions($key, $model = NULL,  $options = NULL, $data = NULL){
    $prm = array('-c', $key);
    if ($model != NULL) {
      $prm = array_merge($prm, array('-m', $model));
    }
    if ($options != NULL) {
      $jopt = json_encode($options);
      if ($jopt == NULL) {
        return array(
          'result'=>NULL,
          'error'=>'Invalid data'
        );
      }
      $prm = array_merge($prm, array('-o', $jopt));
    }
    if ($data != NULL) {
      $jdata = json_encode($data);
      if ($jdata == NULL) {
        return array(
          'result'=>NULL,
          'error'=>'Invalid data'
        );
      }
      $prm = array_merge($prm, array('-d', $jdata));
    }
    return $this->connect($prm);
  }

  public function Database($options = array()){
    return $this->withOptions('database', NULL, $options, NULL);
  }

  public function ResetPassword($options = array()){
    return $this->withOptions('reset', NULL, $options, NULL);
  }

  public function Create($model, $options = array(), $data = array()){
    return $this->withOptions('create', $model, $options, $data);
  }

  public function Update($model, $options = array(), $data = array()){
    return $this->withOptions('update', $model, $options, $data);
  }

  public function Delete($model, $options = array()){
    return $this->withOptions('delete', $model, $options, NULL);
  }

  public function Get($model, $options = array()){
    return $this->withOptions('get', $model, $options, NULL);
  }

  public function Query($model, $options = array()){
    return $this->withOptions('query', $model, $options, NULL);
  }

  public function View($options = array()){
    return $this->withOptions('view', NULL, $options, NULL);
  }

  public function Function($options = array()){
    return $this->withOptions('function', NULL, $options, NULL);
  }
}

?>