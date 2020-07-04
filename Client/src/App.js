import React from 'react';
import {Table} from './Components/Table';
import './App.css';
import {Forms} from './Components/Form';

class App extends React.Component {
  render () {
  return (
    <div className="App">
      <Table/>
      <Forms />
    </div>
  );
  }
}

export default App;
