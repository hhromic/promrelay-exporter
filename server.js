//   Copyright 2021 Hugo Hromic
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

'use strict';

const http = require('http'),
      httpProxy = require('http-proxy');

const proxy = httpProxy.createProxyServer({changeOrigin: true});

const port = process.env.PORT || 8080;

proxy.on('error', function (err, req, res) {
  res.writeHead(502, {'Content-Type': 'text/plain'});
  res.end(err.toString() + '\n');
});

proxy.on('proxyReq', function(proxyReq) {
  const proxyReqURL = new URL(proxyReq.path, 'http://localhost');
  proxyReqURL.searchParams.delete('target');
  proxyReq.path = `${proxyReqURL.pathname}${proxyReqURL.search}`;
});

const server = http.createServer(function(req, res) {
  const reqURL = new URL(req.url, 'http://localhost');
  let target = reqURL.searchParams.get('target');
  if (target === null) {
    res.writeHead(400, {'Content-Type': 'text/plain'});
    res.end('Query search parameter \'target\' is missing.\n')
    return;
  }
  if (!target.startsWith("http://") && !target.startsWith("https://")) {
    target = `http://${target}`;
  }
  proxy.web(req, res, {target: target});
});

console.log('Prometheus relay exporter listening on port %d', port);
server.listen(port);
