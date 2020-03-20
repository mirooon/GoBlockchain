import React from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import "tabler-react/dist/Tabler.css";
import {TabbedCard, Tab, Grid, Card} from 'tabler-react'
import {WalletGenerator} from './components/WalletGenerator'
import {CreateTransaction} from './components/CreateTransaction'
import { Transactions } from './components/Transactions';

function App() {

  return (
    <div className="App container-fluid">
      <Card>
  <Card.Header>
    <Card.Title>Go Blockchain Client</Card.Title>
  </Card.Header>
</Card>
      <TabbedCard initialTab="Wallet Generator">
  <Tab title="Wallet Generator">
  <WalletGenerator/>
  </Tab>
  <Tab title="Create Transaction">
    <CreateTransaction/>
  </Tab>
  <Tab title="Transactions">
  <Transactions/>
  </Tab>
</TabbedCard>
    </div>
  );
}

export default App;
