const http = require('http');

const server = http.createServer((req, res) => {
  res.end('Hello from API service!');
});

server.listen(3000, () => {
  console.log('API server running on port 3000');
});
