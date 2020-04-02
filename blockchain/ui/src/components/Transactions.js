import React from 'react';
import axios from 'axios'
import { Grid, Header, Button } from 'tabler-react'
import { TransactionsTable } from './TransactionsTable';
import { MinedTransactionsTable } from './MinedTransactionsTable';
import { config } from '../config/config'

export class Transactions extends React.Component {
  constructor(props) {
    super(props);
    this.getTransactions = this.getTransactions.bind(this);
    this.mineBlock = this.mineBlock.bind(this);
    this.resolveConflictsAndGetChain = this.resolveConflictsAndGetChain.bind(this);
  }

  state = {
    "transactions": [],
    "minedTransactions": [],
    "blocks": []
  }
  
  componentDidMount(){
    this.getTransactions();
    this.resolveConflictsAndGetChain();
  }

  getTransactions() {
    axios.get('http://' + config.REACT_APP_HOSTNODEIP + '/transactions')
      .then((response) => {
        this.setState({ transactions: response.data});
      }
      )
  }

  mineBlock() {
    axios.post('http://' + config.REACT_APP_HOSTNODEIP + '/mine')
      .then(() => {
        this.getTransactions();
        this.resolveConflictsAndGetChain();
      }
      )
  }

  resolveConflictsAndGetChain() {
    axios.get('http://' + config.REACT_APP_HOSTNODEIP + '/nodes/resolve')
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
            }, () => {
          });
        }
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
        <p><Button color="primary" onClick={this.resolveConflictsAndGetChain}>
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