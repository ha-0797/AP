const fs = require('fs')
const http = require('http')
const socketio = require('socket.io')
let connectCount = 0
const list1 = []
const list2 = []

const readFile = file => new Promise((resolve, reject) =>
    fs.readFile(file, 'utf8', (err, data) => err ? reject(err) : resolve(data)))

const delay = msecs => new Promise(resolve => setTimeout(resolve, msecs))

const server = http.createServer(async (request, response) =>
    response.end(await readFile(request.url.substr(1))))

const io = socketio(server)

io.sockets.on('connection', socket => {
    console.log('a client connected')
   	let game
   	let player
    socket.on('disconnect', () => console.log('a client disconnected'))
    socket.on('begin', data =>{
    	console.log('begin')
    	data[1]['']=85
    	for(var i = 0; i < 10; i ++){
    		for(var j = 0; j < 10; j++){
    			let word = data[0][i][j].substr(5)
    			--data[1][word]
    		}
    	}
    	let okay = true
    	Object.keys(data[1]).forEach(key =>{
    		console.log(data[1][key])
    		if(data[1][key] != 0){
    			socket.emit('begin', false)
    			okay = false
    		}
    	})
    	console.log(okay)
  		if(okay){
  			++connectCount
  			if(connectCount %2 == 1){
  				list2.push(socket)
  			} else {
  				list1.push(socket)
  			}
  			console.log(connectCount)
  			game = list2.indexOf(socket)
  			player = 1
  			if(game = -1){
  				console.log('works')
  				game = list1.indexOf(socket)
  				player = 2
  			}
  			if(player = 1){
  				socket.emit('wait', null)
  			} else {
  				console.log('works game')
  				socket.emit('begin2', 2)
  				list2[game].emit('begin2', 1)
  			}
  		} else{
  			socket.emit('begin2', 0)
  		}
    })
})

server.listen(8000)
