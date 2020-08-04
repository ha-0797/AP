const socket = io()

const delay = secs => new Promise(resolve => setTimeout(resolve, 1000*secs))

const bd = []

const make_bd = () =>{
	for (var i = 0; i < 10; i++) {
		bd.push([])
		for (var j = 0; j < 10; j++) {
			bd[i].push('')
		}
	}
}

const new_bd = (ship) =>{
	for (var i = 0; i < 10; i++) {
		for (var j = 0; j < 10; j++) {
			if(bd[i][j] == ('ship-' + ship)){
				bd[i][j] = ''
			} 
		}
	}
}
let ship = 'aircraft_carrier'
const shipsize = {
    'aircraft_carrier': 5,
    'battleship': 4,
    'cruiser': 3,
    'destroyer': 2,
    'submarine': 1
}
const state = {}

const clicked = (ev,col, row) =>{
	new_bd(ship)
	console.log(row, col)
	console.log(ship)
	console.log(shipsize[ship])
	let valid = true
	if(col <(11-shipsize[ship])){
		for(var i = 0; i < shipsize[ship]; i++){
			if(bd[row][col+i] != ''){
				valid = false
			}
		}
		if(valid){
			for (var i = 0; i < shipsize[ship]; i++) {
				bd[row][col+i] = 'ship-' + ship
			}
		}
	}
	setState({board: bd})
}

const right = (ev, col, row) =>{
	new_bd(ship)
	console.log(row, col)
	console.log(ship)
	console.log(shipsize[ship])
	let valid = true
	if(row < (11-shipsize[ship])){
		for(var i = 0; i < shipsize[ship]; i++){
			if(bd[row+i][col] != ''){
				valid = false
			}
		}
		console.log(valid)
		if(valid){
			for (var i = 0; i < shipsize[ship]; i++) {
				bd[row+i][col] = 'ship-' + ship
			}
		}
	}
	setState({board: bd})
}

const setState = updates => {
    Object.assign(state, updates)
    ReactDOM.render(
    	React.createElement('div', null,
    		React.createElement('div', null, state.msg),
    		[0,1,2,3,4,5,6,7,8,9].map(num2 =>
		    	React.createElement('div',{Id: num2},
		    		[0,1,2,3,4,5,6,7,8,9].map(num =>
		    			React.createElement('div', {
		    				className: 'box' + ' ' + state.board[num2][num], 
		    				onClick: (ev) => clicked(ev,num, num2),
		    				onContextMenu: ev => right(ev,num,num2)},
		    				num + (num2*10))))),
    		React.createElement('select', {onChange: ev => {ship = ev.target.value; console.log(ship)}}, 
    			React.createElement('option', null, 'aircraft_carrier'),
    			React.createElement('option', null, 'battleship'),
    			React.createElement('option', null, 'cruiser'),
    			React.createElement('option', null, 'destroyer'),
    			React.createElement('option', null, 'submarine'))), 
        document.getElementById('root'))
}
make_bd()
setState({msg: 'Hello World', board: bd})
