<?php

use Symfony\Component\HttpClient\HttpClient;

class Rest {
  function request($method, $path, $token, $query = array(), $data = array()){
    $service_url = 'http://localhost:'.getenv('NT_HTTP_PORT').'/api/v6/'.$path;
    $header = array(
      'Content-Type: application/json',
      'Accept: application/json'
    );
    if ($token != '') {
      $header = array_merge($header, array('Authorization: Bearer '.$token));
    }
    if ($token == '') {
      $header = array_merge($header, array('X-API-Key: '.getenv('NT_API_KEY')));
    }

    $client = HttpClient::create();
    $response = $client->request($method, $service_url, [
      'headers' => $header,
      'query' => $query,
      'body' => json_encode($data, JSON_UNESCAPED_UNICODE)
    ]);

    $status = $response->getStatusCode();
    if($status != 200 && $status != 201 && $status != 204) {
      return array(
        'result' => NULL,
        'error' => $response->getContent(false)
      );
    }
    $headers = $response->getHeaders();
    if(array_key_exists('content-type', $headers)) {
      $contentType = $headers['content-type'][0];
      if($contentType == 'application/pdf') {
        return array(
          'result' => $response->getContent(false),
          'error' => NULL
        );
      }
    }
    
    $body = $response->getContent();
    $response = json_decode($body, true);
    return array(
      'result' => $response,
      'error' => NULL
    );
  }

  function error($f3) {
    return array(
      'result' => NULL,
      'error' => $f3->get('ERROR.text')
    );
  }

  public function Get($token='', $path='', $query = array()){
    return $this->request('GET', $path, $token, $query);
  }

  public function Post($token='', $path='', $data = NULL){
    return $this->request('POST', $path, $token, array(), $data);
  }

  public function Put($token='', $path='', $data = NULL){
    return $this->request('PUT', $path, $token, array(), $data);
  }

  public function Delete($token='', $path=''){
    return $this->request('DELETE', $path, $token, array());
  }
}

?>