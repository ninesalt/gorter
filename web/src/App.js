import React, { Component } from 'react';
import { FormGroup } from 'react-bootstrap'
import './App.css';

class App extends Component {

  constructor(props) {
    super(props)
    this.state = { url: '' }
  }

  handleChange = (e) => {
    this.setState({ url: e.target.value })
  }

  render() {
    return (
      <div className="container">

        <a href="https://github.com/ninesalt/gorter">

          <img className="ribbon"
            src="https://s3.amazonaws.com/github/ribbons/forkme_right_orange_ff7600.png"
            alt="Fork me on GitHub" />

        </a>

        <h1>Gorter</h1>

        <FormGroup>

          <input type="text" placeholder="Enter a URL" />

          <button className="button">Shorten</button>


        </FormGroup>


      </div>
    );
  }
}

export default App;
