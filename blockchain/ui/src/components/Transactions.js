import React from 'react';
import axios from 'axios'
import { Grid, Header, Button } from 'tabler-react'
import { TransactionsTable } from './TransactionsTable';
import { MinedTransactionsTable } from './MinedTransactionsTable';

export class Transactions extends React.Component {
  constructor(props) {
    super(props);
    this.getTransactions = this.getTransactions.bind(this);
    this.mineBlock = this.mineBlock.bind(this);
    this.getChain = this.getChain.bind(this);
  }

  state = {
    "transactions": [],
    "minedTransactions": [],
    "blocks": []
  }
  
  componentDidMount(){
    this.getTransactions();
    this.getChain();
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
    axios.post('http://localhost:5001/mine')
      .then((response) => {
        this.getTransactions();
        this.getChain();
      }
      )
  }

  getChain() {
    console.log("getChain");
    axios.get('http://localhost:5001/chain')
      .then((response) => {
        const blocksUpdate = response.data.Chain;
        let minedTransactionsUpdate = []
        blocksUpdate.forEach((block,blockIndex) => {
          if(block.Transactions != null && block.Transactions.length > 0){
          block.Transactions.forEach(transaction => {
            let transactionUpdate = {...transaction, Timestamp: block.Timestamp, BlockNumber: blockIndex+1}
            minedTransactionsUpdate.push(transactionUpdate);
          })
        }
        });
        console.log('minedTransactionsUpdate')
        console.log(minedTransactionsUpdate)
        this.setState({ 
            blocks: blocksUpdate,
            minedTransactions: minedTransactionsUpdate.flat(1),
          }, () => {
          console.log("blocks downloaded")
          console.log("mined transactions")
          console.log(this.state.minedTransactions)
        });
      }
      )
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
          <TransactionsTable transactions={this.state.transactions}></TransactionsTable>
                      </Grid.Col>
                      <Grid.Col md={1}>
                      </Grid.Col>
                    </Grid.Row>
                    
                  <Grid.Row cards deck>
                  <Grid.Col md={5}>
                  </Grid.Col>
                  <Grid.Col md={2}>
                  <Button color="primary" onClick={this.mineBlock}>Mine</Button>
                  </Grid.Col>
                  <Grid.Col md={5}>
                  </Grid.Col>
                  </Grid.Row>
                  <br></br>
                  <Header.H1>Mined transactions</Header.H1>
        <p><Button color="primary" onClick={this.getChain}>
    Refresh
  </Button></p>
        
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