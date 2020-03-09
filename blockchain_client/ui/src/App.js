import React from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import "tabler-react/dist/Tabler.css";
import {TabbedCard, Tab, Grid, Card} from 'tabler-react'
import {WalletGenerator} from './components/WalletGenerator'
import {CreateTransaction} from './components/CreateTransaction'

function App() {

  return (
    <div className="App container-fluid">
      <Card>
  <Card.Header>
    <Card.Title>Go Blockchain Wallet</Card.Title>
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
    <Grid.Row cards deck>
      <Grid.Col md={4}>
        <Card body="Short content" />
      </Grid.Col>
      <Grid.Col md={4}>
        <Card body="Extra long content of card. Lorem ipsum dolor sit amet, consetetur sadipscing elitr" />
      </Grid.Col>
      <Grid.Col md={4}>
        <Card body="Short content" />
      </Grid.Col>
    </Grid.Row>
  </Tab>
</TabbedCard>
    </div>
  );
}

export default App;
