const socket = io()

const delay = secs => new Promise(resolve => setTimeout(resolve, 1000*secs))

const shipsize = {
    'aircraft_carrier': 5,
    'battleship': 4,
    'cruiser': 3,
    'destroyer': 2,
    'submarine': 1
}
const state = {}

const setState = updates => {
    Object.assign(state, updates)
    ReactDOM.render( 
    	
    	React.createElement('div',null,
    		[0,1,2,3,4,5,6,7,8,9].map(num =>
    			React.createElement('div', {className: 'box'}))), 
        document.getElementById('root'))
}

setState({msg: 'Hello World'})
