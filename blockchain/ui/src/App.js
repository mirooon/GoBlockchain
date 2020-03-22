import React from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import "tabler-react/dist/Tabler.css";
import {TabbedCard, Tab, Card} from 'tabler-react'
import { Transactions } from './components/Transactions';
import { Neighbours } from './components/Neighbours';

function App() {
  return (
    <div className="App container-fluid">
      <Card>
  <Card.Header>
    <Card.Title>Go Blockchain Node</Card.Title> <p>IP: {process.env.NODE_IP}</p>
  </Card.Header>
</Card>
      <TabbedCard initialTab="Transactions">
  <Tab title="Transactions">
    <Transactions/>
  </Tab>
  <Tab title="Neighbours">
    <Neighbours/>
  </Tab>
</TabbedCard>
    </div>
  );
}

export default App;
