<?php
/******************************************************************************
 * Icinga 2 API Example: PHP API client with problems                         *
 * Copyright (C) 2016 Icinga Development Team (https://www.icinga.org)        *
 *                                                                            *
 * This program is free software; you can redistribute it and/or              *
 * modify it under the terms of the GNU General Public License                *
 * as published by the Free Software Foundation; either version 2             *
 * of the License, or (at your option) any later version.                     *
 *                                                                            *
 * This program is distributed in the hope that it will be useful,            *
 * but WITHOUT ANY WARRANTY; without even the implied warranty of             *
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the              *
 * GNU General Public License for more details.                               *
 *                                                                            *
 * You should have received a copy of the GNU General Public License          *
 * along with this program; if not, write to the Free Software Foundation     *
 * Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301, USA.             *
 ******************************************************************************/
# inspired by Icinga Director, ./library/Director/Core/RestApiClient.php

class ApiClient {
	// http://php.net/manual/de/function.json-last-error.php#119985
	protected $errorReference = [
		JSON_ERROR_NONE => 'No error has occurred.',
		JSON_ERROR_DEPTH => 'The maximum stack depth has been exceeded.',
		JSON_ERROR_STATE_MISMATCH => 'Invalid or malformed JSON.',
		JSON_ERROR_CTRL_CHAR => 'Control character error, possibly incorrectly encoded.',
		JSON_ERROR_SYNTAX => 'Syntax error.',
		JSON_ERROR_UTF8 => 'Malformed UTF-8 characters, possibly incorrectly encoded.',
		// These last 3 messages are only supported on PHP >= 5.5.
		// See http://php.net/json_last_error#refsect1-function.json-last-error-returnvalues
		JSON_ERROR_RECURSION => 'One or more recursive references in the value to be encoded.',
		JSON_ERROR_INF_OR_NAN => 'One or more NAN or INF values in the value to be encoded.',
		JSON_ERROR_UNSUPPORTED_TYPE => 'A value of a type that cannot be encoded was given.',
	];

	const JSON_UNKNOWN_ERROR = 'Unknown error.';

	protected $version = 'v1';
	protected $peer;
	protected $port;
	protected $user;
	protected $pass;
	protected $curl;

	public function __construct($peer, $port = 5665) {
		$this->peer = $peer;
		$this->port = $port;
	}

	public function setCredentials($user, $pass) {
		$this->user = $user;
		$this->pass = $pass;

		return this;
	}

	protected function url($url) {
		return sprintf('https://%s:%d/%s/%s', $this->peer, $this->port, $this->version, $url);
	}

	protected function curl() {
		if ($this->curl === null) {
			$this->curl = curl_init(sprintf('https://%s:%d', $this->peer, $this->port));
			if (!$this->curl) {
				throw new Exception('CURL INIT ERROR: ' . curl_error($this->curl));
			}
		}

		return $this->curl;
	}

	public function request($method, $url, $headers, $body) {
		$auth = sprintf('%s:%s', $this->user, $this->pass);

		$curlHeaders[] = "Accept: application/json";

		if ($body !== null) {
			$body = json_encode($body);
			$curlHeaders[] = 'Content-Type: application/json';
		}

		if ($headers !== null) {
			$curlFinalHeaders = array_merge($curlHeaders, $headers);
		}

		$curl = $this->curl();
		$opts = array(
			CURLOPT_URL		=> $this->url($url),
			CURLOPT_HTTPHEADER 	=> $curlFinalHeaders,
			CURLOPT_USERPWD		=> $auth,
			CURLOPT_CUSTOMREQUEST	=> strtoupper($method),
			CURLOPT_RETURNTRANSFER 	=> true,
			CURLOPT_CONNECTTIMEOUT 	=> 10,
			//TODO: fix it
			CURLOPT_SSL_VERIFYHOST 	=> false,
			CURLOPT_SSL_VERIFYPEER 	=> false,
		);

		if ($body !== null) {
			$opts[CURLOPT_POSTFIELDS] = $body;
		}

		curl_setopt_array($curl, $opts);

		$res = curl_exec($curl);

		if ($res === false) {
			throw new Exception('CURL ERROR: ' . curl_error($curl));
		}

		$statusCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);

		if ($statusCode === 401) {
			throw new Exception('Unable to authenticate, please check your API credentials');
		}

		return $this->fromJsonResult($res);
	}

	public function fromJsonResult($json) {
		$result = @json_decode($json);

		var_dump($json);

		if ($result === null) {
			throw new Exception('Parsing JSON failed: '.$this->getLastJsonErrorMessage(json_last_error()));
		}

		return $result->results;
	}

	public function getLastJsonErrorMessage($errorCode) {
		if (!array_key_exists($errorCode, $this->errorReference)) {
			return self::JSON_UNKNOWN_ERROR;
		}

		return $this->errorReference[$errorCode];
	}
}

$client = new ApiClient('localhost');
$client->setCredentials('root', 'icinga');
$getHeader = array('X-HTTP-Method-Override: GET');
$body = array(
	'joins' => array(
		'host'
	),
	//'filter' => 'service.state!=service_state && match("networks",host.groups) && host.last_hard_state!=1 && host.last_check!=-1 && host.acknowledgement!=host_ack',
	'filter' => 'service.state!=service_state && host.acknowledgement!=host_ack',
	'filter_vars' => array(
		'service_state' => 'ServiceOK',
		'host_ack' 	=> 2
	)
);

$result = $client->request("post", "objects/services", $getHeader, $body);

var_dump($result);


