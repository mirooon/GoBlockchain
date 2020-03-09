import React from 'react';
import axios from 'axios'
import { Grid, Header, Form, Button } from 'tabler-react'
import {CustomModal} from './Modal'

export class CreateTransaction extends React.Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
    this.createTransaction = this.createTransaction.bind(this);
    this.childHandler = this.childHandler.bind(this);
  }

  state = {
    show: false,
    senderPublicKey: null,
    senderPrivateKey: null,
    recipientPublicKey: null,
    amount: null,
    signature: null,
  };

  handleChange (evt) {
    this.setState({ [evt.target.name]: evt.target.value });
  }

  childHandler(showModal) {
    // log our state before and after we updated it
    console.log('%cPrevious Parent State: ' + JSON.stringify(this.state), "color:orange");
    this.setState({
        show: false
    },() => console.log('Updated Parent State:', this.state));
}

  createTransaction() {
    axios.post('http://localhost:8080/transaction/create', 
    {
      "senderPublicKey":this.state.senderPublicKey,
      "senderPrivateKey":this.state.senderPrivateKey,
      "recipientPublicKey":this.state.recipientPublicKey,
      "amount":this.state.amount,
    }
      )
      .then((response) => {
        console.log(response.data.Signature);
        this.setState({ show: true, signature: response.data.Signature}, () => {console.log(this.state)});
        // this.setState({
        //   privateKey: response.data.privateKey,
        //   publicKey: response.data.publicKey
        // }, function () {
        //   console.log(this.state.privateKey);
        // });
      }
      )
  }

  render() {
    return (
      <div className="container-fluid" >
        <CustomModal object={this.state} action={this.childHandler}/>
        <Header.H1>Create Transaction</Header.H1>
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
          <Form.Group>
  <Form.Input name="senderPublicKey" placeholder="Sender Public Key" onChange={this.handleChange}/>
</Form.Group>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
         
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
          <Form.Group>
  <Form.Input name="senderPrivateKey" placeholder="Sender Private Key" onChange={this.handleChange}/>
</Form.Group>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
         
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
          <Form.Group>
  <Form.Input name="recipientPublicKey" placeholder="Recipient Public Key" onChange={this.handleChange}/>
</Form.Group>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
          <Form.Group>
  <Form.Input name="amount" placeholder="Amount" onChange={this.handleChange}/>
</Form.Group>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
            <Button onClick={this.createTransaction} color="primary" size="lg">Send coins</Button>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
      </div>
    )
  };
};