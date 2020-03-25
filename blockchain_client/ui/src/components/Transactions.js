import React from 'react';
import axios from 'axios'
import { Grid, Header, Button, Form, Alert } from 'tabler-react'
import { MinedTransactionsTable } from './MinedTransactionsTable';

export class Transactions extends React.Component {
  constructor(props) {
    super(props);
    this.getChain = this.getChain.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  state = {
    "errorMessage": null,
    "node": "",
    "minedTransactions": [],
    "blocks": []
  }

  handleChange (evt) {
    this.setState({ [evt.target.name]: evt.target.value });
  }

  getChain() {
    this.setState({ errorMessage: null});
    if(this.state.node !== ""){
    axios.get('http://' + this.state.node +'/chain')
      .then((response) => {
        const blocksUpdate = response.data.Chain;
        if(blocksUpdate != null){
          let minedTransactionsUpdate = []
          blocksUpdate.forEach((block,blockIndex) => {
            if(block.Transactions != null && block.Transactions.length > 0){
            block.Transactions.forEach(transaction => {
              let transactionUpdate = {...transaction, Timestamp: block.Timestamp, BlockNumber: blockIndex+1}
              minedTransactionsUpdate.push(transactionUpdate);
            })
          }
          });
          this.setState({ 
              blocks: blocksUpdate,
              minedTransactions: minedTransactionsUpdate.flat(1),
            });
        }
      }
      ).catch((err) => {    this.setState({ errorMessage: "Error with connection :/"});    })
    }
  }

  ValidateIPaddress(ipaddress) {
 if (/^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?):[0-9]{1,5}$/.test(ipaddress))
  {
    return (true)
  }
return (false)
}

    render() {
    return (
      <div className="container-fluid">
        <Header.H1>Transactions</Header.H1>
        {!this.ValidateIPaddress(this.state.node) &&         <p><Alert type="danger">
  <strong>Provide a node address with port!</strong> (example: 127.0.0.1:5001)
</Alert></p>}
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
          <Form.Group>
  <Form.Input name="node" placeholder="Node address" onChange={this.handleChange}/>
</Form.Group>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
        <p><Button color="primary" onClick={this.getChain}>
    View transactions
  </Button></p>
  {this.state.errorMessage != null &&         <p><Alert type="danger">
    <strong>{this.state.errorMessage}</strong>
</Alert></p>}
        <Grid.Row cards deck>
          <Grid.Col md={1}>
          </Grid.Col>
          <Grid.Col md={10}>
          <MinedTransactionsTable transactions={this.state.minedTransactions}></MinedTransactionsTable>
                      </Grid.Col>
                      <Grid.Col md={1}>
                      </Grid.Col>
                    </Grid.Row>
                  </div>
                  )
}
};