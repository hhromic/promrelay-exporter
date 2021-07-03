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

const http = require('http'),
      httpProxy = require('http-proxy'),
      url = require('url');

const proxy = httpProxy.createProxyServer({});

const port = process.env.PORT || 8080;

proxy.on('error', function (err, req, res) {
  res.writeHead(502, {
    'Content-Type': 'text/plain'
  });
  res.end(err.toString() + '\n');
});

const server = http.createServer(function(req, res) {
  const query = url.parse(req.url, true).query;
  const scheme = query._scheme || 'http';
  const host = query._host || 'localhost';
  const port = query._port || 80;
  proxy.web(req, res, {target: `${scheme}://${host}:${port}`});
});

console.log('Prometheus relay listening on port %d', port);
server.listen(port);
