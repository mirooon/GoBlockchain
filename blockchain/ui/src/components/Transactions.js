import React from 'react';
import axios from 'axios'
import { Grid, Header, Button, Table } from 'tabler-react'

export class Transactions extends React.Component {
  constructor(props) {
    super(props);
    this.getTransactions = this.getTransactions.bind(this);
  }

  state = {
    "transactions": [],
  }
  getTransactions() {
    axios.get('http://localhost:5001/transactions')
      .then((response) => {
        console.log(response.data);
        this.setState({ transactions: response.data}, () => {console.log("transactions downloaded")});
        // this.setState({
        //   privateKey: response.data.privateKey,
        //   publicKey: response.data.publicKey
        // }, function () {
        //   console.log(this.state.privateKey);
        // });
      }
      )
  }

  mineBlock() {
    console.log("mine");
    // axios.get('http://localhost:5001/transactions')
    //   .then((response) => {
    //     this.setState({ transactions: response.data}, () => {console.log("transactions downloaded")});
    //     // this.setState({
    //     //   privateKey: response.data.privateKey,
    //     //   publicKey: response.data.publicKey
    //     // }, function () {
    //     //   console.log(this.state.privateKey);
    //     // });
    //   }
    //   )
  }

    render() {
    return (
      <div className="container-fluid">
        <Header.H1>Transactions</Header.H1>
        <p><Button color="primary" onClick={this.getTransactions}>
    Refresh
  </Button></p>
        
        <Grid.Row cards deck>
          <Grid.Col md={1}>
          </Grid.Col>
          <Grid.Col md={10}>
          <Table>
  <Table.Header>
    <Table.ColHeader>#</Table.ColHeader>
    <Table.ColHeader>Sender Public Key</Table.ColHeader>
    <Table.ColHeader>Recipient Public Key</Table.ColHeader>
    <Table.ColHeader>Amount</Table.ColHeader>
  </Table.Header>
  <Table.Body>
  {this.state.transactions != null && this.state.transactions.length > 0 ? (
    <Table.Row>
    {this.state.transactions.map(function(transaction, index){
                    return (<><Table.Col>{index+1}</Table.Col><Table.Col>{transaction.SenderPublicKey.substring(0, 10)}...{transaction.SenderPublicKey.substring(transaction.SenderPublicKey.length-11, transaction.SenderPublicKey.length-1)}</Table.Col><Table.Col>{transaction.RecipientPublicKey.substring(0, 10)}...{transaction.RecipientPublicKey.substring(transaction.RecipientPublicKey.length-11, transaction.RecipientPublicKey.length-1)}</Table.Col><Table.Col>{transaction.Amount}</Table.Col><Table.Col>
                    <Button color="primary">Edit</Button>
                  </Table.Col></>
                  )}
                  )
                }
                </Table.Row>
                 ) : (
                   <><Table.Col>No transactions</Table.Col></>
                 )}
              
              </Table.Body>
            </Table>
                      </Grid.Col>
                      <Grid.Col md={1}>
                      </Grid.Col>
                    </Grid.Row>
                    
                  <Grid.Row cards deck>
                  <Grid.Col md={5}>
                  </Grid.Col>
                  <Grid.Col md={2}>
                  <Button color="primary">Mine</Button>
                  </Grid.Col>
                  <Grid.Col md={5}>
                  </Grid.Col>
                  </Grid.Row>
                  </div>
                  )
}
};