require 'webrick'

port = ENV['PORT'] || 3000
server = WEBrick::HTTPServer.new(Port: port)

server.mount_proc '/' do |req, res|
  res.status = 200
  res.content_type = 'text/html'
  res.body = 'Welcome, You are at DevOps Solutions!'
end

server.mount_proc '/health' do |req, res|
  res.status = 204
  res.body = ''
end

server.mount_proc '/ping' do |req, res|
  puts 'In ping'
  res.status = 200
  res.content_type = 'text/html'
  res.body = 'pong'
end
server.start