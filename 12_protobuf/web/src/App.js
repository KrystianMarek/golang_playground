import logo from './logo.svg';
import './App.css';
import React from 'react';

// let generated = require('./generated/echo_pb.js') // ToDo: protobuf looks broken for js...
// let request = new proto.EchoRequest()
// request.setName("React!")
//
// console.log(request)

function Square(props) {
    return (
        <button
            className="square"
            onClick={props.onClick}
        >
            {props.value}
        </button>
    );
}

class Board extends React.Component {
    renderSquare(i) {
        return <Square
            value={this.props.squares[i]}
            onClick={() => this.props.onClick(i)}
        />;
    }

    render() {
        return (
            <div>
                <div className="board-row">
                    {this.renderSquare(0)}
                    {this.renderSquare(1)}
                    {this.renderSquare(2)}
                </div>
                <div className="board-row">
                    {this.renderSquare(3)}
                    {this.renderSquare(4)}
                    {this.renderSquare(5)}
                </div>
                <div className="board-row">
                    {this.renderSquare(6)}
                    {this.renderSquare(7)}
                    {this.renderSquare(8)}
                </div>
            </div>
        );
    }
}

class Game extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            history: [{
                squares: Array(9).fill(null)
            }],
            xIsNext: true,
            stepNumber: 0
        }
    }
    jumpTo(step) {
        let tmpState = {
            stepNumber: step,
            xIsNext: (step %2) === 0
        }

        if (step === 0) {
            tmpState.history = [{
                squares: Array(9).fill(null)
            }]
        }

        this.setState(tmpState)
    }
    handleClick(i) {
        const history = this.state.history.slice(0, this.state.stepNumber + 1)
        const current = history[history.length -1]
        const squares = current.squares.slice()
        if (calculateWinner(squares) || squares[i]) {
            return
        }
        squares[i] = this.state.xIsNext ? 'X':'O'
        this.setState({
            history: history.concat([{
                squares: squares
            }]),
            xIsNext: !this.state.xIsNext,
            stepNumber: history.length
        })
    }
    render() {
        const history = this.state.history
        const current = history[this.state.stepNumber]
        const winner = calculateWinner(current.squares)

        const moves = history.map((step, move) => {
            const desc = move ?
                'Go to move #' + move:
                'Go to game start';
            return (
                <li key={move}>
                    <button onClick={() => this.jumpTo(move)}>{desc}</button>
                </li>
            )
        })

        let status
        if (winner) {
            status = 'Winner: ' + winner
        } else {
            status = "Next player: " + (this.state.xIsNext ? 'X':'O')
        }

        return (
            <div className="game">
                <div className="game-board">
                    <Board
                        squares={current.squares}
                        onClick={(i) => this.handleClick(i)}
                    />
                </div>
                <div className="game-info">
                    <div>{status}</div>
                    <ol>{moves}</ol>
                </div>
            </div>
        );
    }
}

function calculateWinner(squares) {
    const lines = [
        [0, 1, 2],
        [3, 4, 5],
        [6, 7, 8],
        [0, 3, 6],
        [1, 4, 7],
        [2, 5, 8],
        [0, 4, 8],
        [2, 4, 6],
    ];
    for (let i = 0; i < lines.length; i++) {
        const [a, b, c] = lines[i];
        if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
            return squares[a];
        }
    }
    return null;
}

class Toy extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            message: 'message',
            response: null,
            msg: ""
        }
    }
    componentDidMount() {
        const request = {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        }
        fetch('http://localhost:8080/v1/echo', request)
            .then(response => response.json())
            .then(data => this.setState({response: data.message}))
    }
    echo(name) {
        const request = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({name: name})
        }
        fetch('http://localhost:8080/v1/echo', request)
            .then(response => response.json())
            .then(data => this.setState({response: data.message}))
    }
    render(){
        return(
            <div>
                <label>Name:</label>
                <input type="text" name="msg" size="50" onChange={e => this.state.msg = e.target.value} />
                <button onClick={() => this.echo(this.state.msg)}>send</button>
                <h3>{this.state.response}</h3>
            </div>
        )
    }
}

function App() {
  return (
    <div className="App">
      {/*<header className="App-header">*/}
      {/*  <img src={logo} className="App-logo" alt="logo" />*/}
      {/*  <p>Edit <code>src/App.js</code> and save to reload.</p>*/}
      {/*</header>*/}
      {/*  <Game />*/}
      {/*  <hr />*/}
        <Toy />
    </div>
  );
}

export default App;
