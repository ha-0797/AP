const fs = require('fs')
const http = require('http')
const socketio = require('socket.io')
let connectCounter = 0
const list = []
const list2 = []

const readFile = f => new Promise((resolve, reject) =>
	fs.readFile(f, (e,d) => e?reject(e):resolve(d)))

const server = http.createServer(async (req, resp) =>
	resp.end(await readFile(req.url.substr(1))))

const io = socketio(server)

io.sockets.on('connection', socket => {
	console.log("user connected")
	if(++connectCounter%2){
		list.push(socket)
		socket.emit('player', 1)
	} else {
		list2.push(socket)
		game = list2.indexOf(socket)
		socket.emit('player', 2)
		list[game].emit('turn', 1)
	}
	socket.on('move', (data) =>{
		console.log('got move')
		console.log(data[1])
		if(data[1] != 0){
			socket.emit('w_move')
		} else {
			let game = list.indexOf(socket)
			let turn = true
			if(game == -1){
				game = list2.indexOf(socket)
				turn = false
			}
			if(turn){
				socket.emit('p_move', [data[0], 1])
				list2[game].emit('p_move', [data[0], 1])	
				socket.emit('turn', 2)
				list2[game].emit('turn', 2)
			}
			else{
				socket.emit('p_move', [data[0], 2])
				list[game].emit('p_move', [data[0], 2])
				socket.emit('turn', 1)
				list[game].emit('turn', 1)
			}
		}	
	})
})

server.listen(8000, () => console.log('Started...'))

io.sockets.on('yourmove', data => console.log(data))