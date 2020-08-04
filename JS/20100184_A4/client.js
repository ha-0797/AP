const socket = io()

let player
let myTurn = false
let board = [[0,0,0,0,0,0],[0,0,0,0,0,0],[0,0,0,0,0,0],[0,0,0,0,0,0],[0,0,0,0,0,0],[0,0,0,0,0,0],[0,0,0,0,0,0]]
let opponent

function win_helper (a,b,c,d) {
	return ((a != 0) && (a ==b) && (a == c) && (a == d))
}

function win () {
	for (var i = 0; i < 4; i++) {
		for (var j = 0; j < 6; j++) {
			var x = win_helper(board[i][j],board[i+1][j],board[i+2][j],board[i+3][j])
			if(x){
				console.log('down is true')
				return(true)
			}
		}
	}
	for (var i = 0; i < 7; i++) {
		for (var j = 0; j < 3; j++) {
			var x = win_helper(board[i][j],board[i][j+1],board[i][j+2],board[i][j+3])
			if(x){			
				console.log('right is true')
				return(true)
			}
		}
	}
	for (var i = 0; i < 4; i++) {
		for (var j = 3; j < 6; j++) {
			var x = win_helper(board[i][j],board[i+1][j-1],board[i+2][j-2],board[i+3][j-3])
			if(x){
				console.log('down-right is true')
				return(true)
			}
		}
	}
	for (var i = 6; i >=3; i--) {
		for (var j = 5; j >=3; j--) {
			var x = win_helper(board[i][j],board[i-1][j-1],board[i-2][j-2],board[i-3][j-3])
			if(x){
				console.log('down-left is true')
				return(true)
			}
		}
	}
	return(false)
}

socket.on('player', data =>{
	console.log(`I am player ${data}`)
	player = data
})

socket.on('player2', sock =>{
	console.log(`match found`)
	opponent = sock
})

socket.on('turn',async data =>{
	if(await win()){
		if(player == data){
			console.log('lose')
			ReactDOM.render(
				React.createElement('div', null, "You lose"),
				document.getElementById('root'))
		} else {
			console.log('win')
			ReactDOM.render(
				React.createElement('div', null, "You win"),
				document.getElementById('root'))
		}
	}
	else if(player == data){
		ReactDOM.render(
				React.createElement('div', null, "Your turn"),
				document.getElementById('root'))
		turn()
	}
	else{
		ReactDOM.render(
				React.createElement('div', null, "Opp turn"),
				document.getElementById('root'))
	}
})

socket.on('p_move', data=>{
	console.log(data[0])
	let x = document.getElementsByClassName(data[0])
	let len = x.length
	console.log(x)
	console.log(x[len-1])
	console.log(len)	
	if(data[1] == 1){
		x[len-1].classList.add('p1')
		board[len-1][data[0]-1] = 1
	}
	else{
		x[len-1].classList.add('p2')	
		board[len-1][data[0]-1] = 2
	}
	x[len-1].classList.remove(data[0])
})

socket.on('w_move', () => {
	myTurn = true
})

const turn = () => {
	console.log('My turn')
	myTurn = true
}

function mousedown(data){
	if(myTurn){
		console.log("mousedown works")
		myTurn = false
		console.log(board)
		socket.emit('move', [data,board[0][data-1]])
	}
}

const mouseover = (data) => {
	if(myTurn){
		console.log("mouseover works")
		let x = document.getElementsByClassName(data)
		let len = x.length
		if(len){	
			board[len-1][data-1] = player
			let opp
			if(player == 1){
				opp = 2
			} else {
				opp = 1
			}
			document.getElementById(data).classList.add('other')
			document.getElementById(data).classList.remove('reg')
			for(var i = 1; i <= 6; i++){		
				var z =document.getElementsByClassName(i)
				var len2 = z.length
				if(x == z){
					--len2
				}
				if(len2){
					board[len2-1][i-1] = opp
				}
				var y = win()
				if(y){
					document.getElementById(data).classList.add('win')
					document.getElementById(data).classList.remove('other')
					board[len2-1][i-1] = 0
					break
				}
				if(len2){
					board[len2-1][i-1] = 0
				}
			}
			board[len-1][data-1] = 0
		}
	}
}

function mouseout(data){
	console.log("mouseout works")
	document.getElementById(data).classList.add('reg')
	document.getElementById(data).classList.remove('other')
	document.getElementById(data).classList.remove('win')
}
