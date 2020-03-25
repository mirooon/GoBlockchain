import React from 'react';
import axios from 'axios'
import { Grid, Header, Form, Button } from 'tabler-react'

export class WalletGenerator extends React.Component {
  constructor(props) {
    super(props);
    this.generate = this.generate.bind(this);
  }

  state = {
    privateKey: null,
    publicKey: null
  };
  generate() {
    axios.get('http://localhost:8080/wallet/generate')
      .then((response) => {
        this.setState({
          privateKey: response.data.privateKey,
          publicKey: response.data.publicKey
        });
      }
      )
  }

  render() {
    return (
      <div className="container-fluid">
        <Header.H1>Wallet Generator</Header.H1>
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
            <Form.Label>Public Key</Form.Label>
            <Form.Textarea value={this.state.publicKey} className="container-fluid" readOnly={true}/>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
        <br></br>
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
            <Form.Label>Private Key</Form.Label>
            <Form.Textarea value={this.state.privateKey} className="container-fluid" readOnly={true}/>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
        <br></br>
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
            <Button onClick={this.generate} color="primary" size="lg">Generate</Button>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
      </div>
    )
  };
};