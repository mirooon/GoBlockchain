import React from 'react';
import { Table } from 'tabler-react'

export class TransactionsTable extends React.Component {
//   constructor(props) {
//     super(props);

//     // this.getTransactions = this.getTransactions.bind(this);
//   }

static getDerivedStateFromProps(props, state) {
    if (props.transactions !== state.transactions) {
      return {
        transactions: props.transactions,
      };
    }
    // Return null to indicate no change to state.
    return null;
  }
  state = {
    "transactions": this.props.transactions,
  }
    render() {
    return (<Table>
        <Table.Header>
          <Table.ColHeader>#</Table.ColHeader>
          <Table.ColHeader>Sender Public Key</Table.ColHeader>
          <Table.ColHeader>Recipient Public Key</Table.ColHeader>
          <Table.ColHeader>Amount</Table.ColHeader>
        </Table.Header>
        <Table.Body>
        {this.state.transactions != null && this.state.transactions.length > 0 ? (
            <>
            {this.state.transactions.map(function(transaction, index){
              
                          return (<><Table.Row>
                              <Table.Col>{index+1}</Table.Col>
                          <Table.Col>{transaction.SenderPublicKey.substring(0, 10)}...{transaction.SenderPublicKey.substring(transaction.SenderPublicKey.length-11, transaction.SenderPublicKey.length-1)}</Table.Col>
                          <Table.Col>{transaction.RecipientPublicKey.substring(0, 10)}...{transaction.RecipientPublicKey.substring(transaction.RecipientPublicKey.length-11, transaction.RecipientPublicKey.length-1)}</Table.Col>
                          <Table.Col>{transaction.Amount}</Table.Col>
                          </Table.Row></>
                        )}
                        )
                    }
                    </>
                       ) : (
                         <><Table.Col>No transactions</Table.Col></>
                       )}
                    
                    </Table.Body>
                  </Table>);
    }
};