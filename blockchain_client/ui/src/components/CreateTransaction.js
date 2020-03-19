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
    console.log("1")

  }

  state = {
    show: false,
    senderPublicKey: "04ef4e7be583c78ed83b46fe2556028ca5832eea2870966896cef7e3c2c9735ee9d4d49a21ec07cf002a80bd7576feab342c7d5d2677047f595ad512d206e94cee",
    senderPrivateKey: "54335069774dc6fbee81f0628e32c6323ee09c1170cdaec0cb64fc0bc37899bb",
    recipientPublicKey: "04afc823f802b407cecc32fd2d4978b8409868bb55a01a67019f84f0fd83fcd2d3e6bfc9c479d00c6adab53a733985f6a2d1e986651cb37a358570331113e6e8ee",
    amount: "45",
    signature: null,
  };

  handleChange (evt) {
    this.setState({ [evt.target.name]: evt.target.value });
  }

  childHandler(showModal) {
    this.setState({
        show: false
    },() => {});
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