import React from 'react';
import './App.css';
import "tabler-react/dist/Tabler.css";
import {TabbedCard, Tab, Grid, Card, Header} from 'tabler-react'
import {WalletGenerator} from './components/WalletGenerator'

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
  <Tab title="Make Transaction">
    <Header.H1>Make Transaction</Header.H1>
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
