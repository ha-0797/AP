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

const new_bd = () =>{
	for (var i = 0; i < 10; i++) {
		for (var j = 0; j < 10; j++) {
			bd[i][j] = ''
		}
	}
}
const shipsize = {
    'aircraft_carrier': 5,
    'battleship': 4,
    'cruiser': 3,
    'destroyer': 2,
    'submarine': 1
}
const state = {}

const makeShip = (col, row) => {
	
}

const clicked = (ev,col, row) =>{
	new_bd()
	console.log(bd)
	console.log(row, col)
	if(col <6){
		[0,1,2,3,4].map(num =>
			bd[row][col+num] = 'ship-aircraft_carrier')
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
		    			React.createElement('div', {className: 'box' + ' ' + state.board[num2][num], onClick: (ev) => clicked(ev,num, num2)},num + (num2*10))))),
    		React.createElement('select', null, 
    			React.createElement('option', null, 'aircraft_carrier'),
    			React.createElement('option', null, 'battleship'),
    			React.createElement('option', null, 'cruiser'),
    			React.createElement('option', null, 'destroyer'),
    			React.createElement('option', null, 'submarine'))), 
        document.getElementById('root'))
}
make_bd()
setState({msg: 'Hello World', board: bd})
